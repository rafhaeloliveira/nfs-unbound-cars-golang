package main

import "github.com/gin-gonic/gin"

type Performance struct {
	TopSpeed int `json:"topSpeed"`  
	Transmission string `json:"transmission"`
}

type Car struct {
	Name string `json:"name"`
	ImageURL string `json:"imageUrl"`
	Performance Performance `json:"performance"`
}

var carList = []Car {
	{	Name: "2018 Lamborghini Aventador LP 740-4 S",
		ImageURL: "https://static.wikia.nocookie.net/nfs/images/b/b9/NFSUB_Garage_Lamborghini_AventadorLP7404S2018.jpg/revision/latest?cb=20230218183107&path-prefix=en",
		Performance: Performance { 
			TopSpeed: 218,
			Transmission: "7-speed Automatic",
		},
 	},
}

func main() {
	router := gin.Default()

	router.Run()
}