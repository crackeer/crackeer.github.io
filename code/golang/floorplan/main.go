package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

}

func calculateFloorPlan(data string) (int64, float64) {
	// floorplanStruct
	type floorplanStruct struct {
		Floorplans []struct {
			Areas []struct {
				SizeWithoutLine float64 `json:"sizeWithoutLine"`
			}
		} `json:"floorplans"`
		Property struct {
			ScanSize float64 `json:"scanSize"`
		} `json:"property"`
	}

	dataStruct := &floorplanStruct{}
	if err := json.Unmarshal([]byte(data), dataStruct); err != nil {
		return -1, -1
	}

	if dataStruct.Property.ScanSize < 1 {
		var areaSize float64
		for _, item := range dataStruct.Floorplans {
			for _, area := range item.Areas {
				areaSize += area.SizeWithoutLine
			}
		}
		dataStruct.Property.ScanSize = areaSize / 1e6
	}

	return int64(len(dataStruct.Floorplans)), dataStruct.Property.ScanSize

}

func getFloorplanData(workID int64) (string, error) {

	url := fmt.Sprintf("http://i.svc.gapi.realsee.com/svc/work/floorplan/detail.json?work_id=%d", workID)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return "", err
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))

}
