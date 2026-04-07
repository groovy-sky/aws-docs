This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IVS::EncoderConfiguration

The `AWS::IVS::EncoderConfiguration` resource specifies an Amazon IVS
encoder configuration. An encoder configuration describes a stream’s video configuration. For more information,
see [Streaming Configuration](https://docs.aws.amazon.com/ivs/latest/LowLatencyUserGuide/streaming-config.html)
in the _Amazon IVS Low-Latency Streaming User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IVS::EncoderConfiguration",
  "Properties" : {
      "Name" : String,
      "Tags" : [ Tag, ... ],
      "Video" : Video
    }
}

```

### YAML

```yaml

Type: AWS::IVS::EncoderConfiguration
Properties:
  Name: String
  Tags:
    - Tag
  Video:
    Video

```

## Properties

`Name`

Encoder cnfiguration name.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9-_]*$`

_Minimum_: `0`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-encoderconfiguration-tag.html).

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ivs-encoderconfiguration-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Video`

Video configuration. Default: video resolution 1280x720, bitrate 2500 kbps, 30 fps. See the [Video](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-encoderconfiguration-video.html) property type for more information.

_Required_: No

_Type_: [Video](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ivs-encoderconfiguration-video.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the encoder-configuration ARN. For example:

`{ "Ref": "myEncoderConfiguration" }`

For the Amazon IVS encoder configuration
`"myEncoderConfiguration"`, `Ref` returns the
encoder-configuration ARN.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The encoder-configuration ARN. For example:
`arn:aws:ivs:us-west-2:123456789012:encoder-configuration/abcdABCDefgh`

## Examples

### EncoderConfiguration Template Examples

The following examples specify an Amazon IVS encoder configuration.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Resources": {
        "EncoderConfiguration": {
            "Type": "AWS::IVS::EncoderConfiguration",
            "Properties": {
                "Name": "myEncoderConfiguration",
                "Video": {
                    "Bitrate": 1500000,
                    "Framerate": 1.5,
                    "Height": 100,
                    "Width": 100
                },
                "Tags": [
                    {
                        "Key": "MyKey",
                        "Value": "MyValue"
                    }
                ]
            }
        }
    }
}

```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  EncoderConfiguration:
    Type: AWS::IVS::EncoderConfiguration
    Properties:
      Video:
        Bitrate: 1500000
        Framerate: 1.5
        Height: 100
        Width: 100
      Name: myEncoderConfiguration
      Tags:
        - Key: myKey
          Value: myValue

```

## See also

- [Getting\
Started with IVS Real-Time Streaming](https://docs.aws.amazon.com/ivs/latest/RealTimeUserGuide/getting-started.html)

- [Server-Side Composition (Real-Time Streaming)](https://docs.aws.amazon.com/ivs/latest/RealTimeUserGuide/server-side-composition.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Tag
