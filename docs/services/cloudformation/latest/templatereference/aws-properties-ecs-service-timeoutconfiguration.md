This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::Service TimeoutConfiguration

An object that represents the timeout configurations for Service Connect.

###### Note

If `idleTimeout` is set to a time that is less than
`perRequestTimeout`, the connection will close when the
`idleTimeout` is reached and not the
`perRequestTimeout`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IdleTimeoutSeconds" : Integer,
  "PerRequestTimeoutSeconds" : Integer
}

```

### YAML

```yaml

  IdleTimeoutSeconds: Integer
  PerRequestTimeoutSeconds: Integer

```

## Properties

`IdleTimeoutSeconds`

The amount of time in seconds a connection will stay active while idle. A value of
`0` can be set to disable `idleTimeout`.

The `idleTimeout` default for
`HTTP`/ `HTTP2`/ `GRPC` is 5 minutes.

The `idleTimeout` default for `TCP` is 1 hour.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `2147483647`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PerRequestTimeoutSeconds`

The amount of time waiting for the upstream to respond with a complete response per
request. A value of `0` can be set to disable `perRequestTimeout`.
`perRequestTimeout` can only be set if Service Connect
`appProtocol` isn't `TCP`. Only `idleTimeout` is
allowed for `TCP` `appProtocol`.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `2147483647`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Associate an Application Load Balancer with a service](../userguide/aws-resource-ecs-service.md#aws-resource-ecs-service--examples--Associate_an_Application_Load_Balancer_with_a_service)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

VpcLatticeConfiguration

All content copied from https://docs.aws.amazon.com/.
