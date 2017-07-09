package gomobilelib

import (
	"fmt"
	"github.com/tidwall/gjson"
	//	"time"
	"encoding/json"
	"math"
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

type GLibInfo struct {
	Gpsloc          *Gps2dLoc // for gmob used for notification
	Testnum         float64
	PrevLoc         Gps4dLoc
	CurrentLoc      Gps4dLoc
	TargetCameraLoc ListGps2dLocTim
	CreateLoc       Gps4dLoc
}

func NewGp2dLoc() *Gps2dLoc {
	a := Gps2dLoc{}
	return &a
}

func (tt *GLibInfo) InitState(id string, ts int, locx float64, locy float64) {
	fmt.Println("show the PrevLoc", tt.PrevLoc)

}

func (tt *GLibInfo) Start(ts int, locx float64, locy float64) int {
	fmt.Println("start now", ts)
	return 0

}
func (tt *GLibInfo) FilterDistance(tarx float64, tary float64) (float64, int) {
	locx := tt.CurrentLoc.Lati
	locy := tt.CurrentLoc.Long
	distcurr := tt.FilterCalDistance(locx, locy, tarx, tary)
	locpx := tt.PrevLoc.Lati
	locpy := tt.PrevLoc.Long
	distprev := tt.FilterCalDistance(locpx, locpy, tarx, tary)
	fmt.Println("dist current, dist prev", distcurr, distprev)
	if distprev < distcurr {
		return distcurr, 0
	} else {
		return distcurr, 1
	}

}

func (tt *GLibInfo) GLibFilter(ts int, gpsx float64, gpsy float64) (float64, int) {
	tarx := 25.080223
	tary := 121.697908
	tt.UpdateCurrent(ts, gpsx, gpsy)
	dist, runflag := tt.FilterDistance(tarx, tary)
	tt.FilterUpdatePrev(ts, gpsx, gpsy)
	if runflag != 0 {
		var notiflag int
		fmt.Println(dist)
		if dist < 500.0 {
			notiflag = 1
		} else {
			notiflag = 0
		}
		return dist, notiflag
	} else {
		return dist, 0
	}
}

func (tt *GLibInfo) UpdateCurrent(ts int, gpsx float64, gpsy float64) {
	tt.CurrentLoc.Lati = gpsx
	tt.CurrentLoc.Long = gpsy
	tt.CurrentLoc.Timestamp = ts
}

func (tt *GLibInfo) FilterUpdatePrev(ts int, gpsx float64, gpsy float64) {
	fmt.Println(ts, tt.PrevLoc.Timestamp)
	if (ts - tt.PrevLoc.Timestamp) > 5 {
		tt.PrevLoc.Lati = gpsx
		tt.PrevLoc.Long = gpsy
		tt.PrevLoc.Timestamp = ts
		fmt.Println("update Prev")
	}

}

func (tt *GLibInfo) FilterCalDistance(locx float64, locy float64, tarx float64, tary float64) float64 {
	//	fmt.Println("currtt loc", tt.CurrentLoc.Long, tt.CurrentLoc.Lati, tarx, tary)
	//	fmt.Println("cal distance", CalDistance(tt.CurrentLoc.Lati, tt.CurrentLoc.Long, tarx, tary))
	return CalDistance(locx, locy, tarx, tary)
}

func CalDistance(lati1 float64, long1 float64, lati2 float64, long2 float64) float64 {
	d2r := 0.0174532925199433
	dlong := (long2 - long1) * d2r
	dlat := (lati2 - lati1) * d2r
	a := math.Pow(math.Sin(dlat/2.0), 2) + math.Cos(lati1*d2r)*math.Cos(lati2*d2r)*math.Pow(math.Sin(dlong/2.0), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	d := 6373 * c
	//return meter
	return d * 1000.
}

func (tt *GLibInfo) ShowData(latidata float64, longdata float64) {
	//	if tt
	fmt.Println("initial testnum", tt.Testnum)
	fmt.Println("testnum put in client", tt.Testnum)
	fmt.Println("lati and long:", latidata, longdata)
	fmt.Println("show testinfo data")
	tt.Testnum = 77.88
	fmt.Println("show Gpsloc", tt.Gpsloc)
}
func (tt *GLibInfo) GetData() float64 {

	return tt.Testnum

}

func NewGLib() *GLibInfo {
	a := &GLibInfo{}
	return a
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
