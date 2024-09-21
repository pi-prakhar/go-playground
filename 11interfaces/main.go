package main

import (
	"fmt"
)

type Speaker interface {
	Speak(string)
}

type JustSpeak func(string)

func (j JustSpeak) Speak(words string) {
	j(words)
}

type Human struct {
	name string
}

func (h *Human) Speak(words string) {
	fmt.Printf("This is %s, a Human, ", h.name)
	fmt.Printf("%s : %s\n", h.name, words)
}

type Cat struct {
	name string
}

func (c *Cat) Speak(words string) {
	fmt.Printf("This is %s, a Cat, ", c.name)
	fmt.Printf("%s : %s \n", c.name, words)
}

type SpeakerMux struct {
}

func NewSpeakerMux() *SpeakerMux {
	return &SpeakerMux{}
}

func (s *SpeakerMux) Handle(speaker Speaker) {
	switch speaker.(type) {
	case *Human:
		speaker.Speak("Hello")
	case *JustSpeak:
		speaker.Speak("#$%^&")
	case *Cat:
		speaker.Speak("Meow!!")
	default:
		break
	}
}
func (s *SpeakerMux) HandleFunc(speak func(string)) {
	speak("Speaking")
}

func main() {
	var speaker1 Speaker = &Cat{
		name: "Garfield",
	}

	speaker2 := Human{
		name: "Jake",
	}

	greet := func(words string) {
		fmt.Printf("This is just a Speaker, ")
		fmt.Printf("Speaking: %s\n", words)

	}

	speaker3 := JustSpeak(greet)

	sm := NewSpeakerMux()

	sm.HandleFunc(speaker1.Speak)

	sm.Handle(speaker1)
	sm.Handle(&speaker2)
	sm.Handle(&speaker3)

	// speaker3.Speak("########")
	// speaker1.Speak("Meow!!")
	// speaker2.Speak("Hello!!")
}
