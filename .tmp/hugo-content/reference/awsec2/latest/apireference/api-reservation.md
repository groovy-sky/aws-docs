# Reservation

Describes a launch request for one or more instances, and includes owner, requester,
and security group information that applies to all instances in the launch
request.

## Contents

**GroupSet.N**

Not supported.

Type: Array of [GroupIdentifier](api-groupidentifier.md) objects

Required: No

**InstancesSet.N**

The instances.

Type: Array of [Instance](api-instance.md) objects

Required: No

**ownerId**

The ID of the AWS account that owns the reservation.

Type: String

Required: No

**requesterId**

The ID of the requester that launched the instances on your behalf (for example,
AWS Management Console or Auto Scaling).

Type: String

Required: No

**reservationId**

The ID of the reservation.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/reservation.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/reservation.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/reservation.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

RequestSpotLaunchSpecification

ReservationFleetInstanceSpecification
