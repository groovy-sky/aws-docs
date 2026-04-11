# InstanceStorageInfo

Describes the instance store features that are supported by the instance type.

## Contents

**Disks.N**

Describes the disks that are available for the instance type.

Type: Array of [DiskInfo](api-diskinfo.md) objects

Required: No

**encryptionSupport**

Indicates whether data is encrypted at rest.

Type: String

Valid Values: `unsupported | required`

Required: No

**nvmeSupport**

Indicates whether non-volatile memory express (NVMe) is supported.

Type: String

Valid Values: `unsupported | supported | required`

Required: No

**totalSizeInGB**

The total size of the disks, in GB.

Type: Long

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/instancestorageinfo.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/instancestorageinfo.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/instancestorageinfo.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

InstanceStatusSummary

InstanceTagNotificationAttribute
