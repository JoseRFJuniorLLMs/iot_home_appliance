/*
arauto server
*/

package main

/*
	#cgo LDFLAGS: -lstdc++ edison.a
	#cgo pkg-config: upm-i2clcd
	#include "edison.hpp"
*/
import "C"

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"unsafe"
)

func initLCD() {
	C.initLCD()
}

func writeLCD(r int, c int, s string) {
	cs := C.CString(s)
	C.writeLCD(C.int(r), C.int(c), cs)
	C.free(unsafe.Pointer(cs))
}

func setColorLCD(r int, g int, b int) {
	C.setColorLCD(C.int(r), C.int(g), C.int(b))
}

func removeLCD() {
	C.removeLCD()
}

type alert int

const (
	green   alert = 0
	blue    alert = 1
	yellow  alert = 2
	red     alert = 3
	gray    alert = 4
	unknown alert = 5
)

var alertTable = []string{"green", "blue", "yellow", "red", "gray", "unknown"}

var statusCode = green
var msgStatus string

func simpleCmdParse(m map[string]interface{}) error {
	/*
		for k, v := range m {
			fmt.Println(k, "->", v)
		}
	*/

	parsed := false

	for i, j := range alertTable[0 : len(alertTable)-1] { // omit the last status (it's a secret!)
		if m["cmd"] == "alert-"+j {
			statusCode = alert(i)
			parsed = true
		}
	}

	if !parsed {
		if str, ok := m["msg"].(string); ok {
			log.Printf("Error: Unable to decode request. Request = \"%s\"", str)
		} else {
			log.Print("Error: Unable to decode request. (And no msg was given!)\n")
		}
		return errors.New("Unknown command.")
	}

	if statusCode != unknown {
		if str, ok := m["msg"].(string); ok {
			msgStatus = str
		} else {
			msgStatus = ""
		}

		if statusCode == red {
			setColorLCD(255, 0, 0)
			writeLCD(0, 0, msgStatus)
		} else if statusCode == green {
			setColorLCD(0, 255, 0)
			writeLCD(0, 0, msgStatus)
		}

		log.Print("Alert: "+alertTable[statusCode]+" condition. ", msgStatus)
		return nil
	}

	return errors.New("Unknown command.")
}

func apiHandle(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	if r.Method != "POST" {
		http.Error(w, "post only", http.StatusMethodNotAllowed)
		log.Print("HTTP Method Not Allowed :" + r.Method)
		return
	}

	if r.Body == nil {
		http.Error(w, "Please send a request body", http.StatusBadRequest)
		log.Print("HTTP request body is nil")
		return
	}

	var f interface{}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
		log.Print("ioutil.ReadAll error: \"" + err.Error() + "\"")
		return
	}

	err = json.Unmarshal(b, &f)
	if err != nil {
		http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
		log.Print("json.Unmarshal error: \"" + err.Error() + "\"")
		return
	}

	m := f.(map[string]interface{})

	err = simpleCmdParse(m)
	if err != nil {
		http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
		log.Print("simpleCmdParse: \"" + err.Error() + "\"")
		return
	}

	fmt.Fprintf(w, "{\"status\": \"ok\"}\r\n")
}

func main() {

	initLCD()

	go func() {
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, os.Interrupt)

		<-sc
		removeLCD()
		os.Exit(0)
	}()

	var port = flag.String("port", "9999", "Define what TCP port")
	flag.Parse()

	http.HandleFunc("/", apiHandle)

	if err := http.ListenAndServe(":"+*port, nil); err != nil {
		log.Fatalf("ListenAndServe error: %v", err)
	}
}
