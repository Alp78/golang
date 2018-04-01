package main

import (

	//"errors"

	"fmt"
	"os"
)

const (
	// iota increments following const by 1
	answer1 = iota
	answer2
	answer3
	answer4
	answer5
)

// Function Definitions
// idiomatic arguments declaration: the two are declared just once with same type string
func printer(msg1, msg2 string, repeat int) {
	for repeat > 0 {
		fmt.Printf("%s", msg1)
		fmt.Printf("%s\n", msg2)
		repeat--
	}
}

// error handling in functions
// specify the function returns only the error
func printError(msg1 string) error {
	msg1 += " ErrorTest\n"
	// we only want the error
	_, err := fmt.Printf(msg1) // assigning the Printf will actually execute it too
	return err
}

// os.Create / os.WriteString / defer: execute code after exit
func createFile(msg string) error {
	f, err := os.Create("gotest.txt")
	// close the file when the function exits only
	defer f.Close()
	// write msg in the text file
	f.WriteString(msg)
	return err
}

// ... takes an undefined number of arguments
func multipleArgs(msgs ...string) {
	for _, msg := range msgs {
		fmt.Printf("%s ", msg)
	}
	fmt.Printf("\n%d arguments taken", len(msgs))
}

// can pass a map as argument
func monthDays(monthMap map[string]int, month string) error {
	_, err := fmt.Printf("%d\n", monthMap[month])
	return err
}

func main() {
	/*
		// go determines the variable type, no need to specify
		message := "The answer to life is %d %d %d %d %d \n"

		// message is the format (string) adding integer %d variables taken in order
		fmt.Printf(message, answer1, answer2, answer3, answer4, answer5)

		// equivalent to declaring: var pi float32 = 3.14
		pi := 3.14
		var pi2 float32 = 3.138347637463

		// %f - width and precision: .3 -> default width and 3 numbers after decimals
		fmt.Printf("Value: %.3f %.6f \n", pi, pi2)

		// unsigned integers: uint (larger positive value, and no negative value)
		var unsigned uint = 6
		var signed int = -6

		fmt.Printf("Unsigned %d signed %d \n", unsigned, signed)

		// boolean %t: truth value
		isTrue := true
		isFalse := !true

		fmt.Printf("Truth value %t %t \n", isTrue, isFalse)

		// bytes
		var buffer1 [256]byte // declares an arry of 256 bytes

		// len(array): length of array
		for i := 0; i < len(buffer1); i++ {
			fmt.Printf("%d ", i)
		}

		fmt.Printf("\n \n")

		// slice of an array - the explicit type is not needed and should be dropped for more idiomatic syntax
		var slice1 []byte = buffer1[100:150]

		for i := 0; i < len(slice1); i++ {
			fmt.Printf("%d ", i)
		}

		fmt.Printf("\n \n")

		// a slice of a slice
		slice2 := slice1[0:10]

		for i := 0; i < len(slice2); i++ {
			fmt.Printf("%d ", i)
		}

		fmt.Printf("\n \n")

		// for range loop: iterates over elements
		slice3 := buffer1[10:20]

		fmt.Println("Original slice: ")

		for i := range slice3 {
			fmt.Printf("%x ", slice3[i])
		}

		fmt.Printf("\n \n")

		// range provides both index and value on arrays and slices
		for i, v := range slice3 {
			fmt.Printf("Index: %d / Value: %x \n", i, v)
		}

		fmt.Printf("\n \n")

		// adds 1 to each element of the slice
		fmt.Println("Adds 1 to each element of the slice: ")

		for i := range slice3 {
			slice3[i]++
		}

		for i := range slice3 {
			fmt.Printf("%x ", slice3[i])
		}

		fmt.Printf("\n \n")

		// affects i to each elemnent of the slice
		fmt.Println("Affects i to each element of the slice: ")

		for i := range slice3 {
			slice3[i] = byte(i)
		}

		for i := range slice3 {
			fmt.Printf("%x ", slice3[i])
		}

		fmt.Printf("\n \n")

		// map: hash table
		kvs := map[string]string{"a": "ananas", "b": "banana", "c": "carrot", "d": "dinosaur"}

		for k, v := range kvs {
			fmt.Printf("%s -> %s \n", k, v)
		}

		// strings %s
		sent1 := "the quick brown fox jumps over the lazy dog \n"

		// extracting a substring
		sub1 := sent1[2:15]
		fmt.Printf("%s\n", sub1)

		// each character %c in a string has an index (to ignore index, replace i by _)
		for i, r := range sent1 {
			fmt.Printf("%d %c__", i, r)
		}

		// a string between `accents` is interpreted literally -> no escape chars
		sent2 := `no \n escape`

		fmt.Printf(sent2)

		fmt.Printf("\n \n")

		// error handling / "error" is a built-in go type
		// inline declaration with ;
		// Printf returns a number of bytes (the characters) and an error type
		// \n is counted as character!!
		if n, err := fmt.Printf("Text String tested.\n"); err != nil {
			fmt.Printf("There is an error!")
			os.Exit(1)
		} else {
			fmt.Printf("Printed %d bytes\n", n)
		}
		fmt.Printf("\n \n")

		// formatting an error using fmt.Errorf - usefule for simple check
		// where the error is not needed elswehere

		emptyString := ""
		var err error

		if emptyString == "" {
			err = fmt.Errorf("Nope")
		}

		if err != nil {
			fmt.Printf("The string is empty: %s\n", err)
		} else {
			fmt.Printf("%s\n", emptyString)
		}
		// switch statement
		n, err := fmt.Printf("asdefrhjultovs")

		switch {
		case err != nil:
			os.Exit(1)
		case n == 0:
			fmt.Printf("No bytes output\n")
		case n != 14:
			fmt.Printf("Wrong number of characters\n")
		default:
			fmt.Printf("\nOK!\n")
		}

		fmt.Printf("\n \n")

		fmt.Printf("\n")

		vowels := 0
		consonants := 0
		zeds := 0

		// range on strings: returns 1)index 2)rune - can be expressed in %b, %c, %x...
		sent1 := "the quick brown fox jumps over the lazy dog in the azure \n"

		// index not needed? replace by _
		for _, r := range sent1 {
			fmt.Printf("%c", r)
			switch r {
			// inline fallthrough: no need of a separate case for each char
			case 'a', 'e', 'i', 'o', 'u':
				vowels++
			// fallthrough: even if case is valid, still go to next
			case 'z':
				zeds++
				fallthrough
			default:
				consonants++
			}
		}
		fmt.Printf("\n")

		fmt.Printf("There are %d vowels and %d consonants of which %d are 'z'\n", vowels, consonants, zeds)

		fmt.Printf("\n \n")

		// for loop

		// infinite loop - runs forever
		counter := 0

		for counter < 10 {
			fmt.Printf("Hello\n")
			counter++
		}

		// compact way ";" separating the clauses
		for counter := 0; counter < 10; counter++ {
			fmt.Printf("Hello\n")
		}

		// simultaneous assignment

		for i, j := 0, 1; i < 10; i, j = i+1, j*2 {
			fmt.Printf("i-%d j-%d Hello\n", i, j)
		}

		// mimicking WHILE loop with FOR - no need of any sorts of other loop. FOR covers all
		var stop bool
		i := 0

		// unless stop is "true", keep looping
		for !stop {
			fmt.Printf("Hello\n")
			i++
			if i == 10 {
				stop = true
			}
		}

		// Calling functions

		printer("Hello ", "I'm a function", 10)

		// assigning the return error type of printError
		err := printError("Hello")

		// if there is no error.. == nil // if error.. != nil
		if err == nil {
			// defining an error message from errors package function New
			// turns the error type into string
			err = errors.New("Doesn't work.")
			// must invoke Println - won't work with Printf
			fmt.Println(err)
		}

		// calling createFile function - print a message if no error

		err1 := createFile("Golang is awesome.")

		if err1 == nil {
			err1 = errors.New("There was no error.")
			fmt.Println(err1)
		}

		// calling multipleArgs with 6 arguments
		multipleArgs("I'm", "a function", "that's not taking", "a fixed number", "of arguments", "which is really cool")

		// Array and Slices

		// declare an array with ... undefined size then affects a number of values
		words := [...]string{"the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}
		for _, s := range words {
			fmt.Printf("%s ", s)
		}
		fmt.Printf("\n")

		// extract a slice form an array: a window on the array, not another array -> resource friendly
		sliceWords := words[0:3]
		for _, s := range sliceWords {
			fmt.Printf("%s ", s)
		}

		// make: declare a an array and add elements later - must declare size as second argument
		words2 := make([]string, 4)

		// cap: capacity
		fmt.Printf("\nLength %d\nCapacity %d\n", len(words2), cap(words2))

		words2[0] = "the"
		words2[1] = "quick"
		words2[2] = "brown"
		words2[3] = "fox"

		// can append to exceed the set length - capacity doubled automatically
		words2 = append(words2, "jumps")

		for _, s := range words2 {
			fmt.Printf("%s ", s)
		}

		fmt.Printf("\nLength %d\nCapacity %d\n", len(words2), cap(words2))
		fmt.Printf("\n")

		// setting the capacity without the length: empty array
		words3 := make([]string, 0, 4)

		fmt.Printf("\nLength %d\nCapacity %d\n", len(words3), cap(words3))

		// append method - if the current capacity is exceeded, it is automatically doubled
		words3 = append(words3, "the")
		words3 = append(words3, "quick")
		words3 = append(words3, "brown")
		words3 = append(words3, "fox")

		// exceeding the capacity by just one element
		words3 = append(words3, "jumps")

		for _, s := range words3 {
			fmt.Printf("%s ", s)
		}
		fmt.Printf("\n")

		fmt.Printf("\nLength %d\nCapacity %d\n", len(words3), cap(words3))

		// copy a slice: creates a new set of data - the recipient must be declared with a set length
		wordsCopy := make([]string, 3)
		copy(wordsCopy, words3)

		for _, s := range wordsCopy {
			fmt.Printf("%s ", s)
		}
		fmt.Printf("\n")

		fmt.Printf("\nLength %d\nCapacity %d\n", len(wordsCopy), cap(wordsCopy))

		// maps: key/value pairs - using "make" if you don't know the values in advance...
		// otherwise normal declaration "key1" : value1, "key2" : value2 ...
		// no need of declaring the size

		dayMonths := make(map[string]int)

		dayMonths["Jan"] = 31
		dayMonths["Feb"] = 28
		dayMonths["Mar"] = 31
		dayMonths["Apr"] = 30
		dayMonths["May"] = 31
		dayMonths["Jun"] = 30
		dayMonths["Jul"] = 31
		dayMonths["Aug"] = 30

		for k, v := range dayMonths {
			fmt.Printf("Key: %s - Value: %d\n", k, v)
		}

		fmt.Printf("Days in February: %d\n", dayMonths["Feb"])

		// if the key doesn't exist in the map, it returns 0
		fmt.Printf("Days in February: %d\n", dayMonths["Febraury"])

		// "ok" keyword: if you need to distinguish between 0 and an actual inexisting key -
		// error handling won't work here as if the key is not found, it returns 0, which is not an error
		// ok is boolean
		days, ok := dayMonths["February"]

		if !ok {
			fmt.Printf("Key doesn't exist.")
		} else {
			fmt.Printf("Days: %d\n", days)
		}
		fmt.Printf("\n%t\n", ok)

		fmt.Printf("\n")
		// range over a map: it has no order

		for month, days := range dayMonths {
			fmt.Printf("Month: %s - Days: %d\n", month, days)
		}

		// open a file with write access
		f, err := os.OpenFile("test.txt", os.O_WRONLY, 0666)
		if err != nil {
			fmt.Printf("%s\n", err)
			os.Exit(1)
		}

		f.WriteString("Golang is awesome. Yeah.")
		defer f.Close()

		fmt.Printf("\n\n")
	*/
}
