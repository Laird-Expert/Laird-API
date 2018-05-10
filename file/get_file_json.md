# Get File JSON

This will get the contents of a file you request

Important
----

1. You will need your task ID number to submit this request
2. You Will need the file ID from the files listing endpoint
3. The API will return the contents of the file base64 encoded so you can save the file name and then decode the contents to the file.


*  This is a **GET** request

Example Request
------

```
GET /api/v2/YOURAPKEY/file/FILEID.json  HTTP/1.1
Host: test2-lairdassessors.swiftcase.co.uk
Cache-Control: no-cache
```

Example Response
--------

```

{
    "id": 3426441,
    "name": "DSC08428.JPG",
    "type": "image/jpeg",
    "data": "/9j/4AAQSkZJRgABAQEBXgFeAAD/4UVtRXhpZgAASUAAAAAAAAAAAAAAAAAAAAAAAAA"
 }
```
