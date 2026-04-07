# DeprovisionByoipCidr

Releases the specified address range that you provisioned for use with your AWS resources
through bring your own IP addresses (BYOIP) and deletes the corresponding address pool.

Before you can release an address range, you must stop advertising it and you must not
have any IP addresses allocated from its address range.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**Cidr**

The address range, in CIDR notation. The prefix must be the same prefix
that you specified when you provisioned the address range.

Type: String

Required: Yes

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

## Response Elements

The following elements are returned by the service.

**byoipCidr**

Information about the address range.

Type: [ByoipCidr](api-byoipcidr.md) object

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DeprovisionByoipCidr)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DeprovisionByoipCidr)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/deprovisionbyoipcidr.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/deprovisionbyoipcidr.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/deprovisionbyoipcidr.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/deprovisionbyoipcidr.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/deprovisionbyoipcidr.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/deprovisionbyoipcidr.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DeprovisionByoipCidr)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/deprovisionbyoipcidr.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DeleteVpnGateway

DeprovisionIpamByoasn
