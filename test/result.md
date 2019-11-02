FORMAT: 1A
HOST: 127.0.0.1

# Minimum-Template
this is the minimum backend powered by minimum

## Ping [/ping]


 + GET: result


## Login [/v1/login]


 + POST: Login


### Login [POST]

+ Request admin login for test

    + Headers

            Content-Type: application/json

    + Body

            {
                "id": 1,
                "user_name": "",
                "phone": "",
                "password": "admin"
            }


+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0,
                "identity": null,
                "phone": "1234567891011",
                "id": 1,
                "nick_name": "admin_context",
                "name": "admin_context",
                "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzI3MDE1MTYsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjk3OTA2LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.Yuyew7UJIYTpb586sn9lMQhtegV_u1AouZlH8LKLOsA",
                "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzMzMDI3MTYsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjk3OTA2LCJJc1JlZnJlc2hUb2tlbiI6dHJ1ZSwiUmVmcmVzaFRhcmdldCI6eyJleHAiOjE1NzI3MDE1MTYsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjk3OTA2LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19LCJDdXN0b21GaWVsZCI6eyJVSUQiOjF9fQ.Ap2aFH1PCn2UnCh8BA7Rpci7SqyQn8EsX53V9H2ueqU"
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzI3MDE1MTYsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjk3OTA2LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.Yuyew7UJIYTpb586sn9lMQhtegV_u1AouZlH8LKLOsA
            Content-Type: application/json

    + Body

            {
                "id": 2,
                "user_name": "",
                "phone": "",
                "password": "11122222"
            }


+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0,
                "identity": null,
                "phone": "10086111",
                "id": 2,
                "nick_name": "tan chan",
                "name": "chan tan",
                "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzI3MDE1MTYsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjk3OTA2LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19.SHCiXZyY6LIPYpzQ-gpuKkJKEbETzLBO0yqDS1GWaOI",
                "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzMzMDI3MTYsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjk3OTA2LCJJc1JlZnJlc2hUb2tlbiI6dHJ1ZSwiUmVmcmVzaFRhcmdldCI6eyJleHAiOjE1NzI3MDE1MTYsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjk3OTA2LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19LCJDdXN0b21GaWVsZCI6eyJVSUQiOjJ9fQ.H5EuRENKq0X830r9k-ffCQWAGSumD-kEKrrIdO9B-FE"
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzI3MDE1MTYsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjk3OTA2LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.Yuyew7UJIYTpb586sn9lMQhtegV_u1AouZlH8LKLOsA
            Content-Type: application/json

    + Body

            {
                "id": 0,
                "user_name": "chan tan",
                "phone": "",
                "password": "11122222"
            }


+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0,
                "identity": null,
                "phone": "10086111",
                "id": 2,
                "nick_name": "tan chan",
                "name": "chan tan",
                "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzI3MDE1MTYsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjk3OTA2LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19.SHCiXZyY6LIPYpzQ-gpuKkJKEbETzLBO0yqDS1GWaOI",
                "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzMzMDI3MTYsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjk3OTA2LCJJc1JlZnJlc2hUb2tlbiI6dHJ1ZSwiUmVmcmVzaFRhcmdldCI6eyJleHAiOjE1NzI3MDE1MTYsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjk3OTA2LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19LCJDdXN0b21GaWVsZCI6eyJVSUQiOjJ9fQ.H5EuRENKq0X830r9k-ffCQWAGSumD-kEKrrIdO9B-FE"
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzI3MDE1MTYsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjk3OTA2LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.Yuyew7UJIYTpb586sn9lMQhtegV_u1AouZlH8LKLOsA
            Content-Type: application/json

    + Body

            {
                "id": 0,
                "user_name": "",
                "phone": "10086111",
                "password": "11122222"
            }


+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0,
                "identity": null,
                "phone": "10086111",
                "id": 2,
                "nick_name": "tan chan",
                "name": "chan tan",
                "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzI3MDE1MTYsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjk3OTA2LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19.SHCiXZyY6LIPYpzQ-gpuKkJKEbETzLBO0yqDS1GWaOI",
                "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzMzMDI3MTYsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjk3OTA2LCJJc1JlZnJlc2hUb2tlbiI6dHJ1ZSwiUmVmcmVzaFRhcmdldCI6eyJleHAiOjE1NzI3MDE1MTYsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjk3OTA2LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19LCJDdXN0b21GaWVsZCI6eyJVSUQiOjJ9fQ.H5EuRENKq0X830r9k-ffCQWAGSumD-kEKrrIdO9B-FE"
            }


## Register [/v1/user]


 + POST: Register


### Register [POST]

+ Request 

    + Headers

            Content-Type: application/json

    + Body

            {
                "name": "admin_context",
                "password": "admin",
                "nick_name": "admin_context",
                "phone": "1234567891011",
                "register_city": "Qing Dao S.D."
            }


+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0,
                "id": 1
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzI3MDE1MTYsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjk3OTA2LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.Yuyew7UJIYTpb586sn9lMQhtegV_u1AouZlH8LKLOsA
            Content-Type: application/json

    + Body

            {
                "name": "chan tan",
                "password": "11122222",
                "nick_name": "tan chan",
                "phone": "10086111",
                "register_city": "tan arch"
            }


+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0,
                "id": 2
            }


## List [/v1/user-list]


 + GET: List User


### List [GET]

+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzI3MDE1MTYsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjk3OTA2LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.Yuyew7UJIYTpb586sn9lMQhtegV_u1AouZlH8LKLOsA
            Content-Type: text/plain

    + Body

            

+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0,
                "users": [
                    {
                        "ID": 1,
                        "created_at": "2019-11-02T20:31:56.3320422+08:00",
                        "updated_at": "2019-11-02T20:31:56.3320422+08:00",
                        "last_login": "2019-11-02T12:31:56Z",
                        "NickName": "admin_context",
                        "Name": "admin_context",
                        "Password": "$2a$10$i9aTebzwPJd9MfLspBSKXenzG34qdXYmbO9LJHgp4x5dMZZzl8N0u",
                        "Phone": "1234567891011",
                        "RegisterCity": "Qing Dao S.D."
                    }
                ]
            }


## Get [/v1/user/:id]


 + GET: Get User
 + PUT: Put User
 + DELETE: Delete User


### Get [GET]

+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzI3MDE1MTYsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjk3OTA2LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.Yuyew7UJIYTpb586sn9lMQhtegV_u1AouZlH8LKLOsA
            Content-Type: text/plain

    + Body

            

+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0,
                "user": {
                    "ID": 1,
                    "created_at": "2019-11-02T20:31:56.3320422+08:00",
                    "updated_at": "2019-11-02T20:31:56.3320422+08:00",
                    "last_login": "2019-11-02T12:31:56Z",
                    "NickName": "admin_context",
                    "Name": "admin_context",
                    "Password": "$2a$10$i9aTebzwPJd9MfLspBSKXenzG34qdXYmbO9LJHgp4x5dMZZzl8N0u",
                    "Phone": "1234567891011",
                    "RegisterCity": "Qing Dao S.D."
                }
            }


## ChangePassword [/v1/user/:id/password]


 + POST: change password of user

