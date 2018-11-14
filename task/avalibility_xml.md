# availability XML

availability is where you would  pick a date that we have for the post code you have provided for the case and book a engineer for a physical inspection

*  This is a **GET** request

You need to submit the following:
---

*  API Key
*  TaskID
*  DATEFROM is a UTC date which is normally the next working day i.e 2018-05-05
*  MAX results is how many dates you would like to be returned the max is 15.


Example Request
-----

```
GET /api/v2/YOURAPIKEY/task/TASKID/availability/DATEFROM/MAXRESULTS.xml HTTP/1.1
Host: test-lairdassessors.swiftcase.co.uk
Content-Type: application/x-www-form-urlencoded
Cache-Control: no-cache
```

Example Response
-----

```
<?xml version="1.0" encoding="UTF-8"?>
<resourceAvailability>
    <id>29991</id>
    <availability>
        <date>2018-05-11</date>
    </availability>
    <availability>
        <date>2018-05-14</date>
    </availability>
    <availability>
        <date>2018-05-15</date>
    </availability>
    <availability>
        <date>2018-05-16</date>
    </availability>
    <availability>
        <date>2018-05-17</date>
    </availability>
    <availability>
        <date>2018-05-18</date>
    </availability>
    <availability>
        <date>2018-05-21</date>
    </availability>
    <availability>
        <date>2018-05-22</date>
    </availability>
    <availability>
        <date>2018-05-23</date>
    </availability>
    <availability>
        <date>2018-05-24</date>
    </availability>
</resourceAvailability>
```

Notes
----

Your next request will require the ID number that has been returned and the date you wish to use.

so in this example the ID is 29991 and the next date avaliable would be 2018-05-11

If no dates are returned then we might not have an engineer who has set avalaiblity so this might need a call to our logistics team to check.
