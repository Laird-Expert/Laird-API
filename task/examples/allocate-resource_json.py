import requests

session = requests.Session()

rawBody = "\r\n\r\n{\r\n    \"id\": 29965,\r\n    \"resource_availabilities\": [\r\n        {\r\n            \"date\": \"2018-05-14\"\r\n        }\r\n    ]\r\n}\r\n"
response = session.post("https://test2-lairdassessors.swiftcase.co.uk/api/v2/YOURAPIKEY/task/TASKID/allocate-resource.json", data=rawBody)

print("Status code:   %i" % response.status_code)
print("Response body: %s" % response.content)

