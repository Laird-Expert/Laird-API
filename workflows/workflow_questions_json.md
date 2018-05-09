# Workflow Questions JSON

Once you have the workflow ID of the report your wish to instruct on you will then need to work out what details are accepted for each workflow.

This will list all the parameters that are expected to be submitted.

You can leave the value of the parameters blank but ever detail helps.

*  This is a **GET** request

Example Request
------

```
GET /api/v2/YOURAPIKEY/workflow/128.json HTTP/1.1
User-Agent: Laird Assessors API Example
Accept: */*
Host: test2-lairdassessors.swiftcase.co.uk
Accept-Encoding: gzip, deflate
Connection: close
```

Example Response
--------

```
{
    "workflow_questions": [
        {
            "question_key": "internal_reference", 
            "name": "Internal Reference", 
            "id": 6
        }, 
        {
            "question_key": "principal_reference", 
            "name": "Principal Reference", 
            "id": 9
        }, 
        {
            "question_key": "type_of_document_collection", 
            "name": "Type of Document Collection", 
            "id": 273
        }, 
        {
            "question_key": "client_salutation", 
            "name": "Client Salutation", 
            "id": 45
        }, 
        {
            "question_key": "client_first_name", 
            "name": "Client First Name", 
            "id": 46
        }, 
        {
            "question_key": "client_second_name", 
            "name": "Client Second Name", 
            "id": 47
        }, 
        {
            "question_key": "client_postcode", 
            "name": "Client Address", 
            "id": 53
        }, 
        {
            "question_key": "client_email_address", 
            "name": "Client Email Address", 
            "id": 54
        }, 
        {
            "question_key": "client_daytime_phone", 
            "name": "Client Daytime Phone", 
            "id": 55
        }, 
        {
            "question_key": "client_evening_phone", 
            "name": "Client Evening Phone", 
            "id": 56
        }, 
        {
            "question_key": "client_mobile", 
            "name": "Client Mobile", 
            "id": 57
        }, 
        {
            "question_key": "client_phone_notes", 
            "name": "Client Phone Notes", 
            "id": 58
        }, 
        {
            "question_key": "standard_appointment_date_time", 
            "name": "Appointment Date & Time", 
            "id": 271
        }, 
        {
            "question_key": "appointment_location_postcode", 
            "name": "Appointment Location", 
            "id": 272
        }, 
        {
            "question_key": "nature_of_case", 
            "name": "Nature of Case", 
            "id": 27
        }, 
        {
            "question_key": "agent_special_instructions", 
            "name": "Expert Instructions", 
            "id": 41
        }, 
        {
            "question_key": "additional_information", 
            "name": "Additional Information", 
            "id": 42
        }, 
        {
            "question_key": "report_handler", 
            "name": "Report handler", 
            "id": 5
        }
    ]
}

```

