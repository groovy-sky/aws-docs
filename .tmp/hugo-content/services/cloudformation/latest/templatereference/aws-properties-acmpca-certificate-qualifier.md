This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ACMPCA::Certificate Qualifier

Defines a `PolicyInformation` qualifier. AWS Private CA supports the [certification\
practice statement (CPS) qualifier](https://datatracker.ietf.org/doc/html/rfc5280) defined in RFC 5280.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CpsUri" : String
}

```

### YAML

```yaml

  CpsUri: String

```

## Properties

`CpsUri`

Contains a pointer to a certification practice statement (CPS) published by the
CA.

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PolicyQualifierInfo

Subject

All content copied from https://docs.aws.amazon.com/.
