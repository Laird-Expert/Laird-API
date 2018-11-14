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


### Example for the Booking of an Engineer


```
https://bridges.laird-assessors.com/test/engineer.php?APIKEY=YOURAPIKEY&TASK=YOURTASKID&IDATE=YOURDATE&maxresult=15
```



### Example of Workflow Questions


```
https://bridges.laird-assessors.com/test/workflowsq.php?APIKEY=YOURAPIKEY&ID=WORKFLOWID
```


### Example of Status


```
https://bridges.laird-assessors.com/test/status.php?APIKEY=YOURAPIKEY&TASK=TASKID
```
