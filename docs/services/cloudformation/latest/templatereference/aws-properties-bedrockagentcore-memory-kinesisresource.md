This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::Memory KinesisResource

The `KinesisResource` property type specifies Property description not available. for an [AWS::BedrockAgentCore::Memory](aws-resource-bedrockagentcore-memory.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ContentConfigurations" : [ ContentConfiguration, ... ],
  "DataStreamArn" : String
}

```

### YAML

```yaml

  ContentConfigurations:
    - ContentConfiguration
  DataStreamArn: String

```

## Properties

`ContentConfigurations`

Property description not available.

_Required_: Yes

_Type_: Array of [ContentConfiguration](aws-properties-bedrockagentcore-memory-contentconfiguration.md)

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataStreamArn`

Property description not available.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:(aws(?:-cn|-us-gov|-iso(?:-[bef])?)?):[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[^/].{0,1023}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InvocationConfigurationInput

MemoryStrategy

All content copied from https://docs.aws.amazon.com/.
