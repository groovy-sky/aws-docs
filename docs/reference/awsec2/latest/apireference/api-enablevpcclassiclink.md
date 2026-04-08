# EnableVpcClassicLink

###### Note

This action is deprecated.

Enables a VPC for ClassicLink. You can then link EC2-Classic instances to your
ClassicLink-enabled VPC to allow communication over private IP addresses. You cannot
enable your VPC for ClassicLink if any of your VPC route tables have existing routes for
address ranges within the `10.0.0.0/8` IP address range, excluding local
routes for VPCs in the `10.0.0.0/16` and `10.1.0.0/16` IP address
ranges.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**VpcId**

The ID of the VPC.

Type: String

Required: Yes

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/enablevpcclassiclink.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/enablevpcclassiclink.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/enablevpcclassiclink.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/enablevpcclassiclink.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/enablevpcclassiclink.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/enablevpcclassiclink.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/enablevpcclassiclink.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/enablevpcclassiclink.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/enablevpcclassiclink.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/enablevpcclassiclink.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

EnableVolumeIO

EnableVpcClassicLinkDnsSupport
