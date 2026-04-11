This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Events::Endpoint Secondary

The secondary Region that processes events when failover is triggered or replication is
enabled.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Route" : String
}

```

### YAML

```yaml

  Route: String

```

## Properties

`Route`

Defines the secondary Region.

_Required_: Yes

_Type_: String

_Pattern_: `^[\-a-z0-9]+$`

_Minimum_: `9`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RoutingConfig

AWS::Events::EventBus

All content copied from https://docs.aws.amazon.com/.
