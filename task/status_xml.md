# status XML

This will advise you on the status of your case

You will need to provide the task ID that you originally recived when you submitted the task.

*  This is a **GET** request

Example Request
------

```
GET GET /api/v2/YOURAPIKEY/task/YOURTASKID/status.xml HTTP/1.1
Host: test2-lairdassessors.swiftcase.co.uk
Cache-Control: no-cache
```

Example Response
--------

```
<?xml version="1.0" encoding="UTF-8"?>
<taskStatus>
    <status>Engineer Validation</status>
</taskStatus>

```
