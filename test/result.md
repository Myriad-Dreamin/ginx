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
                "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzI3MDE5ODIsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjk4MzcyLCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.SQvbhD7o2D7RZ7ib51xyDiWythyOE0nYOtbObdFv6tY",
                "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzMzMDMxODIsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjk4MzcyLCJJc1JlZnJlc2hUb2tlbiI6dHJ1ZSwiUmVmcmVzaFRhcmdldCI6eyJleHAiOjE1NzI3MDE5ODIsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjk4MzcyLCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19LCJDdXN0b21GaWVsZCI6eyJVSUQiOjF9fQ.OdkdkD6APOFd6yjZQQ4zW9HMuX1i1THq39p77Bho0Kg"
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzI3MDE5ODIsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjk4MzcyLCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.SQvbhD7o2D7RZ7ib51xyDiWythyOE0nYOtbObdFv6tY
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
                "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzI3MDE5ODIsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjk4MzcyLCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19.xz6xlVy-3-Z2_QYXCCrtA3pwywdrG1lTJtu9Uh_W5Kw",
                "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzMzMDMxODIsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjk4MzcyLCJJc1JlZnJlc2hUb2tlbiI6dHJ1ZSwiUmVmcmVzaFRhcmdldCI6eyJleHAiOjE1NzI3MDE5ODIsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjk4MzcyLCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19LCJDdXN0b21GaWVsZCI6eyJVSUQiOjJ9fQ.lAUkAgudMrb-SQWoUMTrP2MqcKnlnI-2DkmrFrEcsIM"
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzI3MDE5ODIsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjk4MzcyLCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.SQvbhD7o2D7RZ7ib51xyDiWythyOE0nYOtbObdFv6tY
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
                "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzI3MDE5ODIsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjk4MzcyLCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19.xz6xlVy-3-Z2_QYXCCrtA3pwywdrG1lTJtu9Uh_W5Kw",
                "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzMzMDMxODIsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjk4MzcyLCJJc1JlZnJlc2hUb2tlbiI6dHJ1ZSwiUmVmcmVzaFRhcmdldCI6eyJleHAiOjE1NzI3MDE5ODIsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjk4MzcyLCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19LCJDdXN0b21GaWVsZCI6eyJVSUQiOjJ9fQ.lAUkAgudMrb-SQWoUMTrP2MqcKnlnI-2DkmrFrEcsIM"
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzI3MDE5ODIsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjk4MzcyLCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.SQvbhD7o2D7RZ7ib51xyDiWythyOE0nYOtbObdFv6tY
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
                "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzI3MDE5ODIsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjk4MzcyLCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19.xz6xlVy-3-Z2_QYXCCrtA3pwywdrG1lTJtu9Uh_W5Kw",
                "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzMzMDMxODIsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjk4MzcyLCJJc1JlZnJlc2hUb2tlbiI6dHJ1ZSwiUmVmcmVzaFRhcmdldCI6eyJleHAiOjE1NzI3MDE5ODIsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjk4MzcyLCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19LCJDdXN0b21GaWVsZCI6eyJVSUQiOjJ9fQ.lAUkAgudMrb-SQWoUMTrP2MqcKnlnI-2DkmrFrEcsIM"
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzI3MDE5ODIsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjk4MzcyLCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.SQvbhD7o2D7RZ7ib51xyDiWythyOE0nYOtbObdFv6tY
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzI3MDE5ODIsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjk4MzcyLCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.SQvbhD7o2D7RZ7ib51xyDiWythyOE0nYOtbObdFv6tY
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
                        "created_at": "2019-11-02T20:39:42.545343+08:00",
                        "updated_at": "2019-11-02T20:39:42.545343+08:00",
                        "last_login": "2019-11-02T12:39:42Z",
                        "NickName": "admin_context",
                        "Name": "admin_context",
                        "Password": "$2a$10$grfHhra3G/HXrTui059FmuMV9rgl.JfTEhg8sJaDtLM6YoPUNozFu",
                        "Phone": "1234567891011",
                        "RegisterCity": "Qing Dao S.D."
                    }
                ]
            }


## Delete [/v1/user/:id]


 + DELETE: Delete User
 + GET: Get User
 + PUT: Put User


### Get [GET]

+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzI3MDE5ODIsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjk4MzcyLCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.SQvbhD7o2D7RZ7ib51xyDiWythyOE0nYOtbObdFv6tY
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
                    "created_at": "2019-11-02T20:39:42.545343+08:00",
                    "updated_at": "2019-11-02T20:39:42.545343+08:00",
                    "last_login": "2019-11-02T12:39:42Z",
                    "NickName": "admin_context",
                    "Name": "admin_context",
                    "Password": "$2a$10$grfHhra3G/HXrTui059FmuMV9rgl.JfTEhg8sJaDtLM6YoPUNozFu",
                    "Phone": "1234567891011",
                    "RegisterCity": "Qing Dao S.D."
                }
            }


## ChangePassword [/v1/user/:id/password]


 + POST: change password of user

