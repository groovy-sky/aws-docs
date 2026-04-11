This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GameLift::MatchmakingRuleSet

Creates a new rule set for FlexMatch matchmaking. A rule set describes the type of match
to create, such as the number and size of teams. It also sets the parameters for
acceptable player matches, such as minimum skill level or character type.

To create a matchmaking rule set, provide unique rule set name and the rule set body
in JSON format. Rule sets must be defined in the same Region as the matchmaking
configuration they are used with.

Since matchmaking rule sets cannot be edited, it is a good idea to check the rule
set syntax.

**Learn more**

- [Build a rule\
set](../../../gamelift/latest/flexmatchguide/match-rulesets.md)

- [Design a\
matchmaker](../../../gamelift/latest/flexmatchguide/match-configuration.md)

- [Matchmaking with\
FlexMatch](../../../gamelift/latest/flexmatchguide/match-intro.md)

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::GameLift::MatchmakingRuleSet",
  "Properties" : {
      "Name" : String,
      "RuleSetBody" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::GameLift::MatchmakingRuleSet
Properties:
  Name: String
  RuleSetBody: String
  Tags:
    - Tag

```

## Properties

`Name`

A unique identifier for the matchmaking rule set. A matchmaking configuration identifies the rule set it uses by this name
value. Note that the rule set name is different from the optional `name`
field in the rule set body.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9-\.]*`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RuleSetBody`

A collection of matchmaking rules, formatted as a JSON string. Comments are not
allowed in JSON, but most elements support a description field.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `65535`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

A list of labels to assign to the new matchmaking rule set resource. Tags are developer-defined
key-value pairs. Tagging
AWS resources are useful for resource management, access management and cost allocation.
For more information, see [Tagging AWS Resources](../../../../general/latest/gr/aws-tagging.md) in the
_AWS General Reference_. Once the resource is created, you can
use TagResource, UntagResource, and
ListTagsForResource to add, remove, and view tags. The
maximum tag limit may be lower than stated. See the AWS General Reference for actual
tagging limits.

_Required_: No

_Type_: Array of [Tag](aws-properties-gamelift-matchmakingruleset-tag.md)

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the rule set name, which is unique within each Region.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The unique Amazon Resource Name (ARN) assigned to the rule set.

`CreationTime`

A time stamp indicating when this data object was created. Format is a number expressed in Unix time as milliseconds (for example `"1469498468.057"`).

`Name`

The unique name of the rule set.

## Examples

### Create a Matchmaking Rule Set

The following example creates a matchmaking rule set for GameLift FlexMatch named `MyRuleSet`.
The simple rule set defines a match with one team containing 1 to 20 players. In the YAML example,
since RuleSetBody must be in JSON format, the !Sub command is used to specify JSON content within the YAML format.

#### JSON

```json

{
  "Resources": {
    "MatchmakingRuleSet": {
      "Type": "AWS::GameLift::MatchmakingRuleSet",
      "Properties": {
        "Name": "MyRuleSet",
        "RuleSetBody": {
          "Fn::Sub": "{\"name\": \"MyMatchmakingRuleSet\",\"ruleLanguageVersion\": \"1.0\", \"teams\": [{\"name\": \"MyTeam\",\"minPlayers\": 1,\"maxPlayers\": 20}]}"
        }
      }
    }
  }
}

```

#### YAML

```yaml

Resources:
  MatchmakingRuleSet:
    Type: "AWS::GameLift::MatchmakingRuleSet"
    Properties:
      Name: "MyRuleSet"
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
```

## See also

- [Create GameLift Resources Using Amazon CloudFront](../../../gamelift/latest/developerguide/resources-cloudformation.md) in the _Amazon_
_GameLift Developer Guide_

- [Build a FlexMatch Rule Set](../../../gamelift/latest/flexmatchguide/match-rulesets.md) in the _Amazon GameLift Developer Guide_

- [CreateMatchmakingRuleSet](../../../../reference/gamelift/latest/apireference/api-creatematchmakingruleset.md) in the _Amazon GameLift API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
