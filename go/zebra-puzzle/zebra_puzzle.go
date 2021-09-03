package zebra

import (
	"fmt"
)

type Solution struct {
	DrinksWater string
	OwnsZebra   string
}

// Location starts with 1
type Location int

type LocationRelative int

const (
	right LocationRelative = iota + 1
	nextTo
)

// Color is: red, green, ivory, yellow, blue
type Color int

const (
	red Color = iota + 1
	green
	ivory
	yellow
	blue
)

// Nationality is: Englishman, Spaniard, Ukrainian, Norwegian, Japanese
type Nationality int

const (
	Englishman Nationality = iota + 1
	Spaniard
	Ukrainian
	Norwegian
	Japanese
)

// Pet is: dog, snails, fox, horse, zebra
type Pet int

const (
	dog Pet = iota + 1
	snails
	fox
	horse
	zebra
)

// Beverage is: coffee, tea, milk, orange juice, water
type Beverage int

const (
	coffee Beverage = iota + 1
	tea
	milk
	orangeJuice
	water
)

// CigarBrand is: old gold, kools, chesterfields, lucky strike, parliaments
type CigarBrand int

const (
	oldGold CigarBrand = iota + 1
	kools
	cheserfields
	luckyStrike
	parliaments
)

type House struct {
	color            Color
	nationality      Nationality
	pet              Pet
	beverage         Beverage
	cigarBrand       CigarBrand
	location         Location
	locationRelative LocationRelative
}

type Clue struct {
	input  House
	output House
	solved bool
}

func SolvePuzzle() Solution {
	s := Solution{}
	var c *Clue

	houses := []*House{}

	// There are five houses.
	for i := 0; i < 5; i++ {
		h := &House{
			location: Location(i + 1),
		}
		houses = append(houses, h)
	}

	clues := []*Clue{}

	// The Englishman lives in the red house.
	c = &Clue{
		input: House{
			color: red,
		},
		output: House{
			nationality: Englishman,
		},
	}
	clues = append(clues, c)

	// The Spaniard owns the dog.
	c = &Clue{
		input: House{
			nationality: Spaniard,
		},
		output: House{
			pet: dog,
		},
	}
	clues = append(clues, c)

	// Coffee is drunk in the green house.
	c = &Clue{
		input: House{
			color: green,
		},
		output: House{
			beverage: coffee,
		},
	}
	clues = append(clues, c)

	// The Ukrainian drinks tea.
	c = &Clue{
		input: House{
			nationality: Ukrainian,
		},
		output: House{
			beverage: tea,
		},
	}
	clues = append(clues, c)

	// The green house is immediately to the right of the ivory house.
	c = &Clue{
		input: House{
			color:            ivory,
			locationRelative: right, // to the right of the input house
		},
		output: House{
			color: green,
		},
	}
	clues = append(clues, c)

	// The Old Gold smoker owns snails.
	c = &Clue{
		input: House{
			cigarBrand: oldGold,
		},
		output: House{
			pet: snails,
		},
	}
	clues = append(clues, c)

	// Kools are smoked in the yellow house.
	c = &Clue{
		input: House{
			color: yellow,
		},
		output: House{
			cigarBrand: kools,
		},
	}
	clues = append(clues, c)

	// Milk is drunk in the middle house (start with 1).
	c = &Clue{
		input: House{
			location: 3,
		},
		output: House{
			beverage: milk,
		},
	}
	clues = append(clues, c)

	// The Norwegian lives in the first house.
	c = &Clue{
		input: House{
			location: 1,
		},
		output: House{
			nationality: Norwegian,
		},
	}
	clues = append(clues, c)

	// The man who smokes Chesterfields lives in the house next to the man with the fox.
	c = &Clue{
		input: House{
			pet:              fox,
			locationRelative: nextTo, // next to the input house
		},
		output: House{
			cigarBrand: cheserfields,
		},
	}
	clues = append(clues, c)

	// Kools are smoked in the house next to the house where the horse is kept.
	// TODO

	// The Lucky Strike smoker drinks orange juice.
	c = &Clue{
		input: House{
			cigarBrand: luckyStrike,
		},
		output: House{
			beverage: orangeJuice,
		},
	}
	clues = append(clues, c)

	// The Japanese smokes Parliaments.
	c = &Clue{
		input: House{
			nationality: Japanese,
		},
		output: House{
			cigarBrand: parliaments,
		},
	}
	clues = append(clues, c)

	// The Norwegian lives next to the blue house.

	solve(clues, houses)

	return s
}

func solve(clues []*Clue, houses []*House) {
	fmt.Println("Solving...")
	fmt.Println()
	fmt.Printf("Houses (start):\n")
	printHouses(houses)

	cluesUnsolved := unsolvedClues(clues)

	for cluesUnsolved > 0 {
		for i := 0; i < len(clues); i++ {
			processClue(clues[i], houses)
		}
		cluesUnsolved = unsolvedClues(clues)
		fmt.Println()
		fmt.Printf("Clues :\n")
		printClues(clues)
	}

	fmt.Println()
	fmt.Printf("Houses (end):\n")
	printHouses(houses)
}

func unsolvedClues(clues []*Clue) int {
	unsolved := 0

	for _, c := range clues {
		if !c.solved {
			unsolved++
		}
	}

	return unsolved
}

func processClue(c *Clue, houses []*House) {
	var solved bool
	var h *House
	var err error

	relative := false
	// get the house from input
	switch {
	case c.input.pet != 0:
		h, err = houseByPet(c.input.pet, houses)
	case c.input.cigarBrand != 0:
		h, err = houseByCigarBrand(c.input.cigarBrand, houses)
	case c.input.beverage != 0:
		h, err = houseByBeverage(c.input.beverage, houses)
	case c.input.nationality != 0:
		h, err = houseByNationality(c.input.nationality, houses)
	case c.input.location != 0:
		h, err = houseByLocation(c.input.location, houses)
	case c.input.color != 0:
		h, err = houseByColor(c.input.color, houses)
	}

	if c.input.locationRelative != 0 {
		relative = true
	}

	// failed to find house
	if err != nil {
		return
	}

	solved = true

	if relative {
		var offset Location
		switch c.input.locationRelative {
		case right:
			offset = 1
		}

		h, err = houseByLocation(h.location+offset, houses)
		if err != nil {
			panic(err)

		}
	}

	// set attribute in output
	switch {
	case c.output.pet != 0:
		h.pet = c.output.pet
	case c.output.cigarBrand != 0:
		h.cigarBrand = c.output.cigarBrand
	case c.output.beverage != 0:
		h.beverage = c.output.beverage
	case c.output.nationality != 0:
		h.nationality = c.output.nationality
	case c.output.location != 0:
		h.location = c.output.location
	case c.output.color != 0:
		h.color = c.output.color
	default:
		panic("no output set!")
	}

	// set solved to true if success
	c.solved = solved

}

func printHouses(houses []*House) {
	for _, h := range houses {
		fmt.Printf("%#v\n", h)
	}
}

func printClues(clues []*Clue) {
	for _, c := range clues {
		fmt.Printf("%#v\n", c)
	}
}

func houseByAttribute(attr string, value interface{}, houses []*House) (*House, error) {
	switch attr {
	case "color":
		return houseByColor(value.(Color), houses)
	case "nationality":
		return houseByNationality(value.(Nationality), houses)
	case "beverage":
		return houseByBeverage(value.(Beverage), houses)
	case "cigarbrand":
		return houseByCigarBrand(value.(CigarBrand), houses)
	case "pet":
		return houseByPet(value.(Pet), houses)
	}

	return nil, fmt.Errorf("invalid attribute")
}

func houseByColor(c Color, houses []*House) (*House, error) {
	for _, h := range houses {
		if h.color == c {
			return h, nil
		}
	}

	return nil, fmt.Errorf("unknown by color")
}

func houseByNationality(n Nationality, houses []*House) (*House, error) {
	for _, h := range houses {
		if h.nationality == n {
			return h, nil
		}
	}

	return nil, fmt.Errorf("unknown by nationality")
}

func houseByPet(p Pet, houses []*House) (*House, error) {
	for _, h := range houses {
		if h.pet == p {
			return h, nil
		}
	}

	return nil, fmt.Errorf("unknown by pet")
}

func houseByBeverage(b Beverage, houses []*House) (*House, error) {
	for _, h := range houses {
		if h.beverage == b {
			return h, nil
		}
	}

	return nil, fmt.Errorf("unknown by beverage")
}

func houseByCigarBrand(c CigarBrand, houses []*House) (*House, error) {
	for _, h := range houses {
		if h.cigarBrand == c {
			return h, nil
		}
	}

	return nil, fmt.Errorf("unknown by cigar brand")
}

func houseByLocation(l Location, houses []*House) (*House, error) {
	for _, h := range houses {
		if h.location == l {
			return h, nil
		}
	}

	return nil, fmt.Errorf("unknown by location")
}
