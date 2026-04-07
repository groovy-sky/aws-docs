This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GameLift::GameSessionQueue

The `AWS::GameLift::GameSessionQueue` resource creates a placement queue
that processes requests for new game sessions. A queue uses FleetIQ algorithms to determine
the best placement locations and find an available game server, then prompts the game server
to start a new game session. Queues can have destinations (GameLift fleets or aliases), which
determine where the queue can place new game sessions. A queue can have destinations with
varied fleet type (Spot and On-Demand), instance type, and AWS Region.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::GameLift::GameSessionQueue",
  "Properties" : {
      "CustomEventData" : String,
      "Destinations" : [ GameSessionQueueDestination, ... ],
      "FilterConfiguration" : FilterConfiguration,
      "Name" : String,
      "NotificationTarget" : String,
      "PlayerLatencyPolicies" : [ PlayerLatencyPolicy, ... ],
      "PriorityConfiguration" : PriorityConfiguration,
      "Tags" : [ Tag, ... ],
      "TimeoutInSeconds" : Integer
    }
}

```

### YAML

```yaml

Type: AWS::GameLift::GameSessionQueue
Properties:
  CustomEventData: String
  Destinations:
    - GameSessionQueueDestination
  FilterConfiguration:
    FilterConfiguration
  Name: String
  NotificationTarget: String
  PlayerLatencyPolicies:
    - PlayerLatencyPolicy
  PriorityConfiguration:
    PriorityConfiguration
  Tags:
    - Tag
  TimeoutInSeconds: Integer

```

## Properties

`CustomEventData`

Information to be added to all events that are related to this game session
queue.

_Required_: No

_Type_: String

_Pattern_: `[\s\S]*`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Destinations`

A list of fleets and/or fleet aliases that can be used to fulfill game session placement requests in the queue.
Destinations are identified by either a fleet ARN or a fleet alias ARN, and are listed in order of placement preference.

_Required_: No

_Type_: Array of [GameSessionQueueDestination](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-gamelift-gamesessionqueue-gamesessionqueuedestination.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FilterConfiguration`

A list of locations where a queue is allowed to place new game sessions. Locations
are specified in the form of AWS Region codes, such as `us-west-2`. If this parameter is
not set, game sessions can be placed in any queue location.

_Required_: No

_Type_: [FilterConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-gamelift-gamesessionqueue-filterconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

A descriptive label that is associated with game session queue. Queue names must be unique within each Region.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9-]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NotificationTarget`

An SNS topic ARN that is set up to receive game session placement notifications. See
[Setting up\
notifications for game session placement](https://docs.aws.amazon.com/gamelift/latest/developerguide/queue-notification.html).

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9:_-]*(\.fifo)?`

_Minimum_: `1`

_Maximum_: `300`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PlayerLatencyPolicies`

A set of policies that enforce a sliding cap on player latency when processing game sessions placement requests.
Use multiple policies to gradually relax the cap over time if Amazon GameLift Servers can't make a placement.
Policies are evaluated in order starting with the lowest maximum latency value.

_Required_: No

_Type_: Array of [PlayerLatencyPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-gamelift-gamesessionqueue-playerlatencypolicy.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PriorityConfiguration`

Custom settings to use when prioritizing destinations and locations for game session placements. This
configuration replaces the FleetIQ default prioritization process. Priority types that are not explicitly
named will be automatically applied at the end of the prioritization process.

_Required_: No

_Type_: [PriorityConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-gamelift-gamesessionqueue-priorityconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A list of labels to assign to the new game session queue resource. Tags are developer-defined
key-value pairs. Tagging
AWS resources are useful for resource management, access management and cost allocation.
For more information, see [Tagging AWS Resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the
_AWS General Reference_. Once the resource is created, you can
use TagResource, UntagResource, and
ListTagsForResource to add, remove, and view tags. The
maximum tag limit may be lower than stated. See the AWS General Reference for actual
tagging limits.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-gamelift-gamesessionqueue-tag.html)

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeoutInSeconds`

The maximum time, in seconds, that a new game session placement request remains in the queue. When a request exceeds this time, the game session placement changes to a `TIMED_OUT` status. If you don't specify a request timeout, the queue uses a default value.

###### Note

The minimum value is 10 and the maximum value is 600.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the game session queue, which is unique within each
Region.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The unique Amazon Resource Name (ARN) for the `GameSessionQueue`.

`Name`

A descriptive label that is associated with a game session queue. Names are unique within
each Region.

## Examples

### Create a Game Session Queue

The following example creates a GameLift game session queue named
`MyGameSessionQueue`. The queue is configured with two destinations, one
using a fleet ID and one using an alias ID. The queue has a latency policy.

#### JSON

```json

{
    "Resources": {
        "Queue": {
            "Type": "AWS::GameLift::GameSessionQueue",
            "Properties": {
                "Name": "MyGameSessionQueue",
                "TimeoutInSeconds": 60,
                "NotificationTarget": "arn:aws:sns:us-west-2:111122223333:My_Placement_SNS_Topic",
                "Destinations": [
                    {
                        "DestinationArn": "arn:aws:gamelift:us-west-2:012345678912:fleet/fleet-id"
                    },
                    {
                        "DestinationArn": "arn:aws:gamelift:us-west-2:012345678912:alias/alias-id"
                    }
                ],
                "PlayerLatencyPolicies": [
                    {
                        "MaximumIndividualPlayerLatencyMilliseconds": 1000,
                        "PolicyDurationSeconds": 60
                    }
                ],
                "PriorityConfiguration": {
                    "LocationOrder": [
                        "us-west-2",
                        "us-east-1"
                    ],
                    "PriorityOrder": [
                        "COST",
                        "LATENCY",
                        "LOCATION",
                        "DESTINATION"
                    ]
                },
                "FilterConfiguration": {
                    "AllowedLocations": [
                        "us-east-1",
                        "us-west-2"
                    ]
                }
            }
        }
    }
}
```

#### YAML

```yaml

Resources:
  Queue:
    Type: "AWS::GameLift::GameSessionQueue"
    Properties:
      Name: "MyGameSessionQueue"
      TimeoutInSeconds: 60
      NotificationTarget: "arn:aws:sns:us-west-2:111122223333:My_Placement_SNS_Topic"
      Destinations:
        # DestinationArn can be either an Alias arn or Fleet arn that you own
        - DestinationArn: "arn:aws:gamelift:us-west-2:012345678912:fleet/fleet-id"
        - DestinationArn: "arn:aws:gamelift:us-west-2:012345678912:alias/alias-id"
      PlayerLatencyPolicies:
        - MaximumIndividualPlayerLatencyMilliseconds: 1000
          PolicyDurationSeconds: 60
          PriorityConfiguration:
          LocationOrder:
          - 'us-west-2'
          - 'us-east-1'
          PriorityOrder:
          - 'COST'
          - 'LATENCY'
          - 'LOCATION'
          - 'DESTINATION'
      FilterConfiguration:
        AllowedLocations:
          - 'us-east-1'
          - 'us-west-2'
```

## See also

- [Create GameLift resources Using Amazon CloudFront](https://docs.aws.amazon.com/gamelift/latest/developerguide/resources-cloudformation.html) in the _Amazon_
_GameLift Developer Guide_

- [Setting up GameLift queues for game session placement](https://docs.aws.amazon.com/gamelift/latest/developerguide/queues-intro.html) in the _Amazon GameLift Developer Guide_

- [CreateGameSessionQueue](https://docs.aws.amazon.com/gamelift/latest/apireference/API_CreateGameSessionQueue.html) in the _Amazon GameLift API Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TargetTrackingConfiguration

FilterConfiguration
