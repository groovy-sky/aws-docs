# DescribeDBClusterAutomatedBackups

Displays backups for both current and deleted DB clusters. For example, use this operation to find details
about automated backups for previously deleted clusters. Current clusters are returned for both the
`DescribeDBClusterAutomatedBackups` and `DescribeDBClusters` operations.

All parameters are optional.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**DBClusterIdentifier**

(Optional) The user-supplied DB cluster identifier. If this parameter is specified, it must
match the identifier of an existing DB cluster. It returns information from the
specific DB cluster's automated backup. This parameter isn't case-sensitive.

Type: String

Required: No

**DbClusterResourceId**

The resource ID of the DB cluster that is the source of the automated backup. This parameter isn't case-sensitive.

Type: String

Required: No

**Filters.Filter.N**

A filter that specifies which resources to return based on status.

Supported filters are the following:

- `status`

- `retained` \- Automated backups for deleted clusters and after backup replication is stopped.

- `db-cluster-id` \- Accepts DB cluster identifiers and Amazon Resource Names (ARNs).
The results list includes only information about the DB cluster automated backups identified by these ARNs.

- `db-cluster-resource-id` \- Accepts DB resource identifiers and Amazon Resource Names (ARNs).
The results list includes only information about the DB cluster resources identified by these ARNs.

Returns all resources by default. The status for each resource is specified in the response.

Type: Array of [Filter](api-filter.md) objects

Required: No

**Marker**

The pagination token provided in the previous request. If this parameter is specified the response includes only
records beyond the marker, up to `MaxRecords`.

Type: String

Required: No

**MaxRecords**

The maximum number of records to include in the response. If more records exist than the specified `MaxRecords`
value, a pagination token called a marker is included in the response so that you can retrieve the remaining results.

Type: Integer

Required: No

## Response Elements

The following elements are returned by the service.

**DBClusterAutomatedBackups.DBClusterAutomatedBackup.N**

A list of `DBClusterAutomatedBackup` backups.

Type: Array of [DBClusterAutomatedBackup](api-dbclusterautomatedbackup.md) objects

**Marker**

The pagination token provided in the previous request. If this parameter is specified the response includes only
records beyond the marker, up to `MaxRecords`.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**DBClusterAutomatedBackupNotFoundFault**

No automated backup for this DB cluster was found.

HTTP Status Code: 404

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/describedbclusterautomatedbackups.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/describedbclusterautomatedbackups.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/describedbclusterautomatedbackups.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/describedbclusterautomatedbackups.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/describedbclusterautomatedbackups.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/describedbclusterautomatedbackups.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/describedbclusterautomatedbackups.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/describedbclusterautomatedbackups.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/describedbclusterautomatedbackups.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/describedbclusterautomatedbackups.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeCertificates

DescribeDBClusterBacktracks

All content copied from https://docs.aws.amazon.com/.
