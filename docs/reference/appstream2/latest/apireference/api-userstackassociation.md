# UserStackAssociation

Describes a user in the user pool and the associated stack.

## Contents

**AuthenticationType**

The authentication type for the user.

Type: String

Valid Values: `API | SAML | USERPOOL | AWS_AD`

Required: Yes

**StackName**

The name of the stack that is associated with the user.

Type: String

Length Constraints: Minimum length of 1.

Required: Yes

**UserName**

The email address of the user who is associated with the stack.

###### Note

Users' email addresses are case-sensitive.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Pattern: `[\p{L}\p{M}\p{S}\p{N}\p{P}]+`

Required: Yes

**SendEmailNotification**

Specifies whether a welcome email is sent to a user after the user is created in the user pool.

Type: Boolean

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appstream-2016-12-01/userstackassociation.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appstream-2016-12-01/userstackassociation.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appstream-2016-12-01/userstackassociation.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UserSetting

UserStackAssociationError

All content copied from https://docs.aws.amazon.com/.
