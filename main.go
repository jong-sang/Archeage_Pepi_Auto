package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/Archeage_Pepi_Auto/pkg"
)

func main() {
	b, err := ioutil.ReadFile("./Account.json")
	if err != nil {
		log.Fatal(err)
	}

	var data []pkg.Account // JSON 문서의 데이터를 저장할 구조체 슬라이스 선언

	json.Unmarshal(b, &data)

	err = pkg.Login(data[0].AccType, data[0].ID, data[0].PW)
	if err != nil {
		log.Fatal(err)
	}
}
