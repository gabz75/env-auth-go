# env-auth

JSON RESTFul auth API with Json Web Token

## Signup

    curl
      -X POST
      -H 'Content-Type: application/json'
      -d '{ "email":"<email>", "password":"<password>" }'
      http://localhost:8080/users

## Login

    curl
      -X POST
      -H 'Content-Type: application/json'
      -d '{ "email":"<email>", "password":"<password>" }'
      http://localhost:8080/sessions

## List Sessions

    curl
      -H 'Authorization: Bearer <token>'
      http://localhost:8080/sessions

## Logout

    curl
      -X DELETE
      -H 'Authorization: Bearer <token>'
      http://localhost:8080/sessions
