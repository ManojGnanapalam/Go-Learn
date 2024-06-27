package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type users struct {
	Name     string `json:"Username"`
	Email    string
	Password string   `json:"-"`
	Role     []string `json:"Role,omitempty"`
}

func EncodeJson() {
	appUsers := []users{
		{"viper", "viper@mail", "2134bb", []string{"user", "editer"}},
		{"spark", "spark@mail", "mmkkd", []string{"user", "admin"}},
		{"keethi", "keethi@mail", "9922o", nil},
	}

	data, err := json.MarshalIndent(appUsers, "", "\t")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", data)
}

func DecodeJson() {
	jsonData := []byte(`
	{
        "Username": "keethi",
        "Email": "keethi@mail",
        "Password": "9922o",
        "Role": null
    }
	`)

	var appuser users

	checkValid := json.Valid(jsonData)

	if checkValid {
		json.Unmarshal(jsonData, &appuser)
		fmt.Printf("%#v\n", appuser)
	} else {
		fmt.Println("Json was not valid")
	}

	var goodLookData map[string]interface{}
	json.Unmarshal(jsonData, &goodLookData)
	fmt.Printf("%v#\n", goodLookData)

	for k, v := range goodLookData {
		fmt.Printf("Key is %v and value is %v and Type is : %T \n", k, v, v)
	}
}

func main() {
	fmt.Println("Here were convert")
	EncodeJson()
	DecodeJson()
}
