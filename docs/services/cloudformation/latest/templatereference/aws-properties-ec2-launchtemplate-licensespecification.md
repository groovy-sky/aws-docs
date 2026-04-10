This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::LaunchTemplate LicenseSpecification

Specifies a license configuration for an instance.

`LicenseSpecification` is a property of [AWS::EC2::LaunchTemplate LaunchTemplateData](../userguide/aws-properties-ec2-launchtemplate-launchtemplatedata.md).

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

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [LaunchTemplateLicenseConfigurationRequest](../../../../reference/awsec2/latest/apireference/api-launchtemplatelicenseconfigurationrequest.md) in the _Amazon EC2 API_
_Reference_

- [Create a launch template using advanced settings](../../../autoscaling/ec2/userguide/advanced-settings-for-your-launch-template.md) in the _Amazon EC2 Auto Scaling User Guide_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LaunchTemplateTagSpecification

MaintenanceOptions

All content copied from https://docs.aws.amazon.com/.
