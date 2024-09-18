package main

import(
"fmt"
"sort"
"strings"
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

}
