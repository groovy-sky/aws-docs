# ProvisionIpamByoasn

Provisions your Autonomous System Number (ASN) for use in your AWS account. This action requires authorization context for Amazon to bring the ASN to an AWS account. For more information, see [Tutorial: Bring your ASN to IPAM](https://docs.aws.amazon.com/vpc/latest/ipam/tutorials-byoasn.html) in the _Amazon VPC IPAM guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**Asn**

A public 2-byte or 4-byte ASN.

Type: String

Required: Yes

**AsnAuthorizationContext**

An ASN authorization context.

Type: [AsnAuthorizationContext](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_AsnAuthorizationContext.html) object

Required: Yes

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**IpamId**

An IPAM ID.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**byoasn**

An ASN and BYOIP CIDR association.

Type: [Byoasn](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Byoasn.html) object

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/ProvisionIpamByoasn)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/ProvisionIpamByoasn)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/ProvisionIpamByoasn)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ec2-2016-11-15/ProvisionIpamByoasn)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/ProvisionIpamByoasn)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ec2-2016-11-15/ProvisionIpamByoasn)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ec2-2016-11-15/ProvisionIpamByoasn)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ec2-2016-11-15/ProvisionIpamByoasn)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/ProvisionIpamByoasn)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/ProvisionIpamByoasn)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ProvisionByoipCidr

ProvisionIpamPoolCidr
