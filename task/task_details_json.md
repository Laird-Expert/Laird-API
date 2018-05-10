
# task_details JSON

This is like on v1 API get-case-details where it returns all the details about a case

*  This is a **GET** request

You need to submit the following:
---

*  API Key
*  TaskID



Example Request
-----

```
GET /api/v2/YOURAPIKEY/task/TASKID.json HTTP/1.1
Host: test2-lairdassessors.swiftcase.co.uk
Cache-Control: no-cache
```

Example Response
-----

```
{
    "id": 602821,
    "data": [
        {
            "name": "principal_reference",
            "value": "168dsfsdfsdf"
        },
        {
            "name": "inspection_contact_phone",
            "value": "075234234234"
        },
        {
            "name": "contact_for_inspection",
            "value": "Mr John Smith"
        },
        {
            "name": "inspection_address",
            "value": "L1 0BG"
        },
        {
            "name": "accident_date",
            "value": "2018-05-05"
        },
        {
            "name": "area_of_damage_to_vehicle_reported",
            "value": "rear off side door has been smashed, off side rear lights smashed, hinge pushed into framework"
        },
        {
            "name": "vat_registered",
            "value": "No"
        },
        {
            "name": "principal_comments",
            "value": "Preferred Labour Rate: 56.00 Our clients vehicle was parked and unattended, your insured has collided with our clients vehicle while trying to drive past"
        },
        {
            "name": "agent_special_instructions",
            "value": "<p>please contact client to arrange inspection, clients work address -</p>"
        },
        {
            "name": "repairing_garage_name",
            "value": ""
        },
        {
            "name": "repairing_garage_phone_number",
            "value": "01345345345354"
        },
        {
            "name": "agreed_terms",
            "value": "No"
        },
        {
            "name": "client_salutation",
            "value": "Mr"
        },
        {
            "name": "client_first_name",
            "value": "John"
        },
        {
            "name": "client_second_name",
            "value": "Smith"
        },
        {
            "name": "client_postcode",
            "value": "L1 0BG"
        },
        {
            "name": "client_email_address",
            "value": "client@gmsil.com"
        },
        {
            "name": "client_evening_phone",
            "value": "077865867509"
        },
        {
            "name": "client_mobile",
            "value": "07548457864876"
        },
        {
            "name": "vehicle_registration",
            "value": "W705ACW"
        },
        {
            "name": "internal_reference",
            "value": "17-23423425"
        },
        {
            "name": "instruction_date",
            "value": ""
        },
        {
            "name": "client_type",
            "value": ""
        },
        {
            "name": "inspection_contact_mobile",
            "value": ""
        },
        {
            "name": "approximate_mileage",
            "value": ""
        },
        {
            "name": "approximate_mileage_units",
            "value": ""
        },
        {
            "name": "agent_important_special_instructions",
            "value": ""
        },
        {
            "name": "authorised",
            "value": ""
        },
        {
            "name": "labour_rate_agreed_with_principal",
            "value": ""
        },
        {
            "name": "file_referrer",
            "value": ""
        },
        {
            "name": "appointment_date_time",
            "value": "2018-05-17"
        },
        {
            "name": "client_daytime_phone",
            "value": ""
        },
        {
            "name": "client_phone_notes",
            "value": ""
        },
        {
            "name": "on_site_in_use",
            "value": ""
        },
        {
            "name": "repairable_total_loss",
            "value": "Repairable"
        },
        {
            "name": "roadworthy",
            "value": "Unroadworthy"
        },
        {
            "name": "impact_magnitude",
            "value": "Moderate"
        },
        {
            "name": "agreed_labour_rate",
            "value": ""
        },
        {
            "name": "confirmed_mileage",
            "value": "6630"
        },
        {
            "name": "confirmed_mileage_units",
            "value": "Miles"
        },
        {
            "name": "reading_provided_by",
            "value": "Viewed by the engineer"
        },
        {
            "name": "glass_evaluator_code",
            "value": "K9RX"
        },
        {
            "name": "retail_value",
            "value": "17180.00"
        },
        {
            "name": "trade_value",
            "value": "15180.00"
        },
        {
            "name": "value_of_vehicle",
            "value": "17180.00"
        },
        {
            "name": "salvage_value",
            "value": ""
        },
        {
            "name": "salvage_category",
            "value": "N/A"
        },
        {
            "name": "vehicle_condition",
            "value": "Very Good"
        },
        {
            "name": "unrelated_damage",
            "value": ""
        },
        {
            "name": "vehicle_extras",
            "value": "Security locks, Ply lining."
        },
        {
            "name": "road_tax",
            "value": "March 2018"
        },
        {
            "name": "agreed_authorised_by",
            "value": "Without Prejudice"
        },
        {
            "name": "reserve_price",
            "value": "3249.13"
        },
        {
            "name": "area_of_damage_to_vehicle_assessed",
            "value": "Right hand rear"
        },
        {
            "name": "near_side_rear_tyre_depth",
            "value": "7.6"
        },
        {
            "name": "near_side_front_tyre_depth",
            "value": "6.6"
        },
        {
            "name": "offside_rear_tyre_depth",
            "value": "7.4"
        },
        {
            "name": "offside_front_tyre_depth",
            "value": "6.4"
        },
        {
            "name": "brakes_working",
            "value": "In Order"
        },
        {
            "name": "steering_working",
            "value": "In Order"
        },
        {
            "name": "air_bag_deployed",
            "value": "No"
        },
        {
            "name": "gearbox",
            "value": "Manual"
        },
        {
            "name": "recovery_costs",
            "value": ""
        },
        {
            "name": "storage_costs",
            "value": ""
        },
        {
            "name": "temporary_repairs",
            "value": ""
        },
        {
            "name": "vehicle_comments",
            "value": "The vehicle is still in use with the owner, however, the right hand back door is damaged and difficult to close."
        },
        {
            "name": "repairs_hours_to_repair",
            "value": "30.5"
        },
        {
            "name": "repairs_labour_rate",
            "value": "45.00"
        },
        {
            "name": "repairs_total_labour_cost",
            "value": "1372.50"
        },
        {
            "name": "repairs_total_parts_cost",
            "value": "628.58"
        },
        {
            "name": "repairs_total_paints_cost",
            "value": "546.28"
        },
        {
            "name": "repairs_total_other_costs",
            "value": "135.50"
        },
        {
            "name": "repairs_sub_total",
            "value": "2682.86"
        },
        {
            "name": "repairs_total_vat",
            "value": "536.57"
        },
        {
            "name": "repairs_grand_total",
            "value": "3219.43"
        },
        {
            "name": "repairs_principal_labour_rate",
            "value": ""
        },
        {
            "name": "new_parts_required",
            "value": "Right hand back door\r\nRight hand back door badges\r\nRight hand lower back door hinge\r\nRight hand rear lamp\r\nsundries"
        },
        {
            "name": "repair_comments",
            "value": ""
        },
        {
            "name": "repaired_parts",
            "value": "Right hand rear body panel\r\nBlend left hand back door\r\nBlend back door handle\r\nClamp & pull\r\nDress the sill clamp marks\r\nRight hand outer \"D\" pillar\r\nRight hand inner \"D\" pillar\r\nRear lower panel\r\nRear lower panel\r\nRight hand rear bumper section\r\nDrill holes & refit security lock\r\nFinal check & road test\r\nRepair method research\r\nDe-nib & polish\r\nApply corrosion protection materials\r\nShut down vehicle\r\nWash & clean \r\nRemove & refit ply lining"
        },
        {
            "name": "additional_repairs_and_information",
            "value": "E P A charge\r\nCorrosion protection materials\r\nCar care kit\r\nPre repair diagnostic checks\r\nPost diagnostic checks"
        },
        {
            "name": "savings",
            "value": ""
        },
        {
            "name": "savings_notes",
            "value": ""
        },
        {
            "name": "repairs_original_hours_to_repair",
            "value": ""
        },
        {
            "name": "repairs_original_labour_rate",
            "value": ""
        },
        {
            "name": "repairs_original_total_labour_cost",
            "value": ""
        },
        {
            "name": "repairs_original_total_parts_cost",
            "value": ""
        },
        {
            "name": "repairs_original_total_paints_cost",
            "value": ""
        },
        {
            "name": "repairs_original_total_other_costs",
            "value": ""
        },
        {
            "name": "repairs_original_sub_total",
            "value": ""
        },
        {
            "name": "repairs_original_total_vat",
            "value": ""
        },
        {
            "name": "repairs_original_grand_total",
            "value": ""
        },
        {
            "name": "estimate",
            "value": ""
        }
    ],
    "fees": {
        "net": 80,
        "vat": 16,
        "gross": 96
    },
    "status": {
        "status": "Instruction Received"
    },
    "taskUsers": [
        {
            "user": {
                "id": 1401,
                "name": "CLaims Systems"
            },
            "relationship": "Client"
        },
        {
            "user": {
                "id": 2519,
                "name": "Reports, Engineer"
            },
            "relationship": "Client staff"
        },
        {
            "user": {
                "id": 6006576,
                "name": "Larkin, Sarah"
            },
            "relationship": "Previously Allocated Owner"
        },
        {
            "user": {
                "id": 6000260,
                "name": "Jones, Sean"
            },
            "relationship": "Previously allocated supplier"
        },
        {
            "user": {
                "id": 6006576,
                "name": "Larkin, Sarah"
            },
            "relationship": "Previously Allocated Owner"
        },
        {
            "user": {
                "id": 1000155,
                "name": "Peek, Alexandra"
            },
            "relationship": "Previously Allocated Owner"
        },
        {
            "user": {
                "id": 6001885,
                "name": "Adams, Liam"
            },
            "relationship": "Previously Allocated Owner"
        },
        {
            "user": {
                "id": 6000390,
                "name": "Williams, David"
            },
            "relationship": "Previously allocated supplier"
        },
        {
            "user": {
                "id": 6000390,
                "name": "Williams, David"
            },
            "relationship": "Previously allocated supplier"
        },
        {
            "user": {
                "id": 6000390,
                "name": "Williams, David"
            },
            "relationship": "Previously allocated supplier"
        },
        {
            "user": {
                "id": 6000390,
                "name": "Williams, David"
            },
            "relationship": "Previously allocated supplier"
        },
        {
            "user": {
                "id": 6001946,
                "name": "Sharples, Hannah"
            },
            "relationship": "Previously Allocated Owner"
        },
        {
            "user": {
                "id": 6000386,
                "name": "Goodhead, Julian"
            },
            "relationship": "Previously allocated supplier"
        },
        {
            "user": {
                "id": 6006887,
                "name": "Markey, Sarah"
            },
            "relationship": "Previously Allocated Owner"
        },
        {
            "user": {
                "id": 6000386,
                "name": "Goodhead, Julian"
            },
            "relationship": "Supplier"
        },
        {
            "user": {
                "id": 6001946,
                "name": "Sharples, Hannah"
            },
            "relationship": "Owner"
        }
    ]
}
```

Notes
----

Internal Reference is the Laird Assessors Reference.


