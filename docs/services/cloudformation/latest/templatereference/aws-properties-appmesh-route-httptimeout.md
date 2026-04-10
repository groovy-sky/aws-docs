This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::Route HttpTimeout

An object that represents types of timeouts.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Idle" : Duration,
  "PerRequest" : Duration
}

```

### YAML

```yaml

  Idle:
    Duration
  PerRequest:
    Duration

```

## Properties

`Idle`

An object that represents an idle timeout. An idle timeout bounds the amount of time that a connection may be idle. The default value is none.

_Required_: No

_Type_: [Duration](aws-properties-appmesh-route-duration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PerRequest`

An object that represents a per request timeout. The default value is 15 seconds. If you set a higher timeout, then make sure that the higher value is set for each App Mesh
resource in a conversation. For example, if a virtual node backend uses a virtual router provider to route to another virtual node, then the timeout should be greater than 15
seconds for the source and destination virtual node and the route.

_Required_: No

_Type_: [Duration](aws-properties-appmesh-route-duration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HttpRouteMatch

MatchRange

All content copied from https://docs.aws.amazon.com/.
