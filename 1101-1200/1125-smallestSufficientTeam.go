package main

import "fmt"

func main() {
	fmt.Println(
		smallestSufficientTeam(
			[]string{"java", "nodejs", "reactjs"},
			[][]string{{"java"}, {"nodejs"}, {"nodejs", "reactjs"}},
		),
		[]int{0, 2})
	fmt.Println(
		smallestSufficientTeam(
			[]string{"algorithms", "math", "java", "reactjs", "csharp", "aws"},
			[][]string{{"algorithms", "math", "java"}, {"algorithms", "math", "reactjs"},
				{"java", "csharp", "aws"}, {"reactjs", "csharp"}, {"csharp", "math"}, {"aws", "java"}},
		),
		[]int{1, 2})
	fmt.Println(
		smallestSufficientTeam(
			[]string{"a", "b", "c", "d", "e", "f", "g", "h"},
			[][]string{{"b"}, {"a", "c", "e", "f"}, {"g"}, {"h"}, {"c", "d"}, {"b"}, {"g"}, {"d", "f"}, {"g"}, {"d", "f"}},
		),
		[]int{1, 3, 5, 8, 9}, []int{0, 1, 2, 3, 4})
}

/*
https://leetcode.com/problems/smallest-sufficient-team/

In a project, you have a list of required skills req_skills, and a list of people.  The i-th person people[i] contains
a list of skills that person has.
Consider a sufficient team: a set of people such that for every required skill in req_skills, there is at least
one person in the team who has that skill.  We can represent these teams by the index of each person: for example,
team = [0, 1, 3] represents the people with skills people[0], people[1], and people[3].

Return any sufficient team of the smallest possible size, represented by the index of each person.
You may return the answer in any order.  It is guaranteed an answer exists.

Example 1:

Input: req_skills = ["java","nodejs","reactjs"], people = [["java"],["nodejs"],["nodejs","reactjs"]]
Output: [0,2]

Example 2:

Input: req_skills = ["algorithms","math","java","reactjs","csharp","aws"], people = [["algorithms","math","java"],
["algorithms","math","reactjs"],["java","csharp","aws"],["reactjs","csharp"],["csharp","math"],["aws","java"]]
Output: [1,2]

Constraints:

1 <= req_skills.length <= 16
1 <= people.length <= 60
1 <= people[i].length, req_skills[i].length, people[i][j].length <= 16
Elements of req_skills and people[i] are (respectively) distinct.
req_skills[i][j], people[i][j][k] are lowercase English letters.
Every skill in people[i] is a skill in req_skills.
It is guaranteed a sufficient team exists.
*/

func smallestSufficientTeam(req_skills []string, people [][]string) []int {
	sl := len(req_skills)
	results := make(map[int][]int) // mask => set of people indices
	results[0] = []int{}

	skillsBitMap := make(map[string]int)
	for i, s := range req_skills {
		skillsBitMap[s] = 1 << i
	}

	seenPeople := make(map[int]bool)

	for i, ppl := range people {
		cSkill := 0

		for _, p := range ppl {
			cSkill |= skillsBitMap[p]
		}
		if seenPeople[cSkill] {
			continue
		}
		seenPeople[cSkill] = true

		for rKey, rValue := range results {
			newMask := rKey | cSkill

			rLen := len(rValue) + 1
			nValue, ex := results[newMask]
			if !ex || len(nValue) > rLen {
				// since Golang modifies slice in append operation, we need to copy the slice over to a new one
				nValue = make([]int, rLen, rLen)
				copy(nValue, rValue)
				nValue[rLen-1] = i
				results[newMask] = nValue
			}
		}
	}

	return results[(1<<sl)-1]
}
