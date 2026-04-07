# ResourceRequirement

The type and amount of a resource to assign to a container. The supported resource
types are GPUs and Elastic Inference accelerators. For more information, see [Working with\
GPUs on Amazon ECS](../../../../services/amazonecs/latest/developerguide/ecs-gpu.md) or [Working with Amazon Elastic\
Inference on Amazon ECS](../../../../services/amazonecs/latest/developerguide/ecs-inference.md) in the _Amazon Elastic Container Service_
_Developer Guide_

## Contents

**type**

The type of resource to assign to a container.

Type: String

Valid Values: `GPU | InferenceAccelerator`

Required: Yes

**value**

The value for the specified resource type.

When the type is `GPU`, the value is the number of physical
`GPUs` the Amazon ECS container agent reserves for the container. The
number of GPUs that's reserved for all containers in a task can't exceed the number of
available GPUs on the container instance that the task is launched on.

When the type is `InferenceAccelerator`, the `value` matches the
`deviceName` for an [InferenceAccelerator](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_InferenceAccelerator.html) specified in a task definition.

Type: String

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ecs-2014-11-13/ResourceRequirement)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ecs-2014-11-13/ResourceRequirement)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ecs-2014-11-13/ResourceRequirement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Resource

Rollback
