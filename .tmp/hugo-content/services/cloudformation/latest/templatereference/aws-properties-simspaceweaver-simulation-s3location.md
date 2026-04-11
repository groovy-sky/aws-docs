This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SimSpaceWeaver::Simulation S3Location

A location in Amazon Simple Storage Service (Amazon S3) where SimSpace Weaver stores simulation data, such as your app .zip
files and schema file. For more information about Amazon S3, see the [_Amazon Simple Storage Service User Guide_](../../../s3/latest/userguide/welcome.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BucketName" : String,
  "ObjectKey" : String
}

```

### YAML

```yaml

  BucketName: String
  ObjectKey: String

```

## Properties

`BucketName`

The name of an Amazon S3 bucket. For more information about buckets, see [Creating,\
configuring, and working with Amazon S3 buckets](../../../s3/latest/userguide/creating-buckets-s3.md) in the _Amazon Simple Storage Service User_
_Guide_.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9_\-]{3,63}$`

_Minimum_: `3`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ObjectKey`

The key name of an object in Amazon S3. For more information about Amazon S3 objects and object
keys, see [Uploading,\
downloading, and working with objects in Amazon S3](../../../s3/latest/userguide/uploading-downloading-objects.md) in the _Amazon Simple Storage Service User_
_Guide_.

_Required_: Yes

_Type_: String

_Minimum_: `3`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::SimSpaceWeaver::Simulation

Next

All content copied from https://docs.aws.amazon.com/.
