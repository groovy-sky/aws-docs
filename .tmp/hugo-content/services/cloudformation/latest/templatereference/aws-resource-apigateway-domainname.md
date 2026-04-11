This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGateway::DomainName

The `AWS::ApiGateway::DomainName` resource specifies a public custom domain name for your API in API Gateway.

To create a custom domain name for private APIs, use
[AWS::ApiGateway::DomainNameV2](../userguide/aws-resource-apigateway-domainnamev2.md).

You can use a custom domain name to provide a URL that's more intuitive and easier to recall. For more information about using custom domain names, see [Set up Custom Domain Name for an API in API Gateway](../../../apigateway/latest/developerguide/how-to-custom-domains.md) in the _API Gateway Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ApiGateway::DomainName",
  "Properties" : {
      "CertificateArn" : String,
      "DomainName" : String,
      "EndpointAccessMode" : String,
      "EndpointConfiguration" : EndpointConfiguration,
      "MutualTlsAuthentication" : MutualTlsAuthentication,
      "OwnershipVerificationCertificateArn" : String,
      "RegionalCertificateArn" : String,
      "RoutingMode" : String,
      "SecurityPolicy" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::ApiGateway::DomainName
Properties:
  CertificateArn: String
  DomainName: String
  EndpointAccessMode: String
  EndpointConfiguration:
    EndpointConfiguration
  MutualTlsAuthentication:
    MutualTlsAuthentication
  OwnershipVerificationCertificateArn: String
  RegionalCertificateArn: String
  RoutingMode: String
  SecurityPolicy: String
  Tags:
    - Tag

```

## Properties

`CertificateArn`

The reference to an AWS-managed certificate that will be used by edge-optimized endpoint or private endpoint for this domain name. AWS Certificate Manager is the only supported source.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DomainName`

The custom domain name as an API host name, for example, `my-api.example.com`.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EndpointAccessMode`

The endpoint access mode for your DomainName.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EndpointConfiguration`

The endpoint configuration of this DomainName showing the endpoint types and IP address types of the domain name.

_Required_: No

_Type_: [EndpointConfiguration](aws-properties-apigateway-domainname-endpointconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MutualTlsAuthentication`

The mutual TLS authentication configuration for a custom domain name. If specified, API Gateway
performs two-way authentication between the client and the server. Clients must present a
trusted certificate to access your API.

_Required_: No

_Type_: [MutualTlsAuthentication](aws-properties-apigateway-domainname-mutualtlsauthentication.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OwnershipVerificationCertificateArn`

The ARN of the public certificate issued by ACM to validate ownership of your custom domain. Only required when configuring mutual TLS and using an ACM imported or private CA certificate ARN as the RegionalCertificateArn.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RegionalCertificateArn`

The reference to an AWS-managed certificate that will be used for validating the regional domain name. AWS Certificate Manager is the only supported source.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoutingMode`

The routing mode for this domain name. The routing mode determines how API Gateway sends traffic from your custom domain name to your public APIs.

_Required_: No

_Type_: String

_Allowed values_: `BASE_PATH_MAPPING_ONLY | ROUTING_RULE_THEN_BASE_PATH_MAPPING | ROUTING_RULE_ONLY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecurityPolicy`

The Transport Layer Security (TLS) version + cipher suite for this DomainName.

_Required_: No

_Type_: String

_Allowed values_: `TLS_1_0 | TLS_1_2 | SecurityPolicy_TLS13_1_3_2025_09 | SecurityPolicy_TLS13_1_3_FIPS_2025_09 | SecurityPolicy_TLS13_1_2_PFS_PQ_2025_09 | SecurityPolicy_TLS13_1_2_FIPS_PQ_2025_09 | SecurityPolicy_TLS13_1_2_PQ_2025_09 | SecurityPolicy_TLS13_1_2_2021_06 | SecurityPolicy_TLS13_2025_EDGE | SecurityPolicy_TLS12_PFS_2025_EDGE | SecurityPolicy_TLS12_2018_EDGE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The collection of tags. Each tag element is associated with a given resource.

_Required_: No

_Type_: Array of [Tag](aws-properties-apigateway-domainname-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the domain name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`DistributionDomainName`

The Amazon CloudFront distribution domain name that's mapped to the custom domain name. This is only applicable for endpoints whose type is `EDGE`.

Example: `d111111abcdef8.cloudfront.net`

`DistributionHostedZoneId`

The region-agnostic Amazon Route 53 Hosted Zone ID of the edge-optimized endpoint. The only valid value is `Z2FDTNDATAQYW2` for all regions.

`DomainNameArn`

The ARN of the domain name.

`RegionalDomainName`

The domain name associated with the regional endpoint for this custom domain name. You set up this association by adding a DNS record that points the custom domain name to this regional domain name.

`RegionalHostedZoneId`

The region-specific Amazon Route 53 Hosted Zone ID of the regional endpoint.

## Examples

- [Create Custom Domain](#aws-resource-apigateway-domainname--examples--Create_Custom_Domain)

- [Create Custom Domain from Parameters](#aws-resource-apigateway-domainname--examples--Create_Custom_Domain_from_Parameters)

- [Create domain name with EndpointConfiguration](#aws-resource-apigateway-domainname--examples--Create_domain_name_with_EndpointConfiguration)

- [Create Domain Names and Zone IDs as Outputs](#aws-resource-apigateway-domainname--examples--Create_Domain_Names_and_Zone_IDs_as_Outputs)

- [Create Domain Name with routing mode ROUTING\_RULE\_ONLY](#aws-resource-apigateway-domainname--examples--Create_Domain_Name_with_routing_mode_ROUTING_RULE_ONLY)

### Create Custom Domain

The following example creates a custom domain name of `api.mydomain.com`.

#### JSON

```json

{
    "MyDomainName": {
        "Type": "AWS::ApiGateway::DomainName",
        "Properties": {
            "DomainName": "api.mydomain.com",
            "CertificateArn": "arn:aws:acm:us-east-1:111122223333:certificate/fb1b9770-a305-495d-aefb-27e5e101ff3"
        }
    }
}
```

#### YAML

```yaml

MyDomainName:
  Type: 'AWS::ApiGateway::DomainName'
  Properties:
    DomainName: api.mydomain.com
    CertificateArn: >-
      arn:aws:acm:us-east-1:111122223333:certificate/fb1b9770-a305-495d-aefb-27e5e101ff3

```

### Create Custom Domain from Parameters

The following example creates a custom domain name of `example.mydomain.com`.

#### JSON

```json

{
    "Parameters": {
        "basePath": {
            "Type": "String",
            "Default": "examplepath"
        },
        "domainName": {
            "Type": "String",
            "Default": "example.mydomain.com"
        },
        "restApiName": {
            "Type": "String",
            "Default": "exampleapi"
        }
    },
    "Resources": {
        "myCertificate": {
            "Type": "AWS::CertificateManager::Certificate",
            "Properties": {
                "DomainName": {
                    "Ref": "domainName"
                }
            }
        },
        "myDomainName": {
            "Type": "AWS::ApiGateway::DomainName",
            "Properties": {
                "CertificateArn": {
                    "Ref": "myCertificate"
                },
                "DomainName": {
                    "Ref": "domainName"
                }
            }
        },
        "myMapping": {
            "Type": "AWS::ApiGateway::BasePathMapping",
            "Properties": {
                "BasePath": {
                    "Ref": "basePath"
                },
                "DomainName": {
                    "Ref": "myDomainName"
                },
                "RestApiId": {
                    "Ref": "myRestApi"
                }
            }
        },
        "myRestApi": {
            "Type": "AWS::ApiGateway::RestApi",
            "Properties": {
                "Name": {
                    "Ref": "restApiName"
                }
            }
        }
    },
    "Outputs": {
        "domainName": {
            "Value": {
                "Fn::GetAtt": [
                    "myDomainName",
                    "DistributionDomainName"
                ]
            }
        }
    }
}
```

#### YAML

```yaml

Parameters:
  basePath:
    Type: String
    Default: examplepath
  domainName:
    Type: String
    Default: example.mydomain.com
  restApiName:
    Type: String
    Default: exampleapi
Resources:
  myCertificate:
    Type: 'AWS::CertificateManager::Certificate'
    Properties:
      DomainName: !Ref domainName
  myDomainName:
    Type: 'AWS::ApiGateway::DomainName'
    Properties:
      CertificateArn: !Ref myCertificate
      DomainName: !Ref domainName
  myMapping:
    Type: 'AWS::ApiGateway::BasePathMapping'
    Properties:
      BasePath: !Ref basePath
      DomainName: !Ref myDomainName
      RestApiId: !Ref myRestApi
  myRestApi:
    Type: 'AWS::ApiGateway::RestApi'
    Properties:
      Name: !Ref restApiName
Outputs:
  domainName:
    Value: !GetAtt
      - myDomainName
      - DistributionDomainName

```

### Create domain name with EndpointConfiguration

The following example creates a custom domain name that specifies a regional certificate ARN and an endpoint type.

#### JSON

```json

{
  "Parameters": {
    "cfnDomainName": {
      "Type": "String"
    },
    "certificateArn": {
      "Type": "String"
    },
    "type": {
      "Type": "String"
    }
  },
  "Resources": {
    "myDomainName": {
      "Type": "AWS::ApiGateway::DomainName",
      "Properties": {
        "CertificateArn": {
          "Ref": "certificateArn"
        },
        "DomainName": {
          "Ref": "cfnDomainName"
        },
        "EndpointConfiguration": {
          "Types": [
            {
              "Ref": "type"
            }
          ]
        },
        "RegionalCertificateArn": {
          "Ref": "certificateArn"
        }
      }
    }
  },
  "Outputs": {
    "DomainName": {
      "Value": {
        "Ref": "myDomainName"
      }
    }
  }
}
```

#### YAML

```yaml

Parameters:
  cfnDomainName:
    Type: String
  certificateArn:
    Type: String
  type:
    Type: String
Resources:
  myDomainName:
    Type: AWS::ApiGateway::DomainName
    Properties:
      CertificateArn: !Ref certificateArn
      DomainName: !Ref cfnDomainName
      EndpointConfiguration:
        Types:
          - !Ref type
      RegionalCertificateArn: !Ref certificateArn
Outputs:
  DomainName:
    Value: !Ref myDomainName
```

### Create Domain Names and Zone IDs as Outputs

The following example defines the distribution and regional domain names, as well as the distribution and regional hosted zone IDs, as outputs from the stack.

#### JSON

```json

{
    "Resources": {
        "myDomainName": {
            "Type": "AWS::ApiGateway::DomainName",
            "Properties": {
                "CertificateArn": {
                    "Ref": "certificateArn"
                },
                "DomainName": {
                    "Ref": "cfnDomainName"
                },
                "EndpointConfiguration": {
                    "Types": [
                        {
                            "Ref": "type"
                        }
                    ]
                },
                "RegionalCertificateArn": {
                    "Ref": "certificateArn"
                }
            }
        }
    },
    "Outputs": {
        "DistributionDomainName": {
            "Value": {
                "Fn::GetAtt": [
                    "myDomainName",
                    "DistributionDomainName"
                ]
            }
        },
        "DistributionHostedZoneId": {
            "Value": {
                "Fn::GetAtt": [
                    "myDomainName",
                    "DistributionHostedZoneId"
                ]
            }
        },
        "RegionalDomainName": {
            "Value": {
                "Fn::GetAtt": [
                    "myDomainName",
                    "RegionalDomainName"
                ]
            }
        },
        "RegionalHostedZoneId": {
            "Value": {
                "Fn::GetAtt": [
                    "myDomainName",
                    "RegionalHostedZoneId"
                ]
            }
        }
    }
}
```

#### YAML

```yaml

Resources:
  myDomainName:
    Type: 'AWS::ApiGateway::DomainName'
    Properties:
      CertificateArn: !Ref certificateArn
      DomainName: !Ref cfnDomainName
      EndpointConfiguration:
        Types:
          - !Ref type
      RegionalCertificateArn: !Ref certificateArn
Outputs:
  DistributionDomainName:
    Value: !GetAtt
      - myDomainName
      - DistributionDomainName
  DistributionHostedZoneId:
    Value: !GetAtt
      - myDomainName
      - DistributionHostedZoneId
  RegionalDomainName:
    Value: !GetAtt
      - myDomainName
      - RegionalDomainName
  RegionalHostedZoneId:
    Value: !GetAtt
      - myDomainName
      - RegionalHostedZoneId

```

### Create Domain Name with routing mode ROUTING\_RULE\_ONLY

The following example creates a domain name with a routing mode of `ROUTING_RULE_ONLY`.

#### JSON

```json

{
  "MyDomainName": {
    "Type": "AWS::ApiGateway::DomainName",
    "Properties": {
      "DomainName": "api.mydomain.com",
      "EndpointConfiguration": {
        "Types": [
          "REGIONAL"
        ]
      },
      "SecurityPolicy": "TLS_1_2",
      "RegionalCertificateArn": "arn:aws:acm:us-east-1:111122223333:certificate/fb1b9770-a305-495d-aefb-27e5e101ff3",
      "RoutingMode": "ROUTING_RULE_ONLY"
    }
  }
}
```

#### YAML

```yaml

MyDomainName:
    Type: AWS::ApiGateway::DomainName
    Properties:
      DomainName: api.mydomain.com
      EndpointConfiguration:
        Types:
          - REGIONAL
      SecurityPolicy: TLS_1_2
      RegionalCertificateArn: arn:aws:acm:us-east-1:111122223333:certificate/fb1b9770-a305-495d-aefb-27e5e101ff3
      RoutingMode: ROUTING_RULE_ONLY
```

## See also

- [domainname:create](../../../apigateway/latest/api/api-createdomainname.md) in the _Amazon API Gateway REST API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::ApiGateway::DocumentationVersion

EndpointConfiguration

All content copied from https://docs.aws.amazon.com/.
