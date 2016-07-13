package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type alert int

const (
	red    alert = 0xFF0000
	green  alert = 0x00FF00
	blue   alert = 0x0000FF
	yellow alert = 0xFFFF00
	grey   alert = 0x808080
)

var status = green
var msgStatus string

func simpleCmdParse(m map[string]interface{}) error {
	for k, v := range m {
		fmt.Println(k, "->", v)
	}

	fmt.Println(m["cmd"])

	switch m["cmd"] {
	case "alert-green":
		status = green
		if str, ok := m["msg"].(string); ok {
			msgStatus = str
		} else {
			msgStatus = ""
		}
		log.Print("Alert: Normal condition. ", msgStatus)
		return nil
	case "alert-red":
		status = red
		if str, ok := m["msg"].(string); ok {
			msgStatus = str
		} else {
			msgStatus = ""
		}
		log.Print("Alert: Red condition. ", msgStatus)
		return nil
	case "alert-yellow":
		status = yellow
		if str, ok := m["msg"].(string); ok {
			msgStatus = str
		} else {
			msgStatus = ""
		}
		log.Print("Alert: Yellow condition. ", msgStatus)
		return nil
	case "alert-blue":
		status = blue
		if str, ok := m["msg"].(string); ok {
			msgStatus = str
		} else {
			msgStatus = ""
		}
		log.Print("Alert: Blue condition. ", msgStatus)
		return nil
	case "alert-gray":
		status = grey
		if str, ok := m["msg"].(string); ok {
			msgStatus = str
		} else {
			msgStatus = ""
		}
		log.Print("Alert: Gray condition. ", msgStatus)
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

	http.HandleFunc("/", apiHandle)

	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatalf("ListenAndServe error: %v", err)
	}
}
