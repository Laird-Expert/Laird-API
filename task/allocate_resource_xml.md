# allocate-resource XML

Allocate resource is where you would book a engineer for a physical inspection.

*  This is a **POST** request

You need to submit the following:
---

*  API Key
*  TaskID
*  ID is taken from avaliblity resources.
*  Date in UTC format


Example Request
-----

```
POST /api/v2/YOURAPIKEY/task/TASKID/allocate-resource.xml HTTP/1.1
Host: test-lairdassessors.swiftcase.co.uk
Content-Type: text/xml

<?xml version="1.0" encoding="UTF-8"?>
<resourceAvailability>
    <id>29988</id>
    <availability>
        <date>2018-05-11</date>
    </availability>
</resourceAvailability>
```


Example Response
-----


```
HTTP/1.1 201 Created
Server: nginx
Content-Type: text/xml
Connection: close
X-Powered-By: PHP/7.0.30
Cache-Control: max-age=0, must-revalidate, private
Date: Tue, 15 May 2018 14:14:26 GMT
Strict-Transport-Security: max-age=31536000;
Content-Length: 0
```

Notes
----

If you get a HTTP status code of 204 this means there has been an error and the inspection has not been booked for that date.

If you require a specific date and get this result please contact the our logistics department to arrange the inspection.
