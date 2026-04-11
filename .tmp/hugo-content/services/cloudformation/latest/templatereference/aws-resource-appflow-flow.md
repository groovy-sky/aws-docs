This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::Flow

The `AWS::AppFlow::Flow` resource is an Amazon AppFlow resource type that
specifies a new flow.

###### Note

If you want to use CloudFormation to create a connector profile for connectors
that implement OAuth (such as Salesforce, Slack, Zendesk, and Google Analytics), you must
fetch the access and refresh tokens. You can do this by implementing your own UI for OAuth,
or by retrieving the tokens from elsewhere. Alternatively, you can use the Amazon AppFlow
console to create the connector profile, and then use that connector profile in the flow
creation CloudFormation template.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AppFlow::Flow",
  "Properties" : {
      "Description" : String,
      "DestinationFlowConfigList" : [ DestinationFlowConfig, ... ],
      "FlowName" : String,
      "FlowStatus" : String,
      "KMSArn" : String,
      "MetadataCatalogConfig" : MetadataCatalogConfig,
      "SourceFlowConfig" : SourceFlowConfig,
      "Tags" : [ Tag, ... ],
      "Tasks" : [ Task, ... ],
      "TriggerConfig" : TriggerConfig
    }
}

```

### YAML

```yaml

Type: AWS::AppFlow::Flow
Properties:
  Description: String
  DestinationFlowConfigList:
    - DestinationFlowConfig
  FlowName: String
  FlowStatus: String
  KMSArn: String
  MetadataCatalogConfig:
    MetadataCatalogConfig
  SourceFlowConfig:
    SourceFlowConfig
  Tags:
    - Tag
  Tasks:
    - Task
  TriggerConfig:
    TriggerConfig

```

## Properties

`Description`

A user-entered description of the flow.

_Required_: No

_Type_: String

_Pattern_: `[\w!@#\-.?,\s]*`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DestinationFlowConfigList`

The configuration that controls how Amazon AppFlow places data in the destination
connector.

_Required_: Yes

_Type_: Array of [DestinationFlowConfig](aws-properties-appflow-flow-destinationflowconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FlowName`

The specified name of the flow. Spaces are not allowed. Use underscores (\_) or hyphens
(-) only.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9][\w!@#.-]+`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FlowStatus`

Sets the status of the flow. You can specify one of the following values:

Active

The flow runs based on the trigger settings that you defined. Active scheduled flows run
as scheduled, and active event-triggered flows run when the specified change event occurs.
However, active on-demand flows run only when you manually start them by using Amazon
AppFlow.

Suspended

You can use this option to deactivate an active flow. Scheduled and event-triggered
flows will cease to run until you reactive them. This value only affects scheduled and
event-triggered flows. It has no effect for on-demand flows.

If you omit the FlowStatus parameter, Amazon AppFlow creates the flow with a default
status. The default status for on-demand flows is Active. The default status for scheduled and
event-triggered flows is Draft, which means they’re not yet active.

_Required_: No

_Type_: String

_Allowed values_: `Active | Suspended | Draft`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KMSArn`

The ARN (Amazon Resource Name) of the Key Management Service (KMS) key you provide for
encryption. This is required if you do not want to use the Amazon AppFlow-managed KMS
key. If you don't provide anything here, Amazon AppFlow uses the Amazon AppFlow-managed KMS key.

_Required_: No

_Type_: String

_Pattern_: `arn:aws:kms:.*:[0-9]+:.*`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MetadataCatalogConfig`

Specifies the configuration that Amazon AppFlow uses when it catalogs your data. When
Amazon AppFlow catalogs your data, it stores metadata in a data catalog.

_Required_: No

_Type_: [MetadataCatalogConfig](aws-properties-appflow-flow-metadatacatalogconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceFlowConfig`

Contains information about the configuration of the source connector used in the flow.

_Required_: Yes

_Type_: [SourceFlowConfig](aws-properties-appflow-flow-sourceflowconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags used to organize, track, or control access for your flow.

_Required_: No

_Type_: Array of [Tag](aws-properties-appflow-flow-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tasks`

A list of tasks that Amazon AppFlow performs while transferring the data in the flow
run.

_Required_: Yes

_Type_: Array of [Task](aws-properties-appflow-flow-task.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TriggerConfig`

The trigger settings that determine how and when Amazon AppFlow runs the specified
flow.

_Required_: Yes

_Type_: [TriggerConfig](aws-properties-appflow-flow-triggerconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the flow name. For example:

`{ "Ref": "myFlowName" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`FlowArn`

The flow's Amazon Resource Name (ARN).

## Examples

### Test flow for CloudFormation from Salesforce to Amazon S3

The following example shows a test event flow for CloudFormation using Salesforce as
the source and Amazon S3 as the destination.

#### JSON

```json

{
    "AWSTemplateFormatVersion":"2010-09-09",
    "Resources": {
        "TestFlow": {
            "Type": "AWS::AppFlow::Flow",
            "Properties": {
                "flowName": "MyEventFlow",
                "description": "Test event flow for CloudFormation from salesforce to s3",
                "triggerConfig": {
                    "triggerType": "Event"
                 },
                "sourceFlowConfig": {
                    "connectorType": "Salesforce",
                    "connectorProfileName": "TestConnectorProfile",
                    "sourceConnectorProperties": {
                        "Salesforce": {
                            "object": "Account",
                            "enableDynamicFieldUpdate": false,
                            "includeDeletedRecords": true
                        }
                    }
                },
                "destinationFlowConfigList": [
                    {
                        "connectorType": "S3",
                        "destinationConnectorProperties": {
                            "S3": {
                                "bucketName": "TestOutputBucket",
                                "s3OutputFormatConfig": {
                                    "fileType": "JSON",
                                    "aggregationConfig": {
                                        "aggregationType": "None"
                                    }
                                }
                            }
                        }
                    }
                ],
                "tasks": [
                    {
                        "taskType": "Filter",
                        "sourceFields": [
                            "Id",
                            "Name"
                        ],
                        "connectorOperator": {
"Salesforce": "PROJECTION"
                            }
                     },
                    {
                        "taskType": "Map",
                        "sourceFields": [
                            "Id"
                        ],
                        "taskProperties": [
                            {
                                "Key": "SOURCE_DATA_TYPE",
                                "Value": "id"
                            },
                             {
                                "Key": "DESTINATION_DATA_TYPE",
                                "Value": "id"
                            }
                        ],
                        "destinationField": "Id",
                        "connectorOperator": {
                            "Salesforce": "NO_OP"
                        }
                    },
                    {
                        "taskType": "Map",
                        "sourceFields": [
                            "Name"
                        ],
                        "taskProperties": [
                            {
                                "Key": "SOURCE_DATA_TYPE",
                                "Value": "string"
                            },
                            {
                                "Key": "DESTINATION_DATA_TYPE",
                                "Value": "string"
                            }
                        ],
                        "destinationField": "Name",
                        "connectorOperator": {
                            "Salesforce": "NO_OP"
                        }
                    }
                ],
                "tags": [
                    {
                        "Key": "testKey",
                        "Value": "testValue"
                    }
                ]
            }
        }
    }
}

```

#### YAML

```yaml

AWSTemplateFormatVersion: '2010-09-09'
Resources:
  TestFlow:
    Type: AWS::AppFlow::Flow
    Properties:
      FlowName: MyEventFlow
      Description: Test flow for CloudFormation from salesforce to s3
      TriggerConfig:
        TriggerType: Event
      SourceFlowConfig:
        ConnectorType: Salesforce
        ConnectorProfileName: TestConnectorProfile
        SourceConnectorProperties:
          Salesforce:
            Object: Account
            EnableDynamicFieldUpdate: false
            IncludeDeletedRecords: true
      DestinationFlowConfigList:
        - ConnectorType: S3
          DestinationConnectorProperties:
            S3:
              BucketName: TestOutputBucket
              S3OutputFormatConfig:
                FileType: JSON
                AggregationConfig:
                  AggregationType: None
      Tasks:
        - TaskType: Filter
          ConnectorOperator:
            Salesforce: PROJECTION
          SourceFields:
            - Id
        - TaskType: Map
          SourceFields:
            - Id
          TaskProperties:
            - Key: SOURCE_DATA_TYPE
              Value: id
            - Key: DESTINATION_DATA_TYPE
              Value: id
          DestinationField: Id
          ConnectorOperator:
            Salesforce: NO_OP
        - TaskType: Map
          SourceFields:
            - Name
          TaskProperties:
            - Key: SOURCE_DATA_TYPE
              Value: string
            - Key: DESTINATION_DATA_TYPE
              Value: string
          DestinationField: Name
          ConnectorOperator:
            Salesforce: NO_OP

```

## See also

- [CreateFlow](../../../../reference/appflow/1-0/apireference/api-createflow.md) in the _Amazon AppFlow API Reference_.

- [DescribeFlow](../../../../reference/appflow/1-0/apireference/api-describeflow.md) in the _Amazon AppFlow API Reference_.

- [DeleteFlow](../../../../reference/appflow/1-0/apireference/api-deleteflow.md) in the _Amazon AppFlow API Reference_.

- [UpdateFlow](../../../../reference/appflow/1-0/apireference/api-updateflow.md) in the _Amazon AppFlow API Reference_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ZendeskConnectorProfileProperties

AggregationConfig

All content copied from https://docs.aws.amazon.com/.
