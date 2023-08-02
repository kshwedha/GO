package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/bigkevmcd/go-configparser"
)

type global struct {
	base_url string
}

func getConfig(fn string) *configparser.ConfigParser {
	p, err := configparser.NewConfigParserFromFile(fn)
	if err != nil {
		fmt.Print("!! Error reading a file.", err)
	}
	return p
}

func getConfigSections(config *configparser.ConfigParser) []string {
	return config.Sections()
}

func getConfigVal(config *configparser.ConfigParser, section string, value string) string {
	v, err := config.Get("URL", "baseUrl")
	if err != nil {
		fmt.Print("!! Error reading a file.", err)
	}
	return v

}

func setConfig(config *configparser.ConfigParser) {
	err := config.Set("section", "newoption", "value")
	if err != nil {
		fmt.Print("!! Error reading a file.", err)
	}
}

func typeOf(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

func drawLine(count int, symbol string) {
	for i := 0; i < count; i++ {
		fmt.Print(symbol)
	}
	fmt.Print("\n")
}

func methodLevelCallRequest(url string, value []string, method string) {
	// func callRequest(w *http.ResponseWriter, r *http.Response)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func (GLOBAL global) initialize_global() *global {
	gbl := &GLOBAL
	gbl.base_url = ""
	return gbl
}

func addHeaders(req *http.Request) *http.Request {
	token := ""
	req.Header.Add("content-type", "application/json")
	req.Header.Add("dealerid", "4")
	req.Header.Add("tenantid", "undefined")
	req.Header.Add("tenantname", "stg-cacargroup")
	req.Header.Add("tekion-api-token", token)
	return req
}

func callRequest(url string, method string, body *strings.Reader) []byte {
	// body can be *string.Reader or io.Reader
	var empty_byte []byte
	client := &http.Client{}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		fmt.Println(err)
		return empty_byte
	}
	req = addHeaders(req)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return empty_byte
	}
	defer res.Body.Close()

	_body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return empty_byte
	}
	return _body
}

func readConfigSectionItems(config *configparser.ConfigParser, section string) (configparser.Dict, error) {
	return config.Items(section)
}

func readConfigSection(config *configparser.ConfigParser) []string {
	sections := getConfigSections(config)
	return sections
}

func getFileConfig(fn string) *configparser.ConfigParser {
	url_fn := fn
	config := getConfig(url_fn)
	return config
}

func (sm ScanManagement) search() {
}

type ScanManagement struct {
}

func createFile(file_name string) *os.File {
	// var (
	// 	newFile *os.File
	// 	err     error
	// )
	newFile, err := os.Create(file_name)
	if err != nil {
		log.Fatal(err)
	}
	return newFile
}

func _getPathOrContent(path string) {
	const WINDOWS_PATH_SEPARATOR = '\\'
	const PATH_SEPARATOR = '/'

	// checking if the path is not empty
	if path != "" {
		path_list := strings.Split(path, "/")
		if runtime.GOOS == "windows" {
			path = strings.Join(path_list, string(WINDOWS_PATH_SEPARATOR))
		} else {
			// darwin
			path = strings.Join(path_list, string(PATH_SEPARATOR))
		}

		// checking the stats of the path provided
		info, err := os.Stat(path)

		if os.IsNotExist(err) {
			fmt.Println("\t+ The file path provided is not valid.")
		}

		// is path a file?
		if info.IsDir() {
			fmt.Println("The path provided is a directory.")
		} else {
			fmt.Println("The path provided is a file, reading content.")
			// file chunk size
			file_size := 1024

			// file, err = os.Open(path, os.O_RDONLY)
			file, err := os.OpenFile(path, os.O_RDONLY, 0666)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()
			data := make([]byte, file_size)
			// fmt.Print(data)

			for {
				n, err := file.Read(data)
				if err == io.EOF {
					break
				}
				if err != nil {
					log.Fatal(err)
				}
				string_data := string(data[:n])
				fmt.Println("Read", n, "bytes:\n")
				fmt.Print(string_data)
				fmt.Print("\n")
			}

		}
	} else {
		log.Fatal("\t!! Invalid file/path value..")
	}
}

func getPathOrContent(path string, is_file bool) []byte {
	const WINDOWS_PATH_SEPARATOR = '\\'
	const PATH_SEPARATOR = '/'

	file_size := 1024
	if path != "" {
		path_list := strings.Split(path, "/")
		if runtime.GOOS == "windows" {
			path = strings.Join(path_list, string(WINDOWS_PATH_SEPARATOR))
		} else {
			// darwin
			path = strings.Join(path_list, string(PATH_SEPARATOR))
		}
		if is_file {
			// file, err = os.Open(path, os.O_RDONLY)
			file, err := os.OpenFile(path, os.O_RDONLY, 0666)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()
			data := make([]byte, file_size)
			return data
			// for {
			// 	n, err := file.Read(data)
			// 	if err == io.EOF {
			// 		break
			// 	}
			// 	if err != nil {
			// 		log.Fatal(err)
			// 	}
			// 	fmt.Println("Read", n, "bytes:\n", string(data[:n]))
			// }
		}
	} else {
		log.Fatal("File/Path does not exist.")
	}
	var empty_byte []byte
	return empty_byte
}

func getRandomNumberOrSequence() {
	upperLimit := 99
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	r.Intn(upperLimit)
}

func main() {
	fmt.Print("[*] Initiating process.../\n")
	_getPathOrContent("dir1/dir2/some.txt")
	// json_data := getPathOrContent("del.json", true)

}

func _main1() {
	config := getFileConfig("config.cfg")
	sections := readConfigSection(config)
	drawLine(90, "=")
	for _, section := range sections {
		drawLine(90, "=")
		fmt.Println(section)
		drawLine(90, "=")
		dict, err := readConfigSectionItems(config, section)
		if err != nil {
			panic(err)
		}
		for integ, key := range dict {
			fmt.Println(integ, ":", key)
			drawLine(90, "=")
		}
	}
	drawLine(90, "=")
}

func _main2() {
	g := &global{}
	global := g.initialize_global()

	method := "POST"
	api := "/api/documentV2/u/docset/search?locale=en_US"
	url := global.base_url + api

	payload := strings.NewReader(`{"sort":[{"field":"createdTime","order":"DESC"},{"field":"id","order":"ASC"}],"filters":[],"searchText":"","groupBy":[],"includeFields":[],"searchableFields":[],"excludeFields":[],"pageInfo":{"start":0,"rows":50}}`)

	// // this below two lines are for jsoning key:value payload (map) datatype of io.Writer datattype payload
	// m := new(bytes.Buffer)
	// json.NewEncoder(payload).Encode(m)

	// // this is for array, dict, ...
	// // body := []string{"i-056b856c37791a7dd", "i-0af5647ab921bce8a"}
	// // postBody, _ := json.Marshal(body)
	// // b := bytes.NewBuffer(postBody)

	fmt.Println(string(callRequest(url, method, payload)))
}
