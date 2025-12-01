package lasagna

func PreparationTime(layers []string, avgTimePerLayer int) int {
	if avgTimePerLayer == 0 {
		avgTimePerLayer = 2
	}
	return len(layers) * avgTimePerLayer
}

func Quantities(layers []string) (int, float64) {
	var noodles int
	var sauce float64
	for _, v := range layers {
		switch v {
		case "noodles":
			noodles += 50
		case "sauce":
			sauce += 0.2
		}
	}
	return noodles, sauce
}

func AddSecretIngredient(friendsList, myList []string) {
	for _, v := range friendsList {
		for ii, vv := range myList {
			if v == vv {
				break
			} else if vv == "?" {
				myList[ii] = v
				break
			}
		}
	}
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	scaled := make([]float64, len(quantities))
	copy(scaled, quantities)
	for i := range scaled {
		scaled[i] = scaled[i] * float64(portions) / 2.0
	}
	return scaled
}
