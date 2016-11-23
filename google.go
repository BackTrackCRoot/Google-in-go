package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "log"
    "net/http"
    "regexp"
    "strings"
)

var SourceUrl = "https://www.google.com/"

type Agent struct{}

func (a Agent) ServeHTTP(
    w http.ResponseWriter,
    r *http.Request) {
    url := SourceUrl + r.RequestURI
    
    //Force Simplified Chinese
    if r.RequestURI == "/" {
        url = SourceUrl + "/?hl=zh-CN"
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
        html = strings.Replace(html, SourceUrl, "//"+r.Host+"/", -1)

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
        
        html = strings.Replace(html, "Google.com in English", "RazerNiz Go", -1)
        htmlData = []byte(html)
        //*/
        w.Write(htmlData)
    } else {
        w.WriteHeader(500)
    }
}

func main() {
    fmt.Println("哦~这个网站跑起来了，请访问您的" + os.Getenv("PORT") + "端口")
    fmt.Println("按Ctrl + C 结束该程序 ...")
    var a = Agent{}
    err := http.ListenAndServe("0.0.0.0:" + os.Getenv("PORT"), a)
    if err != nil {
        log.Fatal(err)
    }
}
