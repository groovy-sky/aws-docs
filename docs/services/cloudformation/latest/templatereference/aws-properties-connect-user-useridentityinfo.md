This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::User UserIdentityInfo

Contains information about the identity of a user.

###### Note

For Amazon Connect instances that are created with the `EXISTING_DIRECTORY` identity management
type, `FirstName`, `LastName`, and `Email` cannot be updated from within Amazon Connect because they are managed by the directory.

###### Important

The `FirstName` and `LastName` length constraints below apply only to instances using
SAML for identity management. If you are using Amazon Connect for identity management, the length constraints
are 1-255 for `FirstName`, and 1-256 for `LastName`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Email" : String,
  "FirstName" : String,
  "LastName" : String,
  "Mobile" : String,
  "SecondaryEmail" : String
}

```

### YAML

```yaml

  Email: String
  FirstName: String
  LastName: String
  Mobile: String
  SecondaryEmail: String

```

## Properties

`Email`

The email address. If you are using SAML for identity management and include this parameter, an error is
returned.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FirstName`

The first name. This is required if you are using Amazon Connect or SAML for identity management. Inputs
must be in Unicode Normalization Form C (NFC). Text containing characters in a non-NFC form (for example, decomposed
characters or combining marks) are not accepted.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LastName`

The last name. This is required if you are using Amazon Connect or SAML for identity management. Inputs must
be in Unicode Normalization Form C (NFC). Text containing characters in a non-NFC form (for example, decomposed
characters or combining marks) are not accepted.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `300`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Mobile`

The user's mobile number.

_Required_: No

_Type_: String

_Pattern_: `^\+[1-9]\d{1,14}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecondaryEmail`

The user's secondary email address. If you provide a secondary email, the user
receives email notifications -- other than password reset notifications -- to this email
address instead of to their primary email address.

_Pattern_:
`(?=^.{0,265}$)[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,63}`

_Required_: No

_Type_: String

_Pattern_: `(?=^.{0,265}$)[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,63}`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

UserPhoneConfig

All content copied from https://docs.aws.amazon.com/.
