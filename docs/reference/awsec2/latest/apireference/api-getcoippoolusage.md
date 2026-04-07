# GetCoipPoolUsage

Describes the allocations from the specified customer-owned address pool.

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

- `coip-address-usage.allocation-id` \- The allocation ID of the address.

- `coip-address-usage.aws-account-id` \- The ID of the AWS account that is using the customer-owned IP address.

- `coip-address-usage.aws-service` \- The AWS service that is using the customer-owned IP address.

- `coip-address-usage.co-ip` \- The customer-owned IP address.

Type: Array of [Filter](api-filter.md) objects

Required: No

**MaxResults**

The maximum number of results to return with a single call.
To retrieve the remaining results, make another call with the returned `nextToken` value.

Type: Integer

Valid Range: Minimum value of 5. Maximum value of 1000.

Required: No

**NextToken**

The token for the next page of results.

Type: String

Required: No

**PoolId**

The ID of the address pool.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**coipAddressUsageSet**

Information about the address usage.

Type: Array of [CoipAddressUsage](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CoipAddressUsage.html) objects

**coipPoolId**

The ID of the customer-owned address pool.

Type: String

**localGatewayRouteTableId**

The ID of the local gateway route table.

Type: String

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

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/GetCoipPoolUsage)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/GetCoipPoolUsage)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/GetCoipPoolUsage)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ec2-2016-11-15/GetCoipPoolUsage)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/GetCoipPoolUsage)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ec2-2016-11-15/GetCoipPoolUsage)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ec2-2016-11-15/GetCoipPoolUsage)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ec2-2016-11-15/GetCoipPoolUsage)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/GetCoipPoolUsage)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/GetCoipPoolUsage)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

GetCapacityReservationUsage

GetConsoleOutput
