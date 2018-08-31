package web

import (
	"encoding/gob"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type myElement struct {
	Name    string
	Surname string
	Id      string
}

var DATA = make(map[string]myElement)
var DATAFILE = "/tmp/dataFile.gob"

func save() error {
	fmt.Println("Saving", DATAFILE)
	err := os.Remove(DATAFILE)
	if err != nil {
		fmt.Println(err)
		return err
	}

	saveTo, err := os.Create(DATAFILE)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer saveTo.Close()

	encoder := gob.NewEncoder(saveTo)
	err = encoder.Encode(DATA)
	if err != nil {
		fmt.Println("Cannt save to", DATAFILE)
		return err
	}
	return nil
}

func load() error {
	fmt.Println("Loading", DATAFILE)
	loadForm, err := os.Open(DATAFILE)
	defer loadForm.Close()
	if err != nil {
		fmt.Println("Empty key/value store!")
		return err
	}
	decoder := gob.NewDecoder(loadForm)
	decoder.Decode(&DATA)
	return nil
}

func ADD(k string, n myElement) bool {
	if k == "" {
		return false
	}

	if LOOKUP(k) == nil {
		DATA[k] = n
		return true
	}
	return false
}

func DELETE(k string) bool {
	if LOOKUP(k) != nil {
		delete(DATA, k)
		return true
	}
	return false
}

func LOOKUP(k string) *myElement {
	data, ok := DATA[k]
	if ok {
		return &data
	} else {
		return nil
	}
}

func CHANGE(k string, n myElement) bool {
	DATA[k] = n
	return true
}

func PRINT() {
	for k, d := range DATA {
		fmt.Printf("key: %s value: %v\n", k, d)
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Serving", r.Host, "for", r.URL.Path)
	myT := template.Must(template.ParseGlob("home.gohtml"))
	myT.ExecuteTemplate(w, "home.gohtml", nil)

}

func listAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Listing the contents of the KV store!")
	fmt.Fprintf(w, `<a href="/" style="margin-right:20px;"> Home sweet</a>`)
	fmt.Fprintf(w, `<a href="/list" style="margin-right:20px;">List all element!</a>`)
	fmt.Fprintf(w, `<a href="/change" style="margin-right:20px;">Change An Element!</a>`)
	fmt.Fprintf(w, `<a href="/insert" style="margin-right:20px;">Insert new Element!</a>`)
	fmt.Fprintf(w, `<h1>The content of the  KV store are:</h1>`)
	fmt.Fprintf(w, "<ul>")
	for k, v := range DATA {
		fmt.Fprintf(w, "<li>")
		fmt.Fprintf(w, `<strong>%s</strong> with value: %v`, k, v)
		fmt.Fprintf(w, "</li>")
	}
	fmt.Fprintf(w, "</li>")
}

func changeElement(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Changing an element of the KV store")
	tmpl := template.Must(template.ParseFiles("update.gohtml"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
	}
	key := r.FormValue("key")
	n := myElement{
		Name:    r.FormValue("name"),
		Surname: r.FormValue("surname"),
		Id:      r.FormValue("id"),
	}

	if !CHANGE(key, n) {
		fmt.Println("Update operation failed!")
	} else {
		err := save()
		if err != nil {
			fmt.Println(err)
			return
		}
		tmpl.Execute(w, struct{ Success bool }{true})
	}

}

func insertElement(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inserting an element to the KV store")
	tmpl := template.Must(template.ParseFiles("insert.gohtml"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}
	key := r.FormValue("key")
	n := myElement{
		Name:    r.FormValue("name"),
		Surname: r.FormValue("surname"),
		Id:      r.FormValue("id"),
	}

	if !ADD(key, n) {
		fmt.Println("Add Operation failed!")
	} else {
		err := save()
		if err != nil {
			fmt.Println(err)
			return
		}
		tmpl.Execute(w, struct{ Success bool }{true})
	}
}

func KVWeb() {
	err := load()
	if err != nil {
		fmt.Println(err)
	}

	PORT := ":8001"

	http.HandleFunc("/", homePage)
	http.HandleFunc("/change", changeElement)
	http.HandleFunc("/list", listAll)
	http.HandleFunc("/insert", insertElement)
	err = http.ListenAndServe(PORT, nil)
	if err != nil {
		fmt.Println(err)
	}
}
