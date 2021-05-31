package main

import (
	"fmt"
	"math"
	"strings"
)

/**
 * <p>Implement <code><span>atoi</span></code> which&nbsp;converts a string to an integer.</p>
<p>The function first discards as many whitespace characters as necessary until the first non-whitespace character is found. Then, starting from this character takes an optional initial plus or minus sign followed by as many numerical digits as possible, and interprets them as a numerical value.</p>
<p>The string can contain additional characters after those that form the integral number, which are ignored and have no effect on the behavior of this function.</p>
<p>If the first sequence of non-whitespace characters in str is not a valid integral number, or if no such sequence exists because either str is empty or it contains only whitespace characters, no conversion is performed.</p>
<p>If no valid conversion could be performed, a zero value is returned.</p>
<p><strong>Note:</strong></p>
<ul>
	<li>Only the space character <code>&#39; &#39;</code> is considered a whitespace character.</li>
	<li>Assume we are dealing with an environment that could only store integers within the 32-bit signed integer range: <code>[&minus;2<sup>31</sup>,&nbsp; 2<sup>31&nbsp;</sup>&minus; 1]</code>. If the numerical value is out of the range of representable values,&nbsp;<code>2<sup>31&nbsp;</sup>&minus; 1</code>&nbsp;or <code>&minus;2<sup>31</sup></code>&nbsp;is returned.</li>
</ul>
<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>
<pre>
<strong>Input:</strong> str = &quot;42&quot;
<strong>Output:</strong> 42
</pre>
<p><strong>Example 2:</strong></p>
<pre>
<strong>Input:</strong> str = &quot;   -42&quot;
<strong>Output:</strong> -42
<strong>Explanation:</strong> The first non-whitespace character is &#39;-&#39;, which is the minus sign. Then take as many numerical digits as possible, which gets 42.
</pre>
<p><strong>Example 3:</strong></p>
<pre>
<strong>Input:</strong> str = &quot;4193 with words&quot;
<strong>Output:</strong> 4193
<strong>Explanation:</strong> Conversion stops at digit &#39;3&#39; as the next character is not a numerical digit.
</pre>
<p><strong>Example 4:</strong></p>
<pre>
<strong>Input:</strong> str = &quot;words and 987&quot;
<strong>Output:</strong> 0
<strong>Explanation:</strong> The first non-whitespace character is &#39;w&#39;, which is not a numerical digit or a +/- sign. Therefore no valid conversion could be performed.
</pre>
<p><strong>Example 5:</strong></p>
<pre>
<strong>Input:</strong> str = &quot;-91283472332&quot;
<strong>Output:</strong> -2147483648
<strong>Explanation:</strong> The number &quot;-91283472332&quot; is out of the range of a 32-bit signed integer. Thefore INT_MIN (&minus;2<sup>31</sup>) is returned.
</pre>
<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>
<ul>
	<li><code>0 &lt;= s.length &lt;= 200</code></li>
	<li><code>s</code> consists of English letters (lower-case and upper-case), digits, <code>&#39; &#39;</code>, <code>&#39;+&#39;</code>, <code>&#39;-&#39;</code> and <code>&#39;.&#39;</code>.</li>
</ul>
**/
/**
 * "42"
**/

// Stackは[]runeのエイリアス
type Stack []rune

// Push adds an element
func (s *Stack) Push(v rune) {
	*s = append(*s, v)
}

// Pop removes the top element and return it
func (s *Stack) Pop() (rune, error) {
	if s.Empty() {
		return 0, fmt.Errorf("stack is empty")
	}

	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v, nil
}

// Peek returns the top value
func (s *Stack) Peek() (rune, error) {
	if s.Empty() {
		return 0, fmt.Errorf("stack is empty")
	}

	return (*s)[len(*s)-1], nil
}

// Peek returns the top value
func (s *Stack) Bottom() (rune, error) {
	if s.Empty() {
		return 0, fmt.Errorf("stack is empty")
	}

	return (*s)[0], nil
}

// Size returns the length of stack
func (s *Stack) Size() int {
	return len(*s)
}

// Empty returns true when stack is empty
func (s *Stack) Empty() bool {
	return s.Size() == 0
}

// NewStack generates stack
func NewStack() *Stack {
	s := new(Stack)
	return s
}

func myAtoi(s string) (int, error) {

	MaxInt := int64(math.Pow(2, 31) - 1)
	MinInt := int64(math.Pow(2, 31) * -1)
	MaxDigit := int(math.Log10(float64(MaxInt)))
	runeList := NewStack()
	sList := strings.Fields(s)
	if len(sList) == 0 {
		return 0, nil
	}
	s = sList[0]
	for i, r := range s {
		if (i == 0 && r == '-') || ('0' <= r && r <= '9') {
			runeList.Push(r)
		} else if i == 0 && r == '+' {
			continue
		} else {
			break
		}
	}

	var result int64
	minus := false

	if runeList.Size() != 0 {

		start, err := runeList.Bottom()

		if err != nil {
			return 0, err
		}

		length := runeList.Size()
		if start == '-' {
			minus = true
			length--
		}

		for i := 0; i < length; i++ {

			r, rerr := runeList.Pop()
			if rerr != nil {
				err = rerr
				return 0, err
			}
			number := float64(r - 48)

			if number > 0 && i > MaxDigit {
				if minus {
					result = MinInt * -1
				} else {
					result = MaxInt * -1
				}
			}
			result += int64(number * math.Pow10(i))
		}
	}

	if minus {
		result *= -1
	}

	if result < MinInt {
		result = MinInt
	} else if result > MaxInt {
		result = MaxInt
	}

	return int(result), nil
}
