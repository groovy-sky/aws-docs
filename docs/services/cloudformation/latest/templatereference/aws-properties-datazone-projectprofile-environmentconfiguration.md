This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::ProjectProfile EnvironmentConfiguration

The configuration of an environment.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AwsAccount" : AwsAccount,
  "AwsRegion" : Region,
  "ConfigurationParameters" : EnvironmentConfigurationParametersDetails,
  "DeploymentMode" : String,
  "DeploymentOrder" : Number,
  "Description" : String,
  "EnvironmentBlueprintId" : String,
  "EnvironmentConfigurationId" : String,
  "Name" : String
}

```

### YAML

```yaml

  AwsAccount:
    AwsAccount
  AwsRegion:
    Region
  ConfigurationParameters:
    EnvironmentConfigurationParametersDetails
  DeploymentMode: String
  DeploymentOrder: Number
  Description: String
  EnvironmentBlueprintId: String
  EnvironmentConfigurationId: String
  Name: String

```

## Properties

`AwsAccount`

The AWS account of the environment.

_Required_: No

_Type_: [AwsAccount](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-datazone-projectprofile-awsaccount.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AwsRegion`

The AWS Region of the environment.

_Required_: Yes

_Type_: [Region](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-datazone-projectprofile-region.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConfigurationParameters`

The configuration parameters of the environment.

_Required_: No

_Type_: [EnvironmentConfigurationParametersDetails](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-datazone-projectprofile-environmentconfigurationparametersdetails.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeploymentMode`

The deployment mode of the environment.

_Required_: No

_Type_: String

_Allowed values_: `ON_CREATE | ON_DEMAND`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeploymentOrder`

The deployment order of the environment.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `16`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The environment description.

_Required_: No

_Type_: String

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnvironmentBlueprintId`

The environment blueprint ID.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]{1,36}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnvironmentConfigurationId`

The configuration ID with which the environment is created.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]{1,36}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The environment name.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w -]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AwsAccount

EnvironmentConfigurationParameter
