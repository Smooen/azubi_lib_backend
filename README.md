# Azubiprojekt Backend

Implemented in Go.

## How to run

We are using [air](https://github.com/air-verse/air) to enable hot-reloading for an easier time in development.
Simply run 

```
air
```

in the project root.


## Technologies used

This project uses sqlite as database (for now - might change to MS SQL Server)

[Database driver](https://github.com/mattn/go-sqlite3)

[ORM](https://gorm.io/docs/)

Web framework for building web apps in Go.

[Echo](https://echo.labstack.com/docs/)

All of these technologies are very popular, modern and have excellent documentation.

## TODO:

- Models feddich (Nurnoch User, Favorites)
- Wie geht der Datenbank zugriff (mehr oder weniger geschafft)
- Basic Queries 
    -> GET Books, Get Book mit ID/Name etc.
- CRUD Operations mit ECHO
    - GET
    - Save
    - Update? (weil availability)
    - (Delete)
- Authentication ??
