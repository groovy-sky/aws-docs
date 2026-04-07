This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IVS::PlaybackKeyPair

The `AWS::IVS::PlaybackKeyPair` resource specifies an Amazon IVS playback key pair. Amazon IVS
uses a public playback key to validate playback tokens that have been signed with the
corresponding private key. For more information, see [Setting Up Private Channels](https://docs.aws.amazon.com/ivs/latest/LowLatencyUserGuide/private-channels.html)
in the _Amazon IVS Low-Latency Streaming User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IVS::PlaybackKeyPair",
  "Properties" : {
      "Name" : String,
      "PublicKeyMaterial" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::IVS::PlaybackKeyPair
Properties:
  Name: String
  PublicKeyMaterial: String
  Tags:
    - Tag

```

## Properties

`Name`

Playback-key-pair name. The value does not need to be unique.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9-_]*$`

_Minimum_: `0`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PublicKeyMaterial`

The public portion of a customer-generated key pair. Note that this field is required to create the AWS::IVS::PlaybackKeyPair resource.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-playbackkeypair-tag.html).

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ivs-playbackkeypair-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the playback key pair ARN. For example:

`{ "Ref": "myPlaybackKeyPair" }`

For the Amazon IVS playback key pair
`myPlaybackKeyPair`, `Ref` returns the playback key pair
ARN.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

Key-pair ARN. For example:
`arn:aws:ivs:us-west-2:693991300569:playback-key/f99cde61-c2b0-4df3-8941-ca7d38acca1a`

`Fingerprint`

Key-pair identifier. For example:
`98:0d:1a:a0:19:96:1e:ea:0a:0a:2c:9a:42:19:2b:e7`

## Examples

### Playback Key Pair Template Examples

The following examples specify an Amazon IVS
playback key pair.

#### JSON

```json

 {
     "AWSTemplateFormatVersion": "2010-09-09",
     "Resources": {
         "PlaybackKeyPair": {
             "Type": "AWS::IVS::PlaybackKeyPair",
             "Properties": {
                 "PublicKeyMaterial": "-----BEGIN PUBLIC KEY-----\nMHYwEAYHKoZIzj0CAQYFK4EEACIDYgAEwOR43ETwEoWif1i14aL8GtDMNkT/kBQm\nh4sas9P//bjCU988rmQQXVBfftKT9xngg+W6hzOEpeUlCRlAtz6b6U79naYYRaSk\nK/UhYGWkXlbJlc9zn13imYWgVGe/BMFp\n-----END PUBLIC KEY-----\n",
                 "Name": "MyPlaybackKeyPair",
                 "Tags": [
                     {
                         "Key": "MyKey",
                         "Value": "MyValue"
                     }
                 ]
             }
         }
     },
     "Outputs": {
         "PlaybackKeyPairArn": {
             "Value": {"Ref": "PlaybackKeyPair"}
         },
         "PlaybackKeyPairFingerprint": {
             "Value": {
                 "Fn::GetAtt": [
                     "PlaybackKeyPair",
                     "Fingerprint"
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
  PlaybackKeyPair:
    Type: AWS::IVS::PlaybackKeyPair
    Properties:
      PublicKeyMaterial: |
        -----BEGIN PUBLIC KEY-----
        MHYwEAYHKoZIzj0CAQYFK4EEACIDYgAEwOR43ETwEoWif1i14aL8GtDMNkT/kBQm
        h4sas9P//bjCU988rmQQXVBfftKT9xngg+W6hzOEpeUlCRlAtz6b6U79naYYRaSk
        K/UhYGWkXlbJlc9zn13imYWgVGe/BMFp
        -----END PUBLIC KEY-----
      Name: MyPlaybackKeyPair
      Tags:
        - Key: MyKey
          Value: MyValue
 Outputs:
  PlaybackKeyPairArn:
    Value: !Ref PlaybackKeyPair
  PlaybackKeyPairFingerprint:
    Value: !GetAtt PlaybackKeyPair.Fingerprint

```

## See also

- [Setting Up Private\
Channels](https://docs.aws.amazon.com/ivs/latest/LowLatencyUserGuide/private-channels.html)

- [PlaybackKeyPair](https://docs.aws.amazon.com/ivs/latest/LowLatencyAPIReference/API_PlaybackKeyPair.html) data type

- [ImportPlaybackKeyPair](https://docs.aws.amazon.com/ivs/latest/LowLatencyAPIReference/API_ImportPlaybackKeyPair.html) API endpoint

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Tag
