# VolumeStatusDetails

Describes a volume status.

## Contents

**name**

The name of the volume status.

- `io-enabled` \- Indicates the volume I/O status. For more
information, see [Amazon EBS volume \
status checks](../../../../services/ebs/latest/userguide/monitoring-volume-checks.md).

- `io-performance` \- Indicates the volume performance status.
For more information, see [Amazon EBS volume \
status checks](../../../../services/ebs/latest/userguide/monitoring-volume-checks.md).

- `initialization-state` \- Indicates the status of the volume
initialization process. For more information, see [Initialize Amazon EBS volumes](../../../../services/ebs/latest/userguide/initalize-volume.md).

Type: String

Valid Values: `io-enabled | io-performance | initialization-state`

Required: No

**status**

The intended status of the volume status.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/volumestatusdetails.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/volumestatusdetails.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/volumestatusdetails.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

VolumeStatusAttachmentStatus

VolumeStatusEvent
