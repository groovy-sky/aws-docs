This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Events::Connection InvocationConnectivityParameters

For connections to private APIs, the parameters to use for invoking the API.

For more information, see [Connecting to private APIs](https://docs.aws.amazon.com/eventbridge/latest/userguide/connection-private.html) in the _Amazon EventBridge User Guide_.

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

_Type_: [ResourceParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-events-connection-resourceparameters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ConnectivityParameters

OAuthParameters
