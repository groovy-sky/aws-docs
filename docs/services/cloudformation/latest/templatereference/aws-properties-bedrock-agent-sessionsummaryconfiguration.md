This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Agent SessionSummaryConfiguration

Configuration for SESSION\_SUMMARY memory type enabled for the agent.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaxRecentSessions" : Number
}

```

### YAML

```yaml

  MaxRecentSessions: Number

```

## Properties

`MaxRecentSessions`

Maximum number of recent session summaries to include in the agent's prompt context.

_Required_: No

_Type_: Number

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3Identifier

AWS::Bedrock::AgentAlias

All content copied from https://docs.aws.amazon.com/.
