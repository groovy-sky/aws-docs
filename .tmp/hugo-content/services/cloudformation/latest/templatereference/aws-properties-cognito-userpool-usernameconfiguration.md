This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cognito::UserPool UsernameConfiguration

Case sensitivity of the username input for the selected sign-in option. When case
sensitivity is set to `False` (case insensitive), users can sign in with any
combination of capital and lowercase letters. For example, `username`,
`USERNAME`, or `UserName`, or for email,
`email@example.com` or `EMaiL@eXamplE.Com`. For most use
cases, set case sensitivity to `False` (case insensitive) as a best practice.
When usernames and email addresses are case insensitive, Amazon Cognito treats any
variation in case as the same user, and prevents a case variation from being assigned to
the same attribute for a different user.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CaseSensitive" : Boolean
}

```

### YAML

```yaml

  CaseSensitive: Boolean

```

## Properties

`CaseSensitive`

Specifies whether user name case sensitivity will be applied for all users in the user
pool through Amazon Cognito APIs. For most use cases, set case sensitivity to `False`
(case insensitive) as a best practice. When usernames and email addresses are case
insensitive, users can sign in as the same user when they enter a different
capitalization of their user name.

Valid values include:

true

Enables case sensitivity for all username input. When this option is set
to `true`, users must sign in using the exact capitalization of
their given username, such as “UserName”. This is the default value.

false

Enables case insensitivity for all username input. For example, when this
option is set to `false`, users can sign in using
`username`, `USERNAME`, or `UserName`.
This option also enables both `preferred_username` and
`email` alias to be case insensitive, in addition to the
`username` attribute.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UserAttributeUpdateSettings

UserPoolAddOns

All content copied from https://docs.aws.amazon.com/.
