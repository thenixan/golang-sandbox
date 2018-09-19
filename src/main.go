package main

import (
	"fmt"
	"userdatabase"
)

func main() {
	fmt.Println(userdatabase.CheckAuth("dial","ESXbgvqY"))
	fmt.Println(userdatabase.ChangePassword("dial","ESXbgvqY", "ESXbgvqYa"))
	fmt.Println(userdatabase.CheckAuth("dial","ESXbgvqY"))
}
