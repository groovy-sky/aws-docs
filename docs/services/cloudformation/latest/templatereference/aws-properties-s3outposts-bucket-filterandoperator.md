This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3Outposts::Bucket FilterAndOperator

The `FilterAndOperator` property type specifies Property description not available. for an [AWS::S3Outposts::Bucket](aws-resource-s3outposts-bucket.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Prefix" : String,
  "Tags" : [ FilterTag, ... ]
}

```

### YAML

```yaml

  Prefix: String
  Tags:
    - FilterTag

```

## Properties

`Prefix`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Property description not available.

_Required_: Yes

_Type_: Array of [FilterTag](aws-properties-s3outposts-bucket-filtertag.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Filter

FilterTag

All content copied from https://docs.aws.amazon.com/.
