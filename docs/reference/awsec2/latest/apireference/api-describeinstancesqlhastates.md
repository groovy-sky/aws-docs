# DescribeInstanceSqlHaStates

Describes the SQL Server High Availability states for Amazon EC2 instances that are
enabled for Amazon EC2 High Availability for SQL Server monitoring.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action,
without actually making the request, and provides an error response. If you have the
required permissions, the error response is `DryRunOperation`. Otherwise,
it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**Filter.N**

One or more filters to apply to the results. Supported filters
include:

- `tag:<key>` \- The tag key and value pair assigned to the
instance. For example, to find all instances tagged with `Owner:TeamA`, specify
`tag:Owner` for the filter name and `TeamA` for the filter value.

- `tag-key` \- The tag key assigned to the instance.

- `haStatus` \- The SQL Server High Availability status of the SQL Server High Availability instance ( `processing` \|
`active` \| `standby` \| `invalid`).

- `sqlServerLicenseUsage` \- The license type for the SQL Server license
( `full` \| `waived`).

Type: Array of [Filter](api-filter.md) objects

Required: No

**InstanceId.N**

The IDs of the SQL Server High Availability instances to describe. If omitted, the API returns
SQL Server High Availability states for all SQL Server High Availability instances.

Type: Array of strings

Required: No

**MaxResults**

The maximum number of results to return for the request in a
single page. The remaining results can be seen by sending another request with the
returned `nextToken` value.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 1000.

Required: No

**NextToken**

The token to use to retrieve the next page of results.

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**instanceSet**

Information about the SQL Server High Availability instances.

Type: Array of [RegisteredInstance](api-registeredinstance.md) objects

**nextToken**

The token to use to retrieve the next page of results.
This value is `null` when there are no more results to return.

Type: String

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DescribeInstanceSqlHaStates)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DescribeInstanceSqlHaStates)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describeinstancesqlhastates.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describeinstancesqlhastates.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describeinstancesqlhastates.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describeinstancesqlhastates.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describeinstancesqlhastates.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describeinstancesqlhastates.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DescribeInstanceSqlHaStates)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describeinstancesqlhastates.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeInstanceSqlHaHistoryStates

DescribeInstanceStatus
