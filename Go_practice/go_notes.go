package main

import (
	"errors"
)

func main() {
	/*
		// unicode functions
		var s = "rabbit 73: äàü[&*}$;#~§°;...<>ç+-/*漢字"

		for _, r := range s {
			switch true {
			case unicode.IsLetter(r):
				fmt.Println(string(r), "letter (", reflect.TypeOf(r), ") ")
			case unicode.IsDigit(r):
				fmt.Println(string(r), "digit (", reflect.TypeOf(r), ") ")
			case unicode.IsSpace(r):
				fmt.Println(string(r), "space (", reflect.TypeOf(r), ") ")
			case unicode.IsSymbol(r):
				fmt.Println(string(r), "symbol (", reflect.TypeOf(r), ") ")
			case unicode.IsMark(r):
				fmt.Println(string(r), "mark (", reflect.TypeOf(r), ") ")
			case unicode.IsGraphic(r):
				fmt.Println(string(r), "graphic (", reflect.TypeOf(r), ") ")
			case unicode.IsPunct(r):
				fmt.Println(string(r), "punct (", reflect.TypeOf(r), ") ")
			default:
				fmt.Printf("???? %q %v\n", string(r), reflect.TypeOf(r))
			}
		}
	*/
	/* STRINGS */
	/*
		var s = "This, is a nice evening."
		// Prefix and Suffix check
		fmt.Println(strings.HasPrefix(s, "Th"))
		fmt.Println(strings.HasSuffix(s, "ing."))
		// Contains
		fmt.Println(strings.Contains(s, "nice"))
		// First Index
		fmt.Println(strings.Index(s, "is"))
		// Last Index
		fmt.Println(strings.LastIndex(s, "is"))
		// Replace: -1 -> all occurences
		old := "nice"
		new := "boring"
		fmt.Println(strings.Replace(s, old, new, -1))
		// Count
		fmt.Println(strings.Count(s, "is"))
		// Repeat
		fmt.Println(strings.Repeat("He", 5))
		// ToLower/ToUpper
		fmt.Println(strings.ToLower(s))
		fmt.Println(strings.ToUpper(s))
		// Trim
		// TrimSpace() remove leading and trailing whitespace
		fmt.Println(strings.TrimSpace(" sdhhs skjk   "))
		// Trim() specifies the characters to remove
		// TrimLeft() / TrimRight()
		fmt.Println(strings.Trim(s, "This is "))

		// Splitting a string
		// on whitespace
		strSlice := strings.Fields(s)
		for _, str := range strSlice {
			fmt.Println(str)
		}
		// on a separator
		strSlice = strings.Split(s, ",")
		for _, str := range strSlice {
			fmt.Println(str)
		}
		// Joining over a slice
		fmt.Println(strings.Join(strSlice, ","))

		// Conversion
		// int -> string
		fmt.Println(strconv.Itoa(7))
		// float -> string
		// float, fmt format, precision, bizSize
		fmt.Println(strconv.FormatFloat(6.745, 'b', 3, 64))
		fmt.Println(strconv.FormatFloat(6.745, 'e', 3, 64))
		fmt.Println(strconv.FormatFloat(6.745, 'f', 3, 64))
		fmt.Println(strconv.FormatFloat(6.745, 'g', 3, 64))

		// string -> int
		convInt, _ := strconv.Atoi("7")
		fmt.Println(convInt)
		// string -> float
		convFloat, _ := strconv.ParseFloat("24.56454", 64)
		fmt.Println(convFloat)
	*/

	/* ERROR HANDLING */
	/*
		// no try/catch -> defer-panic-and-recover
		// -> functions return an error object

		// create a new error type
		var errNotFound = errors.New("Not Found")
		fmt.Printf("%s\n", errNotFound)

		// handling error as return from a function in an if statement
		var fin float64 = -1
		if fout, err := Sqrt(fin); err != nil {
			fmt.Printf("Error: %s\n", err)
		} else {
			fmt.Printf("%f\n", fout)
		}

		// run-time panic -> e.g. index out of range
		// can also be initiated from code
		// programs/function exits after a panic -> unlike errors
		var user = os.Getenv("USER")
		if user == "" {
			panic("No value for $USER")
		} else {
			fmt.Println(user)
		}

		// recover: recover from panic and resume execution of program/function
		// -> only useful inside a deferred functions
	*/

	/* Goroutines and Channels */
	/*
		built-in support for communication between concurrent applications
		Process: entity independently executing, that runs in its own address space in memory
		Thread: executing entities sharing the same address space
			-> a process is composed of one or more threads
		Parallelism: ability to make threads run on multiple processors
			-> a concurrent program may or may not be parallel

		Racing conditions: problems occuring with shared data by multithreaded applications
			-> threads manipulate the data unpredictably
			-> DO NOT USE SHARED MEMORY OR GLOBAL VARIABLES
			-> solution: synchronize the threads and lock the data: only 1 thread at a time can change data

		-> GO solution: CSP = Communicating Sequential Processes == message passing-model

		Go routine: part of an application that run concurrently
			-> mapped onto one or more threads, according to their availability
			-> this is accomplished by the goroutine-scheduler in the Go runtime
			-> go routines run in the same address space
				-> access to shared memory must be synchronized
				-> use Go Channels to synchronize goroutines
			-> when a goroutine is blocked by a system call, other goroutines continue to run on other threads
			-> lightweight (lighter than threads): 4k memory stack-space on the heap
			-> use segmented stack for dynamic memory usage
			-> implemented as a function/method, invoked with "go"

		2 styles of concurrency
			1/ deterministic (well-defined ordering)
			2/ non-deterministic (order undefined)

		- Go routines are executed concurrently one 1 core by default
			-> use GOMAXPROCS to enforce parallelism
				-> how many go routines shall execute simultaneously
				-> runtime.GOMAXPROCS(n): goroutines distributed among the n processors
				-> GOMAXPROCS = 8: 8 concurrent threads

		- stop a goroutine (rarely necessary): runtime.Goexit(1)

		Channel: communication pipe between goroutines
			-> ability to read/write is passed with typed values
				-> data race impossible
			-> guarantees synchronization
			-> can only transmit data-items of 1 datatyoe
			-> typed message queue
				-> data can be passed through
			-> FIFO structure
				-> similar to two-ways pipe in Unix shell
			-> reference type
				-> use make() to allocate memory for it
			-> 2 go routines communicate through the same channel
				-> channel as argument of both methods


		/* Characters / Runes */
	// characters are stored in byte -> alias of unint8
	// var ch byte

	// // conversion of a byte to string gives character of ASCII table
	// for i := 0; i < 255; i++ {
	// 	ch++
	// 	fmt.Printf(string(ch))
	// }

	// // runes (or unicode chars) are stored as integers -> alias of int32
	// var ru int

	// for i := 0; i < 65535; i++ {
	// 	ru++
	// 	// conversion of an integer to string gives char of Unicode UTF-8
	// 	fmt.Printf(string(ru))
	//}

	/*String concatenation*/
	// -> use bytes.Buffer

	// var bytbuf bytes.Buffer

	// a := "Hello "
	// b := "World!"

	// bytbuf.WriteString(a)
	// bytbuf.WriteString(b)

	// fmt.Println(bytbuf.String())

	// change a character in a string: convert to []byte first
	// str := "hello"
	// c := []byte(str)
	// c[0] = 'c'
	// str2 := string(c)
	// fmt.Println(str2)

	// // substring
	// sub := str[0:3]
	// fmt.Println(sub)

	// // number of chars in a string
	// fmt.Println(utf8.RuneCountInString(str))
	// fmt.Println(len(str))
	// looping over a string with str[i] will yield the bytes - range will yield the runes

	/* initialization */

	// new() -> arrays, structs, and other value types
	// make() -> slices, maps, channels

}

//Sqrt returns a float64 and an error
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("square root of negative number is illegal")
	}
	return f * f, nil
}
