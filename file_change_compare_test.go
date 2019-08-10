package filechangemonitor

import (
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func testFileIsNotChanged(t *testing.T, fileChangeCompare *FileChangeCompare) {
	f, err := ioutil.TempFile("", "*")
	if err != nil {
		t.Error("fail to create temp file")
		return
	}
	f.Write([]byte("first line"))
	defer os.Remove(f.Name())
	fileChangeCompare.SetInitialFiles([]string{f.Name()})
	time.Sleep(1 * time.Second)
	changeEvent := fileChangeCompare.UpdateFiles([]string{f.Name()})
	if len(changeEvent) != 0 {
		t.Error("file is changed")
	}
}

func testFileIsChanged(t *testing.T, fileChangeCompare *FileChangeCompare) {
	f, err := ioutil.TempFile("", "*")
	if err != nil {
		t.Error("fail to create temp file")
		return
	}
	defer os.Remove(f.Name())
	f.Write([]byte("first line"))
	fileChangeCompare.SetInitialFiles([]string{f.Name()})
	time.Sleep(1 * time.Second)
	f.Write([]byte("change it"))
	changeEvent := fileChangeCompare.UpdateFiles([]string{f.Name()})
	if len(changeEvent) != 1 {
		t.Error("file is not changed")
	}

}

func TestFileTimeIsSame(t *testing.T) {
	testFileIsNotChanged(t, NewFileTimeChangeCompare())
}

func TestFileMD5IsSame(t *testing.T) {
	testFileIsNotChanged(t, NewFileMD5ChangeCompare())
}

func TestFileTimeIsNotSame(t *testing.T) {
	testFileIsChanged(t, NewFileTimeChangeCompare())
}

func TestFileMD5IsNotSame(t *testing.T) {
	testFileIsChanged(t, NewFileMD5ChangeCompare())
}
