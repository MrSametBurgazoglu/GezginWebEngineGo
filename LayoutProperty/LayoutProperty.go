package LayoutProperty

type LayoutProperty struct {
	Parent           *LayoutProperty
	Children         []*LayoutProperty
	Position         string
	XPosition        int
	YPosition        int
	ContentXPosition int
	ContentYPosition int
	Width            int
	Height           int
	ContentWidth     int
	ContentHeight    int
	PaddingLeft      int
	PaddingRight     int
	PaddingTop       int
	PaddingBottom    int
	MarginLeft       int
	MarginRight      int
	MarginTop        int
	MarginBottom     int
}
