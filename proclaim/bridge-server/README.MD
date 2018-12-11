Bridge Server
---

Due to some issues with some clients having SSL errors I have built a small server that will act like the bridge on your server for you.

I have attached some pre-compiled binarys but if you have go installed you can compile it your self with go build.

You can run the binary and add --port 5555 if you wish to start the bridge server on another port other than 5050


### Endpoints

---


### Example for Engineer Availability


```
https://localhost:5050/test/engineer?APIKEY=YOURAPIKEY&TASK=YOURTASKID&IDATE=YOURDATE&maxresult=15
```

```
https://localhost:5050/live/engineer?APIKEY=YOURAPIKEY&TASK=YOURTASKID&IDATE=YOURDATE&maxresult=15
```

```

### Example for Engineer Booking


```
https://localhost:5050/test/bookengineer?APIKEY=YOURAPIKEY&TASK=YOURTASKID&IDATE=YOURDATE&BOOKID=1512312
```

```
https://localhost:5050/live/bookengineer?APIKEY=YOURAPIKEY&TASK=YOURTASKID&IDATE=YOURDATE&BOOKID=1231231
```

```




### Example of Workflow Questions


```
https://localhost:5050/test/workflow.php?APIKEY=YOURAPIKEY&ID=WORKFLOWID
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
