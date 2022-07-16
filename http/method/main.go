package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type UserInfo struct {
	Name string `json:"name"`
	Age  uint8  `json:"age"`
}

func UserHandlerFunc(w http.ResponseWriter, r *http.Request) {

	// Handle different request method

	switch r.Method {
	case http.MethodGet:
		HandlerGet(w, r)
	case http.MethodPost:
		HandlerPost(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Unsupported request method"))
	}
}

func HandlerGet(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")
	age := r.FormValue("age")

	m := make(map[string]interface{})
	m["name"] = name
	m["age"] = age

	b, _ := json.Marshal(&m)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func HandlerPost(w http.ResponseWriter, r *http.Request) {

	ct := r.Header.Get("Content-Type")
	ctv := strings.Split(ct, ";")

	switch ctv[0] {
	case "application/x-www-form-urlencoded":

		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() form err: %s\n", err)
		}

		// Get form value can use following method:
		// v := r.PostForm.Get("name")
		// v := r.FormValue("name")
		name := r.PostFormValue("name")
		age := r.PostFormValue("age")

		m := make(map[string]interface{})
		m["name"] = name
		m["age"] = age

		b, _ := json.Marshal(&m)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(b)

	case "multipart/form-data":

		// limit your max input length!
		r.ParseMultipartForm(32 << 20)

		name := r.FormValue("name")

		// upload file
		f, fh, err := r.FormFile("file")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "read file err: %s\n", err)
			break
		}
		defer f.Close()

		var buf bytes.Buffer
		io.Copy(&buf, f)

		// handler file
		m := make(map[string]interface{})
		m["name"] = name
		m["filename"] = fh.Filename
		m["size"] = fh.Size
		m["content"] = buf.String()

		b, _ := json.Marshal(&m)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(b)

	case "application/json":

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "read json err: %s\n", err)
			break
		}

		// handler json data
		u := UserInfo{}
		_ = json.Unmarshal(b, &u)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(b)

	default:

	}
}

func main() {

	http.HandleFunc("/user", UserHandlerFunc)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
