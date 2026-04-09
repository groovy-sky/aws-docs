# ResetConnectorMetadataCache

Resets metadata about your connector entities that Amazon AppFlow stored in its
cache. Use this action when you want Amazon AppFlow to return the latest information
about the data that you have in a source application.

Amazon AppFlow returns metadata about your entities when you use the
ListConnectorEntities or DescribeConnectorEntities actions. Following these actions, Amazon AppFlow caches the metadata to reduce the number of API requests that it must send to
the source application. Amazon AppFlow automatically resets the cache once every hour,
but you can use this action when you want to get the latest metadata right away.

## Request Syntax

```nohighlight

POST /reset-connector-metadata-cache HTTP/1.1
Content-type: application/json

{
   "apiVersion": "string",
   "connectorEntityName": "string",
   "connectorProfileName": "string",
   "connectorType": "string",
   "entitiesPath": "string"
}
```

## URI Request Parameters

The request does not use any URI parameters.

## Request Body

The request accepts the following data in JSON format.

**[apiVersion](#API_ResetConnectorMetadataCache_RequestSyntax)**

The API version that you specified in the connector profile that you’re resetting cached
metadata for. You must use this parameter only if the connector supports multiple API versions
or if the connector type is CustomConnector.

To look up how many versions a connector supports, use the DescribeConnectors action. In
the response, find the value that Amazon AppFlow returns for the connectorVersion
parameter.

To look up the connector type, use the DescribeConnectorProfiles action. In the response,
find the value that Amazon AppFlow returns for the connectorType parameter.

To look up the API version that you specified in a connector profile, use the
DescribeConnectorProfiles action.

Type: String

Length Constraints: Maximum length of 256.

Pattern: `\S+`

Required: No

**[connectorEntityName](#API_ResetConnectorMetadataCache_RequestSyntax)**

Use this parameter if you want to reset cached metadata about the details for an
individual entity.

If you don't include this parameter in your request, Amazon AppFlow only resets
cached metadata about entity names, not entity details.

Type: String

Length Constraints: Maximum length of 1024.

Pattern: `\S+`

Required: No

**[connectorProfileName](#API_ResetConnectorMetadataCache_RequestSyntax)**

The name of the connector profile that you want to reset cached metadata for.

You can omit this parameter if you're resetting the cache for any of the following
connectors: Amazon Connect, Amazon EventBridge, Amazon Lookout for Metrics, Amazon S3, or Upsolver. If you're resetting the cache for any other connector, you must include this
parameter in your request.

Type: String

Length Constraints: Maximum length of 256.

Pattern: `[\w/!@#+=.-]+`

Required: No

**[connectorType](#API_ResetConnectorMetadataCache_RequestSyntax)**

The type of connector to reset cached metadata for.

You must include this parameter in your request if you're resetting the cache for any of
the following connectors: Amazon Connect, Amazon EventBridge, Amazon Lookout for Metrics,
Amazon S3, or Upsolver. If you're resetting the cache for any other connector, you
can omit this parameter from your request.

Type: String

Valid Values: `Salesforce | Singular | Slack | Redshift | S3 | Marketo | Googleanalytics | Zendesk | Servicenow | Datadog | Trendmicro | Snowflake | Dynatrace | Infornexus | Amplitude | Veeva | EventBridge | LookoutMetrics | Upsolver | Honeycode | CustomerProfiles | SAPOData | CustomConnector | Pardot`

Required: No

**[entitiesPath](#API_ResetConnectorMetadataCache_RequestSyntax)**

Use this parameter only if you’re resetting the cached metadata about a nested entity.
Only some connectors support nested entities. A nested entity is one that has another entity
as a parent. To use this parameter, specify the name of the parent entity.

To look up the parent-child relationship of entities, you can send a ListConnectorEntities
request that omits the entitiesPath parameter. Amazon AppFlow will return a list of
top-level entities. For each one, it indicates whether the entity has nested entities. Then,
in a subsequent ListConnectorEntities request, you can specify a parent entity name for the
entitiesPath parameter. Amazon AppFlow will return a list of the child entities for that
parent.

Type: String

Length Constraints: Maximum length of 256.

Pattern: `[\s\w/!@#+=,.-]*`

Required: No

## Response Syntax

```

HTTP/1.1 200

```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ConflictException**

There was a conflict when processing the request (for example, a flow with the given name
already exists within the account. Check for conflicting resource names and try again.

HTTP Status Code: 409

**InternalServerException**

An internal service error occurred during the processing of your request. Try again
later.

HTTP Status Code: 500

**ResourceNotFoundException**

The resource specified in the request (such as the source or destination connector
profile) is not found.

HTTP Status Code: 404

**ValidationException**

The request has invalid or missing parameters.

HTTP Status Code: 400

## Examples

### Reset cache of entity names

The following example shows the request body for a call to reset the cached metadata
for entity names.

#### Sample Request

```

{
   "apiVersion": "v1.0",
   "connectorProfileName": "MyCustomConnector"
}
```

#### Sample Response

```

{}
```

### Reset cache of entity details

The following example shows the request body for a call to reset the cached metadata
about the details for a specific entity.

#### Sample Request

```

{
   "apiVersion": "v1.0",
   "connectorProfileName": "MyCustomConnector",
   "connectorEntityName": "EntityName"
}
```

#### Sample Response

```

{}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/appflow-2020-08-23/resetconnectormetadatacache.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/appflow-2020-08-23/resetconnectormetadatacache.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/resetconnectormetadatacache.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/appflow-2020-08-23/resetconnectormetadatacache.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/resetconnectormetadatacache.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/appflow-2020-08-23/resetconnectormetadatacache.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/appflow-2020-08-23/resetconnectormetadatacache.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/appflow-2020-08-23/resetconnectormetadatacache.md)

- [AWS SDK for Python](../../../../services/goto/boto3/appflow-2020-08-23/resetconnectormetadatacache.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/resetconnectormetadatacache.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RegisterConnector

StartFlow

All content copied from https://docs.aws.amazon.com/.
