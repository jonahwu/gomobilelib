package gomobilelib

import (
	"fmt"
	"github.com/tidwall/gjson"
	//	"time"
	"encoding/json"
)

type Gps2dLoc struct {
	Long float64
	Lati float64
}

type Gps4dLoc struct {
	Long      float64
	Lati      float64
	Timestamp int
}

type ListGps2dLocTim struct {
	Loc       []Gps2dLoc
	Timestamp int
}

type TestInfo struct {
	Testnum    float64
	Gpsloc     *Gps2dLoc
	PrevLoc    Gps4dLoc
	CurrentLoc Gps4dLoc
}

func NewGp2dLoc() *Gps2dLoc {
	a := Gps2dLoc{}
	return &a
}

func (tt *TestInfo) InitState() {
	fmt.Println("show the PrevLoc", tt.PrevLoc)
}

func (tt *TestInfo) ShowData(latidata float64, longdata float64) {
	//	if tt
	fmt.Println("initial testnum", tt.Testnum)
	fmt.Println("testnum put in client", tt.Testnum)
	fmt.Println("lati and long:", latidata, longdata)
	fmt.Println("show testinfo data")
	tt.Testnum = 77.88
	fmt.Println("show Gpsloc", tt.Gpsloc)
}
func (tt *TestInfo) GetData() float64 {

	return tt.Testnum

}

type UserGpsInfo struct {
	CurrentLoc      Gps4dLoc
	CameraLoc       ListGps2dLocTim
	TargetCameraLoc ListGps2dLocTim
	// > 5 sec to update
	PrevLoc Gps4dLoc
	// as map filter used
	CreateLoc Gps4dLoc
}

func NewCamera() *TestInfo {
	a := &TestInfo{}
	return a
}

func (f *UserGpsInfo) DetectCamera() {
	fmt.Println("into struct detectcamera")
	return
}

func CalGpsDistance(a string) string {
	//	a := `{"a":"b"}`
	//	timestamp := time.Now().Unix()
	fmt.Println("into glib")
	fmt.Println(a)
	//return gjson.Get(a, "0.a").String()
	arr := gjson.Get(a, "gdata").Array()

	fmt.Println("glib array", arr[0].String())
	for i := range arr {
		fmt.Println("glib loop", i)
	}
	//return "0"

	var ll []string
	ll = append(ll, "aa")
	ll = append(ll, "bb")
	m := map[string]interface{}{}
	m["gdata"] = ll
	b, _ := json.Marshal(m)
	fmt.Println("glib b", string(b))
	return string(b)

	//	return gjson.Get(a, "a").String()
}
