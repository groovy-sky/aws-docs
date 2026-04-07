# DescribeAccountAttributes

Describes attributes of your AWS account. The following are the supported account attributes:

- `default-vpc`: The ID of the default VPC for your account, or `none`.

- `max-instances`: This attribute is no longer supported. The returned
value does not reflect your actual vCPU limit for running On-Demand Instances.
For more information, see [On-Demand Instance Limits](../../../../services/ec2/latest/userguide/ec2-on-demand-instances.md#ec2-on-demand-instances-limits) in the
_Amazon Elastic Compute Cloud User Guide_.

- `max-elastic-ips`: The maximum number of Elastic IP addresses that you can allocate.

- `supported-platforms`: This attribute is deprecated.

- `vpc-max-elastic-ips`: The maximum number of Elastic IP addresses that you can allocate.

- `vpc-max-security-groups-per-interface`: The maximum number of security groups
that you can assign to a network interface.

###### Note

The order of the elements in the response, including those within nested
structures, might vary. Applications should not assume the elements appear in a
particular order.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**AttributeName.N**

The account attribute names.

Type: Array of strings

Valid Values: `supported-platforms | default-vpc`

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

## Response Elements

The following elements are returned by the service.

**accountAttributeSet**

Information about the account attributes.

Type: Array of [AccountAttribute](api-accountattribute.md) objects

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DescribeAccountAttributes)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DescribeAccountAttributes)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describeaccountattributes.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describeaccountattributes.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describeaccountattributes.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describeaccountattributes.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describeaccountattributes.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describeaccountattributes.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DescribeAccountAttributes)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describeaccountattributes.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DeregisterTransitGatewayMulticastGroupSources

DescribeAddresses
