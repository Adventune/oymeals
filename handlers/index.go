package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Juvenes restaurants
var juvenesUrls = [][]string{
	{
		"https://fi.jamix.cloud/apps/menuservice/rest/haku/menu/93077/70?lang=fi",
		"Kerttu",
		"Voltti",
	},
	{
		"https://fi.jamix.cloud/apps/menuservice/rest/haku/menu/93077/49?lang=fi",
		"Mara",
	},
	{
		"https://fi.jamix.cloud/apps/menuservice/rest/haku/menu/93077/69?lang=fi",
		"Foobar",
	},
}

// Uniresta restaurants
var powerestaUrls = [][]string{
	{
		"https://api.fi.poweresta.com/publicmenu/dates/uniresta/lipasto/?menu=ravintolalipasto&dates=",
		"Lipasto & Julinia",
	},
	{
		"https://api.fi.poweresta.com/publicmenu/dates/uniresta/preludi/?menu=ravintolapreludi&dates=",
		"Preludi",
	},
	{
		"https://api.fi.poweresta.com/publicmenu/dates/uniresta/pekuri/?menu=ravintolapekuri&dates=",
		"Pekuri",
	},
}

var SORT_ORDER = []string{
	"Lipasto & Julinia",
	"Mara",
	"Kerttu",
	"Voltti",
	"Foobar",
	"Pekuri",
	"Preludi",
}

func Index(c *fiber.Ctx) error {
	title := "OYMeals"
	if os.Getenv("TITLE") != "" {
		title = os.Getenv("TITLE")
	}

	date := time.Now()
	if date.Hour() >= 17 {
		date = date.Add(24 * time.Hour)
	}

	// Render the page
	return c.Render("index.tmpl", fiber.Map{
		"Title":       title,
		"Restaurants": Data(date),
		"Date":        date.Format("02.01.2006"),
	})
}

func Data(date time.Time) []Restaurant {
	var restaurants []Restaurant
	// Fetch menus from Juvenes and Uniresta
	for _, url := range juvenesUrls {
		response := fetchJuvenes(date, url[0], url[1:]...)
		restaurants = append(restaurants, response...)
	}

	for _, url := range powerestaUrls {
		response := fetchPoweresta(date, url[0], url[1])
		restaurants = append(restaurants, response...)
	}

	sortedRestaurants := make([]Restaurant, len(restaurants))
	for _, restaurant := range restaurants {
		for i, name := range SORT_ORDER {
			if restaurant.Name == name {
				sortedRestaurants[i] = restaurant
				break
			}
		}
	}

	return sortedRestaurants
}

// Fetch menus from Juvenes
func fetchJuvenes(date time.Time, url string, names ...string) []Restaurant {
	var response JuvenesResponse
	var restaurants []Restaurant

	// Get the menu data
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return restaurants
	}
	defer res.Body.Close()
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return restaurants
	}
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		fmt.Println(err)
		return restaurants
	}

	currentDate := date.Format("20060102")

	// Parse the menu data
	for i, restaurant := range response[0].MenuTypes {
		if i >= len(names) {
			break
		}

		var menus []Menu
		// Loop through the menus
		for _, menu := range restaurant.Menus {
			// Find the menu for the current date
			for _, day := range menu.Days {
				if strconv.Itoa(day.Date) == currentDate {
					// Loop through the meal options
					for _, mealOption := range day.Mealoptions {
						var items []Item
						// Check if the menu is lunch
						if strings.Contains(strings.ToLower(mealOption.Name), "lounas") ||
							strings.Contains(strings.ToLower(mealOption.Name), "classic") {
							for _, menuItem := range mealOption.MenuItems {
								items = append(items, Item{Name: menuItem.Name})
							}
							// Append the menu to the list
							menus = append(menus, Menu{
								Name:  mealOption.Name,
								Items: items,
							})
						}
					}
				}
			}

			// Append the restaurant to the list
			restaurants = append(restaurants, Restaurant{
				Name:  names[i],
				Menus: menus,
			})
		}
	}

	return restaurants
}

// Fetch menus from Uniresta
func fetchPoweresta(date time.Time, url, name string) []Restaurant {
	var response PowerestaResponse
	var restaurants []Restaurant

	currentDate := date.Format("2006-01-02")

	// Get the menu data
	res, err := http.Get(url + currentDate)
	if err != nil {
		fmt.Println(err)
		return restaurants
	}
	defer res.Body.Close()
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return restaurants
	}
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		fmt.Println(err)
		return restaurants
	}

	var menus []Menu
	// Loop through the menus
	for _, menu := range response[0].Data.MealOptions {
		var items []Item
		// Get all menu items
		for _, row := range menu.Rows {
			for _, name := range row.Names {
				// Check if the item name is in Finnish
				if name.Language == "fi" {
					items = append(items, Item{Name: name.Name})
				}
			}
		}
		for _, name := range menu.Names {
			// Get the menu name in Finnish
			if name.Language == "fi" {
				// Check if the menu is lunch
				if strings.Contains(strings.ToLower(name.Name), "lounas") {
					// Append the menu to the list
					menus = append(menus, Menu{
						Name:  name.Name,
						Items: items,
					})
					break
				}
			}
		}
	}
	// Append the restaurant to the list
	restaurants = append(restaurants, Restaurant{
		Name:  name,
		Menus: menus,
	})

	return restaurants
}
