#!/usr/bin/env python
# -*- coding: UTF-8 -*-
import requests

session = requests.Session()

YOURAPIKEY = "INSERT API KEY"

headers = {"User-Agent":"Laird Assessors API Example","Connection":"close","Accept-Encoding":"gzip, deflate","Accept":"*/*"}
response = session.get("https://test-lairdassessors.swiftcase.co.uk/api/v2/"+YOURAPIKEY+"/workflow/70.xml", headers=headers)

print("Status code:   %i" % response.status_code)
print("Response body: %s" % response.content)
