package main

func solution(my_strings []string, parts [][]int) string {
  var result string
  for i,v := range my_strings{
    s,e := parts[i][0], parts[i][1]
    result += v[s:e+1]
  }
  return result
}

func main() {
}