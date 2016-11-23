# Google-in-go
A simple google search agent server

#### How to Use ?

* On Linux or Windows

  You should download go runtime.

  then

  ``` shell
  git clone https://github.com/BackTrackCRoot/Google-in-go
  cd Google-in-go
  go build google.go -o google # Windows use -o google.exe
  ./google & # or google.exe
  ```

  ​

* On Heroku platform

``` shell
git clone https://github.com/BackTrackCRoot/Google-in-go
cd Google-in-go
heroku create
git add -A
git commit -m "init"
git push heroku master
heroku open
```

#### Update

1. Fix User-Agent

2. Support search in Chinese

3. Remove ad etc.

4. Optimization request header

#### Problem

1. The image search is not support
2. Thin agent to load is very slow