# DescribeDBProxyEndpoints

Returns information about DB proxy endpoints.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**DBProxyEndpointName**

The name of a DB proxy endpoint to describe. If you omit this parameter,
the output includes information about all DB proxy endpoints associated with
the specified proxy.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 63.

Pattern: `[a-zA-Z](?:-?[a-zA-Z0-9]+)*`

Required: No

**DBProxyName**

The name of the DB proxy whose endpoints you want to describe. If you omit
this parameter, the output includes information about all DB proxy endpoints
associated with all your DB proxies.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 63.

Pattern: `[a-zA-Z](?:-?[a-zA-Z0-9]+)*`

Required: No

**Filters.Filter.N**

This parameter is not currently supported.

Type: Array of [Filter](api-filter.md) objects

Required: No

**Marker**

An optional pagination token provided by a previous request.
If this parameter is specified, the response includes only records beyond the marker,
up to the value specified by `MaxRecords`.

Type: String

Required: No

**MaxRecords**

The maximum number of records to include in the response. If more records exist
than the specified `MaxRecords` value, a pagination token called a marker is
included in the response so that the remaining results can be retrieved.

Default: 100

Constraints: Minimum 20, maximum 100.

Type: Integer

Valid Range: Minimum value of 20. Maximum value of 100.

Required: No

## Response Elements

The following elements are returned by the service.

**DBProxyEndpoints.member.N**

The list of `ProxyEndpoint` objects returned by the API operation.

Type: Array of [DBProxyEndpoint](api-dbproxyendpoint.md) objects

**Marker**

An optional pagination token provided by a previous request.
If this parameter is specified, the response includes only records beyond the marker,
up to the value specified by `MaxRecords`.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**DBProxyEndpointNotFoundFault**

The DB proxy endpoint doesn't exist.

HTTP Status Code: 404

**DBProxyNotFoundFault**

The specified proxy name doesn't correspond to a proxy owned by your AWS account in the specified AWS Region.

HTTP Status Code: 404

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/describedbproxyendpoints.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/describedbproxyendpoints.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/describedbproxyendpoints.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/describedbproxyendpoints.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/describedbproxyendpoints.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/describedbproxyendpoints.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/describedbproxyendpoints.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/describedbproxyendpoints.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/describedbproxyendpoints.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/describedbproxyendpoints.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeDBProxies

DescribeDBProxyTargetGroups

All content copied from https://docs.aws.amazon.com/.
