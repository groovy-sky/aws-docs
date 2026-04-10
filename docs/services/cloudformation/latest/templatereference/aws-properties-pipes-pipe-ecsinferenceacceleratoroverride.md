This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pipes::Pipe EcsInferenceAcceleratorOverride

Details on an Elastic Inference accelerator task override. This parameter is used to
override the Elastic Inference accelerator specified in the task definition. For more
information, see [Working with Amazon Elastic\
Inference on Amazon ECS](../../../amazonecs/latest/userguide/ecs-inference.md) in the _Amazon Elastic Container Service_
_Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DeviceName" : String,
  "DeviceType" : String
}

```

### YAML

```yaml

  DeviceName: String
  DeviceType: String

```

## Properties

`DeviceName`

The Elastic Inference accelerator device name to override for the task. This parameter
must match a `deviceName` specified in the task definition.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeviceType`

The Elastic Inference accelerator type to use.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EcsEphemeralStorage

EcsResourceRequirement

All content copied from https://docs.aws.amazon.com/.
