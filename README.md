# 🧪 Projeto - Sistema de Estoque e Faturamento

Este projeto é composto por dois serviços backend (`stock-service` e `billing-service`) e um frontend. Abaixo estão as instruções para rodar cada parte do sistema.

---

## ⚙️ Pré-requisitos

- [Go](https://golang.org/) instalado
- [Node.js](https://nodejs.org/) e npm ou yarn instalados
- [PostgreSQL](https://www.postgresql.org/) em execução
- Bancos de dados já criados:
  - `stock_service`
  - `billing_service`

> ⚠️ **Atenção:** Os bancos de dados **não são criados automaticamente**, portanto você precisa criá-los manualmente no PostgreSQL antes de iniciar os serviços.

---

## 🚀 Como rodar o projeto

### 📦 Backend

#### 1. Iniciando o Stock Service

```bash
cd backend/stock-service
go run main.go
```

---

#### 2. Iniciando o Billing Service

```bash
cd backend/billing-service
go run main.go
```

---

### 💻 Frontend

```bash
cd frontend
```

A partir daqui, você pode instalar as dependências e iniciar o frontend com:

```bash
npm install     # ou yarn install
npm start       # ou yarn start
```

---

## 📌 Observações

- Certifique-se de que os serviços backend estão rodando antes de iniciar o frontend.
- Verifique as portas utilizadas por cada serviço e se estão livres no seu sistema.

---

Feito com 💻 por MateusQ7
