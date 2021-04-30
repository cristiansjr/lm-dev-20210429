package encoder

type BinaryStep struct {
	NextStep ChainEncoder
	BinaryEncoder BinaryEncoder
}

func NewBinaryStep(nextStep ChainEncoder,
	binaryEncoder BinaryEncoder) *BinaryStep {
	return &BinaryStep{
		NextStep:      nextStep,
		BinaryEncoder: binaryEncoder,
	}
}

func (bstep *BinaryStep) Next(text string) {
	bstep.BinaryEncoder.EncodeBinary(text)
	if (bstep.NextStep != nil) {
		bstep.NextStep.Next(text)
	}
}