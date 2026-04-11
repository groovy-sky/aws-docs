This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Kendra::DataSource InlineCustomDocumentEnrichmentConfiguration

Provides the configuration information for applying basic logic to alter document
metadata and content when ingesting documents into Amazon Kendra. To apply advanced
logic, to go beyond what you can do with basic logic, see [HookConfiguration](../../../kendra/latest/dg/api-hookconfiguration.md).

For more information, see [Customizing document metadata\
during the ingestion process](../../../kendra/latest/dg/custom-document-enrichment.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Condition" : DocumentAttributeCondition,
  "DocumentContentDeletion" : Boolean,
  "Target" : DocumentAttributeTarget
}

```

### YAML

```yaml

  Condition:
    DocumentAttributeCondition
  DocumentContentDeletion: Boolean
  Target:
    DocumentAttributeTarget

```

## Properties

`Condition`

Configuration of the condition used for the target document attribute or metadata
field when ingesting documents into Amazon Kendra.

_Required_: No

_Type_: [DocumentAttributeCondition](aws-properties-kendra-datasource-documentattributecondition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DocumentContentDeletion`

`TRUE` to delete content if the condition used for the target attribute is
met.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Target`

Configuration of the target document attribute or metadata field when ingesting
documents into Amazon Kendra. You can also include a value.

_Required_: No

_Type_: [DocumentAttributeTarget](aws-properties-kendra-datasource-documentattributetarget.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HookConfiguration

OneDriveConfiguration

All content copied from https://docs.aws.amazon.com/.
