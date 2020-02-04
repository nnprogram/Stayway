package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
  "strconv"
)

var N int
var M int

// マージソートの関数
func merge_reverse(left, right []int) []int {
    result := []int{}
    for len(left) > 0 && len(right) > 0 {
        if left[0] < right[0] {
            result = append(result, right[0])
            right = right[1:]
        } else {
            result = append(result, left[0])
            left = left[1:]
        }
    }
    return append(result, append(left, right...)...)
}

func merge_sort_reverse(arr []int) []int {
    if len(arr) > 1 {
        q := len(arr) / 2
        arr = merge_reverse(merge_sort_reverse(arr[:q]), merge_sort_reverse(arr[q:]))
    }
    return arr
}

// 入力としてデータ、×1 or ×(-1)か、Mをもらい最大値を返す
// pn = positive and negative
func maximize(matrix [][]int, pn [3]int, M int) int {
  var answer int
  var arr []int
  for l:=0; l<N; l++ {
    arr = append(arr, 0)
    for m:=0; m<3; m++ {
      arr[l] += matrix[l][m] * pn[m]
    }
  }
  merged_arr := merge_sort_reverse(arr)
  for n:=0; n<M; n++ {
    answer += merged_arr[n]
  }
  return answer
}

// main関数
func main(){
  var answer int
  reader := bufio.NewReader(os.Stdin)
  line, _ := reader.ReadString('\n')
  line = strings.TrimSpace(line)
  arr := strings.Split(line, " ")
  N, _ = strconv.Atoi(arr[0])
  M, _ = strconv.Atoi(arr[1])
  matrix := make([][]int, N)
  for j:=0; j<N; j++ {
    matrix[j] = make([]int, 3)
    line, _ = reader.ReadString('\n')
    line = strings.TrimSpace(line)
    arr = strings.Split(line, " ")
    for k:=0; k<3; k++ {
      matrix[j][k], _ = strconv.Atoi(arr[k])
    }
  }
  var candidate_arr []int
  var pn_matrix [8][3]int
  pn_matrix[0] = [3]int{1,1,1}
  pn_matrix[1] = [3]int{1,1,-1}
  pn_matrix[2] = [3]int{1,-1,1}
  pn_matrix[3] = [3]int{-1,1,1}
  pn_matrix[4] = [3]int{1,-1,-1}
  pn_matrix[5] = [3]int{-1,1,-1}
  pn_matrix[6] = [3]int{-1,-1,1}
  pn_matrix[7] = [3]int{-1,-1,-1}
  for i:=0; i<8; i++ {
    candidate_arr = append(candidate_arr, maximize(matrix, pn_matrix[i], M))
  }
  answer = merge_sort_reverse(candidate_arr)[0]
  fmt.Println(answer)
}
