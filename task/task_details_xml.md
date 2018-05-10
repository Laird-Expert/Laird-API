
# task_details XML

This is like on v1 API get-case-details where it returns all the details about a case

*  This is a **GET** request

You need to submit the following:
---

*  API Key
*  TaskID



Example Request
-----

```
GET /api/v2/YOURAPIKEY/task/TASKID.xml HTTP/1.1
Host: test2-lairdassessors.swiftcase.co.uk
Cache-Control: no-cache
```

Example Response
-----

```
<?xml version="1.0" encoding="UTF-8"?>
<task>
    <id>332821</id>
    <data>
        <item name="principal_reference">1354</item>
        <item name="inspection_contact_phone">073455445</item>
        <item name="contact_for_inspection">yrasdsa</item>
        <item name="inspection_address">L1 0BG</item>
        <item name="accident_date"></item>
        <item name="area_of_damage_to_vehicle_reported">rear off side door has been smashed, off side rear lights smashed, hinge pushed into framework</item>
        <item name="vat_registered">No</item>
        <item name="principal_comments">Preferred Labour Rate: 56222.00 Our clients vehicle was parked and unattended, your insured has collided with our clients vehicle while trying to drive past</item>
        <item name="agent_special_instructions">&lt;p&gt;please contact client to arrange inspection, clients work address -&lt;/p&gt;</item>
        <item name="repairing_garage_name">some garage L1 0BG</item>
        <item name="repairing_garage_phone_number">01123123132123</item>
        <item name="agreed_terms">No</item>
        <item name="client_salutation">Mr</item>
        <item name="client_first_name">Matthew</item>
        <item name="client_second_name">Smith</item>
        <item name="client_postcode">L1 0BG</item>
        <item name="client_email_address">client@gmail.com</item>
        <item name="client_evening_phone">012312312</item>
        <item name="client_mobile">0712312313</item>
        <item name="vehicle_registration">L1 0BG</item>
        <item name="internal_reference">14-522205</item>
        <item name="instruction_date"></item>
        <item name="client_type"></item>
        <item name="inspection_contact_mobile"></item>
        <item name="approximate_mileage"></item>
        <item name="approximate_mileage_units"></item>
        <item name="agent_important_special_instructions"></item>
        <item name="authorised"></item>
        <item name="labour_rate_agreed_with_principal"></item>
        <item name="file_referrer"></item>
        <item name="appointment_date_time">2018-05-17</item>
        <item name="client_daytime_phone"></item>
        <item name="client_phone_notes"></item>
        <item name="on_site_in_use"></item>
        <item name="repairable_total_loss">Repairable</item>
        <item name="roadworthy">Unroadworthy</item>
        <item name="impact_magnitude">Moderate</item>
        <item name="agreed_labour_rate"></item>
        <item name="confirmed_mileage">6630</item>
        <item name="confirmed_mileage_units">Miles</item>
        <item name="reading_provided_by">Viewed by the engineer</item>
        <item name="glass_evaluator_code">K9RX</item>
        <item name="retail_value">17180.00</item>
        <item name="trade_value">15180.00</item>
        <item name="value_of_vehicle">17180.00</item>
        <item name="salvage_value"></item>
        <item name="salvage_category">N/A</item>
        <item name="vehicle_condition">Very Good</item>
        <item name="unrelated_damage"></item>
        <item name="vehicle_extras">Security locks, Ply lining.</item>
        <item name="road_tax">March 2018</item>
        <item name="agreed_authorised_by">Without Prejudice</item>
        <item name="reserve_price">3249.13</item>
        <item name="area_of_damage_to_vehicle_assessed">Right hand rear</item>
        <item name="near_side_rear_tyre_depth">7.6</item>
        <item name="near_side_front_tyre_depth">6.6</item>
        <item name="offside_rear_tyre_depth">7.4</item>
        <item name="offside_front_tyre_depth">6.4</item>
        <item name="brakes_working">In Order</item>
        <item name="steering_working">In Order</item>
        <item name="air_bag_deployed">No</item>
        <item name="gearbox">Manual</item>
        <item name="recovery_costs"></item>
        <item name="storage_costs"></item>
        <item name="temporary_repairs"></item>
        <item name="vehicle_comments">The vehicle is still in use with the owner, however, the right hand back door is damaged and difficult to close.</item>
        <item name="repairs_hours_to_repair">30.5</item>
        <item name="repairs_labour_rate">45.00</item>
        <item name="repairs_total_labour_cost">1372.50</item>
        <item name="repairs_total_parts_cost">628.58</item>
        <item name="repairs_total_paints_cost">546.28</item>
        <item name="repairs_total_other_costs">135.50</item>
        <item name="repairs_sub_total">2682.86</item>
        <item name="repairs_total_vat">536.57</item>
        <item name="repairs_grand_total">3219.43</item>
        <item name="repairs_principal_labour_rate"></item>
        <item name="new_parts_required">Right hand back door&#13;
Right hand back door badges&#13;
Right hand lower back door hinge&#13;
Right hand rear lamp&#13;
sundries</item>
        <item name="repair_comments"></item>
        <item name="repaired_parts">Right hand rear body panel&#13;
Blend left hand back door&#13;
Blend back door handle&#13;
Clamp &amp; pull&#13;
Dress the sill clamp marks&#13;
Right hand outer "D" pillar&#13;
Right hand inner "D" pillar&#13;
Rear lower panel&#13;
Rear lower panel&#13;
Right hand rear bumper section&#13;
Drill holes &amp; refit security lock&#13;
Final check &amp; road test&#13;
Repair method research&#13;
De-nib &amp; polish&#13;
Apply corrosion protection materials&#13;
Shut down vehicle&#13;
Wash &amp; clean &#13;
Remove &amp; refit ply lining</item>
        <item name="additional_repairs_and_information">E P A charge&#13;
Corrosion protection materials&#13;
Car care kit&#13;
Pre repair diagnostic checks&#13;
Post diagnostic checks</item>
        <item name="savings"></item>
        <item name="savings_notes"></item>
        <item name="repairs_original_hours_to_repair"></item>
        <item name="repairs_original_labour_rate"></item>
        <item name="repairs_original_total_labour_cost"></item>
        <item name="repairs_original_total_parts_cost"></item>
        <item name="repairs_original_total_paints_cost"></item>
        <item name="repairs_original_total_other_costs"></item>
        <item name="repairs_original_sub_total"></item>
        <item name="repairs_original_total_vat"></item>
        <item name="repairs_original_grand_total"></item>
        <item name="estimate"></item>
    </data>
    <fees>
        <net>803.00</net>
        <vat>1336.00</vat>
        <gross>396.00</gross>
    </fees>
    <status>
        <status>Instruction Received</status>
    </status>
    <taskUsers>
        <taskUser>
            <user>
                <id>14021</id>
                <name>Management Company</name>
            </user>
            <relationship>Client</relationship>
        </taskUser>
        <taskUser>
            <user>
                <id>255419</id>
                <name>Reports, Engineer</name>
            </user>
            <relationship>Client staff</relationship>
        </taskUser>
        <taskUser>
            <user>
                <id>6006576</id>
                <name>Larkin, Sarah</name>
            </user>
            <relationship>Previously Allocated Owner</relationship>
        </taskUser>
        <taskUser>
            <user>
                <id>6000260</id>
                <name>Smith, Sean</name>
            </user>
            <relationship>Previously allocated supplier</relationship>
        </taskUser>
        <taskUser>
            <user>
                <id>6006576</id>
                <name>Larkin, Sarah</name>
            </user>
            <relationship>Previously Allocated Owner</relationship>
        </taskUser>
        <taskUser>
            <user>
                <id>1000155</id>
                <name>Peek, Alexandra</name>
            </user>
            <relationship>Previously Allocated Owner</relationship>
        </taskUser>
        <taskUser>
            <user>
                <id>6001885</id>
                <name>Adams, Liam</name>
            </user>
            <relationship>Previously Allocated Owner</relationship>
        </taskUser>
        <taskUser>
            <user>
                <id>6000390</id>
                <name>Williams, David</name>
            </user>
            <relationship>Previously allocated supplier</relationship>
        </taskUser>
        <taskUser>
            <user>
                <id>6000390</id>
                <name>Williams, David</name>
            </user>
            <relationship>Previously allocated supplier</relationship>
        </taskUser>
        <taskUser>
            <user>
                <id>6000390</id>
                <name>Williams, David</name>
            </user>
            <relationship>Previously allocated supplier</relationship>
        </taskUser>
        <taskUser>
            <user>
                <id>6000390</id>
                <name>Williams, David</name>
            </user>
            <relationship>Previously allocated supplier</relationship>
        </taskUser>
        <taskUser>
            <user>
                <id>6001946</id>
                <name>Sharples, Hannah</name>
            </user>
            <relationship>Previously Allocated Owner</relationship>
        </taskUser>
        <taskUser>
            <user>
                <id>6000386</id>
                <name>Goodhead, Julian</name>
            </user>
            <relationship>Previously allocated supplier</relationship>
        </taskUser>
        <taskUser>
            <user>
                <id>6006887</id>
                <name>Markey, Sarah</name>
            </user>
            <relationship>Previously Allocated Owner</relationship>
        </taskUser>
        <taskUser>
            <user>
                <id>6000386</id>
                <name>Goodhead, Julian</name>
            </user>
            <relationship>Supplier</relationship>
        </taskUser>
        <taskUser>
            <user>
                <id>6001946</id>
                <name>Sharples, Hannah</name>
            </user>
            <relationship>Owner</relationship>
        </taskUser>
    </taskUsers>
</task>
```

Notes
----

Internal Reference is the Laird Assessors Reference.


