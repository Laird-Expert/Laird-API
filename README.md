# Laird-API
Examples for the Laird Assessors API


Please note all examples will have YOURAPIKEY which is where you need to replace that with your actual API key.

API Keys will be provided to you on request and you will be provided with a test and live api key.

The API will support request and responses in JSON or XML format.

Documentation from swiftcase - [Google Docs](https://docs.google.com/document/d/1K12qIap1dvEQLvJigxjzOE5EH9VgYWicSsTfFjpKn0M/edit?ts=5afa9b0f)


Test Endpoint
----

```
https://test2-lairdassessors.swiftcase.co.uk
```

Live Endpoint
----

```
TBC
```


Getting Started
-------

Once you have your API keys you will need to work out all the different report types.

[Grab Workflow IDs](https://github.com/Laird-Expert/Laird-API/blob/master/workflows/workflow_types_json.md) - JSON

[Grab Workflow IDs](https://github.com/Laird-Expert/Laird-API/blob/master/workflows/workflow_types_xml.md) - XML

[Grab Workflow Questions](https://github.com/Laird-Expert/Laird-API/blob/master/workflows/workflow_questions_json.md) - JSON

[Grab Workflow Questions](https://github.com/Laird-Expert/Laird-API/blob/master/workflows/workflow_questions_xml.md) - XML

[Create a Task](https://github.com/Laird-Expert/Laird-API/blob/master/task/task_json.md) -JSON

[Create a Task](https://github.com/Laird-Expert/Laird-API/blob/master/task/task_xml.md) - XML

[Avaliblity](https://github.com/Laird-Expert/Laird-API/blob/master/task/avalibility_json.md)

[Avaliblity](https://github.com/Laird-Expert/Laird-API/blob/master/task/avalibility_xml.md) - XML

[Allocate Resource](https://github.com/Laird-Expert/Laird-API/blob/master/task/allocate-resource_json.md) - JSON

[Allocate Resource](https://github.com/Laird-Expert/Laird-API/blob/master/task/allocate-resource_xml.md) - XML

[Check Status](https://github.com/Laird-Expert/Laird-API/blob/master/task/status_json.md) - JSON

[Check Status](https://github.com/Laird-Expert/Laird-API/blob/master/task/status_xml.md) - XML

[Task Details](https://github.com/Laird-Expert/Laird-API/blob/master/task/task_details_json.md) - JSON

[Task Details](https://github.com/Laird-Expert/Laird-API/blob/master/task/task_details_xml.md) - XML

[List Files](https://github.com/Laird-Expert/Laird-API/blob/master/file/get_files_json.md) - JSON

[List Files](https://github.com/Laird-Expert/Laird-API/blob/master/file/get_files_xml.md) - XML

[Grab File](https://github.com/Laird-Expert/Laird-API/blob/master/file/get_file_json.md) - JSON

[Grab File](https://github.com/Laird-Expert/Laird-API/blob/master/file/get_file_xml.md) - XML

Workflow
----

1. Get ID of report (workflow)
2. Get Questions for submission of report (workflow questions)
3. Submit details (submit task)
4. if physical inspection
   Check avaliablity
   allocate resource
5. check status
6. task details (details of the inspection/report)
7. list files
8. grab file (grab images and report)

[![VDAR](https://s31.postimg.cc/wxd1cuqyj/Capture.png)](https://postimg.cc/image/hon3z2x9z/)



