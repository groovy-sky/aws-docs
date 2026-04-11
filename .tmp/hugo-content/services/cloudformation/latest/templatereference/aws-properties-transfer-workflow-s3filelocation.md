This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Transfer::Workflow S3FileLocation

Specifies the S3 details for the file being used, such as bucket, ETag, and so
forth.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "S3FileLocation" : S3InputFileLocation
}

```

### YAML

```yaml

  S3FileLocation:
    S3InputFileLocation

```

## Properties

`S3FileLocation`

Specifies the details for the file location for the file that's being used in the
workflow. Only applicable if you are using Amazon S3 storage.

_Required_: No

_Type_: [S3InputFileLocation](aws-properties-transfer-workflow-s3inputfilelocation.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InputFileLocation

S3InputFileLocation

All content copied from https://docs.aws.amazon.com/.
