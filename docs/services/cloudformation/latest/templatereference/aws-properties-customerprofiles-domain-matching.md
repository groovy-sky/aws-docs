This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CustomerProfiles::Domain Matching

The process of matching duplicate profiles. If `Matching = true`, Amazon Connect Customer Profiles starts a weekly batch process called _Identity_
_Resolution Job_. If you do not specify a date and time for the
_Identity Resolution Job_ to run, by default it runs every
Saturday at 12AM UTC to detect duplicate profiles in your domains. After the
_Identity Resolution Job_ completes, use the
`GetMatches` API to return and review the results. Or, if you have
configured `ExportingConfig` in the `MatchingRequest`, you can
download the results from S3.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AutoMerging" : AutoMerging,
  "Enabled" : Boolean,
  "ExportingConfig" : ExportingConfig,
  "JobSchedule" : JobSchedule
}

```

### YAML

```yaml

  AutoMerging:
    AutoMerging
  Enabled: Boolean
  ExportingConfig:
    ExportingConfig
  JobSchedule:
    JobSchedule

```

## Properties

`AutoMerging`

Configuration information about the auto-merging process.

_Required_: No

_Type_: [AutoMerging](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-customerprofiles-domain-automerging.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enabled`

The flag that enables the matching process of duplicate profiles.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExportingConfig`

The S3 location where Identity Resolution Jobs write result files.

_Required_: No

_Type_: [ExportingConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-customerprofiles-domain-exportingconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JobSchedule`

The day and time when do you want to start the Identity Resolution Job every
week.

_Required_: No

_Type_: [JobSchedule](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-customerprofiles-domain-jobschedule.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

JobSchedule

MatchingRule
