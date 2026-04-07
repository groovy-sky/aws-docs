This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElasticLoadBalancingV2::Listener MutualAuthentication

The mutual authentication configuration information.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AdvertiseTrustStoreCaNames" : String,
  "IgnoreClientCertificateExpiry" : Boolean,
  "Mode" : String,
  "TrustStoreArn" : String
}

```

### YAML

```yaml

  AdvertiseTrustStoreCaNames: String
  IgnoreClientCertificateExpiry: Boolean
  Mode: String
  TrustStoreArn: String

```

## Properties

`AdvertiseTrustStoreCaNames`

Indicates whether trust store CA certificate names are advertised.

_Required_: No

_Type_: String

_Allowed values_: `on | off`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IgnoreClientCertificateExpiry`

Indicates whether expired client certificates are ignored.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Mode`

The client certificate handling method. Options are `off`,
`passthrough` or `verify`. The default value on initial resource creation is `off`.
After mutual authentication is turned on, you must explicitly set the `Mode` to `off` to turn it off;
removing the property from your template will not turn it off.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TrustStoreArn`

The Amazon Resource Name (ARN) of the trust store.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

The following example configures mutual TLS verify mode.

#### YAML

```yaml

myHTTPSListener:
  Type: 'AWS::ElasticLoadBalancingV2::Listener'
  Properties:
    LoadBalancerArn: !Ref myLoadBalancer
    Protocol: HTTPS
    Port: 443
    DefaultActions:
      - Type: forward
        TargetGroupArn: !Ref myTargetGroup
    SslPolicy: "ELBSecurityPolicy-TLS13-1-2-2021-06"
    Certificates:
      - CertificateArn: "arn:aws:acm:us-west-2:123456789012:certificate/88ca7932-756c-46f1-a70d-03fa7EXAMPLE"
    MutualAuthentication:
      - Mode: verify
        TrustStoreArn: !Ref myTrustStore
```

#### JSON

```json

{
    "myHTTPSListener": {
        "Type": "AWS::ElasticLoadBalancingV2::Listener",
        "Properties": {
            "LoadBalancerArn": {
                "Ref": "myLoadBalancer"
            },
            "Protocol": "HTTPS",
            "Port": 443,
            "DefaultActions": [
                {
                    "Type": "forward",
                    "TargetGroupArn": {
                        "Ref": "myTargetGroup"
                    }
                }
            ],
            "SslPolicy": "ELBSecurityPolicy-TLS13-1-2-2021-06",
            "Certificates": [
                {
                    "CertificateArn": "arn:aws:acm:us-west-2:123456789012:certificate/88ca7932-756c-46f1-a70d-03fa7EXAMPLE"
                }
            ],
            "MutualAuthentication": [
                {
                    "Mode": "verify",
                    "TrustStoreArn": {
                        "Ref": "myTrustStore"
                    }
                }
            ]
        }
    }
}
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ListenerAttribute

RedirectConfig
