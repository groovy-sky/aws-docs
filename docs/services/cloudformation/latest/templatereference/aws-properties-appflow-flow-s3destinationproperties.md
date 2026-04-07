This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::Flow S3DestinationProperties

The properties that are applied when Amazon S3 is used as a destination.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BucketName" : String,
  "BucketPrefix" : String,
  "S3OutputFormatConfig" : S3OutputFormatConfig
}

```

### YAML

```yaml

  BucketName: String
  BucketPrefix: String
  S3OutputFormatConfig:
    S3OutputFormatConfig

```

## Properties

`BucketName`

The Amazon S3 bucket name in which Amazon AppFlow places the transferred
data.

_Required_: Yes

_Type_: String

_Pattern_: `\S+`

_Minimum_: `3`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BucketPrefix`

The object key for the destination bucket in which Amazon AppFlow places the files.

_Required_: No

_Type_: String

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3OutputFormatConfig`

The configuration that determines how Amazon AppFlow should format the flow output
data when Amazon S3 is used as the destination.

_Required_: No

_Type_: [S3OutputFormatConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-flow-s3outputformatconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [S3DestinationProperties](../../../../reference/appflow/1-0/apireference/api-s3destinationproperties.md) in the _Amazon AppFlow API_
_Reference_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

RedshiftDestinationProperties

S3InputFormatConfig
