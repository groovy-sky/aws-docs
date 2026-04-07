# DescribeMovingAddresses

###### Note

This action is deprecated.

Describes your Elastic IP addresses that are being moved from or being restored to the EC2-Classic platform.
This request does not return information about any other Elastic IP addresses in your account.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**Filter.N**

One or more filters.

- `moving-status` \- The status of the Elastic IP address
( `MovingToVpc` \| `RestoringToClassic`).

Type: Array of [Filter](api-filter.md) objects

Required: No

**MaxResults**

The maximum number of results to return for the request in a single page. The remaining
results of the initial request can be seen by sending another request with the returned
`NextToken` value. This value can be between 5 and 1000; if
`MaxResults` is given a value outside of this range, an error is returned.

Default: If no value is provided, the default is 1000.

Type: Integer

Valid Range: Minimum value of 5. Maximum value of 1000.

Required: No

**NextToken**

The token for the next page of results.

Type: String

Required: No

**PublicIp.N**

One or more Elastic IP addresses.

Type: Array of strings

Required: No

## Response Elements

The following elements are returned by the service.

**movingAddressStatusSet**

The status for each Elastic IP address.

Type: Array of [MovingAddressStatus](api-movingaddressstatus.md) objects

**nextToken**

The token to use to retrieve the next page of results. This value is `null` when there are no more results to return.

Type: String

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DescribeMovingAddresses)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DescribeMovingAddresses)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describemovingaddresses.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describemovingaddresses.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describemovingaddresses.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describemovingaddresses.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describemovingaddresses.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describemovingaddresses.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DescribeMovingAddresses)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describemovingaddresses.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeManagedPrefixLists

DescribeNatGateways
