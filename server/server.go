package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/simonmarton/common-colors/models"
)

// CommonColorsResp format
type CommonColorsResp struct {
	Colors []ColorResp `json:"colors"`
}

// ColorResp ...
type ColorResp struct {
	Weight      int     `json:"weight"`
	Value       string  `json:"value"`
	HueDistance float64 `json:"hueDistance"`
}

// APIHandler interface
type APIHandler interface {
	// GetCommonColors(io.Reader) CommonColorsResp
	ProcessImage(file io.Reader, imageType string, config models.CalculatorConfig) (CommonColorsResp, error)
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

		var config models.CalculatorConfig
		err = json.Unmarshal([]byte(r.FormValue("config")), &config)
		if err != nil {
			panic(err)
		}

		colors, err := h.ProcessImage(file, header.Header.Get("Content-Type"), config)
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

	fmt.Println("Ready on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
