## Interfaces and Dependency injection lab

## Objectives
- The initial code is broken, see if you fix it with minimal changes.
- Associate the product service behaviour with the struct.
- Product service is dependent on Repository, see if you can find the abstractions and extract them out as an interface
- Make your product service depend on abstraction rather than concrete implementation.
- Wire application from the main function

## Running mysql db container with default tables and data

```sh
docker run --rm -p 3306:3306 --name customers-db -e MYSQL_ROOT_PASSWORD=thecodecamp thecodecamp/customers-db
```