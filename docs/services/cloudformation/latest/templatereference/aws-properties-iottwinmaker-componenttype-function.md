This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTTwinMaker::ComponentType Function

The function body.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ImplementedBy" : DataConnector,
  "RequiredProperties" : [ String, ... ],
  "Scope" : String
}

```

### YAML

```yaml

  ImplementedBy:
    DataConnector
  RequiredProperties:
    - String
  Scope: String

```

## Properties

`ImplementedBy`

The data connector.

_Required_: No

_Type_: [DataConnector](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iottwinmaker-componenttype-dataconnector.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RequiredProperties`

The required properties of the function.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Scope`

The scope of the function.

_Required_: No

_Type_: String

_Allowed values_: `ENTITY | WORKSPACE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Error

LambdaFunction
