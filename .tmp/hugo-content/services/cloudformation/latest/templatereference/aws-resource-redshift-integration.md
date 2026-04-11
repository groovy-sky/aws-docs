This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Redshift::Integration

Describes a zero-ETL or S3 integration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Redshift::Integration",
  "Properties" : {
      "AdditionalEncryptionContext" : {Key: Value, ...},
      "IntegrationName" : String,
      "KMSKeyId" : String,
      "SourceArn" : String,
      "Tags" : [ Tag, ... ],
      "TargetArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::Redshift::Integration
Properties:
  AdditionalEncryptionContext:
    Key: Value
  IntegrationName: String
  KMSKeyId: String
  SourceArn: String
  Tags:
    - Tag
  TargetArn: String

```

## Properties

`AdditionalEncryptionContext`

The encryption context for the integration. For more information, see
[Encryption context](../../../../general/index.md) in the _AWS Key Management Service Developer Guide_.

_Required_: No

_Type_: Object of String

_Pattern_: `^[\s\S]*$`

_Minimum_: `0`

_Maximum_: `131072`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IntegrationName`

The name of the integration.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KMSKeyId`

The AWS Key Management Service (AWS KMS) key identifier for the key used to encrypt the integration.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourceArn`

The Amazon Resource Name (ARN) of the database used as the source for replication.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The list of tags associated with the integration.

_Required_: No

_Type_: Array of [Tag](aws-properties-redshift-integration-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetArn`

The Amazon Resource Name (ARN) of the Amazon Redshift data warehouse to use as the target for replication.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

### Fn::GetAtt

`CreateTime`

The time (UTC) when the integration was created.

`IntegrationArn`

The Amazon Resource Name (ARN) of the integration.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
