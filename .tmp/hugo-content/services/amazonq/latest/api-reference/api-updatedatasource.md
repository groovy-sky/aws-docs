# UpdateDataSource

Updates an existing Amazon Q Business data source connector.

## Request Syntax

```nohighlight

PUT /applications/applicationId/indices/indexId/datasources/dataSourceId HTTP/1.1
Content-type: application/json

{
   "configuration": JSON value,
   "description": "string",
   "displayName": "string",
   "documentEnrichmentConfiguration": {
      "inlineConfigurations": [
         {
            "condition": {
               "key": "string",
               "operator": "string",
               "value": { ... }
            },
            "documentContentOperator": "string",
            "target": {
               "attributeValueOperator": "string",
               "key": "string",
               "value": { ... }
            }
         }
      ],
      "postExtractionHookConfiguration": {
         "invocationCondition": {
            "key": "string",
            "operator": "string",
            "value": { ... }
         },
         "lambdaArn": "string",
         "roleArn": "string",
         "s3BucketName": "string"
      },
      "preExtractionHookConfiguration": {
         "invocationCondition": {
            "key": "string",
            "operator": "string",
            "value": { ... }
         },
         "lambdaArn": "string",
         "roleArn": "string",
         "s3BucketName": "string"
      }
   },
   "mediaExtractionConfiguration": {
      "audioExtractionConfiguration": {
         "audioExtractionStatus": "string"
      },
      "imageExtractionConfiguration": {
         "imageExtractionStatus": "string"
      },
      "videoExtractionConfiguration": {
         "videoExtractionStatus": "string"
      }
   },
   "roleArn": "string",
   "syncSchedule": "string",
   "vpcConfiguration": {
      "securityGroupIds": [ "string" ],
      "subnetIds": [ "string" ]
   }
}
```

## URI Request Parameters

The request uses the following URI parameters.

**[applicationId](#API_UpdateDataSource_RequestSyntax)**

The identifier of the Amazon Q Business application the data source is attached
to.

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

Required: Yes

**[dataSourceId](#API_UpdateDataSource_RequestSyntax)**

The identifier of the data source connector.

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

Required: Yes

**[indexId](#API_UpdateDataSource_RequestSyntax)**

The identifier of the index attached to the data source connector.

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

Required: Yes

## Request Body

The request accepts the following data in JSON format.

**[configuration](#API_UpdateDataSource_RequestSyntax)**

Provides the configuration information for an Amazon Q Business data source.

Type: JSON value

Required: No

**[description](#API_UpdateDataSource_RequestSyntax)**

The description of the data source connector.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1000.

Pattern: `[\s\S]*`

Required: No

**[displayName](#API_UpdateDataSource_RequestSyntax)**

A name of the data source connector.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1000.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9_-]*`

Required: No

**[documentEnrichmentConfiguration](#API_UpdateDataSource_RequestSyntax)**

Provides the configuration information for altering document metadata and content
during the document ingestion process.

For more information, see [Custom document\
enrichment](../business-use-dg/custom-document-enrichment.md).

Type: [DocumentEnrichmentConfiguration](api-documentenrichmentconfiguration.md) object

Required: No

**[mediaExtractionConfiguration](#API_UpdateDataSource_RequestSyntax)**

The configuration for extracting information from media in documents for your data source.

Type: [MediaExtractionConfiguration](api-mediaextractionconfiguration.md) object

Required: No

**[roleArn](#API_UpdateDataSource_RequestSyntax)**

The Amazon Resource Name (ARN) of an IAM role with permission to access
the data source and required resources.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1284.

Pattern: `arn:[a-z0-9-\.]{1,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[^/].{0,1023}`

Required: No

**[syncSchedule](#API_UpdateDataSource_RequestSyntax)**

The chosen update frequency for your data source.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 998.

Pattern: `[\s\S]*`

Required: No

**[vpcConfiguration](#API_UpdateDataSource_RequestSyntax)**

Provides configuration information needed to connect to an Amazon VPC (Virtual
Private Cloud).

Type: [DataSourceVpcConfiguration](api-datasourcevpcconfiguration.md) object

Required: No

## Response Syntax

```

HTTP/1.1 200

```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**AccessDeniedException**

You don't have access to perform this action. Make sure you have the required
permission policies and user accounts and try again.

HTTP Status Code: 403

**ConflictException**

You are trying to perform an action that conflicts with the current status of your
resource. Fix any inconsistencies with your resources and try again.

**message**

The message describing a `ConflictException`.

**resourceId**

The identifier of the resource affected.

**resourceType**

The type of the resource affected.

HTTP Status Code: 409

**InternalServerException**

An issue occurred with the internal server used for your Amazon Q Business service. Wait
some minutes and try again, or contact [Support](http://aws.amazon.com/contact-us) for help.

HTTP Status Code: 500

**ResourceNotFoundException**

The application or plugin resource you want to use doesn’t exist. Make sure you have
provided the correct resource and try again.

**message**

The message describing a `ResourceNotFoundException`.

**resourceId**

The identifier of the resource affected.

**resourceType**

The type of the resource affected.

HTTP Status Code: 404

**ThrottlingException**

The request was denied due to throttling. Reduce the number of requests and try
again.

HTTP Status Code: 429

**ValidationException**

The input doesn't meet the constraints set by the Amazon Q Business service. Provide the
correct input and try again.

**fields**

The input field(s) that failed validation.

**message**

The message describing the `ValidationException`.

**reason**

The reason for the `ValidationException`.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/qbusiness-2023-11-27/updatedatasource.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/qbusiness-2023-11-27/updatedatasource.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/updatedatasource.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/qbusiness-2023-11-27/updatedatasource.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/updatedatasource.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/qbusiness-2023-11-27/updatedatasource.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/qbusiness-2023-11-27/updatedatasource.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/qbusiness-2023-11-27/updatedatasource.md)

- [AWS SDK for Python](../../../goto/boto3/qbusiness-2023-11-27/updatedatasource.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/updatedatasource.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UpdateDataAccessor

UpdateIndex

All content copied from https://docs.aws.amazon.com/.
