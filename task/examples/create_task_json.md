import requests

session = requests.Session()

rawBody = "{\r\n   \"workflow_type_id\": 70,\r\n   \"data\" : [\r\n\x09   { \"name\" : \"client_staff_email\", \"value\": \"file.handler@yourcompany.com\" },\r\n\x09   { \"name\" : \"principal_reference\", \"value\": \"your reference\" },\r\n\x09   { \"name\" : \"inspection_contact_phone\", \"value\": \"01513429961\" },\r\n\x09   { \"name\" : \"contact_for_inspection\", \"value\": \"Mr Smith\" },\r\n\x09   { \"name\" : \"accident_date\", \"value\": \"2018-07-01\" },\r\n\x09   { \"name\" : \"area_of_damage_to_vehicle_reported\", \"value\": \"rear\" }\r\n\x09]\r\n}"
headers = {"Cache-Control":"no-cache","Content-Type":"application/json"}
response = session.post("https://test2-lairdassessors.swiftcase.co.uk/api/v2/YOURAPIKEY/task.json", data=rawBody, headers=headers)

print("Status code:   %i" % response.status_code)
print("Response body: %s" % response.content)
