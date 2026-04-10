This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ObservabilityAdmin::OrganizationCentralizationRule SourceLogsConfiguration

Configuration for selecting and handling source log groups for centralization.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DataSourceSelectionCriteria" : String,
  "EncryptedLogGroupStrategy" : String,
  "LogGroupSelectionCriteria" : String
}

```

### YAML

```yaml

  DataSourceSelectionCriteria: String
  EncryptedLogGroupStrategy: String
  LogGroupSelectionCriteria: String

```

## Properties

`DataSourceSelectionCriteria`

The selection criteria that specifies which data sources to centralize. The selection criteria uses the same filter expression format
as `LogGroupSelectionCriteria`, but operates on `DataSourceName` and
`DataSourceType` operands. When both `LogGroupSelectionCriteria` and
`DataSourceSelectionCriteria` are specified, a log event must match both criteria
to be centralized.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EncryptedLogGroupStrategy`

A strategy determining whether to centralize source log groups that are encrypted with
customer managed KMS keys (CMK). ALLOW will consider CMK encrypted source log groups for
centralization while SKIP will skip CMK encrypted source log groups from
centralization.

_Required_: Yes

_Type_: String

_Allowed values_: `ALLOW | SKIP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogGroupSelectionCriteria`

The selection criteria that specifies which source log groups to centralize. The selection
criteria uses the same format as OAM link filters.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LogsEncryptionConfiguration

Tag

All content copied from https://docs.aws.amazon.com/.
