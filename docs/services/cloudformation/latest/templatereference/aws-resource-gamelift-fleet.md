This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GameLift::Fleet

The `AWS::GameLift::Fleet` resource creates an Amazon GameLift (GameLift) fleet to host
custom game server or Realtime Servers. A fleet is a set of EC2 instances, configured with instructions to
run game servers on each instance.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::GameLift::Fleet",
  "Properties" : {
      "AnywhereConfiguration" : AnywhereConfiguration,
      "ApplyCapacity" : String,
      "BuildId" : String,
      "CertificateConfiguration" : CertificateConfiguration,
      "ComputeType" : String,
      "Description" : String,
      "EC2InboundPermissions" : [ IpPermission, ... ],
      "EC2InstanceType" : String,
      "FleetType" : String,
      "InstanceRoleARN" : String,
      "InstanceRoleCredentialsProvider" : String,
      "Locations" : [ LocationConfiguration, ... ],
      "MetricGroups" : [ String, ... ],
      "Name" : String,
      "NewGameSessionProtectionPolicy" : String,
      "PeerVpcAwsAccountId" : String,
      "PeerVpcId" : String,
      "PlayerGatewayConfiguration" : PlayerGatewayConfiguration,
      "PlayerGatewayMode" : String,
      "ResourceCreationLimitPolicy" : ResourceCreationLimitPolicy,
      "RuntimeConfiguration" : RuntimeConfiguration,
      "ScalingPolicies" : [ ScalingPolicy, ... ],
      "ScriptId" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::GameLift::Fleet
Properties:
  AnywhereConfiguration:
    AnywhereConfiguration
  ApplyCapacity: String
  BuildId: String
  CertificateConfiguration:
    CertificateConfiguration
  ComputeType: String
  Description: String
  EC2InboundPermissions:
    - IpPermission
  EC2InstanceType: String
  FleetType: String
  InstanceRoleARN: String
  InstanceRoleCredentialsProvider: String
  Locations:
    - LocationConfiguration
  MetricGroups:
    - String
  Name: String
  NewGameSessionProtectionPolicy: String
  PeerVpcAwsAccountId: String
  PeerVpcId: String
  PlayerGatewayConfiguration:
    PlayerGatewayConfiguration
  PlayerGatewayMode: String
  ResourceCreationLimitPolicy:
    ResourceCreationLimitPolicy
  RuntimeConfiguration:
    RuntimeConfiguration
  ScalingPolicies:
    - ScalingPolicy
  ScriptId: String
  Tags:
    - Tag

```

## Properties

`AnywhereConfiguration`

Amazon GameLift Servers Anywhere configuration options.

_Required_: No

_Type_: [AnywhereConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-gamelift-fleet-anywhereconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ApplyCapacity`

Current resource capacity settings for managed EC2 fleets and managed container fleets. For
multi-location fleets, location values might refer to a fleet's remote location or its
home Region.

**Returned by:** [DescribeFleetCapacity](https://docs.aws.amazon.com/gamelift/latest/apireference/API_DescribeFleetCapacity.html), [DescribeFleetLocationCapacity](https://docs.aws.amazon.com/gamelift/latest/apireference/API_DescribeFleetLocationCapacity.html), [UpdateFleetCapacity](https://docs.aws.amazon.com/gamelift/latest/apireference/API_UpdateFleetCapacity.html)

_Required_: No

_Type_: String

_Allowed values_: `ON_UPDATE | ON_CREATE_AND_UPDATE | ON_CREATE_AND_UPDATE_WITH_AUTOSCALING`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BuildId`

A unique identifier for a build to be deployed on the new fleet. If you are
deploying the fleet with a custom game build, you must specify this property. The build must
have been successfully uploaded to Amazon GameLift and be in a `READY` status. This
fleet setting cannot be changed once the fleet is created.

_Required_: Conditional

_Type_: String

_Pattern_: `^build-\S+|^arn:.*:build/build-\S+`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CertificateConfiguration`

Prompts Amazon GameLift Servers to generate a TLS/SSL certificate for the fleet. Amazon GameLift Servers uses the
certificates to encrypt traffic between game clients and the game servers running on
Amazon GameLift Servers. By default, the `CertificateConfiguration` is `DISABLED`.
You can't change this property after you create the fleet.

AWS Certificate Manager (ACM) certificates expire after 13 months.
Certificate expiration can cause fleets to fail, preventing players from connecting to
instances in the fleet. We recommend you replace fleets before 13 months, consider using
fleet aliases for a smooth transition.

###### Note

ACM isn't available in all AWS regions. A fleet creation request
with certificate generation enabled in an unsupported Region, fails with a 4xx
error. For more information about the supported Regions, see [Supported\
Regions](https://docs.aws.amazon.com/acm/latest/userguide/acm-regions.html) in the _AWS Certificate Manager User_
_Guide_.

_Required_: No

_Type_: [CertificateConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-gamelift-fleet-certificateconfiguration.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ComputeType`

The type of compute resource used to host your game servers.

- `EC2` – The game server build is deployed to Amazon EC2 instances for
cloud hosting. This is the default setting.

- `ANYWHERE` – Game servers
and supporting software are deployed to compute resources that you provide and
manage. With this compute type, you can also set the
`AnywhereConfiguration` parameter.

_Required_: No

_Type_: String

_Allowed values_: `EC2 | ANYWHERE`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

A description for the fleet.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EC2InboundPermissions`

The IP address ranges and port settings that allow inbound traffic to access game
server processes and other processes on this fleet. Set this parameter for managed EC2 fleets. You can leave this parameter empty when creating the fleet, but you must call
[https://docs.aws.amazon.com/gamelift/latest/apireference/API\_UpdateFleetPortSettings](https://docs.aws.amazon.com/gamelift/latest/apireference/API_UpdateFleetPortSettings) to set it before players can connect to game sessions.
As a best practice, we recommend
opening ports for remote access only when you need them and closing them when you're finished.
For Amazon GameLift Servers Realtime fleets, Amazon GameLift Servers automatically sets TCP and UDP ranges.

_Required_: No

_Type_: Array of [IpPermission](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-gamelift-fleet-ippermission.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EC2InstanceType`

The Amazon GameLift Servers-supported Amazon EC2 instance type to use with managed EC2 fleets.
Instance type determines the computing resources that will be used to host your game
servers, including CPU, memory, storage, and networking capacity. See [Amazon Elastic Compute Cloud Instance Types](https://aws.amazon.com/ec2/instance-types) for
detailed descriptions of Amazon EC2 instance types.

_Required_: No

_Type_: String

_Pattern_: `^.*..*$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FleetType`

Indicates whether to use On-Demand or Spot instances for this fleet. By default, this
property is set to `ON_DEMAND`. Learn more about when to use [On-Demand versus Spot Instances](https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-ec2-instances.html#gamelift-ec2-instances-spot). This fleet property can't be changed after the fleet is created.

_Required_: No

_Type_: String

_Allowed values_: `ON_DEMAND | SPOT`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InstanceRoleARN`

A unique identifier for an IAM role that manages access to your AWS services.
With an instance role ARN set, any application that runs on an instance in this fleet can assume the role,
including install scripts, server processes, and daemons (background processes). Create a role or look up a role's
ARN by using the [IAM dashboard](https://console.aws.amazon.com/iam) in the AWS Management Console.
Learn more about using on-box credentials for your game servers at
[Access external resources from a game server](https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-sdk-server-resources.html). This attribute is used with fleets where `ComputeType` is
`EC2`.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws(-.*)?:[a-z-]+:(([a-z]+-)+[0-9])?:([0-9]{12})?:[^.]+$`

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InstanceRoleCredentialsProvider`

Indicates that fleet instances maintain a shared credentials file for the IAM role defined in `InstanceRoleArn`. Shared credentials allow
applications that are deployed with the game server executable to communicate with other
AWS resources. This property is used only when the game server is integrated with the
server SDK version 5.x. For more information about using shared credentials, see [Communicate\
with other AWS resources from your fleets](https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-sdk-server-resources.html). This attribute is used with
fleets where `ComputeType` is `EC2`.

_Required_: No

_Type_: String

_Allowed values_: `SHARED_CREDENTIAL_FILE`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Locations`

A set of remote locations to deploy additional instances to and manage as a
multi-location fleet. Use this parameter when creating a fleet in AWS Regions that
support multiple locations. You can add any AWS Region or Local Zone that's supported
by Amazon GameLift Servers. Provide a list of one or more AWS Region codes, such as
`us-west-2`, or Local Zone names. When using this parameter, Amazon GameLift Servers
requires you to include your home location in the request. For a list of supported
Regions and Local Zones, see
[Amazon GameLift Servers service locations](https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-regions.html) for managed hosting.

_Required_: No

_Type_: Array of [LocationConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-gamelift-fleet-locationconfiguration.html)

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetricGroups`

The name of an AWS CloudWatch metric group to add this fleet to. A metric group is
used to aggregate the metrics for multiple fleets. You can specify an existing metric
group name or set a new name to create a new metric group. A fleet can be included in
only one metric group at a time.

_Required_: No

_Type_: Array of String

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

A descriptive label that is associated with a fleet. Fleet names do not need to be unique.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NewGameSessionProtectionPolicy`

The status of termination protection for active game sessions on the fleet. By
default, this property is set to `NoProtection`.

- **NoProtection** \- Game sessions can be terminated
during active gameplay as a result of a scale-down event.

- **FullProtection** \- Game sessions in
`ACTIVE` status cannot be terminated during a scale-down
event.

_Required_: No

_Type_: String

_Allowed values_: `FullProtection | NoProtection`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PeerVpcAwsAccountId`

Used when peering your Amazon GameLift Servers fleet with a VPC, the unique identifier for the AWS
account that owns the VPC. You can find your account ID in the AWS Management Console under account
settings.

_Required_: No

_Type_: String

_Pattern_: `^[0-9]{12}$`

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PeerVpcId`

A unique identifier for a VPC with resources to be accessed by your Amazon GameLift Servers fleet. The
VPC must be in the same Region as your fleet. To look up a VPC ID, use the
[VPC Dashboard](https://console.aws.amazon.com/vpc) in the AWS Management Console.
Learn more about VPC peering in [VPC Peering with Amazon GameLift Servers Fleets](https://docs.aws.amazon.com/gamelift/latest/developerguide/vpc-peering.html).

_Required_: No

_Type_: String

_Pattern_: `^vpc-\S+`

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PlayerGatewayConfiguration`

Configuration settings for player gateway. Use this to specify advanced options for how player gateway handles connections.

_Required_: No

_Type_: [PlayerGatewayConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-gamelift-fleet-playergatewayconfiguration.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PlayerGatewayMode`

Configures player gateway for your fleet. Player gateway provides benefits such as DDoS protection by rate limiting and validating traﬃc before it reaches game servers, hiding game server IP addresses from players, and providing updated endpoints when relay endpoints become unhealthy. Note, player gateway is only available for fleets using server SDK 5.x or later game server builds.

**How it works:** When enabled, game clients connect to relay endpoints instead of to your game servers. Player gateway validates player gateway tokens and routes traffic to the appropriate game server. Your game backend calls [GetPlayerConnectionDetails](https://docs.aws.amazon.com/gamelift/latest/apireference/API_GetPlayerConnectionDetails.html) to retrieve relay endpoints and player gateway tokens for your game clients. To learn more about this topic, see [DDoS protection with Amazon GameLift Servers player gateway](https://docs.aws.amazon.com/gameliftservers/latest/developerguide/ddos-protection-intro.html).

Possible values include:

- `DISABLED` (default) -- Game clients connect to the game server endpoint. Use this when you do not intend to integrate your game with player gateway.

- `ENABLED` \-\- Player gateway is available in fleet locations where it is supported. Your game backend can call [GetPlayerConnectionDetails](https://docs.aws.amazon.com/gamelift/latest/apireference/API_GetPlayerConnectionDetails.html) to obtain a player gateway token and endpoints for game clients.

- `REQUIRED` \-\- Player gateway is available in fleet locations where it is supported, and the fleet can only use locations that support this feature. Attempting to add a remote location to your fleet which does not support player gateway will result in an `InvalidRequestException`.

_Required_: No

_Type_: String

_Allowed values_: `DISABLED | ENABLED | REQUIRED`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ResourceCreationLimitPolicy`

A policy that limits the number of game sessions that an individual player can create
on instances in this fleet within a specified span of time.

_Required_: No

_Type_: [ResourceCreationLimitPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-gamelift-fleet-resourcecreationlimitpolicy.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuntimeConfiguration`

Instructions for how to launch and maintain server processes on instances in the
fleet. The runtime configuration defines one or more server process configurations, each
identifying a build executable or Realtime script file and the number of processes of
that type to run concurrently.

###### Note

The `RuntimeConfiguration` parameter is required unless the fleet is
being configured using the older parameters `ServerLaunchPath` and
`ServerLaunchParameters`, which are still supported for backward
compatibility.

_Required_: Conditional

_Type_: [RuntimeConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-gamelift-fleet-runtimeconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScalingPolicies`

Rule that controls how a fleet is scaled. Scaling policies are uniquely identified by
the combination of name and fleet ID.

_Required_: No

_Type_: Array of [ScalingPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-gamelift-fleet-scalingpolicy.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScriptId`

The unique identifier for a Realtime configuration script to be deployed on fleet
instances. You can use either the script ID or ARN. Scripts must be uploaded to Amazon GameLift Servers
prior to creating the fleet. This fleet property cannot be changed later.

###### Note

You can't use the `!Ref` command to reference a
script created with a CloudFormation template for the fleet property `ScriptId`.
Instead, use `Fn::GetAtt Script.Arn` or `Fn::GetAtt Script.Id` to
retrieve either of these properties as input for `ScriptId`. Alternatively, enter a
`ScriptId` string manually.

_Required_: Conditional

_Type_: String

_Pattern_: `^script-\S+|^arn:.*:script/script-\S+`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

A list of labels to assign to the new fleet resource. Tags are developer-defined
key-value pairs. Tagging AWS resources are useful for resource management, access
management and cost allocation. For more information, see [Tagging AWS Resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the
_AWS General Reference_.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-gamelift-fleet-tag.html)

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the fleet ID, such as
`fleet-1111aaaa-22bb-33cc-44dd-5555eeee66ff`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`FleetArn`

The Amazon Resource Name ( [ARN](https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-arn-format.html)) that is assigned to a Amazon GameLift Servers fleet resource and uniquely identifies it. ARNs are unique across all Regions. Format is `arn:aws:gamelift:<region>::fleet/fleet-a1234567-b8c9-0d1e-2fa3-b45c6d7e8912`. In a GameLift fleet ARN, the resource ID matches the `FleetId`
value.

`FleetId`

A unique identifier for the fleet.

## Examples

- [Create GameLift fleet with a Build](#aws-resource-gamelift-fleet--examples--Create_GameLift_fleet_with_a_Build)

- [Create GameLift fleet with a Script](#aws-resource-gamelift-fleet--examples--Create_GameLift_fleet_with_a_Script)

### Create GameLift fleet with a Build

The following example creates and configures a GameLift fleet for a custom game build.
The fleet uses a `Ref` intrinsic function to specify a build, which can be
declared elsewhere in the same template. The example syntax for log path and server launch
path uses values for a Windows build.

Note: the JSON example shows how to escape the slashes ( `\\`).

#### JSON

```json

{
    "Resources": {
        "FleetResource": {
            "Type": "AWS::GameLift::Fleet",
            "Properties": {
                "BuildId": {
                    "Ref": "BuildResource"
                },
                "CertificateConfiguration": {
                    "CertificateType": "DISABLED"
                },
                "Description": "Description of my Fleet",
                "DesiredEC2Instances": 1,
                "EC2InboundPermissions": [
                    {
                        "FromPort": 1234,
                        "ToPort": 1324,
                        "IpRange": "0.0.0.0/24",
                        "Protocol": "TCP"
                    },
                    {
                        "FromPort": 1356,
                        "ToPort": 1578,
                        "IpRange": "192.168.0.0/24",
                        "Protocol": "UDP"
                    }
                ],
                "EC2InstanceType": "c4.large",
                "FleetType": "SPOT",
                "LogPaths": [
                    "c:\\game\\testlog.log",
                    "c:\\game\\testlog2.log"
                ],
                "MetricGroups": [
                    "MetricGroupName"
                ],
                "Name": "MyGameFleet",
                "NewGameSessionProtectionPolicy": "FullProtection",
                "ResourceCreationLimitPolicy": {
                    "NewGameSessionsPerCreator": 5,
                    "PolicyPeriodInMinutes": 2
                },
                "RuntimeConfiguration": {
                    "GameSessionActivationTimeoutSeconds": 300,
                    "MaxConcurrentGameSessionActivations": 1,
                    "ServerProcesses": [
                        {
                            "ConcurrentExecutions": 1,
                            "LaunchPath": "c:\\game\\TestApplicationServer.exe"
                        }
                    ]
                },
                "Locations": [
                  "us-west-2",
                  "us-east-1",
                  "eu-west-1"
                ]
            }
        }
    }
}
```

#### YAML

```yaml

Resources:
  FleetResource:
    Type: AWS::GameLift::Fleet
    Properties:
      BuildId: !Ref BuildResource
      CertificateConfiguration:
        CertificateType: DISABLED
      Description: Description of my Game Fleet
      DesiredEC2Instances: 1
      EC2InboundPermissions:
        - FromPort: 1234
          ToPort: 1324
          IpRange: 0.0.0.0/24
          Protocol: TCP
        - FromPort: 1356
          ToPort: 1578
          IpRange: 192.168.0.0/24
          Protocol: UDP
      EC2InstanceType: c4.large
      FleetType: SPOT
      LogPaths:
        - c:\game\testlog.log
        - c:\game\testlog2.log
      MetricGroups:
        - MetricGroupName
      Name: MyGameFleet
      NewGameSessionProtectionPolicy: FullProtection
      ResourceCreationLimitPolicy:
        NewGameSessionsPerCreator: 5
        PolicyPeriodInMinutes: 2
      RuntimeConfiguration:
        GameSessionActivationTimeoutSeconds: 300
        MaxConcurrentGameSessionActivations: 1
        ServerProcesses:
          - ConcurrentExecutions: 1
            LaunchPath: c:\game\TestApplicationServer.exe
        Locations:
          - Location: 'us-west-2'
          - Location: 'us-east-1'
          - Location: 'eu-west-1'
```

### Create GameLift fleet with a Script

The following example creates and configures a GameLift fleet to run Realtime Servers.
The fleet uses a `GetAtt` intrinsic function to specify a script, which can be
declared elsewhere in the same template. Note that the syntax example for the log path and
server launch path are for Linux, as all Realtime Servers are deployed with Linux.

#### JSON

```json

{
    "Resources": {
        "FleetResource": {
            "Type": "AWS::GameLift::Fleet",
            "Properties": {
                "CertificateConfiguration": {
                    "CertificateType": "DISABLED"
                },
                "Description": "Description of my Game Fleet",
                "DesiredEC2Instances": 1,
                "EC2InboundPermissions": [
                    {
                        "FromPort": 1234,
                        "ToPort": 1324,
                        "IpRange": "0.0.0.0/24",
                        "Protocol": "TCP"
                    },
                    {
                        "FromPort": 1356,
                        "ToPort": 1578,
                        "IpRange": "192.168.0.0/24",
                        "Protocol": "UDP"
                    }
                ],
                "EC2InstanceType": "c4.large",
                "FleetType": "SPOT",
                "LogPaths": [
                    "/local/game/testlog.log",
                    "/local/game/testlog2.log"
                ],
                "MetricGroups": [
                    "MetricGroupName"
                ],
                "Name": "MyGameFleet",
                "NewGameSessionProtectionPolicy": "FullProtection",
                "ResourceCreationLimitPolicy": {
                    "NewGameSessionsPerCreator": 5,
                    "PolicyPeriodInMinutes": 2
                },
                "RuntimeConfiguration": {
                    "GameSessionActivationTimeoutSeconds": 300,
                    "MaxConcurrentGameSessionActivations": 1,
                    "ServerProcesses": [
                        {
                            "ConcurrentExecutions": 1,
                            "LaunchPath": "/local/game/myscript.js"
                        }
                    ]
                },
                "ScriptId": {
                    "Fn::GetAtt": [
                        "ScriptResource",
                        "Id"
                    ]
                },
                "Locations": [
                  "us-west-2",
                  "us-east-1",
                  "eu-west-1"
                ]
            }
        }
    }
}
```

#### YAML

```yaml

Resources:
  FleetResource:
    Type: AWS::GameLift::Fleet
    Properties:
      CertificateConfiguration:
        CertificateType: DISABLED
      Description: Description of my game fleet
      DesiredEC2Instances: 1
      EC2InboundPermissions:
        - FromPort: 1234
          ToPort: 1324
          IpRange: 0.0.0.0/24
          Protocol: TCP
        - FromPort: 1356
          ToPort: 1578
          IpRange: 192.168.0.0/24
          Protocol: UDP
      EC2InstanceType: c4.large
      FleetType: SPOT
      LogPaths:
        - '/local/game/testlog.log'
        - '/local/game/testlog2.log'
      MetricGroups:
        - MetricGroupName
      Name: MyGameFleet
      NewGameSessionProtectionPolicy: FullProtection
      ResourceCreationLimitPolicy:
        NewGameSessionsPerCreator: 5
        PolicyPeriodInMinutes: 2
      RuntimeConfiguration:
        GameSessionActivationTimeoutSeconds: 300
        MaxConcurrentGameSessionActivations: 1
        ServerProcesses:
          - ConcurrentExecutions: 1
            LaunchPath: '/local/game/myscript.js'
      ScriptId: !GetAtt ScriptResource.Id
      Locations:
          - Location: 'us-west-2'
          - Location: 'us-east-1'
          - Location: 'eu-west-1'
```

## See also

- [Create GameLift resources using Amazon CloudFront](https://docs.aws.amazon.com/gamelift/latest/developerguide/resources-cloudformation.html) in the _Amazon_
_GameLift Developer Guide_

- [Setting up GameLift fleets](https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html) in the _Amazon GameLift Developer_
_Guide_

- [CreateFleet](https://docs.aws.amazon.com/gamelift/latest/apireference/API_CreateFleet.html) in the _Amazon GameLift API Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AnywhereConfiguration
