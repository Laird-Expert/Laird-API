Bridge Server
---

Due to some issues with some clients having SSL errors I have built a small server that will act like the bridge on your server for you.

I have attached some pre-compiled binarys but if you have go installed you can compile it your self with go build.

You can run the binary and add --port 5555 if you wish to start the bridge server on another port other than 5050

Logging can be seen via the CLI to see if there are any issues.


### Endpoints

---


### Example for Engineer Availability


```
https://localhost:5050/test/engineer?APIKEY=YOURAPIKEY&TASK=YOURTASKID&IDATE=YOURDATE&maxresult=15
```

```
https://localhost:5050/live/engineer?APIKEY=YOURAPIKEY&TASK=YOURTASKID&IDATE=YOURDATE&maxresult=15
```


### Example for Engineer Booking


```
https://localhost:5050/test/bookengineer?APIKEY=YOURAPIKEY&TASK=YOURTASKID&IDATE=YOURDATE&BOOKID=1512312
```

```
https://localhost:5050/live/bookengineer?APIKEY=YOURAPIKEY&TASK=YOURTASKID&IDATE=YOURDATE&BOOKID=1231231
```

### Example of Get Files


```
https://localhost:5050/test/getfiles?APIKEY=YOURAPIKEY&TASK=YOURTASKID
```

```
https://localhost:5050/live/getfiles?APIKEY=YOURAPIKEY&TASK=YOURTASKID
```


### Example of Get File


```
https://localhost:5050/test/getfile?APIKEY=YOURAPIKEY&FILEID=YOURFILEID
```

```
https://localhost:5050/live/getfile?APIKEY=YOURAPIKEY&FILEID=YOURFILEID
```


### Example of Download File


```
https://localhost:5050/test/downloadfile?APIKEY=YOURAPIKEY&FILEID=YOURFILEID
```

```
https://localhost:5050/live/downloadfile?APIKEY=YOURAPIKEY&FILEID=YOURFILEID
```



### Example of Task Details


```
https://localhost:5050/test/taskdetails?APIKEY=YOURAPIKEY&TASK=YOURTASKID
```

```
https://localhost:5050/live/taskdetails?APIKEY=YOURAPIKEY&TASK=YOURTASKID
```


### Example of Workflows


```
https://localhost:5050/test/workflows?APIKEY=YOURAPIKEY
```

```
https://localhost:5050/live/workflows?APIKEY=YOURAPIKEY
```


### Example of Workflow Questions


```
https://localhost:5050/test/workflow?APIKEY=YOURAPIKEY&ID=WORKFLOWID
```

```
https://localhost:5050/live/workflow?APIKEY=YOURAPIKEY&ID=WORKFLOWID
```

### Example of Status


```
https://localhost:5050/test/status?APIKEY=YOURAPIKEY&TASK=TASKID
```

```
https://localhost:5050/live/status?APIKEY=YOURAPIKEY&TASK=TASKID
```


### Example of Task

```
POST /test/task?APIKEY=YOURAPIKEY HTTP/1.1
Content-Type: text/xml
cache-control: no-cache
User-Agent: PostmanRuntime/7.6.0
Accept: */*
Host: 127.0.0.1
Accept-Encoding: gzip, deflate
Content-Length: 1382
Connection: close

<?xml version="1.0" encoding="UTF-8"?>
<request>
   <workflow_type_id>70</workflow_type_id>
   <data>
	   <item name="client_staff_email">lee.baty@laird-assessors.com</item>
	   <item name="principal_reference">LBATY CRASH</item>
	   <item name="inspection_contact_phone">01513429961</item>
	  <item name="inspection_contact_mobile">01513429961</item>
	   <item name="contact_for_inspection">Lee Baty</item>
	   <item name="accident_date">01/01/2019</item>
	   <item name="area_of_damage_to_vehicle_reported">rear</item>
	   <item name="contact_for_inspection">Mr Lee Baty</item>
	   <item name="client_first_name">Lee</item>
	   <item name="client_second_name">Baty</item>
	   <item name="client_mobile">0777777777</item>
	   <item name="client_phone_notes">ring after 5pm</item>
	   <item name="vehicle_registration">&lt;registration&gt;W705ACW&lt;/registration&gt;</item>
	   <item name="agent_special_instructions">Knock around the rear</item>
	   <item name="principal_comments">Car was turning left when other car collideded in to the side</item>
	   <item name="file_referrer">Turner and Hooch</item>
	   <item name="client_salutation">Mr</item>
	   <item name="inspection_address"><postcode>PO1 0DE</postcode><first_line>First line</first_line><second_line>Second Line</second_line><town>TOWN</town><county>COUNTY</county></item>
   </data>
</request>
```


```
POST /live/task?APIKEY=YOURAPIKEY HTTP/1.1
Host: 127.0.0.1
Content-Type: text/xml
cache-control: no-cache
Content-Length: 768

<?xml version="1.0" encoding="UTF-8"?>
<request>
   <workflow_type_id>70</workflow_type_id>
   <data>
	   <item name="client_staff_email">lee.baty@laird-assessors.com</item>
	   <item name="principal_reference">LBATY CRASH</item>
	   <item name="inspection_contact_phone">01513429961</item>
	  <item name="inspection_contact_mobile">01513429961</item>
	   <item name="contact_for_inspection">Lee Baty</item>
	   <item name="accident_date">01/01/2019</item>
	   <item name="area_of_damage_to_vehicle_reported">rear</item>
	   <item name="contact_for_inspection">Mr Lee Baty</item>
	   <item name="client_first_name">Lee</item>
	   <item name="client_second_name">Baty</item>
	   <item name="client_mobile">0777777777</item>
	   <item name="client_phone_notes">ring after 5pm</item>
	   <item name="vehicle_registration">&lt;registration&gt;W705ACW&lt;/registration&gt;</item>
	   <item name="agent_special_instructions">Knock around the rear</item>
	   <item name="principal_comments">Car was turning left when other car collideded in to the side</item>
	   <item name="file_referrer">Turner and Hooch</item>
	   <item name="client_salutation">Mr</item>
	   <item name="inspection_address"><postcode>PO1 0DE</postcode><first_line>First line</first_line><second_line>Second Line</second_line><town>TOWN</town><county>COUNTY</county></item>
   </data>
</request>
```

### Health Check

This will confirm the laird api is up and running.

Returning healthy or unhealthy depending on status.

```
http://localhost:5050/live/health
```

```
http://localhost:5050/test/health
```

