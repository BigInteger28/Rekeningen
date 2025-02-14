package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
    "strings"
)

var (
	x1000     = []string{"", "k", "M", "G", "T", "P", "E", "Z", "Y", "R", "Q", "X11", "X12", "X13", "X14", "X15", "X16", "X17", "X18", "X19", "X20", "X21"}
	x1000text = []string{"", " Duizend ", " Miljoen ", " Miljard ", " Biljoen ", " Biljard ", " Triljoen ", " Triljard ", " Quadriljoen ", " Quadriljard ", " Quintiljoen ", " Quintiljard ", " Sextiljoen ", " Sextiljard ", " Septiljoen ", " Septiljard ", " Octiljoen ", " Octiljard ", " Noniljoen ", " Noniljard ", " Deciljoen ", " Deciljard "}
)

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

func formatBigNumber(numberStr string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	parts := []string{}
	for i := n; i > 0; i -= 3 {
		start := max(0, i-3)
		parts = append([]string{s[start:i]}, parts...)
	}
	return strings.Join(parts, " ")
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

func money(grootte string, hoeveelheid string) big.Int {
	var indexGrootte int
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
	return bedrag
}

func askUserAmount() (string, string) {
	var grootte string
	var hoeveelheid string
	fmt.Print("Grootte van het bedrag ", x1000, ": ")
	fmt.Scanln(&grootte)
	fmt.Print("Hoeveelheid van het bedrag: ")
	fmt.Scanln(&hoeveelheid)
	return grootte, hoeveelheid
}

func addMoneyTo(rekening *Geld, grootte string, hoeveelheid string) {
	(*rekening).vorigAmount.SetString((*rekening).amount.String(), 10)
	var totaal big.Int
	bedrag := money(grootte, hoeveelheid)
	totaal.Add(&rekening.amount, &bedrag)
	setMoney(rekening, totaal.String())
}

func income(rekening *Geld) {
	var keuze int
	var keuzearray = []string{"90", "100", "1000", "10", "10", "10", "50", "25", "10", "20"}
	var aantal string
	var money, number2 big.Int
	fmt.Println("1. Dagen gewerkt voor werkgever")
	fmt.Println("2. Spaargeld storting")
	fmt.Println("3. Level behaald")
	fmt.Println("4. Pompen")
	fmt.Println("5. Huishoudelijke taak of nuttig werk (minuten)")
	fmt.Println("6. Stuk fruit eten")
	fmt.Println("7. Gezonde maaltijd eten")
	fmt.Println("8. Wandelen")
	fmt.Println("9. Lopen (Sneller dan 10 km/u) (m)")
  fmt.Println("10. Fietsen (m)")
	fmt.Print("Keuze: ")
	fmt.Scanln(&keuze)
	if keuze != 1 && keuze != 8 && keuze != 9 && keuze != 10 {
		fmt.Print("Aantal: ")
		fmt.Scanln(&aantal)
		money.SetString(aantal, 10)
		number2.SetString(keuzearray[keuze-1], 10)
		money.Mul(&money, &number2)
	} else if keuze == 1 {
		var lvl string
		fmt.Print("Hoogst behaald level: ")
		fmt.Scanln(&lvl)
		fmt.Print("Aantal dagen gewerkt: ")
		fmt.Scanln(&aantal)
		var add, dagen big.Int
		money.SetString(lvl, 10)
		dagen.SetString(aantal, 10)
		add.SetString(keuzearray[0], 10)
		number2.SetString("10", 10)
		money.Mul(&money, &number2)
		money.Add(&money, &add)
		money.Mul(&money, &dagen)
	} else if keuze == 8 {
		var keuze2 int
		fmt.Println("1: Stappen")
		fmt.Println("2: Meter")
		fmt.Print("Keuze: ")
		fmt.Scanln(&keuze2)
		fmt.Print("Aantal: ")
		fmt.Scanln(&aantal)
		if keuze2 == 1 {
			money.SetString(aantal, 10)
			number2.SetString(keuzearray[keuze-1], 10)
			money.Div(&money, &number2)
		} else {
			money.SetString(aantal, 10)
			number2.SetString("20", 10)
			money.Div(&money, &number2)
		}
	} else if keuze == 9 {
		fmt.Print("Aantal meter: ")
		fmt.Scanln(&aantal)
		money.SetString(aantal, 10)
		number2.SetString(keuzearray[keuze-1], 10)
		money.Div(&money, &number2)
	} else if keuze == 10 {
    fmt.Print("Aantal meter: ")
		fmt.Scanln(&aantal)
		money.SetString(aantal, 10)
		number2.SetString(keuzearray[keuze-1], 10)
		money.Div(&money, &number2)
  }
	addMoneyTo(rekening, "", money.String())
}

func subMoneyFrom(rekening *Geld, grootte string, hoeveelheid string) {
	(*rekening).vorigAmount.SetString((*rekening).amount.String(), 10)
	var totaal big.Int
	bedrag := money(grootte, hoeveelheid)
	totaal.Sub(&rekening.amount, &bedrag)
	setMoney((rekening), totaal.String())
}

func expense(rekening *Geld) {
	var keuze int
	var keuzearray = []string{"120", "2000", "100", "20"}
	var aantal string
	var money, number2 big.Int
	fmt.Println("1. Afhaling spaargeld")
	fmt.Println("2. Start level")
	fmt.Println("3. Ongezonde maaltijd")
	fmt.Println("4. Ongezonde snack")
	fmt.Print("Keuze: ")
	fmt.Scanln(&keuze)
	fmt.Print("Aantal: ")
	fmt.Scanln(&aantal)
	money.SetString(aantal, 10)
	number2.SetString(keuzearray[keuze-1], 10)
	money.Mul(&money, &number2)
	subMoneyFrom(rekening, "", money.String())
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
	money := formatBigNumber(rekening.amount.String())
	fmt.Println("\n", rekening.naam, "\n", money, "\n", rekening.dig.String(), x1000text[rekening.x1000], x1000[rekening.x1000])
}

func nieuweRekening(naam string) Geld {
	var rekening Geld
	setMoney(&rekening, "0")
	rekening.naam = naam
	return rekening
}

func kiesRekening(tekst string, rekeningen *[]Geld) int {
	var kies int
	for i := range *rekeningen {
		fmt.Println(i, (*rekeningen)[i].naam)
	}
	fmt.Print(tekst)
	fmt.Scanln(&kies)
	return kies
}

func main() {
	var rekeningen []Geld
	var huidigeRekening int
	var transactie bool
	var transactieVan, transactieNaar int
  rekeningen = importeerRekeningen()
	for {
    fmt.Println("0. Importeer money.txt")
		fmt.Println("1. Maak rekening aan")
		fmt.Println("2. Verwijder rekening")
		fmt.Println("3. Wijzig actieve rekening")
		fmt.Println("4. Verander totaal geld")
		fmt.Println("5. + geld")
		fmt.Println("6. Inkomsten")
		fmt.Println("7. - geld")
		fmt.Println("8. Van rekening naar rekening")
		fmt.Println("9. Uitgaven")
		fmt.Println("10. CANCEL vorige wijziging")
		fmt.Println("11. Sla op")
		var keuze int
		fmt.Print("Keuze: ")
		fmt.Scanln(&keuze)
   if keuze == 0 {
		  rekeningen = importeerRekeningen()
   } else if keuze == 1 {
			var naam string
			fmt.Print("\nNaam: ")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			naam = scanner.Text()
			rekeningen = append(rekeningen, nieuweRekening(naam))
		} else if keuze == 2 {
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
		} else if keuze == 3 {
			huidigeRekening = kiesRekening("Kies: ", &rekeningen)
		} else if keuze == 4 {
			var nieuw string
			fmt.Print("\nNieuw bedrag: ")
			fmt.Scanln(&nieuw)
			setMoney(&rekeningen[huidigeRekening], nieuw)
		} else if keuze == 5 {
			grootte, hoeveelheid := askUserAmount()
			addMoneyTo(&rekeningen[huidigeRekening], grootte, hoeveelheid)
		} else if keuze == 6 {
			income(&rekeningen[huidigeRekening])
		} else if keuze == 7 {
			grootte, hoeveelheid := askUserAmount()
			subMoneyFrom(&rekeningen[huidigeRekening], grootte, hoeveelheid)
		} else if keuze == 8 {
			var vanRekening int
			var naarRekening int
			vanRekening = kiesRekening("Van rekening: ", &rekeningen)
			transactieVan = vanRekening
			naarRekening = kiesRekening("Naar rekening: ", &rekeningen)
			transactieNaar = naarRekening
			grootte, hoeveelheid := askUserAmount()			
			subMoneyFrom(&rekeningen[vanRekening], grootte, hoeveelheid)
			addMoneyTo(&rekeningen[naarRekening], grootte, hoeveelheid)
			transactie = true
		} else if keuze == 9 {
			expense(&rekeningen[huidigeRekening])
		} else if keuze == 10 {
			if !transactie {
				setMoney(&rekeningen[huidigeRekening], rekeningen[huidigeRekening].vorigAmount.String())
			} else {
				setMoney(&rekeningen[transactieVan], rekeningen[transactieVan].vorigAmount.String())
				setMoney(&rekeningen[transactieNaar], rekeningen[transactieNaar].vorigAmount.String())
				transactie = false
			}
		} else if keuze == 11 {
			saveFile(&rekeningen)
		}
		balans(rekeningen[huidigeRekening])
	}
}
