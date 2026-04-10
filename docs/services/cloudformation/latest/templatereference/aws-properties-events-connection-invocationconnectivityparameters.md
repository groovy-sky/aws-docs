This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Events::Connection InvocationConnectivityParameters

For connections to private APIs, the parameters to use for invoking the API.

For more information, see [Connecting to private APIs](../../../eventbridge/latest/userguide/connection-private.md) in the _Amazon EventBridge User Guide_.

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

ConnectivityParameters

OAuthParameters

All content copied from https://docs.aws.amazon.com/.
