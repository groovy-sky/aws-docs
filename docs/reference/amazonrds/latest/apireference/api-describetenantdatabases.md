# DescribeTenantDatabases

Describes the tenant databases in a DB instance that uses the multi-tenant
configuration. Only RDS for Oracle CDB instances are supported.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**DBInstanceIdentifier**

The user-supplied DB instance identifier, which must match the identifier of an
existing instance owned by the AWS account. This parameter isn't
case-sensitive.

Type: String

Required: No

**Filters.Filter.N**

A filter that specifies one or more database tenants to describe.

Supported filters:

- `tenant-db-name` \- Tenant database names. The results list only
includes information about the tenant databases that match these tenant DB
names.

- `tenant-database-resource-id` \- Tenant database resource
identifiers.

- `dbi-resource-id` \- DB instance resource identifiers. The results
list only includes information about the tenants contained within the DB
instances identified by these resource identifiers.

Type: Array of [Filter](api-filter.md) objects

Required: No

**Marker**

An optional pagination token provided by a previous
`DescribeTenantDatabases` request. If this parameter is specified, the
response includes only records beyond the marker, up to the value specified by
`MaxRecords`.

Type: String

Required: No

**MaxRecords**

The maximum number of records to include in the response. If more records exist than
the specified `MaxRecords` value, a pagination token called a marker is
included in the response so that you can retrieve the remaining results.

Type: Integer

Required: No

**TenantDBName**

The user-supplied tenant database name, which must match the name of an existing
tenant database on the specified DB instance owned by your AWS account. This parameter
isn’t case-sensitive.

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**Marker**

An optional pagination token provided by a previous
`DescribeTenantDatabases` request. If this parameter is specified, the
response includes only records beyond the marker, up to the value specified by
`MaxRecords`.

Type: String

**TenantDatabases.TenantDatabase.N**

An array of the tenant databases requested by the `DescribeTenantDatabases`
operation.

Type: Array of [TenantDatabase](api-tenantdatabase.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**DBInstanceNotFound**

`DBInstanceIdentifier` doesn't refer to an existing DB instance.

HTTP Status Code: 404

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/describetenantdatabases.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/describetenantdatabases.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/describetenantdatabases.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/describetenantdatabases.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/describetenantdatabases.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/describetenantdatabases.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/describetenantdatabases.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/describetenantdatabases.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/describetenantdatabases.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/describetenantdatabases.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeSourceRegions

DescribeValidDBInstanceModifications
