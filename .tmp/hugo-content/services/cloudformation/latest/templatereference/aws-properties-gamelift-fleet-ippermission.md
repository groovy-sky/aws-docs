This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GameLift::Fleet IpPermission

A range of IP addresses and port settings that allow inbound traffic to connect to
server processes on an instance in a fleet. New game sessions are assigned an IP
address/port number combination, which must fall into the fleet's allowed ranges. Fleets
with custom game builds must have permissions explicitly set. For Realtime Servers fleets, GameLift
automatically opens two port ranges, one for TCP messaging and one for UDP.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FromPort" : Integer,
  "IpRange" : String,
  "Protocol" : String,
  "ToPort" : Integer
}

```

### YAML

```yaml

  FromPort: Integer
  IpRange: String
  Protocol: String
  ToPort: Integer

```

## Properties

`FromPort`

A starting value for a range of allowed port numbers.

For fleets using Linux builds, only ports `22` and `1026-60000` are valid.

For fleets using Windows builds, only ports `1026-60000` are valid.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `60000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IpRange`

A range of allowed IP addresses. This value must be expressed in CIDR notation.
Example: " `000.000.000.000/[subnet mask]`" or optionally the shortened
version " `0.0.0.0/[subnet mask]`".

_Required_: Yes

_Type_: String

_Pattern_: `(^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])(/([0-9]|[1-2][0-9]|3[0-2]))$)`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Protocol`

The network communication protocol used by the fleet.

_Required_: Yes

_Type_: String

_Allowed values_: `TCP | UDP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ToPort`

An ending value for a range of allowed port numbers. Port numbers are end-inclusive.
This value must be equal to or greater than `FromPort`.

For fleets using Linux builds, only ports `22` and `1026-60000` are valid.

For fleets using Windows builds, only ports `1026-60000` are valid.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `60000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Create GameLift resources using Amazon CloudFront](../../../gamelift/latest/developerguide/resources-cloudformation.md) in the _Amazon_
_GameLift Developer Guide_

- [Deploy a GameLift fleet for a custom game build](../../../gamelift/latest/developerguide/fleets-creating.md) in the _Amazon_
_GameLift Developer Guide_

- [IpPermission](../../../../reference/gamelift/latest/apireference/api-ippermission.md) in the _Amazon GameLift API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CertificateConfiguration

LocationCapacity

All content copied from https://docs.aws.amazon.com/.
