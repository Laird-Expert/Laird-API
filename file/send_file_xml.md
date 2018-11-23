# Send File XML

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
POST /api/v2/YOURAPIKEY/task/TASKID/file.xml
Content-Type: text/xml
Accept: */*
Host: test-lairdassessors.swiftcase.co.uk
Accept-Encoding: gzip, deflate

<?xml version="1.0" encoding="UTF-8"?>
<file>
    <name>DSC08428.JPG</name>
    <type>image/jpeg</type>
    <data>/9j/4AAQSkZJRgABAQEBXgFeAAD/4UVtRXhpZgAASUkqAAgAAAANAA4BAgAgAAAAqgAAAA8BAgAFAAAAygAAABABAgAJAAAA0AAAABIBAwABAAAAAQAAABoBBQABAAAA2gAAABsBBQABAAAA4gAAACgBAwABAAAAAgAAADEBAgAJAAAA6gAAADIBAgAUAAAA9AAAABMCAwABAAAAAgAAAJiCAgAPAAAACAEAAGmHBAABAAAANAEAAKXEBwAcAAAAGAEAAE4dAAAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgAFNPTlkAAERTQy1XODEwAABeAQAAAQAAAF4BAAABAAAAICAxLjA1MDAAADIwMTc6MDg6MDEgMTM6MzU6MzUAQ29weXJpZ2h0IDIwMDkAAFByaW50SU0AMDMwMAAAAgACAAEAAAABAQEAAAAhAJqCBQABAAAAxgIAAJ2CBQABAAAAzgIAACKIAwABAAAAAgAAACeIAwABAAAAoAAAAACQBwAEAAAAMDIzMAOQAgAUAAAA1gIAA</data>
</file>
```

Example Response
--------

```
HTTP/1.1 201 Created
Server: nginx
Content-Type: text/xml
Connection: close
X-Powered-By: PHP/7.0.30
Cache-Control: max-age=0, must-revalidate, private
Date: Tue, 15 May 2018 12:31:35 GMT
Strict-Transport-Security: max-age=31536000;
Content-Length: 0
```
