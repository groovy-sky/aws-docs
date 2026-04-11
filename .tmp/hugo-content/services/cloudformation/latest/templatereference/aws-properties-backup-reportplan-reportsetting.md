This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Backup::ReportPlan ReportSetting

Contains detailed information about a report setting.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Accounts" : [ String, ... ],
  "FrameworkArns" : [ String, ... ],
  "OrganizationUnits" : [ String, ... ],
  "Regions" : [ String, ... ],
  "ReportTemplate" : String
}

```

### YAML

```yaml

  Accounts:
    - String
  FrameworkArns:
    - String
  OrganizationUnits:
    - String
  Regions:
    - String
  ReportTemplate: String

```

## Properties

`Accounts`

These are the accounts to be included in the report.

Use string value of `ROOT` to include all organizational units.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FrameworkArns`

The Amazon Resource Names (ARNs) of the frameworks a report covers.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OrganizationUnits`

These are the Organizational Units to be included in the report.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Regions`

These are the Regions to be included in the report.

Use the wildcard as the string value to include all Regions.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReportTemplate`

Identifies the report template for the report. Reports are built using a report
template. The report templates are:

`RESOURCE_COMPLIANCE_REPORT | CONTROL_COMPLIANCE_REPORT | BACKUP_JOB_REPORT |
            COPY_JOB_REPORT | RESTORE_JOB_REPORT | SCAN_JOB_REPORT`

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ReportDeliveryChannel

Tag

All content copied from https://docs.aws.amazon.com/.
