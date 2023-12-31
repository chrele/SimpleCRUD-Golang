# CRUD WebServer using RESTApi by GoLang and PostgreSQL

This program create a RESTApi CRUD application using GoLang and establish connection to PostgreSQL.

To use:
1. ```git clone``` this repository or download it as zip. (Recommended) Fork this repository, then clone it from your own github.
2. Download [GO](https://go.dev/doc/install) and install it on your PC.
3. Open the folder on your IDE. (Recommended) Download [Visual Studio Code](https://code.visualstudio.com/).
4. Run ```go mod download``` to automatically install any necessary package from GoLang.
5. Create ```app.env``` file and set DB connection:
   ```
    postgres_host = "localhost"
    postgres_user = "postgres"
    postgres_password = "postgres"
    postgres_db = "test"
    postgres_port = "5432"
   ```
   > This part is still buggy as the current version of ```viper.Unmarshal()``` will not convert the config to struct without ```viper.BindEnv()``` and will just return empty struct. If you want to change the DB connection, change it manually on ```./config/database.go```.
6. To check any Postgre details, you can go to pgAdmin Xv.Y > Servers > Postgre SQL 1X. Right click and select Properties. Go to Connection tab, and save all details there.
7. Run ```go run .\main.go``` and if everything goes well, you will see 
    ```
    Run service . . .Connection to DB is Successful.
    ┌───────────────────────────────────────────────────┐ 
    │                   Fiber v2.48.0                   │
    │               http://127.0.0.1:8000               │
    │       (bound on host 0.0.0.0 and port 8000)       │
    │                                                   │
    │ Handlers ............. 9  Processes ........... 1 │
    │ Prefork ....... Disabled  PID ............. 16260 │
    └───────────────────────────────────────────────────┘ 
    ```
8. Check the API with Postman or any tools. It is ready to be used for basic CRUD.
9. I hope this repository can help you understand how to use GoLang to make CRUD with RESTApi to PostgreSQL. Please let me know if you have any question by raising issues~
