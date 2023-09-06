# Comments

### Run

```
docker-compose up --build
```

### Test

```
curl http://localhost:8010/api/v1/users -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluQGV4YW1wbGUuY29tIiwiZXhwIjo0ODQ3NTI1OTcwLCJpc19hZG1pbiI6dHJ1ZSwibmFtZSI6IlRoZSBBZG1pbmlzdHJhdG9yIiwib3JpZ19pYXQiOjE2OTM5MjU5NzB9.BXg9ucp8Lv1IE-18q93aqpRu7FqHQQjEkbtgsj6-i1k"
```

```
curl http://localhost:8010/api/v1/users -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVzZXJAZXhhbXBsZS5jb20iLCJleHAiOjQ4NDc1MjYwMDcsImlzX2FkbWluIjpmYWxzZSwibmFtZSI6IlRoZSBVc2VyIiwib3JpZ19pYXQiOjE2OTM5MjYwMDd9.UdHFhYOgOm3UOGDRJEn3jHTAndDgwejL0rsKMIXLOlc"
```

```
curl http://localhost:8010/api/v1/user/1 -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluQGV4YW1wbGUuY29tIiwiZXhwIjo0ODQ3NTI1OTcwLCJpc19hZG1pbiI6dHJ1ZSwibmFtZSI6IlRoZSBBZG1pbmlzdHJhdG9yIiwib3JpZ19pYXQiOjE2OTM5MjU5NzB9.BXg9ucp8Lv1IE-18q93aqpRu7FqHQQjEkbtgsj6-i1k"
```

```
curl http://localhost:8010/api/v1/user/1 -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVzZXJAZXhhbXBsZS5jb20iLCJleHAiOjQ4NDc1MjYwMDcsImlzX2FkbWluIjpmYWxzZSwibmFtZSI6IlRoZSBVc2VyIiwib3JpZ19pYXQiOjE2OTM5MjYwMDd9.UdHFhYOgOm3UOGDRJEn3jHTAndDgwejL0rsKMIXLOlc"
```

```
curl http://localhost:8010/api/v1/comments -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVzZXJAZXhhbXBsZS5jb20iLCJleHAiOjQ4NDc1MjYwMDcsImlzX2FkbWluIjpmYWxzZSwibmFtZSI6IlRoZSBVc2VyIiwib3JpZ19pYXQiOjE2OTM5MjYwMDd9.UdHFhYOgOm3UOGDRJEn3jHTAndDgwejL0rsKMIXLOlc"
```

```
curl http://localhost:8010/api/v1/comment/3 -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVzZXJAZXhhbXBsZS5jb20iLCJleHAiOjQ4NDc1MjYwMDcsImlzX2FkbWluIjpmYWxzZSwibmFtZSI6IlRoZSBVc2VyIiwib3JpZ19pYXQiOjE2OTM5MjYwMDd9.UdHFhYOgOm3UOGDRJEn3jHTAndDgwejL0rsKMIXLOlc"
```

```
curl http://localhost:8010/api/v1/comment/4 -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVzZXJAZXhhbXBsZS5jb20iLCJleHAiOjQ4NDc1MjYwMDcsImlzX2FkbWluIjpmYWxzZSwibmFtZSI6IlRoZSBVc2VyIiwib3JpZ19pYXQiOjE2OTM5MjYwMDd9.UdHFhYOgOm3UOGDRJEn3jHTAndDgwejL0rsKMIXLOlc"
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
wrk http://localhost:8010/api/v1/users -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluQGV4YW1wbGUuY29tIiwiZXhwIjo0ODQ3NTI1OTcwLCJpc19hZG1pbiI6dHJ1ZSwibmFtZSI6IlRoZSBBZG1pbmlzdHJhdG9yIiwib3JpZ19pYXQiOjE2OTM5MjU5NzB9.BXg9ucp8Lv1IE-18q93aqpRu7FqHQQjEkbtgsj6-i1k"
```
