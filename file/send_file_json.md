# Send File JSON

This will send the file contents to the case

Important
----

1. You will need your task ID number to submit this request
2. You will need the API key
3. The contents of the file needs to be base64 encoded.
4. 201 will be the HTTP status code to confirm it has uploaded.


*  This is a **POST** request

Example Request
------

```
POST /api/v2/YOURAPIKEY/task/TASKID/files.json HTTP/1.1
Content-Type: application/json
cache-control: no-cache
Accept: */*
Host: test2-lairdassessors.swiftcase.co.uk
Accept-Encoding: gzip, deflate

{
    "name": "DSC08428.JPG",
    "type": "image/jpeg",
    "data": "///Z"
}
```

Example Response
--------

```
HTTP/1.1 201 Created
Server: nginx
Content-Type: application/json
Connection: close
X-Powered-By: PHP/7.0.30
Cache-Control: max-age=0, must-revalidate, private
Date: Tue, 15 May 2018 12:31:35 GMT
Strict-Transport-Security: max-age=31536000;
Content-Length: 0
```
