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
	Vel  int
}

type NearCamLocVel struct {
	Distance    float64
	Vel         int
	Timestamp   int
	Flag        int
	OldDistance float64
}

type OldNearCamLocVel struct {
	Distance       float64
	Vel            int
	Timestamp      int
	Flag           int
	OldDistance    float64
	DenoiseCounter int
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
	Gpsloc           *Gps2dLoc // for gmob used for notification
	Testnum          float64
	PrevLoc          Gps4dLoc
	CurrentLoc       Gps4dLoc
	TargetCameraLoc  ListGps2dLocTim
	CreateLoc        Gps4dLoc
	NearestCamera    NearCamLocVel
	OldNearestCamera OldNearCamLocVel
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
	ta.Vel = 50
	tt.TargetCameraLoc.Loc = append(tt.TargetCameraLoc.Loc, ta)

	//	25.0674 121.66832
	ta.Lati = 25.114392
	ta.Long = 121.685341
	ta.Vel = 50
	tt.TargetCameraLoc.Loc = append(tt.TargetCameraLoc.Loc, ta)

	ta.Lati = 25.112894
	ta.Long = 121.690519
	ta.Vel = 40
	tt.TargetCameraLoc.Loc = append(tt.TargetCameraLoc.Loc, ta)

	ta.Lati = 25.078060
	ta.Long = 121.691110
	ta.Vel = 50
	tt.TargetCameraLoc.Loc = append(tt.TargetCameraLoc.Loc, ta)

	ta.Lati = 25.087720
	ta.Long = 121.695398
	ta.Vel = 50
	tt.TargetCameraLoc.Loc = append(tt.TargetCameraLoc.Loc, ta)

	ta.Lati = 25.101782
	ta.Long = 121.696034
	ta.Vel = 50
	tt.TargetCameraLoc.Loc = append(tt.TargetCameraLoc.Loc, ta)

	ta.Lati = 25.121502
	ta.Long = 121.697590
	ta.Vel = 80
	tt.TargetCameraLoc.Loc = append(tt.TargetCameraLoc.Loc, ta)

	ta.Lati = 25.121832
	ta.Long = 121.697869
	ta.Vel = 80
	tt.TargetCameraLoc.Loc = append(tt.TargetCameraLoc.Loc, ta)

	ta.Lati = 25.092500
	ta.Long = 121.699440
	ta.Vel = 50
	tt.TargetCameraLoc.Loc = append(tt.TargetCameraLoc.Loc, ta)

	ta.Lati = 25.080700
	ta.Long = 121.701497
	ta.Vel = 50
	tt.TargetCameraLoc.Loc = append(tt.TargetCameraLoc.Loc, ta)

	ta.Lati = 25.090313
	ta.Long = 121.706649
	ta.Vel = 50
	tt.TargetCameraLoc.Loc = append(tt.TargetCameraLoc.Loc, ta)

	ta.Lati = 25.097341
	ta.Long = 121.716191
	ta.Vel = 50
	tt.TargetCameraLoc.Loc = append(tt.TargetCameraLoc.Loc, ta)

	ta.Lati = 25.103060
	ta.Long = 121.725280
	ta.Vel = 50
	tt.TargetCameraLoc.Loc = append(tt.TargetCameraLoc.Loc, ta)

	ta.Lati = 25.111110
	ta.Long = 121.735000
	ta.Vel = 50
	tt.TargetCameraLoc.Loc = append(tt.TargetCameraLoc.Loc, ta)

	ta.Lati = 25.102713
	ta.Long = 121.754820
	ta.Vel = 80
	tt.TargetCameraLoc.Loc = append(tt.TargetCameraLoc.Loc, ta)

	ta.Lati = 25.102885
	ta.Long = 121.755578
	ta.Vel = 80
	tt.TargetCameraLoc.Loc = append(tt.TargetCameraLoc.Loc, ta)

	ta.Lati = 25.059990
	ta.Long = 121.640210
	ta.Vel = 50
	tt.TargetCameraLoc.Loc = append(tt.TargetCameraLoc.Loc, ta)

	ta.Lati = 25.061110
	ta.Long = 121.642590
	ta.Vel = 50
	tt.TargetCameraLoc.Loc = append(tt.TargetCameraLoc.Loc, ta)

	ta.Lati = 25.061570
	ta.Long = 121.651350
	ta.Vel = 50
	tt.TargetCameraLoc.Loc = append(tt.TargetCameraLoc.Loc, ta)

	ta.Lati = 25.066710
	ta.Long = 121.664570
	ta.Vel = 50
	tt.TargetCameraLoc.Loc = append(tt.TargetCameraLoc.Loc, ta)

	ta.Lati = 25.076990
	ta.Long = 121.671550
	ta.Vel = 50
	tt.TargetCameraLoc.Loc = append(tt.TargetCameraLoc.Loc, ta)

	ta.Lati = 25.065942
	ta.Long = 121.674905
	ta.Vel = 50
	tt.TargetCameraLoc.Loc = append(tt.TargetCameraLoc.Loc, ta)

	ta.Lati = 25.068420
	ta.Long = 121.681260
	ta.Vel = 50
	tt.TargetCameraLoc.Loc = append(tt.TargetCameraLoc.Loc, ta)

	ta.Lati = 25.072780
	ta.Long = 121.683330
	ta.Vel = 50
	tt.TargetCameraLoc.Loc = append(tt.TargetCameraLoc.Loc, ta)
	//大安區
	ta.Lati = 25.057780
	ta.Long = 121.532220
	ta.Vel = 50
	tt.TargetCameraLoc.Loc = append(tt.TargetCameraLoc.Loc, ta)

	ta.Lati = 25.017220
	ta.Long = 121.533330
	ta.Vel = 50
	tt.TargetCameraLoc.Loc = append(tt.TargetCameraLoc.Loc, ta)

	ta.Lati = 25.017220
	ta.Long = 121.533330
	ta.Vel = 50
	tt.TargetCameraLoc.Loc = append(tt.TargetCameraLoc.Loc, ta)

	ta.Lati = 25.030830
	ta.Long = 121.533330
	ta.Vel = 50
	tt.TargetCameraLoc.Loc = append(tt.TargetCameraLoc.Loc, ta)

	ta.Lati = 25.041670
	ta.Long = 121.536940
	ta.Vel = 50
	tt.TargetCameraLoc.Loc = append(tt.TargetCameraLoc.Loc, ta)

	ta.Lati = 25.025280
	ta.Long = 121.537780
	ta.Vel = 70
	tt.TargetCameraLoc.Loc = append(tt.TargetCameraLoc.Loc, ta)

	ta.Lati = 25.027892
	ta.Long = 121.537864
	ta.Vel = 50
	tt.TargetCameraLoc.Loc = append(tt.TargetCameraLoc.Loc, ta)

	ta.Lati = 25.044170
	ta.Long = 121.540560
	ta.Vel = 40
	tt.TargetCameraLoc.Loc = append(tt.TargetCameraLoc.Loc, ta)

	ta.Lati = 25.021670
	ta.Long = 121.543330
	ta.Vel = 50
	tt.TargetCameraLoc.Loc = append(tt.TargetCameraLoc.Loc, ta)

	ta.Lati = 25.024440
	ta.Long = 121.548610
	ta.Vel = 50
	tt.TargetCameraLoc.Loc = append(tt.TargetCameraLoc.Loc, ta)

	ta.Lati = 25.021670
	ta.Long = 121.548890
	ta.Vel = 60
	tt.TargetCameraLoc.Loc = append(tt.TargetCameraLoc.Loc, ta)

	ta.Lati = 25.021670
	ta.Long = 121.548890
	ta.Vel = 60
	tt.TargetCameraLoc.Loc = append(tt.TargetCameraLoc.Loc, ta)

	ta.Lati = 25.017220
	ta.Long = 121.549170
	ta.Vel = 50
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

func (tt *GLibInfo) BeepOrNot(nc NearCamLocVel) bool {
	if nc.Flag == 0 {
		regionBeep := false
		if tt.OldNearestCamera.DenoiseCounter >= 2 {
			tt.OldNearestCamera.Timestamp = nc.Timestamp
			tt.OldNearestCamera.Distance = 0
			tt.OldNearestCamera.DenoiseCounter = 0
		}
		if nc.Timestamp == 0 && tt.OldNearestCamera.Timestamp != 0 {
			tt.OldNearestCamera.DenoiseCounter = tt.OldNearestCamera.DenoiseCounter + 1
		}
		return regionBeep
	} else {

		if nc.Distance < 100 {
			regionBeep := true
			return regionBeep
		}
		fmt.Println("into calculate Region Alarm")
		fmt.Println(tt.OldNearestCamera.Distance, nc.Distance)
		regionBeep := tt.RegionAlarm(tt.OldNearestCamera.Distance, nc.Distance)
		fmt.Println(regionBeep)
		if nc.Timestamp != 0 && tt.OldNearestCamera.Timestamp == 0 {
			regionBeep = true
		}
		if regionBeep {
			tt.OldNearestCamera.Distance = nc.Distance
			tt.OldNearestCamera.Flag = nc.Flag
			tt.OldNearestCamera.Timestamp = nc.Timestamp
		}
		return regionBeep

	}
	return false
}

//func (tt *GLibInfo) Start(ts int, locx float64, locy float64) (float64, int, int) {
func (tt *GLibInfo) Start(ts int, locx float64, locy float64) string {
	tt.UpdateCamera()
	NearestCamera := tt.GLibFilter(ts, locx, locy)

	fmt.Println("start now", ts)
	beep := tt.BeepOrNot(NearestCamera)
	//	vel := NearestCamera.Vel
	jm := make(map[string]interface{})
	jm["dist"] = NearestCamera.Distance
	jm["flag"] = NearestCamera.Flag
	jm["vel"] = NearestCamera.Vel
	jm["beep"] = beep
	ret, _ := json.Marshal(jm)
	//return dist, flag, vel
	return string(ret)

}

func isApproach(dx float64, dx1 float64) bool {
	if dx <= dx1 {
		return true
	} else {
		return false
	}
	return true
}

//func (tt *GLibInfo) FilterDistance(tarx float64, tary float64) (float64, int) {
func (tt *GLibInfo) FilterDistance() NearCamLocVel {
	//now we can only serve one camera
	tt.NearestCamera.Distance = 9999999.88
	tt.NearestCamera.Vel = 8888888
	tt.NearestCamera.Timestamp = 0
	tt.NearestCamera.Flag = 0

	fmt.Println("show near camera", tt.NearestCamera.Distance, tt.NearestCamera.Vel)
	for i := 0; i < len(tt.TargetCameraLoc.Loc); i++ {
		//fmt.Println("-----------calculate on camera:", i)
		tarx := tt.TargetCameraLoc.Loc[i].Lati
		tary := tt.TargetCameraLoc.Loc[i].Long

		locx := tt.CurrentLoc.Lati
		locy := tt.CurrentLoc.Long
		distcurr := tt.FilterCalDistance(locx, locy, tarx, tary)
		locpx := tt.PrevLoc.Lati
		locpy := tt.PrevLoc.Long
		distprev := tt.FilterCalDistance(locpx, locpy, tarx, tary)
		//fmt.Println("dist current, dist prev", distcurr, distprev)
		// where 30 is faraway distance
		if distcurr < 50.0 {
			//return distcurr, 0
			fmt.Println("--------  run into fit camera ----------")
			tt.NearestCamera.Distance = distcurr
			tt.NearestCamera.Timestamp = tt.TargetCameraLoc.Timestamp
			tt.NearestCamera.Vel = tt.TargetCameraLoc.Loc[i].Vel
			tt.NearestCamera.Flag = 2
		} else {
			// calculate Nearest Camera that is satisfy condition
			// set up arround camera
			//if (distcurr < distprev) && distcurr < 1000.0 {
			if isApproach(distcurr, distprev) && distcurr < 300.0 {
				if distcurr < tt.NearestCamera.Distance {
					fmt.Println("--------  run into fit camera ----------")
					tt.NearestCamera.Distance = distcurr
					tt.NearestCamera.Timestamp = tt.TargetCameraLoc.Timestamp
					tt.NearestCamera.Vel = tt.TargetCameraLoc.Loc[i].Vel
					tt.NearestCamera.Flag = 1
				}
			}

			//return distcurr, 1
		}
	}
	//end of for loop end redefine the stauts to more emergency
	if tt.NearestCamera.Distance < 100.0 {
		tt.NearestCamera.Flag = 2
	}
	//return tt.NearestCamera.Distance, tt.NearestCamera.Flag
	return tt.NearestCamera
	//	return distcurr, 0
}

func (tt *GLibInfo) GLibFilter(ts int, gpsx float64, gpsy float64) NearCamLocVel {
	//	tarx := 25.080223
	//	tary := 121.697908
	tt.UpdateCurrent(ts, gpsx, gpsy)

	//	tt.FilterInitMap()
	NearestCamera := tt.FilterDistance()
	tt.FilterUpdatePrev(ts, gpsx, gpsy)
	return NearestCamera
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

// related to beep while dr1=200 dr2=150 not alram but dr2<=100 will alarm
// same as 90, 85 will not alarm, <80 will alarm
func (tt *GLibInfo) RegionAlarm(drold float64, drnew float64) bool {
	//_ := math.Log10(drold)
	b := math.Log10(drold)
	dr := math.Abs(drnew - drold)
	ldr := math.Log10(dr)
	fmt.Println("drold and ldr", b, ldr, int(b), int(ldr))
	if int(b) == int(ldr) {
		return true
	} else {
		return false
	}
	return false
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
