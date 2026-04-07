This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PCAConnectorAD::Connector

Creates a connector between AWS Private CA and an Active Directory. You must specify the private CA,
directory ID, and security groups.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::PCAConnectorAD::Connector",
  "Properties" : {
      "CertificateAuthorityArn" : String,
      "DirectoryId" : String,
      "Tags" : {Key: Value, ...},
      "VpcInformation" : VpcInformation
    }
}

```

### YAML

```yaml

Type: AWS::PCAConnectorAD::Connector
Properties:
  CertificateAuthorityArn: String
  DirectoryId: String
  Tags:
    Key: Value
  VpcInformation:
    VpcInformation

```

## Properties

`CertificateAuthorityArn`

The Amazon Resource Name (ARN) of the certificate authority being used.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:[\w-]+:acm-pca:[\w-]+:[0-9]+:certificate-authority\/[0-9a-f]{8}(-[0-9a-f]{4}){3}-[0-9a-f]{12}$`

_Minimum_: `5`

_Maximum_: `200`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DirectoryId`

The identifier of the Active Directory.

_Required_: Yes

_Type_: String

_Pattern_: `^d-[0-9a-f]{10}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Metadata assigned to a connector consisting of a key-value pair.

_Required_: No

_Type_: Object of String

_Pattern_: `.+`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcInformation`

Information of the VPC and security group(s) used with the connector.

_Required_: Yes

_Type_: [VpcInformation](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pcaconnectorad-connector-vpcinformation.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ConnectorArn`

The Amazon Resource Name (ARN) that was returned when you called [CreateConnector](https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateConnector.html).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS Private Certificate Authority for Active Directory

VpcInformation
