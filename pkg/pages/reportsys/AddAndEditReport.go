package reportsys

import (
	"net/http"
	"log"
	"fmt"
	"html/template"
	"path/filepath"
	"strings"
	"os"
	"io"
	"encoding/json"
	"../../reports"
)

func handleAddAndEditReport(){
	http.HandleFunc("/report/upload", upload)
	http.HandleFunc("/report/add", addReport)
	http.HandleFunc("/report/add/submit", submitAddReport)
}


type RtUploadMsg struct {
	RName string
	Msg   string
}

//report template head
func addReport(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("web/report/addReport.html")
	if err != nil {
		log.Println(err)
	}

	t.Execute(w, nil)
}

//submit add report
func submitAddReport(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	rName := r.PostFormValue("reportName")
	timeType := r.PostFormValue("timeType")
	timeDefault := r.PostFormValue("timeDefault")
	fmtDefault := r.PostFormValue("fmtDefault")

	fmt.Println(rName, timeType, timeDefault, fmtDefault)

}

func upload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("uf")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	rfile := filepath.Base(handler.Filename)
	ldi := strings.LastIndex(rfile, ".")
	rname := rfile[0:ldi]

	dst := reports.RtDir(rname)

	fn := dst + rfile

	fmt.Println(fn)

	f, err := os.OpenFile(fn, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	io.Copy(f, file)

	reports.Parse(f.Name())

	upMsg := RtUploadMsg{rname, "模板文件上传成功"}

	jsObj, err := json.Marshal(upMsg)
	if err != nil {
		log.Println(err)
	}

	fmt.Fprintln(w, string(jsObj))
}
