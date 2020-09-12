package main

import "time"
import "fmt"

func main() {
  // fmt.Printf(time.Date(0,0,0,0,0,0))
  fmt.Printf(time.Unix(0,0).UTC().String())
  fmt.Println(onSameDay(time.Date(2020,1,1,0,0,0,0,time.UTC), time.Date(2020,1,1,0,0,1,0,time.UTC)))
  fmt.Println(onSameDay(time.Date(2020,1,1,0,0,0,0,time.UTC), time.Date(2020,1,2,0,0,0,0,time.UTC)))
}

func onSameDay(timeA time.Time, timeB time.Time) bool {
  // rounded := time.Date(toRound.Year(), toRound.Month(), toRound.Day(), 0, 0, 0, 0, toRound.Location())
  return( timeA.Year() == timeB.Year() ) && ( timeA.Month() == timeB.Month() ) && ( timeA.Day() == timeB.Day() )
}
