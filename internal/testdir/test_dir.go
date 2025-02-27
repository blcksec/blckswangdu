package testdir

import (
	"io/ioutil"
	"os"
)

// CreateTestDir creates test dir structure
func CreateTestDir() func() {
	os.MkdirAll("test_dir/nested/subnested", os.ModePerm)
	ioutil.WriteFile("test_dir/nested/subnested/file", []byte("hello"), 0644)
	ioutil.WriteFile("test_dir/nested/file2", []byte("go"), 0644)
	return func() {
		err := os.RemoveAll("test_dir")
		if err != nil {
			panic(err)
		}
	}
}
