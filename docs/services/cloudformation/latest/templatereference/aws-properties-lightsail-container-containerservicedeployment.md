This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lightsail::Container ContainerServiceDeployment

`ContainerServiceDeployment` is a property of the [AWS::Lightsail::Container](../userguide/aws-resource-lightsail-container.md) resource. It describes a container deployment
configuration of a container service.

A deployment specifies the settings, such as the ports and launch command, of containers
that are deployed to your container service.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Containers" : [ Container, ... ],
  "PublicEndpoint" : PublicEndpoint
}

```

### YAML

```yaml

  Containers:
    - Container
  PublicEndpoint:
    PublicEndpoint

```

## Properties

`Containers`

An object that describes the configuration for the containers of the deployment.

_Required_: No

_Type_: Array of [Container](aws-properties-lightsail-container-container.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PublicEndpoint`

An object that describes the endpoint of the deployment.

_Required_: No

_Type_: [PublicEndpoint](aws-properties-lightsail-container-publicendpoint.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Container

EcrImagePullerRole

All content copied from https://docs.aws.amazon.com/.
