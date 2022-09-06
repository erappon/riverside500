package main

import (
	"fmt"
	"time"
)

func main() {
	current_time := time.Now()
	fmt.Println(current_time.Format(time.RFC1123))
	var expmileage float64 = 2000 // here 2000 is expected mileage of a chain in kilometers
	var X9 = [24]float64{0, 20, 210, 142, 56, 24, 60, 73, 130, 56, 103, 220, 126}
	//fmt.Printf("X9 bridge series in kilometer ")
	//fmt.Println(X9)
	fmt.Printf("X9 :")
	var sumX9 float64 = 0
	for _, value := range X9 {
		sumX9 += value
	}
	fmt.Printf("%.0f kilometers as on 06 september 2022.\n", sumX9)
	var X9percent float64
	X9percent = (sumX9 / expmileage) * 100
	fmt.Printf("chain X9 reached %d percent wear\n", int64(X9percent))
	if X9percent >= 100 {
		fmt.Printf("Warning! Replace the chain X9 with new one\n")
	}

	var Z9 = [24]float64{0, 66, 32, 45, 77, 73, 60, 192, 60, 105, 161, 234}
	//fmt.Printf("Z9 narrow series in kilometer ")
	//fmt.Println(Z9)
	fmt.Printf("Z9 :")
	var sumZ9 float64 = 0
	for _, value := range Z9 {
		sumZ9 += value
	}
	fmt.Printf("%.0f kilometers as on 24 august 2022.\n", sumZ9)
	var sumfZ9 float64 = float64(sumZ9)
	var Z9percent float64
	Z9percent = (sumfZ9 / expmileage) * 100
	fmt.Printf("chain Z9 reached %d percent wear\n", int64(Z9percent))
	if Z9percent >= 100 {
		fmt.Printf("Warning! Replace the chain Z9 with new one\n")
	}
	{
		fmt.Printf("RIVERSIDE-500 total traveling distance is %.0f kilometers as on 06 september 2022,", sumX9+sumZ9)
	}
	pdate := time.Date(2021, time.Month(11), 10, 0, 0, 0, 0, time.UTC)
	cdate := time.Now()
	y, m, d := calcAge(pdate, cdate)
	fmt.Println(" since")
	fmt.Printf(" %d years, %d months %d days\n", y, m, d)
}

// pdate : purchase date of riverside500
func calcAge(pdate, cdate time.Time) (int, time.Month, int) {
	if cdate.Year() < pdate.Year() {
		return -1, -1, -1
	}
	py, pm, pd := pdate.Date()
	cy, cm, cd := cdate.Date()
	if cd < pd {
		cd += 30
		cm--
	}
	if cm < pm {
		cm += 12
		cy--
	}
	return cy - py, cm - pm, cd - pd

}
