This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppSync::ChannelNamespace LambdaConfig

The `LambdaConfig` property type specifies the integration configuration for a Lambda data source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InvokeType" : String
}

```

### YAML

```yaml

  InvokeType: String

```

## Properties

`InvokeType`

The invocation type for a Lambda data source.

_Required_: Yes

_Type_: String

_Allowed values_: `REQUEST_RESPONSE | EVENT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Integration

Tag

All content copied from https://docs.aws.amazon.com/.
