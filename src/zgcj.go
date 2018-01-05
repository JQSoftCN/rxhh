package main

import (
	"net/http"
	"os"
	"io"
	"github.com/tealeg/xlsx"
	"fmt"
	"io/ioutil"
)

func readRankFile() {
	url := "http://stock.gtimg.cn/data/get_hs_xls.php?id=rankash&type=1&metric=chr"

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	f, err := os.Create("tx.csv")
	if err != nil {
		panic(err)
	}
	io.Copy(f, res.Body)
}

func readRankData() {

	bs, err := ioutil.ReadFile("tx2.xlsx")
	if err != nil {
		panic(err)
		return
	}

	xlsFile, err := xlsx.OpenBinary(bs)
	if err != nil {
		panic(err)
		return
	}

	sheet := xlsFile.Sheets[0]
	for ri, row := range sheet.Rows {
		fmt.Println(ri, row.Cells)
	}
}

func main() {
	//读取排名文件
	readRankFile()
	//读取排名数据
	//readRankData()
}
