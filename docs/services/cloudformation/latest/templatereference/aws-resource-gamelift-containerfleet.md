This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GameLift::ContainerFleet

Describes an Amazon GameLift Servers managed container fleet.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::GameLift::ContainerFleet",
  "Properties" : {
      "BillingType" : String,
      "DeploymentConfiguration" : DeploymentConfiguration,
      "Description" : String,
      "FleetRoleArn" : String,
      "GameServerContainerGroupDefinitionName" : String,
      "GameServerContainerGroupsPerInstance" : Integer,
      "GameSessionCreationLimitPolicy" : GameSessionCreationLimitPolicy,
      "InstanceConnectionPortRange" : ConnectionPortRange,
      "InstanceInboundPermissions" : [ IpPermission, ... ],
      "InstanceType" : String,
      "Locations" : [ LocationConfiguration, ... ],
      "LogConfiguration" : LogConfiguration,
      "MetricGroups" : [ String, ... ],
      "NewGameSessionProtectionPolicy" : String,
      "PerInstanceContainerGroupDefinitionName" : String,
      "PlayerGatewayMode" : String,
      "ScalingPolicies" : [ ScalingPolicy, ... ],
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::GameLift::ContainerFleet
Properties:
  BillingType: String
  DeploymentConfiguration:
    DeploymentConfiguration
  Description: String
  FleetRoleArn: String
  GameServerContainerGroupDefinitionName: String
  GameServerContainerGroupsPerInstance: Integer
  GameSessionCreationLimitPolicy:
    GameSessionCreationLimitPolicy
  InstanceConnectionPortRange:
    ConnectionPortRange
  InstanceInboundPermissions:
    - IpPermission
  InstanceType: String
  Locations:
    - LocationConfiguration
  LogConfiguration:
    LogConfiguration
  MetricGroups:
    - String
  NewGameSessionProtectionPolicy: String
  PerInstanceContainerGroupDefinitionName: String
  PlayerGatewayMode: String
  ScalingPolicies:
    - ScalingPolicy
  Tags:
    - Tag

```

## Properties

`BillingType`

Indicates whether the fleet uses On-Demand or Spot instances for this fleet. Learn
more about when to use [On-Demand versus Spot Instances](https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-ec2-instances.html#gamelift-ec2-instances-spot). You can't update this fleet
property.

By default, this property is set to `ON_DEMAND`.

_Required_: No

_Type_: String

_Allowed values_: `ON_DEMAND | SPOT`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DeploymentConfiguration`

Set of rules for processing a deployment for a container fleet update.

_Required_: No

_Type_: [DeploymentConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-gamelift-containerfleet-deploymentconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A meaningful description of the container fleet.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FleetRoleArn`

The unique identifier for an AWS Identity and Access Management (IAM) role with permissions to run your
containers on resources that are managed by Amazon GameLift Servers. See [Set up an IAM service\
role](https://docs.aws.amazon.com/gamelift/latest/developerguide/setting-up-role.html). This fleet property can't be changed.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws(-.*)?:[a-z-]+:(([a-z]+-)+[0-9])?:([0-9]{12})?:[^.]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GameServerContainerGroupDefinitionName`

The name of the fleet's game server container group definition, which describes how to
deploy containers with your game server build and support software onto each fleet
instance.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9\-]+$|^arn:.*:containergroupdefinition\/[a-zA-Z0-9\-]+(:[0-9]+)?$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GameServerContainerGroupsPerInstance`

The number of times to replicate the game server container group on each fleet
instance.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `5000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GameSessionCreationLimitPolicy`

A policy that limits the number of game sessions that each individual player can create
on instances in this fleet. The limit applies for a specified span of time.

_Required_: No

_Type_: [GameSessionCreationLimitPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-gamelift-containerfleet-gamesessioncreationlimitpolicy.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceConnectionPortRange`

The set of port numbers to open on each instance in a container fleet. Connection
ports are used by inbound traffic to connect with processes that are running in
containers on the fleet.

_Required_: No

_Type_: [ConnectionPortRange](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-gamelift-containerfleet-connectionportrange.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceInboundPermissions`

The IP address ranges and port settings that allow inbound traffic to access game
server processes and other processes on this fleet.

_Required_: No

_Type_: Array of [IpPermission](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-gamelift-containerfleet-ippermission.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceType`

The Amazon EC2 instance type to use for all instances in the fleet. Instance type
determines the computing resources and processing power that's available to host your
game servers. This includes including CPU, memory, storage, and networking capacity. You
can't update this fleet property.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Locations`

A set of locations to deploy container fleet instances to. You can add any AWS
Region or Local Zone that's supported by Amazon GameLift Servers. Provide a list of one or more AWS
Region codes, such as `us-west-2`, or Local Zone names. Also include the
fleet's home Region, which is the AWS Region where the fleet is created. For a list of
supported Regions and Local Zones, see [Amazon GameLift Servers service\
locations](https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-regions.html) for managed hosting.

_Required_: No

_Type_: Array of [LocationConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-gamelift-containerfleet-locationconfiguration.html)

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogConfiguration`

The method that is used to collect container logs for the fleet. Amazon GameLift Servers saves all
standard output for each container in logs, including game session logs.

- `CLOUDWATCH` \-\- Send logs to an Amazon CloudWatch log group that you define. Each container
emits a log stream, which is organized in the log group.

- `S3` \-\- Store logs in an Amazon S3 bucket that you define.

- `NONE` \-\- Don't collect container logs.

_Required_: No

_Type_: [LogConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-gamelift-containerfleet-logconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetricGroups`

The name of an AWS CloudWatch metric group to add this fleet to. Metric groups
aggregate metrics for multiple fleets.

_Required_: No

_Type_: Array of String

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NewGameSessionProtectionPolicy`

Determines whether Amazon GameLift Servers can shut down game sessions on the fleet that are actively
running and hosting players. Amazon GameLift Servers might prompt an instance shutdown when scaling down
fleet capacity or when retiring unhealthy instances. You can also set game session
protection for individual game sessions using [UpdateGameSession](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/gamelift/latest/apireference/API_UpdateGameSession.html).

- **NoProtection** \-\- Game sessions can be shut down
during active gameplay.

- **FullProtection** \-\- Game sessions in
`ACTIVE` status can't be shut down.

_Required_: No

_Type_: String

_Allowed values_: `FullProtection | NoProtection`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PerInstanceContainerGroupDefinitionName`

The name of the fleet's per-instance container group definition.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9\-]+$|^arn:.*:containergroupdefinition\/[a-zA-Z0-9\-]+(:[0-9]+)?$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PlayerGatewayMode`

Configures player gateway for your fleet. Player gateway provides benefits such as DDoS protection by rate limiting and validating traﬃc before it reaches game servers, hiding game server IP addresses from players, and providing updated endpoints when relay endpoints become unhealthy.

**How it works:** When enabled, game clients connect to relay endpoints instead of to your game servers. Player gateway validates player gateway tokens and routes traffic to the appropriate game server. Your game backend calls [GetPlayerConnectionDetails](https://docs.aws.amazon.com/gamelift/latest/apireference/API_GetPlayerConnectionDetails.html) to retrieve relay endpoints and player gateway tokens for your game clients. To learn more about this topic, see [DDoS protection with Amazon GameLift Servers player gateway](https://docs.aws.amazon.com/gameliftservers/latest/developerguide/ddos-protection-intro.html).

Possible values include:

- `DISABLED` (default) -- Game clients connect to the game server endpoint. Use this when you do not intend to integrate your game with player gateway.

- `ENABLED` \-\- Player gateway is available in fleet locations where it is supported. Your game backend can call [GetPlayerConnectionDetails](https://docs.aws.amazon.com/gamelift/latest/apireference/API_GetPlayerConnectionDetails.html) to obtain a player gateway token and endpoints for game clients.

- `REQUIRED` \-\- Player gateway is available in fleet locations where it is supported, and the fleet can only use locations that support this feature. Attempting to add a remote location to your fleet which does not support player gateway will result in an `InvalidRequestException`.

_Required_: No

_Type_: String

_Allowed values_: `DISABLED | ENABLED | REQUIRED`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ScalingPolicies`

Rule that controls how a fleet is scaled. Scaling policies are uniquely identified by
the combination of name and fleet ID.

_Required_: No

_Type_: Array of [ScalingPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-gamelift-containerfleet-scalingpolicy.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A list of labels to assign to the new fleet resource. Tags are developer-defined
key-value pairs. Tagging AWS resources are useful for resource management, access
management and cost allocation. For more information, see [Tagging AWS Resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the
_AWS General Reference_.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-gamelift-containerfleet-tag.html)

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`CreationTime`

A time stamp indicating when this data object was created. Format is a number expressed in Unix time as milliseconds (for example `"1469498468.057"`).

`FleetArn`

The Amazon Resource Name ( [ARN](https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-arn-format.html)) that is assigned to a Amazon GameLift Servers fleet resource and uniquely identifies it. ARNs are unique across all Regions. Format is `arn:aws:gamelift:<region>::fleet/fleet-a1234567-b8c9-0d1e-2fa3-b45c6d7e8912`. In a GameLift fleet ARN, the resource ID matches the `FleetId`
value.

`FleetId`

A unique identifier for the container fleet to retrieve.

`GameServerContainerGroupDefinitionArn`

The Amazon Resource Name ( [ARN](https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-arn-format.html)) that is assigned to the fleet's game server container group. The ARN value
also identifies the specific container group definition version in use.

`MaximumGameServerContainerGroupsPerInstance`

The calculated maximum number of game server container group that can be deployed on
each fleet instance. The calculation depends on the resource needs of the container
group and the CPU and memory resources of the fleet's instance type.

`PerInstanceContainerGroupDefinitionArn`

The Amazon Resource Name ( [ARN](https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-arn-format.html)) that is assigned to the fleet's per-instance container group. The ARN value
also identifies the specific container group definition version in use.

`Status`

The current status of the container fleet.

- `PENDING` \-\- A new container fleet has been requested.

- `CREATING` \-\- A new container fleet resource is being created.

- `CREATED` \-\- A new container fleet resource has been created. No fleet instances
have been deployed.

- `ACTIVATING` \-\- New container fleet instances are being deployed.

- `ACTIVE` \-\- The container fleet has been deployed and is ready to host game
sessions.

- `UPDATING` \-\- Updates to the container fleet is being updated. A deployment is in
progress.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

ConnectionPortRange
