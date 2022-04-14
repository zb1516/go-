package main

import (
	"calculator/until"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

//定义全局模版变量
var tmp *template.Template

func index(w http.ResponseWriter,r *http.Request)  {
	if r.Method == "GET" {
		tmps,err:=template.ParseFiles("view/index.html")
		if err != nil {
			fmt.Println("模版加载失败:", err)
		}
		tmp = tmps
		errs := tmp.Execute(w, nil)
		if errs != nil{
			fmt.Println("执行模版失败:", errs)
		}
	}else{
		//解析form表单
		r.ParseForm()
		//取出form表单中提交的值
		x,_:=strconv.Atoi(r.Form.Get("num1"))
		y,_:=strconv.Atoi(r.Form.Get("num2"))
		euma := r.Form.Get("submit")
		data := until.NewGetNum(x , y)
		var res int
		switch euma {
		case "+":
			res =data.Add(data)
		case "-":
			res =data.Minus(data)
		case "*":
			res =data.Take(data)
		case "/":
			res =data.Division(data)
		}
		tmpErr :=tmp.Execute(w, &until.GetNum{Result: res})
		if tmpErr != nil {
			fmt.Println("模版显示错误:",tmpErr)
		}
	}
}

func main()  {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080",nil)
}