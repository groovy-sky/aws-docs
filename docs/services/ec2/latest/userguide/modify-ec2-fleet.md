# Modify an EC2 Fleet

You can modify the total target capacity, Spot capacity, and On-Demand capacity of an
EC2 Fleet. You can also modify whether running instances should be terminated if the new
total target capacity is reduced below the current size of the fleet.

## Considerations

Consider the following when modifying an EC2 Fleet:

- **Fleet type** – You can only modify
an EC2 Fleet of type `maintain`. You can't modify an EC2 Fleet of type
`request` or `instant`.

- **Fleet parameters** – You can modify
the following parameters of an EC2 Fleet:

- `target-capacity-specification` – Increase or
decrease the target capacity for:

- `TotalTargetCapacity`

- `OnDemandTargetCapacity`

- `SpotTargetCapacity`

- `excess-capacity-termination-policy` – Whether
running instances should be terminated if the total target capacity
of the EC2 Fleet is decreased below the current size of the fleet. Valid
values are:

- `no-termination`

- `termination`

- **Fleet behavior when increasing total target**
**capacity** – When you increase the total target
capacity, the EC2 Fleet launches the additional instances according to the
instance purchasing option specified for
`DefaultTargetCapacityType`, which is either On-Demand Instances or Spot Instances,
and according to the specified [allocation strategy](ec2-fleet-allocation-strategy.md).

- **Fleet behavior when decreasing Spot target capacity**
– When you decrease the Spot target capacity, the EC2 Fleet deletes any
open requests that exceed the new target capacity. You can request that the
fleet terminate Spot Instances until the size of the fleet reaches the new target
capacity. When an EC2 Fleet terminates a Spot Instance because the target capacity was
decreased, the instance receives a Spot Instance interruption notice.

Instances are selected for termination based on the allocation
strategy:

- `capacity-optimized` – Terminates instances from pools with the least
available capacity.

- `price-capacity-optimized` – Uses a combination of price and available
capacity: terminates instances from pools with the least available
capacity and which are the highest-priced among these pools.

- `diversified` – Terminates instances across all pools.

- `lowest-price` – Terminates instances from highest-priced pools.

Alternatively, you can request that EC2 Fleet keep the fleet at its current size, but not
replace any Spot Instances that are interrupted or that you terminate
manually.

- **Fleet state** – You can modify an
EC2 Fleet that is in the `submitted` or `active` state.
When you modify a fleet, it enters the `modifying` state.

## Commands for modifying an EC2 Fleet

AWS CLI

###### To modify the total target capacity of an EC2 Fleet

Use the [modify-fleet](../../../cli/latest/reference/ec2/modify-fleet.md)
command.

```nohighlight

aws ec2 modify-fleet \
    --fleet-id fleet-73fbd2ce-aa30-494c-8788-1cee4EXAMPLE \
    --target-capacity-specification TotalTargetCapacity=20
```

If you are decreasing the target capacity but want to keep the fleet at its
current size, you can modify the previous example as follows.

```nohighlight

aws ec2 modify-fleet \
    --fleet-id fleet-73fbd2ce-aa30-494c-8788-1cee4EXAMPLE \
    --target-capacity-specification TotalTargetCapacity=10 \
    --excess-capacity-termination-policy no-termination
```

PowerShell

###### To modify the total target capacity of an EC2 Fleet

Use the [Edit-EC2Fleet](../../../powershell/latest/reference/items/edit-ec2fleet.md)
cmdlet.

```powershell

Edit-EC2Fleet `
    -FleetId "fleet-73fbd2ce-aa30-494c-8788-1cee4EXAMPLE" `
    -TargetCapacitySpecification_TotalTargetCapacity 20
```

If you are decreasing the target capacity but want to keep the fleet at its
current size, you can modify the previous example as follows.

```powershell

Edit-EC2Fleet `
    -FleetId "fleet-73fbd2ce-aa30-494c-8788-1cee4EXAMPLE" `
    -TargetCapacitySpecification_TotalTargetCapacity 20 `
    -ExcessCapacityTerminationPolicy "NoTermination"
```

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

Describe an EC2 Fleet

Delete an EC2 Fleet
