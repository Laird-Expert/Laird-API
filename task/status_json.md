# status JSON

This will advise you on the status of your case

You will need to provide the task ID that you originally recenived when you submitted the task.

*  This is a **GET** request

Example Request
------

```
GET GET /api/v2/YOURAPIKEY/task/YOURTASKID/status.json HTTP/1.1
Host: test2-lairdassessors.swiftcase.co.uk
Cache-Control: no-cache
```

Example Response
--------

```
{
    "status": "Engineer Validation"
}

```
