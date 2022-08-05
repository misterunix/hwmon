package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

//var Hwmon map[string]map[string]string
//var HwmonName map[string]string

/* main struct holding the hwmon names,vales and paths */
type HwMon struct {
	// map[path]map[name]value
	Values map[string]map[string]string
	//BaseNames map[string]string // map of paths to names
	BasePath string
}

func (h *HwMon) Init(path string) error {

	h.Values = make(map[string]map[string]string, 0)

	h.BasePath = path
	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	if len(files) == 0 {
		return errors.New("no hwmon devices found")
	}

	for _, file := range files {
		tpath := path + "/" + file.Name()
		h.Values[tpath] = map[string]string{}
	}

	return nil
}

func (h *HwMon) GetValues() error {

	for d := range h.Values {
		tmpd, err := os.ReadDir(d)
		if err != nil {
			return err
		}
		for _, newf := range tmpd {
			if strings.Contains(newf.Name(), "_") {
				tfpath := d + "/" + newf.Name()
				tv, err := quickRead(tfpath)
				if err != nil {
					tv = "error"
					continue
				}
				h.Values[d][newf.Name()] = tv
			}

		}
	}
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
