# Vue Golang App

This project has been developed for Kartaca "Çekirdekten Yetişenler" program. The project includes Vuejs, Golang and Sqlite technologies. It consists of a login page and a simple dashboard. All users can share posts and see previously shared posts on the dashboard page.

## Compiling and Running The Project

All development and testing was done on Ubuntu 18.04.4 LTS.

### Build and Run with docker-compose

There is a docker-compose.yaml file in the project main directory. Therefore, you can run the application by following the steps below.

-----------------

1. Docker and docker-compose must be installed.

-----------------

2. Clone the project and go to project home directory.
```
cd your/host/dir/to/vue-golang-app
```
-----------------

3. Build and run with docker-compose.
```
sudo docker-compose up -d --build
```
-----------------

4. Check your containers. There should be 2 containers for server and client.
```
sudo docker container ls
```

-----------------

5. If everything is fine, go to http://localhost:8080 from your browser. Login screen needs to be opened.

You can use the registered users below:

```
1-Kartaca
Username: kartaca
Password: kartaca

2-Enes Karali
Username: eneskarali
Password: 123456

3-Ahmet Ak
Username: ahmetak
Password: 123456

4-Mehmet Sarı
Username: mehmetsarı
Password: 123456

```

-----------------

6. To stop
```
sudo docker-compose stop
```
-----------------

