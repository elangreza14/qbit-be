package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

const rawFruits = `[
    {
        "fruitId": 1,
        "fruitName": "Apel",
        "fruitType": "IMPORT",
        "stock": 10
    },
    {
        "fruitId": 2,
        "fruitName": "Kurma",
        "fruitType": "IMPORT",
        "stock": 20
    },
    {
        "fruitId": 3,
        "fruitName": "apel",
        "fruitType": "IMPORT",
        "stock": 50
    },
    {
        "fruitId": 4,
        "fruitName": "Manggis",
        "fruitType": "LOCAL",
        "stock": 100
    },
    {
        "fruitId": 5,
        "fruitName": "Jeruk Bali",
        "fruitType": "LOCAL",
        "stock": 10
    },
    {
        "fruitId": 6,
        "fruitName": "KURMA",
        "fruitType": "IMPORT",
        "stock": 20
    },
    {
        "fruitId": 7,
        "fruitName": "Salak",
        "fruitType": "LOCAL",
        "stock": 150
    }
]`

type (
	FruitType string

	Fruit struct {
		ID    int       `json:"fruitId"`
		Name  string    `json:"fruitName"`
		Type  FruitType `json:"fruitType"`
		Stock int       `json:"stock"`
	}

	Fruits []Fruit
)

const (
	FruitTypeImport FruitType = "IMPORT"
	FruitTypeLocal  FruitType = "LOCAL"
)

func main() {
	fruits := make(Fruits, 0, 0)
	err := json.Unmarshal([]byte(rawFruits), &fruits)
	if err != nil {
		log.Fatal(err)
	}

	// Jawaban Nomor 1, buah yang dimiliki Andi
	fmt.Println("===== nomor 1 =====")
	fruitNames := GetFruitMap(fruits)
	for fruitName := range fruitNames {
		fmt.Println("ada buah", fruitName)
	}

	// Jawaban Nomor 2, Jumlah Wadah yang dibutuhkan adalah
	fmt.Println("===== nomor 2 =====")
	fruitBaskets := GetFruitBasket(fruitNames)
	fmt.Printf("Jumlah Wadah yang dibutuhkan adalah %d\n", len(fruitBaskets))
	for fruitBasket, fruits := range fruitBaskets {
		fmt.Printf("di dalam wadah %s\n", fruitBasket)
		for _, fruit := range fruits {
			fmt.Println("ada buah", fruit.Name, fruit.Stock)
		}
	}

	// Jawaban Nomor 3, Jumlah buah yang ada di dalam Wadah adalah
	fmt.Println("===== nomor 3 =====")
	for fruitBasket, fruits := range fruitBaskets {
		totalFruits := 0
		for _, fruit := range fruits {
			totalFruits += fruit.Stock
		}
		fmt.Printf("wadah %s memiliki %d buah\n", fruitBasket, totalFruits)
	}

	// Jawaban Nomor 4, Apakah ada komentar terkait kasus di atas?
	fmt.Println("===== nomor 4 =====")
	fmt.Println("- untuk value dari fruitId belum konsisten")
	fmt.Println("- untuk penulisan fruitname belum konsisten sebagian menggunakan backtick sebagian menggunakan doublequote ")
	fmt.Println("- untuk penamaan fruitName masih belum konsisten masih ada yang uppercase dan lowercase")
	fmt.Println("- untuk type Fruits struct bisa diubah menjadi Fruit dan Fruits untuk slice of Fruit")

}

func GetFruitMap(fruits Fruits) map[string]Fruit {
	fruitMap := make(map[string]Fruit)
	for _, fruit := range fruits {
		if v, ok := fruitMap[strings.ToLower(fruit.Name)]; !ok {
			fruitMap[strings.ToLower(fruit.Name)] = fruit
		} else {
			v.Stock += fruit.Stock
			fruitMap[strings.ToLower(fruit.Name)] = v
		}
	}
	return fruitMap
}

func GetFruitBasket(fruitMap map[string]Fruit) map[FruitType]Fruits {
	fruitBaskets := make(map[FruitType]Fruits)
	for _, fruit := range fruitMap {
		if _, ok := fruitBaskets[fruit.Type]; !ok {
			fruitBaskets[fruit.Type] = make(Fruits, 0)
		}
		fruitBaskets[fruit.Type] = append(fruitBaskets[fruit.Type], fruit)
	}

	return fruitBaskets
}
