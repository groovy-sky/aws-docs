This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::ExpressGatewayService ExpressGatewayContainer

Defines the configuration for the primary container in an Express service. This container
receives traffic from the Application Load Balancer and runs your application code.

The container configuration includes the container image, port mapping, logging
settings, environment variables, and secrets. The container image is the only required
parameter, with sensible defaults provided for other settings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AwsLogsConfiguration" : ExpressGatewayServiceAwsLogsConfiguration,
  "Command" : [ String, ... ],
  "ContainerPort" : Integer,
  "Environment" : [ KeyValuePair, ... ],
  "Image" : String,
  "RepositoryCredentials" : ExpressGatewayRepositoryCredentials,
  "Secrets" : [ Secret, ... ]
}

```

### YAML

```yaml

  AwsLogsConfiguration:
    ExpressGatewayServiceAwsLogsConfiguration
  Command:
    - String
  ContainerPort: Integer
  Environment:
    - KeyValuePair
  Image: String
  RepositoryCredentials:
    ExpressGatewayRepositoryCredentials
  Secrets:
    - Secret

```

## Properties

`AwsLogsConfiguration`

The log configuration for the container.

_Required_: No

_Type_: [ExpressGatewayServiceAwsLogsConfiguration](aws-properties-ecs-expressgatewayservice-expressgatewayserviceawslogsconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Command`

The command that is passed to the container.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ContainerPort`

The port number on the container that receives traffic from the load balancer. Default
is 80.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Environment`

The environment variables to pass to the container.

_Required_: No

_Type_: Array of [KeyValuePair](aws-properties-ecs-expressgatewayservice-keyvaluepair.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Image`

The image used to start a container. This string is passed directly to the Docker
daemon. Images in the Docker Hub registry are available by default. Other repositories are
specified with either `repository-url/image:tag` or
`repository-url/image@digest`.

For Express services, the image typically contains a web application that listens on the
specified container port. The image can be stored in Amazon ECR, Docker Hub, or any other
container registry accessible to your execution role.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RepositoryCredentials`

The configuration for repository credentials for private registry authentication.

_Required_: No

_Type_: [ExpressGatewayRepositoryCredentials](aws-properties-ecs-expressgatewayservice-expressgatewayrepositorycredentials.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Secrets`

The secrets to pass to the container.

_Required_: No

_Type_: Array of [Secret](aws-properties-ecs-expressgatewayservice-secret.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ECSManagedResourceArns

ExpressGatewayRepositoryCredentials

All content copied from https://docs.aws.amazon.com/.
