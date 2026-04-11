This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Amplify::Branch BasicAuthConfig

Use the BasicAuthConfig property type to set password protection for a specific
branch.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EnableBasicAuth" : Boolean,
  "Password" : String,
  "Username" : String
}

```

### YAML

```yaml

  EnableBasicAuth: Boolean
  Password: String
  Username: String

```

## Properties

`EnableBasicAuth`

Enables basic authorization for the branch.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Password`

The password for basic authorization.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Username`

Property description not available.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Backend

EnvironmentVariable

All content copied from https://docs.aws.amazon.com/.
