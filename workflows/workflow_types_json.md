# Workflow Types JSON

This will advise you on the workflow type (report type) that is avalible to your account.

*  This is a **GET** request

Example Request
------

```
GET /api/v2/YOURAPIKEY/workflows.json HTTP/1.1
Host: test2-lairdassessors.swiftcase.co.uk
Cache-Control: no-cache
```

Example Response
--------

```
{
    "workflow_types": [
        {
            "id": 193,
            "name": "Car - Desktop Vehicle Damage Accident Report"
        },
        {
            "id": 70,
            "name": "Car - Vehicle Damage Accident Report"
        },
        {
            "id": 101,
            "name": "Consistency Interim Report"
        },
        {
            "id": 245,
            "name": "Desktop Diminution Report"
        },
        {
            "id": 206,
            "name": "Desktop Locus in Quo"
        },
        {
            "id": 90,
            "name": "Diminution Report (Attended)"
        },
        {
            "id": 128,
            "name": "Document Collection"
        },
        {
            "id": 207,
            "name": "Injury Photography"
        },
        {
            "id": 112,
            "name": "Interpretation - Court Att"
        },
        {
            "id": 246,
            "name": "Interpretation Service"
        },
        {
            "id": 244,
            "name": "Legacy Product"
        },
        {
            "id": 205,
            "name": "Locus in Quo Report"
        },
        {
            "id": 92,
            "name": "Post Repair Report"
        },
        {
            "id": 208,
            "name": "Public Liability Report"
        },
        {
            "id": 149,
            "name": "Query - Agree Figures"
        },
        {
            "id": 150,
            "name": "Query - Authorisation"
        },
        {
            "id": 151,
            "name": "Query - Checked Account"
        },
        {
            "id": 156,
            "name": "Query - Client Area of Damage Query"
        },
        {
            "id": 152,
            "name": "Query - Client Estimate"
        },
        {
            "id": 153,
            "name": "Query - Client PAV"
        },
        {
            "id": 154,
            "name": "Query - Client Salvage Query"
        },
        {
            "id": 155,
            "name": "Query - Client Temp Repair Query"
        },
        {
            "id": 157,
            "name": "Query - Diminution Query"
        },
        {
            "id": 158,
            "name": "Query - Diminution Report"
        },
        {
            "id": 159,
            "name": "Query - Forensic Review"
        },
        {
            "id": 160,
            "name": "Query - Garage Invoice"
        },
        {
            "id": 161,
            "name": "Query - Garage Supplementary"
        },
        {
            "id": 162,
            "name": "Query - P35Q"
        },
        {
            "id": 163,
            "name": "Query - Previous Total Loss Query"
        },
        {
            "id": 164,
            "name": "Query - Principal Area of Damage Query"
        },
        {
            "id": 165,
            "name": "Query - Principal PAV"
        },
        {
            "id": 166,
            "name": "Query - Principal Salvage Query"
        },
        {
            "id": 167,
            "name": "Query - Principal Temp Repair Query"
        },
        {
            "id": 171,
            "name": "Query - Report Error LA"
        },
        {
            "id": 172,
            "name": "Query - Report Error Sub"
        },
        {
            "id": 173,
            "name": "Query - RW Principal Query"
        },
        {
            "id": 174,
            "name": "Query - RW TP Query"
        },
        {
            "id": 175,
            "name": "Query - TP Area of Damage Query"
        },
        {
            "id": 176,
            "name": "Query - TP Comment on LA Report"
        },
        {
            "id": 177,
            "name": "Query - TP Engineer Report Review"
        },
        {
            "id": 178,
            "name": "Query - TP PAV"
        },
        {
            "id": 179,
            "name": "Query - TP Salvage Query"
        },
        {
            "id": 180,
            "name": "Query - TP Temp Repair Query"
        },
        {
            "id": 181,
            "name": "Query - URGENT Principal Request"
        },
        {
            "id": 260,
            "name": "Terms Sign-up"
        },
        {
            "id": 119,
            "name": "Translation Service"
        }
    ]
}

```
