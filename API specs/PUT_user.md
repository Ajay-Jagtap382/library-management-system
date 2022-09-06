##Description: This API is used to update name and password of a particular user

### HTTP Request
`PUT /user/{user_id}  `

### URL Parameters
{user_id ,name , password}

### Query Parameters
N/A


### Request Body
| Parameter | Format | Description                                |
|-----------|--------|--------------------------------------------|
| user_id   | int    | user id                                    |
| name      | String | name Id of user                            |
| Password  | String | Password of user                           |  


### Status codes and errors
| Value | Description           |
|-------|-----------------------|
| 200   | OK                    |
| 400   | Bad Request           |
| 403   | Forbidden             |
| 410   | Gone                  |
| 500   | Internal Server Error |

### Response Headers
N/A

### Success Response Body
```
{
    "Message": Field Updated Successfully "
    "User Info" :{}
}
```

### Bad Request Response when field updation failed
```
{
    "Message": "Invalid details. "
}
```

### Forbidden Response when user not present in request
```
{
    "Message": "Unable to verify token. Please contact your administrator"
}
```
