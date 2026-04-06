# NetworkOriginConfiguration

The proposed `InternetConfiguration` or `VpcConfiguration` to
apply to the Amazon S3 access point. You can make the access point accessible from the internet,
or you can specify that all requests made through that access point must originate from a
specific virtual private cloud (VPC). You can specify only one type of network
configuration. For more information, see [Creating access\
points](https://docs.aws.amazon.com/AmazonS3/latest/dev/creating-access-points.html).

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
[VpcConfiguration](https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_VpcConfiguration.html).

Type: [VpcConfiguration](api-vpcconfiguration.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/accessanalyzer-2019-11-01/NetworkOriginConfiguration)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/accessanalyzer-2019-11-01/NetworkOriginConfiguration)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/accessanalyzer-2019-11-01/NetworkOriginConfiguration)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Location

PathElement
