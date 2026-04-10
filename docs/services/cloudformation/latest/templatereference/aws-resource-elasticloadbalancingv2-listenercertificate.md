This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElasticLoadBalancingV2::ListenerCertificate

Specifies an SSL server certificate to add to the certificate list for an HTTPS or TLS
listener.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ElasticLoadBalancingV2::ListenerCertificate",
  "Properties" : {
      "Certificates" : [ Certificate, ... ],
      "ListenerArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::ElasticLoadBalancingV2::ListenerCertificate
Properties:
  Certificates:
    - Certificate
  ListenerArn: String

```

## Properties

`Certificates`

The certificate. You can specify one certificate per resource.

_Required_: Yes

_Type_: Array of [Certificate](aws-properties-elasticloadbalancingv2-listenercertificate-certificate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ListenerArn`

The Amazon Resource Name (ARN) of the listener.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Examples

The following example creates a certificate list.

#### YAML

```yaml

myCertificateList:
  Type: 'AWS::ElasticLoadBalancingV2::ListenerCertificate'
  Properties:
    ListenerArn: !Ref myTLSListener
    Certificates:
      - CertificateArn: "arn:aws:acm:us-west-2:123456789012:certificate/11111111-1234-46f1-abcd-11117EXAMPLE"
      - CertificateArn: "arn:aws:acm:us-west-2:123456789012:certificate/22222222-4567-57f2-cdef-22227EXAMPLE"
      - CertificateArn: "arn:aws:acm:us-west-2:123456789012:certificate/33333333-5678-68f3-a123-33337EXAMPLE"
```

#### JSON

```json

{
    "myCertificateList": {
        "Type": "AWS::ElasticLoadBalancingV2::ListenerCertificate",
        "Properties": {
            "ListenerArn": {
                "Ref": "myTLSListener"
            },
            "Certificates": [
                {
                    "CertificateArn": "arn:aws:acm:us-west-2:123456789012:certificate/11111111-1234-46f1-abcd-11117EXAMPLE"
                },
                {
                    "CertificateArn": "arn:aws:acm:us-west-2:123456789012:certificate/22222222-4567-57f2-cdef-22227EXAMPLE"
                },
                {
                    "CertificateArn": "arn:aws:acm:us-west-2:123456789012:certificate/33333333-5678-68f3-a123-33337EXAMPLE"
                }
            ]
        }
    }
}
```

## See also

- [AddListenerCertificates](../../../../reference/elasticloadbalancing/latest/apireference/api-addlistenercertificates.md) in the _Elastic Load Balancing API_
_Reference (version 2015-12-01)_

- [SSL Certificates](../../../elasticloadbalancing/latest/application/https-listener-certificates.md) in the _User Guide for Application Load_
_Balancers_

- [Server certificates](../../../elasticloadbalancing/latest/network/tls-listener-certificates.md) in the _User Guide for Network Load_
_Balancers_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TargetGroupTuple

Certificate

All content copied from https://docs.aws.amazon.com/.
