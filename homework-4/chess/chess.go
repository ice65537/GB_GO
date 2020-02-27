package chess

import (
	"fmt"
	"strings"
)

//Point -
type Point struct {
	X int //0=A, 1=B, ... 7=H
	Y int //0..7
}

func (pt Point) String() string {
	s := "ABCDEFGH"
	return string(s[pt.X]) + fmt.Sprint(pt.Y+1)
}

//GetPoint -
func GetPoint(s string) (pnt Point, err error) {
	err = nil
	switch strings.ToLower(string(s[0])) {
	case "a":
		pnt.X = 0
	case "b":
		pnt.X = 1
	case "c":
		pnt.X = 2
	case "d":
		pnt.X = 3
	case "e":
		pnt.X = 4
	case "f":
		pnt.X = 5
	case "g":
		pnt.X = 6
	case "h":
		pnt.X = 7
	default:
		pnt.X = -1
		pnt.Y = -1
		return pnt, fmt.Errorf("В переданной строке первый символ %s не является латинской буквой от A до H", string(s[0]))
	}
	switch strings.ToLower(string(s[1])) {
	case "1":
		pnt.Y = 0
	case "2":
		pnt.Y = 1
	case "3":
		pnt.Y = 2
	case "4":
		pnt.Y = 3
	case "5":
		pnt.Y = 4
	case "6":
		pnt.Y = 5
	case "7":
		pnt.Y = 6
	case "8":
		pnt.Y = 7
	default:
		pnt.X = -1
		pnt.Y = -1
		return pnt, fmt.Errorf("В переданной строке второй символ %s не является цифрой от 1 до 8", string(s[1]))
	}
	return
}

//Desk -
type Desk struct {
	cells [8][8]string
}

//Print -
func (d Desk) Print() {
	for i := 7; i >= 0; i-- {
		fmt.Printf("%d  ", i+1)
		for j := 0; j <= 7; j++ {
			if d.cells[j][i] == "" {
				if (i+j)%2 == 0 {
					fmt.Print(" *")
				} else {
					fmt.Print(" *")
				}
			} else {
				fmt.Print(" " + string(d.cells[j][i][0]))
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("    A B C D E F G H\n")
}

//SetCell -
func (d *Desk) SetCell(s string, isReset bool, pnt Point) error {
	if pnt.X < 0 || pnt.X > 7 || pnt.Y < 0 || pnt.Y > 7 {
		return fmt.Errorf("Заданная позиция [%d,%d] на доске отсутствует", pnt.X, pnt.Y)
	}
	if d.cells[pnt.X][pnt.Y] == "" || isReset {
		d.cells[pnt.X][pnt.Y] = s
	} else {
		return fmt.Errorf("Заданная позиция [%d,%d] на доске уже помечена", pnt.X, pnt.Y)
	}
	return nil
}

//SetCells -
func (d *Desk) SetCells(s string, isReset bool, pnt ...Point) (retPnt []Point) {
	retPnt = make([]Point, 0, 64)
	for _, val := range pnt {
		if err := d.SetCell(s, isReset, val); err == nil {
			retPnt = append(retPnt, val)
		}
	}
	return
}

//KnightMoves -
func (d *Desk) KnightMoves(start []Point, isReset bool, stepLabel string) (retPnt []Point) {
	retPnt = make([]Point, 0, 64)
	for _, val := range start {
		tmpPnt := d.SetCells(stepLabel, isReset,
			Point{val.X - 1, val.Y - 2},
			Point{val.X + 1, val.Y - 2},
			Point{val.X - 1, val.Y + 2},
			Point{val.X + 1, val.Y + 2},
			Point{val.X - 2, val.Y - 1},
			Point{val.X + 2, val.Y - 1},
			Point{val.X - 2, val.Y + 1},
			Point{val.X + 2, val.Y + 1})
		if len(tmpPnt) > 0 {
			retPnt = append(retPnt, tmpPnt...)
		}
	}
	return
}
