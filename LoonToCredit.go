package main

import (
	"fmt"
	"math/big"
)

var x1000 = []string{"", "k", "M", "G", "T", "P", "E", "Z", "Y", "R", "Q", "X1", "X2", "X3", "X4", "X5", "X6"}
var x1000text = []string{"", " Duizend ", " Miljoen ", " Miljard ", " Biljoen ", " Biljard ", " Triljoen ", " Triljard ", " Quadriljoen ", " Quadriljard ", " Quintiljoen ", " Quintiljard ", " Septiljoen ", " Septiljard ", "Octiljoen ", "Octiljard ", " Noniljoen "}

func formatBigInt(number *big.Int) string {
    numberStr := number.String()
    length := len(numberStr)
    if length < 4 {
        return numberStr
    }

    unitIndex := (length - 1) / 3
    if unitIndex >= len(x1000) {
        unitIndex = len(x1000) - 1
    }

    divisor := big.NewInt(10)
    divisor.Exp(divisor, big.NewInt(int64(3*unitIndex)), nil)
    quotient := big.NewInt(0)
    quotient.Div(number, divisor)

    return fmt.Sprintf("%s%s", quotient.String(), x1000[unitIndex])
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
        var minimum, multiplier, incomeLeft string
        fmt.Print("\nVoer het minimumloon voor de drempelwaarde in (1000): ")
        _, err := fmt.Scanln(&minimum)
        if err != nil {
            fmt.Println("Fout bij het lezen van het minimumloon:", err)
            continue
        }

        fmt.Print("Voer de vermenigvuldigingsfactor in (1): ")
        _, err = fmt.Scanln(&multiplier)
        if err != nil {
            fmt.Println("Fout bij het lezen van de vermenigvuldigingsfactor:", err)
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

        incomeLeftBig := big.NewInt(0)
        incomeLeftBig.SetString(incomeLeft, 10)

        // Bereken het verschil
        difference := big.NewInt(0)
        difference.Sub(incomeLeftBig, minThreshold)

        // Bereken de som van de eerste n natuurlijke getallen, aangepast voor negatieve waarden
        absoluteDifference := new(big.Int).Abs(difference)
        nPlusOne := big.NewInt(0)
        nPlusOne.Add(absoluteDifference, big.NewInt(1))

        sumN := big.NewInt(0)
        sumN.Mul(absoluteDifference, nPlusOne)
        sumN.Div(sumN, big.NewInt(2))

        if difference.Sign() < 0 {
            sumN.Neg(sumN) // Maak het resultaat negatief als het verschil negatief is
        }

        // Vermenigvuldig de som met de vermenigvuldigingsfactor
        totalCredits := big.NewInt(0)
        totalCredits.Mul(sumN, multFactor)

        // Formatteer de output
        creditsFormatted := formatBigInt(totalCredits)
        creditsText := convertToText(totalCredits)

        // Output resultaat
        fmt.Printf("De berekende credits zijn: %s | %s | %s\n", totalCredits.String(), creditsFormatted, creditsText)
    }
}
