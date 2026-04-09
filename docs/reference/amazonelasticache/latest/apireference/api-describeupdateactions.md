# DescribeUpdateActions

Returns details of the update actions

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**CacheClusterIds.member.N**

The cache cluster IDs

Type: Array of strings

Array Members: Maximum number of 20 items.

Required: No

**Engine**

The Elasticache engine to which the update applies. Either Valkey, Redis OSS or Memcached.

Type: String

Required: No

**Marker**

An optional marker returned from a prior request. Use this marker for pagination of
results from this operation. If this parameter is specified, the response includes only
records beyond the marker, up to the value specified by `MaxRecords`.

Type: String

Required: No

**MaxRecords**

The maximum number of records to include in the response

Type: Integer

Required: No

**ReplicationGroupIds.member.N**

The replication group IDs

Type: Array of strings

Array Members: Maximum number of 20 items.

Required: No

**ServiceUpdateName**

The unique ID of the service update

Type: String

Required: No

**ServiceUpdateStatus.member.N**

The status of the service update

Type: Array of strings

Array Members: Maximum number of 3 items.

Valid Values: `available | cancelled | expired`

Required: No

**ServiceUpdateTimeRange**

The range of time specified to search for service updates that are in available
status

Type: [TimeRangeFilter](api-timerangefilter.md) object

Required: No

**ShowNodeLevelUpdateStatus**

Dictates whether to include node level update status in the response

Type: Boolean

Required: No

**UpdateActionStatus.member.N**

The status of the update action.

Type: Array of strings

Array Members: Maximum number of 9 items.

Valid Values: `not-applied | waiting-to-start | in-progress | stopping | stopped | complete | scheduling | scheduled | not-applicable`

Required: No

## Response Elements

The following elements are returned by the service.

**Marker**

An optional marker returned from a prior request. Use this marker for pagination of
results from this operation. If this parameter is specified, the response includes only
records beyond the marker, up to the value specified by `MaxRecords`.

Type: String

**UpdateActions.UpdateAction.N**

Returns a list of update actions

Type: Array of [UpdateAction](api-updateaction.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/elasticache-2015-02-02/describeupdateactions.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/elasticache-2015-02-02/describeupdateactions.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/elasticache-2015-02-02/describeupdateactions.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/elasticache-2015-02-02/describeupdateactions.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/elasticache-2015-02-02/describeupdateactions.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/elasticache-2015-02-02/describeupdateactions.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/elasticache-2015-02-02/describeupdateactions.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/elasticache-2015-02-02/describeupdateactions.md)

- [AWS SDK for Python](../../../../services/goto/boto3/elasticache-2015-02-02/describeupdateactions.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/elasticache-2015-02-02/describeupdateactions.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeSnapshots

DescribeUserGroups

All content copied from https://docs.aws.amazon.com/.
