# Cancel a Capacity Reservation Fleet

When you no longer need a Capacity Reservation Fleet and the capacity it reserves, you can cancel
it. When you cancel a Fleet, its status changes to `cancelled` and it can
no longer create new Capacity Reservations. Additionally, all of the individual Capacity Reservations in the Fleet
are canceled. The instances that were previously running in the reserved capacity
continue to run normally in the shared capacity.

AWS CLI

###### To cancel a Capacity Reservation Fleet

Use the [cancel-capacity-reservation-fleets](../../../cli/latest/reference/ec2/cancel-capacity-reservation-fleets.md) command.

```nohighlight

aws ec2 cancel-capacity-reservation-fleets \
    --capacity-reservation-fleet-ids crf-abcdef01234567890
```

The following is example output.

```json

{
    "SuccessfulFleetCancellations": [
        {
            "CurrentFleetState": "cancelling",
            "PreviousFleetState": "active",
            "CapacityReservationFleetId": "crf-abcdef01234567890"
        }
    ],
    "FailedFleetCancellations": []
}
```

PowerShell

###### To cancel a Capacity Reservation Fleet

Use the [Stop-EC2CapacityReservationFleet](../../../powershell/latest/reference/items/stop-ec2capacityreservationfleet.md) cmdlet.

```powershell

Stop-EC2CapacityReservationFleet `
    -CapacityReservationFleetId crf-abcdef01234567890
```

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

Modify

Example configurations
