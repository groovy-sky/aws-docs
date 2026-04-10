This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Agent MemoryConfiguration

Details of the memory configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EnabledMemoryTypes" : [ String, ... ],
  "SessionSummaryConfiguration" : SessionSummaryConfiguration,
  "StorageDays" : Number
}

```

### YAML

```yaml

  EnabledMemoryTypes:
    - String
  SessionSummaryConfiguration:
    SessionSummaryConfiguration
  StorageDays: Number

```

## Properties

`EnabledMemoryTypes`

The type of memory that is stored.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SessionSummaryConfiguration`

Contains the configuration for SESSION\_SUMMARY memory type enabled for the agent.

_Required_: No

_Type_: [SessionSummaryConfiguration](aws-properties-bedrock-agent-sessionsummaryconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StorageDays`

The number of days the agent is configured to retain the conversational context.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `365`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InferenceConfiguration

OrchestrationExecutor

All content copied from https://docs.aws.amazon.com/.
