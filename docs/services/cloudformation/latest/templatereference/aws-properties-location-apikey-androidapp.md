This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Location::APIKey AndroidApp

The `AndroidApp` property type specifies Property description not available. for an [AWS::Location::APIKey](aws-resource-location-apikey.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CertificateFingerprint" : String,
  "Package" : String
}

```

### YAML

```yaml

  CertificateFingerprint: String
  Package: String

```

## Properties

`CertificateFingerprint`

Property description not available.

_Required_: Yes

_Type_: String

_Pattern_: `^([A-Fa-f0-9]{2}:){19}[A-Fa-f0-9]{2}$`

_Minimum_: `59`

_Maximum_: `59`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Package`

Property description not available.

_Required_: Yes

_Type_: String

_Pattern_: `^([A-Za-z][A-Za-z\d_]*\.)+[A-Za-z][A-Za-z\d_]*$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Location::APIKey

ApiKeyRestrictions

All content copied from https://docs.aws.amazon.com/.
