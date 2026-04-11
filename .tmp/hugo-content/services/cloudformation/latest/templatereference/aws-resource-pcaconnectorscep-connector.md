This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PCAConnectorSCEP::Connector

Connector for SCEP is a service that links AWS Private Certificate Authority to your SCEP-enabled devices. The connector brokers the exchange of certificates from AWS Private CA to your SCEP-enabled devices and mobile device management systems. The connector is a complex type that contains the connector's configuration settings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::PCAConnectorSCEP::Connector",
  "Properties" : {
      "CertificateAuthorityArn" : String,
      "MobileDeviceManagement" : MobileDeviceManagement,
      "Tags" : {Key: Value, ...},
      "VpcEndpointId" : String
    }
}

```

### YAML

```yaml

Type: AWS::PCAConnectorSCEP::Connector
Properties:
  CertificateAuthorityArn: String
  MobileDeviceManagement:
    MobileDeviceManagement
  Tags:
    Key: Value
  VpcEndpointId: String

```

## Properties

`CertificateAuthorityArn`

The Amazon Resource Name (ARN) of the certificate authority associated with the connector.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws(-[a-z]+)*:acm-pca:[a-z]+(-[a-z]+)+-[1-9]\d*:\d{12}:certificate-authority\/[0-9a-f]{8}(-[0-9a-f]{4}){3}-[0-9a-f]{12}$`

_Minimum_: `5`

_Maximum_: `200`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MobileDeviceManagement`

Contains settings relevant to the mobile device management system that you chose for the connector. If you didn't configure `MobileDeviceManagement`, then the connector is for general-purpose use and this object is empty.

_Required_: No

_Type_: [MobileDeviceManagement](aws-properties-pcaconnectorscep-connector-mobiledevicemanagement.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Property description not available.

_Required_: No

_Type_: Object of String

_Pattern_: `.+`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcEndpointId`

Property description not available.

_Required_: No

_Type_: String

_Minimum_: `5`

_Maximum_: `200`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

### Fn::GetAtt

`ConnectorArn`

The Amazon Resource Name (ARN) of the connector.

`Endpoint`

The connector's HTTPS public SCEP URL.

`Type`

The connector type.

## Examples

- [Create a general-purpose SCEP connector and challenge resource](#aws-resource-pcaconnectorscep-connector--examples--Create_a_general-purpose_SCEP_connector_and_challenge_resource)

- [Create connector to use with Microsoft Intune](#aws-resource-pcaconnectorscep-connector--examples--Create_connector_to_use_with_Microsoft_Intune)

### Create a general-purpose SCEP connector and challenge resource

The following example creates a AWS Private Certificate Authority (CA) general-purpose connector with a challenge password. Before you create a connector, you must complete a few prerequisites, including creating a private CA in AWS Private Certificate Authority (CA). For more information, see [Set up Connector for SCEP](../../../privateca/latest/userguide/connector-for-scep-setting-up.md).

#### JSON

```json

{
   "AWSTemplateFormatVersion": "2010-09-09",
   "Description": "Cloudformation template to set up a general-purpose connector for SCEP and challenge password.",
   "Resources": {
      "RootCA": {
         "Type": "AWS::ACMPCA::CertificateAuthority",
         "Properties": {
         "Type": "ROOT",
         "KeyAlgorithm": "RSA_2048",
         "SigningAlgorithm": "SHA256WITHRSA",
         "Subject": {
            "Country": "US",
            "Organization": "string",
            "OrganizationalUnit": "string",
            "DistinguishedNameQualifier": "string",
            "State": "string",
            "CommonName": "123",
            "SerialNumber": "string",
            "Locality": "string",
            "Title": "string",
            "Surname": "string",
            "GivenName": "string",
            "Initials": "DG",
            "Pseudonym": "string",
            "GenerationQualifier": "DBG"
         },
         "RevocationConfiguration": {
            "CrlConfiguration": {
               "Enabled": false
            }
         }
         }
      },
      "RootCACertificate": {
         "Type": "AWS::ACMPCA::Certificate",
         "Properties": {
         "CertificateAuthorityArn": {
            "Fn::Ref": "RootCA"
         },
         "CertificateSigningRequest": {
            "Fn::GetAtt": [
               "RootCA",
               "CertificateSigningRequest"
            ]
         },
         "SigningAlgorithm": "SHA256WITHRSA",
         "TemplateArn": "arn:aws:acm-pca:::template/RootCACertificate/V1",
         "Validity": {
            "Type": "YEARS",
            "Value": 100
         }
         }
      },
      "RootCAActivation": {
         "Type": "AWS::ACMPCA::CertificateAuthorityActivation",
         "Properties": {
         "CertificateAuthorityArn": {
            "Fn::Ref": "RootCA"
         },
         "Certificate": {
            "Fn::GetAtt": [
               "RootCACertificate",
               "Certificate"
            ]
         },
         "Status": "ACTIVE"
         }
      },
      "RootCAResourceShare": {
         "DependsOn": "RootCAActivation",
         "Type": "AWS::RAM::ResourceShare",
         "Properties": {
         "Name": "RootCAResourceShare",
         "PermissionArns": [
            "arn:aws:ram::aws:permission/AWSRAMBlankEndEntityCertificateAPICSRPassthroughIssuanceCertificateAuthority"
         ],
         "ResourceArns": [
            {
               "Fn::Ref": "RootCA"
            }
         ],
         "Sources": [
            {
               "Fn::Ref": "AWS::AccountId"
            }
         ],
         "Principals": [
            "pca-connector-scep.amazonaws.com"
         ]
         }
      },
      "GeneralPurposeConnector": {
         "DependsOn": "RootCAResourceShare",
         "Type": "AWS::PCAConnectorSCEP::Connector",
         "Properties": {
         "CertificateAuthorityArn": {
            "Fn::Ref": "RootCA"
         }
         }
      },
      "GeneralPurposeConnectorChallenge": {
         "DependsOn": "GeneralPurposeConnector",
         "Type": "AWS::PCAConnectorSCEP::Challenge",
         "Properties": {
         "ConnectorArn": {
            "Fn::Ref": "GeneralPurposeConnector"
         }
         }
      }
   },
   "Outputs": {
      "GeneralPurposeConnector": {
         "Value": {
         "Fn::Ref": "GeneralPurposeConnector"
         }
      },
      "GeneralPurposeConnectorChallenge": {
         "Value": {
         "Fn::Ref": "GeneralPurposeConnectorChallenge"
         }
      }
   }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: '2010-09-09'
Description: Cloudformation template to set up a general-purpose connector for SCEP and challenge password.
Resources:
   RootCA:
      Type: AWS::ACMPCA::CertificateAuthority
      Properties:
      Type: ROOT
      KeyAlgorithm: RSA_2048
      SigningAlgorithm: SHA256WITHRSA
      Subject:
         Country: US
         Organization: string
         OrganizationalUnit: string
         DistinguishedNameQualifier: string
         State: string
         CommonName: '123'
         SerialNumber: string
         Locality: string
         Title: string
         Surname: string
         GivenName: string
         Initials: DG
         Pseudonym: string
         GenerationQualifier: DBG
      RevocationConfiguration:
         CrlConfiguration:
            Enabled: false
   RootCACertificate:
      Type: AWS::ACMPCA::Certificate
      Properties:
      CertificateAuthorityArn: !Ref RootCA
      CertificateSigningRequest: !GetAtt RootCA.CertificateSigningRequest
      SigningAlgorithm: SHA256WITHRSA
      TemplateArn: arn:aws:acm-pca:::template/RootCACertificate/V1
      Validity:
         Type: YEARS
         Value: 100
   RootCAActivation:
      Type: AWS::ACMPCA::CertificateAuthorityActivation
      Properties:
      CertificateAuthorityArn: !Ref RootCA
      Certificate: !GetAtt RootCACertificate.Certificate
      Status: ACTIVE
   RootCAResourceShare:
      DependsOn: RootCAActivation
      Type: AWS::RAM::ResourceShare
      Properties:
      Name: RootCAResourceShare
      PermissionArns:
         - arn:aws:ram::aws:permission/AWSRAMBlankEndEntityCertificateAPICSRPassthroughIssuanceCertificateAuthority
      ResourceArns:
         - !Ref RootCA
      Sources:
         - !Ref AWS::AccountId
      Principals:
         - pca-connector-scep.amazonaws.com
   GeneralPurposeConnector:
      DependsOn: RootCAResourceShare
      Type: AWS::PCAConnectorSCEP::Connector
      Properties:
      CertificateAuthorityArn: !Ref RootCA
   GeneralPurposeConnectorChallenge:
      DependsOn: GeneralPurposeConnector
      Type: AWS::PCAConnectorSCEP::Challenge
      Properties:
      ConnectorArn: !Ref GeneralPurposeConnector
Outputs:
   GeneralPurposeConnector:
      Value: !Ref GeneralPurposeConnector
   GeneralPurposeConnectorChallenge:
      Value: !Ref GeneralPurposeConnectorChallenge
```

### Create connector to use with Microsoft Intune

The following example creates a AWS Private Certificate Authority (CA) connector to use with Microsoft Intune. Before you create a connector, you must complete a few prerequisites, including creating a private CA in AWS Private Certificate Authority (CA). For more information, see [Set up Connector for SCEP](../../../privateca/latest/userguide/connector-for-scep-setting-up.md).

#### JSON

```json

{
      "AWSTemplateFormatVersion": "2010-09-09",
      "Description": "Cloudformation template to set up a connector to use with Microsoft Intune.",
      "Resources": {
         "RootCA": {
            "Type": "AWS::ACMPCA::CertificateAuthority",
            "Properties": {
            "Type": "ROOT",
            "KeyAlgorithm": "RSA_2048",
            "SigningAlgorithm": "SHA256WITHRSA",
            "Subject": {
               "Country": "US",
               "Organization": "string",
               "OrganizationalUnit": "string",
               "DistinguishedNameQualifier": "string",
               "State": "string",
               "CommonName": "123",
               "SerialNumber": "string",
               "Locality": "string",
               "Title": "string",
               "Surname": "string",
               "GivenName": "string",
               "Initials": "DG",
               "Pseudonym": "string",
               "GenerationQualifier": "DBG"
            },
            "RevocationConfiguration": {
               "CrlConfiguration": {
                  "Enabled": false
               }
            }
            }
         },
         "RootCACertificate": {
            "Type": "AWS::ACMPCA::Certificate",
            "Properties": {
            "CertificateAuthorityArn": {
               "Fn::Ref": "RootCA"
            },
            "CertificateSigningRequest": {
               "Fn::GetAtt": [
                  "RootCA",
                  "CertificateSigningRequest"
               ]
            },
            "SigningAlgorithm": "SHA256WITHRSA",
            "TemplateArn": "arn:aws:acm-pca:::template/RootCACertificate/V1",
            "Validity": {
               "Type": "YEARS",
               "Value": 100
            }
            }
         },
         "RootCAActivation": {
            "Type": "AWS::ACMPCA::CertificateAuthorityActivation",
            "Properties": {
            "CertificateAuthorityArn": {
               "Fn::Ref": "RootCA"
            },
            "Certificate": {
               "Fn::GetAtt": [
                  "RootCACertificate",
                  "Certificate"
               ]
            },
            "Status": "ACTIVE"
            }
         },
         "RootCAResourceShare": {
            "DependsOn": "RootCAActivation",
            "Type": "AWS::RAM::ResourceShare",
            "Properties": {
            "Name": "RootCAResourceShare",
            "PermissionArns": [
               "arn:aws:ram::aws:permission/AWSRAMBlankEndEntityCertificateAPICSRPassthroughIssuanceCertificateAuthority"
            ],
            "ResourceArns": [
               {
                  "Fn::Ref": "RootCA"
               }
            ],
            "Sources": [
               {
                  "Fn::Ref": "AWS::AccountId"
               }
            ],
            "Principals": [
               "pca-connector-scep.amazonaws.com"
            ]
            }
         },
         "IntuneConnector": {
            "DependsOn": "RootCAResourceShare",
            "Type": "AWS::PCAConnectorSCEP::Connector",
            "Properties": {
            "CertificateAuthorityArn": {
               "Fn::Ref": "RootCA"
            },
            "MobileDeviceManagement": {
               "Intune": {
                  "AzureApplicationId": "222-222-222-222-222",
                  "Domain": "example.onmicrosoft.com"
               }
            }
            }
         }
      },
      "Outputs": {
         "IntuneConnector": {
            "Value": {
            "Fn::Ref": "IntuneConnector"
            }
         }
      }
   }
```

#### YAML

```yaml

AWSTemplateFormatVersion: '2010-09-09'
Description: Cloudformation template to set up a connector to use with Microsoft Intune.
Resources:
   RootCA:
      Type: AWS::ACMPCA::CertificateAuthority
      Properties:
      Type: ROOT
      KeyAlgorithm: RSA_2048
      SigningAlgorithm: SHA256WITHRSA
      Subject:
         Country: US
         Organization: string
         OrganizationalUnit: string
         DistinguishedNameQualifier: string
         State: string
         CommonName: '123'
         SerialNumber: string
         Locality: string
         Title: string
         Surname: string
         GivenName: string
         Initials: DG
         Pseudonym: string
         GenerationQualifier: DBG
      RevocationConfiguration:
         CrlConfiguration:
            Enabled: false
   RootCACertificate:
      Type: AWS::ACMPCA::Certificate
      Properties:
      CertificateAuthorityArn: !Ref RootCA
      CertificateSigningRequest: !GetAtt RootCA.CertificateSigningRequest
      SigningAlgorithm: SHA256WITHRSA
      TemplateArn: arn:aws:acm-pca:::template/RootCACertificate/V1
      Validity:
         Type: YEARS
         Value: 100
   RootCAActivation:
      Type: AWS::ACMPCA::CertificateAuthorityActivation
      Properties:
      CertificateAuthorityArn: !Ref RootCA
      Certificate: !GetAtt RootCACertificate.Certificate
      Status: ACTIVE
   RootCAResourceShare:
      DependsOn: RootCAActivation
      Type: AWS::RAM::ResourceShare
      Properties:
      Name: RootCAResourceShare
      PermissionArns:
         - arn:aws:ram::aws:permission/AWSRAMBlankEndEntityCertificateAPICSRPassthroughIssuanceCertificateAuthority
      ResourceArns:
         - !Ref RootCA
      Sources:
         - !Ref AWS::AccountId
      Principals:
         - pca-connector-scep.amazonaws.com
   IntuneConnector:
      DependsOn: RootCAResourceShare
      Type: AWS::PCAConnectorSCEP::Connector
      Properties:
      CertificateAuthorityArn: !Ref RootCA
      MobileDeviceManagement:
         Intune:
            AzureApplicationId: "222-222-222-222-222"
            Domain: "example.onmicrosoft.com"
   Outputs:
   IntuneConnector:
      Value: !Ref IntuneConnector
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::PCAConnectorSCEP::Challenge

IntuneConfiguration

All content copied from https://docs.aws.amazon.com/.
