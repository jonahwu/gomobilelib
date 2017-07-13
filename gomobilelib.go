package gomobilelib

import (
	//	"encoding/json"
	//	"errors"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	//"os"
	"text/template"
	//	"time"
)

var jsonStrVel = `[{
        "metric": "testgps",
        "timestamp": "{{ .timestampdata  }}",
        "value": "{{ .xdata  }}",
        "tags": {
            "id": "{{ .iddata }}",
            "loc":"x"
                }    },
        {
        "metric": "testgps",
        "timestamp": "{{ .timestampdata }}",
        "value": "{{ .ydata  }}",
        "tags": {
        "id": "{{ .iddata }}",
        "loc":"y"
        }    },
        {
        "metric": "testgps",
        "timestamp": "{{ .timestampdata }}",
        "value": "{{ .vdata  }}",
        "tags": {
        "id": "{{ .iddata }}",
        "loc":"v"
        }    }

]`

var jsonStr = `[{
        "metric": "testgps",
        "timestamp": "{{ .timestampdata  }}",
        "value": "{{ .xdata  }}",
        "tags": {
            "id": "{{ .iddata }}",
            "loc":"x"
                }    },
        {
        "metric": "testgps",
        "timestamp": "{{ .timestampdata }}",
        "value": "{{ .ydata  }}",
        "tags": {
        "id": "{{ .iddata }}",
        "loc":"y"
        }    }

]`

func Test() {
	fmt.Println("I am in GOmobilelib")
}

func Testpassfloat(ff float64) string {
	fff := ff + 10.
	return fmt.Sprint("", fff)
}

func Testpass() string {
	str := `{"page": "1", "fruits": ["apple", "peach"]}`
	return str
}

func SendtoGCPTest() string {
	//url := "http://localhost:8000"
	url := "http://35.189.170.202:8000"
	request, _ := http.NewRequest("GET", url, nil)
	//request.Header.Add("name", "kala")
	resp, _ := http.DefaultClient.Do(request)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	return "haha in mobile"

}
func getDataVel(strTimeStamp string, strlati string, strlong string, strvel string, strid string) []byte {
	m := map[string]interface{}{}
	//  m["xdata"] = "23.5555"
	//  m["ydata"] = "123.5555"
	m["xdata"] = strlati
	m["ydata"] = strlong
	m["vdata"] = strvel
	m["timestampdata"] = strTimeStamp
	m["iddata"] = strid
	// adding id and timestamp and put xdata and ydata and all of them to argument
	t := template.Must(template.New("").Parse(jsonStrVel))
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, m); err != nil {
		panic(err)
	}
	fmt.Println(tpl.String())
	return tpl.Bytes()
}

func getData(strTimeStamp string, strlati string, strlong string, strid string) []byte {
	m := map[string]interface{}{}
	//  m["xdata"] = "23.5555"
	//  m["ydata"] = "123.5555"
	m["xdata"] = strlati
	m["ydata"] = strlong
	m["timestampdata"] = strTimeStamp
	m["iddata"] = strid
	// adding id and timestamp and put xdata and ydata and all of them to argument
	t := template.Must(template.New("").Parse(jsonStr))
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, m); err != nil {
		panic(err)
	}
	fmt.Println(tpl.String())
	return tpl.Bytes()
}

func SendGPS(strTimeStamp string, strlati string, strlong string, strid string) {
	url := "http://35.189.170.202:14242/api/put?details"
	fmt.Println("URL:>", url)
	jsonStrData := getData(strTimeStamp, strlati, strlong, strid)
	//  panic("aaa")
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStrData))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

}

func SendGPSVel(strTimeStamp string, strlati string, strlong string, strvel string, strid string) {
	url := "http://35.189.170.202:14242/api/put?details"
	fmt.Println("URL:>", url)
	jsonStrData := getDataVel(strTimeStamp, strlati, strlong, strvel, strid)
	//  panic("aaa")
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStrData))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

}
