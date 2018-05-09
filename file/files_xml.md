# Files XML

This will get a list of files available for the task

Important
----

1. You will need your task ID number to submit this request
2. the file ID will be used in a later request to obtain the contents of the file.


*  This is a **GET** request

Example Request
------

```
GET /api/v2/YOURAPIKEY/task/YOURTASKID/files.xml HTTP/1.1
Host: test2-lairdassessors.swiftcase.co.uk
Cache-Control: no-cache
```

Example Response
--------

```
<?xml version="1.0" encoding="UTF-8"?>
<fileCollection>
    <file>
        <id>3426441</id>
        <name>DSC08428.JPG</name>
        <type>image/jpeg</type>
    </file>
    <file>
        <id>3426606</id>
        <name>55-55555-Vehicle-Damage-Assessment-Report.pdf</name>
        <type>application/pdf</type>
    </file>
    <file>
        <id>3426608</id>
        <name>55-5555555-Images.pdf</name>
        <type>application/pdf</type>
    </file>
</fileCollection>
```
