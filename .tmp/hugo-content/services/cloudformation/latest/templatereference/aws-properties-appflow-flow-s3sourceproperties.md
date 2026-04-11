This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::Flow S3SourceProperties

The properties that are applied when Amazon S3 is being used as the flow source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BucketName" : String,
  "BucketPrefix" : String,
  "S3InputFormatConfig" : S3InputFormatConfig
}

```

### YAML

```yaml

  BucketName: String
  BucketPrefix: String
  S3InputFormatConfig:
    S3InputFormatConfig

```

## Properties

`BucketName`

The Amazon S3 bucket name where the source files are stored.

_Required_: Yes

_Type_: String

_Pattern_: `\S+`

_Minimum_: `3`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BucketPrefix`

The object key for the Amazon S3 bucket in which the source files are stored.

_Required_: Yes

_Type_: String

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3InputFormatConfig`

When you use Amazon S3 as the source, the configuration format that you provide
the flow input data.

_Required_: No

_Type_: [S3InputFormatConfig](aws-properties-appflow-flow-s3inputformatconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [S3SourceProperties](../../../../reference/appflow/1-0/apireference/api-s3sourceproperties.md)
in the _Amazon AppFlow API Reference_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3OutputFormatConfig

SalesforceDestinationProperties

All content copied from https://docs.aws.amazon.com/.
