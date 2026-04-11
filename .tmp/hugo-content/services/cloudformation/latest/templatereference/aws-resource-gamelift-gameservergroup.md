This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GameLift::GameServerGroup

**This operation is used with the Amazon GameLift FleetIQ solution and game server groups.**

Creates a GameLift FleetIQ game server group for managing game hosting on a collection of
Amazon EC2 instances for game hosting. This operation creates the game server group,
creates an Auto Scaling group in your AWS account, and establishes a link between the
two groups. You can view the status of your game server groups in the GameLift console.
Game server group metrics and events are emitted to Amazon CloudWatch.

Before creating a new game server group, you must have the following:

- An Amazon EC2 launch template that specifies how to launch Amazon EC2 instances
with your game server build. For more information, see [Launching an Instance from a Launch Template](../../../ec2/latest/userguide/ec2-launch-templates.md) in the
_Amazon EC2 User Guide_.

- An IAM role that extends limited access to your AWS account to allow GameLift FleetIQ to create and
interact with the Auto Scaling group. For more information, see [Create IAM roles for cross-service interaction](../../../gamelift/latest/fleetiqguide/gsg-iam-permissions-roles.md) in the _GameLift FleetIQ Developer_
_Guide_.

To create a new game server group, specify a unique group name, IAM role and Amazon EC2
launch template, and provide a list of instance types that can be used in the group. You
must also set initial maximum and minimum limits on the group's instance count. You can
optionally set an Auto Scaling policy with target tracking based on a GameLift FleetIQ
metric.

Once the game server group and corresponding Auto Scaling group are created, you have
full access to change the Auto Scaling group's configuration as needed. Several
properties that are set when creating a game server group, including maximum/minimum
size and auto-scaling policy settings, must be updated directly in the Auto Scaling
group. Keep in mind that some Auto Scaling group properties are periodically updated by
GameLift FleetIQ as part of its balancing activities to optimize for availability and
cost.

**Learn more**

[GameLift FleetIQ Guide](../../../gamelift/latest/fleetiqguide/gsg-intro.md)

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::GameLift::GameServerGroup",
  "Properties" : {
      "AutoScalingPolicy" : AutoScalingPolicy,
      "BalancingStrategy" : String,
      "DeleteOption" : String,
      "GameServerGroupName" : String,
      "GameServerProtectionPolicy" : String,
      "InstanceDefinitions" : [ InstanceDefinition, ... ],
      "LaunchTemplate" : LaunchTemplate,
      "MaxSize" : Number,
      "MinSize" : Number,
      "RoleArn" : String,
      "Tags" : [ Tag, ... ],
      "VpcSubnets" : [ String, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::GameLift::GameServerGroup
Properties:
  AutoScalingPolicy:
    AutoScalingPolicy
  BalancingStrategy: String
  DeleteOption: String
  GameServerGroupName: String
  GameServerProtectionPolicy: String
  InstanceDefinitions:
    - InstanceDefinition
  LaunchTemplate:
    LaunchTemplate
  MaxSize: Number
  MinSize: Number
  RoleArn: String
  Tags:
    - Tag
  VpcSubnets:
    - String

```

## Properties

`AutoScalingPolicy`

Configuration settings to define a scaling policy for the Auto Scaling group that is
optimized for game hosting. The scaling policy uses the metric
`"PercentUtilizedGameServers"` to maintain a buffer of idle game servers
that can immediately accommodate new games and players. After the Auto Scaling group is
created, update this value directly in the Auto Scaling group using the AWS console or
APIs.

_Required_: No

_Type_: [AutoScalingPolicy](aws-properties-gamelift-gameservergroup-autoscalingpolicy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BalancingStrategy`

Indicates how Amazon GameLift Servers FleetIQ balances the use of Spot Instances and On-Demand Instances in the
game server group. Method options include the following:

- `SPOT_ONLY` \- Only Spot Instances are used in the game server group. If Spot
Instances are unavailable or not viable for game hosting, the game server group
provides no hosting capacity until Spot Instances can again be used. Until then,
no new instances are started, and the existing nonviable Spot Instances are
terminated (after current gameplay ends) and are not replaced.

- `SPOT_PREFERRED` \- (default value) Spot Instances are used whenever available in
the game server group. If Spot Instances are unavailable, the game server group
continues to provide hosting capacity by falling back to On-Demand Instances.
Existing nonviable Spot Instances are terminated (after current gameplay ends)
and are replaced with new On-Demand Instances.

- `ON_DEMAND_ONLY` \- Only On-Demand Instances are used in the game
server group. No Spot Instances are used, even when available, while this
balancing strategy is in force.

_Required_: No

_Type_: String

_Allowed values_: `SPOT_ONLY | SPOT_PREFERRED | ON_DEMAND_ONLY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeleteOption`

The type of delete to perform. To delete a game server group, specify the
`DeleteOption`. Options include the following:

- `SAFE_DELETE` – (default) Terminates the game server group and
Amazon EC2 Auto Scaling group only when it has no game servers that are in
`UTILIZED` status.

- `FORCE_DELETE` – Terminates the game server group, including all
active game servers regardless of their utilization status, and the Amazon EC2 Auto
Scaling group.

- `RETAIN` – Does a safe delete of the game server group but retains
the Amazon EC2 Auto Scaling group as is.

_Required_: No

_Type_: String

_Allowed values_: `SAFE_DELETE | FORCE_DELETE | RETAIN`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GameServerGroupName`

A developer-defined identifier for the game server group. The name is unique for each
Region in each AWS account.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9-\.]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GameServerProtectionPolicy`

A flag that indicates whether instances in the game server group are protected
from early termination. Unprotected instances that have active game servers running might
be terminated during a scale-down event, causing players to be dropped from the game.
Protected instances cannot be terminated while there are active game servers running except
in the event of a forced game server group deletion (see ). An exception to this is with Spot
Instances, which can be terminated by AWS regardless of protection status.

_Required_: No

_Type_: String

_Allowed values_: `NO_PROTECTION | FULL_PROTECTION`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceDefinitions`

The set of Amazon EC2 instance types that Amazon GameLift Servers FleetIQ can use when balancing and automatically
scaling instances in the corresponding Auto Scaling group.

_Required_: Yes

_Type_: Array of [InstanceDefinition](aws-properties-gamelift-gameservergroup-instancedefinition.md)

_Minimum_: `2`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LaunchTemplate`

The Amazon EC2 launch template that contains configuration settings and game server code to
be deployed to all instances in the game server group. You can specify the template
using either the template name or ID. For help with creating a launch template, see
[Creating a Launch\
Template for an Auto Scaling Group](../../../autoscaling/ec2/userguide/create-launch-template.md) in the _Amazon Elastic Compute Cloud Auto Scaling_
_User Guide_. After the Auto Scaling group is created, update this value
directly in the Auto Scaling group using the AWS console or APIs.

###### Note

If you specify network interfaces in your launch template, you must explicitly set
the property `AssociatePublicIpAddress` to "true". If no network
interface is specified in the launch template, Amazon GameLift Servers FleetIQ uses your account's default
VPC.

_Required_: No

_Type_: [LaunchTemplate](aws-properties-gamelift-gameservergroup-launchtemplate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxSize`

The maximum number of instances allowed in the Amazon EC2 Auto Scaling group. During
automatic scaling events, Amazon GameLift Servers FleetIQ and EC2 do not scale up the group above this maximum.
After the Auto Scaling group is created, update this value directly in the Auto Scaling
group using the AWS console or APIs.

_Required_: No

_Type_: Number

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinSize`

The minimum number of instances allowed in the Amazon EC2 Auto Scaling group. During
automatic scaling events, Amazon GameLift Servers FleetIQ and Amazon EC2 do not scale down the group below this
minimum. In production, this value should be set to at least 1. After the Auto Scaling
group is created, update this value directly in the Auto Scaling group using the AWS
console or APIs.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The Amazon Resource Name ( [ARN](../../../s3/latest/dev/s3-arn-format.md)) for an IAM role that
allows Amazon GameLift Servers to access your Amazon EC2 Auto Scaling groups.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:.*:role\/[\w+=,.@-]+`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A list of labels to assign to the new game server group resource. Tags are
developer-defined key-value pairs. Tagging AWS resources is useful for resource
management, access management, and cost allocation. For more information, see [Tagging AWS\
Resources](../../../../general/latest/gr/aws-tagging.md) in the _AWS General Reference_. Once the
resource is created, you can use TagResource, UntagResource, and ListTagsForResource to add, remove,
and view tags, respectively. The maximum tag limit may be lower than stated. See the
AWS General Reference for actual tagging limits.

_Required_: No

_Type_: Array of [Tag](aws-properties-gamelift-gameservergroup-tag.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcSubnets`

A list of virtual private cloud (VPC) subnets to use with instances in the game server
group. By default, all Amazon GameLift Servers FleetIQ-supported Availability Zones are used. You can use this
parameter to specify VPCs that you've set up. This property cannot be updated after the
game server group is created, and the corresponding Auto Scaling group will always use
the property value that is set with this request, even if the Auto Scaling group is
updated directly.

_Required_: No

_Type_: Array of String

_Minimum_: `15 | 1`

_Maximum_: `24 | 20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AutoScalingGroupArn`

A unique identifier for the auto scaling group.

`GameServerGroupArn`

A unique identifier for the game server group.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TargetConfiguration

AutoScalingPolicy

All content copied from https://docs.aws.amazon.com/.
