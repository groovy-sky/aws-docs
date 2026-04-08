# DescribeDBShardGroups

Describes existing Aurora Limitless Database DB shard groups.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**DBShardGroupIdentifier**

The user-supplied DB shard group identifier. If this parameter is specified, information for only the specific DB shard group is returned.
This parameter isn't case-sensitive.

Constraints:

- If supplied, must match an existing DB shard group identifier.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 63.

Pattern: `[a-zA-Z](?:-?[a-zA-Z0-9]+)*`

Required: No

**Filters.Filter.N**

A filter that specifies one or more DB shard groups to describe.

Type: Array of [Filter](api-filter.md) objects

Required: No

**Marker**

An optional pagination token provided by a previous `DescribeDBShardGroups` request. If this parameter is
specified, the response includes only records beyond the marker, up to the value specified by `MaxRecords`.

Type: String

Required: No

**MaxRecords**

The maximum number of records to include in the response. If more records exist than the specified `MaxRecords`
value, a pagination token called a marker is included in the response so you can retrieve the remaining results.

Default: 100

Constraints: Minimum 20, maximum 100

Type: Integer

Valid Range: Minimum value of 20. Maximum value of 100.

Required: No

## Response Elements

The following elements are returned by the service.

**DBShardGroups.DBShardGroup.N**

Contains a list of DB shard groups for the user.

Type: Array of [DBShardGroup](api-dbshardgroup.md) objects

**Marker**

A pagination token that can be used in a later `DescribeDBClusters` request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**DBClusterNotFoundFault**

`DBClusterIdentifier` doesn't refer to an existing DB cluster.

HTTP Status Code: 404

**DBShardGroupNotFound**

The specified DB shard group name wasn't found.

HTTP Status Code: 404

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/describedbshardgroups.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/describedbshardgroups.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/describedbshardgroups.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/describedbshardgroups.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/describedbshardgroups.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/describedbshardgroups.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/describedbshardgroups.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/describedbshardgroups.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/describedbshardgroups.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/describedbshardgroups.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeDBSecurityGroups

DescribeDBSnapshotAttributes

All content copied from https://docs.aws.amazon.com/.
