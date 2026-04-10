This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot GrammarSlotTypeSource

Describes the Amazon S3 bucket name and location for the grammar
that is the source for the slot type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "KmsKeyArn" : String,
  "S3BucketName" : String,
  "S3ObjectKey" : String
}

```

### YAML

```yaml

  KmsKeyArn: String
  S3BucketName: String
  S3ObjectKey: String

```

## Properties

`KmsKeyArn`

The AWS KMS key required to decrypt the contents of the grammar, if any.

_Required_: No

_Type_: String

_Pattern_: `^arn:[\w\-]+:kms:[\w\-]+:[\d]{12}:(?:key\/[\w\-]+|alias\/[a-zA-Z0-9:\/_\-]{1,256})$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3BucketName`

The name of the Amazon S3 bucket that contains the grammar source.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-z0-9][\.\-a-z0-9]{1,61}[a-z0-9]$`

_Minimum_: `3`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3ObjectKey`

The path to the grammar in the Amazon S3 bucket.

_Required_: Yes

_Type_: String

_Pattern_: `[\.\-\!\*\_\'\(\)a-zA-Z0-9][\.\-\!\*\_\'\(\)\/a-zA-Z0-9]*$`

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GrammarSlotTypeSetting

ImageResponseCard

All content copied from https://docs.aws.amazon.com/.
