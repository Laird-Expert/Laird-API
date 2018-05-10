# Workflow Questions JSON

Once you have the workflow ID of the report your wish to instruct on you will then need to work out what details are accepted for each workflow.

This will list all the parameters that are expected to be submitted.

You can leave the value of the parameters blank but every detail helps.

*  This is a **GET** request

Example Request
------

```
GET /api/v2/YOURAPIKEY/workflow/70.json HTTP/1.1
User-Agent: Laird Assessors API Example
Accept: */*
Host: test2-lairdassessors.swiftcase.co.uk
Accept-Encoding: gzip, deflate
Connection: close
```

Example Response
--------

```
{
    "workflow_questions": [
        {
            "question_key": "internal_reference", 
            "name": "Internal Reference", 
            "id": 6
        }, 
        {
            "question_key": "principal_reference", 
            "name": "Principal Reference", 
            "id": 9
        }, 
        {
            "question_key": "accident_date", 
            "name": "Accident Date", 
            "id": 13
        }, 
        {
            "question_key": "instruction_date", 
            "name": "Instruction Date Override", 
            "id": 18
        }, 
        {
            "question_key": "client_type", 
            "name": "Client Type", 
            "id": 32
        }, 
        {
            "question_key": "inspection_contact_phone", 
            "name": "Inspection Contact Phone", 
            "id": 39
        }, 
        {
            "question_key": "inspection_contact_mobile", 
            "name": "Inspection Contact Mobile", 
            "id": 40
        }, 
        {
            "question_key": "agent_special_instructions", 
            "name": "Expert Instructions", 
            "id": 41
        }, 
        {
            "question_key": "principal_comments", 
            "name": "Principal Comments", 
            "id": 92
        }, 
        {
            "question_key": "repairing_garage_name", 
            "name": "Repairing Garage Name", 
            "id": 93
        }, 
        {
            "question_key": "repairing_garage_phone_number", 
            "name": "Repairing Garage Phone Number", 
            "id": 94
        }, 
        {
            "question_key": "area_of_damage_to_vehicle_reported", 
            "name": "Area of Damage to Vehicle", 
            "id": 95
        }, 
        {
            "question_key": "approximate_mileage", 
            "name": "Approximate Mileage", 
            "id": 96
        }, 
        {
            "question_key": "approximate_mileage_units", 
            "name": "Approximate Mileage Units", 
            "id": 97
        }, 
        {
            "question_key": "agent_important_special_instructions", 
            "name": "Important Appointment Information for Expert", 
            "id": 101
        }, 
        {
            "question_key": "authorised", 
            "name": "Authorised", 
            "id": 102
        }, 
        {
            "question_key": "labour_rate_agreed_with_principal", 
            "name": "Labour Rate Agreed with Principal", 
            "id": 103
        }, 
        {
            "question_key": "vat_registered", 
            "name": "VAT Registered", 
            "id": 105
        }, 
        {
            "question_key": "file_referrer", 
            "name": "File Referrer", 
            "id": 178
        }, 
        {
            "question_key": "agreed_terms", 
            "name": "Agreed Terms", 
            "id": 257
        }, 
        {
            "question_key": "client_salutation", 
            "name": "Client Salutation", 
            "id": 45
        }, 
        {
            "question_key": "client_first_name", 
            "name": "Client First Name", 
            "id": 46
        }, 
        {
            "question_key": "client_second_name", 
            "name": "Client Second Name", 
            "id": 47
        }, 
        {
            "question_key": "client_postcode", 
            "name": "Client Address", 
            "id": 53
        }, 
        {
            "question_key": "client_email_address", 
            "name": "Client Email Address", 
            "id": 54
        }, 
        {
            "question_key": "client_daytime_phone", 
            "name": "Client Daytime Phone", 
            "id": 55
        }, 
        {
            "question_key": "client_evening_phone", 
            "name": "Client Evening Phone", 
            "id": 56
        }, 
        {
            "question_key": "client_mobile", 
            "name": "Client Mobile", 
            "id": 57
        }, 
        {
            "question_key": "client_phone_notes", 
            "name": "Client Phone Notes", 
            "id": 58
        }, 
        {
            "question_key": "contact_for_inspection", 
            "name": "Contact Name for Inspection", 
            "id": 33
        }, 
        {
            "question_key": "inspection_address", 
            "name": "Inspection Address", 
            "id": 38
        }, 
        {
            "question_key": "vehicle_registration", 
            "name": "Vehicle Registration", 
            "id": 63
        }, 
        {
            "question_key": "on_site_in_use", 
            "name": "On Site/In Use", 
            "id": 91
        }, 
        {
            "question_key": "glass_evaluator_code", 
            "name": "Glass Evaluator Code", 
            "id": 66
        }, 
        {
            "question_key": "vehicle_extras", 
            "name": "Vehicle Extras", 
            "id": 72
        }, 
        {
            "question_key": "vehicle_condition", 
            "name": "Vehicle Condition", 
            "id": 106
        }, 
        {
            "question_key": "unrelated_damage", 
            "name": "Unrelated Damage", 
            "id": 107
        }, 
        {
            "question_key": "confirmed_mileage", 
            "name": "Confirmed Mileage", 
            "id": 111
        }, 
        {
            "question_key": "confirmed_mileage_units", 
            "name": "Confirmed Mileage Units", 
            "id": 112
        }, 
        {
            "question_key": "reading_provided_by", 
            "name": "Reading Provided By", 
            "id": 113
        }, 
        {
            "question_key": "road_tax", 
            "name": "Road Tax", 
            "id": 114
        }, 
        {
            "question_key": "value_of_vehicle", 
            "name": "Value of Vehicle", 
            "id": 133
        }, 
        {
            "question_key": "salvage_value", 
            "name": "Salvage Value", 
            "id": 134
        }, 
        {
            "question_key": "salvage_category", 
            "name": "Salvage Category", 
            "id": 135
        }, 
        {
            "question_key": "agreed_labour_rate", 
            "name": "Agreed Labour Rate", 
            "id": 138
        }, 
        {
            "question_key": "trade_value", 
            "name": "Trade Value", 
            "id": 176
        }, 
        {
            "question_key": "retail_value", 
            "name": "Retail Value", 
            "id": 177
        }, 
        {
            "question_key": "impact_diagram", 
            "name": "Impact Diagram", 
            "id": 2
        }, 
        {
            "question_key": "area_of_damage_to_vehicle_assessed", 
            "name": "Area Of Damage to Vehicle", 
            "id": 116
        }, 
        {
            "question_key": "offside_rear_tyre_depth", 
            "name": "Offside Rear Tyre Depth", 
            "id": 118
        }, 
        {
            "question_key": "offside_front_tyre_depth", 
            "name": "Offside Front Tyre Depth", 
            "id": 120
        }, 
        {
            "question_key": "near_side_front_tyre_depth", 
            "name": "Near Side Front Tyre Depth", 
            "id": 121
        }, 
        {
            "question_key": "near_side_rear_tyre_depth", 
            "name": "Near Side Rear Tyre Depth", 
            "id": 122
        }, 
        {
            "question_key": "agreed_authorised_by", 
            "name": "Agreed/Authorised By", 
            "id": 132
        }, 
        {
            "question_key": "reserve_price", 
            "name": "Reserve Price", 
            "id": 136
        }, 
        {
            "question_key": "brakes_working", 
            "name": "Brakes Working", 
            "id": 126
        }, 
        {
            "question_key": "steering_working", 
            "name": "Steering Working", 
            "id": 127
        }, 
        {
            "question_key": "air_bag_deployed", 
            "name": "Air Bag Deployed", 
            "id": 128
        }, 
        {
            "question_key": "vehicle_comments", 
            "name": "Vehicle Comments", 
            "id": 108
        }, 
        {
            "question_key": "savings", 
            "name": "Savings", 
            "id": 109
        }, 
        {
            "question_key": "repair_comments", 
            "name": "Repair Comments", 
            "id": 110
        }, 
        {
            "question_key": "savings_notes", 
            "name": "Savings Notes", 
            "id": 139
        }, 
        {
            "question_key": "repairs_principal_labour_rate", 
            "name": "Repairs - Principal Labour Rate", 
            "id": 140
        }, 
        {
            "question_key": "repairs_hours_to_repair", 
            "name": "Repairs - Hours to Repair", 
            "id": 141
        }, 
        {
            "question_key": "repairs_labour_rate", 
            "name": "Repairs - Labour Rate", 
            "id": 142
        }, 
        {
            "question_key": "repairs_total_labour_cost", 
            "name": "Repairs - Total Labour Cost", 
            "id": 143
        }, 
        {
            "question_key": "repairs_total_parts_cost", 
            "name": "Repairs - Total Parts Cost", 
            "id": 144
        }, 
        {
            "question_key": "repairs_total_paints_cost", 
            "name": "Repairs - Total Paints Cost", 
            "id": 145
        }, 
        {
            "question_key": "repairs_total_other_costs", 
            "name": "Repairs - Total Other Costs", 
            "id": 146
        }, 
        {
            "question_key": "repairs_sub_total", 
            "name": "Repairs - Sub Total", 
            "id": 147
        }, 
        {
            "question_key": "repairs_total_vat", 
            "name": "Repairs - Total VAT", 
            "id": 148
        }, 
        {
            "question_key": "repairs_grand_total", 
            "name": "Repairs - Grand Total", 
            "id": 149
        }, 
        {
            "question_key": "new_parts_required", 
            "name": "New Parts Required", 
            "id": 160
        }, 
        {
            "question_key": "repaired_parts", 
            "name": "Repaired Parts", 
            "id": 161
        }, 
        {
            "question_key": "additional_repairs_and_information", 
            "name": "Additional Repairs and Information", 
            "id": 162
        }, 
        {
            "question_key": "recovery_costs", 
            "name": "Recovery Costs", 
            "id": 163
        }, 
        {
            "question_key": "storage_costs", 
            "name": "Storage Costs", 
            "id": 164
        }, 
        {
            "question_key": "gearbox", 
            "name": "Gearbox", 
            "id": 214
        }, 
        {
            "question_key": "temporary_repairs", 
            "name": "Temporary Repairs", 
            "id": 217
        }, 
        {
            "question_key": "estimate", 
            "name": "Estimate", 
            "id": 137
        }, 
        {
            "question_key": "impact_magnitude", 
            "name": "Impact Magnitude", 
            "id": 117
        }, 
        {
            "question_key": "roadworthy", 
            "name": "Roadworthy", 
            "id": 130
        }, 
        {
            "question_key": "repairable_total_loss", 
            "name": "Repairable/Total Loss", 
            "id": 131
        }, 
        {
            "question_key": "repairs_original_hours_to_repair", 
            "name": "Repairs - Original Hours to Repair", 
            "id": 150
        }, 
        {
            "question_key": "repairs_original_labour_rate", 
            "name": "Repairs - Original Labour Rate", 
            "id": 151
        }, 
        {
            "question_key": "repairs_original_total_labour_cost", 
            "name": "Repairs - Original Total Labour Cost", 
            "id": 152
        }, 
        {
            "question_key": "repairs_original_total_parts_cost", 
            "name": "Repairs - Original Total Parts Cost", 
            "id": 153
        }, 
        {
            "question_key": "repairs_original_total_paints_cost", 
            "name": "Repairs - Original Total Paints Cost", 
            "id": 154
        }, 
        {
            "question_key": "repairs_original_total_other_costs", 
            "name": "Repairs - Original Total Other Costs", 
            "id": 155
        }, 
        {
            "question_key": "repairs_original_sub_total", 
            "name": "Repairs - Original Sub Total", 
            "id": 156
        }, 
        {
            "question_key": "repairs_original_total_vat", 
            "name": "Repairs - Original Total VAT", 
            "id": 157
        }, 
        {
            "question_key": "repairs_original_grand_total", 
            "name": "Repairs - Original Grand Total", 
            "id": 158
        }, 
        {
            "question_key": "report_date", 
            "name": "Report Date Override", 
            "id": 180
        }, 
        {
            "question_key": "appointment_date_time", 
            "name": "Appointment Date & Time", 
            "id": 24
        }
    ]
}


```

