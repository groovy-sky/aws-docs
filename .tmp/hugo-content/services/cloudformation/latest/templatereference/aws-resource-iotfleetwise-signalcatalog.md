This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTFleetWise::SignalCatalog

Creates a collection of standardized signals that can be reused to create vehicle models.

For more information, see [Signal catalogs](../../../iot-fleetwise/latest/developerguide/signal-catalogs.md) in the _AWS IoT FleetWise Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoTFleetWise::SignalCatalog",
  "Properties" : {
      "Description" : String,
      "Name" : String,
      "NodeCounts" : NodeCounts,
      "Nodes" : [ Node, ... ],
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::IoTFleetWise::SignalCatalog
Properties:
  Description: String
  Name: String
  NodeCounts:
    NodeCounts
  Nodes:
    - Node
  Tags:
    - Tag

```

## Properties

`Description`

A brief description of the signal catalog.

_Required_: No

_Type_: String

_Pattern_: `^[^\u0000-\u001F\u007F]+$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the signal catalog.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z\d\-_:]+$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NodeCounts`

Information about the number of nodes and node types in a vehicle network.

_Required_: No

_Type_: [NodeCounts](aws-properties-iotfleetwise-signalcatalog-nodecounts.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Nodes`

A list of information about nodes, which are a general abstraction of signals.

_Required_: No

_Type_: Array of [Node](aws-properties-iotfleetwise-signalcatalog-node.md)

_Minimum_: `1`

_Maximum_: `5000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Metadata that can be used to manage the signal catalog.

_Required_: No

_Type_: Array of [Tag](aws-properties-iotfleetwise-signalcatalog-tag.md)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the signal catalog.

`CreationTime`

The time the signal catalog was created in seconds since epoch (January 1, 1970 at midnight UTC time).

`LastModificationTime`

The time the signal catalog was last updated in seconds since epoch (January 1, 1970 at midnight UTC time).

`NodeCounts.TotalActuators`

The total number of nodes in a vehicle network that represent actuators.

`NodeCounts.TotalAttributes`

The total number of nodes in a vehicle network that represent attributes.

`NodeCounts.TotalBranches`

The total number of nodes in a vehicle network that represent branches.

`NodeCounts.TotalNodes`

The total number of nodes in a vehicle network.

`NodeCounts.TotalSensors`

The total number of nodes in a vehicle network that represent sensors.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Actuator

All content copied from https://docs.aws.amazon.com/.
