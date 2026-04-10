This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::Route GrpcRoute

An object that represents a gRPC route type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Action" : GrpcRouteAction,
  "Match" : GrpcRouteMatch,
  "RetryPolicy" : GrpcRetryPolicy,
  "Timeout" : GrpcTimeout
}

```

### YAML

```yaml

  Action:
    GrpcRouteAction
  Match:
    GrpcRouteMatch
  RetryPolicy:
    GrpcRetryPolicy
  Timeout:
    GrpcTimeout

```

## Properties

`Action`

An object that represents the action to take if a match is determined.

_Required_: Yes

_Type_: [GrpcRouteAction](aws-properties-appmesh-route-grpcrouteaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Match`

An object that represents the criteria for determining a request match.

_Required_: Yes

_Type_: [GrpcRouteMatch](aws-properties-appmesh-route-grpcroutematch.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RetryPolicy`

An object that represents a retry policy.

_Required_: No

_Type_: [GrpcRetryPolicy](aws-properties-appmesh-route-grpcretrypolicy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Timeout`

An object that represents types of timeouts.

_Required_: No

_Type_: [GrpcTimeout](aws-properties-appmesh-route-grpctimeout.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GrpcRetryPolicy

GrpcRouteAction

All content copied from https://docs.aws.amazon.com/.
