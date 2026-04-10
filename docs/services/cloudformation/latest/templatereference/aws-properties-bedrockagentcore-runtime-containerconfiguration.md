This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::Runtime ContainerConfiguration

Representation of a container configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ContainerUri" : String
}

```

### YAML

```yaml

  ContainerUri: String

```

## Properties

`ContainerUri`

The ECR URI of the container.

_Required_: Yes

_Type_: String

_Pattern_: `^\d{12}\.dkr\.ecr\.([a-z0-9-]+)\.amazonaws\.com/((?:[a-z0-9]+(?:[._-][a-z0-9]+)*/)*[a-z0-9]+(?:[._-][a-z0-9]+)*)([:@]\S+)$`

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CodeConfiguration

CustomClaimValidationType

All content copied from https://docs.aws.amazon.com/.
