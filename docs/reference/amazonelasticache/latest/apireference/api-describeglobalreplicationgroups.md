# DescribeGlobalReplicationGroups

Returns information about a particular global replication group. If no identifier is
specified, returns information about all Global datastores.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**GlobalReplicationGroupId**

The name of the Global datastore

Type: String

Required: No

**Marker**

An optional marker returned from a prior request. Use this marker for pagination of
results from this operation. If this parameter is specified, the response includes only
records beyond the marker, up to the value specified by `MaxRecords`.

Type: String

Required: No

**MaxRecords**

The maximum number of records to include in the response. If more records exist than
the specified MaxRecords value, a marker is included in the response so that the
remaining results can be retrieved.

Type: Integer

Required: No

**ShowMemberInfo**

Returns the list of members that comprise the Global datastore.

Type: Boolean

Required: No

## Response Elements

The following elements are returned by the service.

**GlobalReplicationGroups.GlobalReplicationGroup.N**

Indicates the slot configuration and global identifier for each slice group.

Type: Array of [GlobalReplicationGroup](api-globalreplicationgroup.md) objects

**Marker**

An optional marker returned from a prior request. Use this marker for pagination of
results from this operation. If this parameter is specified, the response includes only
records beyond the marker, up to the value specified by MaxRecords. >

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**GlobalReplicationGroupNotFoundFault**

The Global datastore does not exist

HTTP Status Code: 404

**InvalidParameterCombination**

Two or more incompatible parameters were specified.

**message**

Two or more parameters that must not be used together were used together.

HTTP Status Code: 400

**InvalidParameterValue**

The value for a parameter is invalid.

**message**

A parameter value is invalid.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/elasticache-2015-02-02/describeglobalreplicationgroups.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/elasticache-2015-02-02/describeglobalreplicationgroups.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/elasticache-2015-02-02/describeglobalreplicationgroups.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/elasticache-2015-02-02/describeglobalreplicationgroups.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/elasticache-2015-02-02/describeglobalreplicationgroups.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/elasticache-2015-02-02/describeglobalreplicationgroups.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/elasticache-2015-02-02/describeglobalreplicationgroups.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/elasticache-2015-02-02/describeglobalreplicationgroups.md)

- [AWS SDK for Python](../../../../services/goto/boto3/elasticache-2015-02-02/describeglobalreplicationgroups.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/elasticache-2015-02-02/describeglobalreplicationgroups.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeEvents

DescribeReplicationGroups

All content copied from https://docs.aws.amazon.com/.
