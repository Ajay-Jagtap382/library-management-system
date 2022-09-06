##Description: This API is used to get the information of all the users

### HTTP Request
`GET /users `

### URL Parameters
N/A

### Query Parameters
N/A


### Request Body
 


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

	{This info will be displayed for all the users}
}
```