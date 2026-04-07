This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::Runtime LifecycleConfiguration

The lifecycle configuration for the AgentCore Runtime.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IdleRuntimeSessionTimeout" : Integer,
  "MaxLifetime" : Integer
}

```

### YAML

```yaml

  IdleRuntimeSessionTimeout: Integer
  MaxLifetime: Integer

```

## Properties

`IdleRuntimeSessionTimeout`

The idle session timeout for the AgentCore Runtime.

_Required_: No

_Type_: Integer

_Minimum_: `60`

_Maximum_: `28800`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxLifetime`

The maximum lifetime for the AgentCore Runtime.

_Required_: No

_Type_: Integer

_Minimum_: `60`

_Maximum_: `28800`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

FilesystemConfiguration

NetworkConfiguration
