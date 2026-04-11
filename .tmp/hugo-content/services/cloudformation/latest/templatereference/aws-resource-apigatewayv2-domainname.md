This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGatewayV2::DomainName

The `AWS::ApiGatewayV2::DomainName` resource specifies a custom domain
name for your API in Amazon API Gateway (API Gateway).

You can use a custom domain name to provide a URL that's more intuitive and easier
to recall. For more information about using custom domain names, see [Set up\
Custom Domain Name for an API in API Gateway](../../../apigateway/latest/developerguide/how-to-custom-domains.md) in the _API_
_Gateway Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ApiGatewayV2::DomainName",
  "Properties" : {
      "DomainName" : String,
      "DomainNameConfigurations" : [ DomainNameConfiguration, ... ],
      "MutualTlsAuthentication" : MutualTlsAuthentication,
      "RoutingMode" : String,
      "Tags" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::ApiGatewayV2::DomainName
Properties:
  DomainName: String
  DomainNameConfigurations:
    - DomainNameConfiguration
  MutualTlsAuthentication:
    MutualTlsAuthentication
  RoutingMode: String
  Tags:
    Key: Value

```

## Properties

`DomainName`

The custom domain name for your API in Amazon API Gateway. Uppercase letters and the underscore
( `_`) character are not supported.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DomainNameConfigurations`

The domain name configurations.

_Required_: No

_Type_: Array of [DomainNameConfiguration](aws-properties-apigatewayv2-domainname-domainnameconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MutualTlsAuthentication`

The mutual TLS authentication configuration for a custom domain name.

_Required_: No

_Type_: [MutualTlsAuthentication](aws-properties-apigatewayv2-domainname-mutualtlsauthentication.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoutingMode`

The routing mode API Gateway uses to route traffic to your APIs.

_Required_: No

_Type_: String

_Allowed values_: `API_MAPPING_ONLY | ROUTING_RULE_THEN_API_MAPPING | ROUTING_RULE_ONLY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The collection of tags associated with a domain name.

_Required_: No

_Type_: Object of String

_Pattern_: `.*`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the domain name, such as
`example.mydomain.com`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`DomainNameArn`

Represents an Amazon Resource Name (ARN).

`RegionalDomainName`

The domain name associated with the regional endpoint for this custom domain name.
You set up this association by adding a DNS record that points the custom domain
name to this regional domain name.

`RegionalHostedZoneId`

The region-specific Amazon Route 53 Hosted Zone ID of the regional
endpoint.

## Examples

- [Domain name creation example](#aws-resource-apigatewayv2-domainname--examples--Domain_name_creation_example)

- [Domain name with routing mode creation example](#aws-resource-apigatewayv2-domainname--examples--Domain_name_with_routing_mode_creation_example)

### Domain name creation example

The following example creates a `DomainName` resource
called `MyDomainName`.

#### JSON

```json

{
    "MyDomainName": {
        "Type": "AWS::ApiGatewayV2::DomainName",
        "Properties": {
            "DomainName": "mydomainame.us-east-1.com",
            "DomainNameConfigurations": [
                {
                    "EndpointType": "REGIONAL",
                    "CertificateArn": "arn:aws:acm:us-east-1:123456789012:certificate/1a2b3c4d-aaaa-aaaa-aaaa-1a2b3c4d5e6f",
                    "CertificateName": "testCertificate"
                }
            ]
        }
    }
}
```

#### YAML

```yaml

MyDomainName:
  Type: 'AWS::ApiGatewayV2::DomainName'
  Properties:
    DomainName: mydomainame.us-east-1.com
    DomainNameConfigurations:
      - EndpointType: REGIONAL
        CertificateArn: >-
          arn:aws:acm:us-east-1:123456789012:certificate/1a2b3c4d-aaaa-aaaa-aaaa-1a2b3c4d5e6f
        CertificateName: testCertificate

```

### Domain name with routing mode creation example

The following example creates a `DomainName` resource with the `RoutingMode` parameter of `ROUTING_RULES_ONLY`.

#### JSON

```json

{
  "Resources": {
    "TestDomainName": {
      "Type": "AWS::ApiGatewayV2::DomainName",
      "Properties": {
        "DomainName": "mydomainame.example.com",
        "DomainNameConfigurations": [
          {
            "EndpointType": "REGIONAL",
            "CertificateArn": "arn:aws:acm:us-east-1:111122223333:certificate/fb1b9770-a305-495d-aefb-27e5e101ff3",
            "SecurityPolicy": "TLS_1_2"
          }
        ],
        "RoutingMode": "ROUTING_RULE_ONLY"
      }
    }
  }
}
```

#### YAML

```yaml

Resources:
  TestDomainName:
    Type: AWS::ApiGatewayV2::DomainName
    Properties:
      DomainName: mydomainame.example.com
      DomainNameConfigurations:
        - EndpointType: REGIONAL
          CertificateArn: arn:aws:acm:us-east-1:111122223333:certificate/fb1b9770-a305-495d-aefb-27e5e101ff3
          SecurityPolicy: TLS_1_2
      RoutingMode: "ROUTING_RULE_ONLY"
```

## See also

- [CreateDomainName](../../../apigatewayv2/latest/api-reference/domainnames.md#CreateDomainName) in the _Amazon API_
_Gateway Version 2 API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::ApiGatewayV2::Deployment

DomainNameConfiguration

All content copied from https://docs.aws.amazon.com/.
