package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type CocktailBartender struct {
	url string
}

type DrinkInstructions struct {
	Recipe string `json:"strInstructions"`
}

type DrinksResponsePayload struct {
	DrinkRecipes []DrinkInstructions `json:"drinks"`
}

func NewCocktailBartender(url string) CocktailBartender {
	return CocktailBartender{url: url}
}

func (c *CocktailBartender) Start() {
	var input string
	fmt.Println("Enter drink")
	fmt.Scanln(&input)

	for input != "nothing" {
		url := c.url + "?s=" + input
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatal(err)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}

		if resp.StatusCode != http.StatusOK {
			fmt.Println(resp.StatusCode)
			break
		}

		payload := DrinksResponsePayload{}
		json.NewDecoder(resp.Body).Decode(&payload)

		if len(payload.DrinkRecipes) != 0 {
			recipe := payload.DrinkRecipes[0]
			splittedRecipe := strings.Split(recipe.Recipe, ".")

			for _, line := range splittedRecipe {
				fmt.Println(line)
			}
		} else {
			fmt.Println("We don't have your drink in the menu. Would you like another one?")
		}

		fmt.Println("Enter drink")
		fmt.Scanln(&input)
	}
}

func main() {
	cocktail := NewCocktailBartender("https://www.thecocktaildb.com/api/json/v1/1/search.php")
	cocktail.Start()
}
