This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ACMPCA::Certificate PolicyInformation

Defines the X.509 `CertificatePolicies` extension.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CertPolicyId" : String,
  "PolicyQualifiers" : [ PolicyQualifierInfo, ... ]
}

```

### YAML

```yaml

  CertPolicyId: String
  PolicyQualifiers:
    - PolicyQualifierInfo

```

## Properties

`CertPolicyId`

Specifies the object identifier (OID) of the certificate policy under which the
certificate was issued. For more information, see NIST's definition of [Object Identifier\
(OID)](https://csrc.nist.gov/glossary/term/Object_Identifier).

_Required_: Yes

_Type_: String

_Pattern_: `([0-2])\.([0-9]|([0-3][0-9]))((\.([0-9]+)){0,126})`

_Minimum_: `0`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PolicyQualifiers`

Modifies the given `CertPolicyId` with a qualifier. AWS Private CA supports the
certification practice statement (CPS) qualifier.

_Required_: No

_Type_: Array of [PolicyQualifierInfo](aws-properties-acmpca-certificate-policyqualifierinfo.md)

_Minimum_: `1`

_Maximum_: `20`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OtherName

PolicyQualifierInfo

All content copied from https://docs.aws.amazon.com/.
