This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WorkSpaces::WorkspacesPool TimeoutSettings

Describes the timeout settings for the pool.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DisconnectTimeoutInSeconds" : Integer,
  "IdleDisconnectTimeoutInSeconds" : Integer,
  "MaxUserDurationInSeconds" : Integer
}

```

### YAML

```yaml

  DisconnectTimeoutInSeconds: Integer
  IdleDisconnectTimeoutInSeconds: Integer
  MaxUserDurationInSeconds: Integer

```

## Properties

`DisconnectTimeoutInSeconds`

Specifies the amount of time, in seconds, that a streaming session remains active after
users disconnect. If users try to reconnect to the streaming session after a disconnection
or network interruption within the time set, they are connected to their previous session.
Otherwise, they are connected to a new session with a new streaming instance.

_Required_: No

_Type_: Integer

_Minimum_: `60`

_Maximum_: `36000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IdleDisconnectTimeoutInSeconds`

The amount of time in seconds a connection will stay active while idle.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `36000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxUserDurationInSeconds`

Specifies the maximum amount of time, in seconds, that a streaming session can remain
active. If users are still connected to a streaming instance five minutes before this limit
is reached, they are prompted to save any open documents before being disconnected. After
this time elapses, the instance is terminated and replaced by a new instance.

_Required_: No

_Type_: Integer

_Minimum_: `600`

_Maximum_: `432000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Capacity

Next

All content copied from https://docs.aws.amazon.com/.
