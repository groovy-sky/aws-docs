This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRooms::AnalysisTemplate S3Location

The S3 location.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Bucket" : String,
  "Key" : String
}

```

### YAML

```yaml

  Bucket: String
  Key: String

```

## Properties

`Bucket`

The bucket name.

_Required_: Yes

_Type_: String

_Minimum_: `3`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Key`

The object key.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9!_.*'()-/]+`

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MLSyntheticDataParameters

SyntheticDataColumnProperties

All content copied from https://docs.aws.amazon.com/.
