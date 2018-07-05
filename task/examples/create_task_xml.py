import requests

session = requests.Session()



payload = "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\r\n<request>\r\n   <workflow_type_id>70</workflow_type_id>\r\n   <data>\r\n\t   <item name=\"client_staff_email\">test@test.com</item>\r\n\t   <item name=\"principal_reference\">reference</item>\r\n\t   <item name=\"inspection_contact_phone\">inspection-phone</item>\r\n\t   <item name=\"contact_for_inspection\">inspection-contact-name</item>\r\n\t   <item name=\"accident_date\">1/1/2017</item>\r\n\t   <item name=\"area_of_damage_to_vehicle_reported\">area-of-damage</item>\r\n   </data>\r\n</request>"
headers = {"Cache-Control":"no-cache","Content-Type":"text/xml"}
response = session.post("https://test2-lairdassessors.swiftcase.co.uk/api/v2/YOURAPIKEY/task.xml", headers=headers)

response = requests.request("POST", url, data=payload, headers=headers)


print("Status code:   %i" % response.status_code)
print("Response body: %s" % response.content)
