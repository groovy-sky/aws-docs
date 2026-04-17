---
title: "PutReplicationConfiguration"
---

# PutReplicationConfiguration

Creates or updates the replication configuration for a registry. The existing
replication configuration for a repository can be retrieved with the [DescribeRegistry](api-describeregistry.md) API action. The first time the
PutReplicationConfiguration API is called, a service-linked IAM role is created in
your account for the replication process. For more information, see [Using\
service-linked roles for Amazon ECR](../../../../services/amazonecr/latest/userguide/using-service-linked-roles.md) in the _Amazon Elastic Container Registry User Guide_.
For more information on the custom role for replication, see [Creating an IAM role for replication](../../../../services/amazonecr/latest/userguide/replication-creation-templates.md#roles-creatingrole-user-console).

###### Note

When configuring cross-account replication, the destination account must grant the
source account permission to replicate. This permission is controlled using a
registry permissions policy. For more information, see [PutRegistryPolicy](api-putregistrypolicy.md).

## Request Syntax

```nohighlight

{
   "replicationConfiguration": {
      "rules": [
         {
            "destinations": [
               {
                  "region": "string",
                  "registryId": "string"
               }
            ],
            "repositoryFilters": [
               {
                  "filter": "string",
                  "filterType": "string"
               }
            ]
         }
      ]
   }
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[replicationConfiguration](#API_PutReplicationConfiguration_RequestSyntax)**

An object representing the replication configuration for a registry.

Type: [ReplicationConfiguration](api-replicationconfiguration.md) object

Required: Yes

## Response Syntax

```nohighlight

{
   "replicationConfiguration": {
      "rules": [
         {
            "destinations": [
               {
                  "region": "string",
                  "registryId": "string"
               }
            ],
            "repositoryFilters": [
               {
                  "filter": "string",
                  "filterType": "string"
               }
            ]
         }
      ]
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[replicationConfiguration](#API_PutReplicationConfiguration_ResponseSyntax)**

The contents of the replication configuration for the registry.

Type: [ReplicationConfiguration](api-replicationconfiguration.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidParameterException**

The specified parameter is invalid. Review the available parameters for the API
request.

**message**

The error message associated with the exception.

HTTP Status Code: 400

**ServerException**

These errors are usually caused by a server-side issue.

**message**

The error message associated with the exception.

HTTP Status Code: 500

**ValidationException**

There was an exception validating this request.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ecr-2015-09-21/PutReplicationConfiguration)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ecr-2015-09-21/PutReplicationConfiguration)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ecr-2015-09-21/PutReplicationConfiguration)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ecr-2015-09-21/PutReplicationConfiguration)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ecr-2015-09-21/PutReplicationConfiguration)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ecr-2015-09-21/PutReplicationConfiguration)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ecr-2015-09-21/PutReplicationConfiguration)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ecr-2015-09-21/PutReplicationConfiguration)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ecr-2015-09-21/PutReplicationConfiguration)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ecr-2015-09-21/PutReplicationConfiguration)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PutRegistryScanningConfiguration

PutSigningConfiguration

All content copied from https://docs.aws.amazon.com/.
