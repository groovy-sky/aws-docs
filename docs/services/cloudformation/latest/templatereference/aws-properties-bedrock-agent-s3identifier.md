This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Agent S3Identifier

The identifier information for an Amazon S3 bucket.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "S3BucketName" : String,
  "S3ObjectKey" : String
}

```

### YAML

```yaml

  S3BucketName: String
  S3ObjectKey: String

```

## Properties

`S3BucketName`

The name of the S3 bucket.

_Required_: No

_Type_: String

_Pattern_: `^[a-z0-9][\.\-a-z0-9]{1,61}[a-z0-9]$`

_Minimum_: `3`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3ObjectKey`

The S3 object key for the S3 resource.

_Required_: No

_Type_: String

_Pattern_: `^[\.\-\!\*\_\'\(\)a-zA-Z0-9][\.\-\!\*\_\'\(\)\/a-zA-Z0-9]*$`

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PromptOverrideConfiguration

SessionSummaryConfiguration

All content copied from https://docs.aws.amazon.com/.
