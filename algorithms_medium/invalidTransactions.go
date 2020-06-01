package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	exp := []string{"alice,50,1200,mtv"}
	sort.SliceStable(exp, func(i, j int) bool {
		return exp[i] < exp[j]
	})

	res := invalidTransactions([]string{"alice,20,800,mtv", "alice,50,1200,mtv"})

	fmt.Println(res)
	fmt.Println(exp)

	fmt.Println(invalidTransactionsViaMap([]string{"alice,20,800,mtv", "alice,50,1200,mtv"}))
}

/*
https://leetcode.com/problems/invalid-transactions/

A transaction is possibly invalid if:

the amount exceeds $1000, or;
if it occurs within (and including) 60 minutes of another transaction with the same name in a different city.
Each transaction string transactions[i] consists of comma separated values representing the name, time (in minutes),
amount, and city of the transaction.

Given a list of transactions, return a list of transactions that are possibly invalid.  You may return the answer in any order.

Example 1:

Input: transactions = ["alice,20,800,mtv","alice,50,100,beijing"]
Output: ["alice,20,800,mtv","alice,50,100,beijing"]
Explanation: The first transaction is invalid because the second transaction occurs within a difference of 60 minutes,
have the same name and is in a different city. Similarly the second one is invalid too.

Example 2:

Input: transactions = ["alice,20,800,mtv","alice,50,1200,mtv"]
Output: ["alice,50,1200,mtv"]
Example 3:

Input: transactions = ["alice,20,800,mtv","bob,50,1200,mtv"]
Output: ["bob,50,1200,mtv"]
*/

func invalidTransactions(transactions []string) []string {
	if len(transactions) == 0 {
		return nil
	}

	trs := make([]transaction, len(transactions))

	for i, t := range transactions {
		trs[i] = parseTransaction(t)
	}
	sort.SliceStable(trs, func(i, j int) bool {
		if trs[i].name == trs[j].name {
			return trs[i].time < trs[j].time
		}

		return trs[i].name < trs[j].name
	})

	invalidSet := make(map[string]bool)
	for i := 0; i < len(trs); {
		cName := trs[i].name

		s := i
		for ; i < len(trs) && trs[i].name == cName; i++ {
			if trs[i].amount > 1000 {
				invalidSet[trs[i].str] = true
			}
			for j := s; j < i; j++ {
				if trs[j].amount > 1000 {
					invalidSet[trs[j].str] = true
				}
				if trs[j].time >= trs[i].time-60 && trs[j].city != trs[i].city {
					invalidSet[trs[i].str] = true
					invalidSet[trs[j].str] = true
				}
			}
		}
	}

	invalid := make([]string, len(invalidSet))
	i := 0
	for s := range invalidSet {
		invalid[i] = s
		i++
	}

	return invalid
}

type transaction struct {
	name   string
	time   int
	amount int
	city   string
	str    string
}

func parseTransaction(t string) transaction {
	tr := transaction{}

	parts := strings.Split(t, ",")
	tr.name = parts[0]
	tr.city = parts[3]
	tr.str = t
	tr.time, _ = strconv.Atoi(parts[1])
	tr.amount, _ = strconv.Atoi(parts[2])

	return tr
}

// this is slower
func invalidTransactionsViaMap(transactions []string) []string {
	if len(transactions) == 0 {
		return nil
	}

	trs := make(map[string]map[int][]transaction) // client -> time
	invalidSet := make(map[string]bool)

	for _, t := range transactions {
		tr := parseTransaction(t)
		m := trs[tr.name]
		if m == nil {
			m = make(map[int][]transaction)
			trs[tr.name] = m
		}

		if mt, ex := m[tr.time]; !ex {
			mt = []transaction{tr}
			m[tr.time] = mt
		} else {
			mt = append(mt, tr)
			m[tr.time] = mt
		}

		if tr.amount > 1000 {
			invalidSet[tr.str] = true
		}
	}

	for _, ctrs := range trs {
		for tm, ttr := range ctrs {
			city, sameCity := sameCity(ttr)
			if !sameCity {
				for _, tr := range ttr {
					invalidSet[tr.str] = true
				}
			}

			for t := tm - 60; t <= tm+60; t++ {
				cSameCity := sameCity
				if cttr, ex := ctrs[t]; ex {
					if cSameCity {
						for _, tr := range cttr {
							if tr.city != city {
								cSameCity = false

								for _, tr := range ttr {
									invalidSet[tr.str] = true
								}
								break
							}
						}
					}
					if !cSameCity {
						for _, tr := range cttr {
							invalidSet[tr.str] = true
						}
					}
				}
			}
		}
	}

	invalid := make([]string, len(invalidSet))
	i := 0
	for s := range invalidSet {
		invalid[i] = s
		i++
	}

	return invalid
}

func sameCity(trs []transaction) (string, bool) {
	city := trs[0].city

	for i := 1; i < len(trs); i++ {
		if trs[i].city != city {
			return "", false
		}
	}

	return city, true
}
