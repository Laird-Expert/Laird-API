# Create a Task JSON

Once you have the workflow type id and you have got all the questions that this workflow requires you can submit a post request like the following and it will return with your task id & internal reference (Laird Reference).

Notes
----
1. The returned id is the task ID

2. Internal Rerfence is the Laird Assessors Rerference.

3. Dates are in UTC format yyyy-mm-dd

4. Workflow ID is required and is the type of report you wish us to do.

5. Task User is the Laird Assessors member of staff the case has been automatically sent to. (This may change depending on staff avaliablity)

6. Status is us acknowledging we have rececived the instruction.

7. Not all questions need to be sent over but please send us all fields that can be filled in by yourself so that we can proceed quicker.

*  This is a **POST** request

Example Request
------

```
POST /api/v2/YOURAPIKEY/task.json HTTP/1.1
Content-Type: application/json
User-Agent: API DEMO
Accept: */*
Host: test2-lairdassessors.swiftcase.co.uk
Accept-Encoding: gzip, deflate
Connection: close

{
   "workflow_type_id": 70,
   "data" : [
        
        {
            "value": "Principal Reference",
            "name": "principal_reference"
        },
        {
            "value": "2018-01-01",
            "name": "accident_date"
        },
        {
            "value": "01513429961",
            "name": "inspection_contact_phone"
        },
        {
            "value": "0777777777",
            "name": "inspection_contact_mobile"
        },
        {
            "value": "Ring 2hrs before arrival",
            "name": "agent_special_instructions"
        },
        {
            "value": "Important Client",
            "name": "principal_comments"
        },
        {
            "value": "Rear",
            "name": "area_of_damage_to_vehicle_reported"
        },
        {
            "value": "Car has been in previouse accident on 17th may 2017 and was repaired",
            "name": "agent_important_special_instructions"
        },
        {
            "value": "MD accident management",
            "name": "file_referrer"
        },
        {
            "value": "Mr",
            "name": "client_salutation"
        },
        {
            "value": "Lee",
            "name": "client_first_value"
        },
        {
            "value": "Baty",
            "name": "client_second_value"
        },
        {
            "value": "CH607RJ",
            "name": "client_postcode"
        },
        {
            "value": "lee.baty@laird-assessors.com",
            "name": "client_email_address"
        },
        {
            "value": "01513429961",
            "name": "client_daytime_phone"
        },
        {
            "value": "01513429961",
            "name": "client_evening_phone"
        },
        {
            "value": "0777777777",
            "name": "client_mobile"
        },
        {
            "value": "VIP",
            "name": "client_phone_notes"
        },
        {
            "value": "Lee Baty",
            "name": "contact_for_inspection"
        },
        {
            "value": "Laird Assesors 188 - 200 Whitfield Buildings,heswall,wirral",
            "name": "inspection_address"
        },
        {
            "value": "W705ACW",
            "name": "vehicle_registration"
        },
        {
            "value": "In Use",
            "name": "on_site_in_use"
        }
     

         


	]
}
```

Example Response
--------

```
{
  "id": 700168,
  "data": [
    {
      "name": "principal_reference",
      "value": "Principal Reference"
    },
    {
      "name": "accident_date",
      "value": ""
    },
    {
      "name": "inspection_contact_phone",
      "value": "01513429961"
    },
    {
      "name": "inspection_contact_mobile",
      "value": "0777777777"
    },
    {
      "name": "agent_special_instructions",
      "value": "Ring 2hrs before arrival"
    },
    {
      "name": "principal_comments",
      "value": "Important Client"
    },
    {
      "name": "area_of_damage_to_vehicle_reported",
      "value": "Rear"
    },
    {
      "name": "agent_important_special_instructions",
      "value": "Car has been in previouse accident on 17th may 2017 and was repaired"
    },
    {
      "name": "file_referrer",
      "value": "MD accident management"
    },
    {
      "name": "client_salutation",
      "value": "Mr"
    },
    {
      "name": "client_postcode",
      "value": "C"
    },
    {
      "name": "client_email_address",
      "value": "lee.baty@laird-assessors.com"
    },
    {
      "name": "client_daytime_phone",
      "value": "01513429961"
    },
    {
      "name": "client_evening_phone",
      "value": "01513429961"
    },
    {
      "name": "client_mobile",
      "value": "0777777777"
    },
    {
      "name": "client_phone_notes",
      "value": "VIP"
    },
    {
      "name": "contact_for_inspection",
      "value": "Lee Baty"
    },
    {
      "name": "inspection_address",
      "value": "L"
    },
    {
      "name": "vehicle_registration",
      "value": "W705ACW"
    },
    {
      "name": "on_site_in_use",
      "value": "In Use"
    },
    {
      "name": "internal_reference",
      "value": "18-700162"
    }
  ],
  "fees": {
    "net": 50,
    "vat": 10,
    "gross": 60
  },
  "status": {
    "status": "Instruction Received"
  },
  "taskUsers": [
    {
      "user": {
        "id": 7,
        "name": "Client Testing Principal"
      },
      "relationship": "Client"
    },
    {
      "user": {
        "id": 7,
        "name": "Client Testing Principal"
      },
      "relationship": "Client staff"
    },
    {
      "user": {
        "id": 6006887,
        "name": "Markey, Sarah"
      },
      "relationship": "Owner"
    }
  ]
}

```
