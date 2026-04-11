This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::Flow SuccessResponseHandlingConfig

Determines how Amazon AppFlow handles the success response that it gets from the
connector after placing data.

For example, this setting would determine where to write the response from the destination
connector upon a successful insert operation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BucketName" : String,
  "BucketPrefix" : String
}

```

### YAML

```yaml

  BucketName: String
  BucketPrefix: String

```

## Properties

`BucketName`

The name of the Amazon S3 bucket.

_Required_: No

_Type_: String

_Pattern_: `\S+`

_Minimum_: `3`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BucketPrefix`

The Amazon S3 bucket prefix.

_Required_: No

_Type_: String

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SourceFlowConfig

Tag

All content copied from https://docs.aws.amazon.com/.
