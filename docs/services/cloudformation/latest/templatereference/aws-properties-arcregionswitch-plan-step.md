This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ARCRegionSwitch::Plan Step

Represents a step in a Region switch plan workflow. Each step performs a specific action during the Region switch process.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Description" : String,
  "ExecutionBlockConfiguration" : ExecutionBlockConfiguration,
  "ExecutionBlockType" : String,
  "Name" : String
}

```

### YAML

```yaml

  Description: String
  ExecutionBlockConfiguration:
    ExecutionBlockConfiguration
  ExecutionBlockType: String
  Name: String

```

## Properties

`Description`

The description of a step in a workflow.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExecutionBlockConfiguration`

The configuration for an execution block in a workflow.

_Required_: Yes

_Type_: [ExecutionBlockConfiguration](aws-properties-arcregionswitch-plan-executionblockconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExecutionBlockType`

The type of an execution block in a workflow.

_Required_: Yes

_Type_: String

_Allowed values_: `CustomActionLambda | ManualApproval | AuroraGlobalDatabase | EC2AutoScaling | ARCRoutingControl | ARCRegionSwitchPlan | Parallel | ECSServiceScaling | EKSResourceScaling | Route53HealthCheck | DocumentDb | RdsPromoteReadReplica | RdsCreateCrossRegionReplica`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of a step in a workflow.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Service

Trigger

All content copied from https://docs.aws.amazon.com/.
