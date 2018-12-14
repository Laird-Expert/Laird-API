For Proclaim users some URLS are not possible so we have created some bridges to allow you to use the API.

We have also supplied the SSL certificates for openedge.

If you run in to issues with these endpoints we have made a bridge server you can host yourself to get past these issues.

https://github.com/Laird-Expert/Laird-API/tree/master/proclaim/bridge-server

Endpoints
 ---

### Test Endpoint

```
https://bridges.laird-assessors.com/test/
```

Live Endpoint

```
https://bridges.laird-assessors.com/live/
```


Examples
---


### Example for the Avalibility of engineer

This returns an ID number a date for a engineer please take the ID number and pass it to the next url as the BOOKID.


```
https://bridges.laird-assessors.com/test/engineer.php?APIKEY=YOURAPIKEY&TASK=YOURTASKID&IDATE=YOURDATE&maxresult=15
```




### Example for the Booking of an Engineer


```
https://bridges.laird-assessors.com/test/engineer.php?APIKEY=YOURAPIKEY&TASK=YOURTASKID&IDATE=YOURDATE&maxresult=15&BOOKID=BOOKID
```



### Example of Workflow Questions


```
https://bridges.laird-assessors.com/test/workflowsq.php?APIKEY=YOURAPIKEY&ID=WORKFLOWID
```


### Example of Status


```
https://bridges.laird-assessors.com/test/status.php?APIKEY=YOURAPIKEY&TASK=TASKID
```

### Example of Sending File

```
POST /test/sendfile.php?APIKEY=YOURAPIKEY&TASK=YOURTASKID HTTP/1.1
Host: bridges.laird-assessors.com
Content-Type: application/xml
Content-Length: 136514

<?xml version="1.0" encoding="UTF-8"?>
<file>
    <name>DSC08428.JPG</name>
    <type>image/jpeg</type>
    <data>/9j/4AAQSkZJRgABAQEASABIAAD/2wBDAAMCAgMCAgMDAwMEAwMEBQgFBQQEBQoHBwYIDAoMDAsKCwsNDhIQDQ4RDgsLEBYQERMUFRUVDA8XGBYUGBIUFRT/2wBDAQMEBAUEBQkFBQkUDQsNFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBT/wAARCAJTAlgDASIAAhEBAxEB/8QAHQAAAAcBAQEAAAAAAAAAAAAAAgMEBQYHCAABCf</data>
</file>
```

### Example of task details


```
https://bridges.laird-assessors.com/test/taskdetails.php?APIKEY=YOURAPIKEY&TASK=TASKID
```


### Example of get files

Take the ID number this returns and pass it as a the FILEID in the next URL.


```
https://bridges.laird-assessors.com/test/getfiles.php?APIKEY=YOURAPIKEY&TASK=TASKID
```



### Example of get file


```
https://bridges.laird-assessors.com/test/taskdetails.php?APIKEY=YOURAPIKEY&FILE=FILEID
```



