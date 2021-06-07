## go_static_server
a simple demo of implements static web tcp proto

**Proxy static request if you run it correctly.**

### CLONE

run this command to clone this  repository.



```bash
git clone git@github.com:dengjiawen8955/go_static_server.git
```

### CONFIG

you must config the bathPath in `src/Main.java` 

```go
// basePath is where your static files' folder. 
// you should change it in a property path.
basePath := "/root/go/src/go_static_server"
```


### COMPILE AND RUN

```bash
cd go_satatic_server
# build
go build -o main main.go
# run
./main
```

output:
```bash
2021/06/07 21:53:39 listen=":9000"
# when there are some clients connect.
2021/06/07 21:53:45 NEW CLIENT:[::1]:57956
2021/06/07 21:53:52 NEW CLIENT:[::1]:57958
```



### BROWSER TEST

if you run localhost and test also on localhost


```url
# view a html
localhost:9000/index.html
# view a image.
localhost:9000/bmft.png
```
