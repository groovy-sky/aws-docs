This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataSource S3Location

A storage location in an Amazon S3 bucket.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "URI" : String
}

```

### YAML

```yaml

  URI: String

```

## Properties

`URI`

An object URI starting with `s3://`.

_Required_: Yes

_Type_: String

_Pattern_: `^s3://.{1,128}$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3DataSourceConfiguration

SalesforceCrawlerConfiguration

All content copied from https://docs.aws.amazon.com/.
