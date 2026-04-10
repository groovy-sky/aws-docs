This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lightsail::Container HealthCheckConfig

`HealthCheckConfig` is a property of the [PublicEndpoint](../userguide/aws-properties-lightsail-container-publicendpoint.md) property. It describes the healthcheck configuration of a
container deployment on a container service.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "HealthyThreshold" : Integer,
  "IntervalSeconds" : Integer,
  "Path" : String,
  "SuccessCodes" : String,
  "TimeoutSeconds" : Integer,
  "UnhealthyThreshold" : Integer
}

```

### YAML

```yaml

  HealthyThreshold: Integer
  IntervalSeconds: Integer
  Path: String
  SuccessCodes: String
  TimeoutSeconds: Integer
  UnhealthyThreshold: Integer

```

## Properties

`HealthyThreshold`

The number of consecutive health check successes required before moving the container
to the `Healthy` state. The default value is `2`.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IntervalSeconds`

The approximate interval, in seconds, between health checks of an individual container.
You can specify between `5` and `300` seconds. The default value is
`5`.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Path`

The path on the container on which to perform the health check. The default value is
`/`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SuccessCodes`

The HTTP codes to use when checking for a successful response from a container. You can
specify values between `200` and `499`. You can specify multiple
values (for example, `200,202`) or a range of values (for example,
`200-299`).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeoutSeconds`

The amount of time, in seconds, during which no response means a failed health check.
You can specify between `2` and `60` seconds. The default value is
`2`.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UnhealthyThreshold`

The number of consecutive health check failures required before moving the container to
the `Unhealthy` state. The default value is `2`.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EnvironmentVariable

PortInfo

All content copied from https://docs.aws.amazon.com/.
