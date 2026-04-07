This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket DataExport

Specifies how data related to the storage class analysis for an Amazon S3 bucket should be
exported.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Destination" : Destination,
  "OutputSchemaVersion" : String
}

```

### YAML

```yaml

  Destination:
    Destination
  OutputSchemaVersion: String

```

## Properties

`Destination`

The place to store the data for an analysis.

_Required_: Yes

_Type_: [Destination](aws-properties-s3-bucket-destination.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutputSchemaVersion`

The version of the output schema to use when exporting data. Must be `V_1`.

_Required_: Yes

_Type_: String

_Allowed values_: `V_1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- AWS::S3::Bucket [Examples](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#aws-properties-s3-bucket--examples)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CorsRule

DefaultRetention
