This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Events::Connection ConnectivityParameters

If you specify a private OAuth endpoint, the parameters for EventBridge to use when authenticating against the endpoint.

For more information, see [Authorization methods for connections](../../../eventbridge/latest/userguide/eb-target-connection-auth.md) in the _Amazon EventBridge User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ResourceParameters" : ResourceParameters
}

```

### YAML

```yaml

  ResourceParameters:
    ResourceParameters

```

## Properties

`ResourceParameters`

The parameters for EventBridge to use when invoking the resource endpoint.

_Required_: Yes

_Type_: [ResourceParameters](aws-properties-events-connection-resourceparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConnectionHttpParameters

InvocationConnectivityParameters

All content copied from https://docs.aws.amazon.com/.
