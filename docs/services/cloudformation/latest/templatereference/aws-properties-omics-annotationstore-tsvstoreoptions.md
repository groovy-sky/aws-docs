This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Omics::AnnotationStore TsvStoreOptions

The store's parsing options.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AnnotationType" : String,
  "FormatToHeader" : {Key: Value, ...},
  "Schema" : [ {Key: Value, ...}, ... ]
}

```

### YAML

```yaml

  AnnotationType: String
  FormatToHeader:
    Key: Value
  Schema:
    -
    Key: Value

```

## Properties

`AnnotationType`

The store's annotation type.

_Required_: No

_Type_: String

_Allowed values_: `GENERIC | CHR_POS | CHR_POS_REF_ALT | CHR_START_END_ONE_BASE | CHR_START_END_REF_ALT_ONE_BASE | CHR_START_END_ZERO_BASE | CHR_START_END_REF_ALT_ZERO_BASE`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FormatToHeader`

The store's header key to column name mapping.

_Required_: No

_Type_: Object of String

_Pattern_: `.+`

_Minimum_: `0`

_Maximum_: `1000`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Schema`

The schema of an annotation store.

_Required_: No

_Type_: Array of Object

_Minimum_: `1`

_Maximum_: `5000`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StoreOptions

AWS::Omics::Configuration

All content copied from https://docs.aws.amazon.com/.
