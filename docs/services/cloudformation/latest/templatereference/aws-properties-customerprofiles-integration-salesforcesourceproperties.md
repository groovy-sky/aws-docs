This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CustomerProfiles::Integration SalesforceSourceProperties

The properties that are applied when Salesforce is being used as a source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EnableDynamicFieldUpdate" : Boolean,
  "IncludeDeletedRecords" : Boolean,
  "Object" : String
}

```

### YAML

```yaml

  EnableDynamicFieldUpdate: Boolean
  IncludeDeletedRecords: Boolean
  Object: String

```

## Properties

`EnableDynamicFieldUpdate`

The flag that enables dynamic fetching of new (recently added) fields in the
Salesforce objects while running a flow.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludeDeletedRecords`

Indicates whether Amazon AppFlow includes deleted files in the flow run.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Object`

The object specified in the Salesforce flow source.

_Required_: Yes

_Type_: String

_Pattern_: `\S+`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3SourceProperties

ScheduledTriggerProperties

All content copied from https://docs.aws.amazon.com/.
