This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::StreamingDistribution TrustedSigners

A list of AWS accounts whose public keys CloudFront can use to verify the signatures of
signed URLs and signed cookies.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AwsAccountNumbers" : [ String, ... ],
  "Enabled" : Boolean
}

```

### YAML

```yaml

  AwsAccountNumbers:
    - String
  Enabled: Boolean

```

## Properties

`AwsAccountNumbers`

An AWS account number that contains active CloudFront key pairs that CloudFront can use to
verify the signatures of signed URLs and signed cookies. If the AWS account that owns
the key pairs is the same account that owns the CloudFront distribution, the value of this
field is `self`.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enabled`

This field is `true` if any of the AWS accounts in the list are configured as
trusted signers. If not, this field is `false`.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [TrustedSigners](../../../../reference/cloudfront/latest/apireference/api-trustedsigners.md) in the _Amazon CloudFront API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::CloudFront::TrustStore

All content copied from https://docs.aws.amazon.com/.
