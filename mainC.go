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
  reader.ReadString('\n')

  line, _ := reader.ReadString('\n')
  line = strings.TrimSpace(line)
  arr := strings.Split(line, " ")
  for _, str := range arr {
    value, _ := strconv.Atoi(str)
    sli = append(sli, value)
  }
  var answer int
  for _, element := range sli {
    for {
      if element%2 == 1 {
        break
      }
      element /= 2
      answer += 1
    }
  }
  fmt.Println(answer)
}
