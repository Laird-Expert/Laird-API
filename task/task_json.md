# Create a Task JSON

Once you have the workflow type id and you have got all the questions that this workflow requires you can submit a post request like the following and it will return with your task id & internal reference (Laird Reference).

Notes
----
1. The returned id is the task ID

2. Internal Rerfence is the Laird Assessors Rerference.

3. Dates should be submitted d/m/Y (21/5/2018)

4. Workflow ID is required and is the type of report you wish us to do.

5. Task User is the Laird Assessors member of staff the case has been automatically sent to. (This may change depending on staff avaliablity)

6. Status is us acknowledging we have rececived the instruction.

7. Not all questions need to be sent over but please send us all fields that can be filled in by yourself so that we can proceed quicker.

8. HTTP 201 confirms the request was successful

9. Inspection address and vehicle registration have to be passed with json object stings see example as there are additonal fields for these special keys.

*  This is a **POST** request

Example Request
------

```
POST /api/v2/YOURAPIKEY/task.json HTTP/1.1
Content-Type: application/json
cache-control: no-cache
User-Agent: PostmanRuntime/7.1.1
Accept: */*
Host: test2-lairdassessors.swiftcase.co.uk
Accept-Encoding: gzip, deflate
Connection: close

{
	"workflow_type_id": 70,
	"data": [{
			"name": "client_staff_email",
			"value": "user@sols.com"
		},
		{
			"name": "principal_reference",
			"value": "reference"
		},
		{
			"name": "inspection_contact_phone",
			"value": "inspection-phone"
		},
		{
			"name": "contact_for_inspection",
			"value": "inspection-contact-name"
		},
		{
			"name": "accident_date",
			"value": "5//5/2017"
		},
		{
			"name": "area_of_damage_to_vehicle_reported",
			"value": "area-of-damage"
		},
		{
			"name": "inspection_address",
			"value": "{\"postcode\":\"CH47 4BP\",\"first_line\":\"188 - 200 Whitfield Buildings\",\"second_line\":\"Pensby\",\"town\":\"Pensby\",\"county\":\"Wirral\"}"
		},
		{
			"name": "vehicle_registration",
			"value": "{\"registration\":\"YR57AEJ\",\"vin\":\"WF0BXXBAJBYT66489\",\"make\":\"Ford\",\"model\":\"Fiesta Freestyle\",\"type\":\"Five door hatchback\",\"colour\":\"Black\"}"
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
		}
	]
}
```

Example Response
--------

```
HTTP/1.1 201 Created
Server: nginx
Content-Type: application/json
Connection: close
X-Powered-By: PHP/7.0.30
Cache-Control: max-age=0, must-revalidate, private
Date: Wed, 27 Jun 2018 11:00:11 GMT
Strict-Transport-Security: max-age=31536000;
Content-Length: 1444

{"id":700332,"data":[{"name":"principal_reference","value":"reference"},{"name":"inspection_contact_phone","value":"inspection-phone"},{"name":"accident_date","value":""},{"name":"area_of_damage_to_vehicle_reported","value":"Rear"},{"name":"agent_special_instructions","value":"Ring 2hrs before arrival"},{"name":"principal_comments","value":"Important Client"},{"name":"agent_important_special_instructions","value":"Car has been in previouse accident on 17th may 2017 and was repaired"},{"name":"file_referrer","value":"MD accident management"},{"name":"contact_for_inspection","value":"Lee Baty"},{"name":"inspection_address","value":"CH47 4BP"},{"name":"vehicle_registration","value":"YR57AEJ"},{"name":"client_salutation","value":"Mr"},{"name":"client_postcode","value":"C"},{"name":"client_email_address","value":"lee.baty@laird-assessors.com"},{"name":"client_daytime_phone","value":"01513429961"},{"name":"client_evening_phone","value":"01513429961"},{"name":"client_mobile","value":"0777777777"},{"name":"client_phone_notes","value":"VIP"},{"name":"internal_reference","value":"18-700301"}],"fees":{"net":80,"vat":16,"gross":96},"status":{"status":"Instruction Received"},"taskUsers":[{"user":{"id":1401,"name":"Easi Drive Accident Management"},"relationship":"Client"},{"user":{"id":1401,"name":"Easi Drive Accident Management"},"relationship":"Client staff"},{"user":{"id":6001946,"name":"Sharples, Hannah"},"relationship":"Owner"}]}

```
