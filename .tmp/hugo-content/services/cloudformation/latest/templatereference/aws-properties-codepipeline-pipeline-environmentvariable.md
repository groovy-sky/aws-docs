This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodePipeline::Pipeline EnvironmentVariable

The environment variables for the action.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "Type" : String,
  "Value" : String
}

```

### YAML

```yaml

  Name: String
  Type: String
  Value: String

```

## Properties

`Name`

The environment variable name in the key-value pair.

_Required_: Yes

_Type_: String

_Pattern_: `[A-Za-z0-9_]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

Specifies the type of use for the environment variable value. The value can be either
`PLAINTEXT` or `SECRETS_MANAGER`. If the value is `SECRETS_MANAGER`, provide the Secrets
reference in the EnvironmentVariable value.

_Required_: No

_Type_: String

_Allowed values_: `PLAINTEXT | SECRETS_MANAGER`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The environment variable value in the key-value pair.

_Required_: Yes

_Type_: String

_Pattern_: `.*`

_Minimum_: `1`

_Maximum_: `2000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EncryptionKey

FailureConditions

All content copied from https://docs.aws.amazon.com/.
