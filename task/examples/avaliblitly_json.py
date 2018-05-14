import requests

session = requests.Session()

headers = {"Cache-Control":"no-cach","Content-Type":"application/x-www-form-urlencoded"}
response = session.get("https://test2-lairdassessors.swiftcase.co.uk/api/v2/YOURAPIKEY/task/TASKID/availability/DATEFROM/MAXRESULTS.json", headers=headers)

print("Status code:   %i" % response.status_code)
print("Response body: %s" % response.content)
