# ReportSetting

Contains detailed information about a report setting.

## Contents

**ReportTemplate**

Identifies the report template for the report. Reports are built using a report
template. The report templates are:

`RESOURCE_COMPLIANCE_REPORT | CONTROL_COMPLIANCE_REPORT | BACKUP_JOB_REPORT |
            COPY_JOB_REPORT | RESTORE_JOB_REPORT | SCAN_JOB_REPORT`

Type: String

Required: Yes

**Accounts**

These are the accounts to be included in the report.

Use string value of `ROOT` to include all organizational units.

Type: Array of strings

Required: No

**FrameworkArns**

The Amazon Resource Names (ARNs) of the frameworks a report covers.

Type: Array of strings

Required: No

**NumberOfFrameworks**

The number of frameworks a report covers.

Type: Integer

Required: No

**OrganizationUnits**

These are the Organizational Units to be included in the report.

Type: Array of strings

Required: No

**Regions**

These are the Regions to be included in the report.

Use the wildcard as the string value to include all Regions.

Type: Array of strings

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/backup-2018-11-15/reportsetting.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/backup-2018-11-15/reportsetting.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/backup-2018-11-15/reportsetting.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ReportPlan

ResourceSelection

All content copied from https://docs.aws.amazon.com/.
