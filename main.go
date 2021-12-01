package main

import (
	"fmt"
	"os"
	"strings"

	flag "github.com/ogier/pflag"
)

// global variables
var (
	user  string
	users []string // declare as a slice

)

func main() {
	// Parse the flags.
	flag.Parse()

	// if user does not supply flags, print usage

	if flag.NFlag() == 0 {
		fmt.Printf("Usage: %s [options]\n", os.Args[0])
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// declare := is for declaration + assignment, whereas = is for assignment only.
	// users := strings.Split(user, ",")
	users = strings.Split(user, ",") // see global variables
	fmt.Printf("Searching user(s): %s\n", users)
}

func init() {
	// StringVarP() does 3 things:
	// 1. it tells GO that we will be evaluating expecting a String,
	// 2. it tells GO that we want to bind a Variable to this flag, and
	// 3. it tells GO that we want to have a Posix-compliant flag (e.g. double-dash and single-dash flag)

	//StringVarP() takes 5 arguments in this order:
	// 1. the variable we want to bind this flag to,
	// 2. the double-dash flag,
	// 3. the single-dash flag,
	// 4. the default value to use if flag is not explicitly called,
	// 5. and the description of this flag

	// &user means that we are passing a reference (a.k.a. memory address) of the user variable.
	flag.StringVarP(&user, "user", "u", "", "Search Users")
}
