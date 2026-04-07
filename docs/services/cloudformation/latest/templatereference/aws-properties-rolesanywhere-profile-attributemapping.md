This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RolesAnywhere::Profile AttributeMapping

A mapping applied to the authenticating end-entity certificate.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CertificateField" : String,
  "MappingRules" : [ MappingRule, ... ]
}

```

### YAML

```yaml

  CertificateField: String
  MappingRules:
    - MappingRule

```

## Properties

`CertificateField`

Fields (x509Subject, x509Issuer and x509SAN) within X.509 certificates.

_Required_: Yes

_Type_: String

_Allowed values_: `x509Subject | x509Issuer | x509SAN`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MappingRules`

A list of mapping entries for every supported specifier or sub-field.

_Required_: Yes

_Type_: Array of [MappingRule](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-rolesanywhere-profile-mappingrule.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::RolesAnywhere::Profile

MappingRule
