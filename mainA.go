package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
  "strconv"
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
  if sli[0] <= 8 && sli[1] <= 8 {
    fmt.Println("Yay!")
  } else {
    fmt.Println(":(")
  }
}
