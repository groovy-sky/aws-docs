This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::LaunchTemplate CreditSpecification

Specifies the credit option for CPU usage of a T2, T3, or T3a instance.

`CreditSpecification` is a property of [AWS::EC2::LaunchTemplate LaunchTemplateData](../userguide/aws-properties-ec2-launchtemplate-launchtemplatedata.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CpuCredits" : String
}

```

### YAML

```yaml

  CpuCredits: String

```

## Properties

`CpuCredits`

The credit option for CPU usage of a T instance.

Valid values: `standard` \| `unlimited`

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [CreditSpecificationRequest](../../../../reference/awsec2/latest/apireference/api-creditspecificationrequest.md) in the _Amazon EC2 API_
_Reference_

- [Create a launch template using advanced settings](../../../autoscaling/ec2/userguide/advanced-settings-for-your-launch-template.md) in the _Amazon EC2 Auto Scaling User Guide_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CpuOptions

Ebs

All content copied from https://docs.aws.amazon.com/.
