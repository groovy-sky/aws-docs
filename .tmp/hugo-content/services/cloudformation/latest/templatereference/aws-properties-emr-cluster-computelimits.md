This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMR::Cluster ComputeLimits

The Amazon EC2 unit limits for a managed scaling policy. The managed scaling
activity of a cluster can not be above or below these limits. The limit only applies to the
core and task nodes. The master node cannot be scaled after initial configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaximumCapacityUnits" : Integer,
  "MaximumCoreCapacityUnits" : Integer,
  "MaximumOnDemandCapacityUnits" : Integer,
  "MinimumCapacityUnits" : Integer,
  "UnitType" : String
}

```

### YAML

```yaml

  MaximumCapacityUnits: Integer
  MaximumCoreCapacityUnits: Integer
  MaximumOnDemandCapacityUnits: Integer
  MinimumCapacityUnits: Integer
  UnitType: String

```

## Properties

`MaximumCapacityUnits`

The upper boundary of Amazon EC2 units. It is measured through vCPU cores or
instances for instance groups and measured through units for instance fleets. Managed
scaling activities are not allowed beyond this boundary. The limit only applies to the core
and task nodes. The master node cannot be scaled after initial configuration.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaximumCoreCapacityUnits`

The upper boundary of Amazon EC2 units for core node type in a cluster. It is
measured through vCPU cores or instances for instance groups and measured through units for
instance fleets. The core units are not allowed to scale beyond this boundary. The
parameter is used to split capacity allocation between core and task nodes.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaximumOnDemandCapacityUnits`

The upper boundary of On-Demand Amazon EC2 units. It is measured through vCPU
cores or instances for instance groups and measured through units for instance fleets. The
On-Demand units are not allowed to scale beyond this boundary. The parameter is used to
split capacity allocation between On-Demand and Spot Instances.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinimumCapacityUnits`

The lower boundary of Amazon EC2 units. It is measured through vCPU cores or
instances for instance groups and measured through units for instance fleets. Managed
scaling activities are not allowed beyond this boundary. The limit only applies to the core
and task nodes. The master node cannot be scaled after initial configuration.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UnitType`

The unit type used for specifying a managed scaling policy.

_Required_: Yes

_Type_: String

_Allowed values_: `InstanceFleetUnits | Instances | VCPU`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CloudWatchAlarmDefinition

Configuration

All content copied from https://docs.aws.amazon.com/.
