package encoder

import (
	"reflect"
	"testing"
)

func TestBinaryStep_Next(t *testing.T) {
	type fields struct {
		NextStep      ChainEncoder
		BinaryEncoder BinaryEncoder
	}
	type args struct {
		text string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bstep := &BinaryStep{
				NextStep:      tt.fields.NextStep,
				BinaryEncoder: tt.fields.BinaryEncoder,
			}
			bstep.Next(tt.args.text)
		})
	}
}

func TestNewBinaryStep(t *testing.T) {
	type args struct {
		nextStep      ChainEncoder
		binaryEncoder BinaryEncoder
	}
	tests := []struct {
		name string
		args args
		want *BinaryStep
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBinaryStep(tt.args.nextStep, tt.args.binaryEncoder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBinaryStep() = %v, want %v", got, tt.want)
			}
		})
	}
}