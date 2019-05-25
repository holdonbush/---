package controller

import (
	"net/http"
	"io/ioutil"
	"log"
	"encoding/json"
	"fmt"
	"weixinminiprogram/dbase"
	"time"
)

type data struct {
	Name string
	Num string
	Explain string
}

type data1 struct {
	Name string
}

type back struct {
	Total float64
	YesterdayT float64
	D3t float64
	D4t float64
	D5t float64
	D6t float64
	D7t float64
}

func Cal(w http.ResponseWriter, r *http.Request) {
	fmt.Println("1")
	fmt.Println(r.Method)
	r.ParseForm()
	result, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()
	fmt.Println(string(result))
	var d data
	json.Unmarshal(result,&d)
	fmt.Println(d.Num,d.Explain,d.Name)
	dbase.Insert(d.Name,d.Num,d.Explain,time.Now().Format("2006-01-02"))
	total := dbase.Select(time.Now().Format("2006-01-02"),d.Name)

	//

	var b back
	b.Total = total
	b.YesterdayT = search(d.Name,-1)
	b.D3t = search(d.Name,-2)
	b.D4t = search(d.Name,-3)
	b.D5t = search(d.Name,-4)
	b.D6t = search(d.Name,-5)
	b.D7t = search(d.Name,-6)
	bac, err := json.Marshal(b)
	checkerr(err)
	fmt.Fprintf(w,string(bac))
}

func Thisday(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	result, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	checkerr(err)
	var d data1
	json.Unmarshal(result,&d)
	fmt.Println(d.Name)
	total := dbase.Select(time.Now().Format("2006-01-02"),d.Name)
	fmt.Println(total)
	//yesterday
	nTime := time.Now()
	yesterday := nTime.AddDate(0,0,-1)
	yesterdayd := yesterday.Format("2006-01-02")
	yesterdayt := dbase.Select(yesterdayd,d.Name)
	//

	var b back
	b.Total = total
	b.YesterdayT = yesterdayt
	b.D3t = search(d.Name,-2)
	b.D4t = search(d.Name,-3)
	b.D5t = search(d.Name,-4)
	b.D6t = search(d.Name,-5)
	b.D7t = search(d.Name,-6)
	bac, err := json.Marshal(b)
	checkerr(err)
	fmt.Println(string(bac))
	fmt.Fprintf(w,string(bac))
}

func checkerr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func search(name string, index int) float64 {
	ntime := time.Now()
	d := ntime.AddDate(0,0,index)
	dd := d.Format("2006-01-02")
	dt := dbase.Select(dd,name)
	return dt
}