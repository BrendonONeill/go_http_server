package data

import "fmt"

var People = []string{"Bob", "Ben", "Betty", "Bernie", "Beth"}

func init() {
	fmt.Println("collected People from data package")
}
