package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

// SourceURL is origin web's URL
var SourceURL = "https://www.google.com/"

// Agent is proxy server
type Agent struct{}

func (a Agent) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	url := SourceURL + r.RequestURI

	//Force Simplified Chinese
	if r.RequestURI == "/" {
		url = SourceURL + "/?hl=zh-CN"
	}

	//Advs block listener
	if strings.Index(r.RequestURI, "cdn-dat/adb") != -1 {
		w.Header().Set("Content-Type", "text/html;charset=utf-8")
		w.Header().Set("Cache-Control", "public, max-age=2592000")
		w.Header().Set("Content-Length", "0")
		w.WriteHeader(200)
	}

	//Close Google safe Search
	reg := regexp.MustCompile("safe=strict")
	url = reg.ReplaceAllString(url, "safe=off")

	client := &http.Client{}
	req, err := http.NewRequest(r.Method, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header = r.Header
	//update req Header
	req.Header.Set("Accept", r.Header.Get("Accept"))
	req.Header.Set("Accept-Encoding", "hentai") //DONOT LET GZIP!!!
	req.Header.Set("Accept-Language", r.Header.Get("Accept-Language"))
	req.Header.Set("Cookie", r.Header.Get("Cookie"))
	req.Header.Set("User-Agent", r.Header.Get("User-Agent"))
	resp, err := client.Do(req)
	log.Println(r.Method + " " + fmt.Sprintf("%d", resp.StatusCode) + " " + url)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode < 500 {
		//force cache to speedup
		w.Header().Set("Cache-Control", "public, max-age=86400")
		for k, v := range resp.Header {
			if k != "Cache-Control" {
				w.Header().Set(k, v[0])
			}
		}
		//io.Copy(w, resp.Body)
		htmlData, _ := ioutil.ReadAll(resp.Body)

		//*
		html := string(htmlData)

		//Replace the blocked resource
		html = strings.Replace(html, "https://ajax.googleapis.com/ajax/libs/jquery/", "//cdn.bootcss.com/jquery/", -1)
		html = strings.Replace(html, "https://ajax.googleapis.com/ajax/libs/angularjs/", "//cdn.bootcss.com/angular.js/", -1)
		html = strings.Replace(html, SourceURL, "//"+r.Host+"/", -1)

		//Filter Advs Javascripts
		html = strings.Replace(html, "//imasdk.googleapis.com/", "/cdn-dat/adb?", -1)
		html = strings.Replace(html, "/xjs/_/js/", "/cdn-dat/adb?", -1)
		html = strings.Replace(html, "//pagead2.googlesyndication.com/", "/cdn-dat/adb?", -1)
		html = strings.Replace(html, "//partner.googleadservices.com/", "/cdn-dat/adb?", -1)
		html = strings.Replace(html, "//www.googletagservices.com/", "/cdn-dat/adb?", -1)
		html = strings.Replace(html, "//www.google-analytics.com/", "/cdn-dat/adb?", -1)
		html = strings.Replace(html, "//apis.google.com", "/cdn-dat/adb?", -1)
		html = strings.Replace(html, "//plus.google.com/", "/cdn-dat/adb?", -1)
		html = strings.Replace(html, "//ogs.google.com/", "/cdn-dat/adb?", -1)
		html = strings.Replace(html, "//client5.google.com/", "/cdn-dat/adb?", -1)
		html = strings.Replace(html, "//client4.google.com/", "/cdn-dat/adb?", -1)
		html = strings.Replace(html, "www.gstatic.com", "0.0.0.0", -1)
		html = strings.Replace(html, "ssl.gstatic.com", "0.0.0.0", -1)
		//Clean cookies auth for pc
		html = strings.Replace(html, "<div class=\"gTMtLb fp-nh\" id=\"lb\">", "<div class=\"gTMtLb fp-nh\" id=\"lb\" style=\"display: none;\">", -1)

		//Clean cookies auth for mobile
		re := regexp.MustCompile("<div jsname=\"XKSfm\" id=\".+\" jsaction=\"dBhwS:TvD9Pc;mLt3mc\">")
		html = re.ReplaceAllString(html, "<div style=\"display: none;\">")

		html = strings.Replace(html, "Google.com in English", "RazerNiz Go", -1)
		htmlData = []byte(html)
		//*/
		w.Write(htmlData)
	} else {
		w.WriteHeader(500)
	}
}

func main() {
	addr := flag.String("a", "0.0.0.0", "Listen address.")
	port := flag.String("p", "8080", "Listen port.")
	flag.Parse()
	fmt.Printf("Google proxy server was started. The Server linstened %s \nCtrl + C will close.", *port)
	var a = Agent{}
	server := fmt.Sprintf("%s:%s", *addr, *port)
	err := http.ListenAndServe(server, a)
	if err != nil {
		log.Fatal(err)
	}
}
