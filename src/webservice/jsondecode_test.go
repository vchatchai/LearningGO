package webservice

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestJsonDecoder(t *testing.T) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(dir)
	type args struct {
		filename string
	}
	tests := []struct {
		name     string
		args     args
		wantPost Post
		wantErr  bool
	}{
		// TODO: Add test cases.
		{name: "test convert",
			args: args{filename: "post.json"},
			wantPost: Post{
				Id:      "1",
				Content: "Hello World!",
				Author:  Author{Id: "2", Name: "Sau Sheong"},
				Comments: []Comment{
					Comment{Id: "3", Content: "Have a great day!", Author: Author{Id: "2", Name: "Adam"}},
					Comment{Id: "4", Content: "How are you today?", Author: Author{Id: "2", Name: "Betty"}},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, err := os.Open(tt.args.filename)
			if err != nil {
				t.Errorf("Can't open %s", tt.args.filename)
			}
			defer file.Close()
			gotPost, err := JsonDecoder(file)
			if (err != nil) != tt.wantErr {
				t.Errorf("JsonDecoder() error = %v,\n wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPost, tt.wantPost) {
				t.Errorf("JsonDecoder() = %v, want %v", gotPost, tt.wantPost)
			}
		})
	}
}

func BenchmarkDecode(b *testing.B) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
	file, err := os.Open("/home/chatchai/workspace/LearningGO/src/webservice/post.json")
	if err != nil {
		b.Errorf("Can't open the %s %s", "post.json", err)
	}

	defer file.Close()
	for i := 0; i < b.N; i++ {

		JsonDecoder(file)

	}
}
