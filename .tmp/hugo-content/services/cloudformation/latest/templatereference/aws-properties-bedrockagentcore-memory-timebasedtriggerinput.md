This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::Memory TimeBasedTriggerInput

The memory trigger condition input for the time based trigger.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IdleSessionTimeout" : Integer
}

```

### YAML

```yaml

  IdleSessionTimeout: Integer

```

## Properties

`IdleSessionTimeout`

The memory trigger condition input for the session timeout.

_Required_: No

_Type_: Integer

_Minimum_: `10`

_Maximum_: `3000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SummaryOverrideConsolidationConfigurationInput

TokenBasedTriggerInput

All content copied from https://docs.aws.amazon.com/.
