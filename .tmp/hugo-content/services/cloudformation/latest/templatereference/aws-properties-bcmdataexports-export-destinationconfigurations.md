This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BCMDataExports::Export DestinationConfigurations

The destinations used for data exports.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "S3Destination" : S3Destination
}

```

### YAML

```yaml

  S3Destination:
    S3Destination

```

## Properties

`S3Destination`

An object that describes the destination of the data exports file.

_Required_: Yes

_Type_: [S3Destination](aws-properties-bcmdataexports-export-s3destination.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataQuery

Export

All content copied from https://docs.aws.amazon.com/.
