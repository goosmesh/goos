package controller

import (
	"fmt"
	"github.com/goosmesh/goos/core/env"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var StaticRoot = env.GoosHome + "/desk"

func HandleStaticResource(w http.ResponseWriter, r *http.Request)  {
	path := r.URL.Path
	if strings.Contains(path, ".") {
		requestType := path[strings.LastIndex(path, ".") : ]

		switch requestType {
		case ".css":
			w.Header().Set("Content-Type", "text/css")
		case ".js":
			w.Header().Set("Content-Type", "text/javascript")
		case ".png":
			w.Header().Set("Content-Type", "image/png")
		case ".json":
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		default:
			w.Header().Set("Content-Type", "text/html")
		}
	}

	fin, err := os.Open(StaticRoot + path)
	fmt.Println(StaticRoot + path)
	defer fin.Close()
	if err != nil {
		log.Fatal("static resource:", err)
	}

	fd, _ := ioutil.ReadAll(fin)
	_, _ = w.Write(fd)
}
