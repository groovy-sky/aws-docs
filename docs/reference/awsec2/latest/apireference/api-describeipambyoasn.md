# DescribeIpamByoasn

Describes your Autonomous System Numbers (ASNs), their provisioning statuses, and the BYOIP CIDRs with which they are associated. For more information, see [Tutorial: Bring your ASN to IPAM](../../../../services/vpc/latest/ipam/tutorials-byoasn.md) in the _Amazon VPC IPAM guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**MaxResults**

The maximum number of results to return with a single call.
To retrieve the remaining results, make another call with the returned `nextToken` value.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 100.

Required: No

**NextToken**

The token for the next page of results.

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**byoasnSet**

ASN and BYOIP CIDR associations.

Type: Array of [Byoasn](api-byoasn.md) objects

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/describeipambyoasn.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/describeipambyoasn.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describeipambyoasn.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describeipambyoasn.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describeipambyoasn.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describeipambyoasn.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describeipambyoasn.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describeipambyoasn.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/describeipambyoasn.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describeipambyoasn.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeInternetGateways

DescribeIpamExternalResourceVerificationTokens
