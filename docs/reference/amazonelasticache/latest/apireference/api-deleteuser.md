# DeleteUser

For Valkey engine version 7.2 onwards and Redis OSS 6.0 onwards: Deletes a user. The user will be removed from
all user groups and in turn removed from all replication groups. For more information,
see [Using Role Based Access Control (RBAC)](../../../../services/amazonelasticache/latest/dg/clusters-rbac.md).

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**UserId**

The ID of the user.

Type: String

Length Constraints: Minimum length of 1.

Pattern: `[a-zA-Z][a-zA-Z0-9\-]*`

Required: Yes

## Response Elements

The following elements are returned by the service.

**AccessString**

Access permissions string used for this user.

Type: String

**ARN**

The Amazon Resource Name (ARN) of the user.

Type: String

**Authentication**

Denotes whether the user requires a password to authenticate.

Type: [Authentication](api-authentication.md) object

**Engine**

The options are valkey or redis.

Type: String

Pattern: `[a-zA-Z]*`

**MinimumEngineVersion**

The minimum engine version required, which is Redis OSS 6.0

Type: String

**Status**

Indicates the user status. Can be "active", "modifying" or "deleting".

Type: String

**UserGroupIds.member.N**

Returns a list of the user group IDs the user belongs to.

Type: Array of strings

Length Constraints: Minimum length of 1.

Pattern: `[a-zA-Z][a-zA-Z0-9\-]*`

**UserId**

The ID of the user.

Type: String

**UserName**

The username of the user.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**DefaultUserAssociatedToUserGroup**

The default user assigned to the user group.

HTTP Status Code: 400

**InvalidParameterValue**

The value for a parameter is invalid.

**message**

A parameter value is invalid.

HTTP Status Code: 400

**InvalidUserState**

The user is not in active state.

HTTP Status Code: 400

**ServiceLinkedRoleNotFoundFault**

The specified service linked role (SLR) was not found.

HTTP Status Code: 400

**UserNotFound**

The user does not exist or could not be found.

HTTP Status Code: 404

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/elasticache-2015-02-02/deleteuser.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/elasticache-2015-02-02/deleteuser.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/elasticache-2015-02-02/deleteuser.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/elasticache-2015-02-02/deleteuser.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/elasticache-2015-02-02/deleteuser.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/elasticache-2015-02-02/deleteuser.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/elasticache-2015-02-02/deleteuser.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/elasticache-2015-02-02/deleteuser.md)

- [AWS SDK for Python](../../../../services/goto/boto3/elasticache-2015-02-02/deleteuser.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/elasticache-2015-02-02/deleteuser.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteSnapshot

DeleteUserGroup

All content copied from https://docs.aws.amazon.com/.
