# Hahow專案

## 如何跑起這個server
```
docker-compose build
```
```
docker-compose up -d
```
## 專案架構

```
|__cmd            # todo
|  |__server
|     |__sample-controller.go
|  |__test
|     |__sample-controller.go
|__controler      # todo
|  |__controller_test.go
|  |__controller.go
|__data           # todo
|  |__data.go
|  |__error.go
|__model          # todo
|  |__dao.go
|  |__hahow_test.go
|  |__hahow.go
|  |__mock.go
|__service
|  |__service_test.go
|  |__service.go
|__tools
|  |__tools.go
```