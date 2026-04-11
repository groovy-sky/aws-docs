This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSMGuiConnect::Preferences S3Bucket

The S3 bucket where RDP connection recordings are stored.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BucketName" : String,
  "BucketOwner" : String
}

```

### YAML

```yaml

  BucketName: String
  BucketOwner: String

```

## Properties

`BucketName`

The name of the S3 bucket where RDP connection recordings are stored.

_Required_: Yes

_Type_: String

_Pattern_: `(?=^.{3,63}$)(?!^(\d+\.)+\d+$)(^(([a-z0-9]|[a-z0-9][a-z0-9\-]*[a-z0-9])\.)*([a-z0-9]|[a-z0-9][a-z0-9\-]*[a-z0-9])$)`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BucketOwner`

The AWS account number that owns the S3 bucket.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9]{12}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RecordingDestinations

Next

All content copied from https://docs.aws.amazon.com/.
