This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeBuild::Project ProjectBuildBatchConfig

Contains configuration information about a batch build project.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BatchReportMode" : String,
  "CombineArtifacts" : Boolean,
  "Restrictions" : BatchRestrictions,
  "ServiceRole" : String,
  "TimeoutInMins" : Integer
}

```

### YAML

```yaml

  BatchReportMode: String
  CombineArtifacts: Boolean
  Restrictions:
    BatchRestrictions
  ServiceRole: String
  TimeoutInMins: Integer

```

## Properties

`BatchReportMode`

Specifies how build status reports are sent to the source provider for the batch build. This property is only used
when the source provider for your project is Bitbucket, GitHub, or GitHub Enterprise,
and your project is configured to report build statuses to the source provider.

REPORT\_AGGREGATED\_BATCH

(Default) Aggregate all of the build statuses into a single status report.

REPORT\_INDIVIDUAL\_BUILDS

Send a separate status report for each individual build.

_Required_: No

_Type_: String

_Allowed values_: `REPORT_INDIVIDUAL_BUILDS | REPORT_AGGREGATED_BATCH`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CombineArtifacts`

Specifies if the build artifacts for the batch build should be combined into a single
artifact location.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Restrictions`

A `BatchRestrictions` object that specifies the restrictions for the batch
build.

_Required_: No

_Type_: [BatchRestrictions](aws-properties-codebuild-project-batchrestrictions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceRole`

Specifies the service role ARN for the batch build project.

_Required_: No

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeoutInMins`

Specifies the maximum amount of time, in minutes, that the batch build must be completed in.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LogsConfig

ProjectCache

All content copied from https://docs.aws.amazon.com/.
