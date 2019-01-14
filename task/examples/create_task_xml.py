import requests

url = "https://test-lairdassessors.swiftcase.co.uk/api/v2/YOUR API KEY/task.xml"

payload = "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\r\n<request>\r\n   <workflow_type_id>70</workflow_type_id>\r\n   <data>\r\n\t   <item name=\"client_staff_email\">lee.baty@laird-assessors.com</item>\r\n\t   <item name=\"principal_reference\">LBATY CRASH</item>\r\n\t   <item name=\"inspection_contact_phone\">01513429961</item>\r\n\t   <item name=\"contact_for_inspection\">Lee Baty</item>\r\n\t   <item name=\"accident_date\">01/01/2019</item>\r\n\t   <item name=\"area_of_damage_to_vehicle_reported\">rear</item>\r\n\t   <item name=\"contact_for_inspection\">Mr Lee Baty</item>\r\n\t   <item name=\"client_first_name\">Lee</item>\r\n\t   <item name=\"client_second_name\">Baty</item>\r\n\t   <item name=\"client_mobile\">0777777777</item>\r\n\t   <item name=\"client_phone_notes\">ring after 5pm</item>\r\n\t   <item name=\"vehicle_registration\">&lt;registration&gt;W705ACW&lt;/registration&gt;</item>\r\n\t   <item name=\"agent_special_instructions\">Knock around the rear</item>\r\n\t   <item name=\"principal_comments\">Car was turning left when other car collideded in to the side</item>\r\n\t   <item name=\"file_referrer\">Turner and Hooch</item>\r\n\t   <item name=\"client_salutation\">Mr</item>\r\n\t   <item name=\"inspection_address\">&lt;postcode&gt;PO1 0DE&lt;/postcode&gt;&lt;first_line&gt;First line&lt;/first_line&gt;&lt;second_line&gt;Second Line&lt;/second_line&gt;&lt;town&gt;TOWN&lt;/town&gt;&lt;county&gt;COUNTY&lt;/county&gt;</item>\r\n   </data>\r\n</request>"
headers = {
    'Content-Type': "text/xml",
    'cache-control': "no-cache"
    }

response = requests.request("POST", url, data=payload, headers=headers)

print(response.text)
