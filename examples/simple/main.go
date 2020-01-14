package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/JakubOboza/copycat"
)

const (
	URL      = "https://upload.wikimedia.org/wikipedia/commons/thumb/1/1b/El_sue%C3%B1o_de_Jacob%2C_by_Jos%C3%A9_de_Ribera%2C_from_Prado_in_Google_Earth-x1-y1.jpg/800px-El_sue%C3%B1o_de_Jacob%2C_by_Jos%C3%A9_de_Ribera%2C_from_Prado_in_Google_Earth-x1-y1.jpg"
	FILENAME = "jacobs_dream.jpg"
)

type Observer struct {
	bytesWritten int
	fileSize     int64
}

func (o *Observer) ProgressUpdate(progress int) {
	o.bytesWritten += progress
	percent := (float64(o.bytesWritten) / float64(o.fileSize)) * 100
	fmt.Printf("So far written %d bytes (%3.2f %% of total) done\n", o.bytesWritten, percent)
}

func main() {

	out, err := os.Create(FILENAME)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(URL)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	pm := copycat.NewProgressReader(resp.Body)

	obs := &Observer{fileSize: resp.ContentLength}

	pm.AddListener(obs)

	// Write the body to file but use wrapped progress manager as reader
	_, err = io.Copy(out, pm)
	if err != nil {
		fmt.Println(err)
		return
	}

}
