package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
  "strconv"
  "math"
)

func main(){
  var sli []int
  reader := bufio.NewReader(os.Stdin)
  line, _ := reader.ReadString('\n')
  line = strings.TrimSpace(line)
  arr := strings.Split(line, " ")
  for _, str := range arr {
    value, _ := strconv.Atoi(str)
    sli = append(sli, value)
  }
  answer := int(math.Pow(100, float64(sli[0]))) * sli[1]
  // answer :=  math.Pow(100, sli[0]) * sli[1]
  fmt.Println(answer)
}
