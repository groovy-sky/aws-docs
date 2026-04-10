This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Connection BasicAuthenticationCredentials

For supplying basic auth credentials when not providing a `SecretArn` value.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Password" : String,
  "Username" : String
}

```

### YAML

```yaml

  Password: String
  Username: String

```

## Properties

`Password`

The password to connect to the data source.

_Required_: No

_Type_: String

_Pattern_: `.*`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Username`

The username to connect to the data source.

_Required_: No

_Type_: String

_Pattern_: `\S+`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AuthorizationCodeProperties

ConnectionInput

All content copied from https://docs.aws.amazon.com/.
