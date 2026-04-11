This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::KnowledgeBase

Specifies a knowledge base.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Wisdom::KnowledgeBase",
  "Properties" : {
      "Description" : String,
      "KnowledgeBaseType" : String,
      "Name" : String,
      "RenderingConfiguration" : RenderingConfiguration,
      "ServerSideEncryptionConfiguration" : ServerSideEncryptionConfiguration,
      "SourceConfiguration" : SourceConfiguration,
      "Tags" : [ Tag, ... ],
      "VectorIngestionConfiguration" : VectorIngestionConfiguration
    }
}

```

### YAML

```yaml

Type: AWS::Wisdom::KnowledgeBase
Properties:
  Description: String
  KnowledgeBaseType: String
  Name: String
  RenderingConfiguration:
    RenderingConfiguration
  ServerSideEncryptionConfiguration:
    ServerSideEncryptionConfiguration
  SourceConfiguration:
    SourceConfiguration
  Tags:
    - Tag
  VectorIngestionConfiguration:
    VectorIngestionConfiguration

```

## Properties

`Description`

The description.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KnowledgeBaseType`

The type of knowledge base. Only CUSTOM knowledge bases allow you to upload your own content. EXTERNAL knowledge bases
support integrations with third-party systems whose content is synchronized automatically.

_Required_: Yes

_Type_: String

_Allowed values_: `EXTERNAL | CUSTOM | MESSAGE_TEMPLATES | MANAGED | QUICK_RESPONSES`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the knowledge base.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RenderingConfiguration`

Information about how to render the content.

_Required_: No

_Type_: [RenderingConfiguration](aws-properties-wisdom-knowledgebase-renderingconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServerSideEncryptionConfiguration`

This customer managed key must have a policy that allows
`kms:CreateGrant` and `kms:DescribeKey` permissions to the
IAM identity using the key to invoke Wisdom. For more information
about setting up a customer managed key for Wisdom, see [Enable Amazon Connect Wisdom\
for your instance](../../../connect/latest/adminguide/enable-wisdom.md). For information about valid ID values, see [Key\
identifiers (KeyId)](../../../kms/latest/developerguide/concepts.md#key-id) in the _AWS Key Management Service Developer_
_Guide_.

_Required_: No

_Type_: [ServerSideEncryptionConfiguration](aws-properties-wisdom-knowledgebase-serversideencryptionconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourceConfiguration`

The source of the knowledge base content. Only set this argument for EXTERNAL or Managed knowledge bases.

_Required_: No

_Type_: [SourceConfiguration](aws-properties-wisdom-knowledgebase-sourceconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags used to organize, track, or control access for this resource.

_Required_: No

_Type_: Array of [Tag](aws-properties-wisdom-knowledgebase-tag.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VectorIngestionConfiguration`

Contains details about how to ingest the documents in a data source.

_Required_: No

_Type_: [VectorIngestionConfiguration](aws-properties-wisdom-knowledgebase-vectoringestionconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the knowledge base ID.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`KnowledgeBaseArn`

The Amazon Resource Name (ARN) of the knowledge base.

`KnowledgeBaseId`

The ID of the knowledge base.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AppIntegrationsConfiguration

All content copied from https://docs.aws.amazon.com/.
