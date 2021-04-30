package encoder

type MorseStep struct {
	NextStep ChainEncoder
	MorseEncoder MorseEncoder
}

func NewMorseStep(nextStep ChainEncoder,
	morseEncoder MorseEncoder) *MorseStep {
	return &MorseStep{
		NextStep:     nextStep,
		MorseEncoder: morseEncoder,
	}
}

func (mstep *MorseStep) Next(text string) {
	mstep.MorseEncoder.EncodeMorse(text)
	if mstep.NextStep != nil {
		mstep.NextStep.Next(text)
	}
}