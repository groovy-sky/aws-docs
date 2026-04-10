This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GameLift::MatchmakingConfiguration GameProperty

This key-value pair can store custom data about a game session. For example, you might
use a `GameProperty` to track a game session's map, level of difficulty, or
remaining time. The difficulty level could be specified like this: `{"Key":
                "difficulty", "Value":"Novice"}`.

You can set game properties when creating a game session. You can also modify game
properties of an active game session. When searching for game sessions, you can filter
on game property keys and values. You can't delete game properties from a game session.

For examples of working with game properties, see [Create a game session with properties](../../../../reference/gamelift/latest/developerguide/gamelift-sdk-client-api.md#game-properties).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "Value" : String
}

```

### YAML

```yaml

  Key: String
  Value: String

```

## Properties

`Key`

The game property identifier.

###### Note

- Avoid using periods (".") in property keys if you plan to search for game sessions by properties. Property keys containing periods cannot be searched and will be filtered out from search results due to search index limitations.

- If you use SearchGameSessions API, there is a limit of 500 game property keys across all game sessions and all fleets per region. If the limit is exceeded, there will potentially be game session entries missing from SearchGameSessions API results.

_Required_: Yes

_Type_: String

_Maximum_: `32`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The game property value.

_Required_: Yes

_Type_: String

_Maximum_: `96`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Create GameLift resources using Amazon CloudFront](../../../gamelift/latest/developerguide/resources-cloudformation.md) in the _Amazon_
_GameLift Developer Guide_

- [Design a FlexMatch\
matchmaker](../../../gamelift/latest/flexmatchguide/match-configuration.md) in the _Amazon GameLift Developer Guide_

- [GameProperty](../../../../reference/gamelift/latest/apireference/api-gameproperty.md) in the _Amazon GameLift API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::GameLift::MatchmakingConfiguration

Tag

All content copied from https://docs.aws.amazon.com/.
