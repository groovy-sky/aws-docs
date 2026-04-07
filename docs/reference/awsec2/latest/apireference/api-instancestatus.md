# InstanceStatus

Describes the status of an instance.

## Contents

**attachedEbsStatus**

Reports impaired functionality that stems from an attached Amazon EBS volume that is
unreachable and unable to complete I/O operations.

Type: [EbsStatusSummary](api-ebsstatussummary.md) object

Required: No

**availabilityZone**

The Availability Zone of the instance.

Type: String

Required: No

**availabilityZoneId**

The ID of the Availability Zone of the instance.

Type: String

Required: No

**EventsSet.N**

Any scheduled events associated with the instance.

Type: Array of [InstanceStatusEvent](api-instancestatusevent.md) objects

Required: No

**instanceId**

The ID of the instance.

Type: String

Required: No

**instanceState**

The intended state of the instance. [DescribeInstanceStatus](api-describeinstancestatus.md) requires
that an instance be in the `running` state.

Type: [InstanceState](api-instancestate.md) object

Required: No

**instanceStatus**

Reports impaired functionality that stems from issues internal to the instance, such
as impaired reachability.

Type: [InstanceStatusSummary](api-instancestatussummary.md) object

Required: No

**operator**

The service provider that manages the instance.

Type: [OperatorResponse](api-operatorresponse.md) object

Required: No

**outpostArn**

The Amazon Resource Name (ARN) of the Outpost.

Type: String

Required: No

**systemStatus**

Reports impaired functionality that stems from issues related to the systems that
support an instance, such as hardware failures and network connectivity problems.

Type: [InstanceStatusSummary](api-instancestatussummary.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/instancestatus.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/instancestatus.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/instancestatus.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

InstanceStateChange

InstanceStatusDetails
