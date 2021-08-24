package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func ReceiveData(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Datos de la API"))
	var Field_json []Video
	video := r.FormValue("video")

	if video[0:1] == "/" {
		video = video[1:]
	}

	//url, err := ioutil.ReadAll(r.Body)
	/*
		if err != nil {
			http.Error(w, "Ooops Error in get url", http.StatusBadRequest)
		}
	*/
	url := "https://www.xvideos.com/"

	fmt.Fprintf(w, "Url: %s%s\n", url, video)

	resp, err := http.Get(url + video)
	if err != nil {
		log.Fatal(err)
	}

	v := getLinks(resp.Body)
	var x, y string

	videos := Video{}

	for _, v1 := range v {
		//Para saber la url del video presente y donde encontralo
		if strings.Contains(v1, ".mp4") {
			//fmt.Fprintf(w, "Video: %s\n\n", v1)
			x = v1
		}

		//Para saber las imagenes de las listas de videos presentes
		d := Img(v1)

		if len(d) != 0 {
			//fmt.Fprintf(w, "%s\n\n", v1)
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

	}

	dJson, err := json.Marshal(Field_json)

	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "Archivo json: %s", string(dJson))
}
