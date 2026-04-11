This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::Flow CustomConnectorDestinationProperties

The properties that are applied when the custom connector is being used as a
destination.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomProperties" : {Key: Value, ...},
  "EntityName" : String,
  "ErrorHandlingConfig" : ErrorHandlingConfig,
  "IdFieldNames" : [ String, ... ],
  "WriteOperationType" : String
}

```

### YAML

```yaml

  CustomProperties:
    Key: Value
  EntityName: String
  ErrorHandlingConfig:
    ErrorHandlingConfig
  IdFieldNames:
    - String
  WriteOperationType: String

```

## Properties

`CustomProperties`

The custom properties that are specific to the connector when it's used as a destination
in the flow.

_Required_: No

_Type_: Object of String

_Pattern_: `^[\w]{1,2048}$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EntityName`

The entity specified in the custom connector as a destination in the flow.

_Required_: Yes

_Type_: String

_Pattern_: `\S+`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ErrorHandlingConfig`

The settings that determine how Amazon AppFlow handles an error when placing data in
the custom connector as destination.

_Required_: No

_Type_: [ErrorHandlingConfig](aws-properties-appflow-flow-errorhandlingconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IdFieldNames`

The name of the field that Amazon AppFlow uses as an ID when performing a write
operation such as update, delete, or upsert.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WriteOperationType`

Specifies the type of write operation to be performed in the custom connector when it's
used as destination.

_Required_: No

_Type_: String

_Allowed values_: `INSERT | UPSERT | UPDATE | DELETE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConnectorOperator

CustomConnectorSourceProperties

All content copied from https://docs.aws.amazon.com/.
