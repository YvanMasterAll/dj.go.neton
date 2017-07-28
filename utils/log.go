package utils

import "path/filepath"

const (
	F = "neton.log"
)

type log struct {
	file	string
}

func Info(m string){
	m = GetTime() + " " + "[Info] " + m
	Append(filepath.Join(GetCurrentPath(), F), m)
}

func Debug(m string){
	m = GetTime() + " " + "[Debug] " +  m
	Append(filepath.Join(GetCurrentPath(), F), m)
}

func Error(m string){
	m = GetTime() + " " + "[Error] " + m
	Append(filepath.Join(GetCurrentPath(), F), m)
}