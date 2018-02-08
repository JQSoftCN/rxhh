package main

import (
	"net/http"
	"../pkg/pages"
	"../pkg/pages/reportsys"

	"github.com/qjsoftcn/confs"
)

func main() {
	http.Handle("/css/", http.FileServer(http.Dir("web/static")))
	http.Handle("/js/", http.FileServer(http.Dir("web/static")))
	http.Handle("/ext/", http.FileServer(http.Dir("web/static")))


	pages.Publish()
	reportsys.Publish()

	port:=confs.GetString("top","Web","Port")
	http.ListenAndServe(port, nil)


}
