This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::Instance LaunchTemplateSpecification

Specifies a launch template to use when launching an Amazon EC2 instance.

You must specify the following:

- The ID or the name of the launch template, but not both.

- The version of the launch template.

For information about creating a launch template, see
[AWS::EC2::LaunchTemplate](../userguide/aws-resource-ec2-launchtemplate.md) and
[Create a launch template](../../../ec2/latest/userguide/ec2-launch-templates.md#create-launch-template)
in the _Amazon EC2 User Guide_. For example launch templates, see the
[Examples](../userguide/aws-resource-ec2-launchtemplate.md#aws-resource-ec2-launchtemplate--examples) for `AWS::EC2::LaunchTemplate`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LaunchTemplateId" : String,
  "LaunchTemplateName" : String,
  "Version" : String
}

```

### YAML

```yaml

  LaunchTemplateId: String
  LaunchTemplateName: String
  Version: String

```

## Properties

`LaunchTemplateId`

The ID of the launch template.

You must specify either the launch template ID or the
launch template name, but not both.

_Required_: Conditional

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LaunchTemplateName`

The name of the launch template.

You must specify either the launch template ID or the
launch template name, but not both.

_Required_: Conditional

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Version`

The version number of the launch template. You must specify this property.

To specify the default version of the template, use the `Fn::GetAtt`
intrinsic function to retrieve the `DefaultVersionNumber` attribute
of the launch template. To specify the latest version of the template, use
`Fn::GetAtt` to retrieve the `LatestVersionNumber` attribute.
For more information, see [AWS::EC2:LaunchTemplate return values for Fn::GetAtt](../userguide/aws-resource-ec2-launchtemplate.md#aws-resource-ec2-launchtemplate-return-values-fn--getatt).

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Examples

### Launch an instance using a launch template

This example creates a launch template and uses it to launch a new instance.

#### YAML

```yaml

Resources:
  myInstance:
    Type: 'AWS::EC2::Instance'
    Properties:
      LaunchTemplate:
        LaunchTemplateId: !Ref myLaunchTemplate
        Version: !GetAtt myLaunchTemplate.DefaultVersionNumber
  myLaunchTemplate:
    Type: 'AWS::EC2::LaunchTemplate'
    Properties:
      LaunchTemplateData:
        ImageId: ami-0a70b9d193ae8a799
        InstanceType: t2.micro
        SecurityGroupIds:
          - sg-12a4c434
```

## See also

- [LaunchTemplateSpecification](../../../../reference/awsec2/latest/apireference/api-launchtemplatespecification.md) in the _Amazon EC2 API_
_Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InstanceIpv6Address

LicenseSpecification

All content copied from https://docs.aws.amazon.com/.
