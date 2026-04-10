This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElasticLoadBalancingV2::TrustStoreRevocation

Adds the specified revocation contents to the specified trust store. You must specify `TrustStoreArn`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ElasticLoadBalancingV2::TrustStoreRevocation",
  "Properties" : {
      "RevocationContents" : [ RevocationContent, ... ],
      "TrustStoreArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::ElasticLoadBalancingV2::TrustStoreRevocation
Properties:
  RevocationContents:
    - RevocationContent
  TrustStoreArn: String

```

## Properties

`RevocationContents`

The revocation file to add.

_Required_: No

_Type_: Array of [RevocationContent](aws-properties-elasticloadbalancingv2-truststorerevocation-revocationcontent.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TrustStoreArn`

The Amazon Resource Name (ARN) of the trust store.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the trust store.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`RevocationId`

The revocation ID of the revocation file.

`TrustStoreRevocations`

Information about the revocation file in the trust store. For more information, see [TrustStoreRevocation](aws-properties-elasticloadbalancingv2-truststorerevocation-truststorerevocation.md).

## Examples

The following example creates a revocation list.

#### YAML

```yaml

myRevocationList:
  Type: 'AWS::ElasticLoadBalancingV2::TrustStoreRevocation'
  Properties:
    TrustStoreArn: !Ref myTrustStore
    RevocationContents:
        - RevocationType: CRL
          S3Bucket: amzn-s3-demo-bucket
          S3Key: crl/revoked-list.crl
```

#### JSON

```json

{
    "myRevocationList": {
        "Type": "AWS::ElasticLoadBalancingV2::TrustStoreRevocation",
        "Properties": {
            "TrustStoreArn": {
                "Ref": "myTrustStore"
            },
            "RevocationContents": [
                {
                    "RevocationType": "CRL",
                    "S3Bucket": "amzn-s3-demo-bucket",
                    "S3Key": "crl/revoked-list.crl"
                }
            ]
        }
    }
}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

RevocationContent

All content copied from https://docs.aws.amazon.com/.
