This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CustomerProfiles::Domain S3ExportingConfig

The S3 location where Identity Resolution Jobs write result files.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "S3BucketName" : String,
  "S3KeyName" : String
}

```

### YAML

```yaml

  S3BucketName: String
  S3KeyName: String

```

## Properties

`S3BucketName`

The name of the S3 bucket where Identity Resolution Jobs write result files.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-z0-9.-]+$`

_Minimum_: `3`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3KeyName`

The S3 key name of the location where Identity Resolution Jobs write result
files.

_Required_: No

_Type_: String

_Pattern_: `.*`

_Minimum_: `1`

_Maximum_: `800`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RuleBasedMatching

Tag

All content copied from https://docs.aws.amazon.com/.
