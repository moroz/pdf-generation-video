package main

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
)

var PDFTemplate = template.Must(template.ParseFiles("main.tmpl.tex"))

const listenOn = ":3000"

func main() {
	http.HandleFunc("/", handleGetPDFRequest)
	log.Printf("Listening on %s", listenOn)
	log.Fatal(http.ListenAndServe(listenOn, nil))
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

func handleGetPDFRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
	}

	out := bytes.NewBuffer([]byte{})
	PDFTemplate.Execute(out, nil)
	pdf, err := GeneratePDFFromLatex(out.Bytes())

	if err != nil {
		log.Print(err)
		w.Header().Add("Content-Type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("Content-Type", "application/pdf")
	w.WriteHeader(http.StatusOK)
	w.Write(pdf)
}
