package palindrome

// import "fmt"

func isPalindromeRecursive(value string, i int) bool {

	if i < len(value)/2 {
		firstIndex := value[i]
		lastIndex := value[len(value)-1-i]
		if firstIndex != lastIndex {
			return false
		} else {
			return isPalindromeRecursive(value, i+1)
		}
	} else {
		return true
	}
}

func IsPalindrome(value string) bool {
	// #first solution but not good
	// temp := ""

	// for i := len(value) - 1; i >= 0; i-- {
	// 	temp = temp + string(value[i])
	// }

	// #second solution compare beetwen firstIndex and LastIndex
	// for i := 0; i < len(value)/2; i++ {
	// 	// fmt.Println(fmt.Sprint(i) + " : " + (fmt.Sprint(len(value) - 1 - i)))

	// 	firstIndex := value[i]
	// 	lastIndex := value[len(value)-1-i]
	// 	if firstIndex != lastIndex {
	// 		return false
	// 	}
	// }

	//#third solution
	return isPalindromeRecursive(value, 0)
}
