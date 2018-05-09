# Files JSON

This will get a list of files available for the task

Important
----

1. You will need your task ID number to submit this request
2. the file ID will be used in a later request to obtain the contents of the file.


*  This is a **GET** request

Example Request
------

```
GET /api/v2/YOURAPIKEY/task/YOURTASKID/files.json HTTP/1.1
Host: test2-lairdassessors.swiftcase.co.uk
Cache-Control: no-cache
```

Example Response
--------

```
{
    "files": [
        {
            "id": 3441,
            "name": "DSC08428.JPG",
            "type": "image/jpeg"
        },
        {
            "id": 3606,
            "name": "15-5555-Vehicle-Damage-Assessment-Report.pdf",
            "type": "application/pdf"
        },
        {
            "id": 3608,
            "name": "55-55555-Images.pdf",
            "type": "application/pdf"
        }
    ]
}
```
