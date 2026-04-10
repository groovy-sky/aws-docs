This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lightsail::Container PublicEndpoint

`PublicEndpoint` is a property of the [ContainerServiceDeployment](../userguide/aws-properties-lightsail-container-containerservicedeployment.md) property. It describes describes the settings of the
public endpoint of a container on a container service.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ContainerName" : String,
  "ContainerPort" : Integer,
  "HealthCheckConfig" : HealthCheckConfig
}

```

### YAML

```yaml

  ContainerName: String
  ContainerPort: Integer
  HealthCheckConfig:
    HealthCheckConfig

```

## Properties

`ContainerName`

The name of the container entry of the deployment that the endpoint configuration
applies to.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ContainerPort`

The port of the specified container to which traffic is forwarded to.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HealthCheckConfig`

An object that describes the health check configuration of the container.

_Required_: No

_Type_: [HealthCheckConfig](aws-properties-lightsail-container-healthcheckconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PublicDomainName

Tag

All content copied from https://docs.aws.amazon.com/.
