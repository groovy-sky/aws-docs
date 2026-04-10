This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::Route HttpRoute

An object that represents an HTTP or HTTP/2 route type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Action" : HttpRouteAction,
  "Match" : HttpRouteMatch,
  "RetryPolicy" : HttpRetryPolicy,
  "Timeout" : HttpTimeout
}

```

### YAML

```yaml

  Action:
    HttpRouteAction
  Match:
    HttpRouteMatch
  RetryPolicy:
    HttpRetryPolicy
  Timeout:
    HttpTimeout

```

## Properties

`Action`

An object that represents the action to take if a match is determined.

_Required_: Yes

_Type_: [HttpRouteAction](aws-properties-appmesh-route-httprouteaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Match`

An object that represents the criteria for determining a request match.

_Required_: Yes

_Type_: [HttpRouteMatch](aws-properties-appmesh-route-httproutematch.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RetryPolicy`

An object that represents a retry policy.

_Required_: No

_Type_: [HttpRetryPolicy](aws-properties-appmesh-route-httpretrypolicy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Timeout`

An object that represents types of timeouts.

_Required_: No

_Type_: [HttpTimeout](aws-properties-appmesh-route-httptimeout.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HttpRetryPolicy

HttpRouteAction

All content copied from https://docs.aws.amazon.com/.
