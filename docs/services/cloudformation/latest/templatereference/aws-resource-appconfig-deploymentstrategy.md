This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppConfig::DeploymentStrategy

The `AWS::AppConfig::DeploymentStrategy` resource creates an AWS AppConfig deployment strategy. A deployment strategy defines important criteria for
rolling out your configuration to the designated targets. A deployment strategy includes: the
overall duration required, a percentage of targets to receive the deployment during each
interval, an algorithm that defines how percentage grows, and bake time.

AWS AppConfig requires that you create resources and deploy a configuration in the
following order:

1. Create an application

2. Create an environment

3. Create a configuration profile

4. Choose a pre-defined deployment strategy or create your own

5. Deploy the configuration

For more information, see [AWS AppConfig](../../../appconfig/latest/userguide/what-is-appconfig.md) in the _AWS AppConfig User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AppConfig::DeploymentStrategy",
  "Properties" : {
      "DeploymentDurationInMinutes" : Number,
      "Description" : String,
      "FinalBakeTimeInMinutes" : Number,
      "GrowthFactor" : Number,
      "GrowthType" : String,
      "Name" : String,
      "ReplicateTo" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::AppConfig::DeploymentStrategy
Properties:
  DeploymentDurationInMinutes: Number
  Description: String
  FinalBakeTimeInMinutes: Number
  GrowthFactor: Number
  GrowthType: String
  Name: String
  ReplicateTo: String
  Tags:
    - Tag

```

## Properties

`DeploymentDurationInMinutes`

Total amount of time for a deployment to last.

_Required_: Yes

_Type_: Number

_Minimum_: `0`

_Maximum_: `1440`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description of the deployment strategy.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FinalBakeTimeInMinutes`

Specifies the amount of time AWS AppConfig monitors for Amazon CloudWatch alarms after the
configuration has been deployed to 100% of its targets, before considering the deployment
to be complete. If an alarm is triggered during this time, AWS AppConfig rolls back
the deployment. You must configure permissions for AWS AppConfig to roll back based
on CloudWatch alarms. For more information, see [Configuring permissions for rollback based on Amazon CloudWatch alarms](https://docs.aws.amazon.com/appconfig/latest/userguide/getting-started-with-appconfig-cloudwatch-alarms-permissions.html) in the
_AWS AppConfig User Guide_.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `1440`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GrowthFactor`

The percentage of targets to receive a deployed configuration during each
interval.

_Required_: Yes

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GrowthType`

The algorithm used to define how percentage grows over time. AWS AppConfig
supports the following growth types:

**Linear**: For this type, AWS AppConfig processes
the deployment by dividing the total number of targets by the value specified for
`Step percentage`. For example, a linear deployment that uses a `Step
            percentage` of 10 deploys the configuration to 10 percent of the hosts. After
those deployments are complete, the system deploys the configuration to the next 10
percent. This continues until 100% of the targets have successfully received the
configuration.

**Exponential**: For this type, AWS AppConfig
processes the deployment exponentially using the following formula: `G*(2^N)`.
In this formula, `G` is the growth factor specified by the user and
`N` is the number of steps until the configuration is deployed to all
targets. For example, if you specify a growth factor of 2, then the system rolls out the
configuration as follows:

`2*(2^0)`

`2*(2^1)`

`2*(2^2)`

Expressed numerically, the deployment rolls out as follows: 2% of the targets, 4% of the
targets, 8% of the targets, and continues until the configuration has been deployed to all
targets.

_Required_: No

_Type_: String

_Allowed values_: `EXPONENTIAL | LINEAR`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

A name for the deployment strategy.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ReplicateTo`

Save the deployment strategy to a Systems Manager (SSM) document.

_Required_: Yes

_Type_: String

_Allowed values_: `NONE | SSM_DOCUMENT`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Assigns metadata to an AWS AppConfig resource. Tags help organize and categorize
your AWS AppConfig resources. Each tag consists of a key and an optional value, both
of which you define. You can specify a maximum of 50 tags for a resource.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appconfig-deploymentstrategy-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the deployment strategy ID.

### Fn::GetAtt

`Id`

The deployment strategy ID.

## Examples

### AWS AppConfig deployment strategy example

The following example creates an AWS AppConfig deployment strategy called
MyTestDeploymentStrategy. A deployment strategy defines how a configuration is
deployed.

#### JSON

```json

Resources": {
    "BasicDeploymentStrategy": {
      "Type": "AWS::AppConfig::DeploymentStrategy",
      "Properties": {
        "Name": "MyTestDeploymentStrategy",
        "Description": "A sample test deployment strategy.",
        "DeploymentDurationInMinutes": 3,
        "FinalBakeTimeInMinutes": 4,
        "GrowthFactor": 10,
        "GrowthType": "LINEAR",
        "ReplicateTo": "NONE",
        "Tags": [
          {
            "Key": "Env",
            "Value": "test"
          }
        ]
      }
    }
  }
}

```

#### YAML

```yaml

Resources:
  BasicDeploymentStrategy:
    Type: AWS::AppConfig::DeploymentStrategy
    Properties:
      Name: "MyTestDeploymentStrategy"
      Description: "A sample test deployment strategy."
      DeploymentDurationInMinutes: 3
      FinalBakeTimeInMinutes: 4
      GrowthFactor: 10
      GrowthType: LINEAR
      ReplicateTo: NONE
      Tags:
        - Key: Env
          Value: test

```

## See also

- [AWS AppConfig](../../../appconfig/latest/userguide/what-is-appconfig.md)

- [Creating a deployment strategy](https://docs.aws.amazon.com/systems-manager/latest/userguide/appconfig-creating-deployment-strategy.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Tag
