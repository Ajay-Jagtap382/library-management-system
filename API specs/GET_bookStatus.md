##Description: This API is used to get the status of a particular book

### HTTP Request
`GET /book/{book_id}/status `

### URL Parameters
{book_id}

### Query Parameters
N/A


### Request Body
| Parameter | Format | Description                                |
|-----------|--------|--------------------------------------------|
| book_id   | int    | book id                                    |


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
| id                  | int    | Unique ID of the book                          |
| price               | int    | Price of the book                              |
| borrowed_at         | date   | Date at which book is borrowed                 |
| due_date            | date   | Date at which book should be return            |
| return_date         | date   | Date at which book is returned                 |
| lend_by             | int    | user id of the user who borrowed the book      |
| return_date         | date   | Date at which book is returned                 |  
}
```
### Bad Request Response when ID validation failed
```
{
    "Message": "Invalid ID."
}
```

