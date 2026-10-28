package main

import "fmt"

// checkParentheses checks whether the parentheses in a string are balanced.
// Time complexity: O(n), Space complexity: O(n)
func checkParentheses(input string) bool {
	var stack Stack[rune]

	for _, val := range input {
		if val == '(' {
			stack.push(val)
		} else if val == ')' {
			_, ok := stack.pop()
			if !ok {
				return false
			}
		}
	}

	return stack.size() == 0
}

func main() {
	fmt.Println(checkParentheses("(())")) // true
	fmt.Println(checkParentheses("(()"))  // false
	fmt.Println(checkParentheses("())"))  // false
}
