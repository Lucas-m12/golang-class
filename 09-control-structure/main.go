package main

func weekDays(key int) string {
	switch key {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sabado"
	default:
		return "Day not exist"
	}
}

func weekDays2(key int) string {
	var weekDay string
	switch {
	case key == 1:
		weekDay = "Domingo"
	case key == 2:
		weekDay = "Segunda"
		fallthrough
	case key == 3:
		weekDay = "Terça"
	case key == 4:
		weekDay = "Quarta"
	case key == 5:
		weekDay = "Quinta"
	case key == 6:
		weekDay = "Sexta"
	case key == 7:
		weekDay = "Sabado"
	default:
		weekDay = "Day not exist"
	}
	return weekDay
}

func main() {
	var n1 int8 = 10
	if n1 > 15 {
		println("Maior que 15")
	} else {
		println("Menor ou igual a 15")
	}
	n1 = 0
	if n2 := &n1; *n2 > 0 {
		println("Maior que zero")
	} else {
		println("Menor ou igual a zero")
	}

	n2 := 2
	day := weekDays(n2)
	println(day)
	day2 := weekDays2(n2)
	println(day2)
}
