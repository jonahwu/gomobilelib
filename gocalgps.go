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
type Gps2dLocVel struct {
	Long float64
	Lati float64
	Vel  float64
}

type NearCamLocVel struct {
	Distance  float64
	Vel       float64
	Timestamp int
}

type Gps4dLoc struct {
	Long      float64
	Lati      float64
	Timestamp int
}

type ListGps2dLocTim struct {
	Loc       []Gps2dLocVel
	Timestamp int
}

type GLibInfo struct {
	Gpsloc          *Gps2dLoc // for gmob used for notification
	Testnum         float64
	PrevLoc         Gps4dLoc
	CurrentLoc      Gps4dLoc
	TargetCameraLoc ListGps2dLocTim
	CreateLoc       Gps4dLoc
	NearestCamera   NearCamLocVel
}

func NewGp2dLoc() *Gps2dLoc {
	a := Gps2dLoc{}
	return &a
}

func (tt *GLibInfo) InitState(id string, ts int, locx float64, locy float64) {
	//	tarx := 25.080223
	//	tary := 121.697908

	// set up fake camra
	ta := Gps2dLocVel{}
	ta.Lati = 25.080223
	ta.Long = 121.697908
	ta.Vel = 50.0
	tt.TargetCameraLoc.Loc = append(tt.TargetCameraLoc.Loc, ta)

	//	25.0674 121.66832
	ta.Lati = 25.0674
	ta.Long = 121.66832
	ta.Vel = 100.0
	tt.TargetCameraLoc.Loc = append(tt.TargetCameraLoc.Loc, ta)

	tt.TargetCameraLoc.Timestamp = ts
	fmt.Println("show all initial camera", tt.TargetCameraLoc, len(tt.TargetCameraLoc.Loc))
	// set up Prev status
	tt.PrevLoc.Lati = locx
	tt.PrevLoc.Long = locy
	tt.PrevLoc.Timestamp = ts
	fmt.Println("show the initial  PrevLoc", tt.PrevLoc)
	//
	tt.CurrentLoc.Lati = locx
	tt.CurrentLoc.Long = locy
	tt.CurrentLoc.Timestamp = ts

}

func (tt *GLibInfo) UpdateCamera() {
}

func (tt *GLibInfo) Start(ts int, locx float64, locy float64) (float64, int, float64) {
	tt.UpdateCamera()
	dist, flag := tt.GLibFilter(ts, locx, locy)
	fmt.Println("start now", ts)
	vel := tt.NearestCamera.Vel
	return dist, flag, vel

}

//func (tt *GLibInfo) FilterDistance(tarx float64, tary float64) (float64, int) {
func (tt *GLibInfo) FilterDistance() (float64, int) {
	//now we can only serve one camera
	tt.NearestCamera.Distance = 9999999.88
	tt.NearestCamera.Vel = 9999999.88
	fmt.Println("show near camera", tt.NearestCamera.Distance, tt.NearestCamera.Vel)
	for i := 0; i < len(tt.TargetCameraLoc.Loc); i++ {
		fmt.Println("-----------calculate on camera:", i)
		tarx := tt.TargetCameraLoc.Loc[i].Lati
		tary := tt.TargetCameraLoc.Loc[i].Long

		locx := tt.CurrentLoc.Lati
		locy := tt.CurrentLoc.Long
		distcurr := tt.FilterCalDistance(locx, locy, tarx, tary)
		locpx := tt.PrevLoc.Lati
		locpy := tt.PrevLoc.Long
		distprev := tt.FilterCalDistance(locpx, locpy, tarx, tary)
		fmt.Println("dist current, dist prev", distcurr, distprev)
		// where 30 is faraway distance
		if distprev < distcurr && distcurr > 30.0 {
			//return distcurr, 0
			fmt.Println("not put in candidate")
		} else {
			// calculate Nearest Camera that is satisfy condition
			if distcurr < tt.NearestCamera.Distance {
				fmt.Println("--------  run into fit camera ----------")
				tt.NearestCamera.Distance = distcurr
				tt.NearestCamera.Timestamp = tt.TargetCameraLoc.Timestamp
				tt.NearestCamera.Vel = tt.TargetCameraLoc.Loc[i].Vel
			}

			//return distcurr, 1
		}
	}
	flag := 0
	if tt.NearestCamera.Distance < 300.0 {
		flag = 1
	}
	return tt.NearestCamera.Distance, flag
	//	return distcurr, 0
}

func (tt *GLibInfo) GLibFilter(ts int, gpsx float64, gpsy float64) (float64, int) {
	//	tarx := 25.080223
	//	tary := 121.697908
	tt.UpdateCurrent(ts, gpsx, gpsy)

	//	tt.FilterInitMap()
	dist, runflag := tt.FilterDistance()
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
