# go-webshell
A simple webshell written in Go. Supports Windows and any Unix-like OS with The Bourne shell (sh).

#### Compilation
```
env GOOS=<OS> GOARCH=<ARCHITECTURE> go build shell.go
```

#### Usage
Running with no arguments will launch the HTTP server on **0.0.0.0:8080**
```
./shell -ip <IP> -port <PORT>
```
