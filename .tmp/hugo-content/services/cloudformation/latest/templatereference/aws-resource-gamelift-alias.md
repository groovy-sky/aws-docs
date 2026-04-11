This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GameLift::Alias

The `AWS::GameLift::Alias` resource creates an alias for an Amazon GameLift
(GameLift) fleet destination. There are two types of routing strategies for aliases: simple
and terminal. A simple alias points to an active fleet. A terminal alias displays a message
instead of routing players to an active fleet. For example, a terminal alias might display a
URL link that directs players to an upgrade site. You can use aliases to define destinations
in a game session queue or when requesting new game sessions.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::GameLift::Alias",
  "Properties" : {
      "Description" : String,
      "Name" : String,
      "RoutingStrategy" : RoutingStrategy,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::GameLift::Alias
Properties:
  Description: String
  Name: String
  RoutingStrategy:
    RoutingStrategy
  Tags:
    - Tag

```

## Properties

`Description`

A human-readable description of the alias.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

A descriptive label that is associated with an alias. Alias names do not need to be unique.

_Required_: Yes

_Type_: String

_Pattern_: `.*\S.*`

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoutingStrategy`

The routing configuration, including routing type and fleet target, for the alias.

_Required_: Yes

_Type_: [RoutingStrategy](aws-properties-gamelift-alias-routingstrategy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A list of labels to assign to the new alias resource. Tags are developer-defined
key-value pairs. Tagging AWS resources are useful for resource management, access
management and cost allocation. For more information, see [Tagging AWS Resources](../../../../general/latest/gr/aws-tagging.md) in the
_AWS General Reference_.

_Required_: No

_Type_: Array of [Tag](aws-properties-gamelift-alias-tag.md)

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the alias ID, such as
`alias-1111aaaa-22bb-33cc-44dd-5555eeee66ff`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

`AliasArn`

The Amazon Resource Name ( [ARN](../../../s3/latest/dev/s3-arn-format.md)) that is assigned to a Amazon GameLift Servers alias resource and uniquely identifies it. ARNs are unique across all Regions. Format is `arn:aws:gamelift:<region>::alias/alias-a1234567-b8c9-0d1e-2fa3-b45c6d7e8912`. In a GameLift alias ARN, the resource ID matches the alias ID value.

`AliasId`

A unique identifier for the alias. For example, `arn:aws:gamelift:us-west-1::alias/alias-a1234567-b8c9-0d1e-2fa3-b45c6d7e8912`

Alias IDs are unique within a Region.

## Examples

- [Create a Simple Alias](#aws-resource-gamelift-alias--examples--Create_a_Simple_Alias)

- [Create terminal alias](#aws-resource-gamelift-alias--examples--Create_terminal_alias)

### Create a Simple Alias

The following example creates a simple alias for a fleet. The template uses `!Ref`
to reference a fleet resource, which can be declared elsewhere in the same
template.

#### JSON

```json

{
    "Resources": {
        "AliasResource": {
            "Properties": {
                "Name": "MyAlias",
                "Description": "My Alias Description",
                "RoutingStrategy": {
                    "Type": "SIMPLE",
                    "FleetId": {
                        "Ref": "FleetResource"
                    }
                }
            },
            "Type": "AWS::GameLift::Alias"
        }
    }
}
```

#### YAML

```yaml

Resources:
  AliasResource:
    Properties:
      Name: MyAlias
      Description: My Alias Description
      RoutingStrategy:
        Type: SIMPLE
        FleetId: !Ref FleetResource
    Type: AWS::GameLift::Alias
```

### Create terminal alias

The following example creates a terminal alias with a
generic terminal message.

#### JSON

```json

{
    "Resources": {
        "AliasResource": {
            "Type": "AWS::GameLift::Alias",
            "Properties": {
                "Name": "MyTerminalAlias",
                "Description": "A terminal alias",
                "RoutingStrategy": {
                    "Type": "TERMINAL",
                    "Message": "Terminal routing strategy message"
                }
            }
        }
    }
}
```

#### YAML

```yaml

Resources:
  AliasResource:
    Type: AWS::GameLift::Alias
    Properties:
      Name: MyTerminalAlias
      Description: A terminal alias
      RoutingStrategy:
        Type: TERMINAL
        Message: Terminal routing strategy message
```

## See also

- [Create GameLift resources using Amazon CloudFront](../../../gamelift/latest/developerguide/resources-cloudformation.md) in the _Amazon_
_GameLift Developer Guide_

- [Add an alias to a GameLift fleet](../../../gamelift/latest/developerguide/aliases-creating.md) in the _Amazon GameLift Developer_
_Guide_

- [CreateAlias](../../../../reference/gamelift/latest/apireference/api-createalias.md) in the _Amazon GameLift API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon GameLift Servers

RoutingStrategy

All content copied from https://docs.aws.amazon.com/.
