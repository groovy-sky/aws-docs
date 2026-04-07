This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeDeploy::DeploymentConfig

The `AWS::CodeDeploy::DeploymentConfig` resource creates a set of deployment
rules, deployment success conditions, and deployment failure conditions that AWS CodeDeploy uses during a deployment. The deployment configuration specifies the number
or percentage of instances that must remain available at any time during a deployment.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CodeDeploy::DeploymentConfig",
  "Properties" : {
      "ComputePlatform" : String,
      "DeploymentConfigName" : String,
      "MinimumHealthyHosts" : MinimumHealthyHosts,
      "TrafficRoutingConfig" : TrafficRoutingConfig,
      "ZonalConfig" : ZonalConfig
    }
}

```

### YAML

```yaml

Type: AWS::CodeDeploy::DeploymentConfig
Properties:
  ComputePlatform: String
  DeploymentConfigName: String
  MinimumHealthyHosts:
    MinimumHealthyHosts
  TrafficRoutingConfig:
    TrafficRoutingConfig
  ZonalConfig:
    ZonalConfig

```

## Properties

`ComputePlatform`

The destination platform type for the deployment ( `Lambda`,
`Server`, or `ECS`).

_Required_: No

_Type_: String

_Allowed values_: `Server | Lambda | ECS | Kubernetes`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DeploymentConfigName`

A name for the deployment configuration. If you don't specify a name, CloudFormation generates a unique physical ID and uses that ID for the deployment configuration name. For
more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).

###### Important

If you specify a name, you cannot perform updates that require replacement of this
resource. You can perform updates that require no or some interruption. If you must replace
the resource, specify a new name.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MinimumHealthyHosts`

The minimum number of healthy instances that should be available at any time during the
deployment. There are two parameters expected in the input: type and value.

The type parameter takes either of the following values:

- HOST\_COUNT: The value parameter represents the minimum number of healthy instances as
an absolute value.

- FLEET\_PERCENT: The value parameter represents the minimum number of healthy instances
as a percentage of the total number of instances in the deployment. If you specify
FLEET\_PERCENT, at the start of the deployment, AWS CodeDeploy converts the
percentage to the equivalent number of instance and rounds up fractional instances.

The value parameter takes an integer.

For example, to set a minimum of 95% healthy instance, specify a type of FLEET\_PERCENT and
a value of 95.

For more information about instance health, see [CodeDeploy Instance\
Health](https://docs.aws.amazon.com/codedeploy/latest/userguide/instances-health.html) in the AWS CodeDeploy User Guide.

_Required_: No

_Type_: [MinimumHealthyHosts](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-codedeploy-deploymentconfig-minimumhealthyhosts.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TrafficRoutingConfig`

The configuration that specifies how the deployment traffic is routed.

_Required_: No

_Type_: [TrafficRoutingConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-codedeploy-deploymentconfig-trafficroutingconfig.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ZonalConfig`

Configure the `ZonalConfig` object if you want AWS CodeDeploy to
deploy your application to one [Availability Zone](../../../ec2/latest/userguide/using-regions-availability-zones.md#concepts-availability-zones) at a time, within an AWS Region.

For more information about the zonal configuration feature, see [zonal configuration](https://docs.aws.amazon.com/codedeploy/latest/userguide/deployment-configurations-create.html#zonal-config) in the _CodeDeploy User_
_Guide_.

_Required_: No

_Type_: [ZonalConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-codedeploy-deploymentconfig-zonalconfig.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of an `AWS::CodeDeploy::DeploymentConfig` resource
to the intrinsic `Ref` function, the function returns the deployment configuration
name, such as `mydeploymentconfig-a123d0d1`.

For more information about using the `Ref` function, see [Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref.html).

## Examples

### Specifying minimum healthy hosts

The following example requires at least 75% of the fleet to be healthy. For example,
if you had a fleet of four instances, the deployment proceeds one instance at a
time.

#### JSON

```json

"TwentyFivePercentAtATime" : {
    "Type" : "AWS::CodeDeploy::DeploymentConfig",
    "Properties" : {
        "MinimumHealthyHosts" : {
            "Type" : "FLEET_PERCENT",
            "Value" : "75"
        }
    }
}
```

#### YAML

```yaml

TwentyFivePercentAtATime:
  Type: AWS::CodeDeploy::DeploymentConfig
  Properties:
    MinimumHealthyHosts:
      Type: "FLEET_PERCENT"
      Value: 75
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

MinimumHealthyHosts
