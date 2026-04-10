This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::Flow MarketoDestinationProperties

The properties that Amazon AppFlow applies when you use Marketo as a flow
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

The settings that determine how Amazon AppFlow handles an error when placing data in
the destination. For example, this setting would determine if the flow should fail after one
insertion error, or continue and attempt to insert every record regardless of the initial
failure. `ErrorHandlingConfig` is a part of the destination connector details.

_Required_: No

_Type_: [ErrorHandlingConfig](aws-properties-appflow-flow-errorhandlingconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Object`

The object specified in the Marketo flow destination.

_Required_: Yes

_Type_: String

_Pattern_: `\S+`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LookoutMetricsDestinationProperties

MarketoSourceProperties

All content copied from https://docs.aws.amazon.com/.
