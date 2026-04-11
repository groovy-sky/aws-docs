This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Kendra::DataSource SalesforceStandardObjectConfiguration

Specifies configuration information for indexing a single standard object.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DocumentDataFieldName" : String,
  "DocumentTitleFieldName" : String,
  "FieldMappings" : [ DataSourceToIndexFieldMapping, ... ],
  "Name" : String
}

```

### YAML

```yaml

  DocumentDataFieldName: String
  DocumentTitleFieldName: String
  FieldMappings:
    - DataSourceToIndexFieldMapping
  Name: String

```

## Properties

`DocumentDataFieldName`

The name of the field in the standard object table that contains the document
contents.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DocumentTitleFieldName`

The name of the field in the standard object table that contains the document
title.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FieldMappings`

Maps attributes or field names of the standard object to Amazon Kendra index
field names. To create custom fields, use the `UpdateIndex` API before you
map to Salesforce fields. For more information, see [Mapping data source fields](../../../kendra/latest/dg/field-mapping.md). The
Salesforce data source field names must exist in your Salesforce custom metadata.

_Required_: No

_Type_: Array of [DataSourceToIndexFieldMapping](aws-properties-kendra-datasource-datasourcetoindexfieldmapping.md)

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the standard object.

_Required_: Yes

_Type_: String

_Allowed values_: `ACCOUNT | CAMPAIGN | CASE | CONTACT | CONTRACT | DOCUMENT | GROUP | IDEA | LEAD | OPPORTUNITY | PARTNER | PRICEBOOK | PRODUCT | PROFILE | SOLUTION | TASK | USER`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SalesforceStandardObjectAttachmentConfiguration

ServiceNowConfiguration

All content copied from https://docs.aws.amazon.com/.
