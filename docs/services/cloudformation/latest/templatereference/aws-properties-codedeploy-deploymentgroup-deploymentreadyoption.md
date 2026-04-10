This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeDeploy::DeploymentGroup DeploymentReadyOption

Information about how traffic is rerouted to instances in a replacement environment in
a blue/green deployment.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ActionOnTimeout" : String,
  "WaitTimeInMinutes" : Integer
}

```

### YAML

```yaml

  ActionOnTimeout: String
  WaitTimeInMinutes: Integer

```

## Properties

`ActionOnTimeout`

Information about when to reroute traffic from an original environment to a replacement
environment in a blue/green deployment.

- CONTINUE\_DEPLOYMENT: Register new instances with the load balancer immediately after
the new application revision is installed on the instances in the replacement
environment.

- STOP\_DEPLOYMENT: Do not register new instances with a load balancer unless traffic
rerouting is started using [ContinueDeployment](../../../../reference/codedeploy/latest/apireference/api-continuedeployment.md). If traffic rerouting is not started before the end of the specified wait period,
the deployment status is changed to Stopped.

_Required_: No

_Type_: String

_Allowed values_: `CONTINUE_DEPLOYMENT | STOP_DEPLOYMENT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WaitTimeInMinutes`

The number of minutes to wait before the status of a blue/green deployment is changed
to Stopped if rerouting is not started manually. Applies only to the
`STOP_DEPLOYMENT` option for `actionOnTimeout`.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Deployment

DeploymentStyle

All content copied from https://docs.aws.amazon.com/.
