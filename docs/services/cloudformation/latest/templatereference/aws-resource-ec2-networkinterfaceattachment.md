This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::NetworkInterfaceAttachment

Attaches an elastic network interface (ENI) to an Amazon EC2 instance. You can use this
resource type to attach additional network interfaces to an instance without
interruption.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::NetworkInterfaceAttachment",
  "Properties" : {
      "DeleteOnTermination" : Boolean,
      "DeviceIndex" : String,
      "EnaQueueCount" : Integer,
      "EnaSrdSpecification" : EnaSrdSpecification,
      "InstanceId" : String,
      "NetworkInterfaceId" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::NetworkInterfaceAttachment
Properties:
  DeleteOnTermination: Boolean
  DeviceIndex: String
  EnaQueueCount: Integer
  EnaSrdSpecification:
    EnaSrdSpecification
  InstanceId: String
  NetworkInterfaceId: String

```

## Properties

`DeleteOnTermination`

Whether to delete the network interface when the instance terminates. By default, this
value is set to `true`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeviceIndex`

The network interface's position in the attachment order. For example, the first
attached network interface has a `DeviceIndex` of 0.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EnaQueueCount`

The number of ENA queues created with the instance.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnaSrdSpecification`

Configures ENA Express for the network interface that this action attaches to the
instance.

_Required_: No

_Type_: [EnaSrdSpecification](aws-properties-ec2-networkinterfaceattachment-enasrdspecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceId`

The ID of the instance to which you will attach the ENI.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NetworkInterfaceId`

The ID of the ENI that you want to attach.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AttachmentId`

The ID of the network interface attachment.

## Examples

### Network interface attachment

The following example attaches `MyNetworkInterface` to
`MyInstance`.

#### JSON

```json

"NetworkInterfaceAttachment" : {
   "Type" : "AWS::EC2::NetworkInterfaceAttachment",
      "Properties" : {
         "InstanceId" : {"Ref" : "MyInstance"},
         "NetworkInterfaceId" : {"Ref" : "MyNetworkInterface"},
         "DeviceIndex" : "1"
      }
}
```

#### YAML

```yaml

NetworkInterfaceAttachment:
   Type: AWS::EC2::NetworkInterfaceAttachment
   Properties:
      InstanceId:
         Ref: MyInstance
      NetworkInterfaceId:
         Ref: MyNetworkInterface
      DeviceIndex: 1
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

EnaSrdSpecification

All content copied from https://docs.aws.amazon.com/.
