package encoder


type Process struct {

}
func NewProcess() *Process{
	return &Process{}
}
func (ep *Process) Encode(text string) {
	morseEncoder := NewMorseEncoder()
	morseStep := NewMorseStep(nil, morseEncoder)
	binaryEncoder := NewBinaryEncoder()
	binaryStep := NewBinaryStep(morseStep, binaryEncoder)
	binaryStep.Next(text)
}