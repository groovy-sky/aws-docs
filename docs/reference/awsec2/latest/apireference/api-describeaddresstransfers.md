# DescribeAddressTransfers

Describes an Elastic IP address transfer. For more information, see [Transfer Elastic IP addresses](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-eips.html#transfer-EIPs-intro) in the _Amazon VPC User Guide_.

When you transfer an Elastic IP address, there is a two-step handshake
between the source and transfer AWS accounts. When the source account starts the transfer,
the transfer account has seven days to accept the Elastic IP address
transfer. During those seven days, the source account can view the
pending transfer by using this action. After seven days, the
transfer expires and ownership of the Elastic IP
address returns to the source
account. Accepted transfers are visible to the source account for 14 days
after the transfers have been accepted.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**AllocationId.N**

The allocation IDs of Elastic IP addresses.

Type: Array of strings

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**MaxResults**

The maximum number of address transfers to return in one page of results.

Type: Integer

Valid Range: Minimum value of 5. Maximum value of 1000.

Required: No

**NextToken**

Specify the pagination token from a previous request to retrieve the next page of results.

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**addressTransferSet**

The Elastic IP address transfer.

Type: Array of [AddressTransfer](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_AddressTransfer.html) objects

**nextToken**

Specify the pagination token from a previous request to retrieve the next page of results.

Type: String

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DescribeAddressTransfers)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DescribeAddressTransfers)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/DescribeAddressTransfers)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ec2-2016-11-15/DescribeAddressTransfers)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/DescribeAddressTransfers)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ec2-2016-11-15/DescribeAddressTransfers)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ec2-2016-11-15/DescribeAddressTransfers)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ec2-2016-11-15/DescribeAddressTransfers)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DescribeAddressTransfers)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/DescribeAddressTransfers)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DescribeAddressesAttribute

DescribeAggregateIdFormat
