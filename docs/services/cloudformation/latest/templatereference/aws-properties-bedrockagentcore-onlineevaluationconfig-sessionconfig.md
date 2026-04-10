This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::OnlineEvaluationConfig SessionConfig

The session configuration that defines timeout settings for detecting when agent sessions are complete and ready for evaluation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SessionTimeoutMinutes" : Integer
}

```

### YAML

```yaml

  SessionTimeoutMinutes: Integer

```

## Properties

`SessionTimeoutMinutes`

The number of minutes of inactivity after which an agent session is considered complete and ready for evaluation.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `1440`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SamplingConfig

Tag

All content copied from https://docs.aws.amazon.com/.
