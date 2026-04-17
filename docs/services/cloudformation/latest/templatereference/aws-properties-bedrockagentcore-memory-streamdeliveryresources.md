---
title: "AWS::BedrockAgentCore::Memory StreamDeliveryResources"
---

This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::Memory StreamDeliveryResources

Configuration for streaming memory record data to external resources.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Resources" : [ StreamDeliveryResource, ... ]
}

```

### YAML

```yaml

  Resources:
    - StreamDeliveryResource

```

## Properties

`Resources`

List of stream delivery resource configurations.

_Required_: Yes

_Type_: Array of [StreamDeliveryResource](aws-properties-bedrockagentcore-memory-streamdeliveryresource.md)

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StreamDeliveryResource

SummaryMemoryStrategy

All content copied from https://docs.aws.amazon.com/.
