package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

//var Hwmon map[string]map[string]string
//var HwmonName map[string]string

type Sensors struct {
	Path  string
	Head  string
	Tail  string
	Value string
}

/* main struct holding the hwmon names,vales and paths */
type HwMon struct {
	BasePath string
	Sensor   []Sensors
}

func (h *HwMon) Init(path string) error {

	// Get the list of hwmon directories
	basedirs, err := GetDirEntries(path)
	if err != nil {
		return err
	}

	// Loop through the directories
	for _, d := range basedirs {
		// Get the list of files in the directory
		tmpfilelist, err := GetDirEntries(d)
		if err != nil {
			return err
		}

		// Loop through the files
		for _, f := range tmpfilelist {

			// Get the file name
			if !strings.Contains(f, "_") {
				continue
			}

			split := strings.Split(f, "_")
			y := Sensors{
				Head: split[0],
				Tail: split[1],
			}
			fmt.Println(d, y)
			h.Values[d][y] = "0"

		}

	}

	return nil
}

func GetDirEntries(path string) ([]string, error) {

	var filelist []string

	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		filelist = append(filelist, path+"/"+f.Name())
	}

	if len(filelist) == 0 {
		return nil, errors.New("no files found")
	}

	return filelist, nil

}

func (h *HwMon) GetValues() error {

	return nil
}

func quickRead(path string) (string, error) {
	tmpB, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	tmpS := strings.ToLower(string(tmpB))
	tmpS = tmpS[:len(tmpS)-1]
	tmpS = strings.TrimSpace(tmpS)
	return tmpS, nil
}

func main() {

	os.Exit(0)

	var h HwMon
	err := h.Init("/sys/class/hwmon")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = h.GetValues()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for p, a := range h.Values {
		for k, v := range a {
			fmt.Printf("%s %s %s\n", p, k, v)
		}

	}

	//fmt.Println(h)

}
