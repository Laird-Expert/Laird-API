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
Host: 127.0.0.1
Content-Type: text/xml
cache-control: no-cache
Content-Length: 768

<?xml version="1.0" encoding="UTF-8"?>
<request>
   <workflow_type_id>210</workflow_type_id>
   <data>
	   <item name="principal_reference"></item>
	   <item name="internal_reference"></item>
	   <item name="vehicle_registration"></item>
	   <item name="glass_evaluator_code"></item>
	   <item name="vehicle_comments"></item>
	   <item name="confirmed_mileage"></item>
	   <item name="confirmed_mileage_units"></item>
	   <item name="road_tax"></item>
	   <item name="repairable_total_loss"></item>
	   <item name="value_of_vehicle"></item>
	   <item name="salvage_value"></item>
	   <item name="trade_value"></item>
	   <item name="retail_value"></item>
	   <item name="gearbox"></item>
	   <item name="listed_extras"></item>
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
   <workflow_type_id>210</workflow_type_id>
   <data>
	   <item name="principal_reference"></item>
	   <item name="internal_reference"></item>
	   <item name="vehicle_registration"></item>
	   <item name="glass_evaluator_code"></item>
	   <item name="vehicle_comments"></item>
	   <item name="confirmed_mileage"></item>
	   <item name="confirmed_mileage_units"></item>
	   <item name="road_tax"></item>
	   <item name="repairable_total_loss"></item>
	   <item name="value_of_vehicle"></item>
	   <item name="salvage_value"></item>
	   <item name="trade_value"></item>
	   <item name="retail_value"></item>
	   <item name="gearbox"></item>
	   <item name="listed_extras"></item>
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

