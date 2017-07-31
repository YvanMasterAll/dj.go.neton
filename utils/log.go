package utils

import (
	"path/filepath"
)

type logi struct {
	file	string
}

func Info(m string){
	m = GetTime() + " " + "[Info] " + m
	if E != "" {
		Append(filepath.Join(E, L), m)
	}
}

func Debug(m string){
	m = GetTime() + " " + "[Debug] " +  m
	if E != "" {
		Append(filepath.Join(E, L), m)
	}
}

func Error(m string){
	m = GetTime() + " " + "[Error] " + m
	if E != "" {
		Append(filepath.Join(E, L), m)
	}
}