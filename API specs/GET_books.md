##Description: This API is used to get the information of all the books

### HTTP Request
`GET /books `

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
| name		    | String | Name of the book                               |
| phone               | varchar| phone number of the user                       |
| image               | image  | image of the user                              |  
| total_copies        | int    | total copies of the book                       |
| current_copies      | int    | available copies of the book                   |

	{This info will be displayed for all the books}

}
```