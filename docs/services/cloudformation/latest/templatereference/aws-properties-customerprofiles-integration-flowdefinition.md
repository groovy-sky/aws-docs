This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CustomerProfiles::Integration FlowDefinition

The configurations that control how Customer Profiles retrieves data from the source,
Amazon AppFlow. Customer Profiles uses this information to create an AppFlow flow on
behalf of customers.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Description" : String,
  "FlowName" : String,
  "KmsArn" : String,
  "SourceFlowConfig" : SourceFlowConfig,
  "Tasks" : [ Task, ... ],
  "TriggerConfig" : TriggerConfig
}

```

### YAML

```yaml

  Description: String
  FlowName: String
  KmsArn: String
  SourceFlowConfig:
    SourceFlowConfig
  Tasks:
    - Task
  TriggerConfig:
    TriggerConfig

```

## Properties

`Description`

A description of the flow you want to create.

_Required_: No

_Type_: String

_Pattern_: `[\w!@#\-.?,\s]*`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FlowName`

The specified name of the flow. Use underscores (\_) or hyphens (-) only. Spaces are
not allowed.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9][\w!@#.-]+`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsArn`

The Amazon Resource Name (ARN) of the AWS Key Management Service (KMS) key you provide
for encryption.

_Required_: Yes

_Type_: String

_Pattern_: `arn:aws:kms:.*:[0-9]+:.*`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceFlowConfig`

The configuration that controls how Customer Profiles retrieves data from the
source.

_Required_: Yes

_Type_: [SourceFlowConfig](aws-properties-customerprofiles-integration-sourceflowconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tasks`

A list of tasks that Customer Profiles performs while transferring the data in the
flow run.

_Required_: Yes

_Type_: Array of [Task](aws-properties-customerprofiles-integration-task.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TriggerConfig`

The trigger settings that determine how and when the flow runs.

_Required_: Yes

_Type_: [TriggerConfig](aws-properties-customerprofiles-integration-triggerconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConnectorOperator

IncrementalPullConfig

All content copied from https://docs.aws.amazon.com/.
