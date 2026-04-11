This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DynamoDB::Table ImportSourceSpecification

Specifies the properties of data being imported from the S3 bucket source to the
table.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InputCompressionType" : String,
  "InputFormat" : String,
  "InputFormatOptions" : InputFormatOptions,
  "S3BucketSource" : S3BucketSource
}

```

### YAML

```yaml

  InputCompressionType: String
  InputFormat: String
  InputFormatOptions:
    InputFormatOptions
  S3BucketSource:
    S3BucketSource

```

## Properties

`InputCompressionType`

Type of compression to be used on the input coming from the imported table.

_Required_: No

_Type_: String

_Allowed values_: `GZIP | ZSTD | NONE`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InputFormat`

The format of the source data. Valid values for `ImportFormat` are
`CSV`, `DYNAMODB_JSON` or `ION`.

_Required_: Yes

_Type_: String

_Allowed values_: `DYNAMODB_JSON | ION | CSV`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InputFormatOptions`

Additional properties that specify how the input is formatted,

_Required_: No

_Type_: [InputFormatOptions](aws-properties-dynamodb-table-inputformatoptions.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`S3BucketSource`

The S3 bucket that provides the source for the import.

_Required_: Yes

_Type_: [S3BucketSource](aws-properties-dynamodb-table-s3bucketsource.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GlobalSecondaryIndex

InputFormatOptions

All content copied from https://docs.aws.amazon.com/.
