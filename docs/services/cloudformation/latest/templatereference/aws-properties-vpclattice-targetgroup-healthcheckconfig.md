This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::VpcLattice::TargetGroup HealthCheckConfig

Describes the health check configuration of a target group. Health check configurations
aren't used for target groups of type `LAMBDA` or `ALB`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Enabled" : Boolean,
  "HealthCheckIntervalSeconds" : Integer,
  "HealthCheckTimeoutSeconds" : Integer,
  "HealthyThresholdCount" : Integer,
  "Matcher" : Matcher,
  "Path" : String,
  "Port" : Integer,
  "Protocol" : String,
  "ProtocolVersion" : String,
  "UnhealthyThresholdCount" : Integer
}

```

### YAML

```yaml

  Enabled: Boolean
  HealthCheckIntervalSeconds: Integer
  HealthCheckTimeoutSeconds: Integer
  HealthyThresholdCount: Integer
  Matcher:
    Matcher
  Path: String
  Port: Integer
  Protocol: String
  ProtocolVersion: String
  UnhealthyThresholdCount: Integer

```

## Properties

`Enabled`

Indicates whether health checking is enabled.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HealthCheckIntervalSeconds`

The approximate amount of time, in seconds, between health checks of an individual target.
The range is 5–300 seconds. The default is 30 seconds.

_Required_: No

_Type_: Integer

_Minimum_: `5`

_Maximum_: `300`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HealthCheckTimeoutSeconds`

The amount of time, in seconds, to wait before reporting a target as unhealthy. The range is
1–120 seconds. The default is 5 seconds.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `120`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HealthyThresholdCount`

The number of consecutive successful health checks required before considering an unhealthy
target healthy. The range is 2–10. The default is 5.

_Required_: No

_Type_: Integer

_Minimum_: `2`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Matcher`

The codes to use when checking for a successful response from a target.

_Required_: No

_Type_: [Matcher](aws-properties-vpclattice-targetgroup-matcher.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Path`

The destination for health checks on the targets. If the protocol version is
`HTTP/1.1` or `HTTP/2`, specify a valid URI (for example,
`/path?query`). The default path is `/`. Health checks are not supported
if the protocol version is `gRPC`, however, you can choose `HTTP/1.1` or
`HTTP/2` and specify a valid URI.

_Required_: No

_Type_: String

_Pattern_: `(^/[a-zA-Z0-9@:%_+.~#?&/=-]*$|(^$))`

_Minimum_: `0`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Port`

The port used when performing health checks on targets. The default setting is the port that
a target receives traffic on.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `65535`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Protocol`

The protocol used when performing health checks on targets. The possible protocols are
`HTTP` and `HTTPS`. The default is `HTTP`.

_Required_: No

_Type_: String

_Allowed values_: `HTTP | HTTPS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProtocolVersion`

The protocol version used when performing health checks on targets. The possible protocol
versions are `HTTP1` and `HTTP2`.

_Required_: No

_Type_: String

_Allowed values_: `HTTP1 | HTTP2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UnhealthyThresholdCount`

The number of consecutive failed health checks required before considering a target
unhealthy. The range is 2–10. The default is 2.

_Required_: No

_Type_: Integer

_Minimum_: `2`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::VpcLattice::TargetGroup

Matcher

All content copied from https://docs.aws.amazon.com/.
