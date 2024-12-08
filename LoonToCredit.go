package main

import (
	"fmt"
	"math/big"
)

var x1000 = []string{"", "k", "M", "G", "T", "P", "E", "Z", "Y", "R", "Q", "X1", "X2", "X3", "X4", "X5", "X6"}
var x1000text = []string{"", " Duizend ", " Miljoen ", " Miljard ", " Biljoen ", " Biljard ", " Triljoen ", " Triljard ", " Quadriljoen ", " Quadriljard ", " Quintiljoen ", " Quintiljard ", " Septiljoen ", " Septiljard ", "Octiljoen ", "Octiljard ", " Noniljoen "}

func formatBigInt(number *big.Int) string {
	// Haal het teken van het getal op
	isNegative := number.Sign() < 0
	// Werk met de absolute waarde
	absNumber := new(big.Int).Abs(number)

	numberStr := absNumber.String()
	length := len(numberStr)
	if length < 4 {
		if isNegative {
			return "-" + numberStr
		}
		return numberStr
	}

	// Bereken de index voor de suffix
	unitIndex := (length - 1) / 3
	if unitIndex >= len(x1000) {
		unitIndex = len(x1000) - 1
	}

	// Bereken de schaal en deel het getal
	divisor := big.NewInt(10)
	divisor.Exp(divisor, big.NewInt(int64(3*unitIndex)), nil)
	quotient := new(big.Int).Div(absNumber, divisor)

	// Voeg het teken toe als het een negatief getal is
	result := fmt.Sprintf("%s%s", quotient.String(), x1000[unitIndex])
	if isNegative {
		result = "-" + result
	}
	return result
}

func convertToText(number *big.Int) string {
	numberStr := number.String()
	length := len(numberStr)
	if length < 4 {
		return ""
	}

	unitIndex := (length - 1) / 3
	if unitIndex >= len(x1000text) {
		unitIndex = len(x1000text) - 1
	}

	return x1000text[unitIndex]
}

func main() {
	for {
		// Vraag de gebruiker om input
		var minimum, multiplier, creditAdd, incomeLeft string
		fmt.Print("\nVoer het minimumloon voor de drempelwaarde in (500): ")
		_, err := fmt.Scanln(&minimum)
		if err != nil {
			fmt.Println("Fout bij het lezen van het minimumloon:", err)
			continue
		}

		fmt.Print("Voer de vermenigvuldigingsfactor in (20): ")
		_, err = fmt.Scanln(&multiplier)
		if err != nil {
			fmt.Println("Fout bij het lezen van de vermenigvuldigingsfactor:", err)
			continue
		}

		fmt.Print("Per hoeveel loon komen er credits bij (10): ")
		_, err = fmt.Scanln(&creditAdd)
		if err != nil {
			fmt.Println("Fout bij het lezen van de loon/credit add ratio:", err)
			continue
		}

		fmt.Print("Voer het overgebleven loon in: ")
		_, err = fmt.Scanln(&incomeLeft)
		if err != nil {
			fmt.Println("Fout bij het lezen van het overgebleven loon:", err)
			continue
		}

		// Zet strings om naar big.Int
		minThreshold := big.NewInt(0)
		minThreshold.SetString(minimum, 10)

		multFactor := big.NewInt(0)
		multFactor.SetString(multiplier, 10)

		credRatio := big.NewInt(0)
		credRatio.SetString(creditAdd, 10)

		incomeLeftBig := big.NewInt(0)
		incomeLeftBig.SetString(incomeLeft, 10)

		// Bereken het verschil
		difference := big.NewInt(0)
		difference.Sub(incomeLeftBig, minThreshold)

		// Bereken het aantal stappen
		steps := big.NewInt(0)
		steps.Div(difference, credRatio)

		// Bereken de som van de reeks getallen in stappen van credRatio
		// S = n * (n + 1) / 2 * multFactor
		n := steps
		nPlusOne := big.NewInt(0)
		nPlusOne.Add(n, big.NewInt(1))

		sumN := big.NewInt(0)
		sumN.Mul(n, nPlusOne)
		sumN.Div(sumN, big.NewInt(2))
		sumN.Mul(sumN, multFactor)

		if difference.Sign() < 0 {
			sumN.Neg(sumN) // Maak het resultaat negatief als het verschil negatief is
		}

		// Formatteer de output
		creditsFormatted := formatBigInt(sumN)
		creditsText := convertToText(sumN)

		// Output resultaat
		fmt.Printf("De berekende credits zijn: %s | %s | %s\n", sumN.String(), creditsFormatted, creditsText)
	}
}
