## Add Book API
Deescrription : This API will be used to add new books in library

### HTTP Request
`POST/book`

### URL Parameters
/user/book/create

### Query Parameters
N/A


### Request Headers
```
Content-Type: application/x-www-form-urlencoded
```

### Request Body
| Parameter           | Format | Description                                    |
|---------------------|--------|------------------------------------------------|
| name		    | String | Name of the book                               |
| phone               | varchar| phone number of the user                       |
| image               | image  | image of the user                              |  
| total_copies        | int    | total copies of the book                       |
| current_copies      | int    | available copies of the book                   |



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
    "Message": Added Book Successfully "
```

### Bad Request Response when book addition failed
```
{
    "Message": "Book addition failed, please try again"
}
```

### Forbidden Response when role doesn't match
```
{
    "Message": "Access restricted"
}
```