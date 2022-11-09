package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

var x1000 = []string{"", "k", "M", "G", "T", "P", "E", "Z", "Y", "X", "W", "V", "U", "TD", "S", "R", "Q"}

type Geld struct {
	naam        string
	amount      big.Int
	vorigAmount big.Int
	dig         big.Int
	x1000       int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func setX1000(geld *Geld) {
	var comp = big.NewInt(1000)
	var x = 0
	for geld.amount.Cmp(comp) > 0 {
		comp.Mul(comp, big.NewInt(1000))
		x++
	}
	(*geld).x1000 = x
	comp.Div(comp, big.NewInt(1000))
	(*geld).dig.Div(&geld.amount, comp)
}

func setMoney(geld *Geld, value string) {
	(*geld).amount.SetString(value, 10)
	setX1000(geld)
}

func addMoney(geld *Geld) {
	(*geld).vorigAmount.SetString((*geld).amount.String(), 10)
	var grootte string
	var hoeveelheid string
	var indexGrootte int
	fmt.Print("Grootte van het bedrag ", x1000, ": ")
	fmt.Scanln(&grootte)
	fmt.Print("Hoeveelheid van het bedrag: ")
	fmt.Scanln(&hoeveelheid)
	for i := range x1000 {
		if grootte == x1000[i] {
			indexGrootte = i
			break
		}
	}
	var multiplier = big.NewInt(1000)
	multiplier.Exp(multiplier, big.NewInt(int64(indexGrootte)), nil)
	var bedrag big.Int
	bedrag.SetString(hoeveelheid, 10)
	bedrag.Mul(&bedrag, multiplier)
	var totaal big.Int
	totaal.Add(&geld.amount, &bedrag)
	setMoney(geld, totaal.String())
}

func subMoney(geld *Geld) {
	(*geld).vorigAmount.SetString((*geld).amount.String(), 10)
	var grootte string
	var hoeveelheid string
	var indexGrootte int
	fmt.Print("Grootte van het bedrag ", x1000, ": ")
	fmt.Scanln(&grootte)
	fmt.Print("Hoeveelheid van het bedrag: ")
	fmt.Scanln(&hoeveelheid)
	for i := range x1000 {
		if grootte == x1000[i] {
			indexGrootte = i
			break
		}
	}
	var multiplier = big.NewInt(1000)
	multiplier.Exp(multiplier, big.NewInt(int64(indexGrootte)), nil)
	var bedrag big.Int
	bedrag.SetString(hoeveelheid, 10)
	bedrag.Mul(&bedrag, multiplier)
	var totaal big.Int
	totaal.Sub(&geld.amount, &bedrag)
	setMoney((geld), totaal.String())
}

func loadFile() []string {
	var rek []string
	file, err := os.Open("money.txt")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rek = append(rek, scanner.Text())
	}
	return rek
}

func saveFile(rekeningen *[]Geld) {
	file, err := os.Create("./money.txt")
	check(err)
	for i := range *rekeningen {
		file.WriteString((*rekeningen)[i].naam)
		file.WriteString("\n")
		file.WriteString((*rekeningen)[i].amount.String())
		file.WriteString("\n")
	}
	file.Close()
	file.Sync()
}

func importeerRekeningen() []Geld {
	rek := loadFile()
	var rekeningen = make([]Geld, len(rek)/2)
	for i := range rek {
		if i%2 == 0 {
			rekeningen[i/2].naam = rek[i]
		} else {
			setMoney(&rekeningen[i/2], rek[i])
		}
	}
	return rekeningen
}

func balans(rekening Geld) {
	fmt.Println("\n", rekening.naam, "\n", rekening.amount.String(), "\n", rekening.dig.String(), x1000[rekening.x1000])
}

func nieuweRekening(naam string) Geld {
	var rekening Geld
	setMoney(&rekening, "0")
	rekening.naam = naam
	return rekening
}

func main() {
	var rekeningen []Geld
	var huidigeRekening int
	for {
		fmt.Println("1. Importeer rekeningen")
		fmt.Println("2. Maak rekening aan")
		fmt.Println("3. Verwijder rekening")
		fmt.Println("4. Wijzig actieve rekening")
		fmt.Println("5. Verander totaal geld")
		fmt.Println("6. + geld")
		fmt.Println("7. - geld")
		fmt.Println("8. CANCEL vorige wijziging")
		fmt.Println("9. Sla op")
		var keuze int
		fmt.Print("Keuze: ")
		fmt.Scanln(&keuze)

		if keuze == 1 {
			rekeningen = importeerRekeningen()
		} else if keuze == 2 {
			var naam string
			fmt.Print("\nNaam: ")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			naam = scanner.Text()
			rekeningen = append(rekeningen, nieuweRekening(naam))
		} else if keuze == 3 {
			var nummer int
			for i := range rekeningen {
				fmt.Println(i, rekeningen[i].naam)
			}
			fmt.Print("Nummer: ")
			fmt.Scanln(&nummer)
			rekeningen = append(rekeningen[:nummer], rekeningen[nummer+1:]...)
			if huidigeRekening > len(rekeningen) {
				huidigeRekening = len(rekeningen) - 1
			}
		} else if keuze == 4 {
			for i := range rekeningen {
				fmt.Println(i, rekeningen[i].naam)
			}
			fmt.Print("Kies: ")
			fmt.Scanln(&huidigeRekening)
		} else if keuze == 5 {
			var nieuw string
			fmt.Print("\nNieuw bedrag: ")
			fmt.Scanln(&nieuw)
			setMoney(&rekeningen[huidigeRekening], nieuw)
		} else if keuze == 6 {
			addMoney(&rekeningen[huidigeRekening])
		} else if keuze == 7 {
			subMoney(&rekeningen[huidigeRekening])
		} else if keuze == 8 {
			setMoney(&rekeningen[huidigeRekening], rekeningen[huidigeRekening].vorigAmount.String())
		} else if keuze == 9 {
			saveFile(&rekeningen)
		}
		balans(rekeningen[huidigeRekening])
	}
}
