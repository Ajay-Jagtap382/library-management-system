## Delete user
Description: This API will delete existing user

### HTTP Request
`DELETE/user/{userid}`

### URL Parameters
{userid}

### Query Parameters
N/A


### Request Headers
```
Content-Type: application/x-www-form-urlencoded
```

### Request Body
| Parameter | Format | Description                                    |
|-----------|--------|------------------------------------------------|
| Email     | String | Email Id of the user requesting password reset |
| password  | String | password of user                               |


### Sample cURL request
```

```

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
    "Message": User Deleted Successfully"
}
```

### Bad Request Response when Password validation failed
```
{
    "Message": "Invalid password. "
}
```

### Bad Request Response when the user doesn't exist
```
{
    "Message": "User doesn't exist."
}
```