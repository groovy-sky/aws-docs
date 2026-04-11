This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::Memory CustomConfigurationInput

The memory configuration input.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EpisodicOverride" : EpisodicOverride,
  "SelfManagedConfiguration" : SelfManagedConfiguration,
  "SemanticOverride" : SemanticOverride,
  "SummaryOverride" : SummaryOverride,
  "UserPreferenceOverride" : UserPreferenceOverride
}

```

### YAML

```yaml

  EpisodicOverride:
    EpisodicOverride
  SelfManagedConfiguration:
    SelfManagedConfiguration
  SemanticOverride:
    SemanticOverride
  SummaryOverride:
    SummaryOverride
  UserPreferenceOverride:
    UserPreferenceOverride

```

## Properties

`EpisodicOverride`

The episodic memory strategy override configuration for a custom memory strategy.

_Required_: No

_Type_: [EpisodicOverride](aws-properties-bedrockagentcore-memory-episodicoverride.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SelfManagedConfiguration`

The custom configuration input.

_Required_: No

_Type_: [SelfManagedConfiguration](aws-properties-bedrockagentcore-memory-selfmanagedconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SemanticOverride`

The memory override configuration.

_Required_: No

_Type_: [SemanticOverride](aws-properties-bedrockagentcore-memory-semanticoverride.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SummaryOverride`

The memory configuration override.

_Required_: No

_Type_: [SummaryOverride](aws-properties-bedrockagentcore-memory-summaryoverride.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserPreferenceOverride`

The memory user preference override.

_Required_: No

_Type_: [UserPreferenceOverride](aws-properties-bedrockagentcore-memory-userpreferenceoverride.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ContentConfiguration

CustomMemoryStrategy

All content copied from https://docs.aws.amazon.com/.
