This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Omics::VariantStore ReferenceItem

The read set's genome reference ARN.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ReferenceArn" : String
}

```

### YAML

```yaml

  ReferenceArn: String

```

## Properties

`ReferenceArn`

The reference's ARN.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:.+$`

_Minimum_: `1`

_Maximum_: `127`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Omics::VariantStore

SseConfig
