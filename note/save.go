package note

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func Save(output *bytes.Buffer, path string) {
	fmt.Fprintf(output, "_last updated @ %s_ \n\n", time.Now().Format("2006-01-02 15:04"))

	file, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	defer file.Close()
	if err != nil {
		log.Println("got error", err)
		return
	}

	if _, err := io.Copy(file, output); err != nil {
		log.Println("got error", err)
		return
	}

	log.Println("saved.")
}
