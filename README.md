# Auth API
This is an Auth API service in Go that uses JWT.

## Command to run locally
1. Installing all the required packages.
```go
go mod download
```
2. For running the Auth API service.
```go
go run .
```

## Curl commands
1. Sign Up
```bash
curl --location 'http://127.0.0.1:3000/api/assignment/auth/signup' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email":"email@gmail.com",
    "password":"password@123",
    "confirmPassword":"password@123"
}' 
```
2. Sign In
```bash
curl --location 'http://127.0.0.1:3000/api/assignment/auth/signin' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email":"email@gmail.com",
    "password":"password@123"
}'
```
3. Token Authorization
```bash
curl --location --request POST 'http://127.0.0.1:3000/api/assignment/auth/verify' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzM5ODk5MzgsInJhbmRvbU51bSI6ODQ0NjM2NjczMzAwLCJ1c2VyIjoidmlyYWpsMkBnbWFpbC5jb20ifQ.jcXbQptcFTaO7JE31uIRG-uiLAc_rQhNVLpDN4JXH0U'
```

4. Refresh Token
```bash
curl --location --request POST 'http://127.0.0.1:3000/api/assignment/auth/refreshToken' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzQwMDc2MzgsInJhbmRvbU51bSI6NTA2OTkzMDc0ODI4LCJ1c2VyIjoidmlyYWpsMkBnbWFpbC5jb21yZWZyZXNodG9rZW4ifQ.9ayc_TmDUYcethQbFEGwEG-QVHIBvWfmClYwubJ9cLU'
```

5. Revoke Token
```bash
curl --location --request POST 'http://127.0.0.1:3000/api/assignment/auth/revokeToken' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzM5ODk5MzgsInJhbmRvbU51bSI6ODQ0NjM2NjczMzAwLCJ1c2VyIjoidmlyYWpsMkBnbWFpbC5jb20ifQ.jcXbQptcFTaO7JE31uIRG-uiLAc_rQhNVLpDN4JXH0U'
```
