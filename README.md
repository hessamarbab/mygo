## simple sample project in Go lang

## Description:
This Go application reads a JSON file containing data for too many users and concurrently inserts this data into a PostgreSQL relational database. The application is designed to limit concurrent processes to 10 to manage resource utilization.

The user data is structured into two tables: 'users' and 'addresses'. A one-to-many relationship is established between these tables, allowing for multiple addresses to be associated with a single user.

An HTTP API built using the Echo framework enables fetching detailed information about a specific user by providing their unique ID. The API response includes both the user's basic details and their associated addresses.

The user data model consists of four fields: first name, last name, address, and a unique ID.

To facilitate data transfer between the reading and writing components, Go channels are employed. The GORM ORM (Object-Relational Mapper) is utilized for interacting with the PostgreSQL database, providing a convenient way to map Go structs to database tables.

## run by:
```
    docker compose up  
```

## Api:
#### Url: /user
#### method: GET
#### headers:
```
Accept: application/json
Content-Type: application/json
```
#### body:
```
{
    "id":"35b75a6b-0183-43e1-bdd0-3cd062dc6935"
}
```

  
