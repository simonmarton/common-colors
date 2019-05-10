package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// CommonColorsResp format
type CommonColorsResp struct {
	Colors []ColorResp `json:"colors"`
}

// ColorResp ...
type ColorResp struct {
	Weight int    `json:"weight"`
	Value  string `json:"value"`
}

// APIHandler interface
type APIHandler interface {
	// GetCommonColors(io.Reader) CommonColorsResp
	ProcessImage(file io.Reader, imageType string) (CommonColorsResp, error)
}

// Initialize a new web server
func Initialize(h APIHandler) {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("public/"))))

	http.HandleFunc("/api/upload", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("handle api upload")
		file, header, err := r.FormFile("image")

		if err != nil {
			panic(err)
		}

		fmt.Println(header.Header.Get("Content-Type"))

		colors, err := h.ProcessImage(file, header.Header.Get("Content-Type"))
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json")
		resp, err := json.Marshal(colors)

		if err != nil {
			panic(err)
		}

		w.Write(resp)
	})

	http.ListenAndServe(":8080", nil)
}
