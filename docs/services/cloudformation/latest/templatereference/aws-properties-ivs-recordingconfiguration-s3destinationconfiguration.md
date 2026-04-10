This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IVS::RecordingConfiguration S3DestinationConfiguration

The S3DestinationConfiguration property type describes an S3 location where recorded
videos will be stored.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BucketName" : String
}

```

### YAML

```yaml

  BucketName: String

```

## Properties

`BucketName`

Location (S3 bucket name) where recorded videos will be stored.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-z0-9-.]+$`

_Minimum_: `3`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RenditionConfiguration

Tag

All content copied from https://docs.aws.amazon.com/.
