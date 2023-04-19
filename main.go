package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gonutz/w32/v2"
)

func main() {
	desktop, ok := w32.SHGetSpecialFolderPath(w32.HWND_DESKTOP, w32.CSIDL_DESKTOP, false)
	if !ok {
		log.Fatalf("Could not lookup desktop folder DANG!!!!")
	}
	for i := 1; i <= 1000; i++ {
		err := os.Mkdir(filepath.Join(desktop, fmt.Sprintf("COOL FOLDER #%d", i)), os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}
