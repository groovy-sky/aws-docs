This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaPackage::Channel

Creates a channel to receive content.

After it's created, a channel provides static input URLs. These URLs remain the same throughout the lifetime of the channel, regardless of any failures or upgrades that might
occur. Use these URLs to configure the outputs of your upstream encoder.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::MediaPackage::Channel",
  "Properties" : {
      "Description" : String,
      "EgressAccessLogs" : LogConfiguration,
      "HlsIngest" : HlsIngest,
      "Id" : String,
      "IngressAccessLogs" : LogConfiguration,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::MediaPackage::Channel
Properties:
  Description: String
  EgressAccessLogs:
    LogConfiguration
  HlsIngest:
    HlsIngest
  Id: String
  IngressAccessLogs:
    LogConfiguration
  Tags:
    - Tag

```

## Properties

`Description`

Any descriptive information that you want to add to the channel for future identification purposes.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EgressAccessLogs`

Configures egress access logs.

_Required_: No

_Type_: [LogConfiguration](aws-properties-mediapackage-channel-logconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HlsIngest`

The input URL where the source stream should be sent.

_Required_: No

_Type_: [HlsIngest](aws-properties-mediapackage-channel-hlsingest.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Id`

Unique identifier that you assign to the channel.

_Required_: Yes

_Type_: String

_Pattern_: `\A[0-9a-zA-Z-_]+\Z`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IngressAccessLogs`

Configures ingress access logs.

_Required_: No

_Type_: [LogConfiguration](aws-properties-mediapackage-channel-logconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags to assign to the channel.

_Required_: No

_Type_: Array of [Tag](aws-properties-mediapackage-channel-tag.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the channel.

For example: `{ "Ref": "myChannel" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The channel's unique system-generated resource name, based on the AWS record.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

HlsIngest

All content copied from https://docs.aws.amazon.com/.
