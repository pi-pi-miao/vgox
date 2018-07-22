package controller

import (
	"encoding/json"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"web/defs"
	"web/proxy"
)

type (
	//需要在前端替换的文件
	HomePage struct {
		Name string
	}
)

//proxy转发
func ProxyHandler(w http.ResponseWriter, r *http.Request) {
	//下面是把8000域名转换成了9000
	u, _ := url.Parse("http://127.0.0.1:9000/")
	proxy := httputil.NewSingleHostReverseProxy(u)
	proxy.ServeHTTP(w, r)
}

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		re, _ := json.Marshal(defs.ErrorRequestNotRecognized)
		io.WriteString(w, string(re))
	}

	res, _ := ioutil.ReadAll(r.Body)
	apibody := &defs.ApiBody{}
	if err := json.Unmarshal(res, apibody); err != nil {
		re, _ := json.Marshal(defs.ErrorBodyParseFailed)
		io.WriteString(w, string(re))
	}

	proxy.Request(apibody, w, r)
	defer r.Body.Close()

}

func Index(w http.ResponseWriter, r *http.Request) {
	t, e := template.ParseFiles("./template/index.html")
	if e != nil {
		log.Printf("parsing template home.html err%s", e)
		return
	}
	//将模板和需要渲染的变量加进去
	t.Execute(w, nil)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	//TODO:需要加入验证信息
	t, e := template.ParseFiles("./template/login.html")

	if e != nil {
		log.Printf("parsing template home.html err%s", e)
		return
	}
	t.Execute(w, nil)
}

func UserHomeHandler(w http.ResponseWriter, r *http.Request) {
	//需要加入验证信息
	video, err := template.ParseFiles("./template/index.html")
	if err != nil {
		log.Printf("parsing template video.html err%s", err)
		return
	}
	video.Execute(w, nil)
}
