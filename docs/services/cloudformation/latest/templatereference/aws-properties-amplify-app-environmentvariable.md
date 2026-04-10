This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Amplify::App EnvironmentVariable

Environment variables are key-value pairs that are available at build time. Set
environment variables for all branches in your app.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "Value" : String
}

```

### YAML

```yaml

  Name: String
  Value: String

```

## Properties

`Name`

The environment variable name.

_Required_: Yes

_Type_: String

_Pattern_: `(?s).*`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The environment variable value.

_Required_: Yes

_Type_: String

_Pattern_: `(?s).*`

_Maximum_: `5500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomRule

JobConfig

All content copied from https://docs.aws.amazon.com/.
