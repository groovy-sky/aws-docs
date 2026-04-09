# DescribeConnector

Describes the given custom connector registered in your AWS account. This
API can be used for custom connectors that are registered in your account and also for Amazon
authored connectors.

## Request Syntax

```nohighlight

POST /describe-connector HTTP/1.1
Content-type: application/json

{
   "connectorLabel": "string",
   "connectorType": "string"
}
```

## URI Request Parameters

The request does not use any URI parameters.

## Request Body

The request accepts the following data in JSON format.

**[connectorLabel](#API_DescribeConnector_RequestSyntax)**

The label of the connector. The label is unique for each
`ConnectorRegistration` in your AWS account. Only needed if
calling for CUSTOMCONNECTOR connector type/.

Type: String

Length Constraints: Maximum length of 256.

Pattern: `[a-zA-Z0-9][\w!@#.-]+`

Required: No

**[connectorType](#API_DescribeConnector_RequestSyntax)**

The connector type, such as CUSTOMCONNECTOR, Saleforce, Marketo. Please choose
CUSTOMCONNECTOR for Lambda based custom connectors.

Type: String

Valid Values: `Salesforce | Singular | Slack | Redshift | S3 | Marketo | Googleanalytics | Zendesk | Servicenow | Datadog | Trendmicro | Snowflake | Dynatrace | Infornexus | Amplitude | Veeva | EventBridge | LookoutMetrics | Upsolver | Honeycode | CustomerProfiles | SAPOData | CustomConnector | Pardot`

Required: Yes

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "connectorConfiguration": {
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
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[connectorConfiguration](#API_DescribeConnector_ResponseSyntax)**

Configuration info of all the connectors that the user requested.

Type: [ConnectorConfiguration](api-connectorconfiguration.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

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

### Describing the given connector

This example shows a sample request for the `DescribeConnector` API and a
sample response.

#### Sample Request

```json

{
  "connectorType": "connectorType",
  "connectorLabel": "connectorLabel"
}
```

#### Sample Response

```json

{
  "connectorConfiguration":
  {
    "authenticationConfig": null,
    "canUseAsDestination": true,
    "canUseAsSource": true,
    "connectorArn": null,
    "connectorDescription": "Salesforce",
    "connectorLabel": null,
    "connectorMetadata":
    {
      "Amplitude": null,
      "AppIntegrations": null,
      "CustomerProfiles": null,
      "Datadog": null,
      "Dynatrace": null,
      "EventBridge": null,
      "GoogleAnalytics": null,
      "Honeycode": null,
      "InforNexus": null,
      "Locke": null,
      "Marketo": null,
      "Pardot": null,
      "Redshift": null,
      "S3": null,
      "SAPOData": null,
      "Salesforce":
      {
        "oAuthScopes":
        [
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
      "Upsolver": null,
      "Veeva": null,
      "Zendesk": null
    },
    "connectorModes": null,
    "connectorName": null,
    "connectorOwner": null,
    "connectorProvisioningConfig": null,
    "connectorProvisioningType": null,
    "connectorRuntimeSettings": null,
    "connectorType": "SALESFORCE",
    "connectorVersion": null,
    "isPrivateLinkEnabled": true,
    "isPrivateLinkEndpointUrlRequired": false,
    "logoURL": "insert-logo-here",
    "registeredAt": null,
    "registeredBy": null,
    "supportedApiVersions": null,
    "supportedDestinationConnectors":
    [
      "S3",
      "Snowflake",
      "Redshift",
      "Salesforce",
      "EventBridge",
      "LookoutMetrics",
      "Locke",
      "Honeycode",
      "Upsolver",
      "CustomerProfiles",
      "Zendesk",
      "Marketo",
      "CustomConnector"
    ],
    "supportedOperators": null,
    "supportedSchedulingFrequencies":
    [
      "BYMINUTE",
      "HOURLY",
      "DAILY",
      "WEEKLY",
      "MONTHLY",
      "ONCE"
    ],
    "supportedTriggerTypes":
    [
      "Scheduled",
      "OnDemand",
      "Event"
    ],
    "supportedWriteOperations": null
  }
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/appflow-2020-08-23/describeconnector.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/appflow-2020-08-23/describeconnector.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/describeconnector.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/appflow-2020-08-23/describeconnector.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/describeconnector.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/appflow-2020-08-23/describeconnector.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/appflow-2020-08-23/describeconnector.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/appflow-2020-08-23/describeconnector.md)

- [AWS SDK for Python](../../../../services/goto/boto3/appflow-2020-08-23/describeconnector.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/describeconnector.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteFlow

DescribeConnectorEntity

All content copied from https://docs.aws.amazon.com/.
