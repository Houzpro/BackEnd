package main

import (
	"fmt"
	"math"
	"strconv"
)

func sumTask1_1(number int) (final int) {
	return number%10 + (number/10)%10 + (number/100)%10 + (number/1000)%10

}

func temperatureConvertToFahrenheitTask1_2(degree int) float64 {
	return float64(degree)*1.8 + 32
}

func temperatureConvertToCelsiusTask1_2(degree int) float64 {
	return float64(degree-32) / 1.8
}

func doubleArrayTask1_3(arr [4]int) [4]int {
	arr[0] *= 2
	arr[1] *= 2
	arr[2] *= 2
	arr[3] *= 2
	return arr
}

func concatStringsTask1_4(str [2]string) string {
	return str[0] + " " + str[1]
}

func distanceBetween2PointsTask1_5(x1 int, y1 int, x2 int, y2 int) float64 {
	return math.Sqrt(math.Pow(float64(x2-x1), 2) + math.Pow(float64(y2-y1), 2))
}

func oddOrEvenTask2_1(number int) string {
	if number%2 == 0 {
		return "Четное"
	}
	return "Нечетное"
}

func leapYearTask2_2(year int) string {
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				return "Високосный"
			}
		} else {
			return "Високосный"
		}
	}
	return "Не високосный"
}

func largestNumberTask2_3(n1 int, n2 int, n3 int) int {
	if n1 > n2 && n1 > n3 {
		return n1
	}
	if n2 > n1 && n2 > n3 {
		return n2
	}
	return n3
}

func ageGroupTask2_4(age int) string {
	if age <= 10 {
		return "Ребенок"
	}
	if age >= 11 && age <= 17 {
		return "Подросток"
	}
	if age >= 18 && age <= 59 {
		return "Взрослый"
	}
	return "Пожилой"
}

func delimetersTask2_5(number int) string {
	if number%3 == 0 && number%5 == 0 {
		return "Делится"
	}
	return "Не делится"
}

func factorialTask3_1(number int) (result int) {
	result = 1
	for i := 1; i <= number; i++ {
		result *= i
	}
	return result
}

func fiboTask3_2(number int) []int {
	arr := make([]int, number)
	arr[0] = 0
	arr[1] = 1
	for i := 2; i < number; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr
}

func reverseTask3_3(arr [5]int) [5]int {
	var temp int
	length := len(arr)
	for i := 0; i < length/2; i++ {
		temp = arr[i]
		arr[i] = arr[length-i-1]
		arr[length-i-1] = temp
	}
	return arr
}

func eratosthenesTask3_4(number int) []int {
	arr := make([]bool, number)
	for i := 2; i <= int(math.Sqrt(float64(number)+1)); i++ {
		if !arr[i] {
			for j := i * i; j < number; j += i {
				arr[j] = true
			}
		}
	}
	var primes []int

	for i, isComposite := range arr {
		if i > 1 && !isComposite {
			primes = append(primes, i)
		}
	}

	return primes
}

func sumarrayTask3_5(number [7]int) (result int) {
	for i := 0; i < len(number); i++ {
		result += number[i]
	}
	return result
}

func main() {
	var taskNumber string

	fmt.Print("Выберите номер задания 1.1-1.5, 2.1-2.5, 3.1-3.5\nНапример для задания 1.3: 1.3\nВведите номер задания: ")
	fmt.Scanln(&taskNumber)

	switch taskNumber {
	case "1.1":
		{
			var number int
			fmt.Print("Введите 4-х значное число: ")
			fmt.Scan(&number)
			fmt.Println("Входные данные: ", number, "\nСумма чисел: "+strconv.Itoa(sumTask1_1(number)))
		}
	case "1.2":
		{
			var t string
			var degree int
			fmt.Print("Выполнить перевод в градусы Цельсия или Фаренгейт?(c/f): ")
			fmt.Scanln(&t)
			if t == "f" || t == "F" {
				fmt.Print("Введите целочисленное число: ")
				fmt.Scan(&degree)
				fmt.Print("Входные данные: ", degree, "\nРезультат перевода: "+fmt.Sprintf("%f", temperatureConvertToFahrenheitTask1_2(degree)))
			}
			if t == "c" || t == "C" {
				fmt.Print("Введите целочисленное число: ")
				fmt.Scan(&degree)
				fmt.Print("Входные данные: ", degree, "\nРезультат перевода: "+fmt.Sprintf("%f", temperatureConvertToCelsiusTask1_2(degree)))
			}
		}
	case "1.3":
		{
			var arr [4]int
			fmt.Print("Введите 4 целочисленных числа: ")
			for i := range arr {
				fmt.Scan(&arr[i])
			}
			fmt.Print("Входные данные: ", arr, "\nРезультат удвоения: ", doubleArrayTask1_3(arr))
		}
	case "1.4":
		{
			var str [2]string
			fmt.Print("Введите 2 слова через пробел: ")
			for i := range str {
				fmt.Scan(&str[i])
			}
			fmt.Print("Входные данные: ", str, "\nРезультат сложения строки: "+concatStringsTask1_4(str))
		}
	case "1.5":
		{
			var x1, y1, x2, y2 int
			fmt.Print("Введите координаты двух точек через пробел (x1 y1 x2 y2): ")
			fmt.Scan(&x1)
			fmt.Scan(&y1)
			fmt.Scan(&x2)
			fmt.Scan(&y2)
			fmt.Print("Входные данные: ", x1, " ", y1, " ", x2, " ", y2, "\nРасстояние: ", distanceBetween2PointsTask1_5(x1, y1, x2, y2))
		}
	case "2.1":
		{
			var number int
			fmt.Print("Введите число: ")
			fmt.Scan(&number)
			fmt.Print("Входные данные: ", number, "\nПроверка на четность: ", oddOrEvenTask2_1(number))
		}
	case "2.2":
		{
			var year int
			fmt.Print("Введите год: ")
			fmt.Scan(&year)
			fmt.Print("Входные данные: ", year, "\nВведенный год: ", leapYearTask2_2(year))
		}
	case "2.3":
		{
			var n1, n2, n3 int
			fmt.Print("Введите 3 числа: ")
			fmt.Scan(&n1)
			fmt.Scan(&n2)
			fmt.Scan(&n3)
			fmt.Print("Входные данные: ", n1, " ", n2, " ", n3, "\nНаибольшее среди 3-х: ", largestNumberTask2_3(n1, n2, n3))
		}
	case "2.4":
		{
			var age int
			fmt.Print("Введите возраст (до 10: ребенок, 11-17: подросток, 18-59: взрослый, 59+: пожилой): ")
			fmt.Scan(&age)
			fmt.Print("Входные данные: ", age, "\nВозрастная группа: ", ageGroupTask2_4(age))
		}
	case "2.5":
		{
			var number int
			fmt.Print("Введите число: ")
			fmt.Scan(&number)
			fmt.Print("Входные данные: ", number, "\nПроверка на делимость 3 и 5: ", delimetersTask2_5(number))
		}
	case "3.1":
		{
			var number int
			fmt.Print("Введите число: ")
			fmt.Scan(&number)
			fmt.Print("Входные данные: ", number, "\nФакториал: ", factorialTask3_1(number))
		}
	case "3.2":
		{
			var number int
			fmt.Print("Введите число: ")
			fmt.Scan(&number)
			fmt.Print("Входные данные: ", number, "\nФибонначи: ", fiboTask3_2(number))
		}
	case "3.3":
		{
			var arr [5]int
			fmt.Print("Введите 5 целочисленных числа: ")
			for i := range arr {
				fmt.Scan(&arr[i])
			}
			fmt.Print("Входные данные: ", arr, "\nРеверс массива: ", reverseTask3_3(arr))
		}
	case "3.4":
		{
			var number int
			fmt.Print("Введите число: ")
			fmt.Scan(&number)
			fmt.Print("Входные данные: ", number, "\nСписок простых чисел: ", eratosthenesTask3_4(number))
		}
	case "3.5":
		{
			var arr [7]int
			fmt.Print("Введите 7 чисел: ")
			for i := range arr {
				fmt.Scan(&arr[i])
			}
			fmt.Print("Входные данные: ", arr, "\nСумма чисел массива: ", sumarrayTask3_5(arr))
		}
	}

}
