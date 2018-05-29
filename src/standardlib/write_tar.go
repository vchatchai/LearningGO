package standardlib

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
)

func addFile(filename string, tw *tar.Writer) (err error) {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed opening %s: %s", filename, err)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return fmt.Errorf("fail file stat for %s: %s", filename, err)
	}

	hdr := &tar.Header{
		ModTime: stat.ModTime(),
		Name:    filename,
		Size:    stat.Size(),
		Mode:    int64(stat.Mode().Perm()),
	}

	err = tw.WriteHeader(hdr)

	if err != nil {
		return fmt.Errorf("write header fail %s:%s", filename, err)
	}

	copied, err := io.Copy(tw, file)
	if err != nil {
		return fmt.Errorf("Copy error %s: %s", filename, err)
	}

	if copied < stat.Size() {
		return fmt.Errorf("wrote %d bytes of %s, expected to write %d", copied, filename, stat.Size())
	}

	return
}
