# allocate-resource JSON

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
POST /api/v2/YOURAPIKEY/task/TASKID/allocate-resource.json HTTP/1.1
Host: test2-lairdassessors.swiftcase.co.uk



{
    "id": 29965,
    "resource_availabilities": [
        {
            "date": "2018-05-14"
        }
    ]
}
```

Example Response
-----

```
"Resource allocated."
```

Notes
----

If you get a HTTP status code of 204 this means there has been an error and the inspection has not been booked for that date.

If you require a specific date and get this result please contact the our logistics department to arrange the inspection.
