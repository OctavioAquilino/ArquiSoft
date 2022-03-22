package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//	https://api.mercadolibre.com/sites/MLA/search?q=Motrola

type Categories []Category
type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	cats, err := GetCategories("MLA")
	if err != nil {
		fmt.Println("Error: ", err.Error())
		//exit(1)
		return
	}
	fmt.Println("Las categorias son: ")

	for i := 0; i < len(cats); i++ {

		fmt.Println(cats[i].Name)

	}

}
func GetCategories(siteID string) (Categories, error) {
	response, err1 := http.Get("https://api.mercadolibre.com/sites/MLA/categories")
	if err1 != nil {
		fmt.Println("Error: ", err1.Error())
		//exit(1)
		return nil, nil
	}
	//data, err := response.Bytes()
	bytes, err2 := ioutil.ReadAll(response.Body)
	if err2 != nil {
		fmt.Println("Error: ", err2.Error())
		//exit(1)
		return nil, nil
	}
	var cats Categories
	err := json.Unmarshal(bytes, &cats)

	if err != nil {
		fmt.Println("Error: ", err.Error())
		//exit(1)
		return nil, nil
	}

	return cats, nil

}
