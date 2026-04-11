This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::RouterOutput MediaLiveInputRouterOutputConfiguration

Configuration settings for connecting a router output to a MediaLive input.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DestinationTransitEncryption" : MediaLiveTransitEncryption,
  "MediaLiveInputArn" : String,
  "MediaLivePipelineId" : String
}

```

### YAML

```yaml

  DestinationTransitEncryption:
    MediaLiveTransitEncryption
  MediaLiveInputArn: String
  MediaLivePipelineId: String

```

## Properties

`DestinationTransitEncryption`

The encryption configuration for the MediaLive input when connected to this router output.

_Required_: Yes

_Type_: [MediaLiveTransitEncryption](aws-properties-mediaconnect-routeroutput-medialivetransitencryption.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MediaLiveInputArn`

The ARN of the MediaLive input to connect to this router output.

_Required_: No

_Type_: String

_Pattern_: `^arn:(aws[a-zA-Z-]*):medialive:[a-z0-9-]+:[0-9]{12}:input:[a-zA-Z0-9]+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MediaLivePipelineId`

The index of the MediaLive pipeline to connect to this router output.

_Required_: No

_Type_: String

_Allowed values_: `PIPELINE_0 | PIPELINE_1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MediaConnectFlowRouterOutputConfiguration

MediaLiveTransitEncryption

All content copied from https://docs.aws.amazon.com/.
