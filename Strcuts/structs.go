package main

import "fmt"

type User struct {
	name           string
	email          string
	contact_number int
}

func creatUser(_name string, _email string, _contact_number int) User {

	return User{
		name:           _name,
		email:          _email,
		contact_number: _contact_number,
	}
}

type Order struct {
	quantity int
	price    float32
}

// Methods for order Structs

func (o Order) total() float32 {
	return float32(o.quantity) * o.price
}
func (Order) creatOrder(_quantity int, _price float32) Order {
	return Order{
		quantity: _quantity,
		price:    _price,
	}
}
func main() {

	user1 := creatUser("sam", "sam@zym.com", 15466116)
	user2 := creatUser("kim", "kim@aym.com", 4254)
	user3 := creatUser("will", "will@zym.com", 15466116)

	fmt.Println("User 1", user1)
	fmt.Println("User 2", user2)
	fmt.Println("User 3", user3)

	// Accesing Struct elements

	user2_number := user2.contact_number

	fmt.Println("user2 number", user2_number)

	// Array Of Structs

	users := [3]User{user1, user2, user3}

	// looping thriough arrays of users

	for i, val := range users {
		fmt.Printf("user %d is  number is %v\n", i, val.contact_number)
	}

	// Oder structs Uses With structs

	order1 := Order.creatOrder(Order{}, 10, 20.0)
	order2 := Order.creatOrder(Order{}, 30, 90.0)
	order3 := Order.creatOrder(Order{}, 40, 60.0)

	// Array of Orders Structs

	orders := []Order{order1, order2, order3}

	for i, val := range orders {
		fmt.Printf("Order %d is  Price is %v\n", i, val.price)
	}

	// To total price of the three orders
	var sum float32
	for _, val := range orders {
		sum += val.total()

	}

	fmt.Println("Total value Of The Three Orders is ", sum)

}
