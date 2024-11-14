Pet project to play with [monstache](https://github.com/rwynn/monstache)

# First run

```sh
docker compose up -d

# insert on mongo
db.users.insertOne({name: "Marcelo"})

# Check on elastic
curl "http://localhost:9200/app.users/_search?pretty=true"

```

# References

- https://github.com/rwynn/monstache
  - https://github.com/rwynn/monstache-showcase/blob/master/docker-compose.sc.yml
