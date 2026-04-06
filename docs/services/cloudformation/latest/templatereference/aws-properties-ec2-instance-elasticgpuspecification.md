This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::Instance ElasticGpuSpecification

###### Note

Amazon Elastic Graphics reached end of life on January 8, 2024.

Specifies the type of Elastic GPU. An Elastic GPU is a GPU resource that you can attach
to your Amazon EC2 instance to accelerate the graphics performance of your applications.

`ElasticGpuSpecification` is a property of the [AWS::EC2::Instance](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html) resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : String
}

```

### YAML

```yaml

  Type: String

```

## Properties

`Type`

The type of Elastic Graphics accelerator.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Ebs

ElasticInferenceAccelerator
