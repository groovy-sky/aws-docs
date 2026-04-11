This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::Route TcpRouteAction

An object that represents the action to take if a match is determined.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "WeightedTargets" : [ WeightedTarget, ... ]
}

```

### YAML

```yaml

  WeightedTargets:
    - WeightedTarget

```

## Properties

`WeightedTargets`

An object that represents the targets that traffic is routed to when a request matches the route.

_Required_: Yes

_Type_: Array of [WeightedTarget](aws-properties-appmesh-route-weightedtarget.md)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TcpRoute

TcpRouteMatch

All content copied from https://docs.aws.amazon.com/.
