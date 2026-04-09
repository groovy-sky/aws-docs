# UserGroup

## Contents

###### Note

In the following list, the required parameters are described first.

**ARN**

The Amazon Resource Name (ARN) of the user group.

Type: String

Required: No

**Engine**

The options are valkey or redis.

Type: String

Pattern: `[a-zA-Z]*`

Required: No

**MinimumEngineVersion**

The minimum engine version required, which is Redis OSS 6.0

Type: String

Required: No

**PendingChanges**

A list of updates being applied to the user group.

Type: [UserGroupPendingChanges](api-usergrouppendingchanges.md) object

Required: No

**ReplicationGroups.member.N**

A list of replication groups that the user group can access.

Type: Array of strings

Required: No

**ServerlessCaches.member.N**

Indicates which serverless caches the specified user group is associated with. Available for Valkey, Redis OSS and Serverless Memcached only.

Type: Array of strings

Required: No

**Status**

Indicates user group status. Can be "creating", "active", "modifying",
"deleting".

Type: String

Required: No

**UserGroupId**

The ID of the user group.

Type: String

Required: No

**UserIds.member.N**

The list of user IDs that belong to the user group.

Type: Array of strings

Length Constraints: Minimum length of 1.

Pattern: `[a-zA-Z][a-zA-Z0-9\-]*`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/elasticache-2015-02-02/usergroup.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/elasticache-2015-02-02/usergroup.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/elasticache-2015-02-02/usergroup.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

User

UserGroupPendingChanges

All content copied from https://docs.aws.amazon.com/.
