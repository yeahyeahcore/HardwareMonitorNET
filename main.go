package main

import "fmt"

func main() {
	fmt.Println("Hello Утка")

	var age int = 12
	var b int = 135
	var num = 2.3456
	var str = "gaga"
	var res int
	res = age + b
	var arr [2]int
	arr[0] = 45
	arr[1] = 53
	nums := [3]float64{4.23, 5.32, 5.21}
	webSites := make(map[string]int)
	webSites["yeahyeahcore"] = 2
	webSites["psychonaut4"] = 8

	//Объект структуры
	bob := Cats{"Bob", 7, 0.2}
	fmt.Println("Bob age is", bob.age)
	fmt.Println("Bob function is", bob.test())

	//откладывание(выполнение функции после всей хуеты)
	defer two()
	one()

	//ССЫЛКИ БЛЯТЬ
	pointer(&b)
	fmt.Println(b)

	//замыкание
	multiple := func() int {
		arr[0] *= 2
		return arr[0]
	}

	fmt.Println(multiple())
	fmt.Println(webSites["psychonaut4"])

	fmt.Println(age, num, str, res, len(str))
	fmt.Printf("%.2f \n", num)

	if age < 7 {
		fmt.Println("Ты лох")
	} else if (age < 5) && (age > 12) {
		var grade = age - 5
		fmt.Println("{0} крч ты ещё лох", grade)
	} else {
		fmt.Println("Ну ты красава ёпт")
	}

	switch age {
	case 5:
		fmt.Println("Вам 5 лет")
	case 12:
		fmt.Println("Вам 12 лет")
	case 18:
		fmt.Println("Вам 18 лет")
	case 15:
		fmt.Println("Вам 15 лет")
	default:
		fmt.Println("Неизвестно")
	}

	for age <= 14 {
		age++
		fmt.Println(age)
	}

	for i := 0; i <= len(arr); i++ {
		fmt.Println(i)
	}

	for i, value := range nums {
		fmt.Println(value, i)
	}

	age += summ(b, arr[0])
	fmt.Println(age)
	fmt.Println(summ2(age, b))
}

func summ(num1 int, num2 int) int {
	var res int
	res = num1 + num2
	return res
}

func summ2(num1 int, num2 int) (int, int) {
	var res, res2 int
	res = num1 + num2
	res2 = num1 * num2
	return res, res2
}

func one() {
	fmt.Println("1")
}

func two() {
	fmt.Println("2")
}

//пример с ссылками(this,base)
func pointer(b *int) {
	*b += 2
}

//Cats ёпта бля
type Cats struct {
	name      string
	age       int
	happiness float64
}

func (cat *Cats) test() float64 {
	return float64(cat.age) * cat.happiness
}
