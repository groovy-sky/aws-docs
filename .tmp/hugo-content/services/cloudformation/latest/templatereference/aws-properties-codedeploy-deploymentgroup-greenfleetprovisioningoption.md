This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeDeploy::DeploymentGroup GreenFleetProvisioningOption

Information about the instances that belong to the replacement environment in a
blue/green deployment.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Action" : String
}

```

### YAML

```yaml

  Action: String

```

## Properties

`Action`

The method used to add instances to a replacement environment.

- `DISCOVER_EXISTING`: Use instances that already exist or will be
created manually.

- `COPY_AUTO_SCALING_GROUP`: Use settings from a specified Auto Scaling group to define and create instances in a new Auto Scaling
group.

_Required_: No

_Type_: String

_Allowed values_: `DISCOVER_EXISTING | COPY_AUTO_SCALING_GROUP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GitHubLocation

LoadBalancerInfo

All content copied from https://docs.aws.amazon.com/.
