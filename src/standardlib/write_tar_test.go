package standardlib

import (
	"archive/tar"
	"log"
	"os"
	"testing"
)

func Test_addFile(t *testing.T) {
	type args struct {
		filename string
		flags    int
		files    []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test add files", args{filename: "go.tar", flags: os.O_WRONLY | os.O_CREATE | os.O_TRUNC, files: []string{"write_tar.go", "write_tar_test.go"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			file, err := os.OpenFile(tt.args.filename, tt.args.flags, 0644)
			if err != nil {
				log.Fatalf("failed opening tar for writing: %s", err)
			}

			tw := tar.NewWriter(file)

			defer tw.Close()
			for _, file := range tt.args.files {

				if err := addFile(file, tw); (err != nil) != tt.wantErr {
					t.Errorf("addFile() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}
