This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::Service ForceNewDeployment

Determines whether to force a new deployment of the service. By default, deployments
aren't forced. You can use this option to start a new deployment with no service
definition changes. For example, you can update a service's tasks to use a newer Docker
image with the same image/tag combination ( `my_image:latest`) or to roll
Fargate tasks onto a newer platform version.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EnableForceNewDeployment" : Boolean,
  "ForceNewDeploymentNonce" : String
}

```

### YAML

```yaml

  EnableForceNewDeployment: Boolean
  ForceNewDeploymentNonce: String

```

## Properties

`EnableForceNewDeployment`

Determines whether to force a new deployment of the service. By default, deployments
aren't forced. You can use this option to start a new deployment with no service
definition changes. For example, you can update a service's tasks to use a newer Docker
image with the same image/tag combination ( `my_image:latest`) or to roll
Fargate tasks onto a newer platform version.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ForceNewDeploymentNonce`

When you change the ` ForceNewDeploymentNonce` value in your
template, it signals Amazon ECS to start a new deployment even though no
other service parameters have changed. The value must be a unique, time-
varying value like a timestamp, random string, or sequence number.
Use this property when you want to ensure your tasks pick up the
latest version of a Docker image that uses the same tag but has been updated in
the registry.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EBSTagSpecification

LinearConfiguration

All content copied from https://docs.aws.amazon.com/.
