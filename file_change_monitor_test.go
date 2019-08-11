package filechangemonitor

import (
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func testNoFileChanged(t *testing.T, fileChangeMonitor *FileChangeMonitor) {
	dir, err := ioutil.TempDir("", "example")
	if err != nil {
		t.Error("fail to create temp directory for testing")
	}
	defer os.RemoveAll(dir) // clean up
	_, err = ioutil.TempFile(dir, "*")
	if err != nil {
		t.Error("fail to create temp file")
	}
	changedFiles := 0
	fileChangeCb := func(path string, mode FileChangeMode) {
		changedFiles += 1
	}

	fileChangeMonitor.AddMonitorFile(dir,
		true,
		NewMatchAllFile(),
		NewFileChangeCallbackWrapper(fileChangeCb),
		NewFileModTimeCompareInfo())
	time.Sleep(3 * time.Second)
	fileChangeMonitor.Stop()
	fileChangeMonitor.Wait()
	if changedFiles != 0 {
		t.Error("No file should be changed")
	}
}

func testFileChanged(t *testing.T, fileChangeMonitor *FileChangeMonitor) {
	dir, err := ioutil.TempDir("", "example")
	if err != nil {
		t.Error("fail to create temp directory for testing")
	}
	defer os.RemoveAll(dir) // clean up
	_, err = ioutil.TempFile(dir, "*")
	if err != nil {
		t.Error("fail to create temp file")
	}
	changedFiles := 0
	fileChangeCb := func(path string, mode FileChangeMode) {
		changedFiles += 1
	}
	fileChangeMonitor.AddMonitorFile(dir,
		true,
		NewMatchAllFile(),
		NewChainedFileChangeCallback(NewFileChangeCallbackWrapper(fileChangeCb), NewPrintFileChangeCallback()),
		NewFileModTimeCompareInfo())
	_, err = ioutil.TempFile(dir, "*")
	time.Sleep(3 * time.Second)
	fileChangeMonitor.Stop()
	fileChangeMonitor.Wait()
	if changedFiles == 0 {
		t.Error("there should be some files changed")
	}
}

func TestNoFileChanged(t *testing.T) {
	testNoFileChanged(t, NewFileChangeMonitor(1))
}

func TestFileChanged(t *testing.T) {
	testFileChanged(t, NewFileChangeMonitor(1))
}
