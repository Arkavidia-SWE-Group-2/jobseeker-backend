package main

import (
	"fmt"
	"jobseeker/database"
	"jobseeker/internal/config"
)

func main() {
	fmt.Print("Are you sure you want to migrate the database? (N/y) >>> ")

	var answer string

	fmt.Scanln(&answer)

	if answer != "y" {
		fmt.Println("Migration cancelled")
		return
	}

	config := config.NewViper()
	db := database.New(config)

	err := database.Migrate(db)

	if err != nil {
		fmt.Println("Migration failed")
		fmt.Println(err)
		return
	}

	fmt.Println("Migration success")
}
