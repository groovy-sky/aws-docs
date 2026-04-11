This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ServiceDiscovery::Service HealthCheckCustomConfig

A complex type that contains information about an optional custom health check. A custom
health check, which requires that you use a third-party health checker to evaluate the health of
your resources, is useful in the following circumstances:

- You can't use a health check that's defined by `HealthCheckConfig` because the
resource isn't available over the internet. For example, you can use a custom health check when
the instance is in an Amazon VPC. (To check the health of resources in a VPC, the health
checker must also be in the VPC.)

- You want to use a third-party health checker regardless of where your resources are
located.

###### Important

If you specify a health check configuration, you can specify either
`HealthCheckCustomConfig` or `HealthCheckConfig` but not both.

To change the status of a custom health check, submit an
`UpdateInstanceCustomHealthStatus` request. AWS Cloud Map doesn't monitor the status of the
resource, it just keeps a record of the status specified in the most recent
`UpdateInstanceCustomHealthStatus` request.

Here's how custom health checks work:

1. You create a service.

2. You register an instance.

3. You configure a third-party health checker to monitor the resource that's associated with
    the new instance.

###### Note

AWS Cloud Map doesn't check the health of the resource directly.

4. The third-party health-checker determines that the resource is unhealthy and notifies your
    application.

5. Your application submits an `UpdateInstanceCustomHealthStatus` request.

6. AWS Cloud Map waits for 30 seconds.

7. If another `UpdateInstanceCustomHealthStatus` request doesn't arrive during
    that time to change the status back to healthy, AWS Cloud Map stops routing traffic to the
    resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FailureThreshold" : Number
}

```

### YAML

```yaml

  FailureThreshold: Number

```

## Properties

`FailureThreshold`

###### Important

This parameter is no longer supported and is always set to 1. AWS Cloud Map waits for
approximately 30 seconds after receiving an `UpdateInstanceCustomHealthStatus`
request before changing the status of the service instance.

The number of 30-second intervals that you want AWS Cloud Map to wait after receiving an
`UpdateInstanceCustomHealthStatus` request before it changes the health status of a
service instance.

Sending a second or subsequent `UpdateInstanceCustomHealthStatus` request with
the same value before 30 seconds has passed doesn't accelerate the change. AWS Cloud Map still waits
`30` seconds after the first request to make the change.

_Required_: No

_Type_: Number

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## See also

- [Return values](../userguide/aws-resource-servicediscovery-service.md#aws-resource-servicediscovery-service-return-values) in the topic [AWS::ServiceDiscovery::Service](../userguide/aws-resource-servicediscovery-service.md)

- [HealthCheckCustomConfig](../../../cloud-map/latest/api/api-healthcheckcustomconfig.md) in the _AWS Cloud Map API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HealthCheckConfig

Tag

All content copied from https://docs.aws.amazon.com/.
