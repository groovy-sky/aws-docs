This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataPipeline::Pipeline

The AWS::DataPipeline::Pipeline resource specifies a data pipeline that you can use to
automate the movement and transformation of data.

###### Important

AWS Data Pipeline is no longer available to new customers. Existing customers of AWS Data Pipeline can continue to use the service as normal. [Learn more](https://aws.amazon.com/blogs/big-data/migrate-workloads-from-aws-data-pipeline)

In each pipeline, you define pipeline
objects, such as activities, schedules, data nodes, and resources.

The `AWS::DataPipeline::Pipeline` resource adds tasks, schedules, and
preconditions to the specified pipeline. You can use `PutPipelineDefinition` to
populate a new pipeline.

`PutPipelineDefinition` also validates the configuration as it adds it to the pipeline. Changes to the pipeline are saved unless one
of the following validation errors exist in the pipeline.

- An object is missing a name or identifier field.

- A string or reference field is empty.

- The number of objects in the pipeline exceeds the allowed maximum number of objects.

- The pipeline is in a FINISHED state.

Pipeline object definitions are passed to the [PutPipelineDefinition](https://docs.aws.amazon.com/datapipeline/latest/APIReference/API_PutPipelineDefinition.html) action and returned by the [GetPipelineDefinition](https://docs.aws.amazon.com/datapipeline/latest/APIReference/API_GetPipelineDefinition.html) action.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DataPipeline::Pipeline",
  "Properties" : {
      "Activate" : Boolean,
      "Description" : String,
      "Name" : String,
      "ParameterObjects" : [ ParameterObject, ... ],
      "ParameterValues" : [ ParameterValue, ... ],
      "PipelineObjects" : [ PipelineObject, ... ],
      "PipelineTags" : [ PipelineTag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::DataPipeline::Pipeline
Properties:
  Activate: Boolean
  Description: String
  Name: String
  ParameterObjects:
    - ParameterObject
  ParameterValues:
    - ParameterValue
  PipelineObjects:
    - PipelineObject
  PipelineTags:
    - PipelineTag

```

## Properties

`Activate`

Indicates whether to validate and start the pipeline or stop an active pipeline. By
default, the value is set to `true`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description of the pipeline.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the pipeline.

_Required_: Yes

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\n\t]*`

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ParameterObjects`

The parameter objects used with the pipeline.

_Required_: No

_Type_: Array of [ParameterObject](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-datapipeline-pipeline-parameterobject.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParameterValues`

The parameter values used with the pipeline.

_Required_: No

_Type_: Array of [ParameterValue](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-datapipeline-pipeline-parametervalue.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PipelineObjects`

The objects that define the pipeline. These objects overwrite the existing pipeline definition. Not all objects, fields, and values
can be updated. For information about restrictions, see
[Editing Your Pipeline](https://docs.aws.amazon.com/datapipeline/latest/DeveloperGuide/dp-manage-pipeline-modify-console.html)
in the _AWS Data Pipeline Developer Guide_.

_Required_: No

_Type_: Array of [PipelineObject](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-datapipeline-pipeline-pipelineobject.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PipelineTags`

A list of arbitrary tags (key-value pairs) to associate with the pipeline, which you
can use to control permissions. For more information, see [Controlling Access to\
Pipelines and Resources](https://docs.aws.amazon.com/datapipeline/latest/DeveloperGuide/dp-control-access.html) in the
_AWS Data Pipeline Developer Guide_.

_Required_: No

_Type_: Array of [PipelineTag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-datapipeline-pipeline-pipelinetag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the pipeline ID.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

`PipelineId`

The ID of the pipeline.

## Examples

The following data pipeline backs up data from an Amazon DynamoDB table to an Amazon
S3 bucket. The pipeline uses the `HiveCopyActivity` activity to copy the data,
and runs it once a day. The [roles](https://docs.aws.amazon.com/datapipeline/latest/DeveloperGuide/dp-iam-roles.html) for the
pipeline and the pipeline resource are declared elsewhere in the same template.

#### JSON

```json

"DynamoDBInputS3OutputHive": {
  "Type": "AWS::DataPipeline::Pipeline",
  "Properties": {
    "Name": "DynamoDBInputS3OutputHive",
    "Description": "Pipeline to backup DynamoDB data to S3",
    "Activate": "true",
    "ParameterObjects": [
      {
        "Id": "myDDBReadThroughputRatio",
        "Attributes": [
          {
            "Key": "description",
            "StringValue": "DynamoDB read throughput ratio"
          },
          {
            "Key": "type",
            "StringValue": "Double"
          },
          {
            "Key": "default",
            "StringValue": "0.2"
          }
        ]
      },
      {
        "Id": "myOutputS3Loc",
        "Attributes": [
          {
            "Key": "description",
            "StringValue": "S3 output bucket"
          },
          {
            "Key": "type",
            "StringValue": "AWS::S3::ObjectKey"
          },
          {
            "Key": "default",
            "StringValue": { "Fn::Join" : [ "", [ "s3://", { "Ref": "S3OutputLoc" } ] ] }
          }
        ]
      },
      {
        "Id": "myDDBTableName",
        "Attributes": [
          {
            "Key": "description",
            "StringValue": "DynamoDB Table Name "
          },
          {
            "Key": "type",
            "StringValue": "String"
          }
        ]
      }
    ],
    "ParameterValues": [
      {
        "Id": "myDDBTableName",
        "StringValue": { "Ref": "TableName" }
      }
    ],
    "PipelineObjects": [
      {
        "Id": "S3BackupLocation",
        "Name": "Copy data to this S3 location",
        "Fields": [
          {
            "Key": "type",
            "StringValue": "S3DataNode"
          },
          {
            "Key": "dataFormat",
            "RefValue": "DDBExportFormat"
          },
          {
            "Key": "directoryPath",
            "StringValue": "#{myOutputS3Loc}/#{format(@scheduledStartTime, 'YYYY-MM-dd-HH-mm-ss')}"
          }
        ]
      },
      {
        "Id": "DDBSourceTable",
        "Name": "DDBSourceTable",
        "Fields": [
          {
            "Key": "tableName",
            "StringValue": "#{myDDBTableName}"
          },
          {
            "Key": "type",
            "StringValue": "DynamoDBDataNode"
          },
          {
            "Key": "dataFormat",
            "RefValue": "DDBExportFormat"
          },
          {
            "Key": "readThroughputPercent",
            "StringValue": "#{myDDBReadThroughputRatio}"
          }
        ]
      },
      {
        "Id": "DDBExportFormat",
        "Name": "DDBExportFormat",
        "Fields": [
          {
            "Key": "type",
            "StringValue": "DynamoDBExportDataFormat"
          }
        ]
      },
      {
        "Id": "TableBackupActivity",
        "Name": "TableBackupActivity",
        "Fields": [
          {
            "Key": "resizeClusterBeforeRunning",
            "StringValue": "true"
          },
          {
            "Key": "type",
            "StringValue": "HiveCopyActivity"
          },
          {
            "Key": "input",
            "RefValue": "DDBSourceTable"
          },
          {
            "Key": "runsOn",
            "RefValue": "EmrClusterForBackup"
          },
          {
            "Key": "output",
            "RefValue": "S3BackupLocation"
          }
        ]
      },
      {
        "Id": "DefaultSchedule",
        "Name": "RunOnce",
        "Fields": [
          {
            "Key": "occurrences",
            "StringValue": "1"
          },
          {
            "Key": "startAt",
            "StringValue": "FIRST_ACTIVATION_DATE_TIME"
          },
          {
            "Key": "type",
            "StringValue": "Schedule"
          },
          {
            "Key": "period",
            "StringValue": "1 Day"
          }
        ]
      },
      {
        "Id": "Default",
        "Name": "Default",
        "Fields": [
          {
            "Key": "type",
            "StringValue": "Default"
          },
          {
            "Key": "scheduleType",
            "StringValue": "cron"
          },
          {
            "Key": "failureAndRerunMode",
            "StringValue": "CASCADE"
          },
          {
            "Key": "role",
            "StringValue": "DataPipelineDefaultRole"
          },
          {
            "Key": "resourceRole",
            "StringValue": "DataPipelineDefaultResourceRole"
          },
          {
            "Key": "schedule",
            "RefValue": "DefaultSchedule"
          }
        ]
      },
      {
        "Id": "EmrClusterForBackup",
        "Name": "EmrClusterForBackup",
        "Fields": [
          {
            "Key": "terminateAfter",
            "StringValue": "2 Hours"
          },
          {
            "Key": "amiVersion",
            "StringValue": "3.3.2"
          },
          {
            "Key": "masterInstanceType",
            "StringValue": "m1.medium"
          },
          {
            "Key": "coreInstanceType",
            "StringValue": "m1.medium"
          },
          {
            "Key": "coreInstanceCount",
            "StringValue": "1"
          },
          {
            "Key": "type",
            "StringValue": "EmrCluster"
          }
        ]
      }
    ]
  }
}
```

#### YAML

```yaml

DynamoDBInputS3OutputHive:
  Type: AWS::DataPipeline::Pipeline
  Properties:
    Name: DynamoDBInputS3OutputHive
    Description: "Pipeline to backup DynamoDB data to S3"
    Activate: true
    ParameterObjects:
      -
        Id: "myDDBReadThroughputRatio"
        Attributes:
          -
            Key: "description"
            StringValue: "DynamoDB read throughput ratio"
          -
            Key: "type"
            StringValue: "Double"
          -
            Key: "default"
            StringValue: "0.2"
      -
        Id: "myOutputS3Loc"
        Attributes:
          -
            Key: "description"
            StringValue: "S3 output bucket"
          -
            Key: "type"
            StringValue: "AWS::S3::ObjectKey"
          -
            Key: "default"
            StringValue:
              Fn::Join:
                - ""
                -
                  - "s3://"
                  -
                    Ref: "S3OutputLoc"
      -
        Id: "myDDBTableName"
        Attributes:
          -
            Key: "description"
            StringValue: "DynamoDB Table Name "
          -
            Key: "type"
            StringValue: "String"
    ParameterValues:
      -
        Id: "myDDBTableName"
        StringValue:
          Ref: "TableName"
    PipelineObjects:
      -
        Id: "S3BackupLocation"
        Name: "Copy data to this S3 location"
        Fields:
          -
            Key: "type"
            StringValue: "S3DataNode"
          -
            Key: "dataFormat"
            RefValue: "DDBExportFormat"
          -
            Key: "directoryPath"
            StringValue: "#{myOutputS3Loc}/#{format(@scheduledStartTime, 'YYYY-MM-dd-HH-mm-ss')}"
      -
        Id: "DDBSourceTable"
        Name: "DDBSourceTable"
        Fields:
          -
            Key: "tableName"
            StringValue: "#{myDDBTableName}"
          -
            Key: "type"
            StringValue: "DynamoDBDataNode"
          -
            Key: "dataFormat"
            RefValue: "DDBExportFormat"
          -
            Key: "readThroughputPercent"
            StringValue: "#{myDDBReadThroughputRatio}"
      -
        Id: "DDBExportFormat"
        Name: "DDBExportFormat"
        Fields:
          -
            Key: "type"
            StringValue: "DynamoDBExportDataFormat"
      -
        Id: "TableBackupActivity"
        Name: "TableBackupActivity"
        Fields:
          -
            Key: "resizeClusterBeforeRunning"
            StringValue: "true"
          -
            Key: "type"
            StringValue: "HiveCopyActivity"
          -
            Key: "input"
            RefValue: "DDBSourceTable"
          -
            Key: "runsOn"
            RefValue: "EmrClusterForBackup"
          -
            Key: "output"
            RefValue: "S3BackupLocation"
      -
        Id: "DefaultSchedule"
        Name: "RunOnce"
        Fields:
          -
            Key: "occurrences"
            StringValue: "1"
          -
            Key: "startAt"
            StringValue: "FIRST_ACTIVATION_DATE_TIME"
          -
            Key: "type"
            StringValue: "Schedule"
          -
            Key: "period"
            StringValue: "1 Day"
      -
        Id: "Default"
        Name: "Default"
        Fields:
          -
            Key: "type"
            StringValue: "Default"
          -
            Key: "scheduleType"
            StringValue: "cron"
          -
            Key: "failureAndRerunMode"
            StringValue: "CASCADE"
          -
            Key: "role"
            StringValue: "DataPipelineDefaultRole"
          -
            Key: "resourceRole"
            StringValue: "DataPipelineDefaultResourceRole"
          -
            Key: "schedule"
            RefValue: "DefaultSchedule"
      -
        Id: "EmrClusterForBackup"
        Name: "EmrClusterForBackup"
        Fields:
          -
            Key: "terminateAfter"
            StringValue: "2 Hours"
          -
            Key: "amiVersion"
            StringValue: "3.3.2"
          -
            Key: "masterInstanceType"
            StringValue: "m1.medium"
          -
            Key: "coreInstanceType"
            StringValue: "m1.medium"
          -
            Key: "coreInstanceCount"
            StringValue: "1"
          -
            Key: "type"
            StringValue: "EmrCluster"
```

## See also

- [Pipeline Object\
Reference](https://docs.aws.amazon.com/datapipeline/latest/DeveloperGuide/dp-pipeline-objects.html) in the _AWS Data Pipeline Developer_
_Guide_.

- [PutPipelineDefinition](https://docs.aws.amazon.com/datapipeline/latest/APIReference/API_PutPipelineDefinition.html) in the
_AWS Data Pipeline API Reference_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS Data Pipeline

Field
