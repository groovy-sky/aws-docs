This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Transfer::Workflow InputFileLocation

Specifies the location for the file that's being processed.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EfsFileLocation" : EfsInputFileLocation,
  "S3FileLocation" : S3InputFileLocation
}

```

### YAML

```yaml

  EfsFileLocation:
    EfsInputFileLocation
  S3FileLocation:
    S3InputFileLocation

```

## Properties

`EfsFileLocation`

Specifies the details for the Amazon Elastic File System (Amazon EFS) file that's
being decrypted.

_Required_: No

_Type_: [EfsInputFileLocation](aws-properties-transfer-workflow-efsinputfilelocation.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`S3FileLocation`

Specifies the details for the Amazon S3 file that's being copied or decrypted.

_Required_: No

_Type_: [S3InputFileLocation](aws-properties-transfer-workflow-s3inputfilelocation.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EfsInputFileLocation

S3FileLocation

All content copied from https://docs.aws.amazon.com/.
