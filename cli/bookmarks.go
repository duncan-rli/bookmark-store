//
// bookmark commandline interface
//

package main

import (
	"../client"
	"flag"
	"fmt"
	"io/ioutil"
)

func usageText() {
	fmt.Println("bookmarks")
	fmt.Println("   bookmarks store filename")
	fmt.Println("		read the json contents from the file and store")
	fmt.Println("   bookmarks retrieve")
	fmt.Println("		retrieve the contents of the bookmark store")
	fmt.Println("   bookmarks removeurls")
	fmt.Println("		retrieve the contents of the store, remove urls and replace in the store")
	fmt.Println("")
	fmt.Println("")
}

func main() {
	flag.Usage = usageText
	flag.Parse()
	args := flag.Args()

	var clientObj client.ClientStruct

	// TODO validate parameters
	if args[0] == "retrieve" {
		data, err := clientObj.Retrieve()
		if nil != err {
			fmt.Println(err)
		} else {
			fmt.Println(string(data))
		}
	} else if args[0] == "store" {
		// get file name
		filename := string([]byte(args[1]))
		file, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Println(err)
		}
		err = clientObj.Store(file)
		if nil != err {
			fmt.Println(err.Error())
		}
	} else if args[0] == "removeurls" {
		err := clientObj.RemoveUrls()
		if nil != err {
			fmt.Println(err.Error())
		}
	} else {
		fmt.Println("Unknown command.")
		usageText()
		return
	}

}
