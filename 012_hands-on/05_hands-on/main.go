package main

import (
	"log"
	"os"
	"text/template"
)

type menu []meal

type meal struct {
	Name  string
	Items []item
}

type item struct {
	Name, Description string
	Price             float64
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	m := menu{
		meal{
			Name: "Breakfast",
			Items: []item{
				item{
					Name:        "Oatmeal",
					Description: "heart healthy",
					Price:       2.99,
				},
				item{
					Name:        "2 Eggs",
					Description: "Cooked to order",
					Price:       4.00,
				},
				item{
					Name:        "4 Pancakes",
					Description: "Served with maple syurp.",
					Price:       5.50,
				},
			},
		},

		meal{
			Name: "Lunch",
			Items: []item{
				item{
					Name:        "Hamburger",
					Description: "Comes with all the toppings.",
					Price:       5.50,
				},
				item{
					Name:        "Fried Chicken",
					Description: "1 piece white 1 piece dark meat",
					Price:       6.79,
				},
				item{
					Name:        "Turkey Sandwich",
					Description: "With mayo and mustard on toast",
					Price:       3.50,
				},
			},
		},

		meal{
			Name: "Dinner",
			Items: []item{
				item{
					Name:        "Mac and Cheese",
					Description: "Three cheese mac and cheese.",
					Price:       7.50,
				},
				item{
					Name:        "Fried Chicken",
					Description: "Whole chicken.",
					Price:       8.79,
				},
				item{
					Name:        "Ham and Mashed Potatoes",
					Description: "It's our specialty",
					Price:       9.50,
				},
			},
		},
	}
	err := tpl.Execute(os.Stdout, m)
	if err != nil {
		log.Fatalln(err)
	}
}
