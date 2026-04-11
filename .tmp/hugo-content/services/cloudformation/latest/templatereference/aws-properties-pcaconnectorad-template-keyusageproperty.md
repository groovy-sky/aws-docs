This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PCAConnectorAD::Template KeyUsageProperty

The key usage property defines the purpose of the private key contained in the
certificate. You can specify specific purposes using property flags or all by using
property type ALL.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PropertyFlags" : KeyUsagePropertyFlags,
  "PropertyType" : String
}

```

### YAML

```yaml

  PropertyFlags:
    KeyUsagePropertyFlags
  PropertyType: String

```

## Properties

`PropertyFlags`

You can specify key usage for encryption, key agreement, and signature. You can use
property flags or property type but not both.

_Required_: No

_Type_: [KeyUsagePropertyFlags](aws-properties-pcaconnectorad-template-keyusagepropertyflags.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PropertyType`

You can specify all key usages using property type ALL. You can use property type or
property flags but not both.

_Required_: No

_Type_: String

_Allowed values_: `ALL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

KeyUsageFlags

KeyUsagePropertyFlags

All content copied from https://docs.aws.amazon.com/.
