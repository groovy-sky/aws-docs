# ReservedCapacityOptionsRequest

Defines EC2 Fleet preferences for utilizing reserved capacity when DefaultTargetCapacityType is set to `reserved-capacity`.

###### Note

This configuration can only be used if the EC2 Fleet is of type
`instant`.

When you specify `ReservedCapacityOptions`, you must also set
`DefaultTargetCapacityType` to `reserved-capacity` in the
`TargetCapacitySpecification`.

For more information about Interruptible Capacity Reservations, see [Launch\
instances into an Interruptible Capacity Reservation](../../../../services/ec2/latest/userguide/ec2-fleet-launch-instances-interruptible-cr-walkthrough.md) in the _Amazon EC2 User Guide_.

## Contents

**ReservationType.N**

The types of Capacity Reservations to use for fulfilling the EC2 Fleet request.

Type: Array of strings

Valid Values: `interruptible-capacity-reservation`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/reservedcapacityoptionsrequest.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/reservedcapacityoptionsrequest.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/reservedcapacityoptionsrequest.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ReservedCapacityOptions

ReservedInstanceLimitPrice
