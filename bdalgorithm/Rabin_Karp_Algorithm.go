package main

import "fmt"

func main() {
	var T = "Rabin–Karp string search algorithm: Rabin-Karp"
	var P = "Rabin"
	var q = 101 // A prime number
	var d = 16
	Rabin_Karp_search(T, P, d, q)

	fmt.Println("pause")

}

func Rabin_Karp_search(T, P string, d, q int) {
	m := len(P)
	n := len(T)
	var i, j int
	p := 0 // hash value for pattern
	t := 0 // hash value for txt
	h := 1

	TRune := []rune(T)
	PRune := []rune(P)

	// The value of h would be "pow(d, M-1)%q"
	for i = 0; i < m-1; i++ {
		h = (h * d) % q
	}

	// Calculate the hash value of pattern and first window of text
	for i = 0; i < m; i++ {
		p = (d*p + PRune[i]) % q
		t = (d*t + TRune[i]) % q
	}

	// Slide the pattern over text one by one
	for i = 0; i <= n-m; i++ {

		// Chaeck the hash values of current window of text and pattern
		// If the hash values match then only check for characters on by one
		if p == t {
			/* Check for characters one by one */
			for j = 0; j < m; j++ {
				if TRune[i+j] != PRune[j] {
					break
				}
			}
			// if p == t and pat[0...M-1] = txt[i, i+1, ...i+M-1]
			if j == m {
				fmt.Println("Pattern found at index :", i)
			}

		}

		// Calulate hash value for next window of text: Remove leading digit,
		// add trailing digit
		if i < n-m {
			t = (d*(t-TRune[i]*h) + TRune[i+m]) % q

			// We might get negative value of t, converting it to positive
			if t < 0 {
				t = (t + q)
			}

		}
	}

}
