package encoder

import (
	"reflect"
	"testing"
)

func TestNewBinaryEncoder(t *testing.T) {
	tests := []struct {
		name string
		want *binaryEncoder
	}{
		{
			name: "Test NewBinaryEncoder",
			want: &binaryEncoder{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBinaryEncoder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBinaryEncoder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_binaryEncoder_EncodeBinary(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test EncodeBinary",
			args: args{text: "Kennia"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			be := &binaryEncoder{}
			be.EncodeBinary(tt.args.text)
		})
	}
}

func Test_convertToBinary(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test convertToBinary Success",
			args: args{text: "Kennia"},
			want: convertToBinary("Kennia"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToBinary(tt.args.text); got != tt.want {
				t.Errorf("convertToBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}