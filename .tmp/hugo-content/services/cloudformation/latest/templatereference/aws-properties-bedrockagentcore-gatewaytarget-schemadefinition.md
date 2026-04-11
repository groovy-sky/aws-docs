This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::GatewayTarget SchemaDefinition

A schema definition for a gateway target. This structure defines the structure of the API that the target exposes.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Description" : String,
  "Items" : SchemaDefinition,
  "Properties" : {Key: Value, ...},
  "Required" : [ String, ... ],
  "Type" : String
}

```

### YAML

```yaml

  Description: String
  Items:
    SchemaDefinition
  Properties:
    Key: Value
  Required:
    - String
  Type: String

```

## Properties

`Description`

The description of the schema definition. This description provides information about the purpose and usage of the schema.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Items`

The items in the schema definition. This field is used for array types to define the structure of the array elements.

_Required_: No

_Type_: [SchemaDefinition](aws-properties-bedrockagentcore-gatewaytarget-schemadefinition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Properties`

The properties of the schema definition. These properties define the fields in the schema.

_Required_: No

_Type_: Object of [SchemaDefinition](aws-properties-bedrockagentcore-gatewaytarget-schemadefinition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Required`

The required fields in the schema definition. These fields must be provided when using the schema.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of the schema definition. This field specifies the data type of the schema.

_Required_: Yes

_Type_: String

_Allowed values_: `string | number | object | array | boolean | integer`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3Configuration

TargetConfiguration

All content copied from https://docs.aws.amazon.com/.
