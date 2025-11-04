package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"goPassword/internal/generator"
)

func main() {
	mux := http.NewServeMux()

	// rota principal (HTML)
	mux.HandleFunc("/", handleIndex)
	// rota de geração (POST)
	mux.HandleFunc("/generate", handleGenerate)

	log.Println("🚀 Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	tmplPath := filepath.Join("templates", "index.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Erro ao carregar template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func handleGenerate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	lengthStr := r.FormValue("length")
	length, err := strconv.Atoi(lengthStr)
	if err != nil || length <= 0 {
		length = 12 // padrão
	}

	useUpper := r.FormValue("upper") == "on"
	useLower := r.FormValue("lower") == "on"
	useDigits := r.FormValue("digits") == "on"
	useSymbols := r.FormValue("symbols") == "on"

	password := generator.GeneratePassword(length, useUpper, useLower, useDigits, useSymbols)

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(password))
}
