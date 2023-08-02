package main

func solution(arr []int) []int {
  stk := make([]int,0)
  var i int
  for i < len(arr){
    if len(stk) == 0{
      stk = append(stk, arr[i])
      i += 1
    }else{
      if stk[len(stk)-1] < arr[i]{
        stk = append(stk, arr[i])
        i += 1
      }else{
        stk = stk[:len(stk)-1]
      }
    }
  }
  return stk
}

func main() {
}