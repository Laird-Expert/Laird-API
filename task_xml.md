# Create a Task XML

Once you have the workflow type id and you have got all the questions that this workflow requires you can submit a post request like the following and it will return with your task id & internal reference (Laird Reference).

Notes
----
1. The returned id is the task ID

2. Internal Rerfence is the Laird Assessors Rerference.

3. Dates are in UTC format yyyy-mm-dd

4. Workflow ID is required and is the type of report you wish us to do.

5. Task User is the Laird Assessors member of staff the case has been automatically sent to. (This may change depending on staff avaliablity)

6. Status is us acknowledging we have rececived the instruction.

*  This is a **POST** request

Example Request
------

```
POST /api/v2/YOURAPIKEY/task.xml HTTP/1.1
Host: test2-lairdassessors.swiftcase.co.uk
Content-Type: text/xml
Cache-Control: no-cache


<?xml version="1.0" encoding="UTF-8"?>
<request>
   <workflow_type_id>70</workflow_type_id>
   <data>
	   <item name="client_staff_email">test@test.com</item>
	   <item name="principal_reference">YOUR REFERENCE/item>
	   <item name="inspection_contact_phone">inspection-phone</item>
	   <item name="contact_for_inspection">inspection-contact-name</item>
	   <item name="accident_date">2018-04-01</item>
	   <item name="area_of_damage_to_vehicle_reported">area-of-damage</item>
   </data>
</request>
```

Example Response
--------

```
<task>
    <id>700110</id>
    <data>
        <item name="principal_reference">reference</item>
        <item name="inspection_contact_phone">inspection-phone</item>
        <item name="accident_date"></item>
        <item name="area_of_damage_to_vehicle_reported">area-of-damage</item>
        <item name="contact_for_inspection">inspection-contact-name</item>
        <item name="internal_reference">18-700106</item>
    </data>
    <fees>
        <net>100.00</net>
        <vat>100.00</vat>
        <gross>100.00</gross>
    </fees>
    <status>
        <status>Instruction Received</status>
    </status>
    <taskUsers>
        <taskUser>
            <user>
                <id>1401</id>
                <name>Your Company Name</name>
            </user>
            <relationship>Client</relationship>
        </taskUser>
        <taskUser>
            <user>
                <id>1401</id>
                <name>Your Company Name</name>
            </user>
            <relationship>Client staff</relationship>
        </taskUser>
        <taskUser>
            <user>
                <id>6001885</id>
                <name>Expert, Laird</name>
            </user>
            <relationship>Owner</relationship>
        </taskUser>
    </taskUsers>
</task>

```
