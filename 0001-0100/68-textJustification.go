package main

import (
	"fmt"
	"strings"
)

func main() {
	printJustification(fullJustify([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16))
	printJustification(fullJustify([]string{"What", "must", "be", "acknowledgment", "shall", "be"}, 16))
	printJustification(fullJustify([]string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}, 20))
}

func printJustification(s []string) {
	for _, l := range s {
		fmt.Printf("\"%s\" %d\n", l, len(l))
	}
}

/*
https://leetcode.com/problems/text-justification/

Given an array of strings words and a width maxWidth, format the text such that each line has exactly maxWidth
characters and is fully (left and right) justified.

You should pack your words in a greedy approach; that is, pack as many words as you can in each line.
Pad extra spaces ' ' when necessary so that each line has exactly maxWidth characters.

Extra spaces between words should be distributed as evenly as possible. If the number of spaces on a line
does not divide evenly between words, the empty slots on the left will be assigned more spaces than the slots
on the right.

For the last line of text, it should be left-justified and no extra space is inserted between words.

Note:

A word is defined as a character sequence consisting of non-space characters only.
Each word's length is guaranteed to be greater than 0 and not exceed maxWidth.
The input array words contains at least one word.

Example 1:

Input: words = ["This", "is", "an", "example", "of", "text", "justification."], maxWidth = 16
Output:
[
   "This    is    an",
   "example  of text",
   "justification.  "
]

Example 2:

Input: words = ["What","must","be","acknowledgment","shall","be"], maxWidth = 16
Output:
[
  "What   must   be",
  "acknowledgment  ",
  "shall be        "
]

Explanation: Note that the last line is "shall be    " instead of "shall     be", because the last line must
be left-justified instead of fully-justified.
Note that the second line is also left-justified becase it contains only one word.

Example 3:

Input: words = ["Science","is","what","we","understand","well","enough","to","explain","to","a","computer.","Art",
"is","everything","else","we","do"], maxWidth = 20
Output:
[
  "Science  is  what we",
  "understand      well",
  "enough to explain to",
  "a  computer.  Art is",
  "everything  else  we",
  "do                  "
]


Constraints:

1 <= words.length <= 300
1 <= words[i].length <= 20
words[i] consists of only English letters and symbols.
1 <= maxWidth <= 100
words[i].length <= maxWidth
*/

func justify(words []string, currentWidth, maxWidth int) string {
	if len(words) == 1 {
		sb := strings.Builder{}
		sb.WriteString(words[0])
		for currentWidth < maxWidth {
			sb.WriteByte(' ')
			currentWidth++
		}

		return sb.String()
	}

	spaces := (maxWidth-currentWidth)/(len(words)-1) + 1
	r := (maxWidth - currentWidth) % (len(words) - 1)
	s := strings.Builder{}
	for i := 0; i < spaces; i++ {
		s.WriteByte(' ')
	}
	sb := strings.Builder{}

	i := 0
	for ; i < len(words)-1; i++ {
		sb.WriteString(words[i])
		sb.WriteString(s.String())
		if r > 0 {
			sb.WriteByte(' ')
			r--
			maxWidth--
		}
		maxWidth = maxWidth - len(words[i]) - spaces
	}
	for j := maxWidth - len(words[i]); j > 0; j-- {
		sb.WriteByte(' ')
	}
	sb.WriteString(words[i])

	return sb.String()
}

func fullJustify(words []string, maxWidth int) []string {
	res := []string{}
	c := 0
	s := 0
	buf := []string{}
	for _, w := range words {
		c = c + len(w) + s
		if c > maxWidth {
			c = c - len(w) - s
			res = append(res, justify(buf, c, maxWidth))
			buf = []string{}
			c = len(w)
		}

		buf = append(buf, w)
		s = 1
	}
	if len(buf) > 0 {
		sb := strings.Builder{}
		s = 0
		for _, w := range buf {
			if s > 0 {
				sb.WriteByte(' ')
			}
			sb.WriteString(w)
			s = 1
		}
		for c < maxWidth {
			sb.WriteByte(' ')
			c++
		}
		res = append(res, sb.String())
	}

	return res
}
