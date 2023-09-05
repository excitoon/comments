# Comments

### Run

```
docker-compose up --build
```

### Test

```
curl http://localhost:8010/api/v1/users -H "Authorization: Bearer TOKEN"
```

```
curl http://localhost:8010/api/v1/comments -H "Authorization: Bearer TOKEN"
```

#### Log in as admin

```
curl http://localhost:8010/api/v1/auth -X POST -H "Content-Type: application/x-www-form-urlencoded" -d "username=The%20Administrator&password=admin"
```

#### Log in as user

```
curl http://localhost:8010/api/v1/auth -X POST -H "Content-Type: application/x-www-form-urlencoded" -d "username=The%20User&password=user"
```

### Performance

```
wrk http://localhost:8010/api/v1/users -H "Authorization: Bearer TOKEN"
```
