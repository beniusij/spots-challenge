# Spots Challenge

This projects provides solution to the challenges described below.

## Challenge 1

Write a Query (or as PL/SQL function) that includes the following points below.

- Return spots which have a domain with a count greater than 1.
- Change the website field, so it only contains the domain.
  - Example:https://domain.com/index.php→domain.com
- Count how many spots have the same domain.
- Return 3 columns : spot name, domain and count number for domain.

## Challenge 2

Create an endpoint which returns spots in a circle or square area.
Although the task is short, clean code and a good project structure is still very important. This task must be completed in Golang.

1. Endpoint should receive 4 parameters
  - Latitude
  - Longitude
  - Radius(in meters)
  - Type(circle or square)
2. Find all spots in the table (spots.sql) using the received parameters.
3. Order results by distance.
  - If distance between two spots is smaller than 50m, then order by rating.
4. Endpoint should return an array of objects containing all fields in the data set.

## Getting Started

To test solutions to the first challenge, import each query and function into database containing "MY_TABLE" table 
and proceed executing each query or function.

To test solution to the second challenge, perform following steps:

1. After downloading the project, `cd` into the project's directory
2. `make start-db`
3. Start microservice with `go run main.go`

**Note:** You might need to wait ~10 seconds after container startup for it to finish importing data.

To test the endpoint, run following command in the terminal:
```
curl "localhost:3000/spots?longitude=-8.473656&latitude=51.899216&radius=100&type=circle"
```
**Disclaimer**: Implementation for querying spots in square area is incomplete, thus results may be inaccurate or incomplete. 

To run tests, follow these steps inside project's directory:

1. `make test-db`
2. Run `make test`
   
**Note:** You might need to wait ~10 seconds after container startup for it to finish importing data.