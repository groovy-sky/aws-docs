# MoveAddressToVpc

###### Note

This action is deprecated.

Moves an Elastic IP address from the EC2-Classic platform to the EC2-VPC platform. The
Elastic IP address must be allocated to your account for more than 24 hours, and it must not
be associated with an instance. After the Elastic IP address is moved, it is no longer
available for use in the EC2-Classic platform. You cannot move an Elastic IP address that was
originally allocated for use in the EC2-VPC platform to the EC2-Classic platform.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**PublicIp**

The Elastic IP address.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**allocationId**

The allocation ID for the Elastic IP address.

Type: String

**requestId**

The ID of the request.

Type: String

**status**

The status of the move of the IP address.

Type: String

Valid Values: `MoveInProgress | InVpc | InClassic`

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/MoveAddressToVpc)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/MoveAddressToVpc)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/moveaddresstovpc.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/moveaddresstovpc.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/moveaddresstovpc.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/moveaddresstovpc.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/moveaddresstovpc.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/moveaddresstovpc.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/MoveAddressToVpc)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/moveaddresstovpc.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

MonitorInstances

MoveByoipCidrToIpam
