This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::Memory TokenBasedTriggerInput

The token based trigger input.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TokenCount" : Integer
}

```

### YAML

```yaml

  TokenCount: Integer

```

## Properties

`TokenCount`

The token based trigger token count.

_Required_: No

_Type_: Integer

_Minimum_: `100`

_Maximum_: `500000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TimeBasedTriggerInput

TriggerConditionInput

All content copied from https://docs.aws.amazon.com/.
