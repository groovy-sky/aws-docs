This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::Connection S3PropertiesInput

The Amazon S3 properties of a connection.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "S3AccessGrantLocationId" : String,
  "S3Uri" : String
}

```

### YAML

```yaml

  S3AccessGrantLocationId: String
  S3Uri: String

```

## Properties

`S3AccessGrantLocationId`

The Amazon S3 Access Grant location ID that's part of the Amazon S3 properties of a
connection.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9\-]+`

_Minimum_: `0`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3Uri`

The Amazon S3 URI that's part of the Amazon S3 properties of a connection.

_Required_: Yes

_Type_: String

_Pattern_: `s3://.+`

_Minimum_: `0`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RedshiftStorageProperties

SparkEmrPropertiesInput

All content copied from https://docs.aws.amazon.com/.
