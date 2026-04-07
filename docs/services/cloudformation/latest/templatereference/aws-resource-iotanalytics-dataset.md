This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTAnalytics::Dataset

The AWS::IoTAnalytics::Dataset resource stores data retrieved from a data store by applying a
`queryAction` (an SQL query) or a `containerAction` (executing a containerized application).
The data set can be populated manually by calling `CreateDatasetContent` or automatically according
to a `trigger` you specify. For more information, see
[How to Use AWS IoT Analytics](https://docs.aws.amazon.com/iotanalytics/latest/userguide/welcome.html#aws-iot-analytics-how) in the _AWS IoT Analytics User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoTAnalytics::Dataset",
  "Properties" : {
      "Actions" : [ Action, ... ],
      "ContentDeliveryRules" : [ DatasetContentDeliveryRule, ... ],
      "DatasetName" : String,
      "LateDataRules" : [ LateDataRule, ... ],
      "RetentionPeriod" : RetentionPeriod,
      "Tags" : [ Tag, ... ],
      "Triggers" : [ Trigger, ... ],
      "VersioningConfiguration" : VersioningConfiguration
    }
}

```

### YAML

```yaml

Type: AWS::IoTAnalytics::Dataset
Properties:
  Actions:
    - Action
  ContentDeliveryRules:
    - DatasetContentDeliveryRule
  DatasetName: String
  LateDataRules:
    - LateDataRule
  RetentionPeriod:
    RetentionPeriod
  Tags:
    - Tag
  Triggers:
    - Trigger
  VersioningConfiguration:
    VersioningConfiguration

```

## Properties

`Actions`

The `DatasetAction` objects that automatically create the dataset
contents.

_Required_: Yes

_Type_: Array of [Action](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotanalytics-dataset-action.html)

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ContentDeliveryRules`

When dataset contents are created they are delivered to destinations specified
here.

_Required_: No

_Type_: Array of [DatasetContentDeliveryRule](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotanalytics-dataset-datasetcontentdeliveryrule.html)

_Minimum_: `0`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DatasetName`

The name of the dataset.

_Required_: No

_Type_: String

_Pattern_: `(^(?!_{2}))(^[a-zA-Z0-9_]+$)`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LateDataRules`

A list of data rules that send notifications to CloudWatch, when data arrives late. To specify `lateDataRules`, the dataset must use a [DeltaTimer](https://docs.aws.amazon.com/iotanalytics/latest/APIReference/API_DeltaTime.html) filter.

_Required_: No

_Type_: Array of [LateDataRule](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotanalytics-dataset-latedatarule.html)

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RetentionPeriod`

Optional. How long, in days, message data is kept for the dataset.

_Required_: No

_Type_: [RetentionPeriod](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotanalytics-dataset-retentionperiod.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Metadata which can be used to manage the data set.

For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotanalytics-dataset-tag.html)

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Triggers`

The `DatasetTrigger` objects that specify when the dataset is automatically
updated.

_Required_: No

_Type_: Array of [Trigger](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotanalytics-dataset-trigger.html)

_Minimum_: `0`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VersioningConfiguration`

Optional. How many versions of dataset contents are kept. If not specified or set to null,
only the latest version plus the latest succeeded version (if they are different) are kept for
the time period specified by the `retentionPeriod` parameter. For more information,
see [Keeping Multiple Versions of AWS IoT Analytics datasets](https://docs.aws.amazon.com/iotanalytics/latest/userguide/getting-started.html#aws-iot-analytics-dataset-versions) in the
_AWS IoT Analytics User Guide_.

_Required_: No

_Type_: [VersioningConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotanalytics-dataset-versioningconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

- [Simple SQL Dataset](#aws-resource-iotanalytics-dataset--examples--Simple_SQL_Dataset)

- [Complex SQL Dataset](#aws-resource-iotanalytics-dataset--examples--Complex_SQL_Dataset)

- [Simple Container Dataset](#aws-resource-iotanalytics-dataset--examples--Simple_Container_Dataset)

- [Complex Container Dataset](#aws-resource-iotanalytics-dataset--examples--Complex_Container_Dataset)

### Simple SQL Dataset

The following example creates a simple SQL dataset.

#### JSON

```json

{
    "Description": "Create a simple SQL Dataset",
    "Resources": {
        "Dataset": {
            "Type": "AWS::IoTAnalytics::Dataset",
            "Properties": {
                "DatasetName": "SimpleSQLDataset",
                "Actions": [
                    {
                        "ActionName": "SqlAction",
                        "QueryAction": {
                            "SqlQuery": "select * from Datastore"
                        }
                    }
                ],
                "Triggers": [
                    {
                        "Schedule": {
                            "ScheduleExpression": "cron(0 12 * * ? *)"
                        }
                    }
                ]
            }
        }
    }
}
```

#### YAML

```yaml

---
Description: "Create a simple SQL Dataset"
Resources:
  Dataset:
    Type: "AWS::IoTAnalytics::Dataset"
    Properties:
      DatasetName: "SimpleSQLDataset"
      Actions:
        -
          ActionName: "SqlAction"
          QueryAction:
            SqlQuery: "select * from Datastore"
      Triggers:
        -
          Schedule:
            ScheduleExpression: "cron(0 12 * * ? *)"
```

### Complex SQL Dataset

The following example creates a complex SQL dataset.

#### JSON

```json

{
    "Description": "Create a complex SQL Dataset",
    "Resources": {
        "Dataset": {
            "Type": "AWS::IoTAnalytics::Dataset",
            "Properties": {
                "DatasetName": "ComplexSQLDataset",
                "Actions": [
                    {
                        "ActionName": "SqlAction",
                        "QueryAction": {
                            "SqlQuery": "select * from Datastore",
                            "Filters": [
                                {
                                    "DeltaTime": {
                                        "OffsetSeconds": 1,
                                        "TimeExpression": "timestamp"
                                    }
                                }
                            ]
                        }
                    }
                ],
                "Triggers": [
                    {
                        "Schedule": {
                            "ScheduleExpression": "cron(0 12 * * ? *)"
                        }
                    }
                ],
                "RetentionPeriod": {
                    "Unlimited": false,
                    "NumberOfDays": 10
                },
                "Tags": [
                    {
                        "Key": "keyname1",
                        "Value": "value1"
                    },
                    {
                        "Key": "keyname2",
                        "Value": "value2"
                    }
                ]
            }
        }
    }
}
```

#### YAML

```yaml

---
Description: "Create a complex SQL Dataset"
Resources:
  Dataset:
    Type: "AWS::IoTAnalytics::Dataset"
    Properties:
      DatasetName: "ComplexSQLDataset"
      Actions:
        -
          ActionName: "SqlAction"
          QueryAction:
            SqlQuery: "select * from Datastore"
            Filters:
              -
                DeltaTime:
                  OffsetSeconds: 1
                  TimeExpression: "timestamp"
      Triggers:
        -
          Schedule:
            ScheduleExpression: "cron(0 12 * * ? *)"
      RetentionPeriod:
        Unlimited: false
        NumberOfDays: 10
      Tags:
        -
          Key: "keyname1"
          Value: "value1"
        -
          Key: "keyname2"
          Value: "value2"
```

### Simple Container Dataset

The following example creates a simple Container dataset.

#### JSON

```json

{
    "Description": "Create a simple container Dataset",
    "Resources": {
        "ContainerDataset": {
            "Type": "AWS::IoTAnalytics::Dataset",
            "Properties": {
                "DatasetName": "SimpleContainerDataset",
                "Actions": [
                    {
                        "ActionName": "ContainerAction",
                        "ContainerAction": {
                            "Image": "<your_Account_Id>.dkr.ecr.us-east-1.amazonaws.com/sampleimage",
                            "ExecutionRoleArn": "arn:aws:iam::<your_Account_Id>:role/ExecutionRole",
                            "ResourceConfiguration": {
                                "ComputeType": "ACU_1",
                                "VolumeSizeInGB": 10
                            },
                            "Variables": [
                                {
                                    "VariableName": "Variable1",
                                    "StringValue": "StringValue"
                                }
                            ]
                        }
                    }
                ],
                "Triggers": [
                    {
                        "Schedule": {
                            "ScheduleExpression": "cron(0 12 * * ? *)"
                        }
                    }
                ]
            }
        }
    }
}
```

#### YAML

```yaml

---
Description: "Create a simple container Dataset"
Resources:
  ContainerDataset:
    Type: "AWS::IoTAnalytics::Dataset"
    Properties:
      DatasetName: "SimpleContainerDataset"
      Actions:
        - ActionName: ContainerAction
          ContainerAction:
            Image: "<your_Account_Id>.dkr.ecr.us-east-1.amazonaws.com/sampleimage"
            ExecutionRoleArn: "arn:aws:iam::<your_Account_Id>:role/ExecutionRole"
            ResourceConfiguration:
              ComputeType: "ACU_1"
              VolumeSizeInGB: 10
            Variables:
              - VariableName: "Variable1"
                StringValue: StringValue
      Triggers:
        -
          Schedule:
            ScheduleExpression: "cron(0 12 * * ? *)"
```

### Complex Container Dataset

The following example creates a complex Container dataset.

#### JSON

```json

{
    "Description": "Create a complex container Dataset",
    "Resources": {
        "TriggeringDataset": {
            "Type": "AWS::IoTAnalytics::Dataset",
            "Properties": {
                "DatasetName": "TriggeringDataset",
                "Actions": [
                    {
                        "ActionName": "SqlAction",
                        "QueryAction": {
                            "SqlQuery": "select * from Datastore"
                        }
                    }
                ]
            }
        },
        "ContainerDataset": {
            "Type": "AWS::IoTAnalytics::Dataset",
            "DependsOn": "TriggeringDataset",
            "Properties": {
                "DatasetName": "ComplexContainerDataset",
                "Actions": [
                    {
                        "ActionName": "ContainerAction",
                        "ContainerAction": {
                            "Image": "<your_Account_Id>.dkr.ecr.us-east-1.amazonaws.com/sampleimage",
                            "ExecutionRoleArn": "arn:aws:iam::<your_Account_Id>:role/ExecutionRole",
                            "ResourceConfiguration": {
                                "ComputeType": "ACU_1",
                                "VolumeSizeInGB": 10
                            },
                            "Variables": [
                                {
                                    "VariableName": "Variable1",
                                    "StringValue": "StringValue"
                                },
                                {
                                    "VariableName": "Variable2",
                                    "DoubleValue": 1
                                },
                                {
                                    "VariableName": "Variable3",
                                    "DatasetContentVersionValue": {
                                        "DatasetName": "BasicDataset"
                                    }
                                },
                                {
                                    "VariableName": "Variable4",
                                    "OutputFileUriValue": {
                                        "FileName": "fileName"
                                    }
                                }
                            ]
                        }
                    }
                ],
                "Triggers": [
                    {
                        "TriggeringDataset": {
                            "DatasetName": "TriggeringDataset"
                        }
                    }
                ]
            }
        }
    }
}
```

#### YAML

```yaml

---
Description: "Create a complex container Dataset"
Resources:
  TriggeringDataset:
    Type: "AWS::IoTAnalytics::Dataset"
    Properties:
      DatasetName: "TriggeringDataset"
      Actions:
        -
          ActionName: "SqlAction"
          QueryAction:
            SqlQuery: "select * from Datastore"

  ContainerDataset:
    Type: "AWS::IoTAnalytics::Dataset"
    DependsOn: TriggeringDataset
    Properties:
      DatasetName: "ComplexContainerDataset"
      Actions:
        -
          ActionName: "ContainerAction"
          ContainerAction:
            Image: "<your_Account_Id>.dkr.ecr.us-east-1.amazonaws.com/sampleimage"
            ExecutionRoleArn: "arn:aws:iam::<your_Account_Id>:role/ExecutionRole"
            ResourceConfiguration:
              ComputeType: "ACU_1"
              VolumeSizeInGB: 10
            Variables:
              -
                VariableName: "Variable1"
                StringValue: "StringValue"
              -
                VariableName: "Variable2"
                DoubleValue: 1
              -
                VariableName: "Variable3"
                DatasetContentVersionValue:
                  DatasetName: "BasicDataset"
              -
                VariableName: "Variable4"
                OutputFileUriValue:
                  FileName: "fileName"
      Triggers:
        -
          TriggeringDataset:
            DatasetName: "TriggeringDataset"
```

## See also

- [How to Use AWS IoT Analytics](https://docs.aws.amazon.com/iotanalytics/latest/userguide/welcome.html#aws-iot-analytics-how)
in the _AWS IoT Analytics User Guide_

- [CreateDataset](https://docs.aws.amazon.com/iotanalytics/latest/APIReference/API_CreateDataset.html) in the
_AWS IoT Analytics API Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Action
