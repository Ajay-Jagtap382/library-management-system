##Description: This API is used to get the information of a particular user

### HTTP Request
`GET /user `

### URL Parameters
{user_id}

### Query Parameters
N/A


### Request Body
| Parameter | Format | Description                                |
|-----------|--------|--------------------------------------------|
| user_id   | int    | user id                                    |


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
| Parameter           | Format | Description                                    |
|---------------------|--------|------------------------------------------------|
| id                  | int    | Unique ID of the user                          |
| Email               | String | Email Id of the user requesting password reset |
| name		    | String | Name of the user                               |
| phone               | varchar| phone number of the user                       |
| profile_picture     | image  | image of the user                              |  
}

### Bad Request Response when ID validation failed
```
{
    "Message": "Invalid ID."
}
```