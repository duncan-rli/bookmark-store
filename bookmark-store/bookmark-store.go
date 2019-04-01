// bookmark store server
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type Bookmarks = map[string]interface{}

var bookmarks Bookmarks

func StoreBookmarks(w http.ResponseWriter, req *http.Request) {

	defer req.Body.Close()

	jData, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println("Error ReadAll: ", err)
		return
	}

	err = json.Unmarshal(jData, &bookmarks)
	if err != nil {
		// err to client
		io.WriteString(w, "Error JSON: "+err.Error())
		return
	}
	fmt.Println("store: ", bookmarks)
	return
}

func RetrieveStore(w http.ResponseWriter, req *http.Request) {

	defer req.Body.Close()

	data, err := json.Marshal(&bookmarks)
	if err != nil {
		s := `{"Error JSON Marshall: ` + err.Error() + `"}`
		io.WriteString(w, s)
		return
	} else {
		// reply to client with data
		s := string(data)
		io.WriteString(w, s)
	}

}

func main() {
	http.HandleFunc("/store", StoreBookmarks)
	http.HandleFunc("/retrieve", RetrieveStore)
	err := http.ListenAndServe(":8443", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
