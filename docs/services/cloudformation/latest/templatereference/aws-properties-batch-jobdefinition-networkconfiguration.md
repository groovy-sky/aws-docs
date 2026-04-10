This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::JobDefinition NetworkConfiguration

The network configuration for jobs that are running on Fargate resources. Jobs that are
running on Amazon EC2 resources must not specify this parameter.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AssignPublicIp" : String
}

```

### YAML

```yaml

  AssignPublicIp: String

```

## Properties

`AssignPublicIp`

Indicates whether the job has a public IP address. For a job that's running on Fargate
resources in a private subnet to send outbound traffic to the internet (for example, to pull
container images), the private subnet requires a NAT gateway be attached to route requests to the
internet. For more information, see [Amazon ECS task networking](../../../amazonecs/latest/developerguide/task-networking.md) in the
_Amazon Elastic Container Service Developer Guide_. The default value is " `DISABLED`".

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MultiNodeEcsTaskProperties

NodeProperties

All content copied from https://docs.aws.amazon.com/.
