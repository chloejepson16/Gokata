package main

import(
"fmt"
"sort"
"strings"
"unicode"
"math/rand"
"time"
) 

//1. write a program that calculates the factorial of a number
func factorial(n int) int{
	if n < 0{
		return 0
	}
	if n == 1 || n == 0{
		return 1
	}
	result := 1

	for i :=1; i <= n; i++{
		result *= i
	}
	return result
}
//2. implement a function that reverses a string
func reverseString(stringToReverse string) string{
	runes := []rune(stringToReverse)
	length := len(runes)
	for i:= 0; i < length/2; i++{
		runes[i], runes[(length -1) - i] = runes[(length -1)- i], runes[i]
	}

	newStr := string(runes)
	return newStr
}
//3. create a function that checks if a given string is a palindrome
func findPalindrome(potentialPalindrome string) bool{
	reversedString := reverseString(potentialPalindrome)
	if reversedString == potentialPalindrome{
		return true
	}else{
		return false
	}
}
//4. write a program that calculates the nth fib number
func fib(n int) int{
	if n <= 1{
		return n
	}else{
		return fib(n -1) + fib(n -2)
	}
}

//5. implement a function that sorts an array of integers in ascending order
func sortAscending(arr []int) []int{
	sort.Ints(arr)
	return arr
}
//6. create a program that finds the largest element in an array of integers
func findLargest(arr []int) int{
	result := 0
	for i := 0; i < len(arr); i++{
		if(arr[i] > result){
			result= arr[i]
		}
	}
	return result
}
//7. implement a function that removes duplicates from an array/slice
func removeFirstOccurrence(str string, charToRemove rune) string{
	index := strings.IndexRune(str, charToRemove)

	if index == -1{
		return str
	}
	return str[:index] + str[index + 1:]
}
func allValuesEqualToOne(m map[rune]int) bool {
    for _, value := range m {
        if value != 1 {
            return false
        }
    }
    return true
}
func deleteDuplicates(str string) string{
	charCount:= make(map[rune]int)

	for _, char := range str{
		//add one occurance to the character in each map each time found
		charCount[char]++
	}

	for !allValuesEqualToOne(charCount){
		for key, value := range charCount{
			if value > 1{
				str= removeFirstOccurrence(str, key)
				charCount[key]--
			}
			if charCount[key] == 0{
				delete(charCount, key)
			}
		}
	}

	return str
}

//8. write a program to implement a basic calculator that can perform addition, subtraction, mult, and division

func calculator(operation int) float64{
	var num1, num2 float64
	switch operation {
	case 0:
		fmt.Println("Addition (0) selected")
		fmt.Print("Enter first number: ")
		fmt.Scanln(&num1)
		fmt.Print("Enter second number: ")
		fmt.Scanln(&num2)
		result := num1 + num2
		return result
	case 1:
		fmt.Println("Subtraction (1) selected")
		fmt.Print("Enter first number: ")
		fmt.Scanln(&num1)
		fmt.Print("Enter second number: ")
		fmt.Scanln(&num2)
		result := num1 - num2
		return result
	case 2:
		fmt.Println("Multiplicaiton (2) selected")
		fmt.Print("Enter first number: ")
		fmt.Scanln(&num1)
		fmt.Print("Enter second number: ")
		fmt.Scanln(&num2)
		result := num1 * num2
		return result
	case 3:
		fmt.Println("Division (2) selected")
		fmt.Print("Enter first number: ")
		fmt.Scanln(&num1)
		fmt.Print("Enter second number: ")
		fmt.Scanln(&num2)
		result := num1 / num2
		return result
	}
	return 0;
}

//9. find the longest word in the sentence
func findLongestWord(str string) string{
	runes:= []rune(str)
	longestWord:= ""
	currentWord:= ""
	charMaxLength:= 0

	for _, character := range runes{
		if unicode.IsLetter(character){
			currentWord+= string(character)
		}else{
			if(len(currentWord) > charMaxLength){
				charMaxLength= len(currentWord)
				longestWord= currentWord
			}
			currentWord= ""
		}
	}

	if len(currentWord) > charMaxLength {
		longestWord = currentWord
	}

	return longestWord
	
}

//10. create a program that converts a decimal number to binary
func decimalToBinary(n int) string{
	if n == 0{
		return "0"
	}

	binaryNum:= ""
	for n > 0{
		remainder := n % 2
		binaryNum = fmt.Sprintf("%d%s", remainder, binaryNum)
		n= n/2
	}

	return binaryNum;
}
//11. write a function that checks if a number is a prime number
func isPrimeNumber(n int) bool{
	if n<= 1{
		return false
	}
	if n <= 3{
		return true
	}

	if n % 2 == 0 || n % 3 == 0{
		return false
	}

	for i:= 5; i+i <= n; i+=6{
		if (n % i == 0) || ((n % (i + 2)) == 0){
			return false
		}
	}

	return true
}
//12. implement a program that generates a random number within a given range
func generateWithinRange(min int, max int) int{
	rand.Seed(time.Now().UnixNano())
	randomInt:= rand.Intn(max - min + 1) + min
	return randomInt
}

func main() {
    fmt.Println("Hello, World!")
	//1. factorial
	fmt.Println("Please enter the number you would like to find the factorial of: ")
	var number int
	fmt.Scanln(&number)
	factorial_of_number := factorial(number)
	fmt.Println("Factorial: ", factorial_of_number)

	//2. Reversing a string
	str := "hello"
	newString := reverseString(str)
	fmt.Println(newString)
	
	//3. find a palindrome
	palindrome := "madam"
	notPalindrome := "not palindrome"
	fmt.Println("Is madam a plaindrome", findPalindrome(palindrome))
	fmt.Println("Is plaindrome a palindrome", findPalindrome(notPalindrome))

	//4. calculates nth fib number
	fmt.Println("Fib of 10", fib(10))

	//5. sort in ascending order
	values := []int{3, 6, 1, 7, 0}
	sortAscending(values)
	fmt.Println("Sorted values: ", values)

	//6. find largest number
	fmt.Println("Largest value in the list above: ", findLargest(values))

	//7. removing duplicates
	stringWithDuplicates := "aabbcc"
	stringWithDuplicatesRemoved:= deleteDuplicates(stringWithDuplicates)
	fmt.Println("removing duplicates: ", stringWithDuplicatesRemoved)

	//8. calculator
	fmt.Println("Please input a number that corresponds to which operation you would like the calculator to perform,")
	fmt.Println("Addition (0), Subtraciton (1), Multiplication (2), Division (3)")
	var operationIndicator int
	fmt.Scanln(&operationIndicator)
	claculatorResult:= calculator(operationIndicator)
	fmt.Println("Result: ", claculatorResult)

	//9. finding the longest word
	sentence:= "Try to find the longest word."
	longestWord:= findLongestWord(sentence)
	fmt.Println("Finding the longest word: ", longestWord)

	//10. decimal to binary num
	decimalNum:= 12
	binaryNum:= decimalToBinary(decimalNum)
	fmt.Println("Binary num conversion of 12: ", binaryNum)

	//11. function to find if a number is a primary number
	num1:= 2
	num2:= 16
	fmt.Println("is 2 a prime number: ", isPrimeNumber(num1))
	fmt.Println("is 16 a prime number: ", isPrimeNumber(num2))

	//12. Generate a random number between two numbers
	var max int
	var min int
	fmt.Println("To generate a random number between two number please enter a max and a min value")
	fmt.Print("Enter the min value: ")
	fmt.Scanln(&min)
	fmt.Print("Enter the max value: ")
	fmt.Scanln(&max)
	randInt:= generateWithinRange(min, max)
	fmt.Println("Random Value: ", randInt)

}
