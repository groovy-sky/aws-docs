This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PCAConnectorAD::Template KeyUsage

The key usage extension defines the purpose (e.g., encipherment, signature) of the key
contained in the certificate.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Critical" : Boolean,
  "UsageFlags" : KeyUsageFlags
}

```

### YAML

```yaml

  Critical: Boolean
  UsageFlags:
    KeyUsageFlags

```

## Properties

`Critical`

Sets the key usage extension to critical.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UsageFlags`

The key usage flags represent the purpose (e.g., encipherment, signature) of the key
contained in the certificate.

_Required_: Yes

_Type_: [KeyUsageFlags](aws-properties-pcaconnectorad-template-keyusageflags.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GeneralFlagsV4

KeyUsageFlags

All content copied from https://docs.aws.amazon.com/.
