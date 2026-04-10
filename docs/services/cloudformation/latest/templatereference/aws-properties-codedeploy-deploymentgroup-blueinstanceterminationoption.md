This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeDeploy::DeploymentGroup BlueInstanceTerminationOption

Information about whether instances in the original environment are terminated when a
blue/green deployment is successful. `BlueInstanceTerminationOption` does not
apply to Lambda deployments.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Action" : String,
  "TerminationWaitTimeInMinutes" : Integer
}

```

### YAML

```yaml

  Action: String
  TerminationWaitTimeInMinutes: Integer

```

## Properties

`Action`

The action to take on instances in the original environment after a successful
blue/green deployment.

- `TERMINATE`: Instances are terminated after a specified wait
time.

- `KEEP_ALIVE`: Instances are left running after they are
deregistered from the load balancer and removed from the deployment
group.

_Required_: No

_Type_: String

_Allowed values_: `TERMINATE | KEEP_ALIVE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TerminationWaitTimeInMinutes`

For an Amazon EC2 deployment, the number of minutes to wait after a successful
blue/green deployment before terminating instances from the original environment.

For an Amazon ECS deployment, the number of minutes before deleting the
original (blue) task set. During an Amazon ECS deployment, CodeDeploy shifts
traffic from the original (blue) task set to a replacement (green) task set.

The maximum setting is 2880 minutes (2 days).

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BlueGreenDeploymentConfiguration

Deployment

All content copied from https://docs.aws.amazon.com/.
