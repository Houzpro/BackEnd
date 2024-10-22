package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {

	cleaned := regexp.MustCompile(`[^\w\s]|[\s]`).ReplaceAllString(s, "")

	cleanedLower := strings.ToLower(cleaned)

	return strings.EqualFold(cleanedLower, reverse(cleanedLower))
}
func ConvertInt(val string, base int, toBase int) (string, error) {
	i, err := strconv.ParseInt(val, base, 64)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, toBase), nil
}
func solveQuadratic(a, b, c float64) (complex128, complex128) {
	d := b*b - 4*a*c

	x1 := (-complex(b, 0) + cmplx.Sqrt(complex(d, 0))) / (2 * complex(a, 0))
	x2 := (-complex(b, 0) - cmplx.Sqrt(complex(d, 0))) / (2 * complex(a, 0))

	return x1, x2
}
func sortAbsValues(arr []int) {

	pairs := make([]struct {
		val int
		abs int
	}, len(arr))
	for i, v := range arr {
		pairs[i] = struct {
			val int
			abs int
		}{v, abs(v)}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].abs < pairs[j].abs
	})

	for i, _ := range arr {
		arr[i] = pairs[i].val
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func findSubstring(mainString, substring string) int {
	if len(substring) == 0 {
		return 0
	}

	length := len(substring)
	for i := 0; i <= len(mainString)-length; i++ {
		match := true
		for j := 0; j < length; j++ {
			if mainString[i+j] != substring[j] {
				match = false
				break
			}
		}
		if match {
			return i
		}
	}
	return -1
}
func mergeSortedArrays(arr1, arr2 []int) []int {

	result := make([]int, 0, len(arr1)+len(arr2))

	i, j := 0, 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] <= arr2[j] {
			result = append(result, arr1[i])
			i++
		} else {
			result = append(result, arr2[j])
			j++
		}
	}

	for i < len(arr1) {
		result = append(result, arr1[i])
		i++
	}

	for j < len(arr2) {
		result = append(result, arr2[j])
		j++
	}

	return result
}
func calculate(num1 string, num2 string, operator string) float64 {
	var result float64
	var err error

	switch operator {
	case "+":
		result, err = add(num1, num2)
	case "-":
		result, err = subtract(num1, num2)
	case "*":
		result, err = multiply(num1, num2)
	case "/":
		result, err = divide(num1, num2)
	case "^":
		result, err = power(num1, num2)
	case "%":
		result, err = modulus(num1, num2)
	default:
		err = fmt.Errorf("недопустимая операция")
	}

	if err != nil {
		return 0
	}

	return result
}

func add(a, b string) (float64, error) {
	num1, err := strconv.ParseFloat(a, 64)
	if err != nil {
		return 0, fmt.Errorf("ошибка при парсинге первого числа: %v", err)
	}
	num2, err := strconv.ParseFloat(b, 64)
	if err != nil {
		return 0, fmt.Errorf("ошибка при парсинге второго числа: %v", err)
	}
	return num1 + num2, nil
}

func subtract(a, b string) (float64, error) {
	num1, err := strconv.ParseFloat(a, 64)
	if err != nil {
		return 0, fmt.Errorf("ошибка при парсинге первого числа: %v", err)
	}
	num2, err := strconv.ParseFloat(b, 64)
	if err != nil {
		return 0, fmt.Errorf("ошибка при парсинге второго числа: %v", err)
	}
	return num1 - num2, nil
}

func multiply(a, b string) (float64, error) {
	num1, err := strconv.ParseFloat(a, 64)
	if err != nil {
		return 0, fmt.Errorf("ошибка при парсинге первого числа: %v", err)
	}
	num2, err := strconv.ParseFloat(b, 64)
	if err != nil {
		return 0, fmt.Errorf("ошибка при парсинге второго числа: %v", err)
	}
	return num1 * num2, nil
}

func divide(a, b string) (float64, error) {
	num1, err := strconv.ParseFloat(a, 64)
	if err != nil {
		return 0, fmt.Errorf("ошибка при парсинге первого числа: %v", err)
	}
	num2, err := strconv.ParseFloat(b, 64)
	if err != nil {
		return 0, fmt.Errorf("ошибка при парсинге второго числа: %v", err)
	}
	if num2 == 0 {
		return 0, fmt.Errorf("деление на ноль")
	}
	return num1 / num2, nil

}

func power(a, b string) (float64, error) {
	num1, err := strconv.ParseFloat(a, 64)
	if err != nil {
		return 0, fmt.Errorf("ошибка при парсинге первого числа: %v", err)
	}
	num2, err := strconv.ParseFloat(b, 64)
	if err != nil {
		return 0, fmt.Errorf("ошибка при парсинге второго числа: %v", err)
	}
	return math.Pow(num1, num2), nil
}

func modulus(a, b string) (float64, error) {
	num1, err := strconv.ParseFloat(a, 64)
	if err != nil {
		return 0, fmt.Errorf("ошибка при парсинге первого числа: %v", err)
	}
	num2, err := strconv.ParseFloat(b, 64)
	if err != nil {
		return 0, fmt.Errorf("ошибка при парсинге второго числа: %v", err)
	}

	result := math.Mod(num1, num2)

	return result, nil
}
func hasIntersection(a1, a2, b1, b2, c1, c2 int) bool {

	if a1 > a2 {
		a1, a2 = a2, a1
	}
	if b1 > b2 {
		b1, b2 = b2, b1
	}
	if c1 > c2 {
		c1, c2 = c2, c1
	}

	maxStart := max(a1, b1, c1)
	minEnd := min(a2, b2, c2)

	return maxStart <= minEnd
}

func findLongestWord(input string) string {
	words := strings.FieldsFunc(input, func(r rune) bool {
		if unicode.IsPunct(r) || r == ' ' {
			return true
		}
		return false
	})

	maxLength := 0
	for _, word := range words {
		if len(word) > maxLength {
			maxLength = len(word)
		}
	}

	for _, word := range words {
		if len(word) == maxLength {
			return word
		}
	}

	return ""
}

func isLeapYear(year int64) bool {

	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

func fibonacci(limit int64) {
	var a, b int64 = 0, 1

	fmt.Print("Fibonacci sequence up to ")
	fmt.Print(limit)
	fmt.Println(": ")

	for a <= limit {
		fmt.Printf("%d ", a)
		a, b = b, a+b
	}

	fmt.Println()
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func findPrimes(start, end int) []int {
	primes := make([]int, 0)
	for num := start; num <= end; num++ {
		if isPrime(num) {
			primes = append(primes, num)
		}
	}
	return primes
}
func countDigits(num int) int {
	count := 0
	for num != 0 {
		num /= 10
		count++
	}
	return count
}

func isArmstrong(num int) bool {
	original := num
	sum := 0
	digits := countDigits(num)

	for num != 0 {
		digit := num % 10
		sum += int(math.Pow(float64(digit), float64(digits)))
		num /= 10
	}

	return original == sum
}

func findArmstrongNumbers(start, end int) []int {
	armstrongs := make([]int, 0)
	for num := start; num <= end; num++ {
		if isArmstrong(num) {
			armstrongs = append(armstrongs, num)
		}
	}
	return armstrongs
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func gcd(a, b int) int {
	if b > a {
		a, b = b, a
	}

	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {

	var taskNumber string

	fmt.Print("Выберите номер задания 1.1-1.5  2.1-2.5 3.1-3.5 \nВведите номер задания: ")
	fmt.Scanln(&taskNumber)

	switch taskNumber {
	case "1.1":
		{
			var num string
			var fromBase int
			var toBase int

			fmt.Print("Enter number to convert: ")
			fmt.Scanln(&num)

			fmt.Print("Enter base to convert from (2-36): ")
			fmt.Scanln(&fromBase)

			fmt.Print("Enter target base (2-36): ")
			fmt.Scanln(&toBase)

			result, err := ConvertInt(num, fromBase, toBase)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Printf("%s in base %d converted to base %d is %s\n", num, fromBase, toBase, result)
		}
	case "1.2":
		{
			var a, b, c float64
			fmt.Print("Введите коэффициент a: ")
			fmt.Scan(&a)
			fmt.Print("Введите коэффициент b: ")
			fmt.Scan(&b)
			fmt.Print("Введите коэффициент c: ")
			fmt.Scan(&c)

			x1, x2 := solveQuadratic(a, b, c)

			fmt.Printf("��орни уравнения: %.2f и %.2f\n", x1, x2)
		}
	case "1.3":
		{
			var n int
			fmt.Print("Введите количество чисел: ")
			fmt.Scan(&n)

			arr := make([]int, n)
			for i := 0; i < n; i++ {
				fmt.Print("Введите число ", i+1, ": ")
				_, err := fmt.Scan(&arr[i])
				if err != nil {
					fmt.Println("Ошибка при вводе числа")
					return
				}
			}

			fmt.Println("\nИсходный массив:")
			fmt.Println(arr)

			sortAbsValues(arr)

			fmt.Println("\nОтсортированный массив по а��солютным значениям:")
			fmt.Println(arr)
		}
	case "1.4":
		{
			arr1 := []int{1, 3, 5, 7}
			arr2 := []int{2, 4, 6, 8}

			mergedArr := mergeSortedArrays(arr1, arr2)
			fmt.Println(mergedArr)
		}
	case "1.5":
		{
			mainString := "Hello, World!"
			substring := "World"

			index := findSubstring(mainString, substring)
			if index != -1 {
				fmt.Printf("Подстрока '%s' найдена в строке '%s' начиная с индекса %d\n",
					substring, mainString, index)
			} else {
				fmt.Println("Подстрока не найдена")
			}
		}
	case "2.1":
		{
			for {
				var num1, num2, operator string
				fmt.Print("Введите первое число: ")
				fmt.Scanln(&num1)
				fmt.Print("Введите оператор (+, -, *, /, ^, %): ")
				fmt.Scanln(&operator)
				fmt.Print("Введите второе число: ")
				fmt.Scanln(&num2)

				result := calculate(num1, num2, operator)
				if result != 0 {
					fmt.Printf("%s %s %s = %.2f\n", num1, operator, num2, result)
				} else {
					fmt.Println("Ошибка выполнения операции.")
				}
			}
		}
	case "2.2":
		{
			testCases := []string{
				"A man, a plan, a canal: Panama",
				"Abba",
			}

			for _, tc := range testCases {
				result := isPalindrome(tc)
				var isPalindrome string
				if result {
					isPalindrome = "a"
				} else {
					isPalindrome = "not"
				}
				fmt.Printf("%s is %s a palindrome\n", tc, isPalindrome)
			}
		}
	case "2.3":
		{

			testCases := []struct {
				a1, a2, b1, b2, c1, c2 int
				expected               bool
			}{
				{1, 2, 3, 4, 5, 6, false},
				{1, 2, 2, 3, 3, 4, false},
				{0, 0, 0, 0, 0, 0, true},
				{-1, 1, 0, 1, -1, 0, true},
				{5, 5, 5, 5, 5, 5, true},
			}

			for _, tc := range testCases {
				fmt.Printf("Test Case: A1A2(%d,%d), B1B2(%d,%d), C1C2(%d,%d)\n", tc.a1, tc.a2, tc.b1, tc.b2, tc.c1, tc.c2)
				result := hasIntersection(tc.a1, tc.a2, tc.b1, tc.b2, tc.c1, tc.c2)
				fmt.Printf("Expected: %t, Actual: %t\n\n", tc.expected, result)
			}
		}
	case "2.4":
		{
			sentence := "No punctuation here"
			fmt.Println(findLongestWord(sentence))
		}
	case "2.5":
		{
			var year int64
			fmt.Print("Введите год : ")

			fmt.Scanln(&year)
			fmt.Println("Год", year, " ", isLeapYear(year))
		}
	case "3.1":
		{
			var limit int64
			fmt.Print("Enter a limit: ")
			fmt.Scanln(&limit)
			fibonacci(limit)
		}
	case "3.2":
		{
			var start, end int
			fmt.Print("Enter start of range: ")
			fmt.Scanln(&start)
			fmt.Print("Enter end of range: ")
			fmt.Scanln(&end)

			primes := findPrimes(start, end)
			fmt.Printf("Prime numbers between %d and %d:\n", start, end)
			fmt.Println(primes)
		}
	case "3.3":
		{
			var start, end int
			fmt.Print("Enter start of range: ")
			fmt.Scanln(&start)
			fmt.Print("Enter end of range: ")
			fmt.Scanln(&end)

			armstrongs := findArmstrongNumbers(start, end)
			fmt.Printf("Armstrong numbers between %d and %d:\n", start, end)
			fmt.Println(armstrongs)
		}
	case "3.4":
		{
			var input string
			fmt.Print("Enter a string: ")
			fmt.Scanln(&input)

			reversed := reverse(input)
			fmt.Printf("Reversed string: %s\n", reversed)
		}
	case "3.5":
		{
			var num1, num2 int
			fmt.Print("Enter first number: ")
			fmt.Scan(&num1)
			fmt.Print("Enter second number: ")
			fmt.Scan(&num2)

			result := gcd(num1, num2)
			fmt.Printf("GCD of %d and %d is: %d\n", num1, num2, result)
		}

	}
}
