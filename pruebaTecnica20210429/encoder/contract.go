package encoder

type Encoder interface {
	Encode(text string)
}

type ChainEncoder interface {
	Next(text string)
}

type MorseEncoder interface {
	EncodeMorse(text string)
}

type BinaryEncoder interface {
	EncodeBinary(text string)
}