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
	fmt.Println("Las categorias son: ", cats)

}
func GetCategories(siteID string) (Categories, error) {
	response, err1 := http.Get("https://api.mercadolibre.com/sites/MLA/search?q=Motrola")
	if err1 != nil {
		fmt.Println("Error: ", err1.Error())
		//exit(1)
		return nil, nil
	}
	bytes, err2 := ioutil.ReadAll(response.Body)
	if err2 != nil {
		fmt.Println("Error: ", err2.Error())
		//exit(1)
		return nil, nil
	}
	var categories Categories
	err := json.Unmarshal(bytes, &categories)

	if err != nil {
		fmt.Println("Error: ", err.Error())
		//exit(1)
		return nil, nil
	}

	return categories, nil

}
