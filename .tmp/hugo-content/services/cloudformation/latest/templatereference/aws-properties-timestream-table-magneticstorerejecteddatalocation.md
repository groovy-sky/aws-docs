This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Timestream::Table MagneticStoreRejectedDataLocation

The location to write error reports for records rejected, asynchronously, during magnetic store writes.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "S3Configuration" : S3Configuration
}

```

### YAML

```yaml

  S3Configuration:
    S3Configuration

```

## Properties

`S3Configuration`

Configuration of an S3 location to write error reports for records rejected, asynchronously, during magnetic
store writes.

_Required_: No

_Type_: [S3Configuration](aws-properties-timestream-table-s3configuration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Timestream::Table

MagneticStoreWriteProperties

All content copied from https://docs.aws.amazon.com/.
