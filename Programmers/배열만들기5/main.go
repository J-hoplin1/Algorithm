package main

import "strconv"

/*
Itoa, Atoi

https://pyrasis.com/book/GoForTheReallyImpatient/Unit47
*/

func solution(intStrs []string, k int, s int, l int) []int {
  result := make([]int,0)
  for _,value := range intStrs{
    if strSlice,_ := strconv.Atoi(value[s : s + l]);  strSlice > k{
      result = append(result, strSlice)
    }
  }
  return result
}