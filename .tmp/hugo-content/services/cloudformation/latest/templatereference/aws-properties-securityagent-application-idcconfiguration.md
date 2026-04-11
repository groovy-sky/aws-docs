This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityAgent::Application IdCConfiguration

The IAM Identity Center configuration for an application.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IdCApplicationArn" : String,
  "IdCInstanceArn" : String
}

```

### YAML

```yaml

  IdCApplicationArn: String
  IdCInstanceArn: String

```

## Properties

`IdCApplicationArn`

The Amazon Resource Name (ARN) of the IAM Identity Center application.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IdCInstanceArn`

The Amazon Resource Name (ARN) of the IAM Identity Center instance.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::SecurityAgent::Application

Tag

All content copied from https://docs.aws.amazon.com/.
