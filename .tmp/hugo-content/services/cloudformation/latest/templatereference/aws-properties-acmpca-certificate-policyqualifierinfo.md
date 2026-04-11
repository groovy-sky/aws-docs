This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ACMPCA::Certificate PolicyQualifierInfo

Modifies the `CertPolicyId` of a `PolicyInformation` object with
a qualifier. AWS Private CA supports the certification practice statement (CPS)
qualifier.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PolicyQualifierId" : String,
  "Qualifier" : Qualifier
}

```

### YAML

```yaml

  PolicyQualifierId: String
  Qualifier:
    Qualifier

```

## Properties

`PolicyQualifierId`

Identifies the qualifier modifying a `CertPolicyId`.

_Required_: Yes

_Type_: String

_Allowed values_: `CPS`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Qualifier`

Defines the qualifier type. AWS Private CA supports the use of a URI for a CPS qualifier
in this field.

_Required_: Yes

_Type_: [Qualifier](aws-properties-acmpca-certificate-qualifier.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PolicyInformation

Qualifier

All content copied from https://docs.aws.amazon.com/.
