# CreateWorkGroup

Creates a workgroup with the specified name. A workgroup can be an Apache Spark
enabled workgroup or an Athena SQL workgroup.

## Request Syntax

```nohighlight

{
   "Configuration": {
      "AdditionalConfiguration": "string",
      "BytesScannedCutoffPerQuery": number,
      "CustomerContentEncryptionConfiguration": {
         "KmsKey": "string"
      },
      "EnableMinimumEncryptionConfiguration": boolean,
      "EnforceWorkGroupConfiguration": boolean,
      "EngineConfiguration": {
         "AdditionalConfigs": {
            "string" : "string"
         },
         "Classifications": [
            {
               "Name": "string",
               "Properties": {
                  "string" : "string"
               }
            }
         ],
         "CoordinatorDpuSize": number,
         "DefaultExecutorDpuSize": number,
         "MaxConcurrentDpus": number,
         "SparkProperties": {
            "string" : "string"
         }
      },
      "EngineVersion": {
         "EffectiveEngineVersion": "string",
         "SelectedEngineVersion": "string"
      },
      "ExecutionRole": "string",
      "IdentityCenterConfiguration": {
         "EnableIdentityCenter": boolean,
         "IdentityCenterInstanceArn": "string"
      },
      "ManagedQueryResultsConfiguration": {
         "Enabled": boolean,
         "EncryptionConfiguration": {
            "KmsKey": "string"
         }
      },
      "MonitoringConfiguration": {
         "CloudWatchLoggingConfiguration": {
            "Enabled": boolean,
            "LogGroup": "string",
            "LogStreamNamePrefix": "string",
            "LogTypes": {
               "string" : [ "string" ]
            }
         },
         "ManagedLoggingConfiguration": {
            "Enabled": boolean,
            "KmsKey": "string"
         },
         "S3LoggingConfiguration": {
            "Enabled": boolean,
            "KmsKey": "string",
            "LogLocation": "string"
         }
      },
      "PublishCloudWatchMetricsEnabled": boolean,
      "QueryResultsS3AccessGrantsConfiguration": {
         "AuthenticationType": "string",
         "CreateUserLevelPrefix": boolean,
         "EnableS3AccessGrants": boolean
      },
      "RequesterPaysEnabled": boolean,
      "ResultConfiguration": {
         "AclConfiguration": {
            "S3AclOption": "string"
         },
         "EncryptionConfiguration": {
            "EncryptionOption": "string",
            "KmsKey": "string"
         },
         "ExpectedBucketOwner": "string",
         "OutputLocation": "string"
      }
   },
   "Description": "string",
   "Name": "string",
   "Tags": [
      {
         "Key": "string",
         "Value": "string"
      }
   ]
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[Configuration](#API_CreateWorkGroup_RequestSyntax)**

Contains configuration information for creating an Athena SQL workgroup or
Spark enabled Athena workgroup. Athena SQL workgroup
configuration includes the location in Amazon S3 where query and calculation
results are stored, the encryption configuration, if any, used for encrypting query
results, whether the Amazon CloudWatch Metrics are enabled for the workgroup, the
limit for the amount of bytes scanned (cutoff) per query, if it is specified, and
whether workgroup's settings (specified with `EnforceWorkGroupConfiguration`)
in the `WorkGroupConfiguration` override client-side settings. See [WorkGroupConfiguration:EnforceWorkGroupConfiguration](api-workgroupconfiguration-athena-type-workgroupconfiguration-enforceworkgroupconfiguration.md).

Type: [WorkGroupConfiguration](api-workgroupconfiguration.md) object

Required: No

**[Description](#API_CreateWorkGroup_RequestSyntax)**

The workgroup description.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1024.

Required: No

**[Name](#API_CreateWorkGroup_RequestSyntax)**

The workgroup name.

Type: String

Pattern: `[a-zA-Z0-9._-]{1,128}`

Required: Yes

**[Tags](#API_CreateWorkGroup_RequestSyntax)**

A list of comma separated tags to add to the workgroup that is created.

Type: Array of [Tag](api-tag.md) objects

Required: No

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InternalServerException**

Indicates a platform issue, which may be due to a transient condition or
outage.

HTTP Status Code: 500

**InvalidRequestException**

Indicates that something is wrong with the input to the request. For example, a
required parameter may be missing or out of range.

**AthenaErrorCode**

The error code returned when the query execution failed to process, or when the
processing request for the named query failed.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/athena-2017-05-18/createworkgroup.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/athena-2017-05-18/createworkgroup.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/athena-2017-05-18/createworkgroup.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/athena-2017-05-18/createworkgroup.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/athena-2017-05-18/createworkgroup.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/athena-2017-05-18/createworkgroup.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/athena-2017-05-18/createworkgroup.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/athena-2017-05-18/createworkgroup.md)

- [AWS SDK for Python](../../../../services/goto/boto3/athena-2017-05-18/createworkgroup.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/athena-2017-05-18/createworkgroup.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreatePresignedNotebookUrl

DeleteCapacityReservation
