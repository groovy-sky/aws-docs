This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::Connection BasicAuthenticationCredentials

The basic authentication credentials of a connection.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Password" : String,
  "UserName" : String
}

```

### YAML

```yaml

  Password: String
  UserName: String

```

## Properties

`Password`

The password for a connection.

_Required_: No

_Type_: String

_Pattern_: `^.*$`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserName`

The user name for the connecion.

_Required_: No

_Type_: String

_Pattern_: `^\S+$`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AwsLocation

ConnectionPropertiesInput

All content copied from https://docs.aws.amazon.com/.
