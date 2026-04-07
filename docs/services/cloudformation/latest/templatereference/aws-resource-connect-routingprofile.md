This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::RoutingProfile

Creates a new routing profile.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Connect::RoutingProfile",
  "Properties" : {
      "AgentAvailabilityTimer" : String,
      "DefaultOutboundQueueArn" : String,
      "Description" : String,
      "InstanceArn" : String,
      "ManualAssignmentQueueConfigs" : [ RoutingProfileManualAssignmentQueueConfig, ... ],
      "MediaConcurrencies" : [ MediaConcurrency, ... ],
      "Name" : String,
      "QueueConfigs" : [ RoutingProfileQueueConfig, ... ],
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Connect::RoutingProfile
Properties:
  AgentAvailabilityTimer: String
  DefaultOutboundQueueArn: String
  Description: String
  InstanceArn: String
  ManualAssignmentQueueConfigs:
    - RoutingProfileManualAssignmentQueueConfig
  MediaConcurrencies:
    - MediaConcurrency
  Name: String
  QueueConfigs:
    - RoutingProfileQueueConfig
  Tags:
    - Tag

```

## Properties

`AgentAvailabilityTimer`

Whether agents with this routing profile will have their routing order calculated based on _time since_
_their last inbound contact_ or _longest idle time_.

_Required_: No

_Type_: String

_Allowed values_: `TIME_SINCE_LAST_ACTIVITY | TIME_SINCE_LAST_INBOUND`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultOutboundQueueArn`

The Amazon Resource Name (ARN) of the default outbound queue for the routing
profile.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*/queue/[-a-zA-Z0-9]*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the routing profile.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `250`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceArn`

The identifier of the Amazon Connect instance.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ManualAssignmentQueueConfigs`

Contains information about the queue and channel for manual assignment behaviour can
be enabled.

_Required_: No

_Type_: Array of [RoutingProfileManualAssignmentQueueConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-routingprofile-routingprofilemanualassignmentqueueconfig.html)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MediaConcurrencies`

The channels agents can handle in the Contact Control Panel (CCP) for this routing profile.

_Required_: Yes

_Type_: Array of [MediaConcurrency](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-routingprofile-mediaconcurrency.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the routing profile.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `127`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QueueConfigs`

The inbound queues associated with the routing profile. If no queue is added, the
agent can make only outbound calls.

_Required_: No

_Type_: Array of [RoutingProfileQueueConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-routingprofile-routingprofilequeueconfig.html)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags used to organize, track, or control access for this resource. For example, { "Tags": {"key1":"value1", "key2":"value2"} }.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-routingprofile-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the routing profile name. For example:

`{ "Ref": "myRoutingProfileName" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`RoutingProfileArn`

The Amazon Resource Name (ARN) of the routing profile.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

UserQuickConnectConfig

CrossChannelBehavior
