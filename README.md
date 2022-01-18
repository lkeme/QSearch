# QSearch

...

## Development-environment

- [**Editor: Goland 2021.3.x**](https://www.jetbrains.com/go/)
- [**GoVersion: 1.17.6 windows/amd64**](https://go.dev/)
- [**GoProxy: Aliyun**](https://developer.aliyun.com/mirror/goproxy)

## Routes

| Method | Path                                                                | Desc       |
|--------|---------------------------------------------------------------------|------------|
| GET    | [/swagger/index.html](http://localhost:3000/swagger/index.html)     | SwaggerApi |
| POST   | [/api/v1/system/health](http://localhost:3000/api/v1/system/health) | health     |
| POST   | ...                                                                 | ...        |

## Install SwagApi

```shell
# cli
$ go get -u github.com/swaggo/swag/cmd/swag
# middleware
$ go get github.com/swaggo/gin-swagger
# swagger embed files
$ go get github.com/swaggo/gin-swagger/swaggerFiles
# version
$ swag --version
# generate docs/docs.go
$ swag init
```

```text
// add
"github.com/swaggo/gin-swagger"
"github.com/swaggo/gin-swagger/swaggerFiles"
_ "github.com/swaggo/gin-swagger/example/docs"
```

```text
// SwagAPI register swagger handler
Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
```