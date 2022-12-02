package main

import "testing"

func TestScore(t *testing.T) {
	win, shape := Score("A", "Y")
	if win+shape != 8 {
		t.Errorf("%v %v", win, shape)
	}
}

func TestScoreAll(t *testing.T) {
	win, shape := ScoreAll("A Y\nB X\nC Z")
	if win+shape != 15 {
		t.Errorf("%v %v", win, shape)
	}
}

func TestScore2(t *testing.T) {
	win, shape := Score2("A", "Y")
	if win+shape != 4 {
		t.Errorf("%v %v", win, shape)
	}
	win, shape = Score2("B", "X")
	if win+shape != 1 {
		t.Errorf("%v %v", win, shape)
	}
	win, shape = Score2("C", "Z")
	if win+shape != 7 {
		t.Errorf("%v %v", win, shape)
	}
}

func TestScoreAll2(t *testing.T) {
	win, shape := ScoreAll2("A Y\nB X\nC Z")
	if win+shape != 12 {
		t.Errorf("%v %v", win, shape)
	}
}
