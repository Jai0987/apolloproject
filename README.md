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

# Prereqs:
1. **Go**: Version 1.19 or higher.
2. **MongoDB**: A local or remote MongoDB instance (Make sure mongoDB is up and running; [MongoDB Installation Guide](https://www.mongodb.com/docs/manual/administration/install-community/#std-label-install-mdb-community-edition))

---

# Run the project

1. Clone the project to your local machine:
  ```bash
  git clone https://github.com/Jai0987/apolloproject.git
  cd apolloproject
  ```

2. Ensure proper rights to the shell scripts
  ```bash
  chmod +x run.sh
  chmod +x setup.sh
  ```

3. Now, run the script

```bash
./run.sh
```

It will automatically setup the environment for you if not installed already (setting up Go and MongoDB), will build the Go APIs and create the DB with the Vehicle Collection and then add some sample data too.

---

# Testing

For testing, Postman, Insomnia or any other API testing tool can be easily used.
For testing on Terminal:

1. Get All Vehicles:
```
bash
curl -X GET http://localhost:8080/vehicle
```

2. Get a Vehicle with a specific VIN
```bash
curl -X GET http://localhost:8080/vehicle/<VIN>
Eg: Sample VIN: 1HGCM82633A123456
```

3. Add a Vehicle
```bash
curl -X POST http://localhost:8080/vehicle \
-H "Content-Type: application/json" \
-d '{
    "vin": "1HGCM82633A123456",
    "manufacturer_name": "Honda",
    "description": "Compact Sedan",
    "horse_power": 150,
    "model_name": "Civic",
    "model_year": 2023,
    "purchase_price": 22000,
    "fuel_type": "Hybrid"
}'
```

4. Update a pre-existing vehicle record
```bash
curl -X PUT http://localhost:8080/vehicle/1HGBH41JXMN109186 \
-H "Content-Type: application/json" \
-d '{
    "manufacturer": "Toyota",
    "description": "Updated compact car",
    "horsePower": 130,
    "modelName": "Corolla",
    "model_year": 2020,
    "purchase_price": 20000,
    "fuel_type": "Petrol"
}'
```

5. Delete a Vehicle Record
```bash
curl -X DELETE http://localhost:8080/vehicle/3C6UR5FL2JG301234
```

---

## References and Support Material:
1. Go Official Documentation
2. gonic-gin documentation
3. MongoDB Documentation
4. Youtube
5. GenAI (Solely to support and fix Testing functionality)
