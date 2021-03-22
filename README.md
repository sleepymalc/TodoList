# README

## Abstract

#### 	This project is based on following elements:

#### 		1. Golang

#### 		2. Gin Framework

#### 		3. JWT(Json Web Token)

#### 		4. Restful API

#### 		5. Mongodb 

#### 	To build a simple todo list project, while using Postman to test all the API functionality, and finally build on Docker.

## Features

#### 	1. users/singup : Post method

#### 		Sample input:

```json
{
		"user_id": "admin",
		"password": "88888888"
}
```

#### 		Sample output:

```json
{
		"InsertedID": "60581eb9cc2216b508d7477d"
}
```

#### 	2. users/login : Post method

#### 		Sample input:

```json
{
		"user_id": "admin",
		"password": "88888888"
}
```

#### 		Sample output:

```json
{
"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyX2lkIjoiYWRtaW4iLCJleHAiOjE2MTY0NzQxNjl9.aAlh_LuxqLfWWSzCd3uA3C2RoBOTnC3HSqPaAvzYIkE"
}
```

#### 	3. users/todo_list : Post method

#### 		Sample input:

```json
{
  	"Authorization": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyX2lkIjoiYWRtaW4iLCJleHAiOjE2MTY0NzQxNjl9.aAlh_LuxqLfWWSzCd3uA3C2RoBOTnC3HSqPaAvzYIkE",
		"todo_list": "do homework",
		"todo_list": "write diary"
}
```

#### 		Sample output:

```json
{
    "MatchedCount": 1,
    "ModifiedCount": 1,
    "UpsertedCount": 0,
    "UpsertedID": null
}
```

#### 	4. User/todo_list : Get method

#### 		Sample input:

```json
{
		"Authorization" = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyX2lkIjoiYWRtaW4iLCJleHAiOjE2MTYxNzMyODZ9.8nri7jkBQPr-0ygZe5a8W2OB1K8TdmNJtc0W-XAqVkk
}
```

#### 		Sample output:

```json
[
    "do homework",
    "write diary"
]
```

#### 	5. User/todo_list : Delete method

#### 		Sample input:

```json
{
		"Authorization" = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyX2lkIjoiYWRtaW4iLCJleHAiOjE2MTYxNzMyODZ9.8nri7jkBQPr-0ygZe5a8W2OB1K8TdmNJtc0W-XAqVkk",
		"delete_element" = "do homework"
}
```

#### 	Sample output:

```json
{
    "MatchedCount": 1,
    "ModifiedCount": 1,
    "UpsertedCount": 0,
    "UpsertedID": null
}
```





#### 