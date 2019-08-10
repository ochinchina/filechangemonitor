package filechangemonitor

import (
	"fmt"
)

type FileChangeCallback func(path string, mode FileChangeMode)

func PrintFileChangeCallback(path string, mode FileChangeMode) {
	switch mode {
	case Create:
		fmt.Printf("%s is created\n", path)
	case Delete:
		fmt.Printf("%s is deleted\n", path)
	case Modify:
		fmt.Printf("%s is changed\n", path)
	}
}
