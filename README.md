# 🔐 GoPassword

**GoPassword** é um gerador de senhas simples e rápido desenvolvido em **Go (Golang)** com **Tailwind CSS** no front-end.  
O objetivo é permitir que o usuário crie senhas seguras, personalizadas e aleatórias de forma prática, direto no navegador.

---

## 🚀 Funcionalidades

- Geração de senhas seguras e aleatórias  
- Opção de incluir letras maiúsculas, minúsculas, números e símbolos  
- Definição do tamanho da senha  
- Interface simples e responsiva com Tailwind CSS  
- Backend leve e rápido utilizando Go (`net/http`)

---

## 🛠️ Tecnologias Utilizadas

- **Go (Golang)** — servidor e lógica principal  
- **HTML + Tailwind CSS** — interface web  
- **JavaScript** — interação dinâmica no front-end  

---

## 📦 Como Executar o Projeto

### 1. Clonar o repositório
```bash
git clone https://github.com/CaioHentz/GoPassword.git
cd GoPassword
```

### 2. Instalar dependências (caso tenha)
```bash
go mod tidy
```

### 3. Rodar o servidor
```bash
go run .
```

O servidor iniciará em:
👉 [http://localhost:8080](http://localhost:8080)

---

## 🧠 Estrutura do Projeto

```
GoPassword/
│
├── main.go                # Servidor principal em Go
├── go.mod                 # Configuração do módulo Go
│
├── static/
│   ├── css/
│   │   └── styles.css     # Estilos com Tailwind
│   └── js/
│       └── script.js      # Lógica de geração no front
│
└── templates/
    └── index.html         # Página principal
```

---

## 💡 Melhorias Futuras

- Histórico de senhas geradas  
- Opção de copiar senha com um clique  
- Modo escuro  
- API REST para integração externa  

---

## 🧑‍💻 Autor

**Caio Hentz**  
📎 [github.com/CaioHentz](https://github.com/CaioHentz)

---

## 📄 Licença

Este projeto está sob a licença MIT — veja o arquivo [LICENSE](LICENSE) para mais detalhes.
