This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CustomerProfiles::Integration IncrementalPullConfig

Specifies the configuration used when importing incremental records from the
source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DatetimeTypeFieldName" : String
}

```

### YAML

```yaml

  DatetimeTypeFieldName: String

```

## Properties

`DatetimeTypeFieldName`

A field that specifies the date time or timestamp field as the criteria to use when
importing incremental records from the source.

_Required_: No

_Type_: String

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FlowDefinition

MarketoSourceProperties

All content copied from https://docs.aws.amazon.com/.
