This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EKS::Nodegroup NodeRepairConfigOverrides

Specify granular overrides for specific repair actions. These overrides control the
repair action and the repair delay time before a node is considered eligible for repair.
If you use this, you must specify all the values.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MinRepairWaitTimeMins" : Integer,
  "NodeMonitoringCondition" : String,
  "NodeUnhealthyReason" : String,
  "RepairAction" : String
}

```

### YAML

```yaml

  MinRepairWaitTimeMins: Integer
  NodeMonitoringCondition: String
  NodeUnhealthyReason: String
  RepairAction: String

```

## Properties

`MinRepairWaitTimeMins`

Specify the minimum time in minutes to wait before attempting to repair a node
with this specific `nodeMonitoringCondition` and
`nodeUnhealthyReason`.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NodeMonitoringCondition`

Specify an unhealthy condition reported by the node monitoring agent that this
override would apply to.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NodeUnhealthyReason`

Specify a reason reported by the node monitoring agent that this
override would apply to.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RepairAction`

Specify the repair action to take for nodes when all of the specified conditions are
met.

_Required_: No

_Type_: String

_Allowed values_: `Replace | Reboot | NoAction`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NodeRepairConfig

RemoteAccess

All content copied from https://docs.aws.amazon.com/.
