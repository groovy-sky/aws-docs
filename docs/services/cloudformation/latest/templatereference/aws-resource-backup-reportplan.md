This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Backup::ReportPlan

Creates a report plan. A report plan is a document that contains information about the
contents of the report and where AWS Backup will deliver it.

If you call `CreateReportPlan` with a plan that already exists, you receive
an `AlreadyExistsException` exception.

For a sample CloudFormation template, see the [AWS Backup Developer Guide](../../../aws-backup/latest/devguide/assigning-resources.md#assigning-resources-cfn).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Backup::ReportPlan",
  "Properties" : {
      "ReportDeliveryChannel" : ReportDeliveryChannel,
      "ReportPlanDescription" : String,
      "ReportPlanName" : String,
      "ReportPlanTags" : [ Tag, ... ],
      "ReportSetting" : ReportSetting
    }
}

```

### YAML

```yaml

Type: AWS::Backup::ReportPlan
Properties:
  ReportDeliveryChannel:
    ReportDeliveryChannel
  ReportPlanDescription: String
  ReportPlanName: String
  ReportPlanTags:
    - Tag
  ReportSetting:
    ReportSetting

```

## Properties

`ReportDeliveryChannel`

Contains information about where and how to deliver your reports, specifically your
Amazon S3 bucket name, S3 key prefix, and the formats of your reports.

_Required_: Yes

_Type_: [ReportDeliveryChannel](aws-properties-backup-reportplan-reportdeliverychannel.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReportPlanDescription`

An optional description of the report plan with a maximum 1,024 characters.

_Required_: No

_Type_: String

_Pattern_: `.*\S.*`

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReportPlanName`

The unique name of the report plan. This name is between 1 and 256 characters starting
with a letter, and consisting of letters (a-z, A-Z), numbers (0-9), and underscores
(\_).

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z][_a-zA-Z0-9]*`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ReportPlanTags`

The tags to assign to your report plan.

_Required_: No

_Type_: Array of [Tag](aws-properties-backup-reportplan-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReportSetting`

Identifies the report template for the report. Reports are built using a report
template. The report templates are:

`RESOURCE_COMPLIANCE_REPORT | CONTROL_COMPLIANCE_REPORT | BACKUP_JOB_REPORT |
            COPY_JOB_REPORT | RESTORE_JOB_REPORT`

If the report template is `RESOURCE_COMPLIANCE_REPORT` or
`CONTROL_COMPLIANCE_REPORT`, this API resource also describes the report
coverage by AWS Regions and frameworks.

_Required_: Yes

_Type_: [ReportSetting](aws-properties-backup-reportplan-reportsetting.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`ReportPlanArn`

The Amazon Resource Name (ARN) of your report plan.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NotificationObjectType

ReportDeliveryChannel

All content copied from https://docs.aws.amazon.com/.
