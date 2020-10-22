## GO-API-Docker

Setup with docker
* docker-compose build
* docker-compose --env-file config.env up
* access mysql container with
    * `docker exec it <containerid> bash -l`
    * `mysql -u <username> -p`
    * insert <password>
    * `use <databasename>`
    * ```
      insert into products (name, price, quantity, state) values
      ('Hammer', 25.34, 50, 'available'),
      ('Screwdriver', 12.30, 25, 'available'),
      ('Saw', 100.23, 5, 'available'),
      ('Drill', 79.99, 10, 'available'),
      ('Nail', 1.99, 100, 'available'),
      ('Ladder', 67.88, 5, 'available'),
      ('Rail', 78.45, 34, 'available'),
      ('Pi', 15.00, 20, 'available'),
      ('Wires', 2.00, 0, 'unavailable');
      ``` 
    * verify products inserted into db `select * from products`
      
* use postman or curl `http://localhost:8080/v1/products`
* close containers`docker-compose down`
---

Setup local
* Update the config.env app `db-host` to `localhost` and `db-user` to `root` 
* Ensure your local mysql server started and running at port 3306
* make setup
* make build
* make test
* make run
* insert data into database using insert command above
* test endpoint

