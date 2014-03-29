package main

import (
  "fmt"
  "strconv"
  "os"
)

// It's a example without recursive function
func not_recursive_3n1(f int) int {
  var counter int = 1

  for {
    if f % 2 == 0 {
      f = f / 2
    } else {
      f = (f * 3) + 1
    }

    counter = counter + 1

    // when f is equal to 1, program exit
    if f == 1 {
      break
    }
  }

  return counter
}

func main() {
  from, e := strconv.Atoi(os.Args[1])
  to, e := strconv.Atoi(os.Args[2])
  aux := from

  if e != nil {
    fmt.Println(e)
  }

  max_circle_lenght := 0

  for aux <= to {
    var circle_lenght int = not_recursive_3n1(aux)

    if circle_lenght > max_circle_lenght {
      max_circle_lenght = circle_lenght
    }

    aux = aux + 1
  }

  fmt.Println(from, to, max_circle_lenght)
}
