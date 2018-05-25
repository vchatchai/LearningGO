package main

import (
	"fmt"
	"math"
	"net/http" 
	"webservice"
	"time"
)

// var plus = add

// func add(x int, y int) int {
// 	return x + y
// }

// func init() {
// 	fmt.Println("Initial1")
// }

func main() {

	fmt.Println("Start Testing")
	// testFor()
	// testForNoInitial()

	// s := []int{2, 3, 5, 7, 11, 13}

	// s = s[1:4]
	// fmt.Println(s)

	// s = s[:2]
	// fmt.Println(s)

	// s = s[1:]
	// fmt.Println(s)

	// s = s[2:1]
	// fmt.Println(s)

	// pic.Show(Pic)
	// testMap()
	// echoServer()
	// deferFunction()
	// runGoRoutine()
	// channel()

	// var c Circle
	// c = new(Circle)

	// c := Circle{x: 0, y: 0, r: 5}

	// fmt.Println(c)
	// fmt.Println(c.r)
	// fmt.Println(c.area())
	// MaxParallelism()

	// c2 := circle{radius: 10}
	// r2 := rect{height: 10, width: 10}
	// measure(c2)
	// measure(r2)

	// word := "Chatchai Vichai"

	// fmt.Println(strings.Contains("test", "es"))
	// fmt.Println(strings.Count(strings.ToLower(word), "c"))
	// fmt.Println(strings.Join(strings.Split(word, ""), "-"))

	// arr := []byte("atest")
	// str := string([]byte{98, 't', 'e', 's', 't'})
	// fmt.Println(arr)
	// fmt.Println(str)

	// file, err := os.Open("test.txt")
	// if err != nil {

	// 	file, err := os.Create("test.txt")

	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}

	// 	file.WriteString("test for golang")

	// }
	// defer file.Close()

	// //get the file size
	// stat, err := file.Stat()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// //read the file
	// bs := make([]byte, stat.Size())
	// _, err = file.Read(bs)
	// if err != nil {
	// 	return
	// }

	// str = string(bs)
	// fmt.Println(str)

	// bs, err = ioutil.ReadFile("test.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// str = string(bs)
	// fmt.Println(str)

	// dir, err := os.Open(".")
	// if err != nil {
	// 	return
	// }

	// defer dir.Close()

	// fileInfos, err := dir.Readdir(-1)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// for _, fi := range fileInfos {
	// 	fmt.Println(fi.Name(), fi.Mode(), fi.IsDir())
	// }

	// filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
	// 	fmt.Println(path)
	// 	return nil
	// })

	// h := crc32.NewIEEE()
	// h.Write([]byte("test"))
	// v := h.Sum32()
	// fmt.Println(v)

	// runcrc32()
	// runHash()
	// runSHA1()

	// // serverRun()
	// // clientRun()
	// // httpRun();
	// go rpcserver()
	// go rpcclient()

	// var input string
	// fmt.Scanln(&input)

	// go f(0)

	// for i := 0; i < 10; i++ {
	// 	routine.F2(i)
	// }

	// routine.Channel()

	// var input string
	// fmt.Scanln(&input)

	// http.HandleFunc("/test", TestHandler)
	// http.HandleFunc("/static", serveStatic)
	// http.HandleFunc("/", serveDynamic)
	// http.FileServer(http.Dir("D:\\"))
	// http.ListenAndServe(PORT, nil)
	// fmt.Println("Everything is set up")

	// web.GenerateSSL()
	// web.ListenAndServeTLS()
	// web.Httprouter()
	// web.Uploading()
	// web.ResponseWriter()

	// Compute()
	// Assertions()
	// Empty()
	// Nil()
	// InterfaceValues()
	// TypeSwitch()
	// Stringers();
	// IPAddrStringer()
	// Errors2()
	// Readers()
	// Rot13Reader()
	// Images()
	// routine.Saying()
	// routine.Channels()
	// routine.BinaryTree()
	// routine.Mutex()
	// web.Template()
	// storage.Gob()
	// storage.CRUD()
		// storage.CRUD2()
	// storage.SQLX()
	// storage.GORM()
	webservice.Unmarshal()
	// web.Cookies()

	// status, err := os.Stat("key.pem")
	// fmt.Println("isExist", status.IsDir(), "err", os.IsExist(err))
}

type Vertex struct {
	Lat, Long float64
}

const (
	PORT = ":8080"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("TestHandler")
}

func serveStatic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static.html")
}

func serveDynamic(w http.ResponseWriter, r *http.Request) {
	response := "The time is now " + time.Now().String()
	fmt.Fprint(w, response)
}

func Pic(dx, dy int) [][]uint8 {
	x := [][]uint8{}
	for i := 0; i < dy; i++ {
		row := []uint8{}
		for j := 0; j < dx; j++ {
			row = append(row, uint8(i^j*2))
		}
		x = append(x, row)
	}
	return x
}

type Circle struct {
	x float64
	y float64
	r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

// Here's a basic interface for geometric shapes.
type geometry interface {
	area() float64
	perim() float64
}

// For our example we'll implement this interface on
// `rect` and `circle` types.
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

// To implement an interface in Go, we just need to
// implement all the methods in the interface. Here we
// implement `geometry` on `rect`s.
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// The implementation for `circle`s.
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// If a variable has an interface type, then we can call
// methods that are in the named interface. Here's a
// generic `measure` function taking advantage of this
// to work on any `geometry`.
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	// fmt.Println(g.perim())
}
