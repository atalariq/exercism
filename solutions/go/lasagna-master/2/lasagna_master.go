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
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	scaled := make([]float64, len(quantities))
	for i := range quantities {
		scaled[i] = quantities[i] * float64(portions) / 2.0
	}
	return scaled
}
