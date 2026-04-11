# CreateFlow

Enables your application to create a new flow using Amazon AppFlow. You must create
a connector profile before calling this API. Please note that the Request Syntax below shows
syntax for multiple destinations, however, you can only transfer data to one item in this list
at a time. Amazon AppFlow does not currently support flows to multiple destinations at
once.

## Request Syntax

```nohighlight

POST /create-flow HTTP/1.1
Content-type: application/json

{
   "clientToken": "string",
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
   "flowName": "string",
   "kmsArn": "string",
   "metadataCatalogConfig": {
      "glueDataCatalog": {
         "databaseName": "string",
         "roleArn": "string",
         "tablePrefix": "string"
      }
   },
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

## URI Request Parameters

The request does not use any URI parameters.

## Request Body

The request accepts the following data in JSON format.

**[clientToken](#API_CreateFlow_RequestSyntax)**

The `clientToken` parameter is an idempotency token. It ensures that your
`CreateFlow` request completes only once. You choose the value to pass. For
example, if you don't receive a response from your request, you can safely retry the request
with the same `clientToken` parameter value.

If you omit a `clientToken` value, the AWS SDK that you are
using inserts a value for you. This way, the SDK can safely retry requests multiple times
after a network error. You must provide your own value for other use cases.

If you specify input parameters that differ from your first request, an error occurs. If
you use a different value for `clientToken`, Amazon AppFlow considers it a new
call to `CreateFlow`. The token is active for 8 hours.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Pattern: `[ -~]+`

Required: No

**[description](#API_CreateFlow_RequestSyntax)**

A description of the flow you want to create.

Type: String

Length Constraints: Maximum length of 2048.

Pattern: `[\w!@#\-.?,\s]*`

Required: No

**[destinationFlowConfigList](#API_CreateFlow_RequestSyntax)**

The configuration that controls how Amazon AppFlow places data in the destination
connector.

Type: Array of [DestinationFlowConfig](api-destinationflowconfig.md) objects

Required: Yes

**[flowName](#API_CreateFlow_RequestSyntax)**

The specified name of the flow. Spaces are not allowed. Use underscores (\_) or hyphens
(-) only.

Type: String

Length Constraints: Maximum length of 256.

Pattern: `[a-zA-Z0-9][\w!@#.-]+`

Required: Yes

**[kmsArn](#API_CreateFlow_RequestSyntax)**

The ARN (Amazon Resource Name) of the Key Management Service (KMS) key you provide for
encryption. This is required if you do not want to use the Amazon AppFlow-managed KMS
key. If you don't provide anything here, Amazon AppFlow uses the Amazon AppFlow-managed KMS key.

Type: String

Length Constraints: Minimum length of 20. Maximum length of 2048.

Pattern: `arn:aws:kms:.*:[0-9]+:.*`

Required: No

**[metadataCatalogConfig](#API_CreateFlow_RequestSyntax)**

Specifies the configuration that Amazon AppFlow uses when it catalogs the data that's
transferred by the associated flow. When Amazon AppFlow catalogs the data from a flow, it
stores metadata in a data catalog.

Type: [MetadataCatalogConfig](api-metadatacatalogconfig.md) object

Required: No

**[sourceFlowConfig](#API_CreateFlow_RequestSyntax)**

The configuration that controls how Amazon AppFlow retrieves data from the source
connector.

Type: [SourceFlowConfig](api-sourceflowconfig.md) object

Required: Yes

**[tags](#API_CreateFlow_RequestSyntax)**

The tags used to organize, track, or control access for your flow.

Type: String to string map

Map Entries: Minimum number of 0 items. Maximum number of 50 items.

Key Length Constraints: Minimum length of 1. Maximum length of 128.

Key Pattern: `^(?!aws:)[a-zA-Z+-=._:/]+$`

Value Length Constraints: Maximum length of 256.

Value Pattern: `[\s\w+-=\.:/@]*`

Required: No

**[tasks](#API_CreateFlow_RequestSyntax)**

A list of tasks that Amazon AppFlow performs while transferring the data in the flow
run.

Type: Array of [Task](api-task.md) objects

Required: Yes

**[triggerConfig](#API_CreateFlow_RequestSyntax)**

The trigger settings that determine how and when the flow runs.

Type: [TriggerConfig](api-triggerconfig.md) object

Required: Yes

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "flowArn": "string",
   "flowStatus": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[flowArn](#API_CreateFlow_ResponseSyntax)**

The flow's Amazon Resource Name (ARN).

Type: String

Length Constraints: Maximum length of 512.

Pattern: `arn:aws:appflow:.*:[0-9]+:.*`

**[flowStatus](#API_CreateFlow_ResponseSyntax)**

Indicates the current status of the flow.

Type: String

Valid Values: `Active | Deprecated | Deleted | Draft | Errored | Suspended`

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**AccessDeniedException**

AppFlow/Requester has invalid or missing permissions.

HTTP Status Code: 403

**ConflictException**

There was a conflict when processing the request (for example, a flow with the given name
already exists within the account. Check for conflicting resource names and try again.

HTTP Status Code: 409

**ConnectorAuthenticationException**

An error occurred when authenticating with the connector endpoint.

HTTP Status Code: 401

**ConnectorServerException**

An error occurred when retrieving data from the connector endpoint.

HTTP Status Code: 400

**InternalServerException**

An internal service error occurred during the processing of your request. Try again
later.

HTTP Status Code: 500

**ResourceNotFoundException**

The resource specified in the request (such as the source or destination connector
profile) is not found.

HTTP Status Code: 404

**ServiceQuotaExceededException**

The request would cause a service quota (such as the number of flows) to be exceeded.

HTTP Status Code: 402

**ValidationException**

The request has invalid or missing parameters.

HTTP Status Code: 400

## Examples

### Salesforce

This example shows a sample request for the `CreateFlow` API using
Salesforce. If you see the error shown in the second sample, it is because of an expired
access or refresh token. Retrieve a new access or refresh token to proceed.

#### Sample Request

```json

{
  "flowName": "testFlowSaleforce",
  "description": "TestFlow",
  "triggerConfig": {
    "triggerType": "Scheduled",
    "triggerProperties": {
        "scheduledTriggerProperties" : {
            "scheduleExpression" : "rate(1minutes)"
        }
    }
  },
  "sourceFlowConfig": {
    "connectorType": "Salesforce",
    "connectorLabel":"MyCustomConnector",
    "connectorProfileName": "Connector_Profile_Name",
    "sourceConnectorProperties": {
      "salesforce": {
         "object": "account"
      }
    }
  },
  "destinationFlowConfigList": [
    {
      "connectorType": "S3",
      "connectorLabel":"MyCustomConnector",
      "destinationConnectorProperties": {
        "s3": {
          "bucketName": "BucketName",
          "bucketKey": "BucketKey"
        }
      }
    }
  ],
  "tasks": [
    {
      "sourceFields": [
        "CreatedDate"
      ],
      "taskProperties": {

      },
      "destinationField": "CreatedDate",
      "taskType": "Map"
    },
    {
      "destinationField": "LastActivityDate",
      "sourceFields": [
        "LastActivityDate"
      ],
      "taskProperties": {

      },
      "taskType": "Map"
    }
  ]
}
```

```json

{
  "message": "Error while authenticating to connector"
}
```

### Zendesk

This example shows a sample request for the `CreateFlow` API using
Zendesk.

#### Sample Request

```json

{
  "useDefaultEncryption": false,
  "flowName": "testFlowZendesk",
  "description": "TestFlow",
  "triggerConfig": {
    "triggerType": "Scheduled",
    "triggerProperties": {
        "scheduledTriggerProperties" : {
            "scheduleExpression" : "rate(1minutes)"
        }
    }
  },
  "sourceFlowConfig": {
    "connectorType": "Zendesk",
    "connectorProfileName": "connector-profile-name",
    "sourceConnectorProperties": {
      "zendesk": {
         "object": "tickets"
      }
    }
  },
  "destinationFlowConfigList": [
    {
      "connectorType": "S3",
      "connectorProfileName": "Test",
      "destinationConnectorProperties": {
        "s3": {
          "bucketName": "BucketName",
          "bucketKey": "BucketKey"
        }
      }
    }
  ],
  "tasks": [
    {
      "destinationField": "url",
      "sourceFields": [
        "url"
      ],
      "taskProperties": {
          "SOURCE_DATA_TYPE": "integer"
      },
      "taskType": "Map"
    }
  ]
}
```

### Google Analytics

This example shows a sample request for the `CreateFlow` API using Google
Analytics.

#### Sample Request

```json

 {
  "useDefaultEncryption": false,
  "flowName": "testFlowGAnalytics",
  "description": "TestFlow",
  "triggerConfig": {
    "triggerType": "Scheduled",
    "triggerProperties": {
        "scheduledTriggerProperties" : {
            "scheduleExpression" : "rate(1minutes)"
        }
    }
  },
  "sourceFlowConfig": {
    "connectorType": "Googleanalytics",
    "connectorProfileName": "connector-profile-name",
    "sourceConnectorProperties": {
      "googleanalytics": {
         "object": "ObjectID"
      }
    }
  },
  "destinationFlowConfigList": [
    {
      "connectorType": "S3",
      "connectorProfileName": "Test",
      "destinationConnectorProperties": {
        "s3": {
          "bucketName": "BucketName",
          "bucketKey": "BucketKey"
        }
      }
    }
  ],
  "tasks": [
    {
      "destinationField": "url",
      "sourceFields": [
        "url"
      ],
      "taskProperties": {
          "SOURCE_DATA_TYPE": "integer"
      },
      "taskType": "Map"
    }
  ]
}
```

### Marketo

This example shows a sample request for the `CreateFlow` API using
Marketo.

#### Sample Request

```json

{
  "flowName": "testMarketo",
  "description": "TestFlow",
  "triggerConfig": {
    "triggerType": "Scheduled",
    "triggerProperties": {
        "scheduledTriggerProperties" : {
            "scheduleExpression" : "rate(1minutes)"
        }
    }
  },
  "sourceFlowConfig": {
    "connectorType": "Marketo",
    "connectorProfileName": "Connector-profile-new",
    "sourceConnectorProperties": {
      "marketo": {
         "object": "leads"
      }
    }
  },
  "destinationFlowConfigList": [
    {
      "connectorType": "S3",
      "connectorProfileName": "Test",
      "destinationConnectorProperties": {
        "s3": {
          "bucketName": "BucketName",
          "bucketKey": "BucketKey"
        }
      }
    }
  ],
  "tasks": [
      {
      "connectorOperator": {
          "marketo":"BETWEEN"
      },
      "sourceFields": [
        "updatedAt"
      ],
      "taskProperties": {
        "DATA_TYPE": "datetime",
        "LOWER_BOUND": "Lower_Bound_value",
        "UPPER_BOUND": "Upper_Bound_value"
      },
      "taskType": "Filter"
    },
    {
      "destinationField": "company",
      "sourceFields": [
        "company"
      ],
      "taskProperties": {
          "SOURCE_DATA_TYPE": "string"
      },
      "taskType": "Map"
    }
  ]
}
```

### Connection credentials

This Java example shows credentials associated with a connection. In this case, the
example uses OAuth. These credentials are stored in secrets manager after the profile is
created.

```java

ConnectorProfileCredentials credentials = new ConnectorProfileCredentials().withSalesforce(
        new SalesforceConnectorProfileCredentials().withAccessToken(accessToken)
           .withRefreshToken(refreshToken));
```

### Connection properties

This Java example shows other properties associated with a connection to Salesforce.

```java

ConnectorProfileProperties properties = new ConnectorProfileProperties().withSalesforce(
                new SalesforceConnectorProfileProperties().withInstanceUrl(instanceUrl));
```

### Connector profile creation

To create a flow, you must first create a connector profile. This Java example shows
information about connecting to Salesforce.

```java

amazonAppflow.createConnectorProfile(new CreateConnectorProfileRequest().withConnectorProfileConfig(
                new ConnectorProfileConfig().withConnectorProfileCredentials(credentials)
                        .withConnectorProfileProperties(properties))
                .withConnectorProfileName(salesforceProfileName)
                .withConnectorType(ConnectorType.Salesforce)
                .withConnectionMode(ConnectionMode.Public));
```

### Source connection properties

This Java example shows properties for a source connection. `Account` is
the object in Salesforce that we want to retrieve. You can find a list of all supported
objects by using the `listConnectorEntity` API. The
`ConnectorProfileName` is the connector profile, the creation of which is
seen in the previous example.

```java

  SourceFlowConfig sourceFlowConfig = new SourceFlowConfig().withSourceConnectorProperties(
                new SourceConnectorProperties().withSalesforce(new SalesforceSourceProperties().withObject("Account")))
                .withConnectorType(ConnectorType.Salesforce)
                .withConnectorProfileName(salesforceProfileName);
```

### Destination connection properties

This Java example shows properties for a destination connection. Note that many
AWS connectors such as Amazon S3 don't require a connector
profile. Amazon AppFlow accesses S3 buckets through a bucket resource policy,
therefore a connector profile isn't needed.

```java

   DestinationFlowConfig destinationFlowConfig = new DestinationFlowConfig().withConnectorType(ConnectorType.S3)
                .withDestinationConnectorProperties(new DestinationConnectorProperties().withS3(
                        new S3DestinationProperties().withBucketName(bucketName).withBucketPrefix("testPrefix")));
```

### Tasks

Tasks describe what to do with the data once it has been retrieved, but before it is
sent to the destination. Most connectors require a projection task. A projection task
describes what fields should be retrieved from the source object. Fields that can be
retrieved can be discovered by making a call to the `DescribeConnectorEntity`
API.

```java

Task projectionTask = new Task().withTaskType(TaskType.Filter)
                .withConnectorOperator(new ConnectorOperator().withSalesforce(SalesforceConnectorOperator.PROJECTION))
                .withSourceFields("Id", "Name");
```

### Mapping task

Most flows also require at least one mapping task. Mapping tasks map a source field to
a destination field. This Java example shows the mapping between the retrieved field,
`Id`, to a new field, `AccountId`.

```java

Task createdMappingTask = new Task().withTaskType(TaskType.Map)
                .withConnectorOperator(new ConnectorOperator().withSalesforce(SalesforceConnectorOperator.NO_OP))
                .withSourceFields("Id")
                .withDestinationField("AccountId");

        amazonAppflow.createFlow(new CreateFlowRequest().withFlowName(salesforceFlowName)
                .withTriggerConfig(new TriggerConfig().withTriggerType(TriggerType.OnDemand))
                .withSourceFlowConfig(sourceFlowConfig)
                .withDestinationFlowConfigList(destinationFlowConfig)
                .withTasks(projectionTask, idMappingTask, createdMappingTask));
```

### Sample imports

This Java example shows sample imports.

```java

import com.amazonaws.services.appflow.AmazonAppflow;
import com.amazonaws.services.appflow.model.ConnectionMode;
import com.amazonaws.services.appflow.model.ConnectorOperator;
import com.amazonaws.services.appflow.model.ConnectorProfileConfig;
import com.amazonaws.services.appflow.model.ConnectorProfileCredentials;
import com.amazonaws.services.appflow.model.ConnectorProfileProperties;
import com.amazonaws.services.appflow.model.ConnectorType;
import com.amazonaws.services.appflow.model.CreateConnectorProfileRequest;
import com.amazonaws.services.appflow.model.CreateFlowRequest;
import com.amazonaws.services.appflow.model.DeleteConnectorProfileRequest;
import com.amazonaws.services.appflow.model.DeleteFlowRequest;
import com.amazonaws.services.appflow.model.DestinationConnectorProperties;
import com.amazonaws.services.appflow.model.DestinationFlowConfig;
import com.amazonaws.services.appflow.model.S3DestinationProperties;
import com.amazonaws.services.appflow.model.SalesforceConnectorOperator;
import com.amazonaws.services.appflow.model.SalesforceConnectorProfileCredentials;
import com.amazonaws.services.appflow.model.SalesforceConnectorProfileProperties;
import com.amazonaws.services.appflow.model.SalesforceSourceProperties;
import com.amazonaws.services.appflow.model.SourceConnectorProperties;
import com.amazonaws.services.appflow.model.SourceFlowConfig;
import com.amazonaws.services.appflow.model.Task;
import com.amazonaws.services.appflow.model.TaskType;
import com.amazonaws.services.appflow.model.TriggerConfig;
import com.amazonaws.services.appflow.model.TriggerType;

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/appflow-2020-08-23/createflow.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/appflow-2020-08-23/createflow.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/createflow.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/appflow-2020-08-23/createflow.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/createflow.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/appflow-2020-08-23/createflow.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/appflow-2020-08-23/createflow.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/appflow-2020-08-23/createflow.md)

- [AWS SDK for Python](../../../../services/goto/boto3/appflow-2020-08-23/createflow.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/createflow.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateConnectorProfile

DeleteConnectorProfile

All content copied from https://docs.aws.amazon.com/.
