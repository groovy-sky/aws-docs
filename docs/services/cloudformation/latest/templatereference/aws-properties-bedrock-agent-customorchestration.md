This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Agent CustomOrchestration

Contains details of the custom orchestration configured for the agent.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Executor" : OrchestrationExecutor
}

```

### YAML

```yaml

  Executor:
    OrchestrationExecutor

```

## Properties

`Executor`

The structure of the executor invoking the actions in custom orchestration.

_Required_: No

_Type_: [OrchestrationExecutor](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-agent-orchestrationexecutor.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

APISchema

Function
