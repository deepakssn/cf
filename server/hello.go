package main

import "fmt"

func age() {
	var age = 10
	fmt.Printf("hello, world\n")
	fmt.Println(age)

	type domains struct {
		AllowedDomain []string
		BlockedDomain []string
	}
	var AllowedDomains = []string{"infosys.com", "skidata.com"}
	var BlockedDomains = []string{"test.com", "gmail.com"}

	fmt.Println(domains{AllowedDomain: AllowedDomains, BlockedDomain: BlockedDomains})
}
