package main

import (
	"net/http"
	"../pkg/pages"
	"../pkg/confs"
	"../pkg/pages/reportsys"

)

func main() {
	http.Handle("/css/", http.FileServer(http.Dir("web/static")))
	http.Handle("/js/", http.FileServer(http.Dir("web/static")))
	http.Handle("/ext/", http.FileServer(http.Dir("web/static")))


	pages.Publish()
	reportsys.Publish()

	http.ListenAndServe(confs.WebPort(), nil)


}
