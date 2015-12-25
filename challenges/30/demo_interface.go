package main

import "fmt"

//extraPaid is an interface and has a single method overTimePaid that returns an int
type extraPaid interface {
	overTimePaid(int) int
}

//define a Dev struct that has the info of a developer
type Dev struct {
	name		string
	yearInService	int
	progLanguage	string
	overTimeRate	int
}

//a function to calculate the overtime rate for a developer.
func (d Dev) overTimePaid(hour int) int {
	return d.overTimeRate * hour
}

//define a Ops struct that has the info for a operator 
type Ops struct {
	name		string
	yearInService	int
	certification	string
	overTimeRate	int
}

//a function to calculate the overtime rate for an operator.
// Operation has an extra $10000 as bonus because they are smile more
func (o Ops) overTimePaid(hour int) int {
	return o.overTimeRate * hour + 10000
}

func main() {
	var e extraPaid
	const numHours = 40

	fmt.Printf("\nDuring Christmas season everyone worked an extra %d hours\n\n", numHours)

	d := Dev{"Tom", 3, "Go, C, C++, Java, Python", 120}
	//print out the details of developer
	fmt.Println("Developer details are:", d)
	fmt.Printf("  Name: %s\n", d.name)
	fmt.Printf("  Year in Service: %d\n", d.yearInService)
	fmt.Printf("  Skill: %s\n", d.progLanguage)
	e = d
	fmt.Printf("Christmas extra paid = $%d", e.overTimePaid(numHours))
	fmt.Printf("\n\n")

	o := Ops{"Jerry", 3, "VCP, CCIE, NPX", 100}
	//print out the details of operator
	fmt.Println("Operator details are: ", o)
	fmt.Printf("  Name: %s\n", o.name)
	fmt.Printf("  Year in Service: %d\n", o.yearInService)
	fmt.Printf("  Skill: %s\n", o.certification)
	e = o
	fmt.Printf("Christmas extra paid = $%d", e.overTimePaid(numHours))
	fmt.Printf("\n\n")
}
