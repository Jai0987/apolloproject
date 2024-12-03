# Apollo Coding Challenge
I designed this project solely to complete the coding challenge sent to me. It provides a RESTful API to manage vehicle data collection, supporting Create, Read, Update, and Delete (CRUD) operations. This project is designed to be easily extensible and deployable.  The application also includes automation scripts to set up the environment, populate the database with sample data, and manage the project's lifecycle.

## **Tech Stack**
- **Programming Language**: Go (Golang)
- **Database**: MongoDB
- **API Framework**: Gin
- **Build Tool**: Makefile

## Endpoints:

| HTTP Method | Endpoint                     | Description                            |
|-------------|------------------------------|----------------------------------------|
| GET         | `/vehicle`                   | Fetch all vehicles.                   |
| GET         | `/vehicle/:vin`              | Fetch a specific vehicle by VIN.      |
| POST        | `/vehicle`                   | Add a new vehicle.                    |
| PUT         | `/vehicle/:vin`              | Update a specific vehicle by VIN.     |
| DELETE      | `/vehicle/:vin`              | Delete a specific vehicle by VIN.     |
