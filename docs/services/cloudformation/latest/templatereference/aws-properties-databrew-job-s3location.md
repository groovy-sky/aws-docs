This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataBrew::Job S3Location

Represents an Amazon S3 location (bucket name, bucket owner, and object key) where DataBrew can read
input data, or write output from a job.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Bucket" : String,
  "BucketOwner" : String,
  "Key" : String
}

```

### YAML

```yaml

  Bucket: String
  BucketOwner: String
  Key: String

```

## Properties

`Bucket`

The Amazon S3 bucket name.

_Required_: Yes

_Type_: String

_Minimum_: `3`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BucketOwner`

The AWS account ID of the bucket owner.

_Required_: No

_Type_: String

_Minimum_: `12`

_Maximum_: `12`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Key`

The unique name of the object in the bucket.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1280`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Recipe

S3TableOutputOptions

All content copied from https://docs.aws.amazon.com/.
