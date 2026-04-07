This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PCAConnectorAD::Template

Creates an Active Directory compatible certificate template. The connectors issues certificates
using these templates based on the requester’s Active Directory group membership.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::PCAConnectorAD::Template",
  "Properties" : {
      "ConnectorArn" : String,
      "Definition" : TemplateDefinition,
      "Name" : String,
      "ReenrollAllCertificateHolders" : Boolean,
      "Tags" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::PCAConnectorAD::Template
Properties:
  ConnectorArn: String
  Definition:
    TemplateDefinition
  Name: String
  ReenrollAllCertificateHolders: Boolean
  Tags:
    Key: Value

```

## Properties

`ConnectorArn`

The Amazon Resource Name (ARN) that was returned when you called [CreateConnector](https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateConnector.html).

_Required_: Yes

_Type_: String

_Pattern_: `^arn:[\w-]+:pca-connector-ad:[\w-]+:[0-9]+:connector\/[0-9a-f]{8}(-[0-9a-f]{4}){3}-[0-9a-f]{12}$`

_Minimum_: `5`

_Maximum_: `200`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Definition`

Template configuration to define the information included in certificates. Define
certificate validity and renewal periods, certificate request handling and enrollment
options, key usage extensions, application policies, and cryptography settings.

_Required_: Yes

_Type_: [TemplateDefinition](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pcaconnectorad-template-templatedefinition.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

Name of the templates. Template names must be unique.

_Required_: Yes

_Type_: String

_Pattern_: `^(?!^\s+$)((?![\x5c'\x2b,;<=>#\x22])([\x20-\x7E]))+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ReenrollAllCertificateHolders`

This setting allows the major version of a template to be increased automatically. All
members of Active Directory groups that are allowed to enroll with a template will
receive a new certificate issued using that template.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Metadata assigned to a template consisting of a key-value pair.

_Required_: No

_Type_: Object of String

_Pattern_: `.+`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`TemplateArn`

The Amazon Resource Name (ARN) that was returned when you called [CreateTemplate](https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateTemplate.html) .

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::PCAConnectorAD::ServicePrincipalName

ApplicationPolicies
