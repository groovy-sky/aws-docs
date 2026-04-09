# DescribeUsers

Returns a list of users.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**Engine**

The engine.

Type: String

Pattern: `[a-zA-Z]*`

Required: No

**Filters.member.N**

Filter to determine the list of User IDs to return.

Type: Array of [Filter](api-filter.md) objects

Required: No

**Marker**

An optional marker returned from a prior request. Use this marker for pagination of
results from this operation. If this parameter is specified, the response includes only
records beyond the marker, up to the value specified by MaxRecords. >

Type: String

Required: No

**MaxRecords**

The maximum number of records to include in the response. If more records exist than
the specified MaxRecords value, a marker is included in the response so that the
remaining results can be retrieved.

Type: Integer

Required: No

**UserId**

The ID of the user.

Type: String

Length Constraints: Minimum length of 1.

Pattern: `[a-zA-Z][a-zA-Z0-9\-]*`

Required: No

## Response Elements

The following elements are returned by the service.

**Marker**

An optional marker returned from a prior request. Use this marker for pagination of
results from this operation. If this parameter is specified, the response includes only
records beyond the marker, up to the value specified by MaxRecords. >

Type: String

**Users.member.N**

A list of users.

Type: Array of [User](api-user.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidParameterCombination**

Two or more incompatible parameters were specified.

**message**

Two or more parameters that must not be used together were used together.

HTTP Status Code: 400

**ServiceLinkedRoleNotFoundFault**

The specified service linked role (SLR) was not found.

HTTP Status Code: 400

**UserNotFound**

The user does not exist or could not be found.

HTTP Status Code: 404

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/elasticache-2015-02-02/describeusers.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/elasticache-2015-02-02/describeusers.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/elasticache-2015-02-02/describeusers.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/elasticache-2015-02-02/describeusers.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/elasticache-2015-02-02/describeusers.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/elasticache-2015-02-02/describeusers.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/elasticache-2015-02-02/describeusers.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/elasticache-2015-02-02/describeusers.md)

- [AWS SDK for Python](../../../../services/goto/boto3/elasticache-2015-02-02/describeusers.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/elasticache-2015-02-02/describeusers.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeUserGroups

DisassociateGlobalReplicationGroup

All content copied from https://docs.aws.amazon.com/.
