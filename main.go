package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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

var Field_json []Video

func handler() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		url, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Error(w, "Ooops Error in get url", http.StatusBadRequest)
		}
		fmt.Fprintf(w, "Url: %s\n", url)

		resp, err := http.Get(string(url))
		if err != nil {
			log.Fatal(err)
		}

		v := getLinks(resp.Body)
		var x, y string

		videos := Video{}

		for _, v1 := range v {
			//Para saber la url del video presente y donde encontralo
			if strings.Contains(v1, ".mp4") {
				fmt.Fprintf(w, "Video: %s\n\n", v1)
				x = v1
			}

			//Para saber las imagenes de las listas de videos presentes
			d := Img(v1)

			if len(d) != 0 {
				fmt.Fprintf(w, "%s\n\n", v1)
				for _, v2 := range d {
					if strings.Contains(v2, ".jpg") {
						y = v2
					}
				}
			}
			u_v := Url{
				Url: v1,
				Img: y,
			}

			videos = Video{
				Video:      x,
				Url_videos: u_v,
			}

			Field_json = append(Field_json, videos)

			fmt.Println(Field_json)
		}
	})

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

	resp, err := http.Get(string("https://xvideos.com" + url))

	if err != nil {
		fmt.Println("Error al recibir la url")
	}

	v = getImg(resp.Body)
	return v
}
