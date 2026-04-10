This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::Runtime AgentRuntimeArtifact

The artifact of the agent.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CodeConfiguration" : CodeConfiguration,
  "ContainerConfiguration" : ContainerConfiguration
}

```

### YAML

```yaml

  CodeConfiguration:
    CodeConfiguration
  ContainerConfiguration:
    ContainerConfiguration

```

## Properties

`CodeConfiguration`

The code configuration for the agent runtime artifact, including the source code location and execution settings.

_Required_: No

_Type_: [CodeConfiguration](aws-properties-bedrockagentcore-runtime-codeconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ContainerConfiguration`

The container configuration for the agent artifact.

_Required_: No

_Type_: [ContainerConfiguration](aws-properties-bedrockagentcore-runtime-containerconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::BedrockAgentCore::Runtime

AuthorizerConfiguration

All content copied from https://docs.aws.amazon.com/.
