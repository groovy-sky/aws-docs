This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTWireless::TaskDefinition

Creates a gateway task definition.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoTWireless::TaskDefinition",
  "Properties" : {
      "AutoCreateTasks" : Boolean,
      "LoRaWANUpdateGatewayTaskEntry" : LoRaWANUpdateGatewayTaskEntry,
      "Name" : String,
      "Tags" : [ Tag, ... ],
      "TaskDefinitionType" : String,
      "Update" : UpdateWirelessGatewayTaskCreate
    }
}

```

### YAML

```yaml

Type: AWS::IoTWireless::TaskDefinition
Properties:
  AutoCreateTasks: Boolean
  LoRaWANUpdateGatewayTaskEntry:
    LoRaWANUpdateGatewayTaskEntry
  Name: String
  Tags:
    - Tag
  TaskDefinitionType: String
  Update:
    UpdateWirelessGatewayTaskCreate

```

## Properties

`AutoCreateTasks`

Whether to automatically create tasks using this task definition for all gateways with
the specified current version. If `false`, the task must be created by calling
`CreateWirelessGatewayTask`.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LoRaWANUpdateGatewayTaskEntry`

LoRaWANUpdateGatewayTaskEntry object.

_Required_: No

_Type_: [LoRaWANUpdateGatewayTaskEntry](aws-properties-iotwireless-taskdefinition-lorawanupdategatewaytaskentry.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the new resource.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags are an array of key-value pairs to attach to the specified resource. Tags can
have a minimum of 0 and a maximum of 50 items.

_Required_: No

_Type_: Array of [Tag](aws-properties-iotwireless-taskdefinition-tag.md)

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TaskDefinitionType`

A filter to list only the wireless gateway task definitions that use this task
definition type.

_Required_: No

_Type_: String

_Allowed values_: `UPDATE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Update`

Information about the gateways to update.

_Required_: No

_Type_: [UpdateWirelessGatewayTaskCreate](aws-properties-iotwireless-taskdefinition-updatewirelessgatewaytaskcreate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the task definition.

### Fn::GetAtt

`Arn`

The Amazon Resource Name of the resource.

`Id`

The ID of the new wireless gateway task definition.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

LoRaWANGatewayVersion

All content copied from https://docs.aws.amazon.com/.
