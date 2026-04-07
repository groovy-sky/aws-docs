This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::PublicKey PublicKeyConfig

Configuration information about a public key that you can use with [signed URLs and signed cookies](../../../amazoncloudfront/latest/developerguide/privatecontent.md), or with [field-level encryption](../../../amazoncloudfront/latest/developerguide/field-level-encryption.md).

CloudFront supports signed URLs and signed cookies with RSA 2048 or ECDSA 256 key signatures. Field-level encryption is only compatible with RSA 2048 key signatures.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CallerReference" : String,
  "Comment" : String,
  "EncodedKey" : String,
  "Name" : String
}

```

### YAML

```yaml

  CallerReference: String
  Comment: String
  EncodedKey: String
  Name: String

```

## Properties

`CallerReference`

A string included in the request to help make sure that the request can't be
replayed.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Comment`

A comment to describe the public key. The comment cannot be longer than 128
characters.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EncodedKey`

The public key that you can use with [signed URLs and signed cookies](../../../amazoncloudfront/latest/developerguide/privatecontent.md), or with [field-level encryption](../../../amazoncloudfront/latest/developerguide/field-level-encryption.md).

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

A name to help identify the public key.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::CloudFront::PublicKey

AWS::CloudFront::RealtimeLogConfig
