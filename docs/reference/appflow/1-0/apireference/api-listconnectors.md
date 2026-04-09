# ListConnectors

Returns the list of all registered custom connectors in your AWS account.
This API lists only custom connectors registered in this account, not the AWS
authored connectors.

## Request Syntax

```nohighlight

POST /list-connectors HTTP/1.1
Content-type: application/json

{
   "maxResults": number,
   "nextToken": "string"
}
```

## URI Request Parameters

The request does not use any URI parameters.

## Request Body

The request accepts the following data in JSON format.

**[maxResults](#API_ListConnectors_RequestSyntax)**

Specifies the maximum number of items that should be returned in the result set. The
default for `maxResults` is 20 (for all paginated API operations).

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 100.

Required: No

**[nextToken](#API_ListConnectors_RequestSyntax)**

The pagination token for the next page of data.

Type: String

Length Constraints: Maximum length of 2048.

Pattern: `\S+`

Required: No

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "connectors": [
      {
         "applicationType": "string",
         "connectorDescription": "string",
         "connectorLabel": "string",
         "connectorModes": [ "string" ],
         "connectorName": "string",
         "connectorOwner": "string",
         "connectorProvisioningType": "string",
         "connectorType": "string",
         "connectorVersion": "string",
         "registeredAt": number,
         "registeredBy": "string",
         "supportedDataTransferTypes": [ "string" ]
      }
   ],
   "nextToken": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[connectors](#API_ListConnectors_ResponseSyntax)**

Contains information about the connectors supported by Amazon AppFlow.

Type: Array of [ConnectorDetail](api-connectordetail.md) objects

**[nextToken](#API_ListConnectors_ResponseSyntax)**

The pagination token for the next page of data. If nextToken=null, this means that all
records have been fetched.

Type: String

Length Constraints: Maximum length of 2048.

Pattern: `\S+`

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InternalServerException**

An internal service error occurred during the processing of your request. Try again
later.

HTTP Status Code: 500

**ValidationException**

The request has invalid or missing parameters.

HTTP Status Code: 400

## Examples

### Listing connectors

This example shows a sample request for the `ListConnector` API and a
sample response.

#### Sample Request

```json

{
  "maxResults": 1,
  "nextToken": "nextToken_value"
}
```

#### Sample Response

```json

{
  "connectors":
  [
    {
      "connectorArn": "Arn of connector1",
      "connectorDescription": "Some Sample Connector1",
      "connectorName": "Salesforce custom connector1",
      "connectorOwner": "AppFlow",
      "connectorVersion": "1.0",
      "applicationType": "Salesforce",
      "connectorType": "CUSTOMCONNECTOR",
      "connectorLabel": "MyCustomConnector",
      "registeredAt": 1628732168.132,
      "registeredBy": "CUSTOM",
      "tags": null,
      "connectorModes":
      [
        "SOURCE",
        "DESTINATION"
      ],
      "connectorProvisioningType": "LAMBDA"
    },
    {
      "connectorArn": "Arn of connector2",
      "connectorDescription": "Some Sample Connector2",
      "connectorName": "Salesforce custom connector2",
      "connectorOwner": "AppFlow",
      "connectorVersion": "1.0",
      "applicationType": "Salesforce",
      "connectorType": "CUSTOMCONNECTOR",
      "connectorLabel": "MyCustomConnector",
      "registeredAt": 1628732168.132,
      "registeredBy": "CUSTOM",
      "tags": null,
      "connectorModes":
      [
        "SOURCE",
        "DESTINATION"
      ],
      "connectorProvisioningType": "LAMBDA"
    }
  ],
  "nextToken": null
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/appflow-2020-08-23/listconnectors.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/appflow-2020-08-23/listconnectors.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/listconnectors.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/appflow-2020-08-23/listconnectors.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/listconnectors.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/appflow-2020-08-23/listconnectors.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/appflow-2020-08-23/listconnectors.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/appflow-2020-08-23/listconnectors.md)

- [AWS SDK for Python](../../../../services/goto/boto3/appflow-2020-08-23/listconnectors.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/listconnectors.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListConnectorEntities

ListFlows

All content copied from https://docs.aws.amazon.com/.
