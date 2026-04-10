This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::Flow SnowflakeDestinationProperties

The properties that are applied when Snowflake is being used as a destination.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BucketPrefix" : String,
  "ErrorHandlingConfig" : ErrorHandlingConfig,
  "IntermediateBucketName" : String,
  "Object" : String
}

```

### YAML

```yaml

  BucketPrefix: String
  ErrorHandlingConfig:
    ErrorHandlingConfig
  IntermediateBucketName: String
  Object: String

```

## Properties

`BucketPrefix`

The object key for the destination bucket in which Amazon AppFlow places the files.

_Required_: No

_Type_: String

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ErrorHandlingConfig`

The settings that determine how Amazon AppFlow handles an error when placing data in
the Snowflake destination. For example, this setting would determine if the flow should fail
after one insertion error, or continue and attempt to insert every record regardless of the
initial failure. `ErrorHandlingConfig` is a part of the destination connector
details.

_Required_: No

_Type_: [ErrorHandlingConfig](aws-properties-appflow-flow-errorhandlingconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IntermediateBucketName`

The intermediate bucket that Amazon AppFlow uses when moving data into Snowflake.

_Required_: Yes

_Type_: String

_Pattern_: `\S+`

_Minimum_: `3`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Object`

The object specified in the Snowflake flow destination.

_Required_: Yes

_Type_: String

_Pattern_: `\S+`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [SnowflakeDestinationProperties](../../../../reference/appflow/1-0/apireference/api-snowflakedestinationproperties.md) in the _Amazon AppFlow API_
_Reference_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SlackSourceProperties

SourceConnectorProperties

All content copied from https://docs.aws.amazon.com/.
