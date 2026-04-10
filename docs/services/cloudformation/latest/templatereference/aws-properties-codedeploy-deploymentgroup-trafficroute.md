This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeDeploy::DeploymentGroup TrafficRoute

Information about a listener. The listener contains the path used to route traffic
that is received from the load balancer to a target group.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ListenerArns" : [ String, ... ]
}

```

### YAML

```yaml

  ListenerArns:
    - String

```

## Properties

`ListenerArns`

The Amazon Resource Name (ARN) of one listener. The listener identifies the route
between a target group and a load balancer. This is an array of strings with a maximum
size of one.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TargetGroupPairInfo

TriggerConfig

All content copied from https://docs.aws.amazon.com/.
