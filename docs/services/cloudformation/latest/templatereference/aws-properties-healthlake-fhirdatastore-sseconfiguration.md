This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::HealthLake::FHIRDatastore SseConfiguration

The server-side encryption key configuration for a customer-provided encryption
key.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "KmsEncryptionConfig" : KmsEncryptionConfig
}

```

### YAML

```yaml

  KmsEncryptionConfig:
    KmsEncryptionConfig

```

## Properties

`KmsEncryptionConfig`

The server-side encryption key configuration for a customer provided encryption
key.

_Required_: Yes

_Type_: [KmsEncryptionConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-healthlake-fhirdatastore-kmsencryptionconfig.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

PreloadDataConfig

Tag
