This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pipes::Pipe EcsResourceRequirement

The type and amount of a resource to assign to a container. The supported resource types
are GPUs and Elastic Inference accelerators. For more information, see [Working with\
GPUs on Amazon ECS](../../../amazonecs/latest/developerguide/ecs-gpu.md) or [Working with Amazon Elastic\
Inference on Amazon ECS](../../../amazonecs/latest/developerguide/ecs-inference.md) in the _Amazon Elastic Container Service_
_Developer Guide_

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : String,
  "Value" : String
}

```

### YAML

```yaml

  Type: String
  Value: String

```

## Properties

`Type`

The type of resource to assign to a container. The supported values are `GPU`
or `InferenceAccelerator`.

_Required_: Yes

_Type_: String

_Allowed values_: `GPU | InferenceAccelerator`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value for the specified resource type.

If the `GPU` type is used, the value is the number of physical
`GPUs` the Amazon ECS container agent reserves for the container. The
number of GPUs that's reserved for all containers in a task can't exceed the number of
available GPUs on the container instance that the task is launched on.

If the `InferenceAccelerator` type is used, the `value` matches
the `deviceName` for an InferenceAccelerator specified in a task
definition.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EcsInferenceAcceleratorOverride

EcsTaskOverride

All content copied from https://docs.aws.amazon.com/.
