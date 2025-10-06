# 💻 Study Application

## 🧾 Description
A simple **banking application** that allows you to **authenticate using an API key**, view your **account balance**, and **make transfers** securely through RESTful endpoints.

---

## 🎨 Design System
![Bank App UI Preview](https://github.com/user-attachments/assets/92e7ac03-f4e2-4572-8eb7-98316d7b04e5)

---

## 🧠 Technologies Used
- **Golang (Gin Framework)** – Fast and minimal web framework for Go  
- **Nginx** – Reverse proxy and load balancer  
- **PostgreSQL** – Relational database for structured data  
- **Apache Kafka** – Event streaming and asynchronous message handling  

---

## 🚀 Features
- API key–based authentication  
- Retrieve account balance  
- Perform transfers between accounts  
- Kafka integration for event-driven processing  
- Reverse proxy routing via Nginx  

---

## 🛠️ Setup (Quick Start)
```bash
# Clone the repository
git clone https://github.com/yourusername/study-application.git
cd study-application

# Build and run with Docker Compose
docker-compose up --build

```

### 📘 OpenAPI Example

```yaml
openapi: 3.0.0
info:
  title: Bank Go API
  version: 1.0.0
  description: API endpoints for the Bank Go collection.
servers:
  - url: http://localhost:8080
paths:
  /user/balance_no_cache:
    get:
      summary: Get Balance (No Cache)
      description: Returns the balance for a user by CPF, without using cache.
      parameters:
        - in: query
          name: cpf
          schema:
            type: string
          required: true
          description: CPF of the user whose balance is being requested.
      responses:
        '200':
          description: Successful response with balance information
          content:
            application/json:
              schema:
                type: object
                properties:
                  error:
                    type: string
                    nullable: true
                  result:
                    type: string
                    description: Balance amount
              example:
                error: null
                result: "10000"
        '400':
          description: Bad request
        '500':
          description: Internal server error
  # The following endpoint is included as a placeholder since details are not provided
  # Please update with actual details if available
  /user/balance:
    get:
      summary: Get Balance
      description: Returns the balance for a user (details not provided in context).
      responses:
        '200':
          description: Successful response
        '400':
          description: Bad request
        '500':
          description: Internal server error


