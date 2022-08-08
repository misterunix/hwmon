package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type Sensor struct {
	Head  string
	Tail  string
	Value string
	Name  string
	Path  string
}

var Sensors []Sensor

func Init(path string) error {

	// Get the list of hwmon directories
	basedirs, err := GetDirEntries(path)
	if err != nil {
		return err
	}

	// Loop through the hwmon directories
	for _, basedir := range basedirs {
		var name string
		var head string
		var tail string
		var value string
		//get files in hwmon directory
		files, err := GetDirEntries(basedir)
		if err != nil {
			continue
		}

		if len(files) == 0 {
			continue
		}

		//loop through files in hwmon directory
		for _, y := range files {
			if !strings.Contains(y, "_") {
				continue
			}

			//get name of sensor
			split := strings.Split(y, "_")
			name = split[0]
			tail = split[1]
			//value = quickRead()
			z := Sensor{
				Head:  head,
				Tail:  tail,
				Value: value,
				Name:  name,
				Path:  basedir,
			}
			Sensors = append(Sensors, z)

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

	err := Init("/sys/class/hwmon")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(Sensors)

}
