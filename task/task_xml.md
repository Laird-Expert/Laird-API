# Create a Task XML

Once you have the workflow type id and you have got all the questions that this workflow requires you can submit a post request like the following and it will return with your task id & internal reference (Laird Reference).

Notes
----
1. The returned id is the task ID

2. Internal Rerfence is the Laird Assessors Rerference.

3. Dates should be submitted d/m/Y (21/5/2018)

4. Workflow ID is required and is the type of report you wish us to do.

5. Task User is the Laird Assessors member of staff the case has been automatically sent to. (This may change depending on staff avaliablity)

6. Status is us acknowledging we have rececived the instruction.

*  This is a **POST** request

Example Request
------

```
POST /api/v2/YOURAPIKEY/task.xml HTTP/1.1
Host: test-lairdassessors.swiftcase.co.uk
Content-Type: text/xml
Cache-Control: no-cache


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
	   <item name="client_postcode">&lt;first_line&gt;First line&lt;/first_line&gt;&lt;second_line&gt;Second Line&lt;/second_line&gt;&lt;town&gt;TOWN&lt;/town&gt;&lt;county&gt;COUNTY&lt;/county&gt;</item>
	   <item name="client_phone_notes">ring after 5pm</item>
	   <item name="vehicle_registration">&lt;registration&gt;W705ACW&lt;/registration&gt;</item>
	   <item name="agent_special_instructions">Knock around the rear</item>
	   <item name="principal_comments">Car was turning left when other car collideded in to the side</item>
	   <item name="file_referrer">Turner and Hooch</item>
	   <item name="client_salutation">Mr</item>
	   <item name="inspection_address">&lt;postcode&gt;PO1 0DE&lt;/postcode&gt;&lt;first_line&gt;First line&lt;/first_line&gt;&lt;second_line&gt;Second Line&lt;/second_line&gt;&lt;town&gt;TOWN&lt;/town&gt;&lt;county&gt;COUNTY&lt;/county&gt;</item>
   </data>
</request>
```

Example Response
--------

```
<?xml version="1.0" encoding="UTF-8"?>
<task>
    <id>659748</id>
    <data>
        <item name="principal_reference">LBATY CRASH</item>
        <item name="accident_date">2019-01-01</item>
        <item name="area_of_damage_to_vehicle_reported">rear</item>
        <item name="agent_special_instructions">Knock around the rear</item>
        <item name="principal_comments">Car was turning left when other car collideded in to the side</item>
        <item name="file_referrer">Turner and Hooch</item>
        <item name="inspection_contact_phone">01513429961</item>
        <item name="contact_for_inspection">Mr Lee Baty</item>
        <item name="vehicle_registration">W705ACW</item>
        <item name="inspection_address">PO1 0DE</item>
        <item name="client_first_name">Lee</item>
        <item name="client_second_name">Baty</item>
        <item name="client_mobile">0777777777</item>
        <item name="client_phone_notes">ring after 5pm</item>
        <item name="client_salutation">Mr</item>
        <item name="internal_reference">19-626131</item>
    </data>
    <fees>
        <net>0.00</net>
        <vat>0.00</vat>
        <gross>0.00</gross>
    </fees>
    <status>
        <status>Instruction Received</status>
    </status>
    <taskUsers>
        <taskUser>
            <user>
                <id>6000343</id>
                <name>1 High Priority Client Testing Company TEST</name>
            </user>
            <relationship>Client</relationship>
        </taskUser>
        <taskUser>
            <user>
                <id>6000343</id>
                <name>1 High Priority Client Testing Company TEST</name>
            </user>
            <relationship>Client staff</relationship>
        </taskUser>
        <taskUser>
            <user>
                <id>6001885</id>
                <name>Adams, Liam</name>
            </user>
            <relationship>Owner</relationship>
        </taskUser>
    </taskUsers>
</task>

```
