This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GameLift::Alias RoutingStrategy

The routing configuration for a fleet alias.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FleetId" : String,
  "Message" : String,
  "Type" : String
}

```

### YAML

```yaml

  FleetId: String
  Message: String
  Type: String

```

## Properties

`FleetId`

A unique identifier for a fleet that the alias points to. If you specify
`SIMPLE` for the `Type` property, you must specify this
property.

_Required_: Conditional

_Type_: String

_Pattern_: `^[a-z]*fleet-[a-zA-Z0-9\-]+`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Message`

The message text to be used with a terminal routing strategy. If you specify
`TERMINAL` for the `Type` property, you must specify this
property.

_Required_: Conditional

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

A type of routing strategy.

Possible routing types include the following:

- **SIMPLE** \- The alias resolves to one specific fleet. Use
this type when routing to active fleets.

- **TERMINAL** \- The alias does not resolve to a fleet but
instead can be used to display a message to the user. A terminal alias throws a
`TerminalRoutingStrategyException` with the message that you specified in the
`Message` property.

_Required_: Yes

_Type_: String

_Allowed values_: `SIMPLE | TERMINAL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Create GameLift resources using Amazon CloudFront](https://docs.aws.amazon.com/gamelift/latest/developerguide/resources-cloudformation.html) in the _Amazon_
_GameLift Developer Guide_

- [Add an alias to a GameLift fleet](https://docs.aws.amazon.com/gamelift/latest/developerguide/aliases-creating.html) in the _Amazon GameLift Developer_
_Guide_

- [RoutingStrategy](https://docs.aws.amazon.com/gamelift/latest/apireference/API_RoutingStrategy.html) in the _Amazon GameLift API Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::GameLift::Alias

Tag
