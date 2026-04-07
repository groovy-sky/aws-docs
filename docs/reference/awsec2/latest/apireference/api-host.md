# Host

Describes the properties of the Dedicated Host.

## Contents

**allocationTime**

The time that the Dedicated Host was allocated.

Type: Timestamp

Required: No

**allowsMultipleInstanceTypes**

Indicates whether the Dedicated Host supports multiple instance types of the same
instance family. If the value is `on`, the Dedicated Host supports multiple
instance types in the instance family. If the value is `off`, the Dedicated
Host supports a single instance type only.

Type: String

Valid Values: `on | off`

Required: No

**assetId**

The ID of the Outpost hardware asset on which the Dedicated Host is allocated.

Type: String

Required: No

**autoPlacement**

Whether auto-placement is on or off.

Type: String

Valid Values: `on | off`

Required: No

**availabilityZone**

The Availability Zone of the Dedicated Host.

Type: String

Required: No

**availabilityZoneId**

The ID of the Availability Zone in which the Dedicated Host is allocated.

Type: String

Required: No

**availableCapacity**

Information about the instances running on the Dedicated Host.

Type: [AvailableCapacity](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_AvailableCapacity.html) object

Required: No

**clientToken**

Unique, case-sensitive identifier that you provide to ensure the idempotency of the request. For more information, see [Ensuring Idempotency](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html).

Type: String

Required: No

**hostId**

The ID of the Dedicated Host.

Type: String

Required: No

**hostMaintenance**

Indicates whether host maintenance is enabled or disabled for the Dedicated
Host.

Type: String

Valid Values: `on | off`

Required: No

**hostProperties**

The hardware specifications of the Dedicated Host.

Type: [HostProperties](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_HostProperties.html) object

Required: No

**hostRecovery**

Indicates whether host recovery is enabled or disabled for the Dedicated Host.

Type: String

Valid Values: `on | off`

Required: No

**hostReservationId**

The reservation ID of the Dedicated Host. This returns a `null` response if
the Dedicated Host doesn't have an associated reservation.

Type: String

Required: No

**Instances.N**

The IDs and instance type that are currently running on the Dedicated Host.

Type: Array of [HostInstance](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_HostInstance.html) objects

Required: No

**memberOfServiceLinkedResourceGroup**

Indicates whether the Dedicated Host is in a host resource group. If **memberOfServiceLinkedResourceGroup** is `true`, the
host is in a host resource group; otherwise, it is not.

Type: Boolean

Required: No

**outpostArn**

The Amazon Resource Name (ARN) of the AWS Outpost on which the
Dedicated Host is allocated.

Type: String

Required: No

**ownerId**

The ID of the AWS account that owns the Dedicated Host.

Type: String

Required: No

**releaseTime**

The time that the Dedicated Host was released.

Type: Timestamp

Required: No

**state**

The Dedicated Host's state.

Type: String

Valid Values: `available | under-assessment | permanent-failure | released | released-permanent-failure | pending`

Required: No

**TagSet.N**

Any tags assigned to the Dedicated Host.

Type: Array of [Tag](api-tag.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/Host)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/Host)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/Host)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

HistoryRecordEntry

HostInstance
