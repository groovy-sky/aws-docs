# InferenceAccelerator

Details on an Elastic Inference accelerator. For more information, see [Working with Amazon Elastic Inference on Amazon ECS](../../../../services/amazonecs/latest/developerguide/ecs-inference.md) in the _Amazon_
_Elastic Container Service Developer Guide_.

## Contents

**deviceName**

The Elastic Inference accelerator device name. The `deviceName` must also
be referenced in a container definition as a [ResourceRequirement](api-resourcerequirement.md).

Type: String

Required: Yes

**deviceType**

The Elastic Inference accelerator type to use.

Type: String

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ecs-2014-11-13/inferenceaccelerator.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecs-2014-11-13/inferenceaccelerator.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecs-2014-11-13/inferenceaccelerator.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

HostVolumeProperties

InferenceAcceleratorOverride
