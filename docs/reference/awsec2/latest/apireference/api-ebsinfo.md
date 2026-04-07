# EbsInfo

Describes the Amazon EBS features supported by the instance type.

## Contents

**attachmentLimitType**

Indicates whether the instance type features a shared or dedicated Amazon EBS
volume attachment limit. For more information, see [Amazon EBS volume limits for \
Amazon EC2 instances](../../../../services/ec2/latest/userguide/volume-limits.md) in the _Amazon EC2 User Guide_.

Type: String

Valid Values: `shared | dedicated`

Required: No

**EbsCardSet.N**

Describes the EBS cards available for the instance type.

Type: Array of [EbsCardInfo](api-ebscardinfo.md) objects

Required: No

**ebsOptimizedInfo**

Describes the optimized EBS performance for the instance type.

Type: [EbsOptimizedInfo](api-ebsoptimizedinfo.md) object

Required: No

**ebsOptimizedSupport**

Indicates whether the instance type is Amazon EBS-optimized. For more information, see [Amazon EBS-optimized\
instances](../../../../services/ec2/latest/userguide/ebsoptimized.md) in _Amazon EC2 User Guide_.

Type: String

Valid Values: `unsupported | supported | default`

Required: No

**encryptionSupport**

Indicates whether Amazon EBS encryption is supported.

Type: String

Valid Values: `unsupported | supported`

Required: No

**maximumEbsAttachments**

Indicates the maximum number of Amazon EBS volumes that can be attached to
the instance type. For more information, see [Amazon EBS volume limits for \
Amazon EC2 instances](../../../../services/ec2/latest/userguide/volume-limits.md) in the _Amazon EC2 User Guide_.

Type: Integer

Required: No

**maximumEbsCards**

Indicates the number of EBS cards supported by the instance type.

Type: Integer

Required: No

**nvmeSupport**

Indicates whether non-volatile memory express (NVMe) is supported.

Type: String

Valid Values: `unsupported | supported | required`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/ebsinfo.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/ebsinfo.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/ebsinfo.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

EbsCardInfo

EbsInstanceBlockDevice
