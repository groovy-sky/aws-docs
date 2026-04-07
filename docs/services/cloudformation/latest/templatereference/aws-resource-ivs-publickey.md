This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IVS::PublicKey

The `AWS::IVS::PublicKey` resource specifies an Amazon IVS public key used to sign stage participant tokens.
For more information, see [Distribute Participant Tokens](https://docs.aws.amazon.com/ivs/latest/RealTimeUserGuide/getting-started-distribute-tokens.html)
in the _Amazon IVS Real-Time Streaming User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IVS::PublicKey",
  "Properties" : {
      "Name" : String,
      "PublicKeyMaterial" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::IVS::PublicKey
Properties:
  Name: String
  PublicKeyMaterial: String
  Tags:
    - Tag

```

## Properties

`Name`

Public key name. The value does not need to be unique.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9-_]*$`

_Minimum_: `0`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PublicKeyMaterial`

The public portion of a customer-generated key pair. Note that this field is required to create the AWS::IVS::PublicKey resource.

_Required_: No

_Type_: String

_Pattern_: `-----BEGIN PUBLIC KEY-----\r?\n([a-zA-Z0-9+/=\r\n]+)\r?\n-----END PUBLIC KEY-----(\r?\n)?`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An array of key-value pairs to apply to this resource.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ivs-publickey-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the public key ARN. For example:

`{ "Ref": "myPublicKey" }`

For the Amazon IVS public key
`myPublicKey`, `Ref` returns the
public key ARN.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The public key ARN. For example:
`arn:aws:ivs:us-west-2:123456789012:public-key/abcdABCDefgh`

`Fingerprint`

The public key identifier. For example:
`98:0d:1a:a0:19:96:1e:ea:0a:0a:2c:9a:42:19:2b:e7`

## Examples

### PublicKey Template Examples

The following examples specify an Amazon IVS public key.

#### JSON

```json

{
     "AWSTemplateFormatVersion": "2010-09-09",
     "Resources": {
         "PublicKey": {
             "Type": "AWS::IVS::PublicKey",
             "Properties": {
                 "PublicKeyMaterial": "-----BEGIN PUBLIC KEY-----\nMHYwEAYHKoZIzj0CAQYFK4EEACIDYgAEwOR43ETwEoWif1i14aL8GtDMNkT/kBQm\nh4sas9P//bjCU988rmQQXVBfftKT9xngg+W6hzOEpeUlCRlAtz6b6U79naYYRaSk\nK/UhYGWkXlbJlc9zn13imYWgVGe/BMFp\n-----END PUBLIC KEY-----\n",
                 "Name": "MyPublicKey",
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
  PublicKey:
    Type: AWS::IVS::PublicKey
    Properties:
      PublicKeyMaterial: |
        -----BEGIN PUBLIC KEY-----
        MHYwEAYHKoZIzj0CAQYFK4EEACIDYgAEwOR43ETwEoWif1i14aL8GtDMNkT/kBQm
        h4sas9P//bjCU988rmQQXVBfftKT9xngg+W6hzOEpeUlCRlAtz6b6U79naYYRaSk
        K/UhYGWkXlbJlc9zn13imYWgVGe/BMFp
        -----END PUBLIC KEY-----
      Name: MyPublicKey
      Tags:
        - Key: MyKey
          Value: MyValue
```

## See also

- [Distribute Participant Tokens](https://docs.aws.amazon.com/ivs/latest/RealTimeUserGuide/getting-started-distribute-tokens.html)

- [PublicKey](https://docs.aws.amazon.com/ivs/latest/RealTimeAPIReference/API_PublicKey.html) data
type

- [ImportPublicKey](https://docs.aws.amazon.com/ivs/latest/RealTimeAPIReference/API_ImportPublicKey.html) API endpoint

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Tag
