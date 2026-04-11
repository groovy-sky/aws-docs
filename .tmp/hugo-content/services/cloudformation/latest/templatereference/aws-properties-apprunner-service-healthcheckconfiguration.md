This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppRunner::Service HealthCheckConfiguration

Describes the settings for the health check that AWS App Runner performs to monitor the health of a service.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "HealthyThreshold" : Integer,
  "Interval" : Integer,
  "Path" : String,
  "Protocol" : String,
  "Timeout" : Integer,
  "UnhealthyThreshold" : Integer
}

```

### YAML

```yaml

  HealthyThreshold: Integer
  Interval: Integer
  Path: String
  Protocol: String
  Timeout: Integer
  UnhealthyThreshold: Integer

```

## Properties

`HealthyThreshold`

The number of consecutive checks that must succeed before App Runner decides that the service is healthy.

Default: `1`

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Interval`

The time interval, in seconds, between health checks.

Default: `5`

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Path`

The URL that health check requests are sent to.

`Path` is only applicable when you set `Protocol` to `HTTP`.

Default: `"/"`

_Required_: No

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Protocol`

The IP protocol that App Runner uses to perform health checks for your service.

If you set `Protocol` to `HTTP`, App Runner sends health check requests to the HTTP path specified by `Path`.

Default: `TCP`

_Required_: No

_Type_: String

_Allowed values_: `TCP | HTTP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Timeout`

The time, in seconds, to wait for a health check response before deciding it failed.

Default: `2`

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UnhealthyThreshold`

The number of consecutive checks that must fail before App Runner decides that the service is unhealthy.

Default: `5`

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EncryptionConfiguration

ImageConfiguration

All content copied from https://docs.aws.amazon.com/.
