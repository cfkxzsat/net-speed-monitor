package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"
)

const (
	rPath        string = "/sys/class/net/"
	recievePath  string = "/statistics/rx_bytes"
	transmitPath string = "/statistics/tx_bytes"
)

//up/downs[x][0]: prev transmit/recieve bytes
//up/downs[x][1]: current transmit/recieve bytes
var ups, downs [][]int

func main() {

	netNames := getNetNames()

	//init values
	for i, netName := range netNames {
		arr := make([]int, 2)
		ups = append(ups, arr)
		downs = append(downs, arr)
		var err error
		ups[i][0], err = getUp(netName)
		if err != nil {
			log.Fatal(err)
		}
		downs[i][0], err = getDown(netName)
		if err != nil {
			log.Fatal(err)
		}
	}

	//get data per second
	ticker := time.NewTicker(time.Second)
	for {
		<-ticker.C
		//clean screen
// 		cmd := exec.Command("clear")
// 		cmd.Stdout = os.Stdout
// 		cmd.Run()

		for i, netName := range netNames {
			detectNetSpeed(i, netName)
		}
	}
}

func getNetNames() (names []string) {
	nets, err := ioutil.ReadDir(rPath)
	if err != nil {
		log.Fatal(err)
	}
	for _, netName := range nets {
		names = append(names, netName.Name())
	}
	return names
}

func getUp(netName string) (int, error) {
	b, err := ioutil.ReadFile(rPath + netName + transmitPath)
	if err != nil {
		log.Fatal(err)
	}
	b = b[:len(b)-1]
	return strconv.Atoi(string(b))
}
func getDown(netName string) (int, error) {
	b, err := ioutil.ReadFile(rPath + netName + recievePath)
	if err != nil {
		log.Fatal(err)
	}
	b = b[:len(b)-1]
	return strconv.Atoi(string(b))
}

func detectNetSpeed(i int, netName string) {
	var err error
	ups[i][1], err = getUp(netName)
	if err != nil {
		log.Fatal(err)
	}
	downs[i][1], err = getDown(netName)
	if err != nil {
		log.Fatal(err)
	}
	up := withUnit(ups[i][1] - ups[i][0])
	down := withUnit(downs[i][1] - downs[i][0])

	fmt.Printf("%8s\tupload:%5sB/s\tdownload:%5sB/s\r",netname,up,down)
	
	ups[i][0] = ups[i][1]
	downs[i][0] = downs[i][1]

}

func withUnit(b int) (res string) {
	v := float64(b)
	unit := ""
	count := 0

	for v >= 1024 {
		count++
		v = v / 1024
	}

	switch count {
	case 0:
	case 1:
		unit = "K"
	case 2:
		unit = "M"
	case 3:
		unit = "G"
	case 4:
		unit = "T"
	default:
		log.Fatal("can't handle so much data")

	}

	return strconv.FormatFloat(v, 'g', 5, 32) + unit
}
