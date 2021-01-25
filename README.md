# Google-in-go



This tools is the Google search agent only.

And the build utility is depend on [xmake](https://github.com/xmake-io/xmake) . when your platform is Linux/Windows(x64), it can be automatically packed by upx.

### Build

``` shell
xmake
```



### How to use

``` shell
Usage of Google-in-go:
  -a string
        Listen address. (default "0.0.0.0")
  -p string
        Listen port. (default "8080")
```



The old version is support on `Heroku platform`



#### update 20210125

* fix pc/mobile blocked by setting.



#### Problem

1. The image search is not support
2. Thin agent to load is very slow
