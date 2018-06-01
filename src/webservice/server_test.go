package webservice

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"storage"
	"testing"
)

func Test_handleGET(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Normal Case GET", args: args{}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mux := http.NewServeMux()
			mux.HandleFunc("/post/", crudHandler)

			writer := httptest.NewRecorder()
			request, _ := http.NewRequest("GET", "/post/1", nil)
			mux.ServeHTTP(writer, request)

			if writer.Code != 200 {
				t.Errorf("Response code is %v", writer.Code)
			}

			topic := storage.Topic{}
			t.Log("Before Unmarshal", string(writer.Body.Bytes()))
			err := json.Unmarshal(writer.Body.Bytes(), &topic)

			if err != nil {
				t.Errorf(" Unmarshal is error: %v, %v", err, topic)
			}

			// if strings.EqualFold( post.Id ,"1") {
			if topic.Id != 2 {
				t.Errorf("Json Post is wrong id = %v , expected is %v", topic.Id, 2)
			}

		})
	}
}

func Test_handlePost(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Test Posting", args{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Error( tt.name)
			mux := http.NewServeMux()

			http.HandleFunc("/post", crudHandler)


			writer := httptest.NewRecorder()
			request, _ := http.NewRequest("POST", "/post", nil)

			mux.ServeHTTP(writer,request)

			if writer.Code != 200 {
				t.Errorf("Reponse code wrong %d expected 200", writer.Code )
			}
 
		})
	}
}
