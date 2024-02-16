package main

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
)

var PDFTemplate = template.Must(template.ParseFiles("main.tmpl.tex"))

func main() {
	out := bytes.NewBuffer([]byte{})
	PDFTemplate.Execute(out, nil)
	pdf, err := GeneratePDFFromLatex(out.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(pdf)
}

func GeneratePDFFromLatex(src []byte) ([]byte, error) {
	os.MkdirAll("./tmp/", 0o775)
	tmpDir, err := os.MkdirTemp("./tmp/", "")
	if err != nil {
		return nil, err
	}
	errout := bytes.NewBuffer([]byte{})
	cmd := exec.Command("xelatex", "--jobname=main")
	cmd.Dir = tmpDir
	cmd.Stdin = bytes.NewBuffer(src)
	cmd.Stdout = errout
	err = cmd.Run()
	if err != nil {
		return nil, err
	}
	target := filepath.Join(tmpDir, "main.pdf")
	return os.ReadFile(target)
}
