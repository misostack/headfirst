package main

import (
	"fmt"
	"time"
	"strings"
	"sort"
	"strconv"
	"log"
	"math/rand"
)

type user struct {
	id uint64
	username string
	firstName string
	lastName string
	dob string
	age uint8
	active bool
	votes uint64
}

// sort
type byName []string

func (a byName) Len() int { return len(a) }
func (a byName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
// https://golang.org/src/strings/compare.go: The result will be 0 if a==b, -1 if a < b, and +1 if a > b.
func (a byName) Less(i, j int) bool { return strings.Compare(a[i], a[j]) == -1 }

func main() {
	// naming convention : https://github.com/unknwon/go-code-convention/blob/master/en-US/naming_rules.md
	fmt.Println("[Headfirst Series][Game 12]: Structs")
	pLangs := map[string]string{ }
	pLangs["go"] = "Golang"
	pLangs["js"] = "Javascript"
	pLangs["python"] = "Python"
	pLangs["ruby"] = "Ruby"
	pLangs["c"] = "C lang"
	pLangs["cpp"] = "C++"	
	// display user info
	plks := []string{}
	for plk := range pLangs {
		plks = append(plks, plk)
	}
	sort.Sort(byName(plks))
	fmt.Println(plks)
	// new user list
	users := []user{}
	for idx, k := range plks {
		dob := strconv.Itoa(1996 + idx) + "-01-01"
		u := newUser(uint64(idx + 1), k, pLangs[k], "The Programming Language", dob )
		users = append(users, u)
	}
	for _, u := range users {
		randomVotes(&u)
		displayUserInfo(u)
	}
}

func randomVotes( u *user) {
	duration, err := time.ParseDuration("1s")
	if err != nil { log.Fatal(err) }
	time.Sleep(duration)
	rand.Seed(int64(time.Now().Nanosecond()))
	(*u).votes = uint64(rand.Int63n(10000))
}

func displayUserInfo(user user) {
	fmt.Printf("ID: %v\n", user.id)
	fmt.Printf("Username : %v\n", strings.ToUpper(user.username))
	fmt.Printf("Full Name : %v %v\n", user.firstName, user.lastName)
	fmt.Printf("DOB : %v\n", user.dob)
	fmt.Printf("Age : %v\n", user.age)	
	var userStatus string = "Inactive"
	if user.active {
		userStatus = "Active"
	}	
	fmt.Printf("Status : %v\n", userStatus)
	fmt.Printf("Votes : %v\n", user.votes)	
	for i:=0; i< 50; i++ { fmt.Printf("%v", "-") }
	fmt.Printf("\n")
}

func newUser(id uint64, username string, firstName string, lastName string, dob string) user {
	var user user
	user.id = id
	user.username = username
	user.firstName = firstName
	user.lastName = lastName
	user.dob = dob
	currentYear, err := strconv.Atoi(strings.Split(user.dob, "-")[0])
	if err != nil { log.Fatal(err) }
	user.age = uint8(time.Now().Year() - currentYear)
	// rand seed
	rand.Seed(int64(time.Now().Nanosecond()))
	randomNumber := rand.Intn(100)
	if randomNumber % 2 == 0 {
		user.active = false
	} else {
		user.active = true
	}
	return user
}