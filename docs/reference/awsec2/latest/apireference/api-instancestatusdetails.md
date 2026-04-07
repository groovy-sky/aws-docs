# InstanceStatusDetails

Describes the instance status.

## Contents

**impairedSince**

The time when a status check failed. For an instance that was launched and impaired,
this is the time when the instance was launched.

Type: Timestamp

Required: No

**name**

The type of instance status.

Type: String

Valid Values: `reachability`

Required: No

**status**

The status.

Type: String

Valid Values: `passed | failed | insufficient-data | initializing`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/instancestatusdetails.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/instancestatusdetails.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/instancestatusdetails.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

InstanceStatus

InstanceStatusEvent
