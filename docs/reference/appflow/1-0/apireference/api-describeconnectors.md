# DescribeConnectors

Describes the connectors vended by Amazon AppFlow for specified connector types. If
you don't specify a connector type, this operation describes all connectors vended by Amazon AppFlow. If there are more connectors than can be returned in one page, the response
contains a `nextToken` object, which can be be passed in to the next call to the
`DescribeConnectors` API operation to retrieve the next page.

## Request Syntax

```nohighlight

POST /describe-connectors HTTP/1.1
Content-type: application/json

{
   "connectorTypes": [ "string" ],
   "maxResults": number,
   "nextToken": "string"
}
```

## URI Request Parameters

The request does not use any URI parameters.

## Request Body

The request accepts the following data in JSON format.

**[connectorTypes](#API_DescribeConnectors_RequestSyntax)**

The type of connector, such as Salesforce, Amplitude, and so on.

Type: Array of strings

Array Members: Minimum number of 0 items. Maximum number of 100 items.

Valid Values: `Salesforce | Singular | Slack | Redshift | S3 | Marketo | Googleanalytics | Zendesk | Servicenow | Datadog | Trendmicro | Snowflake | Dynatrace | Infornexus | Amplitude | Veeva | EventBridge | LookoutMetrics | Upsolver | Honeycode | CustomerProfiles | SAPOData | CustomConnector | Pardot`

Required: No

**[maxResults](#API_DescribeConnectors_RequestSyntax)**

The maximum number of items that should be returned in the result set. The default is
20.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 100.

Required: No

**[nextToken](#API_DescribeConnectors_RequestSyntax)**

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
   "connectorConfigurations": {
      "string" : {
         "authenticationConfig": {
            "customAuthConfigs": [
               {
                  "authParameters": [
                     {
                        "connectorSuppliedValues": [ "string" ],
                        "description": "string",
                        "isRequired": boolean,
                        "isSensitiveField": boolean,
                        "key": "string",
                        "label": "string"
                     }
                  ],
                  "customAuthenticationType": "string"
               }
            ],
            "isApiKeyAuthSupported": boolean,
            "isBasicAuthSupported": boolean,
            "isCustomAuthSupported": boolean,
            "isOAuth2Supported": boolean,
            "oAuth2Defaults": {
               "authCodeUrls": [ "string" ],
               "oauth2CustomProperties": [
                  {
                     "connectorSuppliedValues": [ "string" ],
                     "description": "string",
                     "isRequired": boolean,
                     "isSensitiveField": boolean,
                     "key": "string",
                     "label": "string",
                     "type": "string"
                  }
               ],
               "oauth2GrantTypesSupported": [ "string" ],
               "oauthScopes": [ "string" ],
               "tokenUrls": [ "string" ]
            }
         },
         "canUseAsDestination": boolean,
         "canUseAsSource": boolean,
         "connectorArn": "string",
         "connectorDescription": "string",
         "connectorLabel": "string",
         "connectorMetadata": {
            "Amplitude": {
            },
            "CustomerProfiles": {
            },
            "Datadog": {
            },
            "Dynatrace": {
            },
            "EventBridge": {
            },
            "GoogleAnalytics": {
               "oAuthScopes": [ "string" ]
            },
            "Honeycode": {
               "oAuthScopes": [ "string" ]
            },
            "InforNexus": {
            },
            "Marketo": {
            },
            "Pardot": {
            },
            "Redshift": {
            },
            "S3": {
            },
            "Salesforce": {
               "dataTransferApis": [ "string" ],
               "oauth2GrantTypesSupported": [ "string" ],
               "oAuthScopes": [ "string" ]
            },
            "SAPOData": {
            },
            "ServiceNow": {
            },
            "Singular": {
            },
            "Slack": {
               "oAuthScopes": [ "string" ]
            },
            "Snowflake": {
               "supportedRegions": [ "string" ]
            },
            "Trendmicro": {
            },
            "Upsolver": {
            },
            "Veeva": {
            },
            "Zendesk": {
               "oAuthScopes": [ "string" ]
            }
         },
         "connectorModes": [ "string" ],
         "connectorName": "string",
         "connectorOwner": "string",
         "connectorProvisioningConfig": {
            "lambda": {
               "lambdaArn": "string"
            }
         },
         "connectorProvisioningType": "string",
         "connectorRuntimeSettings": [
            {
               "connectorSuppliedValueOptions": [ "string" ],
               "dataType": "string",
               "description": "string",
               "isRequired": boolean,
               "key": "string",
               "label": "string",
               "scope": "string"
            }
         ],
         "connectorType": "string",
         "connectorVersion": "string",
         "isPrivateLinkEnabled": boolean,
         "isPrivateLinkEndpointUrlRequired": boolean,
         "logoURL": "string",
         "registeredAt": number,
         "registeredBy": "string",
         "supportedApiVersions": [ "string" ],
         "supportedDataTransferApis": [
            {
               "Name": "string",
               "Type": "string"
            }
         ],
         "supportedDataTransferTypes": [ "string" ],
         "supportedDestinationConnectors": [ "string" ],
         "supportedOperators": [ "string" ],
         "supportedSchedulingFrequencies": [ "string" ],
         "supportedTriggerTypes": [ "string" ],
         "supportedWriteOperations": [ "string" ]
      }
   },
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

**[connectorConfigurations](#API_DescribeConnectors_ResponseSyntax)**

The configuration that is applied to the connectors used in the flow.

Type: String to [ConnectorConfiguration](api-connectorconfiguration.md) object map

Valid Keys: `Salesforce | Singular | Slack | Redshift | S3 | Marketo | Googleanalytics | Zendesk | Servicenow | Datadog | Trendmicro | Snowflake | Dynatrace | Infornexus | Amplitude | Veeva | EventBridge | LookoutMetrics | Upsolver | Honeycode | CustomerProfiles | SAPOData | CustomConnector | Pardot`

**[connectors](#API_DescribeConnectors_ResponseSyntax)**

Information about the connectors supported in Amazon AppFlow.

Type: Array of [ConnectorDetail](api-connectordetail.md) objects

**[nextToken](#API_DescribeConnectors_ResponseSyntax)**

The pagination token for the next page of data.

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

### DescribeConnectors example

This example shows a sample request and response for the
`DescribeConnectors` API using Marketo and Salesforce. The second sample
shows how to proceed if you receive `nextToken`.

#### Sample Request

```json

{
  "connectorTypes": ["Marketo","Salesforce"]
}
```

```json

{
  "connectorTypes": ["Marketo","Salesforce"],
  "nextToken": "nextToken_value"
}
```

#### Sample Response

```json

{
  "connectorConfigurations": {
    "Marketo": {
      "canUseAsDestination": false,
      "canUseAsSource": true,
      "compatibleConnectors": [
        "S3",
        "Snowflake",
        "Salesforce",
        "Redshift"
      ],
      "connectorMetadata": {
        "Amplitude": null,
        "Datadog": null,
        "Dynatrace": null,
        "GoogleAnalytics": null,
        "InforNexus": null,
        "Marketo": {},
        "Redshift": null,
        "S3": null,
        "Salesforce": null,
        "ServiceNow": null,
        "Singular": null,
        "Slack": null,
        "Snowflake": null,
        "Trendmicro": null,
        "Veeva": null,
        "Zendesk": null
      },
      "connectorSpecificConfiguration": {
        "api_version": "v1"
      },
      "credentialKeys": [
        "accessToken",
        "clientId",
        "clientSecret"
      ],
      "displayName": "Marketo",
      "hasLogoImage": false,
      "isPrivateLinkEnabled": false,
      "isPrivateLinkEndpointUrlRequired": false,
      "privateLinkEndpointServiceUrl": null,
      "supportedDestinationConnectors": [
        "S3",
        "Snowflake",
        "Salesforce",
        "Redshift"
      ],
      "supportedFrequencies": [
        "Hour",
        "Day",
        "Week",
        "Month",
        "Once"
      ],
      "supportedRegions": null,
      "supportedSchedulingFrequencies": [
        "Hour",
        "Day",
        "Week",
        "Month",
        "Once"
      ],
      "supportedTriggerTypes": [
        "Scheduled",
        "OnDemand"
      ],
      "supportedTriggers": [
        "Scheduled",
        "OnDemand"
      ]
    },
    "Salesforce": {
      "canUseAsDestination": true,
      "canUseAsSource": true,
      "compatibleConnectors": [
        "S3",
        "Snowflake",
        "Redshift",
        "Salesforce"
      ],
      "connectorMetadata": {
        "Amplitude": null,
        "Datadog": null,
        "Dynatrace": null,
        "GoogleAnalytics": null,
        "InforNexus": null,
        "Marketo": null,
        "Redshift": null,
        "S3": null,
        "Salesforce": {
          "authScopes": [
            "api",
            "refresh_token",
            "id"
          ]
        },
        "ServiceNow": null,
        "Singular": null,
        "Slack": null,
        "Snowflake": null,
        "Trendmicro": null,
        "Veeva": null,
        "Zendesk": null
      },
      "connectorSpecificConfiguration": {
        "api_version": "v47.0",
        "auth_scopes": "[api, refresh_token, id]",
        "bulk_api_path": "services/data",
        "query_api_path": "query",
        "sobject_api_path": "sobjects"
      },
      "credentialKeys": [
        "accessToken",
        "refreshToken",
        "clientId",
        "clientSecret"
      ],
      "displayName": "Salesforce",
      "hasLogoImage": false,
      "isPrivateLinkEnabled": false,
      "isPrivateLinkEndpointUrlRequired": false,
      "privateLinkEndpointServiceUrl": null,
      "supportedDestinationConnectors": [
        "S3",
        "Snowflake",
        "Redshift",
        "Salesforce"
      ],
      "supportedFrequencies": [
        "Minute",
        "Hour",
        "Day",
        "Week",
        "Month",
        "Once"
      ],
      "supportedRegions": null,
      "supportedSchedulingFrequencies": [
        "Minute",
        "Hour",
        "Day",
        "Week",
        "Month",
        "Once"
      ],
      "supportedTriggerTypes": [
        "Scheduled",
        "OnDemand",
        "Event"
      ],
      "supportedTriggers": [
        "Scheduled",
        "OnDemand",
        "Event"
      ]
    }
  },
  "nextToken": null,
  "pageToken": null
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/appflow-2020-08-23/describeconnectors.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/appflow-2020-08-23/describeconnectors.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/describeconnectors.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/appflow-2020-08-23/describeconnectors.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/describeconnectors.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/appflow-2020-08-23/describeconnectors.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/appflow-2020-08-23/describeconnectors.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/appflow-2020-08-23/describeconnectors.md)

- [AWS SDK for Python](../../../../services/goto/boto3/appflow-2020-08-23/describeconnectors.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/describeconnectors.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeConnectorProfiles

DescribeFlow

All content copied from https://docs.aws.amazon.com/.
