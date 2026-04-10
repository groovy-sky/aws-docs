This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IVS::PlaybackRestrictionPolicy

The `AWS::IVS::PlaybackRestrictionPolicy` resource specifies an Amazon IVS
playback restriction policy. A playback restriction policy constrains playback by country and/or origin sites.
For more information, see [Undesired Content and Viewers](../../../ivs/latest/lowlatencyuserguide/undesired-content.md)
in the _Amazon IVS Low-Latency Streaming User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IVS::PlaybackRestrictionPolicy",
  "Properties" : {
      "AllowedCountries" : [ String, ... ],
      "AllowedOrigins" : [ String, ... ],
      "EnableStrictOriginEnforcement" : Boolean,
      "Name" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::IVS::PlaybackRestrictionPolicy
Properties:
  AllowedCountries:
    - String
  AllowedOrigins:
    - String
  EnableStrictOriginEnforcement: Boolean
  Name: String
  Tags:
    - Tag

```

## Properties

`AllowedCountries`

A list of country codes that control geoblocking restrictions. Allowed values are the officially assigned ISO 3166-1 alpha-2 codes. Default: All countries (an empty array).

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AllowedOrigins`

A list of origin sites that control CORS restriction. Allowed values are the same as valid values of the Origin header defined at [https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Origin"](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Origin)

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnableStrictOriginEnforcement`

Whether channel playback is constrained by the origin site.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

Playback-restriction-policy name.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9-_]*$`

_Minimum_: `0`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](../userguide/aws-properties-ivs-playbackrestrictionpolicy-tag.md).

_Required_: No

_Type_: Array of [Tag](aws-properties-ivs-playbackrestrictionpolicy-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the playback-restriction-policy ARN. For example:

`{ "Ref": "myPlaybackRestrictionPolicy" }`

For the Amazon IVS playback restriction policy
`"myPlaybackRestrictionPolicy"`, `Ref` returns the
playback-restriction-policy ARN.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The playback-restriction-policy ARN. For example:
`arn:aws:ivs:us-west-2:123456789012:playback-restriction-policy/abcdABCDefgh`

## Examples

### PlaybackRestrictionPolicy Template Examples

The following examples specify an Amazon IVS playback restriction policy.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Resources": {
        "PlaybackRestrictionPolicy": {
            "Type": "AWS::IVS::PlaybackRestrictionPolicy",
            "Properties": {
                "AllowedCountries" : [ "US" ],
                "AllowedOrigins" : [ "https://aws.amazon.com" ],
                "EnableStrictOriginEnforcement" : true,
                "Name": "myPlaybackRestrictionPolicy",
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
  PlaybackRestrictionPolicy:
    Type: AWS::IVS::PlaybackRestrictionPolicy
    Properties:
      AllowedCountries:
        - US
      AllowedOrigins:
        - https://aws.amazon.com
      EnableStrictOriginEnforcement: true
      Name: myPlaybackRestrictionPolicy
      Tags:
        - Key: myKey
          Value: myValue

```

## See also

- [Getting\
Started with IVS Low-Latency Streaming](../../../ivs/latest/lowlatencyuserguide/getting-started.md)

- [Undesired\
Content and Viewers](../../../ivs/latest/lowlatencyuserguide/undesired-content.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
