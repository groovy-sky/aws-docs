# CreateUserGroup

For Valkey engine version 7.2 onwards and Redis OSS 6.0 to 7.1: Creates a user group. For more
information, see [Using Role Based Access Control (RBAC)](../../../../services/amazonelasticache/latest/dg/clusters-rbac.md)

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**Engine**

Sets the engine listed in a user group. The options are valkey or redis.

Type: String

Pattern: `[a-zA-Z]*`

Required: Yes

**UserGroupId**

The ID of the user group. This value is stored as a lowercase string.

Type: String

Required: Yes

**Tags.Tag.N**

A list of tags to be added to this resource. A tag is a key-value pair. A tag key must
be accompanied by a tag value, although null is accepted. Available for Valkey and Redis OSS only.

Type: Array of [Tag](api-tag.md) objects

Required: No

**UserIds.member.N**

The list of user IDs that belong to the user group.

Type: Array of strings

Array Members: Minimum number of 1 item.

Length Constraints: Minimum length of 1.

Pattern: `[a-zA-Z][a-zA-Z0-9\-]*`

Required: No

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

**DefaultUserRequired**

You must add default user to a user group.

HTTP Status Code: 400

**DuplicateUserName**

A user with this username already exists.

HTTP Status Code: 400

**InvalidParameterValue**

The value for a parameter is invalid.

**message**

A parameter value is invalid.

HTTP Status Code: 400

**ServiceLinkedRoleNotFoundFault**

The specified service linked role (SLR) was not found.

HTTP Status Code: 400

**TagQuotaPerResourceExceeded**

The request cannot be processed because it would cause the resource to have more than
the allowed number of tags. The maximum number of tags permitted on a resource is
50.

HTTP Status Code: 400

**UserGroupAlreadyExists**

The user group with this ID already exists.

HTTP Status Code: 400

**UserGroupQuotaExceeded**

The number of users exceeds the user group limit.

HTTP Status Code: 400

**UserNotFound**

The user does not exist or could not be found.

HTTP Status Code: 404

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/elasticache-2015-02-02/createusergroup.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/elasticache-2015-02-02/createusergroup.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/elasticache-2015-02-02/createusergroup.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/elasticache-2015-02-02/createusergroup.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/elasticache-2015-02-02/createusergroup.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/elasticache-2015-02-02/createusergroup.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/elasticache-2015-02-02/createusergroup.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/elasticache-2015-02-02/createusergroup.md)

- [AWS SDK for Python](../../../../services/goto/boto3/elasticache-2015-02-02/createusergroup.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/elasticache-2015-02-02/createusergroup.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateUser

DecreaseNodeGroupsInGlobalReplicationGroup

All content copied from https://docs.aws.amazon.com/.
