This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Deadline::Queue JobAttachmentSettings

The job attachment settings. These are the Amazon S3 bucket name and the Amazon S3 prefix.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RootPrefix" : String,
  "S3BucketName" : String
}

```

### YAML

```yaml

  RootPrefix: String
  S3BucketName: String

```

## Properties

`RootPrefix`

The root prefix.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3BucketName`

The Amazon S3 bucket name.

_Required_: Yes

_Type_: String

_Pattern_: `(?!^(\d+\.)+\d+$)(^(([a-z0-9]|[a-z0-9][a-z0-9\-]*[a-z0-9])\.)*([a-z0-9]|[a-z0-9][a-z0-9\-]*[a-z0-9])$)`

_Minimum_: `3`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Deadline::Queue

JobRunAsUser

All content copied from https://docs.aws.amazon.com/.
