# DescribeFlow

Provides a description of the specified flow.

## Request Syntax

```nohighlight

POST /describe-flow HTTP/1.1
Content-type: application/json

{
   "flowName": "string"
}
```

## URI Request Parameters

The request does not use any URI parameters.

## Request Body

The request accepts the following data in JSON format.

**[flowName](#API_DescribeFlow_RequestSyntax)**

The specified name of the flow. Spaces are not allowed. Use underscores (\_) or hyphens
(-) only.

Type: String

Length Constraints: Maximum length of 256.

Pattern: `[a-zA-Z0-9][\w!@#.-]+`

Required: Yes

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "createdAt": number,
   "createdBy": "string",
   "description": "string",
   "destinationFlowConfigList": [
      {
         "apiVersion": "string",
         "connectorProfileName": "string",
         "connectorType": "string",
         "destinationConnectorProperties": {
            "CustomConnector": {
               "customProperties": {
                  "string" : "string"
               },
               "entityName": "string",
               "errorHandlingConfig": {
                  "bucketName": "string",
                  "bucketPrefix": "string",
                  "failOnFirstDestinationError": boolean
               },
               "idFieldNames": [ "string" ],
               "writeOperationType": "string"
            },
            "CustomerProfiles": {
               "domainName": "string",
               "objectTypeName": "string"
            },
            "EventBridge": {
               "errorHandlingConfig": {
                  "bucketName": "string",
                  "bucketPrefix": "string",
                  "failOnFirstDestinationError": boolean
               },
               "object": "string"
            },
            "Honeycode": {
               "errorHandlingConfig": {
                  "bucketName": "string",
                  "bucketPrefix": "string",
                  "failOnFirstDestinationError": boolean
               },
               "object": "string"
            },
            "LookoutMetrics": {
            },
            "Marketo": {
               "errorHandlingConfig": {
                  "bucketName": "string",
                  "bucketPrefix": "string",
                  "failOnFirstDestinationError": boolean
               },
               "object": "string"
            },
            "Redshift": {
               "bucketPrefix": "string",
               "errorHandlingConfig": {
                  "bucketName": "string",
                  "bucketPrefix": "string",
                  "failOnFirstDestinationError": boolean
               },
               "intermediateBucketName": "string",
               "object": "string"
            },
            "S3": {
               "bucketName": "string",
               "bucketPrefix": "string",
               "s3OutputFormatConfig": {
                  "aggregationConfig": {
                     "aggregationType": "string",
                     "targetFileSize": number
                  },
                  "fileType": "string",
                  "prefixConfig": {
                     "pathPrefixHierarchy": [ "string" ],
                     "prefixFormat": "string",
                     "prefixType": "string"
                  },
                  "preserveSourceDataTyping": boolean
               }
            },
            "Salesforce": {
               "dataTransferApi": "string",
               "errorHandlingConfig": {
                  "bucketName": "string",
                  "bucketPrefix": "string",
                  "failOnFirstDestinationError": boolean
               },
               "idFieldNames": [ "string" ],
               "object": "string",
               "writeOperationType": "string"
            },
            "SAPOData": {
               "errorHandlingConfig": {
                  "bucketName": "string",
                  "bucketPrefix": "string",
                  "failOnFirstDestinationError": boolean
               },
               "idFieldNames": [ "string" ],
               "objectPath": "string",
               "successResponseHandlingConfig": {
                  "bucketName": "string",
                  "bucketPrefix": "string"
               },
               "writeOperationType": "string"
            },
            "Snowflake": {
               "bucketPrefix": "string",
               "errorHandlingConfig": {
                  "bucketName": "string",
                  "bucketPrefix": "string",
                  "failOnFirstDestinationError": boolean
               },
               "intermediateBucketName": "string",
               "object": "string"
            },
            "Upsolver": {
               "bucketName": "string",
               "bucketPrefix": "string",
               "s3OutputFormatConfig": {
                  "aggregationConfig": {
                     "aggregationType": "string",
                     "targetFileSize": number
                  },
                  "fileType": "string",
                  "prefixConfig": {
                     "pathPrefixHierarchy": [ "string" ],
                     "prefixFormat": "string",
                     "prefixType": "string"
                  }
               }
            },
            "Zendesk": {
               "errorHandlingConfig": {
                  "bucketName": "string",
                  "bucketPrefix": "string",
                  "failOnFirstDestinationError": boolean
               },
               "idFieldNames": [ "string" ],
               "object": "string",
               "writeOperationType": "string"
            }
         }
      }
   ],
   "flowArn": "string",
   "flowName": "string",
   "flowStatus": "string",
   "flowStatusMessage": "string",
   "kmsArn": "string",
   "lastRunExecutionDetails": {
      "mostRecentExecutionMessage": "string",
      "mostRecentExecutionStatus": "string",
      "mostRecentExecutionTime": number
   },
   "lastRunMetadataCatalogDetails": [
      {
         "catalogType": "string",
         "partitionRegistrationOutput": {
            "message": "string",
            "result": "string",
            "status": "string"
         },
         "tableName": "string",
         "tableRegistrationOutput": {
            "message": "string",
            "result": "string",
            "status": "string"
         }
      }
   ],
   "lastUpdatedAt": number,
   "lastUpdatedBy": "string",
   "metadataCatalogConfig": {
      "glueDataCatalog": {
         "databaseName": "string",
         "roleArn": "string",
         "tablePrefix": "string"
      }
   },
   "schemaVersion": number,
   "sourceFlowConfig": {
      "apiVersion": "string",
      "connectorProfileName": "string",
      "connectorType": "string",
      "incrementalPullConfig": {
         "datetimeTypeFieldName": "string"
      },
      "sourceConnectorProperties": {
         "Amplitude": {
            "object": "string"
         },
         "CustomConnector": {
            "customProperties": {
               "string" : "string"
            },
            "dataTransferApi": {
               "Name": "string",
               "Type": "string"
            },
            "entityName": "string"
         },
         "Datadog": {
            "object": "string"
         },
         "Dynatrace": {
            "object": "string"
         },
         "GoogleAnalytics": {
            "object": "string"
         },
         "InforNexus": {
            "object": "string"
         },
         "Marketo": {
            "object": "string"
         },
         "Pardot": {
            "object": "string"
         },
         "S3": {
            "bucketName": "string",
            "bucketPrefix": "string",
            "s3InputFormatConfig": {
               "s3InputFileType": "string"
            }
         },
         "Salesforce": {
            "dataTransferApi": "string",
            "enableDynamicFieldUpdate": boolean,
            "includeDeletedRecords": boolean,
            "object": "string"
         },
         "SAPOData": {
            "objectPath": "string",
            "paginationConfig": {
               "maxPageSize": number
            },
            "parallelismConfig": {
               "maxParallelism": number
            }
         },
         "ServiceNow": {
            "object": "string"
         },
         "Singular": {
            "object": "string"
         },
         "Slack": {
            "object": "string"
         },
         "Trendmicro": {
            "object": "string"
         },
         "Veeva": {
            "documentType": "string",
            "includeAllVersions": boolean,
            "includeRenditions": boolean,
            "includeSourceFiles": boolean,
            "object": "string"
         },
         "Zendesk": {
            "object": "string"
         }
      }
   },
   "tags": {
      "string" : "string"
   },
   "tasks": [
      {
         "connectorOperator": {
            "Amplitude": "string",
            "CustomConnector": "string",
            "Datadog": "string",
            "Dynatrace": "string",
            "GoogleAnalytics": "string",
            "InforNexus": "string",
            "Marketo": "string",
            "Pardot": "string",
            "S3": "string",
            "Salesforce": "string",
            "SAPOData": "string",
            "ServiceNow": "string",
            "Singular": "string",
            "Slack": "string",
            "Trendmicro": "string",
            "Veeva": "string",
            "Zendesk": "string"
         },
         "destinationField": "string",
         "sourceFields": [ "string" ],
         "taskProperties": {
            "string" : "string"
         },
         "taskType": "string"
      }
   ],
   "triggerConfig": {
      "triggerProperties": {
         "Scheduled": {
            "dataPullMode": "string",
            "firstExecutionFrom": number,
            "flowErrorDeactivationThreshold": number,
            "scheduleEndTime": number,
            "scheduleExpression": "string",
            "scheduleOffset": number,
            "scheduleStartTime": number,
            "timezone": "string"
         }
      },
      "triggerType": "string"
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[createdAt](#API_DescribeFlow_ResponseSyntax)**

Specifies when the flow was created.

Type: Timestamp

**[createdBy](#API_DescribeFlow_ResponseSyntax)**

The ARN of the user who created the flow.

Type: String

Length Constraints: Maximum length of 256.

Pattern: `\S+`

**[description](#API_DescribeFlow_ResponseSyntax)**

A description of the flow.

Type: String

Length Constraints: Maximum length of 2048.

Pattern: `[\w!@#\-.?,\s]*`

**[destinationFlowConfigList](#API_DescribeFlow_ResponseSyntax)**

The configuration that controls how Amazon AppFlow transfers data to the destination
connector.

Type: Array of [DestinationFlowConfig](https://docs.aws.amazon.com/appflow/1.0/APIReference/API_DestinationFlowConfig.html) objects

**[flowArn](#API_DescribeFlow_ResponseSyntax)**

The flow's Amazon Resource Name (ARN).

Type: String

Length Constraints: Maximum length of 512.

Pattern: `arn:aws:appflow:.*:[0-9]+:.*`

**[flowName](#API_DescribeFlow_ResponseSyntax)**

The specified name of the flow. Spaces are not allowed. Use underscores (\_) or hyphens
(-) only.

Type: String

Length Constraints: Maximum length of 256.

Pattern: `[a-zA-Z0-9][\w!@#.-]+`

**[flowStatus](#API_DescribeFlow_ResponseSyntax)**

Indicates the current status of the flow.

Type: String

Valid Values: `Active | Deprecated | Deleted | Draft | Errored | Suspended`

**[flowStatusMessage](#API_DescribeFlow_ResponseSyntax)**

Contains an error message if the flow status is in a suspended or error state. This
applies only to scheduled or event-triggered flows.

Type: String

Length Constraints: Maximum length of 2048.

Pattern: `[\s\w/!@#+=.-]*`

**[kmsArn](#API_DescribeFlow_ResponseSyntax)**

The ARN (Amazon Resource Name) of the Key Management Service (KMS) key you provide for
encryption. This is required if you do not want to use the Amazon AppFlow-managed KMS
key. If you don't provide anything here, Amazon AppFlow uses the Amazon AppFlow-managed KMS key.

Type: String

Length Constraints: Minimum length of 20. Maximum length of 2048.

Pattern: `arn:aws:kms:.*:[0-9]+:.*`

**[lastRunExecutionDetails](#API_DescribeFlow_ResponseSyntax)**

Describes the details of the most recent flow run.

Type: [ExecutionDetails](https://docs.aws.amazon.com/appflow/1.0/APIReference/API_ExecutionDetails.html) object

**[lastRunMetadataCatalogDetails](#API_DescribeFlow_ResponseSyntax)**

Describes the metadata catalog, metadata table, and data partitions that Amazon AppFlow used for the associated flow run.

Type: Array of [MetadataCatalogDetail](https://docs.aws.amazon.com/appflow/1.0/APIReference/API_MetadataCatalogDetail.html) objects

**[lastUpdatedAt](#API_DescribeFlow_ResponseSyntax)**

Specifies when the flow was last updated.

Type: Timestamp

**[lastUpdatedBy](#API_DescribeFlow_ResponseSyntax)**

Specifies the user name of the account that performed the most recent update.

Type: String

Length Constraints: Maximum length of 256.

Pattern: `\S+`

**[metadataCatalogConfig](#API_DescribeFlow_ResponseSyntax)**

Specifies the configuration that Amazon AppFlow uses when it catalogs the data that's
transferred by the associated flow. When Amazon AppFlow catalogs the data from a flow, it
stores metadata in a data catalog.

Type: [MetadataCatalogConfig](https://docs.aws.amazon.com/appflow/1.0/APIReference/API_MetadataCatalogConfig.html) object

**[schemaVersion](#API_DescribeFlow_ResponseSyntax)**

The version number of your data schema. Amazon AppFlow assigns this version number.
The version number increases by one when you change any of the following settings in your flow
configuration:

- Source-to-destination field mappings

- Field data types

- Partition keys

Type: Long

**[sourceFlowConfig](#API_DescribeFlow_ResponseSyntax)**

The configuration that controls how Amazon AppFlow retrieves data from the source
connector.

Type: [SourceFlowConfig](https://docs.aws.amazon.com/appflow/1.0/APIReference/API_SourceFlowConfig.html) object

**[tags](#API_DescribeFlow_ResponseSyntax)**

The tags used to organize, track, or control access for your flow.

Type: String to string map

Map Entries: Minimum number of 0 items. Maximum number of 50 items.

Key Length Constraints: Minimum length of 1. Maximum length of 128.

Key Pattern: `^(?!aws:)[a-zA-Z+-=._:/]+$`

Value Length Constraints: Maximum length of 256.

Value Pattern: `[\s\w+-=\.:/@]*`

**[tasks](#API_DescribeFlow_ResponseSyntax)**

A list of tasks that Amazon AppFlow performs while transferring the data in the flow
run.

Type: Array of [Task](https://docs.aws.amazon.com/appflow/1.0/APIReference/API_Task.html) objects

**[triggerConfig](#API_DescribeFlow_ResponseSyntax)**

The trigger settings that determine how and when the flow runs.

Type: [TriggerConfig](https://docs.aws.amazon.com/appflow/1.0/APIReference/API_TriggerConfig.html) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](https://docs.aws.amazon.com/appflow/1.0/APIReference/CommonErrors.html).

**InternalServerException**

An internal service error occurred during the processing of your request. Try again
later.

HTTP Status Code: 500

**ResourceNotFoundException**

The resource specified in the request (such as the source or destination connector
profile) is not found.

HTTP Status Code: 404

## Examples

### DescribeFlow example

This example shows a sample request and response for the `DescribeFlow`
API.

#### Sample Request

```json

{
  "flowName": "name"
}
```

#### Sample Response

```json

{
  "createdAt": "2022-02-22T15:31:41.467000-08:00",
  "createdBy": "user",
  "description": null,
  "destinationFlowConfigList": [
    {
      "aggregationConfig": null,
      "connectorProfileName": "connector_profile_name",
      "connectorType": "S3",
      "destinationConnectorProps":
      {
        "Redshift": null,
        "S3":
        {
          "bucketKey": null,
          "bucketName": "salesforceaccountflowbucket"
        },
        "Salesforce": null,
        "Snowflake": null
      }
    }
  ],
  "flowArn": "flow_arn_value",
  "flowName": "test_flow_ondemand_10",
  "flowStatus": "Active",
  "flowStatusMessage": null,
  "kmsArn": null,
  "lastRunExecutionDetails": null,
  "lastUpdatedAt": "2022-02-22T15:31:41.467000-08:00",
  "lastUpdatedBy": "user",
  "sourceFlowConfig":
  {
    "connectorProfileName": "connectorProfileName",
    "connectorType": "Salesforce",
    "sourceConnectorProps":
    {
      "Amplitude": null,
      "Datadog": null,
      "Dynatrace": null,
      "GoogleAnalytics": null,
      "InforNexus": null,
      "Marketo": null,
      "Redshift": null,
      "S3": null,
      "Salesforce":
      {
        "object": "Account"
      },
      "ServiceNow": null,
      "Singular": null,
      "Slack": null,
      "Snowflake": null,
      "Trendmicro": null,
      "Veeva": null,
      "Zendesk": null
    }
  },
  "tags":
  {
    "internalId": "Internal_Id_value",
    "resourceArn": "resource_arn_value"
  },
  "tasks": [
    {
      "connectorOperator": null,
      "destinationField": "Id",
      "operator": "NO_OP",
      "sourceFields": ["Id"],
      "taskProperties":
      {
        "DESTINATION_DATA_TYPE": "id",
        "SOURCE_DATA_TYPE": "id"
      },
      "taskType": "Mapping"
    },
    {
      "connectorOperator": null,
      "destinationField": "Name",
      "operator": "NO_OP",
      "sourceFields": ["Name"],
      "taskProperties":
      {
        "DESTINATION_DATA_TYPE": "string",
        "SOURCE_DATA_TYPE": "string"
      },
      "taskType": "Mapping"
    },
    {
      "connectorOperator": null,
      "destinationField": null,
      "operator": "PROJECTION",
      "sourceFields": ["Id","Name"],
      "taskProperties":
      {},
      "taskType": "Filtering"
    }
  ],
  "triggerConfig":
  {
    "triggerProps":
    {
      "ScheduledTriggerProps": null
    },
    "triggerType": "OnDemand"
  }
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/appflow-2020-08-23/DescribeFlow)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/appflow-2020-08-23/DescribeFlow)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/appflow-2020-08-23/DescribeFlow)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/appflow-2020-08-23/DescribeFlow)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/appflow-2020-08-23/DescribeFlow)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/appflow-2020-08-23/DescribeFlow)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/appflow-2020-08-23/DescribeFlow)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/appflow-2020-08-23/DescribeFlow)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/appflow-2020-08-23/DescribeFlow)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/appflow-2020-08-23/DescribeFlow)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DescribeConnectors

DescribeFlowExecutionRecords
