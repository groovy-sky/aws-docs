# DescribeDBClusterBacktracks

Returns information about backtracks for a DB cluster.

For more information on Amazon Aurora, see
[What is Amazon Aurora?](../../../../services/amazonrds/latest/aurorauserguide/chap-auroraoverview.md) in the _Amazon Aurora User Guide_.

###### Note

This action only applies to Aurora MySQL DB clusters.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**DBClusterIdentifier**

The DB cluster identifier of the DB cluster to be described. This parameter is
stored as a lowercase string.

Constraints:

- Must contain from 1 to 63 alphanumeric characters or hyphens.

- First character must be a letter.

- Can't end with a hyphen or contain two consecutive hyphens.

Example: `my-cluster1`

Type: String

Required: Yes

**BacktrackIdentifier**

If specified, this value is the backtrack identifier of the backtrack to be
described.

Constraints:

- Must contain a valid universally unique identifier (UUID). For more information about UUIDs, see
[Universally unique \
identifier](https://en.wikipedia.org/wiki/Universally_unique_identifier).

Example: `123e4567-e89b-12d3-a456-426655440000`

Type: String

Required: No

**Filters.Filter.N**

A filter that specifies one or more DB clusters to describe. Supported filters
include the following:

- `db-cluster-backtrack-id` \- Accepts backtrack identifiers. The
results list includes information about only the backtracks identified by these
identifiers.

- `db-cluster-backtrack-status` \- Accepts any of the following backtrack status values:

- `applying`

- `completed`

- `failed`

- `pending`

The results list includes information about only the backtracks identified
by these values.

Type: Array of [Filter](api-filter.md) objects

Required: No

**Marker**

An optional pagination token provided by a previous
`DescribeDBClusterBacktracks` request.
If this parameter is specified, the response includes
only records beyond the marker,
up to the value specified by `MaxRecords`.

Type: String

Required: No

**MaxRecords**

The maximum number of records to include in the response.
If more records exist than the specified `MaxRecords` value,
a pagination token called a marker is included in the response so you can retrieve the remaining results.

Default: 100

Constraints: Minimum 20, maximum 100.

Type: Integer

Required: No

## Response Elements

The following elements are returned by the service.

**DBClusterBacktracks.DBClusterBacktrack.N**

Contains a list of backtracks for the user.

Type: Array of [DBClusterBacktrack](api-dbclusterbacktrack.md) objects

**Marker**

A pagination token that can be used in a later `DescribeDBClusterBacktracks` request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**DBClusterBacktrackNotFoundFault**

`BacktrackIdentifier` doesn't refer to an existing backtrack.

HTTP Status Code: 404

**DBClusterNotFoundFault**

`DBClusterIdentifier` doesn't refer to an existing DB cluster.

HTTP Status Code: 404

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/describedbclusterbacktracks.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/describedbclusterbacktracks.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/describedbclusterbacktracks.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/describedbclusterbacktracks.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/describedbclusterbacktracks.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/describedbclusterbacktracks.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/describedbclusterbacktracks.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/describedbclusterbacktracks.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/describedbclusterbacktracks.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/describedbclusterbacktracks.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeDBClusterAutomatedBackups

DescribeDBClusterEndpoints

All content copied from https://docs.aws.amazon.com/.
