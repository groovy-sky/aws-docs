This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::AssistantAssociation ExternalBedrockKnowledgeBaseConfig

Configuration for an external Bedrock knowledge base.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccessRoleArn" : String,
  "BedrockKnowledgeBaseArn" : String
}

```

### YAML

```yaml

  AccessRoleArn: String
  BedrockKnowledgeBaseArn: String

```

## Properties

`AccessRoleArn`

The Amazon Resource Name (ARN) of the IAM role used to access the external Bedrock knowledge base.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws:iam::[0-9]{12}:role/(?:service-role/)?[a-zA-Z0-9_+=,.@-]{1,64}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`BedrockKnowledgeBaseArn`

The Amazon Resource Name (ARN) of the external Bedrock knowledge base.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws(|-cn|-us-gov):bedrock:[a-zA-Z0-9-]*:[0-9]{12}:knowledge-base/[0-9a-zA-Z]+$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AssociationData

Tag

All content copied from https://docs.aws.amazon.com/.
