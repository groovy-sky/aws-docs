This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::PrimaryTaskSet

Modifies which task set in a service is the primary task set. Any parameters that are
updated on the primary task set in a service will transition to the service. This is
used when a service uses the `EXTERNAL` deployment controller type. For more
information, see [Amazon ECS Deployment\
Types](../../../amazonecs/latest/developerguide/deployment-types.md) in the _Amazon Elastic Container Service Developer_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ECS::PrimaryTaskSet",
  "Properties" : {
      "Cluster" : String,
      "Service" : String,
      "TaskSetId" : String
    }
}

```

### YAML

```yaml

Type: AWS::ECS::PrimaryTaskSet
Properties:
  Cluster: String
  Service: String
  TaskSetId: String

```

## Properties

`Cluster`

The short name or full Amazon Resource Name (ARN) of the cluster that hosts the
service that the task set exists in.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Service`

The short name or full Amazon Resource Name (ARN) of the service that the task set
exists in.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TaskSetId`

The short name or full Amazon Resource Name (ARN) of the task set to set as the
primary task set in the deployment.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::ECS::Service

All content copied from https://docs.aws.amazon.com/.
