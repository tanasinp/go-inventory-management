# Inventory Management System

![Untitled (1)2](https://github.com/user-attachments/assets/685674d8-128a-433f-9e87-754f7bd42ae3)

## Description
I've developed an Inventory Management System REST API. It helps manage products, suppliers, categories, and users. Developed using the Go programming language, it leverages the Fiber framework, GORM for ORM, PostgreSQL for data storage, and Docker for containerization.

## Features
-**Product Management :** Create, update, delete, and retrieve products. Each product can be associated with a category and a supplier.
-**Category Managemen :** Create, update, delete, and retrieve categories. Each category can have multiple products. Organize products into categories.
-**Supplier Management :** Create, update, delete, and retrieve suppliers. Each supplier can supply multiple products. Manage supplier details like name and contact information.
-**User Authentication :** Register and login users with JWT-based authentication.
-**Authorization :** Protect routes with middleware to ensure only authenticated users can access them.

## Installation
1. Clone the repository
2. Create `.env` and Setup `.env` file :
   ```env
   DB_HOST=your_db_host
   DB_PORT=your_db_port
   DB_USER=your_db_user
   DB_PASSWORD=your_db_password
   DB_NAME=your_db_name
   JWT_SECRETKEY=your_jwt_secret_key
   ```
3. Create a `docker-compose.yml` file :
    ```
    version: '3.8'

    services:
    postgres:
        image: postgres:latest
        container_name: postgres
        environment:
        POSTGRES_DB: <Maintenance database>
        POSTGRES_USER: <Username>
        POSTGRES_PASSWORD: <Password>
        volumes:
        - postgres_data:/var/lib/postgresql/data
        ports:
        - "<external : internal>"
        restart: unless-stopped

    pgadmin:
        image: dpage/pgadmin4:latest
        container_name: pgadmin
        environment:
        PGADMIN_DEFAULT_EMAIL: <Email to login>
        PGADMIN_DEFAULT_PASSWORD: <Password>
        ports:
        - "<external : internal>"
        depends_on:
        - postgres
        restart: unless-stopped

    volumes:
    postgres_data:
    ```
4. Run `docker-compose up -d` to start the services.
5. Access the API at `http://localhost:8000` and pgAdmin at `http://localhost:external_port`.