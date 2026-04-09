# ModifyUser

Changes user password(s) and/or access string.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**UserId**

The ID of the user.

Type: String

Length Constraints: Minimum length of 1.

Pattern: `[a-zA-Z][a-zA-Z0-9\-]*`

Required: Yes

**AccessString**

Access permissions string used for this user.

Type: String

Pattern: `.*\S.*`

Required: No

**AppendAccessString**

Adds additional user permissions to the access string.

Type: String

Pattern: `.*\S.*`

Required: No

**AuthenticationMode**

Specifies how to authenticate the user.

Type: [AuthenticationMode](api-authenticationmode.md) object

Required: No

**Engine**

Modifies the engine listed for a user. The options are valkey or redis.

Type: String

Pattern: `[a-zA-Z]*`

Required: No

**NoPasswordRequired**

Indicates no password is required for the user.

Type: Boolean

Required: No

**Passwords.member.N**

The passwords belonging to the user. You are allowed up to two.

Type: Array of strings

Array Members: Minimum number of 1 item.

Required: No

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/elasticache-2015-02-02/modifyuser.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/elasticache-2015-02-02/modifyuser.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/elasticache-2015-02-02/modifyuser.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/elasticache-2015-02-02/modifyuser.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/elasticache-2015-02-02/modifyuser.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/elasticache-2015-02-02/modifyuser.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/elasticache-2015-02-02/modifyuser.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/elasticache-2015-02-02/modifyuser.md)

- [AWS SDK for Python](../../../../services/goto/boto3/elasticache-2015-02-02/modifyuser.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/elasticache-2015-02-02/modifyuser.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ModifyServerlessCache

ModifyUserGroup

All content copied from https://docs.aws.amazon.com/.
