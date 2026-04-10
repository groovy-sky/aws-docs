This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::Flow EventBridgeDestinationProperties

The properties that are applied when Amazon EventBridge is being used as a
destination.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ErrorHandlingConfig" : ErrorHandlingConfig,
  "Object" : String
}

```

### YAML

```yaml

  ErrorHandlingConfig:
    ErrorHandlingConfig
  Object: String

```

## Properties

`ErrorHandlingConfig`

The object specified in the Amplitude flow source.

_Required_: No

_Type_: [ErrorHandlingConfig](aws-properties-appflow-flow-errorhandlingconfig.md)

_Pattern_: `\S+`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Object`

The object specified in the Amazon EventBridge flow destination.

_Required_: Yes

_Type_: String

_Pattern_: `\S+`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [EventBridgeDestinationProperties](../../../../reference/appflow/1-0/apireference/api-eventbridgedestinationproperties.md) in the _Amazon AppFlow API_
_Reference_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ErrorHandlingConfig

GlueDataCatalog

All content copied from https://docs.aws.amazon.com/.
