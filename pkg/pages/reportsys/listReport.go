package reportsys

import (
	"net/http"
	"log"
	"html/template"
)

func handleListReport() {
	http.HandleFunc("/report", listReport)
}

type lrdata struct {
}

//report template head
func listReport(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("web/report/list.html", "web/public/css.html", "web/public/js.html")
	if err != nil {
		log.Println(err)
	}
	index := lrdata{}
	t.Execute(w, index)

}
