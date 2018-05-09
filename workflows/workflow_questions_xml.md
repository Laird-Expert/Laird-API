# Workflow Questions XML

Once you have the workflow ID of the report your wish to instruct on you will then need to work out what details are accepted for each workflow.

This will list all the parameters that are expected to be submitted.

You can leave the value of the parameters blank but ever detail helps.

*  This is a **GET** request

Example Request
------

```
GET /api/v2/YOURAPIKEY/workflow/128.xml HTTP/1.1
User-Agent: Laird Assessors API Example
Accept: */*
Host: test2-lairdassessors.swiftcase.co.uk
Accept-Encoding: gzip, deflate
Connection: close
```

Example Response
--------

```
<?xml version="1.0" encoding="UTF-8"?>
<workflowQuestionCollection>
    <workflowQuestion>
        <id>6</id>
        <name>Internal Reference</name>
        <question_key>internal_reference</question_key>
    </workflowQuestion>
    <workflowQuestion>
        <id>9</id>
        <name>Principal Reference</name>
        <question_key>principal_reference</question_key>
    </workflowQuestion>
    <workflowQuestion>
        <id>273</id>
        <name>Type of Document Collection</name>
        <question_key>type_of_document_collection</question_key>
    </workflowQuestion>
    <workflowQuestion>
        <id>45</id>
        <name>Client Salutation</name>
        <question_key>client_salutation</question_key>
    </workflowQuestion>
    <workflowQuestion>
        <id>46</id>
        <name>Client First Name</name>
        <question_key>client_first_name</question_key>
    </workflowQuestion>
    <workflowQuestion>
        <id>47</id>
        <name>Client Second Name</name>
        <question_key>client_second_name</question_key>
    </workflowQuestion>
    <workflowQuestion>
        <id>53</id>
        <name>Client Address</name>
        <question_key>client_postcode</question_key>
    </workflowQuestion>
    <workflowQuestion>
        <id>54</id>
        <name>Client Email Address</name>
        <question_key>client_email_address</question_key>
    </workflowQuestion>
    <workflowQuestion>
        <id>55</id>
        <name>Client Daytime Phone</name>
        <question_key>client_daytime_phone</question_key>
    </workflowQuestion>
    <workflowQuestion>
        <id>56</id>
        <name>Client Evening Phone</name>
        <question_key>client_evening_phone</question_key>
    </workflowQuestion>
    <workflowQuestion>
        <id>57</id>
        <name>Client Mobile</name>
        <question_key>client_mobile</question_key>
    </workflowQuestion>
    <workflowQuestion>
        <id>58</id>
        <name>Client Phone Notes</name>
        <question_key>client_phone_notes</question_key>
    </workflowQuestion>
    <workflowQuestion>
        <id>271</id>
        <name>Appointment Date &amp; Time</name>
        <question_key>standard_appointment_date_time</question_key>
    </workflowQuestion>
    <workflowQuestion>
        <id>272</id>
        <name>Appointment Location</name>
        <question_key>appointment_location_postcode</question_key>
    </workflowQuestion>
    <workflowQuestion>
        <id>27</id>
        <name>Nature of Case</name>
        <question_key>nature_of_case</question_key>
    </workflowQuestion>
    <workflowQuestion>
        <id>41</id>
        <name>Expert Instructions</name>
        <question_key>agent_special_instructions</question_key>
    </workflowQuestion>
    <workflowQuestion>
        <id>42</id>
        <name>Additional Information</name>
        <question_key>additional_information</question_key>
    </workflowQuestion>
    <workflowQuestion>
        <id>5</id>
        <name>Report handler</name>
        <question_key>report_handler</question_key>
    </workflowQuestion>
</workflowQuestionCollection>

```

