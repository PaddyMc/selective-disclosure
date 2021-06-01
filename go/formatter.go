package main

import "encoding/hex"

type lFormatter struct {
}

func (lf *lFormatter) Format(s []byte) string {
	return string(s)
}

type bFormatter struct {
}

func (bf *bFormatter) Format(s []byte) string {
	address := hex.EncodeToString(s)
	return address
}
