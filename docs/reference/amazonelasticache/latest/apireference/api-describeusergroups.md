# DescribeUserGroups

Returns a list of user groups.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

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

**UserGroupId**

The ID of the user group.

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**Marker**

An optional marker returned from a prior request. Use this marker for pagination of
results from this operation. If this parameter is specified, the response includes only
records beyond the marker, up to the value specified by MaxRecords.>

Type: String

**UserGroups.member.N**

Returns a list of user groups.

Type: Array of [UserGroup](api-usergroup.md) objects

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

**UserGroupNotFound**

The user group was not found or does not exist

HTTP Status Code: 404

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/elasticache-2015-02-02/describeusergroups.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/elasticache-2015-02-02/describeusergroups.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/elasticache-2015-02-02/describeusergroups.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/elasticache-2015-02-02/describeusergroups.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/elasticache-2015-02-02/describeusergroups.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/elasticache-2015-02-02/describeusergroups.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/elasticache-2015-02-02/describeusergroups.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/elasticache-2015-02-02/describeusergroups.md)

- [AWS SDK for Python](../../../../services/goto/boto3/elasticache-2015-02-02/describeusergroups.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/elasticache-2015-02-02/describeusergroups.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeUpdateActions

DescribeUsers

All content copied from https://docs.aws.amazon.com/.
