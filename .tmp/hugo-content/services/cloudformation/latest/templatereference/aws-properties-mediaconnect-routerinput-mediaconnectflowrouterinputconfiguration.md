This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::RouterInput MediaConnectFlowRouterInputConfiguration

Configuration settings for connecting a router input to a flow output.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FlowArn" : String,
  "FlowOutputArn" : String,
  "SourceTransitDecryption" : FlowTransitEncryption
}

```

### YAML

```yaml

  FlowArn: String
  FlowOutputArn: String
  SourceTransitDecryption:
    FlowTransitEncryption

```

## Properties

`FlowArn`

The ARN of the flow to connect to.

_Required_: No

_Type_: String

_Pattern_: `^arn:(aws[a-zA-Z-]*):mediaconnect:[a-z0-9-]+:[0-9]{12}:flow:[a-zA-Z0-9-]+:[a-zA-Z0-9_-]+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FlowOutputArn`

The ARN of the flow output to connect to this router input.

_Required_: No

_Type_: String

_Pattern_: `^arn:(aws[a-zA-Z-]*):mediaconnect:[a-z0-9-]+:[0-9]{12}:output:[a-zA-Z0-9-]+:[a-zA-Z0-9_-]+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceTransitDecryption`

The decryption configuration for the flow source when connected to this router input.

_Required_: Yes

_Type_: [FlowTransitEncryption](aws-properties-mediaconnect-routerinput-flowtransitencryption.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MaintenanceConfiguration

MergeRouterInputConfiguration

All content copied from https://docs.aws.amazon.com/.
