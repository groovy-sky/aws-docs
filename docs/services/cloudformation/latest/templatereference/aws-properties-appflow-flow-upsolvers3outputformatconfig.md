This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::Flow UpsolverS3OutputFormatConfig

The configuration that determines how Amazon AppFlow formats the flow output data
when Upsolver is used as the destination.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AggregationConfig" : AggregationConfig,
  "FileType" : String,
  "PrefixConfig" : PrefixConfig
}

```

### YAML

```yaml

  AggregationConfig:
    AggregationConfig
  FileType: String
  PrefixConfig:
    PrefixConfig

```

## Properties

`AggregationConfig`

The aggregation settings that you can use to customize the output format of your flow
data.

_Required_: No

_Type_: [AggregationConfig](aws-properties-appflow-flow-aggregationconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FileType`

Indicates the file type that Amazon AppFlow places in the Upsolver Amazon S3
bucket.

_Required_: No

_Type_: String

_Allowed values_: `CSV | JSON | PARQUET`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrefixConfig`

Specifies elements that Amazon AppFlow includes in the file and folder names in the flow
destination.

_Required_: Yes

_Type_: [PrefixConfig](aws-properties-appflow-flow-prefixconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [UpsolverS3OutputFormatConfig](../../../../reference/appflow/1-0/apireference/api-upsolvers3outputformatconfig.md) in the _Amazon AppFlow API_
_Reference_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UpsolverDestinationProperties

VeevaSourceProperties

All content copied from https://docs.aws.amazon.com/.
