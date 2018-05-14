import requests

session = requests.Session()

headers = {"Cache-Control":"no-cache","Content-Type":"text/xml"}
response = session.post("https://test2-lairdassessors.swiftcase.co.uk/api/v2/YOURAPIKEY/task.xml", headers=headers)

print("Status code:   %i" % response.status_code)
print("Response body: %s" % response.content)
