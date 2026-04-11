This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::Connection UsernamePassword

The username and password of a connection.

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

The password of a connection.

_Required_: Yes

_Type_: String

_Pattern_: `^[\S]*$`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Username`

The username of a connection.

_Required_: Yes

_Type_: String

_Pattern_: `^[\S]*$`

_Minimum_: `1`

_Maximum_: `127`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SparkGluePropertiesInput

WorkflowsMwaaPropertiesInput

All content copied from https://docs.aws.amazon.com/.
