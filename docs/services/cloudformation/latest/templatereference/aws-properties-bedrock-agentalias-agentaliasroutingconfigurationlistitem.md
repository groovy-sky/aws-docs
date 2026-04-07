This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::AgentAlias AgentAliasRoutingConfigurationListItem

Contains details about the routing configuration of the alias.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AgentVersion" : String
}

```

### YAML

```yaml

  AgentVersion: String

```

## Properties

`AgentVersion`

The version of the agent with which the alias is associated.

_Required_: Yes

_Type_: String

_Pattern_: `^(DRAFT|[0-9]{0,4}[1-9][0-9]{0,4})$`

_Minimum_: `1`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AgentAliasHistoryEvent

AWS::Bedrock::ApplicationInferenceProfile
