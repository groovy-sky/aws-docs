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

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ecs-2014-11-13/InferenceAccelerator)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ecs-2014-11-13/InferenceAccelerator)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ecs-2014-11-13/InferenceAccelerator)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

HostVolumeProperties

InferenceAcceleratorOverride
