This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PCAConnectorAD::Template ExtensionsV4

Certificate extensions for v4 template schema

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ApplicationPolicies" : ApplicationPolicies,
  "KeyUsage" : KeyUsage
}

```

### YAML

```yaml

  ApplicationPolicies:
    ApplicationPolicies
  KeyUsage:
    KeyUsage

```

## Properties

`ApplicationPolicies`

Application policies specify what the certificate is used for and its purpose.

_Required_: No

_Type_: [ApplicationPolicies](aws-properties-pcaconnectorad-template-applicationpolicies.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KeyUsage`

The key usage extension defines the purpose (e.g., encipherment, signature) of the key
contained in the certificate.

_Required_: Yes

_Type_: [KeyUsage](aws-properties-pcaconnectorad-template-keyusage.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ExtensionsV3

GeneralFlagsV2

All content copied from https://docs.aws.amazon.com/.
