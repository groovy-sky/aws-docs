This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Omics::SequenceStore SseConfig

Server-side encryption (SSE) settings for a store.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "KeyArn" : String,
  "Type" : String
}

```

### YAML

```yaml

  KeyArn: String
  Type: String

```

## Properties

`KeyArn`

An encryption key ARN.

_Required_: No

_Type_: String

_Pattern_: `arn:([^:
]*):([^:
]*):([^:
]*):([0-9]{12}):([^:
]*)`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Type`

The encryption type.

_Required_: Yes

_Type_: String

_Allowed values_: `KMS`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Omics::SequenceStore

AWS::Omics::VariantStore

All content copied from https://docs.aws.amazon.com/.
