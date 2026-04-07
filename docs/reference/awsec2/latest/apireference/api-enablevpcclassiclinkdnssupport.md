# EnableVpcClassicLinkDnsSupport

###### Note

This action is deprecated.

Enables a VPC to support DNS hostname resolution for ClassicLink. If enabled, the DNS
hostname of a linked EC2-Classic instance resolves to its private IP address when
addressed from an instance in the VPC to which it's linked. Similarly, the DNS hostname
of an instance in a VPC resolves to its private IP address when addressed from a linked
EC2-Classic instance.

You must specify a VPC ID in the request.

## Request Parameters

For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**VpcId**

The ID of the VPC.

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**return**

Returns `true` if the request succeeds; otherwise, it returns an error.

Type: Boolean

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/EnableVpcClassicLinkDnsSupport)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/EnableVpcClassicLinkDnsSupport)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/enablevpcclassiclinkdnssupport.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/enablevpcclassiclinkdnssupport.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/enablevpcclassiclinkdnssupport.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/enablevpcclassiclinkdnssupport.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/enablevpcclassiclinkdnssupport.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/enablevpcclassiclinkdnssupport.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/EnableVpcClassicLinkDnsSupport)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/enablevpcclassiclinkdnssupport.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

EnableVpcClassicLink

ExportClientVpnClientCertificateRevocationList
