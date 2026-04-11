This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cognito::UserPool SignInPolicy

The policy for allowed types of authentication in a user pool. To activate this
setting, your user pool must be in the [Essentials tier](../../../cognito/latest/developerguide/feature-plans-features-essentials.md) or higher.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AllowedFirstAuthFactors" : [ String, ... ]
}

```

### YAML

```yaml

  AllowedFirstAuthFactors:
    - String

```

## Properties

`AllowedFirstAuthFactors`

The sign-in methods that a user pool supports as the first factor. You can permit
users to start authentication with a standard username and password, or with other
one-time password and hardware factors.

Supports values of `EMAIL_OTP`, `SMS_OTP`,
`WEB_AUTHN` and `PASSWORD`,

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `4`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SchemaAttribute

SmsConfiguration

All content copied from https://docs.aws.amazon.com/.
