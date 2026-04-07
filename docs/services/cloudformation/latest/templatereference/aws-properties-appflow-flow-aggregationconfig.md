This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::Flow AggregationConfig

The aggregation settings that you can use to customize the output format of your flow
data.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AggregationType" : String,
  "TargetFileSize" : Integer
}

```

### YAML

```yaml

  AggregationType: String
  TargetFileSize: Integer

```

## Properties

`AggregationType`

Specifies whether Amazon AppFlow aggregates the flow records into a single file, or
leave them unaggregated.

_Required_: No

_Type_: String

_Allowed values_: `None | SingleFile`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetFileSize`

The desired file size, in MB, for each output file that Amazon AppFlow writes to the
flow destination. For each file, Amazon AppFlow attempts to achieve the size that you
specify. The actual file sizes might differ from this target based on the number and size of
the records that each file contains.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [AggregationConfig](../../../../reference/appflow/1-0/apireference/api-aggregationconfig.md) in
the _Amazon AppFlow API Reference_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::AppFlow::Flow

AmplitudeSourceProperties
