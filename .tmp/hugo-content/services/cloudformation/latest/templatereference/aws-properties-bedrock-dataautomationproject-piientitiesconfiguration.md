This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataAutomationProject PIIEntitiesConfiguration

Configuration for detecting and redacting Personally Identifiable Information (PII) entities.
Specify which PII entity types to detect and the redaction mask mode.
If not provided, defaults to ALL entity types with ENTITY\_TYPE redaction mask mode.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PiiEntityTypes" : [ String, ... ],
  "RedactionMaskMode" : String
}

```

### YAML

```yaml

  PiiEntityTypes:
    - String
  RedactionMaskMode: String

```

## Properties

`PiiEntityTypes`

List of PII entity types to detect/redact in the output.
Choose from specific entity types (such as ADDRESS, NAME, EMAIL, PHONE, US\_SOCIAL\_SECURITY\_NUMBER) or specify ALL to detect all supported PII types.
If not specified, defaults to ALL.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RedactionMaskMode`

Defines how detected PII entities are masked in redacted output files.
Set to PII to replace all detected entities with a generic \[PII\] marker regardless of entity type.
Set to ENTITY\_TYPE to replace each detected entity with its specific type marker (for example, \[NAME\], \[EMAIL\], \[ADDRESS\]).
This setting only applies when detectionMode is set to DETECTION\_AND\_REDACTION.
If not specified, defaults to ENTITY\_TYPE.

_Required_: No

_Type_: String

_Allowed values_: `PII | ENTITY_TYPE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OverrideConfiguration

SensitiveDataConfiguration

All content copied from https://docs.aws.amazon.com/.
