# MacModificationTask

Information about a System Integrity Protection (SIP) modification task or volume
ownership delegation task for an Amazon EC2 Mac instance.

## Contents

**instanceId**

The ID of the Amazon EC2 Mac instance.

Type: String

Required: No

**macModificationTaskId**

The ID of task.

Type: String

Required: No

**macSystemIntegrityProtectionConfig**

\[SIP modification tasks only\] Information about the SIP
configuration.

Type: [MacSystemIntegrityProtectionConfiguration](api-macsystemintegrityprotectionconfiguration.md) object

Required: No

**startTime**

The date and time the task was created, in the UTC timezone
( `YYYY-MM-DDThh:mm:ss.sssZ`).

Type: Timestamp

Required: No

**TagSet.N**

The tags assigned to the task.

Type: Array of [Tag](api-tag.md) objects

Required: No

**taskState**

The state of the task.

Type: String

Valid Values: `successful | failed | in-progress | pending`

Required: No

**taskType**

The type of task.

Type: String

Valid Values: `sip-modification | volume-ownership-delegation`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/macmodificationtask.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/macmodificationtask.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/macmodificationtask.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

MacHost

MacSystemIntegrityProtectionConfiguration
