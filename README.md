# w6_go_2

examples :

## Create Appointment
    curl -X POST http://localhost:5455/appointments/      -H "Content-Type: application/json"      -d '{
           "pname": "Dipansu",
           "date": "2025-10-05",
           "description": "General check-up"
         }'

## Update Appointment
    Specify the id in the url which needs to be modified/updated
    curl -X PUT http://localhost:5455/appointments/2      -H "Content-Type: application/json"      -d '{
        "pname": "Dipansu",
        "date": "2024-10-05",
        "description": "dental check-up"
        }'
## Get All the Appointments

    curl -X ALL http://localhost:5455/appointments/

## Get a specific Appointmnet
    To get 2 nd appointment
    curl -X ALL http://localhost:5455/appointments/2
## Delete a specific Appointment
    To delete 2nd appointment
    curl -X DELETE  http://localhost:5455/appointments/2