# bank_transaction_application
A simple internal financial transaction system built with Golang and PostgreSQL. It provides RESTful APIs to create accounts, check balances, and transfer funds between accounts.

🛠 Features
Create accounts with initial balances
Query account balances
Perform internal fund transfers between accounts

📦 Tech Stack
Language: Go (Golang)
Database: PostgreSQL
Router: Gorilla Mux


### Database Setup

sql
CREATE TABLE accounts (
    account_id BIGINT PRIMARY KEY,
    balance NUMERIC(20, 5) NOT NULL
);

CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    source_account_id BIGINT NOT NULL,
    destination_account_id BIGINT NOT NULL,
    amount NUMERIC(20, 5) NOT NULL,
    created_at TIMESTAMP DEFAULT now()
);


### Run the App

1. Clone repo
2. Update DB credentials in main.go
3. Run:

bash
go mod tidy
go run main.go


### API Endpoints

- POST /accounts – create account
- GET /accounts/{id} – get account balance
- POST /transactions – transfer funds

## ✅ Assumptions

- No authentication needed
- Currency is the same across accounts

## 📄 License

This project is licensed under the [MIT License](LICENSE).
