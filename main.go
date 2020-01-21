package main

import (
	"fmt"

	"github.com/Twitter_stat/config"
)

func main() {

	someToken := config.Credentials{}

	fmt.Println(someToken.GetCredentials())
}
