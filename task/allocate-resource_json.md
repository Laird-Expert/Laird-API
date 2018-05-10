```
POST /api/v2/YOURAPIKEY/task/TASKID/allocate-resource.json HTTP/1.1
Host: test2-lairdassessors.swiftcase.co.uk


ID - is taken from avaliblity resources.

{
    "id": 29965,
    "resource_availabilities": [
        {
            "date": "2018-05-14"
        }
    ]
}
```



```
"Resource allocated."
```



NOTE: 204 No Content means there has been an error and nothing has been booked.
