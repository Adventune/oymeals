package handlers

type Restaurant struct {
	Name  string
	Menus []Menu
}

type Menu struct {
	Name  string
	Items []Item
}

type Item struct {
	Name string
}

type JuvenesResponse []struct {
	KitchenName string `json:"kitchenName"`
	KitchenID   int    `json:"kitchenId"`
	Address     string `json:"address"`
	City        string `json:"city"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Info        string `json:"info"`
	MenuTypes   []struct {
		MenuTypeID   int    `json:"menuTypeId"`
		MenuTypeName string `json:"menuTypeName"`
		Menus        []struct {
			MenuName           string `json:"menuName"`
			MenuAdditionalName string `json:"menuAdditionalName"`
			MenuID             int    `json:"menuId"`
			Days               []struct {
				Date        int `json:"date"`
				Weekday     int `json:"weekday"`
				Mealoptions []struct {
					Name        string `json:"name"`
					OrderNumber int    `json:"orderNumber"`
					ID          int    `json:"id"`
					MenuItems   []struct {
						Name        string `json:"name"`
						OrderNumber int    `json:"orderNumber"`
						PortionSize int    `json:"portionSize"`
						Diets       string `json:"diets"`
						Ingredients string `json:"ingredients"`
					} `json:"menuItems"`
				} `json:"mealoptions"`
			} `json:"days"`
		} `json:"menus"`
	} `json:"menuTypes"`
}

type PowerestaResponse []struct {
	AllSuccessful bool   `json:"allSuccessful"`
	Date          string `json:"date"`
	TenantFound   bool   `json:"tenantFound"`
	SiteFound     bool   `json:"siteFound"`
	IsMenuForDate bool   `json:"isMenuForDate"`
	Data          struct {
		Date        string `json:"date"`
		MenuID      string `json:"menuId"`
		SiteID      string `json:"siteId"`
		Version     int    `json:"version"`
		MealOptions []struct {
			Rows []struct {
				Diets []struct {
					Diets      []string `json:"diets"`
					Language   string   `json:"language"`
					DietShorts []string `json:"dietShorts"`
				} `json:"diets"`
				Names []struct {
					Name     string `json:"name"`
					Language string `json:"language"`
				} `json:"names"`
				Co2Value struct {
					Ranges []struct {
						Max     int    `json:"max"`
						Min     int    `json:"min"`
						Color   string `json:"color"`
						NameKey string `json:"nameKey"`
					} `json:"ranges"`
					ValuePerPortionInKg string `json:"valuePerPortionInKg"`
				} `json:"co2Value"`
				Allergens []struct {
					Language  string   `json:"language"`
					Allergens []string `json:"allergens"`
				} `json:"allergens"`
				Ingredients []struct {
					Language    string `json:"language"`
					Ingredients string `json:"ingredients"`
				} `json:"ingredients"`
				PublicInfos []struct {
					Language string `json:"language"`
				} `json:"publicInfos"`
				NutritiveItem struct {
					Factors []struct {
						Values []struct {
							Value    string `json:"value"`
							Language string `json:"language"`
							Value100 string `json:"value100"`
						} `json:"values"`
						FactorOrder          int         `json:"factorOrder"`
						EnergyContent        interface{} `json:"energyContent"`
						MainGroupCode        string      `json:"mainGroupCode"`
						CompareProportion    string      `json:"compareProportion,omitempty"`
						CalculatePercentage  bool        `json:"calculatePercentage"`
						NutritiveFactorNames []struct {
							Name     string `json:"name"`
							Language string `json:"language"`
						} `json:"nutritiveFactorNames"`
						EnergyProportion string `json:"energyProportion,omitempty"`
						Percentage       string `json:"percentage,omitempty"`
					} `json:"factors"`
					PortionSize                   string `json:"portionSize"`
					PortionUnitOfMeasurementNames []struct {
						Name     string `json:"name"`
						Language string `json:"language"`
					} `json:"portionUnitOfMeasurementNames"`
				} `json:"nutritiveItem"`
			} `json:"rows"`
			Names []struct {
				Name     string `json:"name"`
				Language string `json:"language"`
			} `json:"names"`
		} `json:"mealOptions"`
	} `json:"data"`
}
