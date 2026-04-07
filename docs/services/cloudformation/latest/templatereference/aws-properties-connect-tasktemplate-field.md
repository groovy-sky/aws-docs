This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::TaskTemplate Field

Describes a single task template field.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Description" : String,
  "Id" : FieldIdentifier,
  "SingleSelectOptions" : [ String, ... ],
  "Type" : String
}

```

### YAML

```yaml

  Description: String
  Id:
    FieldIdentifier
  SingleSelectOptions:
    - String
  Type: String

```

## Properties

`Description`

The description of the field.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Id`

The unique identifier for the field.

_Required_: Yes

_Type_: [FieldIdentifier](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-tasktemplate-fieldidentifier.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SingleSelectOptions`

A list of options for a single select field.

_Required_: No

_Type_: Array of String

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

Indicates the type of field. Following are the valid field types:
`NAME` `DESCRIPTION` \| `SCHEDULED_TIME` \|
`QUICK_CONNECT` \| `URL` \| `NUMBER` \|
`TEXT` \| `TEXT_AREA` \| `DATE_TIME` \|
`BOOLEAN` \| `SINGLE_SELECT` \| `EMAIL`

_Required_: Yes

_Type_: String

_Allowed values_: `NAME | DESCRIPTION | SCHEDULED_TIME | QUICK_CONNECT | URL | NUMBER | TEXT | TEXT_AREA | DATE_TIME | BOOLEAN | SINGLE_SELECT | EMAIL | EXPIRY_DURATION | SELF_ASSIGN`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DefaultFieldValue

FieldIdentifier
