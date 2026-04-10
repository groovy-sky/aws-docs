This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::Instance ElasticInferenceAccelerator

###### Note

Amazon Elastic Inference is no longer available.

Specifies the Elastic Inference Accelerator for the instance.

`ElasticInferenceAccelerator` is a property of the [AWS::EC2::Instance](../userguide/aws-properties-ec2-instance.md) resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Count" : Integer,
  "Type" : String
}

```

### YAML

```yaml

  Count: Integer
  Type: String

```

## Properties

`Count`

The number of elastic inference accelerators to attach to the instance.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Type`

The type of elastic inference accelerator. The possible values are `eia1.medium`, `eia1.large`, `eia1.xlarge`, `eia2.medium`, `eia2.large`, and `eia2.xlarge`.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ElasticGpuSpecification

EnaSrdSpecification

All content copied from https://docs.aws.amazon.com/.
