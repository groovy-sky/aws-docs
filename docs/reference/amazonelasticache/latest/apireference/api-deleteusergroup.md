# DeleteUserGroup

For Valkey engine version 7.2 onwards and Redis OSS 6.0 onwards: Deletes a user group. The user group must first
be disassociated from the replication group before it can be deleted. For more
information, see [Using Role Based Access Control (RBAC)](../../../../services/amazonelasticache/latest/dg/clusters-rbac.md).

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**UserGroupId**

The ID of the user group.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**ARN**

The Amazon Resource Name (ARN) of the user group.

Type: String

**Engine**

The options are valkey or redis.

Type: String

Pattern: `[a-zA-Z]*`

**MinimumEngineVersion**

The minimum engine version required, which is Redis OSS 6.0

Type: String

**PendingChanges**

A list of updates being applied to the user group.

Type: [UserGroupPendingChanges](api-usergrouppendingchanges.md) object

**ReplicationGroups.member.N**

A list of replication groups that the user group can access.

Type: Array of strings

**ServerlessCaches.member.N**

Indicates which serverless caches the specified user group is associated with. Available for Valkey, Redis OSS and Serverless Memcached only.

Type: Array of strings

**Status**

Indicates user group status. Can be "creating", "active", "modifying",
"deleting".

Type: String

**UserGroupId**

The ID of the user group.

Type: String

**UserIds.member.N**

The list of user IDs that belong to the user group.

Type: Array of strings

Length Constraints: Minimum length of 1.

Pattern: `[a-zA-Z][a-zA-Z0-9\-]*`

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidParameterValue**

The value for a parameter is invalid.

**message**

A parameter value is invalid.

HTTP Status Code: 400

**InvalidUserGroupState**

The user group is not in an active state.

HTTP Status Code: 400

**ServiceLinkedRoleNotFoundFault**

The specified service linked role (SLR) was not found.

HTTP Status Code: 400

**UserGroupNotFound**

The user group was not found or does not exist

HTTP Status Code: 404

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/elasticache-2015-02-02/deleteusergroup.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/elasticache-2015-02-02/deleteusergroup.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/elasticache-2015-02-02/deleteusergroup.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/elasticache-2015-02-02/deleteusergroup.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/elasticache-2015-02-02/deleteusergroup.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/elasticache-2015-02-02/deleteusergroup.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/elasticache-2015-02-02/deleteusergroup.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/elasticache-2015-02-02/deleteusergroup.md)

- [AWS SDK for Python](../../../../services/goto/boto3/elasticache-2015-02-02/deleteusergroup.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/elasticache-2015-02-02/deleteusergroup.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteUser

DescribeCacheClusters

All content copied from https://docs.aws.amazon.com/.
