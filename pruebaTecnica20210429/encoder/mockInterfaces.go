package encoder

type MockStep struct {}
func NewMockStep() *MockStep{ return &MockStep{}}
func (mockStep *MockStep) Next(text string) {}