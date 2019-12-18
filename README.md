# Headfirst series

- [x] Print alphabet
- [ ] Conversions (strconv)
- [ ] Error & Log ( log )
- [ ] User Input ( os, bufio )
- [ ] Strings ( strings )

## The Headfirst Games

### Game 3

**Multiline string ()**
- Ref : https://stackoverflow.com/questions/7933460/how-do-you-write-multiline-strings-in-go

```golang
	fmt.Println(`
		Game rules: 
		- Generate a random number from 0 - 100.
		- Prompt the player to guess what target number is, and store their response
		- If the player's guess is less than target number, says "LOW" else
		says "HIGH"
		- Allow the player to guess up to 10 times. Before each guess let them know how many
		guess they have left
		- If the player's guess is equal to the target number, tell them "Good job". Then stop asking guess
		- If the player ran out of turns without guessing correctly. Says "Sorry, it was [target]\"
	`)  
```

**math/rand**

- https://golang.org/pkg/math/rand/
- https://gobyexample.com/random-numbers

```golang
func generateNumber() int{
	// Create and seed the generator.
	// Typically a non-fixed seed should be used, such as time.Now().UnixNano().
	// Using a fixed seed will produce the same output on every run.	
	// fixed seed: rand.Seed(123)
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100)
}
```
