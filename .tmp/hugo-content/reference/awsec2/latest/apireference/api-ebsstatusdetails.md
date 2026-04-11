# EbsStatusDetails

Describes the attached EBS status check for an instance.

## Contents

**impairedSince**

The date and time when the attached EBS status check failed.

Type: Timestamp

Required: No

**name**

The name of the attached EBS status check.

Type: String

Valid Values: `reachability`

Required: No

**status**

The result of the attached EBS status check.

Type: String

Valid Values: `passed | failed | insufficient-data | initializing`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/ebsstatusdetails.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/ebsstatusdetails.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/ebsstatusdetails.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

EbsOptimizedInfo

EbsStatusSummary
