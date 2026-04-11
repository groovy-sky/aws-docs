This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTTwinMaker::ComponentType DataConnector

The data connector.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IsNative" : Boolean,
  "Lambda" : LambdaFunction
}

```

### YAML

```yaml

  IsNative: Boolean
  Lambda:
    LambdaFunction

```

## Properties

`IsNative`

A boolean value that specifies whether the data connector is native to IoT TwinMaker.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Lambda`

The Lambda function associated with the data connector.

_Required_: No

_Type_: [LambdaFunction](aws-properties-iottwinmaker-componenttype-lambdafunction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CompositeComponentType

DataType

All content copied from https://docs.aws.amazon.com/.
