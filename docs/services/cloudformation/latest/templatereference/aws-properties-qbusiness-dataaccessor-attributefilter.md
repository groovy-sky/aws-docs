This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QBusiness::DataAccessor AttributeFilter

Enables filtering of responses based on document attributes or metadata fields.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AndAllFilters" : [ AttributeFilter, ... ],
  "ContainsAll" : DocumentAttribute,
  "ContainsAny" : DocumentAttribute,
  "EqualsTo" : DocumentAttribute,
  "GreaterThan" : DocumentAttribute,
  "GreaterThanOrEquals" : DocumentAttribute,
  "LessThan" : DocumentAttribute,
  "LessThanOrEquals" : DocumentAttribute,
  "NotFilter" : AttributeFilter,
  "OrAllFilters" : [ AttributeFilter, ... ]
}

```

### YAML

```yaml

  AndAllFilters:
    - AttributeFilter
  ContainsAll:
    DocumentAttribute
  ContainsAny:
    DocumentAttribute
  EqualsTo:
    DocumentAttribute
  GreaterThan:
    DocumentAttribute
  GreaterThanOrEquals:
    DocumentAttribute
  LessThan:
    DocumentAttribute
  LessThanOrEquals:
    DocumentAttribute
  NotFilter:
    AttributeFilter
  OrAllFilters:
    - AttributeFilter

```

## Properties

`AndAllFilters`

Performs a logical `AND` operation on all supplied filters.

_Required_: No

_Type_: Array of [AttributeFilter](aws-properties-qbusiness-dataaccessor-attributefilter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ContainsAll`

Returns `true` when a document contains all the specified document
attributes or metadata fields. Supported for the following [document attribute value types](../../../amazonq/latest/api-reference/api-documentattributevalue.md):
`stringListValue`.

_Required_: No

_Type_: [DocumentAttribute](aws-properties-qbusiness-dataaccessor-documentattribute.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ContainsAny`

Returns `true` when a document contains any of the specified document
attributes or metadata fields. Supported for the following [document attribute value types](../../../amazonq/latest/api-reference/api-documentattributevalue.md):
`stringListValue`.

_Required_: No

_Type_: [DocumentAttribute](aws-properties-qbusiness-dataaccessor-documentattribute.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EqualsTo`

Performs an equals operation on two document attributes or metadata fields. Supported
for the following [document attribute value types](../../../amazonq/latest/api-reference/api-documentattributevalue.md): `dateValue`,
`longValue`, `stringListValue` and
`stringValue`.

_Required_: No

_Type_: [DocumentAttribute](aws-properties-qbusiness-dataaccessor-documentattribute.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GreaterThan`

Performs a greater than operation on two document attributes or metadata fields.
Supported for the following [document attribute value types](../../../amazonq/latest/api-reference/api-documentattributevalue.md): `dateValue`
and `longValue`.

_Required_: No

_Type_: [DocumentAttribute](aws-properties-qbusiness-dataaccessor-documentattribute.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GreaterThanOrEquals`

Performs a greater or equals than operation on two document attributes or metadata
fields. Supported for the following [document attribute value types](../../../amazonq/latest/api-reference/api-documentattributevalue.md): `dateValue`
and `longValue`.

_Required_: No

_Type_: [DocumentAttribute](aws-properties-qbusiness-dataaccessor-documentattribute.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LessThan`

Performs a less than operation on two document attributes or metadata fields.
Supported for the following [document attribute value types](../../../amazonq/latest/api-reference/api-documentattributevalue.md): `dateValue`
and `longValue`.

_Required_: No

_Type_: [DocumentAttribute](aws-properties-qbusiness-dataaccessor-documentattribute.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LessThanOrEquals`

Performs a less than or equals operation on two document attributes or metadata
fields.Supported for the following [document attribute value type](../../../amazonq/latest/api-reference/api-documentattributevalue.md): `dateValue`
and `longValue`.

_Required_: No

_Type_: [DocumentAttribute](aws-properties-qbusiness-dataaccessor-documentattribute.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NotFilter`

Performs a logical `NOT` operation on all supplied filters.

_Required_: No

_Type_: [AttributeFilter](aws-properties-qbusiness-dataaccessor-attributefilter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OrAllFilters`

Performs a logical `OR` operation on all supplied filters.

_Required_: No

_Type_: Array of [AttributeFilter](aws-properties-qbusiness-dataaccessor-attributefilter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ActionFilterConfiguration

DataAccessorAuthenticationConfiguration

All content copied from https://docs.aws.amazon.com/.
