package euclidian

import "errors"

type Rectangle struct {
	Length int
	Width  int
}

func NewRectangle(length, width int) (Rectangle, error) {
	if length < width {
		return Rectangle{}, errors.New("the length of a rectangle must be the longer side")
	}

	return Rectangle{
		Length: length,
		Width:  width,
	}, nil
}

func (r Rectangle) IsSquare() bool {
	return r.Length == r.Width
}

func FindBiggestSquare(in Rectangle) Rectangle {
	if in.IsSquare() {
		return in
	}

	rest := in.Length % in.Width
	if rest == 0 {
		// This means, in.Width is a multiple of in.Length.
		fit := in.Length / in.Width
		rec, err := NewRectangle(in.Length/fit, in.Width)
		if err != nil {
			panic("you done fucked up")
		}

		return rec

	}

	rec, err := NewRectangle(in.Width, rest)
	if err != nil {
		panic("you done fucked up")
	}

	return FindBiggestSquare(rec)
}
