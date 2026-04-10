This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GameLift::MatchmakingConfiguration

The `AWS::GameLift::MatchmakingConfiguration` resource defines a new
matchmaking configuration for use with FlexMatch. Whether you're using FlexMatch with GameLift hosting or as a
standalone matchmaking service, the matchmaking configuration sets out rules for matching players and forming teams.
If you're using GameLift hosting, it also defines how to start game sessions for each match. Your matchmaking system
can use multiple configurations to handle different game scenarios. All matchmaking requests identify the
matchmaking configuration to use and provide player attributes that are consistent with that configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::GameLift::MatchmakingConfiguration",
  "Properties" : {
      "AcceptanceRequired" : Boolean,
      "AcceptanceTimeoutSeconds" : Integer,
      "AdditionalPlayerCount" : Integer,
      "BackfillMode" : String,
      "CreationTime" : String,
      "CustomEventData" : String,
      "Description" : String,
      "FlexMatchMode" : String,
      "GameProperties" : [ GameProperty, ... ],
      "GameSessionData" : String,
      "GameSessionQueueArns" : [ String, ... ],
      "Name" : String,
      "NotificationTarget" : String,
      "RequestTimeoutSeconds" : Integer,
      "RuleSetArn" : String,
      "RuleSetName" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::GameLift::MatchmakingConfiguration
Properties:
  AcceptanceRequired: Boolean
  AcceptanceTimeoutSeconds: Integer
  AdditionalPlayerCount: Integer
  BackfillMode: String
  CreationTime: String
  CustomEventData: String
  Description: String
  FlexMatchMode: String
  GameProperties:
    - GameProperty
  GameSessionData: String
  GameSessionQueueArns:
    - String
  Name: String
  NotificationTarget: String
  RequestTimeoutSeconds: Integer
  RuleSetArn: String
  RuleSetName: String
  Tags:
    - Tag

```

## Properties

`AcceptanceRequired`

A flag that determines whether a match that was created with this configuration must
be accepted by the matched players. To require acceptance, set to `TRUE`.
With this option enabled, matchmaking tickets use the status
`REQUIRES_ACCEPTANCE` to indicate when a completed potential match is
waiting for player acceptance.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AcceptanceTimeoutSeconds`

The length of time (in seconds) to wait for players to accept a proposed match, if
acceptance is required.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AdditionalPlayerCount`

The number of player slots in a match to keep open for future players. For example, if the configuration's rule set specifies
a match for a single 12-person team, and the additional player count is set to 2, only 10 players are selected for the match. This parameter is not used if `FlexMatchMode` is set to
`STANDALONE`.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BackfillMode`

The method used to backfill game sessions that are created with this matchmaking
configuration. Specify `MANUAL` when your game manages backfill requests manually
or does not use the match backfill feature. Specify `AUTOMATIC` to have GameLift
create a `StartMatchBackfill` request whenever a game session has one or more open
slots. Learn more about manual and automatic backfill in [Backfill Existing Games with\
FlexMatch](../../../gamelift/latest/flexmatchguide/match-backfill.md). Automatic backfill is not
available when `FlexMatchMode` is set to `STANDALONE`.

_Required_: No

_Type_: String

_Allowed values_: `AUTOMATIC | MANUAL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreationTime`

A time stamp indicating when this data object was created. Format is a number expressed in Unix time as milliseconds (for example `"1469498468.057"`).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomEventData`

Information to add to all events related to the matchmaking configuration.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description for the matchmaking configuration.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FlexMatchMode`

Indicates whether this matchmaking configuration is being used with Amazon GameLift Servers hosting or
as a standalone matchmaking solution.

- **STANDALONE** \- FlexMatch forms matches and
returns match information, including players and team assignments, in a [MatchmakingSucceeded](../../../gamelift/latest/flexmatchguide/match-events.md#match-events-matchmakingsucceeded) event.

- **WITH\_QUEUE** \- FlexMatch forms matches and uses
the specified Amazon GameLift Servers queue to start a game session for the match.

_Required_: No

_Type_: String

_Allowed values_: `STANDALONE | WITH_QUEUE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GameProperties`

A set of custom properties for a game session, formatted as key-value pairs. These
properties are passed to a game server process with a request to start a new game session. See
[Start a Game Session](../../../../reference/gamelift/latest/developerguide/gamelift-sdk-server-api.md#gamelift-sdk-server-startsession).
This parameter is not used if `FlexMatchMode` is set to `STANDALONE`.

_Required_: No

_Type_: Array of [GameProperty](aws-properties-gamelift-matchmakingconfiguration-gameproperty.md)

_Maximum_: `16`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GameSessionData`

A set of custom game session properties, formatted as a single string value. This
data is passed to a game server process with a request to start a new game session.
See [Start a Game Session](../../../../reference/gamelift/latest/developerguide/gamelift-sdk-server-api.md#gamelift-sdk-server-startsession).
This parameter is not used if `FlexMatchMode` is set to `STANDALONE`.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GameSessionQueueArns`

The Amazon Resource Name ( [ARN](../../../s3/latest/dev/s3-arn-format.md)) that is assigned to a Amazon GameLift Servers game session queue resource and uniquely identifies it. ARNs are unique across all Regions. Format is `arn:aws:gamelift:<region>::gamesessionqueue/<queue name>`. Queues can be located in any Region. Queues are used to start new
Amazon GameLift Servers-hosted game sessions for matches that are created with this matchmaking
configuration. If `FlexMatchMode` is set to `STANDALONE`, do not
set this parameter.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

A unique identifier for the matchmaking configuration. This name is used to identify the configuration associated with a matchmaking
request or ticket.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9-\.]*`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NotificationTarget`

An SNS topic ARN that is set up to receive matchmaking notifications. See [Setting up notifications for matchmaking](../../../gamelift/latest/flexmatchguide/match-notification.md) for more information.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9:_/-]*(.fifo)?`

_Minimum_: `0`

_Maximum_: `300`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RequestTimeoutSeconds`

The maximum duration, in seconds, that a matchmaking ticket can remain in process
before timing out. Requests that fail due to timing out can be resubmitted as
needed.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `43200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuleSetArn`

The Amazon Resource Name ( [ARN](../../../s3/latest/dev/s3-arn-format.md)) associated with the GameLift matchmaking rule set resource that this
configuration uses.

_Required_: No

_Type_: String

_Pattern_: `^arn:.*:matchmakingruleset\/[a-zA-Z0-9-\.]*`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuleSetName`

A unique identifier for the matchmaking rule set to use with this configuration. You can use either the rule set name or ARN
value. A matchmaking configuration can only use rule sets that are defined in the same
Region.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9-\.]*`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A list of labels to assign to the new matchmaking configuration resource. Tags are developer-defined
key-value pairs. Tagging
AWS resources are useful for resource management, access management and cost allocation.
For more information, see [Tagging AWS Resources](../../../../general/latest/gr/aws-tagging.md) in the
_AWS General Reference_. Once the resource is created, you can
use TagResource, UntagResource, and
ListTagsForResource to add, remove, and view tags. The
maximum tag limit may be lower than stated. See the AWS General Reference for actual
tagging limits.

_Required_: No

_Type_: Array of [Tag](aws-properties-gamelift-matchmakingconfiguration-tag.md)

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `MatchmakingConfiguration` name, which is unique.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The unique Amazon Resource Name (ARN) for the
`MatchmakingConfiguration`.

`Name`

The `MatchmakingConfiguration` name, which is unique.

## Examples

- [Create a matchmaking configuration for use with GameLift managed hosting](#aws-resource-gamelift-matchmakingconfiguration--examples--Create_a_matchmaking_configuration_for_use_with_GameLift_managed_hosting)

- [Create a matchmaking configuration for a standalone FlexMatch system](#aws-resource-gamelift-matchmakingconfiguration--examples--Create_a_matchmaking_configuration_for_a_standalone_FlexMatch_system)

### Create a matchmaking configuration for use with GameLift managed hosting

The following example creates a matchmaking configuration for a game that is being hosted on
GameLift servers, identifying a game session queue and providing a set of game properties to be passed on
to new game sessions. Player acceptance is required, with a 60-second timeout, and auto-backfill is enabled.
The example uses a
`Ref` intrinsic function to specify the game session queue and rule set, which are also in
the template.

#### JSON

```json

{
    "Resources": {
        "QueueResource": {
            "Type": "AWS::GameLift::GameSessionQueue",
            "Properties": {
                "Name": "MyGameSessionQueue"
            }
        },
        "MatchmakingRuleSetResource": {
            "Type": "AWS::GameLift::MatchmakingRuleSet",
            "Properties": {
                "Name": "MyRuleSet",
                "RuleSetBody": {
                  "Fn::Sub": "{\"name\": \"MyMatchmakingRuleSet\",\"ruleLanguageVersion\": \"1.0\", \"teams\": [{\"name\": \"MyTeam\",\"minPlayers\": 1,\"maxPlayers\": 20}]}"
                }
            }
        },
        "MatchmakingConfigurationResource": {
            "Type": "AWS::GameLift::MatchmakingConfiguration",
            "Properties": {
                "Name": "MyMatchmakingConfiguration",
                "AcceptanceRequired": true,
                "AcceptanceTimeoutSeconds": 60,
                "AdditionalPlayerCount": 8,
                "BackfillMode": "AUTOMATIC",
                "CustomEventData": "MyCustomEventData",
                "Description": "A basic matchmaking configuration for a GameLift-hosted game",
                "FlexMatchMode": "WITH_QUEUE",
                "GameSessionData": "MyGameSessionData",
                "GameProperties": [
                    {
                        "Key": "level",
                        "Value": "10"
                    },
                    {
                        "Key": "gameMode",
                        "Value": "hard"
                    }
                ],
                "GameSessionQueueArns": [
                    {
                        "Fn::GetAtt": [
                            "QueueResource",
                            "Arn"
                        ]
                    }
                ],
                "RequestTimeoutSeconds": 100,
                "RuleSetName": {
                    "Ref": "MatchmakingRuleSetResource"
                }
            },
            "DependsOn": [
                "QueueResource",
                "MatchmakingRuleSetResource"
            ]
        }
    }
}
```

#### YAML

```yaml

Resources:
  QueueResource:
    Type: "AWS::GameLift::GameSessionQueue"
    Properties:
      Name: "MyGameSessionQueue"
  MatchmakingRuleSetResource:
    Type: "AWS::GameLift::MatchmakingRuleSet"
    Properties:
      Name: "MyRuleSet"
      # Rule set body for a game of 20 players
      RuleSetBody: !Sub |
        {
          "name": "MyMatchmakingRuleSet",
           "ruleLanguageVersion": "1.0",
           "teams": [{
                       "name": "MyTeam",
                       "minPlayers": 1,
                       "maxPlayers": 20
                     }]
        }
  MatchmakingConfigurationResource:
    Type: "AWS::GameLift::MatchmakingConfiguration"
    Properties:
      Name: "MyMatchmakingConfiguration"
      AcceptanceRequired: true
      AcceptanceTimeoutSeconds: 60
      AdditionalPlayerCount: 8
      BackfillMode: "AUTOMATIC"
      CustomEventData: "MyCustomEventData"
      Description: "A basic matchmaking configuration for a GameLift-hosted game"
      FlexMatchMode: "WITH_QUEUE"
      GameSessionData: "MyGameSessionData"
      GameProperties:
        - Key: "level"
          Value: "10"
        - Key: "gameMode"
          Value: "hard"
      GameSessionQueueArns:
        - !GetAtt QueueResource.Arn
      RequestTimeoutSeconds: 100
      RuleSetName: !Ref MatchmakingRuleSetResource
    DependsOn:
      - QueueResource
      - MatchmakingRuleSetResource
```

### Create a matchmaking configuration for a standalone FlexMatch system

The following example creates a matchmaking configuration for a game that is hosted on resources
other than GameLift game servers. This includes games that are hosted on Amazon EC2 with GameLift FleetIQ.
This configuration omits the game session queue, game properties and session data, and additional player count.
Player acceptance is required, with a 60-second timeout. It uses a
`Ref` intrinsic function to specify the rule set, which is also in
the template.

#### JSON

```json

{
    "Resources": {
        "MatchmakingRuleSetResource": {
            "Type": "AWS::GameLift::MatchmakingRuleSet",
            "Properties": {
                "Name": "MyRuleSet",
                "RuleSetBody": {
                    "Fn::Sub": "{\"name\": \"MyMatchmakingRuleSet\",\"ruleLanguageVersion\": \"1.0\", \"teams\": [{\"name\": \"MyTeam\",\"minPlayers\": 1,\"maxPlayers\": 20}]}"
                }
            }
        },
        "MatchmakingConfigurationResource": {
            "Type": "AWS::GameLift::MatchmakingConfiguration",
            "Properties": {
                "Name": "MyMatchmakingConfiguration",
                "AcceptanceRequired": true,
                "AcceptanceTimeoutSeconds": 60,
                "BackfillMode": "MANUAL",
                "CustomEventData": "MyCustomEventData",
                "Description": "A basic standalone matchmaking configuration",
                "FlexMatchMode": "STANDALONE",
                "RequestTimeoutSeconds": 100,
                "RuleSetName": {
                    "Ref": "MatchmakingRuleSetResource"
                }
            },
            "DependsOn": [
                "MatchmakingRuleSetResource"
            ]
        }
    }
}
```

#### YAML

```yaml

Resources:
    MatchmakingRuleSetResource:
        Type: "AWS::GameLift::MatchmakingRuleSet"
        Properties:
            Name: "MyRuleSet"
            # Rule set body for a game of 20 players
            RuleSetBody: !Sub |
            {
                "name": "MyMatchmakingRuleSet",
                "ruleLanguageVersion": "1.0",
                "teams": [{
                    "name": "MyTeam",
                    "minPlayers": 1,
                    "maxPlayers": 20
                }]
            }
    MatchmakingConfigurationResource:
        Type: "AWS::GameLift::MatchmakingConfiguration"
        Properties:
            Name: "MyMatchmakingConfiguration"
            AcceptanceRequired: true
            AcceptanceTimeoutSeconds: 60
            BackfillMode: "MANUAL"
            CustomEventData: "MyCustomEventData"
            Description: "A basic standalone matchmaking configuration"
            FlexMatchMode: "STANDALONE"
            RequestTimeoutSeconds: 100
            RuleSetName: !Ref MatchmakingRuleSetResource
        DependsOn:
          - MatchmakingRuleSetResource
```

## See also

- [Create GameLift resources using Amazon CloudFront](../../../gamelift/latest/developerguide/resources-cloudformation.md) in the _Amazon_
_GameLift Developer Guide_

- [Setting up GameLift FlexMatch matchmakers](../../../gamelift/latest/flexmatchguide/matchmaker-build.md) in the _Amazon GameLift_
_Developer Guide_

- [MatchmakingConfiguration](../../../../reference/gamelift/latest/apireference/api-matchmakingconfiguration.md) in the _Amazon GameLift API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

GameProperty

All content copied from https://docs.aws.amazon.com/.
