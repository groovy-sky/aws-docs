This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSMGuiConnect::Preferences RecordingDestinations

Determines where recordings of RDP connections are stored.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "S3Buckets" : [ S3Bucket, ... ]
}

```

### YAML

```yaml

  S3Buckets:
    - S3Bucket

```

## Properties

`S3Buckets`

The S3 bucket where RDP connection recordings are stored.

_Required_: Yes

_Type_: Array of [S3Bucket](aws-properties-ssmguiconnect-preferences-s3bucket.md)

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConnectionRecordingPreferences

S3Bucket

All content copied from https://docs.aws.amazon.com/.
