package characters

import (
	"math/rand"
)

var (
	genderNeutralNames []string = []string{
		"Avery", "Alex", "Alexis", "Aubrey", "Bailey", "Blake", "Charlie",
		"Madison", "Riley", "Dakota", "Blair", "Briar", "Carey", "Harper",
		"Adair", "Ali", "Arden", "Aspen", "Casey", "Jamie", "Logan", "Rigel",
		"Finley", "Lindsay", "Robin", "Rowan", "Mackenzie", "Briar", "Kirby",
		"Morgan", "River", "Taylor", "Kat", "Hayden", "Cam", "Brook",
	}

	maleNames []string = append(genderNeutralNames, []string{
		"Henry", "Jack", "James", "Oliver", "William", "Charles", "Noah",
		"George", "Lucas", "Thomas", "John", "Leo", "Theodore", "Arthur",
		"Daniel", "Elijah", "Aaron", "Alexander", "Benjamin", "Edward",
		"Jacob", "Joseph", "Julian", "Liam", "Duncan", "Ryan", "Michael",
		"Matthew", "Hunter",
	}...)

	femaleNames []string = append(genderNeutralNames, []string{
		"Ava", "Mia", "Olivia", "Amelia", "Emma", "Lily", "Evelyn", "Isabella",
		"Isla", "Sophia", "Katherine", "Ivy", "Lucia", "Audrey", "Charlotte",
		"Chloe", "Daisy", "Eleanor", "Emily", "Madeline", "Scarlett", "Jessica",
		"Helena", "Avery", "Elizabeth", "Kate",
	}...)

	surnames []string = []string{
		"Brown", "Johnson", "Jones", "Smith", "Taylor", "Thomas", "Williams",
		"Anderson", "Evans", "Wilson", "Davies", "White", "Davis", "Hughes",
		"Robinson", "Thompson", "Walker", "Allen", "Armstrong", "Bailey",
		"Edwards", "Green", "Harris", "Jackson", "London",
	}
)

func generateName(g gender) string {
	switch g {
	case man:
		return maleNames[rand.Intn(len(maleNames))]
	case woman:
		return femaleNames[rand.Intn(len(femaleNames))]
	case nonbinary:
		fallthrough
	default:
		return genderNeutralNames[rand.Intn(len(genderNeutralNames))]
	}
}

func generateSurname() string {
	return surnames[rand.Intn(len(surnames))]
}
