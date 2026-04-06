This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::Instance LicenseSpecification

Specifies the license configuration to use.

`LicenseSpecification` is a property of the [AWS::EC2::Instance](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html) resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LicenseConfigurationArn" : String
}

```

### YAML

```yaml

  LicenseConfigurationArn: String

```

## Properties

`LicenseConfigurationArn`

The Amazon Resource Name (ARN) of the license configuration.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

LaunchTemplateSpecification

MetadataOptions
