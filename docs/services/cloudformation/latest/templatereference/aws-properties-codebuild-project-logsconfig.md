This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeBuild::Project LogsConfig

`LogsConfig` is a property of the [AWS CodeBuild Project](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html) resource
that specifies information about logs for a build project. These can be logs in Amazon CloudWatch Logs, built in a
specified S3 bucket, or both.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CloudWatchLogs" : CloudWatchLogsConfig,
  "S3Logs" : S3LogsConfig
}

```

### YAML

```yaml

  CloudWatchLogs:
    CloudWatchLogsConfig
  S3Logs:
    S3LogsConfig

```

## Properties

`CloudWatchLogs`

Information about CloudWatch Logs for a build project. CloudWatch Logs are enabled by default.

_Required_: No

_Type_: [CloudWatchLogsConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-codebuild-project-cloudwatchlogsconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3Logs`

Information about logs built to an S3 bucket for a build project. S3 logs are not
enabled by default.

_Required_: No

_Type_: [S3LogsConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-codebuild-project-s3logsconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [LogsConfig](https://docs.aws.amazon.com/codebuild/latest/APIReference/API_LogsConfig.html) in the _AWS CodeBuild API Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

GitSubmodulesConfig

ProjectBuildBatchConfig
