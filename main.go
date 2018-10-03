package main

import (
      "fmt"
      "math"

)

func Calculate (pv float64, rate float64, term float64) float64 {
  var pmt float64
  var rateperiod float64

  rateperiod = rate/1200  /* monthly payments only  */

  pmt = pv *(rateperiod/(1-math.Pow((1+rateperiod),-term)))

  return pmt

}

func main() {
  /* first very simple paymet calculation
  parameters Present Value, anual rate in %, number of payments
  */
  fmt.Println(Calculate(25000, 4, 36))


/* stupid comment, just to check the commit*/
/* xxxx */
}
