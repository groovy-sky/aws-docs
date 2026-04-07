This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::VPCEndpointConnectionNotification

Specifies a connection notification for a VPC endpoint or VPC endpoint service. A
connection notification notifies you of specific endpoint events. You must create an SNS
topic to receive notifications. For more information, see [Create a Topic](https://docs.aws.amazon.com/sns/latest/dg/CreateTopic.html) in the _Amazon_
_Simple Notification Service Developer Guide_.

You can create a connection notification for interface endpoints only.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::VPCEndpointConnectionNotification",
  "Properties" : {
      "ConnectionEvents" : [ String, ... ],
      "ConnectionNotificationArn" : String,
      "ServiceId" : String,
      "VPCEndpointId" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::VPCEndpointConnectionNotification
Properties:
  ConnectionEvents:
    - String
  ConnectionNotificationArn: String
  ServiceId: String
  VPCEndpointId: String

```

## Properties

`ConnectionEvents`

The endpoint events for which to receive notifications. Valid values are
`Accept`, `Connect`, `Delete`, and
`Reject`.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConnectionNotificationArn`

The ARN of the SNS topic for the notifications.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceId`

The ID of the endpoint service.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VPCEndpointId`

The ID of the endpoint.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the notification.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`VPCEndpointConnectionNotificationId`

The ID of the notification.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AWS::EC2::VPCEndpointService
