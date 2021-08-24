package main

import (
	"fmt"
	"net/http"
)

func main() {
	handler()
}

type Video struct {
	Video      string `json:"Video"`
	Url_videos Url    `json:"UrlVideos"`
}

type Url struct {
	Url string `json:"Url"`
	Img string `json:"Img"`
}

func handler() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("HOME DE LA API XVIDEO ILEGAL"))
	})
	http.HandleFunc("/api", ReceiveData)

	http.ListenAndServe(":8000", nil)

}

func Img(url string) []string {

	var v []string

	if len(url) < 6 {
		return v
	}

	if url[0:6] != "/video" {
		return v
	}
	//video46174535/hermana_b._pide_masajes_a_su_hermano

	fmt.Println(url)
	resp, err := http.Get(string("https://xvideos.com" + url))

	if err != nil {
		fmt.Printf("Error al recibir la %s", url)
	}

	v = getImg(resp.Body)
	return v
}
