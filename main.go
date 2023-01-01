package main

import (
	"fmt"
	"os"
)

func main() {
	var lang string
	var day, month, year int
	var monthVal int

	fmt.Println("What Languague do you prefer? / Qual Idioma você prefere?")
	fmt.Println("( EN / PT )")
	fmt.Scanln(&lang)

	if lang == "PT" || lang == "pt" {
		fmt.Println("Idioma escolhido: Português")
	} else if lang == "EN" || lang == "en" {
		fmt.Println("Chosen language: English")
	} else {
		fmt.Println("Invalid Language. The language was set to: English")
		lang = "EN"
	}

	if lang == "PT" || lang == "pt" {
		fmt.Println("Insira o dia (DD): ")
	} else if lang == "EN" || lang == "en" {
		fmt.Println("Insert the day (DD): ")
	}
	fmt.Scanln(&day)
	if day < 1 || day > 31 {
		if lang == "PT" || lang == "pt" {
			fmt.Println("Dia inválido, finalizando programa.")
		} else if lang == "EN" || lang == "en" {
			fmt.Println("Invalid day, exiting program.")
		}
		os.Exit(1)
	}

	if lang == "PT" || lang == "pt" {
		fmt.Println("Insira o mês (MM): ")
	} else if lang == "EN" || lang == "en" {
		fmt.Println("Insert the month (MM): ")
	}
	fmt.Scanln(&month)
	if month < 1 || month > 12 {
		if lang == "PT" || lang == "pt" {
			fmt.Println("Mês inválido, finalizando programa.")
		} else if lang == "EN" || lang == "en" {
			fmt.Println("Invalid month, exiting program.")
		}
		os.Exit(1)
	}
	if (month == 5 || month == 6 || month == 9 || month == 11) && day > 30 {
		if lang == "PT" || lang == "pt" {
			fmt.Println("Esse mês acaba no dia 30")
		} else if lang == "EN" || lang == "en" {
			fmt.Println("This month ends in the 30th")
		}
		os.Exit(1)
	} else if month == 2 && day > 29 {
		if lang == "PT" || lang == "pt" {
			fmt.Println("Esse mês acaba no dia 28, exceto em anos bissextos.")
		} else if lang == "EN" || lang == "en" {
			fmt.Println("This month ends in the 28th, except in leap years.")
		}
		os.Exit(1)
	}

	if lang == "PT" || lang == "pt" {
		fmt.Println("Insira o ano (YYYY): ")
	} else if lang == "EN" || lang == "en" {
		fmt.Println("Insert the year (YYYY): ")
	}
	fmt.Scanln(&year)
	fmt.Println("")

	var leapYear int = year % 4
	var centenaryLeapYear int = year % 400

	if leapYear != 0 && month == 2 && day > 28 {
		if lang == "PT" || lang == "pt" {
			fmt.Println("Fevereiro acaba no dia 29 apenas em anos bissextos.")
		} else if lang == "EN" || lang == "en" {
			fmt.Println("February only ends in the 29th in leap years.")
		}
		os.Exit(1)
	}

	if month == 1 || month == 10 {
		monthVal = day + 5
		if (leapYear == 0 || centenaryLeapYear == 0) && month == 1 {
			monthVal = day + 4
		}
	} else if month == 2 || month == 3 || month == 11 {
		monthVal = day + 1
		if (leapYear == 0 || centenaryLeapYear == 0) && month == 2 {
			monthVal = day
		}
	} else if month == 4 || month == 7 {
		monthVal = day + 4
	} else if month == 5 {
		monthVal = day + 6
	} else if month == 6 {
		monthVal = day + 2
	} else if month == 8 {
		monthVal = day
	} else if month == 9 || month == 12 {
		monthVal = day + 3
	}

	var centenaryValue = year % 100
	var yearValue = ((centenaryValue + (centenaryValue / 4) + 1) % 7)

	var decimalValue = year / 100
	var decimalPoint int

	if decimalValue == 12 || decimalValue == 16 || decimalValue == 20 || decimalValue == 24 {
		decimalPoint = yearValue
	} else if decimalValue == 13 || decimalValue == 17 || decimalValue == 21 || decimalValue == 25 {
		decimalPoint = yearValue + 5
	} else if decimalValue == 14 || decimalValue == 18 || decimalValue == 22 || decimalValue == 26 {
		decimalPoint = yearValue + 3
	} else if decimalValue == 15 || decimalValue == 19 || decimalValue == 23 || decimalValue == 27 {
		decimalPoint = yearValue + 1
	}

	var finalValue = (decimalPoint + monthVal) % 7

	if finalValue == 0 {
		if lang == "PT" || lang == "pt" {
			fmt.Println("Esse dia é um Domingo")
		} else if lang == "EN" || lang == "en" {
			fmt.Println("This day is a Sunday")
		}
	} else if finalValue == 1 {
		if lang == "PT" || lang == "pt" {
			fmt.Println("Esse dia é uma Segunda")
		} else if lang == "EN" || lang == "en" {
			fmt.Println("This day is a Monday")
		}
	} else if finalValue == 2 {
		if lang == "PT" || lang == "pt" {
			fmt.Println("Esse dia é uma Terça")
		} else if lang == "EN" || lang == "en" {
			fmt.Println("This day is a Tuesday")
		}
	} else if finalValue == 3 {
		if lang == "PT" || lang == "pt" {
			fmt.Println("Esse dia é uma Quarta")
		} else if lang == "EN" || lang == "en" {
			fmt.Println("This day is a Wednesday")
		}
	} else if finalValue == 4 {
		if lang == "PT" || lang == "pt" {
			fmt.Println("Esse dia é uma Quinta")
		} else if lang == "EN" || lang == "en" {
			fmt.Println("This day is a Thursday")
		}
	} else if finalValue == 5 {
		if lang == "PT" || lang == "pt" {
			fmt.Println("Esse dia é uma Sexta")
		} else if lang == "EN" || lang == "en" {
			fmt.Println("This day is a Friday")
		}
	} else if finalValue == 6 {
		if lang == "PT" || lang == "pt" {
			fmt.Println("Esse dia é um Sábado")
		} else if lang == "EN" || lang == "en" {
			fmt.Println("This day is a Saturday")
		}
	} else {
		if lang == "PT" || lang == "pt" {
			fmt.Println("Algo deu errado. Finalizando programa.")
		} else if lang == "EN" || lang == "en" {
			fmt.Println("Something went wrong. Exiting program.")
		}
		os.Exit(1)
	}

}
