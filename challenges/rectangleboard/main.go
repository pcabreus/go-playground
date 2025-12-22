package main

/* # Assumptions
 grid size: W x H

 Reactangle: letter, w, y, W, H
   matrix
    dimension: list[w, h]
    sort list of rectangles ->
rectangle
- origin [w, h]
- sizeWidth
- sizeHeight
- letter

itarete over h and over w -> 0
// no rectagle

// add 1 rectagle origin w 1,origin h 1, sizew 2, size h 2 > [T]
// add another to the end of the list rectagle o w 2, h 2, s w 3, h 2 -> [T, L]

// brute-force
iterate over h and over w -> for every position I ask to the list of rectangles in order. If is found it just print the letter
0 0 0 0 0
0 T T 0 0

sort list is [T, L] 0(w*h*l) -> O(n'3)

// another approach is loop into the list in reserve order and update the value in the matrix
[T, L] -> [L, T]
put in the matric the value of the rectangle
iterate over reserve list
mark from origin w, h moving for the size O(n'2)

print
iterate over the matrix and print the value.
*/

/*
	Input
	grid: grid w h -> id or nil
	rect z-order list:

	rect 2: w, h, sw, sh

	Action:
	Add: rect T: w, h, sw, sh
	PopUp: w, h
	Move: w, h, target w , h

	Output:
	Matrix


*/

func main() {

}

func buildMatrix()
