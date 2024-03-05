package main

import "fmt"

func main() {
	//Task 1
	fmt.Println("Task1:")

	str1 := "Hello "
	str2 := "World!"

	result := str1 + str2

	//Task2
	fmt.Println("Task2:")

	num1 := 17
	num2 := 135

	formattedResult := fmt.Sprintf("%s %d", result, num1+num2)
	fmt.Println(formattedResult)

	//Task3
	fmt.Println("Task3:")

	addition := num1 + num2
	fmt.Println("Addition:", addition)

	substraction := num1 - num2
	fmt.Println("Substraction:", substraction)

	multiplication := num1 * num2
	fmt.Println("Multiplication", multiplication)

	division := float64(num1) / float64(num2)
	fmt.Println("Division:", division)

	//Task4
	fmt.Println("Task4:")

	var float1, float2 float64

	fmt.Println("Enter your float numbers:")
	fmt.Scanln(&float1)
	fmt.Scanln(&float2)

	if float1 > float2 {
		fmt.Println("float1 > float2!")
	} else if float1 < float2 {
		fmt.Println("float1 < float2!")
	} else {
		fmt.Println("float1 = float2!")
	}

	if float1 >= float2 || float1 <= float2 {
		fmt.Println("float1 >= float2 || float1 <= float2!")
	} else if float1 == 35.1 && float2 == 26.4 {
		fmt.Println("float1 >= float2 || float1 <= float2!")
	} else {
		fmt.Println("Invalid operation!")
	}

	//Task5
	fmt.Println("Task5:")

	str := "programmer"

	switch str {
	case "programmer":
		fmt.Println("It's a programmer")
	case "footballer":
		fmt.Println("It's a footballer")
	case "doctor":
		fmt.Println("It's a doctor")
	default:
		fmt.Println("I don't know this type of work")
	}

	//Task6
	fmt.Println("Task6:")

	switch {
	case num1 == 17:
		fmt.Println("First case")
		fallthrough
	case num2 == 135:
		fmt.Println("Second case")
	case num1 == num2:
		fmt.Println("Third case")
	}

	//Task7
	fmt.Println("Task7:")

	fruit := "kiwi"

	switch fruit {
	case "banana":
		fmt.Println("This is a banana")
	case "orange":
		fmt.Println("This is an orange")
	case "apple":
		fmt.Println("This is an apple")
	default:
		fmt.Println("I don't know this fruit")
	}

	//Task 8
	fmt.Println("Task8:")

	array := [5]int{1, 2, 3, 4, 5}
	array[3]++
	array[4]--
	fmt.Println(array)

	//task9
	fmt.Println("Task9:")

	makeSlice := make([]string, 0)
	makeSlice = append(makeSlice, result)
	fmt.Println(makeSlice)

	fruitsSlice := []string{"apple", "melon", "banana", "kiwi"}
	fmt.Println("Length of Slice2:", len(fruitsSlice))

}
