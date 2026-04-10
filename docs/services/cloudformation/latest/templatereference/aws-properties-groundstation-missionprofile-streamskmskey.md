This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GroundStation::MissionProfile StreamsKmsKey

KMS key info.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "KmsAliasArn" : String,
  "KmsAliasName" : String,
  "KmsKeyArn" : String
}

```

### YAML

```yaml

  KmsAliasArn: String
  KmsAliasName: String
  KmsKeyArn: String

```

## Properties

`KmsAliasArn`

KMS Alias Arn.

_Required_: No

_Type_: String

_Pattern_: `^(arn:(aws[a-zA-Z-]*)?:[a-z0-9-.]+:.*)|()$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsAliasName`

KMS Alias Name.

_Required_: No

_Type_: String

_Pattern_: `^alias/[a-zA-Z0-9:/_-]+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyArn`

KMS Key Arn.

_Required_: No

_Type_: String

_Pattern_: `^(arn:(aws[a-zA-Z-]*)?:[a-z0-9-.]+:.*)|()$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataflowEdge

Tag

All content copied from https://docs.aws.amazon.com/.
