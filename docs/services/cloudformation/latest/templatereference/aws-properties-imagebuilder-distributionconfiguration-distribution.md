This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ImageBuilder::DistributionConfiguration Distribution

The distribution configuration distribution defines the settings for a specific Region
in the Distribution Configuration. You must specify whether the distribution is for an AMI
or a container image. To do so, include exactly one of the following data types for your
distribution:

- amiDistributionConfiguration

- containerDistributionConfiguration

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AmiDistributionConfiguration" : AmiDistributionConfiguration,
  "ContainerDistributionConfiguration" : ContainerDistributionConfiguration,
  "FastLaunchConfigurations" : [ FastLaunchConfiguration, ... ],
  "LaunchTemplateConfigurations" : [ LaunchTemplateConfiguration, ... ],
  "LicenseConfigurationArns" : [ String, ... ],
  "Region" : String,
  "SsmParameterConfigurations" : [ SsmParameterConfiguration, ... ]
}

```

### YAML

```yaml

  AmiDistributionConfiguration:
    AmiDistributionConfiguration
  ContainerDistributionConfiguration:
    ContainerDistributionConfiguration
  FastLaunchConfigurations:
    - FastLaunchConfiguration
  LaunchTemplateConfigurations:
    - LaunchTemplateConfiguration
  LicenseConfigurationArns:
    - String
  Region: String
  SsmParameterConfigurations:
    - SsmParameterConfiguration

```

## Properties

`AmiDistributionConfiguration`

The specific AMI settings, such as launch permissions and AMI tags. For details,
see example schema below.

_Required_: No

_Type_: [AmiDistributionConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-imagebuilder-distributionconfiguration-amidistributionconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ContainerDistributionConfiguration`

Container distribution settings for encryption, licensing, and sharing
in a specific Region. For details, see example schema below.

_Required_: No

_Type_: [ContainerDistributionConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-imagebuilder-distributionconfiguration-containerdistributionconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FastLaunchConfigurations`

The Windows faster-launching configurations to use for AMI distribution.

_Required_: No

_Type_: Array of [FastLaunchConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-imagebuilder-distributionconfiguration-fastlaunchconfiguration.html)

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LaunchTemplateConfigurations`

A group of launchTemplateConfiguration settings that apply to image distribution for
specified accounts.

_Required_: No

_Type_: Array of [LaunchTemplateConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-imagebuilder-distributionconfiguration-launchtemplateconfiguration.html)

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LicenseConfigurationArns`

The License Manager Configuration to associate with the AMI in the specified Region.
For more information, see the [LicenseConfiguration API](https://docs.aws.amazon.com/license-manager/latest/APIReference/API_LicenseConfiguration.html).

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Region`

The target Region for the Distribution Configuration. For example,
`eu-west-1`.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SsmParameterConfigurations`

Contains settings to update AWS Systems Manager (SSM) Parameter Store Parameters with
output AMI IDs from the build by target Region.

_Required_: No

_Type_: Array of [SsmParameterConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-imagebuilder-distributionconfiguration-ssmparameterconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

- [Example 1: AmiDistributionConfiguration schema with launch permissions](#aws-properties-imagebuilder-distributionconfiguration-distribution--examples--Example_1:_AmiDistributionConfiguration_schema_with_launch_permissions)

- [Example 2: Create a distribution configuration resource for a copied AMI](#aws-properties-imagebuilder-distributionconfiguration-distribution--examples--Example_2:_Create_a_distribution_configuration_resource_for_a_copied_AMI)

- [Example 3: ContainerDistributionConfiguration schema](#aws-properties-imagebuilder-distributionconfiguration-distribution--examples--Example_3:_ContainerDistributionConfiguration_schema)

### Example 1: AmiDistributionConfiguration schema with launch permissions

The following example shows the schema for the AmiDistributionConfiguration
property in both YAML and JSON format.

###### Note

To make an AMI public, set the launch permission authorized accounts to
`all`. See the examples for making an AMI public at [EC2 ModifyImageAttribute](../../../../reference/awsec2/latest/apireference/api-modifyimageattribute.md).

#### YAML

```yaml

Distributions:
  - Region: 'us-west-2'
    AmiDistributionConfiguration:
      Name: 'AmiCopyConfiguration - {{ imagebuilder:buildDate }}'
      Description: 'Share an AMI in the distribution Region by granting launch permissions to specified AWS organizations, OUs, user groups, and accounts.'
      AmiTags:
        AmiTagKey: 'AmiTagValue'
      LaunchPermissionConfiguration:
        OrganizationArns:
          - 'arn:aws:organizations::123456789012:organization/o-myorganization123'
        OrganizationalUnitArns:
          - 'arn:aws:organizations::123456789012:ou/o-123example/ou-1234-myorganizationalunit'
        UserGroups:
          - 'GroupName1'
          - 'GroupName2'
        UserIds:
          - '123456789012'
          - '345678901234'
```

#### JSON

```json

"{
    "Distributions": [{
        AmiDistributionConfiguration": {
        "Name": "AmiCopyConfiguration - {{ imagebuilder:buildDate }}",
        "Description": "Share an AMI in the distribution Region by granting launch permissions to specified user groups and accounts.",
        "AmiTags": {
            "AmiTagKey": "AmiTagValue"
        },
        "LaunchPermissionConfiguration": {
            "OrganizationArns": ["arn:aws:organizations::123456789012:organization/o-myorganization123"],
            "OrganizationalUnitArns": ["arn:aws:organizations::123456789012:ou/o-123example/ou-1234-myorganizationalunit"],
            "UserGroups": [
                "GroupName1",
                "GroupName2"
            ],
            "UserIds": [
                "123456789012",
                "345678901234"
            ]
        }
    }]
}
```

### Example 2: Create a distribution configuration resource for a copied AMI

The following example shows the schema for the AmiDistributionConfiguration
property in both YAML and JSON.

#### YAML

```yaml

Distributions:
  - Region: 'us-west-2'
    AmiDistributionConfiguration:
      Name: AmiCopyConfiguration - {{ imagebuilder:buildDate }}
      Description: 'Distribute a copy of the AMI to specific accounts in the destination Region.'
      AmiTags:
        AmiTagKey: 'AmiTagValue'
      TargetAccountIds:
        - '123456789012'
        - '345678901234'
```

#### JSON

```json

{
    "Distributions": [{
        "AmiDistributionConfiguration": {
            "Name": "AmiCopyConfiguration - {{ imagebuilder:buildDate }}",
            "Description": "Distribute a copy of the AMI to specific accounts in the destination Region.",
            "AmiTags": {
                "AmiTagKey": "AmiTagValue"
            },
            "TargetAccountIds": ["123456789012", "345678901234"]
        }
    }]
}
```

### Example 3: ContainerDistributionConfiguration schema

The following example shows the schema for the ContainerDistributionConfiguration
property in both YAML and JSON format.

#### YAML

```yaml

Distributions:
  - Region: 'us-west-2'
    ContainerDistributionConfiguration:
      Description: 'test distribution cfn template'
      TargetRepository:
        Service: ECR
        RepositoryName: 'cfn-test'
      ContainerTags:
        - 'Tag1'
        - 'Tag2'
```

#### JSON

```json

{
    "Distributions": [{
        "ContainerDistributionConfiguration": {
            "Description": "test distribution cfn template",
            "TargetRepository": {
                "Service": "ECR",
                "RepositoryName": "cfn-test"
            },
            "ContainerTags": ["Tag1", "Tag2"]
        }
    }]
}
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ContainerDistributionConfiguration

FastLaunchConfiguration
