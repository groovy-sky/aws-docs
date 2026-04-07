# InstanceStatus

Describes the status of an instance.

## Contents

**attachedEbsStatus**

Reports impaired functionality that stems from an attached Amazon EBS volume that is
unreachable and unable to complete I/O operations.

Type: [EbsStatusSummary](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_EbsStatusSummary.html) object

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

Type: Array of [InstanceStatusEvent](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_InstanceStatusEvent.html) objects

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

Type: [InstanceStatusSummary](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_InstanceStatusSummary.html) object

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

Type: [InstanceStatusSummary](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_InstanceStatusSummary.html) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/InstanceStatus)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/InstanceStatus)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/InstanceStatus)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

InstanceStateChange

InstanceStatusDetails
