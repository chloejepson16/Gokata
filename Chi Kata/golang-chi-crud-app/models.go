package main

type GroceryItem struct{
	ID string `json:"id"`
	Name string `json:"name"`
	Category string `json:"category"`
	Price string `json:"price"`
}

var groceryList= []*GroceryItem{
	{
		ID: "1",
		Name: "Apple",
		Category: "Produce",
		Price: "$1.79/lb",
	},
	{
		ID: "2",
		Name: "Milk",
		Category: "Dairy",
		Price: "$3.59",
	},
	{
		ID: "3",
		Name: "Lettuce",
		Category: "Produce",
		Price: "$1.29/lb",
	},
	{
		ID: "4",
		Name: "Ground Beef",
		Category: "Meat",
		Price: "$8.57/lb",
	},
}

func listGroceries() []*GroceryItem{
	return groceryList
}

func getGroceries(id string) *GroceryItem{
	for _, grocery:= range groceryList{
		if grocery.ID == id{
			return grocery
		}
	}
	return nil
}

func addGrocery(grocery GroceryItem){
	groceryList= append(groceryList, &grocery)
}

func deleteGrocery(id string) *GroceryItem{
	for i, grocery:= range groceryList{
		if grocery.ID == id{
			deletedGrocery := groceryList[i]
			groceryList = append(groceryList[:i], groceryList[i+1:]...)
			return deletedGrocery
		}
	}

	return nil
}

func updateGrocery(id string, groceryUpdate GroceryItem) *GroceryItem{
	for i, grocery:= range groceryList{
		if grocery.ID == id{
			groceryList[i]= &groceryUpdate
			return grocery
		}
	}
	return nil
}