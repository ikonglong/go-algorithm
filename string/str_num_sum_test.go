package string

import (
	"fmt"
	"strings"
	"testing"
)

func digitSum(s string, k int) string {
	if len(s) <= k {
		return s
	}

	for {
		var sum byte = 0
		var b strings.Builder
		for i := 0; i < len(s); i++ {
			sum += byte(s[i] - '0')
			if (i+1)%k == 0 || i == len(s)-1 {
				fmt.Fprintf(&b, "%d", sum)
				sum = 0
			}
		}
		if b.Len() <= k {
			return b.String()
		} else {
			s = b.String()
		}
	}
}

func TestDigitSum(t *testing.T) {
	testCases := []struct {
		s        string
		k        int
		expected string
	}{
		{"11111222223", 3, "135"},
		{"00000000", 3, "000"},
	}

	for _, tc := range testCases {
		result := digitSum(tc.s, tc.k)
		if result != tc.expected {
			t.Errorf("digitSum(%s, %d) = %s; expected %s", tc.s, tc.k, result, tc.expected)
		}
	}
}
