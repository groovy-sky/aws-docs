This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElasticLoadBalancing::LoadBalancer HealthCheck

Specifies health check settings for your Classic Load Balancer.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "HealthyThreshold" : String,
  "Interval" : String,
  "Target" : String,
  "Timeout" : String,
  "UnhealthyThreshold" : String
}

```

### YAML

```yaml

  HealthyThreshold: String
  Interval: String
  Target: String
  Timeout: String
  UnhealthyThreshold: String

```

## Properties

`HealthyThreshold`

The number of consecutive health checks successes required before moving the instance to the `Healthy` state.

_Required_: Yes

_Type_: String

_Minimum_: `2`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Interval`

The approximate interval, in seconds, between health checks of an individual instance.

_Required_: Yes

_Type_: String

_Minimum_: `5`

_Maximum_: `300`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Target`

The instance being checked. The protocol is either TCP, HTTP, HTTPS, or SSL. The range of valid ports is one (1) through 65535.

TCP is the default, specified as a TCP: port pair, for example "TCP:5000". In this case, a health check simply attempts to open a TCP connection to the instance on the specified port. Failure to connect within the configured timeout is considered unhealthy.

SSL is also specified as SSL: port pair, for example, SSL:5000.

For HTTP/HTTPS, you must include a ping path in the string. HTTP is specified as a HTTP:port;/;PathToPing; grouping, for example "HTTP:80/weather/us/wa/seattle". In this case, a HTTP GET request is issued to the instance on the given port and path. Any answer other than "200 OK" within the timeout period is considered unhealthy.

The total length of the HTTP ping target must be 1024 16-bit Unicode characters or less.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Timeout`

The amount of time, in seconds, during which no response means a failed health check.

This value must be less than the `Interval` value.

_Required_: Yes

_Type_: String

_Minimum_: `2`

_Maximum_: `60`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UnhealthyThreshold`

The number of consecutive health check failures required before moving the instance to the `Unhealthy` state.

_Required_: Yes

_Type_: String

_Minimum_: `2`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [ConfigureHealthCheck](../../../../reference/elasticloadbalancing/2012-06-01/apireference/api-configurehealthcheck.md) in the
_Elastic Load Balancing API Reference (version 2012-06-01)_

- [Configure Health Checks](../../../elasticloadbalancing/latest/classic/elb-healthchecks.md)
in the _User Guide for Classic Load Balancers_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConnectionSettings

LBCookieStickinessPolicy

All content copied from https://docs.aws.amazon.com/.
