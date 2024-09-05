package main

type MockSpeaker struct {
	tested bool
}

func (ms *MockSpeaker) Speak(words string) {
	ms.tested = true
}
