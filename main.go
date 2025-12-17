package main

import (
	"fmt"
	"sort"
	"time"
)

func lengthOfLongestSubstring(s string) int {
	length := len(s)
	if length == 0 {
		return 0
	}
	longest := 0
	// abcabcbb
	for _, i := range s { // TODO: Or remaining chars are same size than the longest // a
		seen := map[byte]struct{}{s[i]: {}} // seen -> a, b , c
		charLength := 1
		for j := i + 1; j < length; j++ {
			currentChar := s[j]                 // b
			if _, ok := seen[currentChar]; ok { // a, b, c
				break
			}

			seen[currentChar] = struct{}{} //c <- c

			charLength++ // 3
		}

		if charLength > longest {
			longest = charLength
		}
	}

	return longest
}

func main() {

	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring(""))
	fmt.Println(lengthOfLongestSubstring(" "))

	// Map
	x := map[string]string{"x": "pepe"}
	fmt.Println(x)
	fmt.Println(x["x"])

	// Add y : "juan to X"
	x["y"] = "juan"
	fmt.Println(x)
	fmt.Println(x)

	fmt.Println(len(x))

	delete(x, "x")
	fmt.Println(x)

	// Slice
	c := []int{1, 2, 3}

	fmt.Println(c)

	c = append(c, 4)
	fmt.Println(c)

	// delete index 1
	c = append(c[:1], c[2:]...)
	fmt.Println(c)

	// Using make with slices and maps
	s := make([]int, 10)
	fmt.Println(len(s))
	s = append(s, 10)
	fmt.Println(s)

	m := make(map[string]int)
	fmt.Println(len(m))
	m["a"] = 1
	fmt.Println(m)
}

func rectangleArea(rectangles [][]int) int {
	const MOD = int(1e9 + 7)

	n := len(rectangles)
	// 1. Recoger y ordenar todos los X
	xs := make([]int, 0, 2*n)
	for _, r := range rectangles {
		xs = append(xs, r[0])
		xs = append(xs, r[2])
	}
	sort.Ints(xs)

	// Eliminar duplicados en xs
	uniqueXs := xs[:0]
	for i, x := range xs {
		if i == 0 || x != xs[i-1] {
			uniqueXs = append(uniqueXs, x)
		}
	}
	xs = uniqueXs

	var area int64 = 0

	// 2. Barrer por franjas entre xs[i] y xs[i+1]
	for i := 0; i < len(xs)-1; i++ {
		xStart := xs[i]
		xEnd := xs[i+1]
		dx := int64(xEnd - xStart)
		if dx == 0 {
			continue
		}

		// 3. Recoger intervalos de Y de los rectángulos activos en esta franja
		intervals := make([][2]int, 0)
		for _, r := range rectangles {
			x1, y1, x2, y2 := r[0], r[1], r[2], r[3]
			if x1 <= xStart && xStart < x2 {
				intervals = append(intervals, [2]int{y1, y2})
			}
		}

		if len(intervals) == 0 {
			continue
		}

		// 4. Unir intervalos de Y
		sort.Slice(intervals, func(i, j int) bool {
			if intervals[i][0] == intervals[j][0] {
				return intervals[i][1] < intervals[j][1]
			}
			return intervals[i][0] < intervals[j][0]
		})

		totalY := int64(0)
		curStart, curEnd := intervals[0][0], intervals[0][1]

		for j := 1; j < len(intervals); j++ {
			y1, y2 := intervals[j][0], intervals[j][1]
			if y1 > curEnd {
				// No se solapan: sumamos el intervalo actual y empezamos uno nuevo
				totalY += int64(curEnd - curStart)
				curStart, curEnd = y1, y2
			} else if y2 > curEnd {
				// Se solapan parcialmente: extendemos el intervalo actual
				curEnd = y2
			}
			// Si y2 <= curEnd, el intervalo está completamente cubierto -> no hacemos nada
		}
		// Añadir el último intervalo
		totalY += int64(curEnd - curStart)

		// 5. Añadir el área de esta franja
		area = (area + dx*totalY) % int64(MOD)
	}

	return int(area % int64(MOD))
}

// To execute Go code, please declare a func main() in a package "main"

// Given a list of bookings with start_time, end_time and user_id,
// return the number of concurrent bookings per user.”

// concurrent is when a Booking start before the other ends
type Booking struct {
	StartTime time.Time
	EndTime   time.Time
	UserID    int
}

func process(bookings []Booking, userID int) int {
	// validate input

	// return 0 if booking lenght is 0 || out-of-scope

	var concurrent int
	// perform a general search
	userBookings := getBookingsByUserID(bookings, userID)
	max := len(userBookings)
	for i := 0; i < max; i++ {
		for j := i + 1; j < max; j++ {
			if userBookings[i].StartTime.Before(userBookings[j].StartTime) && userBookings[i].EndTime.After(userBookings[j].StartTime) {
				concurrent++
			}
		}
	}

	return concurrent
}

// Get the bookings for a given user id.
func getBookingsByUserID(bookings []Booking, userID int) []Booking {
	indexed := []Booking{}

	for _, booking := range bookings {
		if booking.UserID == userID {
			indexed = append(indexed, booking)
		}
	}

	return indexed
}
