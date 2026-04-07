This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel CdiInputSpecification

The input specification for this channel. It specifies the key characteristics of CDI inputs for this channel,
when those characteristics are different from other inputs.

This entity is at the top level in the channel.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Resolution" : String
}

```

### YAML

```yaml

  Resolution: String

```

## Properties

`Resolution`

Maximum CDI input resolution

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CaptionSelectorSettings

ChannelEngineVersionRequest
