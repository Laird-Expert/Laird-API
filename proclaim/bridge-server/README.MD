Bridge Server
---

Due to some issues with some clients having SSL errors I have built a small server that will act like the bridge on your server for you.

I have attached some pre-compiled binarys but if you have go installed you can compile it your self with go build.

You can run the binary and add --port 5555 if you wish to start the bridge server on another port other than 9090


### Endpoints

---


### Example for the Booking of an Engineer


```
https://localhost:9090/test/engineer?APIKEY=YOURAPIKEY&TASK=YOURTASKID&IDATE=YOURDATE&maxresult=15
```

```
https://localhost:9090/live/engineer?APIKEY=YOURAPIKEY&TASK=YOURTASKID&IDATE=YOURDATE&maxresult=15
```



### Example of Workflow Questions


```
https://localhost:9090/test/workflow.php?APIKEY=YOURAPIKEY&ID=WORKFLOWID
```

```
https://localhost:9090/live/workflow?APIKEY=YOURAPIKEY&ID=WORKFLOWID
```

### Example of Status


```
https://localhost:9090/test/status?APIKEY=YOURAPIKEY&TASK=TASKID
```

```
https://localhost:9090/live/status?APIKEY=YOURAPIKEY&TASK=TASKID
```
