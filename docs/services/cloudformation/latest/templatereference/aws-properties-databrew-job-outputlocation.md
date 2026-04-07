This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataBrew::Job OutputLocation

The location in Amazon S3 or AWS Glue Data Catalog where the job
writes its output.

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

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BucketOwner`

Property description not available.

_Required_: No

_Type_: String

_Minimum_: `12`

_Maximum_: `12`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Key`

The unique name of the object in the bucket.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

OutputFormatOptions

ProfileConfiguration
