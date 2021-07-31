package main

import "fmt"

func main() {
	var sell_cost float64
	var advert_rate float64	

	fmt.Println("Enter the sell cost")
	fmt.Scanln(&sell_cost)

	fmt.Println("Enter the advertising rate")
	fmt.Scanln(&advert_rate)


	var output =((((sell_cost*0.1)+sell_cost)*(0.1))+(sell_cost * 0.029)+0.3)+(sell_cost * advert_rate)

	fmt.Printf("%f %f = %f", sell_cost, advert_rate, output)
}