package main

import "testing"

//
//func TestElevator_getCurFloor(t *testing.T) {
//	tests := []Elevator{
//		{
//			curFloor: 1,
//		},
//		{
//			curFloor: 2,
//		},
//		{
//			curFloor: 3,
//		},
//		{
//			curFloor: 4,
//		},
//		{
//			curFloor: 5,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			s := sugg
//			if got := s.GetSuggestion(tt.args.person); got != tt.want {
//				t.Errorf("GetSuggestion() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func TestElevator_getCurFloor(t *testing.T) {
	type fields struct {
		ele *Elevator
	}
	tests := []struct {
		name         string
		fields       fields
		wantCurFloor int
	}{
		{name: "test-1", fields: fields{ele: &Elevator{curFloor: 1}}, wantCurFloor: 1},
		{name: "test-2", fields: fields{ele: &Elevator{curFloor: 2}}, wantCurFloor: 2},
		{name: "test-3", fields: fields{ele: &Elevator{curFloor: 3}}, wantCurFloor: 3},
		{name: "test-4", fields: fields{ele: &Elevator{curFloor: 4}}, wantCurFloor: 4},
		{name: "test-5", fields: fields{ele: &Elevator{curFloor: 5}}, wantCurFloor: 5},
		{name: "test-6", fields: fields{ele: &Elevator{curFloor: 6}}, wantCurFloor: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Elevator{
				curFloor:  tt.fields.ele.curFloor,
				curStatus: tt.fields.ele.curStatus,
			}
			if gotCurFloor := e.getCurFloor(); gotCurFloor != tt.wantCurFloor {
				t.Errorf("getCurFloor() = %v, want %v", gotCurFloor, tt.wantCurFloor)
			}
		})
	}
}
