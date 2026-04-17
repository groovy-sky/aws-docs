---
title: "AWS::BedrockAgentCore::Memory ContentConfiguration"
---

This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::Memory ContentConfiguration

Defines what content to stream and at what level of detail.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Level" : String,
  "Type" : String
}

```

### YAML

```yaml

  Level: String
  Type: String

```

## Properties

`Level`

Level of detail for streamed content.

_Required_: No

_Type_: String

_Allowed values_: `METADATA_ONLY | FULL_CONTENT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

Type of content to stream.

_Required_: Yes

_Type_: String

_Allowed values_: `MEMORY_RECORDS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::BedrockAgentCore::Memory

CustomConfigurationInput

All content copied from https://docs.aws.amazon.com/.
