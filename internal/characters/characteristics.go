package characters

import "github.com/dpsommer/eventstream/internal/utils"

type gender int
type sex int

const (
	nonbinary gender = iota
	man
	woman

	male sex = iota
	female

	// TODO: configurable weights
	// note that these don't need to sum to 100,
	// it's just easier to think of them as percentages
	maleWeight   int = 50
	femaleWeight int = 50

	cisWeight   int = 98
	transWeight int = 1
	nbWeight    int = 1
)

func chooseSex() sex {
	return utils.WeightedChoice(
		utils.Choice[sex]{
			Element: male,
			Weight:  maleWeight,
		},
		utils.Choice[sex]{
			Element: female,
			Weight:  femaleWeight,
		},
	)
}

func chooseGender(s sex) gender {
	switch s {
	case male:
		return utils.WeightedChoice(
			utils.Choice[gender]{
				Element: man,
				Weight:  cisWeight,
			},
			utils.Choice[gender]{
				Element: nonbinary,
				Weight:  nbWeight,
			},
			utils.Choice[gender]{
				Element: woman,
				Weight:  transWeight,
			},
		)
	case female:
		fallthrough
	default:
		return utils.WeightedChoice(
			utils.Choice[gender]{
				Element: woman,
				Weight:  cisWeight,
			},
			utils.Choice[gender]{
				Element: nonbinary,
				Weight:  nbWeight,
			},
			utils.Choice[gender]{
				Element: man,
				Weight:  transWeight,
			},
		)
	}
}
