# DeprovisionIpamByoasn

Deprovisions your Autonomous System Number (ASN) from your AWS account. This action can only be called after any BYOIP CIDR associations are removed from your AWS account with [DisassociateIpamByoasn](api-disassociateipambyoasn.md).
For more information, see [Tutorial: Bring your ASN to IPAM](../../../../services/vpc/latest/ipam/tutorials-byoasn.md) in the _Amazon VPC IPAM guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**Asn**

An ASN.

Type: String

Required: Yes

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**IpamId**

The IPAM ID.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**byoasn**

An ASN and BYOIP CIDR association.

Type: [Byoasn](api-byoasn.md) object

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DeprovisionIpamByoasn)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DeprovisionIpamByoasn)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/deprovisionipambyoasn.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/deprovisionipambyoasn.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/deprovisionipambyoasn.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/deprovisionipambyoasn.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/deprovisionipambyoasn.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/deprovisionipambyoasn.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DeprovisionIpamByoasn)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/deprovisionipambyoasn.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DeprovisionByoipCidr

DeprovisionIpamPoolCidr
