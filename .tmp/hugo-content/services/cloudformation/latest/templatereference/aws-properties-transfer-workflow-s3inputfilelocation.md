This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Transfer::Workflow S3InputFileLocation

Specifies the details for the Amazon S3 location for an input file to a
workflow.

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

Specifies the S3 bucket for the customer input file.

_Required_: No

_Type_: String

_Pattern_: `^[a-z0-9][\.\-a-z0-9]{1,61}[a-z0-9]$`

_Minimum_: `3`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Key`

The name assigned to the file when it was created in Amazon S3. You use the object key to retrieve the object.

_Required_: No

_Type_: String

_Pattern_: `.*`

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3FileLocation

S3Tag

All content copied from https://docs.aws.amazon.com/.
