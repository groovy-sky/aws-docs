This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::SecurityConfiguration

Creates a new security configuration. A security configuration is a set of security properties that can be used by AWS Glue. You can use a security configuration to encrypt data at rest. For information about using security configurations in AWS Glue, see [Encrypting Data Written by Crawlers, Jobs, and Development Endpoints](../../../glue/latest/dg/encryption-security-configuration.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Glue::SecurityConfiguration",
  "Properties" : {
      "EncryptionConfiguration" : EncryptionConfiguration,
      "Name" : String
    }
}

```

### YAML

```yaml

Type: AWS::Glue::SecurityConfiguration
Properties:
  EncryptionConfiguration:
    EncryptionConfiguration
  Name: String

```

## Properties

`EncryptionConfiguration`

The encryption configuration associated with this security configuration.

_Required_: Yes

_Type_: [EncryptionConfiguration](aws-properties-glue-securityconfiguration-encryptionconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the security configuration.

_Required_: Yes

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\t]*`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

### Fn::GetAtt

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Glue::SchemaVersionMetadata

CloudWatchEncryption

All content copied from https://docs.aws.amazon.com/.
