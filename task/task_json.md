# Create a Task

Once you have the workflow type id and you have got all the questions that this workflow requires you can submit a post request like the following and it will return with your task id & internal reference (Laird Reference).

Notes
----
1. The returned id is the task ID

2. Internal Rerfence is the Laird Assessors Rerference.

3. Dates are in UTC format yyyy-mm-dd

4. Workflow ID is required and is the type of report you wish us to do.

*  This is a **POST** request

Example Request
------

```
POST /api/v2/YOURAPIKEY/task.json HTTP/1.1
Host: test2-lairdassessors.swiftcase.co.uk
Content-Type: application/json
Cache-Control: no-cache

{
   "workflow_type_id": 70,
   "data" : [
	   { "name" : "client_staff_email", "value": "file.handler@yourcompany.com" },
	   { "name" : "principal_reference", "value": "your reference" },
	   { "name" : "inspection_contact_phone", "value": "01513429961" },
	   { "name" : "contact_for_inspection", "value": "Mr Smith" },
	   { "name" : "accident_date", "value": "2018-07-01" },
	   { "name" : "area_of_damage_to_vehicle_reported", "value": "rear" }
	]
}
```

Example Response
--------

```
{
    "id": 700109,
    "data": [
        {
            "name": "principal_reference",
            "value": "reference"
        },
        {
            "name": "inspection_contact_phone",
            "value": "inspection-phone"
        },
        {
            "name": "accident_date",
            "value": ""
        },
        {
            "name": "area_of_damage_to_vehicle_reported",
            "value": "area-of-damage"
        },
        {
            "name": "contact_for_inspection",
            "value": "inspection-contact-name"
        },
        {
            "name": "internal_reference",
            "value": "18-700105"
        }
    ],
    "fees": {
        "net": 150,
        "vat": 150,
        "gross": 150
    },
    "status": {
        "status": "Instruction Received"
    },
    "taskUsers": [
        {
            "user": {
                "id": 1401,
                "name": "Your Company Name"
            },
            "relationship": "Client"
        },
        {
            "user": {
                "id": 1401,
                "name": "Your Company Name"
            },
            "relationship": "Client staff"
        },
        {
            "user": {
                "id": 6001885,
                "name": "Expert, Laird"
            },
            "relationship": "Owner"
        }
    ]
}

```
