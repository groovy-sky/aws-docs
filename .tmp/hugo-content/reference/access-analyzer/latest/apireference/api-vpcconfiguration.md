# VpcConfiguration

The proposed virtual private cloud (VPC) configuration for the Amazon S3 access point. VPC
configuration does not apply to multi-region access points. For more information, see
[VpcConfiguration](../../../../services/s3/latest/api/api-control-vpcconfiguration.md).

## Contents

**vpcId**

If this field is specified, this access point will only allow connections from the
specified VPC ID.

Type: String

Pattern: `vpc-([0-9a-f]){8}(([0-9a-f]){9})?`

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/accessanalyzer-2019-11-01/vpcconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/accessanalyzer-2019-11-01/vpcconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/accessanalyzer-2019-11-01/vpcconfiguration.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ValidationExceptionField

Common Parameters
