package encoder

import (
	"reflect"
	"testing"
)

func TestNewMorseEncoder(t *testing.T) {
	tests := []struct {
		name string
		want *morseEncoder
	}{
		{
			name: "Test NewMorseEncoder", want: &morseEncoder{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMorseEncoder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMorseEncoder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertToMorse(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "Test convertToMorse",
			args: args{text: "texto"},
			want: convertToMorse("texto"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToMorse(tt.args.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertToMorse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_morseEncoder_EncodeMorse(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test EncodeMorse",
			args: args{text: "texto"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := &morseEncoder{}
			me.EncodeMorse(tt.args.text)
		})
	}
}