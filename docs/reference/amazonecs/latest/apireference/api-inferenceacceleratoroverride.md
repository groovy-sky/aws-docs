# InferenceAcceleratorOverride

Details on an Elastic Inference accelerator task override. This parameter is used to
override the Elastic Inference accelerator specified in the task definition. For more
information, see [Working with Amazon Elastic Inference on Amazon ECS](../../../../services/amazonecs/latest/developerguide/ecs-inference.md) in the _Amazon_
_Elastic Container Service Developer Guide_.

## Contents

**deviceName**

The Elastic Inference accelerator device name to override for the task. This parameter
must match a `deviceName` specified in the task definition.

Type: String

Required: No

**deviceType**

The Elastic Inference accelerator type to use.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ecs-2014-11-13/inferenceacceleratoroverride.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecs-2014-11-13/inferenceacceleratoroverride.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecs-2014-11-13/inferenceacceleratoroverride.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

InferenceAccelerator

InfrastructureOptimization
