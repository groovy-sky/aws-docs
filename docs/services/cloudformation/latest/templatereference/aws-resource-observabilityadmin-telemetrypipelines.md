This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ObservabilityAdmin::TelemetryPipelines

Creates a telemetry pipeline for processing and transforming telemetry data. The pipeline
defines how data flows from sources through processors to destinations, enabling data
transformation and delivering capabilities.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ObservabilityAdmin::TelemetryPipelines",
  "Properties" : {
      "Configuration" : TelemetryPipelineConfiguration,
      "Name" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::ObservabilityAdmin::TelemetryPipelines
Properties:
  Configuration:
    TelemetryPipelineConfiguration
  Name: String
  Tags:
    - Tag

```

## Properties

`Configuration`

The configuration that defines how the telemetry pipeline processes data, including
sources, processors, and destinations. For more information, see the
[Amazon CloudWatch User Guide](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Creating-pipelines.html).

_Required_: Yes

_Type_: [TelemetryPipelineConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-observabilityadmin-telemetrypipelines-telemetrypipelineconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the telemetry pipeline to create. The name must be unique within your
account.

_Required_: No

_Type_: String

_Pattern_: `[a-z][a-z0-9\-]+`

_Minimum_: `3`

_Maximum_: `28`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The key-value pairs to associate with the telemetry pipeline resource for categorization
and management purposes.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-observabilityadmin-telemetrypipelines-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

The Amazon Resource Name (ARN) of the created telemetry pipeline.

### Fn::GetAtt

`Arn`

The Amazon Resource Name (ARN) of the created telemetry pipeline.

`PipelineIdentifier`

The Amazon Resource Name (ARN) of the telemetry pipeline.

`Status`

The current status of the telemetry pipeline.

## Examples

- [Using CloudWatch as a pipeline source](#aws-resource-observabilityadmin-telemetrypipelines--examples--Using_CloudWatch_as_a_pipeline_source)

- [Using Amazon S3 as a pipeline source](#aws-resource-observabilityadmin-telemetrypipelines--examples--Using_Amazon_S3_as_a_pipeline_source)

### Using CloudWatch as a pipeline source

The following is an example of a `Body` property value for the `Configuration` object.

#### JSON

```json

{
  "Type": "AWS::ObservabilityAdmin::TelemetryPipelines",
  "Properties": {
    "Configuration": {
      "Body": "pipeline:\n  source:\n    cloudwatch_logs:\n      log_event_metadata:\n        data_source_name: \"my_data_source\"\n        data_source_type: \"default\"\n      aws:\n        sts_role_arn: \"arn:aws:iam::123456789012:role/MyPipelineAccessRole\"\n  processor:\n    - parse_json: {}\n  sink:\n    - cloudwatch_logs:\n        log_group: \"@original\""
    }
  }
}
```

#### YAML

```yaml

Type: AWS::ObservabilityAdmin::TelemetryPipelines
Properties:
  Configuration:
    Body: |
      pipeline:
        source:
          cloudwatch_logs:
            log_event_metadata:
              data_source_name: "my_data_source"
              data_source_type: "default"
            aws:
              sts_role_arn: "arn:aws:iam::123456789012:role/MyPipelineAccessRole"
        processor:
          - parse_json: {}
        sink:
          - cloudwatch_logs:
              log_group: "@original"
```

### Using Amazon S3 as a pipeline source

The following is an example of a `Body` property value for the `Configuration` object.

#### JSON

```json

{
  "Type": "AWS::ObservabilityAdmin::TelemetryPipelines",
  "Properties": {
    "Configuration": {
      "Body": "pipeline:\n  source:\n    s3:\n      sqs:\n        visibility_timeout: \"PT60S\"\n        visibility_duplication_protection: true\n        maximum_messages: 10\n        queue_url: \"https://sqs.us-east-1.amazonaws.com/123456789012/my-sqs-queue\"\n      notification_type: \"sqs\"\n      codec:\n        ndjson: {}\n      aws:\n        region: \"us-east-1\"\n        sts_role_arn: \"arn:aws:iam::123456789012:role/MyAccessRole\"\n      data_source_name: \"crowdstrike_falcon\"\n  processor:\n    - ocsf:\n        version: \"1.5\"\n        mapping_version: \"1.5.0\"\n        schema:\n          crowdstrike_falcon:\n  sink:\n    - cloudwatch_logs:\n        log_group: \"my-log-group\""
    }
  }
}
```

#### YAML

```yaml

Type: AWS::ObservabilityAdmin::TelemetryPipelines
Properties:
  Configuration:
    Body: |
      pipeline:
        source:
          s3:
            sqs:
              visibility_timeout: "PT60S"
              visibility_duplication_protection: true
              maximum_messages: 10
              queue_url: "https://sqs.us-east-1.amazonaws.com/123456789012/my-sqs-queue"
            notification_type: "sqs"
            codec:
              ndjson: {}
            aws:
              region: "us-east-1"
              sts_role_arn: "arn:aws:iam::123456789012:role/MyAccessRole"
            data_source_name: "crowdstrike_falcon"
        processor:
          - ocsf:
              version: "1.5"
              mapping_version: "1.5.0"
              schema:
                crowdstrike_falcon:
        sink:
          - cloudwatch_logs:
              log_group: "my-log-group"
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::ObservabilityAdmin::TelemetryEnrichment

Tag
