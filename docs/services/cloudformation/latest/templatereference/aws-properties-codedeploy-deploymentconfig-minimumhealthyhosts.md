This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeDeploy::DeploymentConfig MinimumHealthyHosts

`MinimumHealthyHosts` is a property of the [DeploymentConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentconfig.html) resource that defines how many instances must remain healthy
during an AWS CodeDeploy deployment.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : String,
  "Value" : Integer
}

```

### YAML

```yaml

  Type: String
  Value: Integer

```

## Properties

`Type`

The minimum healthy instance type:

- HOST\_COUNT: The minimum number of healthy instance as an absolute value.

- FLEET\_PERCENT: The minimum number of healthy instance as a percentage of the total
number of instance in the deployment.

In an example of nine instance, if a HOST\_COUNT of six is specified, deploy to up to three
instances at a time. The deployment is successful if six or more instances are deployed to
successfully. Otherwise, the deployment fails. If a FLEET\_PERCENT of 40 is specified, deploy
to up to five instance at a time. The deployment is successful if four or more instance are
deployed to successfully. Otherwise, the deployment fails.

###### Note

In a call to `GetDeploymentConfig`, CodeDeployDefault.OneAtATime returns a
minimum healthy instance type of MOST\_CONCURRENCY and a value of 1. This means a deployment
to only one instance at a time. (You cannot set the type to MOST\_CONCURRENCY, only to
HOST\_COUNT or FLEET\_PERCENT.) In addition, with CodeDeployDefault.OneAtATime, AWS CodeDeploy attempts to ensure that all instances but one are kept in a healthy state
during the deployment. Although this allows one instance at a time to be taken offline for a
new deployment, it also means that if the deployment to the last instance fails, the overall
deployment is still successful.

For more information, see [AWS CodeDeploy Instance\
Health](https://docs.aws.amazon.com/codedeploy/latest/userguide/instances-health.html) in the _AWS CodeDeploy User Guide_.

_Required_: Yes

_Type_: String

_Allowed values_: `HOST_COUNT | FLEET_PERCENT`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Value`

The minimum healthy instance value.

_Required_: Yes

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::CodeDeploy::DeploymentConfig

MinimumHealthyHostsPerZone
