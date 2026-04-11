This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudTrail::Dashboard Widget

Contains information about a widget on a CloudTrail Lake dashboard.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "QueryParameters" : [ String, ... ],
  "QueryStatement" : String,
  "ViewProperties" : {Key: Value, ...}
}

```

### YAML

```yaml

  QueryParameters:
    - String
  QueryStatement: String
  ViewProperties:
    Key: Value

```

## Properties

`QueryParameters`

The optional query parameters. The following query parameters are valid: `$StartTime$`, `$EndTime$`, and `$Period$`.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QueryStatement`

The query statement for the widget. For custom dashboard widgets, you can query across multiple event data stores as long as all event data stores exist in your account.

###### Note

When a query uses `?` with `eventTime`, `?` must be surrounded by single quotes as follows: `'?'`.

_Required_: Yes

_Type_: String

_Pattern_: `(?s).*`

_Minimum_: `1`

_Maximum_: `10000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ViewProperties`

The view properties for the widget. For more information about view properties, see [View properties for widgets](../../../awscloudtrail/latest/userguide/lake-widget-properties.md) in the _AWS CloudTrail User Guide_.

_Required_: No

_Type_: Object of String

_Pattern_: `^[a-zA-Z0-9._-]{3,128}$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::CloudTrail::EventDataStore

All content copied from https://docs.aws.amazon.com/.
