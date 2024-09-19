package main

import(
"fmt"
"sort"
"strings"
"unicode"
"math/rand"
"time"
"math"
"io/ioutil"
"strconv"
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
	if min > max{
		rand.Seed(time.Now().UnixNano())
		randomInt:= rand.Intn(max - min + 1) + min
		return randomInt
	}
	return 0
}

//13. num vowels
func isVowel(ch rune) bool{
	ch= unicode.ToLower(ch)
	return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u'
}
func numVowels(str string) int{
	numVowels:= 0
	for _, ch := range str{
		if(isVowel(ch)){
			numVowels++
		}
	}

	return numVowels
}

//14. find the second smallest element in an array/slice of integers
func secondSmallest(arr []int) int{
	if len(arr) < 2{
		return 0
	}

	smallest:= math.MaxInt64
	secondSmallest:= math.MaxInt64

	for _, num := range arr{
		if num < smallest{
			secondSmallest= smallest
			smallest= num
		}else if num < secondSmallest && num != smallest{
			secondSmallest= num
		}
	}

	if secondSmallest == math.MaxInt64{
		return 0
	}

	return secondSmallest
}

//15. checks if two strings are anagrams of each other
func isAnagram(str1 string, str2 string) bool{
	charCountString1:= make(map[rune] int)
	for _, char:= range str1{
		charCountString1[char]++
	}

	charCountString2:= make(map[rune] int)
	for _, char:= range str2{
		charCountString2[char]++
	}

	if len(charCountString2) != len(charCountString1){
		return false
	}

	for key, value:= range charCountString1{
		if value2, exists:= charCountString2[key]; !exists || value2 != value{
			return false
		}
	}

	return true

}

//16. reads a file and counts the number of words in it
func numWords(str string) int{
	runes:= []rune(str)
	inWord:= false
	wordCount:= 0

	for _, char:= range runes{
		if unicode.IsLetter(char){
			if !inWord{
				wordCount++
				inWord= true
			}
		}else{
			inWord= false
		}
	}
	return wordCount
}
func numWordsInFile(fileName string) int{
	content, err:= ioutil.ReadFile(fileName)
	if err != nil{
		fmt.Println("unable to read file")
	}

	fileContents:= string(content)
	numWordsInFile:= numWords(fileContents)
	return numWordsInFile
}

//17. create a functin that merges two sorted ararys into a single array
func mergeSort(arr1 []int, arr2 []int) []int {
	merged:= make([]int, 0, len(arr1) + len(arr2))
	i, j:= 0, 0

	for i < len(arr1) && j < len(arr2){
		if arr1[i] < arr2[j]{
			merged= append(merged, arr1[i])
			i++
		}else{
			merged= append(merged, arr2[j])
			j++
		}
	}

	//adding any left over from previous loop condition
	for i < len(arr1){
		merged= append(merged, arr1[i])
		i++
	}
	for j < len(arr2){
		merged= append(merged, arr2[j])
		j++
	}

	return merged
}

//18. calculate the sum of digits in a given number
func sumOfDigits(num int) int{
	str:= strconv.Itoa(num)
	sum:= 0
	for _, char:= range str{
		intValue:= int(char - '0')
		sum= sum + intValue
	}

	return sum
}

//19. convert roman numeral to int
func valueOfNumeral(char byte) int{
	if char == 'I'{return 1}
	if char == 'V'{return 5}
	if char == 'X'{return 10}
	if char == 'L'{return 50}
	if char == 'C'{return 100}
	if char == 'D'{return 500}
	if char == 'M'{return 1000}
	return 0
}

func numeralToDecimal(str string) int{
	sum:= 0

	for i:= 0; i < len(str); i++{
		val1:= valueOfNumeral(str[i])

		//comparison 1
		if i + 1 < len(str){
			val2:= valueOfNumeral(str[i+1])

			if val1 >= val2{
				sum+= val1
			}else{
				sum+= (val2 - val1)
				i++
			}
		}else{
			sum+= val1
		}
	}
	return sum
}

func main() {
	/*
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

	//13. write a function that counts the number of vowels in a string
	stringToCountVowels:= "How many vowels in this string?"
	fmt.Println("Num vowels in this string: ", numVowels(stringToCountVowels))

	*/

	//14. create a program that finds the second smallest element in an array/slice of integers
	arr:= []int{9, 20, 22, 1, 4, 15}
	fmt.Println("second smallest in the array: ", secondSmallest(arr))

	//15. implement a function that checks if two strings are anagrams of each other
	iceman:= "iceman"
	cinema:= "cinema"
	fmt.Println("are iceman and cinema anagrams: ", isAnagram(iceman, cinema))

	//16. write a program that reads a file and counts the number of words in it
	fmt.Println("number of words in the text example file: ", numWordsInFile("textExample.txt"))

	//17. create a function that merges two sorted arrays into a single sorted array
	arr1 := []int{1, 3, 5, 7, 9}
    arr2 := []int{2, 4, 6, 8, 10}
	mergedArr:= mergeSort(arr1, arr2)
	fmt.Println("merged array: ", mergedArr)

	//18. implement a program that calculates the sum of digits in a given number
	longNumber:= 12923
	sumDigits:= sumOfDigits(longNumber)
	fmt.Println("sum of digits in 12923: ", sumDigits)

	//19. write a function that converts a roman numeral to an integer
	romanNumeral:= "IV"
	fmt.Println("value of IV: ", numeralToDecimal(romanNumeral))
	
	//20. create a program that sorts a slice of strings based on thir length
	//21. implement a function that generates a multiplication table up to a given number
}
