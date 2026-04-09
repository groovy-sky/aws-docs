# CreateUser

For Valkey engine version 7.2 onwards and Redis OSS 6.0 to 7.1: Creates a user. For more information, see
[Using Role Based Access Control (RBAC)](../../../../services/amazonelasticache/latest/dg/clusters-rbac.md).

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**AccessString**

Access permissions string used for this user.

Type: String

Pattern: `.*\S.*`

Required: Yes

**Engine**

The options are valkey or redis.

Type: String

Pattern: `[a-zA-Z]*`

Required: Yes

**UserId**

The ID of the user. This value is stored as a lowercase string.

Type: String

Length Constraints: Minimum length of 1.

Pattern: `[a-zA-Z][a-zA-Z0-9\-]*`

Required: Yes

**UserName**

The username of the user.

Type: String

Length Constraints: Minimum length of 1.

Required: Yes

**AuthenticationMode**

Specifies how to authenticate the user.

Type: [AuthenticationMode](api-authenticationmode.md) object

Required: No

**NoPasswordRequired**

Indicates a password is not required for this user.

Type: Boolean

Required: No

**Passwords.member.N**

Passwords used for this user. You can create up to two passwords for each user.

Type: Array of strings

Array Members: Minimum number of 1 item.

Required: No

**Tags.Tag.N**

A list of tags to be added to this resource. A tag is a key-value pair. A tag key must
be accompanied by a tag value, although null is accepted.

Type: Array of [Tag](api-tag.md) objects

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

**DuplicateUserName**

A user with this username already exists.

HTTP Status Code: 400

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

**ServiceLinkedRoleNotFoundFault**

The specified service linked role (SLR) was not found.

HTTP Status Code: 400

**TagQuotaPerResourceExceeded**

The request cannot be processed because it would cause the resource to have more than
the allowed number of tags. The maximum number of tags permitted on a resource is
50.

HTTP Status Code: 400

**UserAlreadyExists**

A user with this ID already exists.

HTTP Status Code: 400

**UserQuotaExceeded**

The quota of users has been exceeded.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/elasticache-2015-02-02/createuser.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/elasticache-2015-02-02/createuser.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/elasticache-2015-02-02/createuser.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/elasticache-2015-02-02/createuser.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/elasticache-2015-02-02/createuser.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/elasticache-2015-02-02/createuser.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/elasticache-2015-02-02/createuser.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/elasticache-2015-02-02/createuser.md)

- [AWS SDK for Python](../../../../services/goto/boto3/elasticache-2015-02-02/createuser.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/elasticache-2015-02-02/createuser.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateSnapshot

CreateUserGroup

All content copied from https://docs.aws.amazon.com/.
