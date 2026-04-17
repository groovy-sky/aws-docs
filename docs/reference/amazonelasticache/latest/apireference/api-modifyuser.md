---
title: "ModifyUser"
---

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

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/elasticache-2015-02-02/ModifyUser)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/elasticache-2015-02-02/ModifyUser)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/elasticache-2015-02-02/ModifyUser)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/elasticache-2015-02-02/ModifyUser)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/elasticache-2015-02-02/ModifyUser)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/elasticache-2015-02-02/ModifyUser)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/elasticache-2015-02-02/ModifyUser)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/elasticache-2015-02-02/ModifyUser)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/elasticache-2015-02-02/ModifyUser)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/elasticache-2015-02-02/ModifyUser)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ModifyServerlessCache

ModifyUserGroup

All content copied from https://docs.aws.amazon.com/.
