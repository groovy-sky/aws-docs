This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppSync::ChannelNamespace

The `AWS::AppSync::ChannelNamespace` resource creates a channel namespace
associated with an `Api`. The `ChannelNamespace` contains the
definitions for code handlers for the `Api`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AppSync::ChannelNamespace",
  "Properties" : {
      "ApiId" : String,
      "CodeHandlers" : String,
      "CodeS3Location" : String,
      "HandlerConfigs" : HandlerConfigs,
      "Name" : String,
      "PublishAuthModes" : [ AuthMode, ... ],
      "SubscribeAuthModes" : [ AuthMode, ... ],
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::AppSync::ChannelNamespace
Properties:
  ApiId: String
  CodeHandlers: String
  CodeS3Location: String
  HandlerConfigs:
    HandlerConfigs
  Name: String
  PublishAuthModes:
    - AuthMode
  SubscribeAuthModes:
    - AuthMode
  Tags:
    - Tag

```

## Properties

`ApiId`

The `Api` ID.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CodeHandlers`

The event handler functions that run custom business logic to process published events
and subscribe requests.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32768`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CodeS3Location`

The Amazon S3 endpoint where the code is located.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HandlerConfigs`

The configuration for the `OnPublish` and `OnSubscribe`
handlers.

_Required_: No

_Type_: [HandlerConfigs](aws-properties-appsync-channelnamespace-handlerconfigs.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the channel namespace. This name must be unique within the
`Api`.

_Required_: Yes

_Type_: String

_Pattern_: `([A-Za-z0-9](?:[A-Za-z0-9\-]{0,48}[A-Za-z0-9])?)`

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PublishAuthModes`

The authorization mode to use for publishing messages on the channel namespace. This
configuration overrides the default `Api` authorization configuration.

_Required_: No

_Type_: Array of [AuthMode](aws-properties-appsync-channelnamespace-authmode.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubscribeAuthModes`

The authorization mode to use for subscribing to messages on the channel namespace.
This configuration overrides the default `Api` authorization
configuration.

_Required_: No

_Type_: Array of [AuthMode](aws-properties-appsync-channelnamespace-authmode.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A set of tags (key-value pairs) for this channel namespace.

_Required_: No

_Type_: Array of [Tag](aws-properties-appsync-channelnamespace-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns The Amazon Resource Name (ARN) of the channel namespace.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ChannelNamespaceArn`

The Amazon Resource Name (ARN) of the channel namespace.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::AppSync::ApiKey

AuthMode

All content copied from https://docs.aws.amazon.com/.
