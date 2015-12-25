package main

import "fmt"

//Go function can return multiple parameters - a nice feature
func checkMonthIndexSize(i int) (int, error) {
	// switch is also a form of flow control similar to if-then-else
	switch  {
	case i < 0:
		return i, fmt.Errorf("Error: slice len %d is less than zero\n", i)
	case i > 12:
		return i, fmt.Errorf("Error: slice len %d is greater than 12\n", i)
	default:
		return i, nil
	}
}

func main () {

	// variable declarations
	var a int
	var i, j int
	var s string

	// map declaration - index-value pair
	// may not be stored in memory in this exact order.
	var months = map[int]string{ 1:"January",
				     2:"Fabruary",
				     3:"March",
				     4:"April",
				     5:"May",
				     6:"June",
				     7:"July",
				     8:"August",
				     9:"September",
				    10:"October",
				    11:"Novenber",
				    12:"December",
	}

	// slice declaration
	// It is the preferred in Go over array which is static in size
	var monthIndexSlice = []int{1,2,3,4,5,6,7,8,9,10,11,12}
	var string1 []string
	//
	// slice initialization -  append the value to the slice "string1"
	// there are other smarter ways to do this but for now this serves the purpose
	//
	// My programming philosophy:- 
	//
	//        Keep it simple and make it work first then optimize later.
	//
	string1 = append(string1, "Jan")
	string1 = append(string1, "Feb")
	string1 = append(string1, "Mar")
	string1 = append(string1, "April")
	string1 = append(string1, "May")
	string1 = append(string1, "June")
	string1 = append(string1, "July")
	string1 = append(string1, "Aug")
	string1 = append(string1, "Sept")
	string1 = append(string1, "Oct")
	string1 = append(string1, "Nov")
	string1 = append(string1, "Dec")

  //
  // Code illustration starts here:
  //

    //
    //   1. Assigment
    //
	i = len(monthIndexSlice)
	s = "The commitmas project is very useful!"
	// Go supports "Tuple Assignment" - a very useful feature.
	x, y := 100, 200
	fmt.Printf("X = %d, Y = %d\n", x, y)

    //
    //   2. if then else
    //
	if _, err := checkMonthIndexSize(i); err != nil {
	    fmt.Printf("\n%s\n", err)
	} else {
	    fmt.Printf("\nSlice is initialized correctly (len = %d)\n", i)
	}

	// array and slice are zero based.
	// this is a most common mistake for begineers.
	a = monthIndexSlice[i - 1]
	fmt.Printf("\nIn the month of %s %s\n\n", months[a], s)

	// Is there anything wrong with the below if-the-else statement?
	// Syntax is correct, it compiles ok but it might not perform what
	// it was intended.
	x = 10

	if ( x != 1 ) || (x != 2) {
	    fmt.Printf(" ** x = %d, the evaluation is TRUE\n", x)
	} else {
	    fmt.Printf(" ** x = %d, the evaluation is FALSE\n", x)
	}
	fmt.Printf("\n")

    //
    //   3. Iteration
    //
	// * for loop example #1
	for i := 6; i < len(string1); i++ {
		fmt.Printf("  [%d] := %s\r\n", i, string1[i])
	}
	fmt.Printf("\n")

	// * for loop example #2
	//
	// use a blank identifier "_" to take the value of 
	// the index from the range operation on the slice
	// if we use a variable to take the value and that
	// vaiable is not used, the compile will issue an WARNING
	// and program will not compile
	for _, string2 := range string1 {
		fmt.Printf("  %s\r\n", string2)
	}
	fmt.Printf("\n")

	// * for loop example #3
	// There is no while loop in Go
	//
	// No need to set j to zero as the variable is created and initialized to zero
	//
	// Just like in any while loop, make sure the condition to terminate the execution
	// will be satisfied.  Otherwise, we will get the - INFINITE LOOP.
	// and this is not the continous feedback loop that is in DevOps.
	for j < len(string1) {
		fmt.Printf("  [%d] := %s\r\n", j + 1, string1[j])
		j = j + 1
	}

  //
  // Code illustration ends here, program terminates
  //
}
