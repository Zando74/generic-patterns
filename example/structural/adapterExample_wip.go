package main

import (
	"fmt"
	"generic-patterns/structural"
)

type English string
type French string

type EnglishSpeaker interface {
	SpeakEnglish(message string) English
}

type EnglishExecutor struct {
	englishSpeaker EnglishSpeaker
}

func (e *EnglishExecutor) Execute(message string) {
	messageInEnglish := e.englishSpeaker.SpeakEnglish(message)
	// Do something with the English message
	fmt.Println(messageInEnglish)
}

type FrenchSpeaker struct {
	structural.Adaptable
}

func (f *FrenchSpeaker) SpeakFrench(message string) French {
	return French(message)
}

type FrenchToEnglishAdapter struct {
	fr *FrenchSpeaker
}

func (adapter *FrenchToEnglishAdapter) SpeakEnglish(message string) English {
	messageInFrench := adapter.fr.SpeakFrench(message)
	translatedMessage := English("Unknown")

	if messageInFrench == "Bonjour" {
		translatedMessage = English("Hello")
	}
	return translatedMessage
}

func MainAdapterExample() {

	frenchSpeaker := &FrenchSpeaker{}

	frenchToEnglishAdapter := &FrenchToEnglishAdapter{fr: frenchSpeaker}

	frenchAdaptedExecutor := EnglishExecutor{englishSpeaker: frenchToEnglishAdapter}
	frenchAdaptedExecutor.Execute("Bonjour")

}
