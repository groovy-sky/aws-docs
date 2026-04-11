# NetworkOriginConfiguration

The proposed `InternetConfiguration` or `VpcConfiguration` to
apply to the Amazon S3 access point. You can make the access point accessible from the internet,
or you can specify that all requests made through that access point must originate from a
specific virtual private cloud (VPC). You can specify only one type of network
configuration. For more information, see [Creating access\
points](../../../../services/s3/latest/dev/creating-access-points.md).

## Contents

###### Important

This data type is a UNION, so only one of the following members can be specified when used or returned.

**internetConfiguration**

The configuration for the Amazon S3 access point or multi-region access point with an
`Internet` origin.

Type: [InternetConfiguration](api-internetconfiguration.md) object

Required: No

**vpcConfiguration**

The proposed virtual private cloud (VPC) configuration for the Amazon S3 access point. VPC
configuration does not apply to multi-region access points. For more information, see
[VpcConfiguration](../../../../services/s3/latest/api/api-control-vpcconfiguration.md).

Type: [VpcConfiguration](api-vpcconfiguration.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/accessanalyzer-2019-11-01/networkoriginconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/accessanalyzer-2019-11-01/networkoriginconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/accessanalyzer-2019-11-01/networkoriginconfiguration.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

Location

PathElement
