package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// Globals

const filepath string = "/tmp/voronoi.ppm"

const nbPoints int = 128
const width, height int = 3840, 1600

// Catppuccin Latte color palette
// var colors = []string{
// 	"220 138 120",
// 	"221 120 120",
// 	"234 118 203",
// 	"136 57 239",
// 	"210 15 57",
// 	"230 69 83",
// 	"254 100 11",
// 	"223 142 29",
// 	"64 160 43",
// 	"23 146 153",
// 	"4 165 229",
// 	"32 159 181",
// 	"30 102 245",
// 	"114 135 253",
// 	"76 79 105",
// 	"108 111 133",
// 	"255 255 255",
// }

// Nord color palette
var colors = []string {
    "46 52 64",
    "59 66 82",
    "46 52 64",
    "59 66 82",
    "46 52 64",
    "59 66 82",
    "46 52 64",
    "59 66 82",
    "129 161 193",
    "255 255 255",
}

// // Gruvbox color palette
// var colors = []string {
//     "235 219 178",
//     "168 153 132",
//     "251 73 52",
//     "184 187 38",
//     "250 189 47",
//     "131 165 152",
//     "211 134 155",
//     "142 192 124",
//     "254 128 25",
//     "204 36 29",
//     "152 151 26",
//     "215 153 33",
//     "69 133 136",
//     "177 98 134",
//     "104 157 106",
//     "214 93 14",
// }

// Structs

type coords struct {
	x, y int
}

// Functions

func writeHeader(f *os.File) {
	_, err := f.Write([]byte(fmt.Sprintf("P3\n%d %d\n%d\n", width, height, 255)))
	if err != nil {
		panic(err)
	}
}

func writeArray(f *os.File, array []int) {
	for _, v := range array {
		_, err := f.Write([]byte(fmt.Sprintf("%s\n", colors[v])))
		if err != nil {
			panic(err)
		}
	}
}

func drawPoint(c coords, array []int, r, color int) {
	for i := -r; i <= r; i++ {
		if c.x+i < 0 || c.x+i >= width {
			continue
		}

		for j := -r; j <= r; j++ {
			if c.y+j < 0 || c.y+j >= height {
				continue
			}

			array[(c.y+j)*width+(c.x+i)] = color
		}
	}
}

func getDistance(a, b coords) int {
	return (a.x-b.x)*(a.x-b.x) + (a.y-b.y)*(a.y-b.y)
}

func getClosest(centers []coords, point coords) int {
	var closestId, closestDistance int = 0, getDistance(point, centers[0])
	for k, v := range centers {
		d := getDistance(point, v)
		if d < closestDistance {
			closestId = k
			closestDistance = d
		}
	}
	return closestId
}

func main() {
	rand.Seed(time.Now().UnixNano())

	f, err := os.Create(filepath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Println("Generating data...")

	var pointsAr = make([]coords, nbPoints)
	var colorAr = make([]int, height*width)

	for k := range pointsAr {
		pointsAr[k] = coords{rand.Intn(width), rand.Intn(height)}
		fmt.Println("Debug: ", pointsAr[k])
	}

	for i := 0; i < height*width; i++ {
		colorAr[i] = getClosest(pointsAr, coords{i % width, i / width}) % (len(colors) - 1)
	}

	// Draw center of each cell
	// for _, coordPoint := range pointsAr {
	//     drawPoint(coordPoint, colorAr, 2, 16)
	// }

	fmt.Println("Writing...")

	writeHeader(f)
	writeArray(f, colorAr)

	fmt.Println("Done!")
}
