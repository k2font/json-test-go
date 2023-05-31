package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// ネストしていないオブジェクトの場合
type Bird struct {
	Species     string
	Description string
}

// ネストしているオブジェクトの場合
type Dimensions struct {
	Height int
	Width  int
}
type dBird struct {
	Species     string `json:"birdType"`
	Description string `json:"desc"`
}

//

func main() {
	birdJson := `{"species": "pigeon","description": "likes to perch on rocks"}`
	var bird Bird

	json.Unmarshal([]byte(birdJson), &bird)

	fmt.Printf("Species: %s, Description: %s\n", bird.Species, bird.Description)

	jsonData, err := ioutil.ReadFile("sample2.json")
	if err != nil {
		log.Fatal(err)
	}

	birdsJson := `{"birdType": "pigeon","desc": "likes to perch on rocks"}`
	var birds2 dBird
	json.Unmarshal([]byte(birdsJson), &birds2)
	fmt.Printf("%+v", birds2)

	var birds []Bird
	json.Unmarshal([]byte(jsonData), &birds)
	fmt.Printf("Birds: %+v\n", birds)

	// JSONの構造が予めわからないときは、Mapに突っ込むのがよさそう
	d := `{"birds":{"pigeon":"likes to perch on rocks","eagle":"bird of prey"},"animals":"none"}`
	var result map[string]any
	json.Unmarshal([]byte(d), &result)

	birds3 := result["birds"].(map[string]any) // rangeで取り扱うために型をキャスト

	for i, v := range birds3 {
		fmt.Println(i, v.(string))
	}

	// JSON文字列に変換したいならMarshal
	pigeion := &Bird{
		Species:     "Pigeon",
		Description: "likes to eat seed",
	}
	data, _ := json.Marshal(pigeion)
	fmt.Println(string(data)) // {"Species":"Pigeon","Description":"likes to eat seed"}
}
