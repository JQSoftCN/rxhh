package reportsys

import (
	"net/http"
	"log"
	"html/template"
	"../../reports"
	"strconv"
	"fmt"
	"time"
	"../../funcs"
	"golang.org/x/net/websocket"
	"encoding/json"
	"github.com/qjsoftcn/confs"
)

func handleRunReport() {

	rWsMap = make(map[string]bool)

	http.HandleFunc("/report/run", run)

	http.HandleFunc("/report/rt/head", rtHead)
	http.HandleFunc("/report/rt/mid", rtMid)
	http.HandleFunc("/report/rt/bottom", rtBottom)

}

type RHead struct {
	Name string
	DateTime string
	WsUrl string
}

//report template head
func rtHead(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	reportId := r.FormValue("rId")
	rid, _ := strconv.Atoi(reportId)
	ri := reports.GetReportInfo(rid)

	rh:=RHead{}
	rh.Name=ri.Name
	rh.DateTime=ri.TimeType.Html()
	rh.WsUrl=makeWsUrl(r.Host,reportId)

	t, err := template.ParseFiles("web/report/head.html","web/public/rhead.html")
	if err != nil {
		log.Println(err)
	}

	err = t.Execute(w, ri)
	if err != nil {
		log.Println(err)
	}
}

//report template head
func rtBottom(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	reportId := r.FormValue("rId")
	rid, _ := strconv.Atoi(reportId)
	ri := reports.GetReportInfo(rid)

	rName := ri.Name
	t, err := template.ParseFiles("rts/" + rName + "/tabs.html")
	if err != nil {
		log.Println(err)
	}

	err = t.Execute(w, ri)
	if err != nil {
		log.Println(err)
	}
}

type RtMidData struct {
	WsURL string
}

//report cell style
type RCellStyle struct {
	BgColor   string
	FontAgain string
	FontColor string
}

//report cell val
type RCellVal struct {
	Id    string
	Text  string
	Style RCellStyle
}

func newRCellVal(){

}

func calcReport(ch chan int64, ri reports.ReportInfo, ws *websocket.Conn) {
	start := time.Now()
	end := start
	rconfs := reports.GetRConf(ri.Name)

	fctx := funcs.NewFuncContext(start, end)
	//set const
	for index, cc := range *rconfs {
		if cc.IsConst() {
			fctx.AddConst(cc.Key, cc.Cell)
		} else {
			f, ok := cc.Cell.(string)
			if ok {
				fctx.AddFormula(cc.Key, f)
			} else {
				fmt.Println("No.", index, " Formula:", cc, " convert err")
			}
		}
	}

	//calc run
	t1 := time.Now()
	for _, cc := range *rconfs {
		if cc.IsFormula() {
			fctx.RunFormula(cc.Key)
			r := fctx.GetResult(cc.Key)
			//rText:=fmt.Sprint(r)

			//rcVal:=newRCellVal(cc.key,rText,fctx.get)

			bs, _ := json.Marshal(r)
			ws.Write(bs)
		}
	}

	hs := time.Since(t1).Nanoseconds()

	for _, ce := range *fctx.GetCalcErrs() {
		fmt.Println("err", ce)
	}

	for index, cc := range *rconfs {
		fmt.Println(index, cc.Key, fctx.GetResult(cc.Key))
	}

	ch <- hs

}

var rWsMap map[string]bool

func makeWsUrl(host,rid string)string{
	port:=confs.GetString("top","Web","Port")
	return "ws://" + host + port + "/r" + rid
}

//report template head
func rtMid(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	reportId := r.FormValue("rId")
	rid, _ := strconv.Atoi(reportId)
	ri := reports.GetReportInfo(rid)

	rName := ri.Name

	sheet := r.FormValue("sheet")
	t, err := template.ParseFiles("rts/" + rName + "/" + sheet + ".html")
	if err != nil {
		log.Println(err)
	}

	isHandle, ok := rWsMap[reportId]
	if !ok {
		if !isHandle {
			http.Handle("/r"+reportId, websocket.Handler(func(ws *websocket.Conn) {
				ch := make(chan int64)
				go calcReport(ch, ri, ws)
				hs := <-ch
				fmt.Println("calc time spaced:", hs)
			}))

			rWsMap[reportId] = true
		}
	}

	rtmd := RtMidData{}
	err = t.Execute(w, rtmd)
	if err != nil {
		log.Println(err)
	}
}

//report template head
func run(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	reportId := r.FormValue("reportId")
	rid, _ := strconv.Atoi(reportId)
	ri := reports.GetReportInfo(rid)

	rName := ri.Name

	t, err := template.ParseFiles("rts/" + rName + "/f.html")
	if err != nil {
		log.Println(err)
	}

	t.Execute(w, ri)
}
