# EbsStatusSummary

Provides a summary of the attached EBS volume status for an instance.

## Contents

**Details.N**

Details about the attached EBS status check for an instance.

Type: Array of [EbsStatusDetails](api-ebsstatusdetails.md) objects

Required: No

**status**

The current status.

Type: String

Valid Values: `ok | impaired | insufficient-data | not-applicable | initializing`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/ebsstatussummary.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/ebsstatussummary.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/ebsstatussummary.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

EbsStatusDetails

Ec2InstanceConnectEndpoint
