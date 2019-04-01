// Used by CLI as client to connect to Bookmark server
package client

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type ClientStruct struct{}
type Bookmarks = map[string]interface{}

func postData(url string, content []byte) ([]byte, error) {

	// Create a transport
	tr := &http.Transport{
		DisableCompression: true,
	}

	hClient := &http.Client{Transport: tr}
	req, err := http.NewRequest("POST", "http://localhost:8443/"+url, bytes.NewBuffer(content))
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	// POST
	resp, err := hClient.Do(req)
	if err != nil {
		return nil, err
	}

	// get the body in a string
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)

	return data, err
}

func (ClientStruct) Retrieve() (string, error) {
	var returnedData string
	var err error

	rData, err := postData("retrieve", nil)
	if err == nil {
		returnedData = string(rData)
	}
	return returnedData, err
}

func (ClientStruct) Store(bookmarks []byte) error {

	_, err := postData("store", bookmarks)

	return err
}

func (s ClientStruct) RemoveUrls() error {

	type Bookmark map[string]interface{}
	bookmarks := make(Bookmark)

	// get bookmark data
	rData, err := s.Retrieve()
	if err == nil {
		// reformat json
		var prettyJSON bytes.Buffer
		err = json.Indent(&prettyJSON, []byte(rData), "", "    ")
		if err != nil {
			return err
		}
		err = json.Unmarshal(prettyJSON.Bytes(), &bookmarks)

		if err == nil {
			// remove urls
			for k, val := range bookmarks {
				// find and remove
				removeUrls(k, &val)
				// store back in structure
				bookmarks[k] = val
			}
		}
	}

	jdata, err := json.Marshal(&bookmarks)
	if err == nil {
		// put back book mark data
		err = s.Store([]byte(jdata))
	}

	return err
}

func removeUrls(k string, val *interface{}) {

	switch v := (*val).(type) {
	case []interface{}:
		for i, entry := range v {
			removeUrls("", &entry)
			v[i] = entry
		}
	case map[string]interface{}:
		for ke, va := range v {
			removeUrls(ke, &va)
			v[ke] = va
		}
	case string:
		if k == "url" {
			*val = ""
		}
	default:
		// unknown type
	}

}
