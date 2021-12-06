package main

import (
    "bytes"
    "fmt"
    "log"
    "net/http"
    "os/exec"
    "strings"
)

func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    
        keys, ok := r.URL.Query()["exec"]
        if (!ok || len(keys[0]) < 1) {
            fmt.Fprintf(w, "Missing Exec Parameter, %q", "localhost:8081/?exec=SomethingHere")
        } else {
            key := keys[0]
            
            cmd := exec.Command(strings.Split(key, " ")[0])
            var sOut bytes.Buffer
            cmd.Stdout = &sOut
            var sErr bytes.Buffer
            cmd.Stderr = &sErr
            
            error := cmd.Run()
            if (error != nil) {
                fmt.Fprintf(w, "Command error, %v\n", error)
            } else {
                fmt.Fprintf(w, "StdOut: %q\nStdErr: %q\n", sOut.String(), sErr.String())
            }
        }
    })
    log.Fatal(http.ListenAndServe(":8081", nil))
}
