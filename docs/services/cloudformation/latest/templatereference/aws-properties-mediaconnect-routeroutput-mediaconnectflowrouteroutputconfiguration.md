This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::RouterOutput MediaConnectFlowRouterOutputConfiguration

Configuration settings for connecting a router output to a MediaConnect flow source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DestinationTransitEncryption" : FlowTransitEncryption,
  "FlowArn" : String,
  "FlowSourceArn" : String
}

```

### YAML

```yaml

  DestinationTransitEncryption:
    FlowTransitEncryption
  FlowArn: String
  FlowSourceArn: String

```

## Properties

`DestinationTransitEncryption`

The encryption configuration for the flow destination when connected to this router output.

_Required_: Yes

_Type_: [FlowTransitEncryption](aws-properties-mediaconnect-routeroutput-flowtransitencryption.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FlowArn`

The ARN of the flow to connect to this router output.

_Required_: No

_Type_: String

_Pattern_: `^arn:(aws[a-zA-Z-]*):mediaconnect:[a-z0-9-]+:[0-9]{12}:flow:[a-zA-Z0-9-]+:[a-zA-Z0-9_-]+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FlowSourceArn`

The ARN of the flow source to connect to this router output.

_Required_: No

_Type_: String

_Pattern_: `^arn:(aws[a-zA-Z-]*):mediaconnect:[a-z0-9-]+:[0-9]{12}:source:[a-zA-Z0-9-]+:[a-zA-Z0-9_-]+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MaintenanceConfiguration

MediaLiveInputRouterOutputConfiguration

All content copied from https://docs.aws.amazon.com/.
