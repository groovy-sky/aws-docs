This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lightsail::Container Container

`Container` is a property of the [ContainerServiceDeployment](../userguide/aws-properties-lightsail-container-containerservicedeployment.md) property. It describes the settings of a container
that will be launched, or that is launched, to an Amazon Lightsail container
service.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Command" : [ String, ... ],
  "ContainerName" : String,
  "Environment" : [ EnvironmentVariable, ... ],
  "Image" : String,
  "Ports" : [ PortInfo, ... ]
}

```

### YAML

```yaml

  Command:
    - String
  ContainerName: String
  Environment:
    - EnvironmentVariable
  Image: String
  Ports:
    - PortInfo

```

## Properties

`Command`

The launch command for the container.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ContainerName`

The name of the container.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Environment`

The environment variables of the container.

_Required_: No

_Type_: Array of [EnvironmentVariable](aws-properties-lightsail-container-environmentvariable.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Image`

The name of the image used for the container.

Container images that are sourced from (registered and stored on) your container service
start with a colon ( `:`). For example, if your container service name is `container-service-1`,
the container image label is `mystaticsite`, and you want to use the third version ( `3`) of the
registered container image, then you should specify `:container-service-1.mystaticsite.3`. To use the latest
version of a container image, specify `latest` instead of a version number (for example,
`:container-service-1.mystaticsite.latest`). Your container service will automatically use the highest numbered
version of the registered container image.

Container images that are sourced from a public registry like Docker Hub don’t start with a
colon. For example, `nginx:latest` or `nginx`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Ports`

An object that describes the open firewall ports and protocols of the container.

_Required_: No

_Type_: Array of [PortInfo](aws-properties-lightsail-container-portinfo.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Lightsail::Container

ContainerServiceDeployment

All content copied from https://docs.aws.amazon.com/.
