package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type Language struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

var langs = []Language{
	{
		Id:   "1",
		Name: "GoLang",
		URL:  "go.dev",
	},
	{
		Id:   "2",
		Name: "Dart",
		URL:  "dart.dev",
	},
	{
		Id:   "3",
		Name: "Flutter",
		URL:  "flutter.dev",
	},
}

func main() {

	http.HandleFunc("/", root)
	http.HandleFunc("/langs", GetLangs)
	http.HandleFunc("/getLangById", GetLangById)
	http.HandleFunc("/addLang", AddLang)
	http.HandleFunc("/addLangForm", AddLangForm)
	http.HandleFunc("/fileUpload", FileUpload)

	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func root(w http.ResponseWriter, r *http.Request) {

	// Changing the header map after a call to WriteHeader (or
	// Write) has no effect unless the modified headers are
	// trailers.
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	w.Write([]byte("Language Home"))

	// 该方法必须在最后调用, 否则 w.Header().Set() 不会生效(下文略)
	w.WriteHeader(http.StatusOK)
}

func GetLangs(w http.ResponseWriter, r *http.Request) {

	// Accept Get Method Only
	if r.Method != http.MethodGet {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Unsupported Request Method"))
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(&langs)

	w.WriteHeader(http.StatusOK)
}

func GetLangById(w http.ResponseWriter, r *http.Request) {

	// Only Get Request
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Write([]byte("Unsupported Request Method"))
		return
	}

	// 获取URL中的请求参数
	v := r.URL.Query()
	id := v.Get("id")
	if id == "" {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Write([]byte("Cannot find \"id\" parameter"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 或直接使用 FormValue 获取
	// id := r.FormValue("id")

	for _, lang := range langs {
		if lang.Id == id {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			json.NewEncoder(w).Encode(&lang)
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}

func AddLang(w http.ResponseWriter, r *http.Request) {
	// Only POST Request
	if r.Method != http.MethodPost {
		w.Write([]byte("Unsupported Request Method"))
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	cType := r.Header.Get("Content-Type")
	contains := strings.Contains(cType, "application/json; charset=utf-8")
	if !contains {
		w.Write([]byte("Only support application/json"))
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}

	// 直接将请求体中的数据转换为结构体
	var lang = new(Language)
	json.NewDecoder(r.Body).Decode(lang)

	// 或者直接获取原始数据
	// bytes, _ := ioutil.ReadAll(r.Body)
	// log.Println(string(bytes))

	// add
	langs = append(langs, *lang)

	// response
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(lang)
	w.WriteHeader(http.StatusOK)
}

func FileUpload(w http.ResponseWriter, r *http.Request) {
	// Only POST Request
	if r.Method != http.MethodPost {
		w.Write([]byte("Unsupported Request Method"))
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	cType := r.Header.Get("Content-Type")
	contains := strings.Contains(cType, "multipart/form-data")
	if !contains {
		w.Write([]byte("Only support application/json"))
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}

	// Limit your max input length!
	r.ParseMultipartForm(32 << 20)

	username := r.FormValue("username")
	_, header, err := r.FormFile("avatar")
	if err != nil {

	}

	var info = make(map[string]interface{})
	info["username"] = username
	info["filename"] = header.Filename
	info["filesize"] = header.Size

	// response
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(&info)
	w.WriteHeader(http.StatusOK)
}

func AddLangForm(w http.ResponseWriter, r *http.Request) {

	// Only POST Request
	if r.Method != http.MethodPost {
		w.Write([]byte("Unsupported Request Method"))
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	cType := r.Header.Get("Content-Type")
	contains := strings.Contains(cType, "application/x-www-form-urlencoded")
	if !contains {
		w.Write([]byte("Only support application/x-www-form-urlencoded"))
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}

	id := r.FormValue("id")
	name := r.FormValue("name")
	url := r.FormValue("url")

	// or
	// values := r.PostForm

	var lang = Language{
		Id:   id,
		Name: name,
		URL:  url,
	}

	// response
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(&lang)
	w.WriteHeader(http.StatusOK)
}
