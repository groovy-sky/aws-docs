# DescribeConnectorEntity

Provides details regarding the entity used with the connector, with a description of the
data model for each field in that entity.

## Request Syntax

```nohighlight

POST /describe-connector-entity HTTP/1.1
Content-type: application/json

{
   "apiVersion": "string",
   "connectorEntityName": "string",
   "connectorProfileName": "string",
   "connectorType": "string"
}
```

## URI Request Parameters

The request does not use any URI parameters.

## Request Body

The request accepts the following data in JSON format.

**[apiVersion](#API_DescribeConnectorEntity_RequestSyntax)**

The version of the API that's used by the connector.

Type: String

Length Constraints: Maximum length of 256.

Pattern: `\S+`

Required: No

**[connectorEntityName](#API_DescribeConnectorEntity_RequestSyntax)**

The entity name for that connector.

Type: String

Length Constraints: Maximum length of 1024.

Pattern: `\S+`

Required: Yes

**[connectorProfileName](#API_DescribeConnectorEntity_RequestSyntax)**

The name of the connector profile. The name is unique for each
`ConnectorProfile` in the AWS account.

Type: String

Length Constraints: Maximum length of 256.

Pattern: `[\w/!@#+=.-]+`

Required: No

**[connectorType](#API_DescribeConnectorEntity_RequestSyntax)**

The type of connector application, such as Salesforce, Amplitude, and so on.

Type: String

Valid Values: `Salesforce | Singular | Slack | Redshift | S3 | Marketo | Googleanalytics | Zendesk | Servicenow | Datadog | Trendmicro | Snowflake | Dynatrace | Infornexus | Amplitude | Veeva | EventBridge | LookoutMetrics | Upsolver | Honeycode | CustomerProfiles | SAPOData | CustomConnector | Pardot`

Required: No

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "connectorEntityFields": [
      {
         "customProperties": {
            "string" : "string"
         },
         "defaultValue": "string",
         "description": "string",
         "destinationProperties": {
            "isCreatable": boolean,
            "isDefaultedOnCreate": boolean,
            "isNullable": boolean,
            "isUpdatable": boolean,
            "isUpsertable": boolean,
            "supportedWriteOperations": [ "string" ]
         },
         "identifier": "string",
         "isDeprecated": boolean,
         "isPrimaryKey": boolean,
         "label": "string",
         "parentIdentifier": "string",
         "sourceProperties": {
            "isQueryable": boolean,
            "isRetrievable": boolean,
            "isTimestampFieldForIncrementalQueries": boolean
         },
         "supportedFieldTypeDetails": {
            "v1": {
               "fieldLengthRange": {
                  "maximum": number,
                  "minimum": number
               },
               "fieldType": "string",
               "fieldValueRange": {
                  "maximum": number,
                  "minimum": number
               },
               "filterOperators": [ "string" ],
               "supportedDateFormat": "string",
               "supportedValues": [ "string" ],
               "valueRegexPattern": "string"
            }
         }
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[connectorEntityFields](#API_DescribeConnectorEntity_ResponseSyntax)**

Describes the fields for that connector entity. For example, for an
_account_ entity, the fields would be _account name_,
_account ID_, and so on.

Type: Array of [ConnectorEntityField](api-connectorentityfield.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

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

**ValidationException**

The request has invalid or missing parameters.

HTTP Status Code: 400

## Examples

### DescribeConnectorEntity example

This example shows a sample request and response for the
`DescribeConnectorEntity` API

#### Sample Request

```json

{
  "connectorEntityName": "tickets",
  "connectorProfileName": "connector_profile_name",
  "connectorType": "Zendesk"
}
```

#### Sample Response

```json

{
  "connectorEntityFields": [
    {
      "description": "Account ID",
      "destinationProperties": {
        "isCreatable": false,
        "isNullable": false,
        "isUpdatable": false,
        "isUpsertable": false,
        "supportedWriteOperations": [
          "UPDATE",
          "UPSERT"
        ]
      },
      "identifier": "Id",
      "label": "Account ID",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "id",
          "filterOperators": [
            "EQUAL_TO"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "Deleted",
      "destinationProperties": {
        "isCreatable": false,
        "isNullable": false,
        "isUpdatable": false,
        "isUpsertable": false,
        "supportedWriteOperations": []
      },
      "identifier": "IsDeleted",
      "label": "Deleted",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "boolean",
          "filterOperators": [
            "EQUAL_TO"
          ],
          "supportedValues": [
            "true",
            "false"
          ]
        }
      }
    },
    {
      "description": "Master Record ID",
      "destinationProperties": {
        "isCreatable": false,
        "isNullable": true,
        "isUpdatable": false,
        "isUpsertable": false,
        "supportedWriteOperations": []
      },
      "identifier": "MasterRecordId",
      "label": "Master Record ID",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "reference",
          "filterOperators": [
            "EQUAL_TO"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "Account Name",
      "destinationProperties": {
        "isCreatable": true,
        "isNullable": false,
        "isUpdatable": true,
        "isUpsertable": true,
        "supportedWriteOperations": []
      },
      "identifier": "Name",
      "label": "Account Name",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "string",
          "filterOperators": [
            "CONTAINS",
            "EQUAL_TO",
            "NOT_EQUAL_TO"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "Account Type",
      "destinationProperties": {
        "isCreatable": true,
        "isNullable": true,
        "isUpdatable": true,
        "isUpsertable": true,
        "supportedWriteOperations": []
      },
      "identifier": "Type",
      "label": "Account Type",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "picklist",
          "filterOperators": [
            "EQUAL_TO"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "Parent Account ID",
      "destinationProperties": {
        "isCreatable": true,
        "isNullable": true,
        "isUpdatable": true,
        "isUpsertable": true,
        "supportedWriteOperations": []
      },
      "identifier": "ParentId",
      "label": "Parent Account ID",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "reference",
          "filterOperators": [
            "EQUAL_TO"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "Billing Street",
      "destinationProperties": {
        "isCreatable": true,
        "isNullable": true,
        "isUpdatable": true,
        "isUpsertable": true,
        "supportedWriteOperations": []
      },
      "identifier": "BillingStreet",
      "label": "Billing Street",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "textarea",
          "filterOperators": [
            "EQUAL_TO"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "Billing City",
      "destinationProperties": {
        "isCreatable": true,
        "isNullable": true,
        "isUpdatable": true,
        "isUpsertable": true,
        "supportedWriteOperations": []
      },
      "identifier": "BillingCity",
      "label": "Billing City",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "string",
          "filterOperators": [
            "CONTAINS",
            "EQUAL_TO",
            "NOT_EQUAL_TO"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "Billing State/Province",
      "destinationProperties": {
        "isCreatable": true,
        "isNullable": true,
        "isUpdatable": true,
        "isUpsertable": true,
        "supportedWriteOperations": []
      },
      "identifier": "BillingState",
      "label": "Billing State/Province",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "string",
          "filterOperators": [
            "CONTAINS",
            "EQUAL_TO",
            "NOT_EQUAL_TO"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "Billing Zip/Postal Code",
      "destinationProperties": {
        "isCreatable": true,
        "isNullable": true,
        "isUpdatable": true,
        "isUpsertable": true,
        "supportedWriteOperations": []
      },
      "identifier": "BillingPostalCode",
      "label": "Billing Zip/Postal Code",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "string",
          "filterOperators": [
            "CONTAINS",
            "EQUAL_TO",
            "NOT_EQUAL_TO"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "Billing Country",
      "destinationProperties": {
        "isCreatable": true,
        "isNullable": true,
        "isUpdatable": true,
        "isUpsertable": true,
        "supportedWriteOperations": []
      },
      "identifier": "BillingCountry",
      "label": "Billing Country",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "string",
          "filterOperators": [
            "CONTAINS",
            "EQUAL_TO",
            "NOT_EQUAL_TO"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "Billing Latitude",
      "destinationProperties": {
        "isCreatable": true,
        "isNullable": true,
        "isUpdatable": true,
        "isUpsertable": true,
        "supportedWriteOperations": []
      },
      "identifier": "BillingLatitude",
      "label": "Billing Latitude",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "double",
          "filterOperators": [
            "NOT_EQUAL_TO",
            "EQUAL_TO",
            "LESS_THAN",
            "LESS_THAN_OR_EQUAL_TO",
            "GREATER_THAN",
            "GREATER_THAN_OR_EQUAL_TO"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "Billing Longitude",
      "destinationProperties": {
        "isCreatable": true,
        "isNullable": true,
        "isUpdatable": true,
        "isUpsertable": true,
        "supportedWriteOperations": []
      },
      "identifier": "BillingLongitude",
      "label": "Billing Longitude",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "double",
          "filterOperators": [
            "NOT_EQUAL_TO",
            "EQUAL_TO",
            "LESS_THAN",
            "LESS_THAN_OR_EQUAL_TO",
            "GREATER_THAN",
            "GREATER_THAN_OR_EQUAL_TO"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "Billing Geocode Accuracy",
      "destinationProperties": {
        "isCreatable": true,
        "isNullable": true,
        "isUpdatable": true,
        "isUpsertable": true,
        "supportedWriteOperations": []
      },
      "identifier": "BillingGeocodeAccuracy",
      "label": "Billing Geocode Accuracy",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "picklist",
          "filterOperators": [
            "EQUAL_TO"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "Billing Address",
      "destinationProperties": {
        "isCreatable": false,
        "isNullable": true,
        "isUpdatable": false,
        "isUpsertable": false,
        "supportedWriteOperations": []
      },
      "identifier": "BillingAddress",
      "label": "Billing Address",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "address",
          "filterOperators": [
            "EQUAL_TO"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "Shipping Street",
      "destinationProperties": {
        "isCreatable": true,
        "isNullable": true,
        "isUpdatable": true,
        "isUpsertable": true,
        "supportedWriteOperations": []
      },
      "identifier": "ShippingStreet",
      "label": "Shipping Street",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "textarea",
          "filterOperators": [
            "EQUAL_TO"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "Shipping City",
      "destinationProperties": {
        "isCreatable": true,
        "isNullable": true,
        "isUpdatable": true,
        "isUpsertable": true,
        "supportedWriteOperations": []
      },
      "identifier": "ShippingCity",
      "label": "Shipping City",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "string",
          "filterOperators": [
            "CONTAINS",
            "EQUAL_TO",
            "NOT_EQUAL_TO"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "Shipping State/Province",
      "destinationProperties": {
        "isCreatable": true,
        "isNullable": true,
        "isUpdatable": true,
        "isUpsertable": true,
        "supportedWriteOperations": []
      },
      "identifier": "ShippingState",
      "label": "Shipping State/Province",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "string",
          "filterOperators": [
            "CONTAINS",
            "EQUAL_TO",
            "NOT_EQUAL_TO"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "Shipping Zip/Postal Code",
      "destinationProperties": {
        "isCreatable": true,
        "isNullable": true,
        "isUpdatable": true,
        "isUpsertable": true,
        "supportedWriteOperations": []
      },
      "identifier": "ShippingPostalCode",
      "label": "Shipping Zip/Postal Code",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "string",
          "filterOperators": [
            "CONTAINS",
            "EQUAL_TO",
            "NOT_EQUAL_TO"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "Shipping Country",
      "destinationProperties": {
        "isCreatable": true,
        "isNullable": true,
        "isUpdatable": true,
        "isUpsertable": true,
        "supportedWriteOperations": []
      },
      "identifier": "ShippingCountry",
      "label": "Shipping Country",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "string",
          "filterOperators": [
            "CONTAINS",
            "EQUAL_TO",
            "NOT_EQUAL_TO"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "Shipping Latitude",
      "destinationProperties": {
        "isCreatable": true,
        "isNullable": true,
        "isUpdatable": true,
        "isUpsertable": true,
        "supportedWriteOperations": []
      },
      "identifier": "ShippingLatitude",
      "label": "Shipping Latitude",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "double",
          "filterOperators": [
            "NOT_EQUAL_TO",
            "EQUAL_TO",
            "LESS_THAN",
            "LESS_THAN_OR_EQUAL_TO",
            "GREATER_THAN",
            "GREATER_THAN_OR_EQUAL_TO"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "Shipping Longitude",
      "destinationProperties": {
        "isCreatable": true,
        "isNullable": true,
        "isUpdatable": true,
        "isUpsertable": true,
        "supportedWriteOperations": []
      },
      "identifier": "ShippingLongitude",
      "label": "Shipping Longitude",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "double",
          "filterOperators": [
            "NOT_EQUAL_TO",
            "EQUAL_TO",
            "LESS_THAN",
            "LESS_THAN_OR_EQUAL_TO",
            "GREATER_THAN",
            "GREATER_THAN_OR_EQUAL_TO"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "Shipping Geocode Accuracy",
      "destinationProperties": {
        "isCreatable": true,
        "isNullable": true,
        "isUpdatable": true,
        "isUpsertable": true,
        "supportedWriteOperations": []
      },
      "identifier": "ShippingGeocodeAccuracy",
      "label": "Shipping Geocode Accuracy",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "picklist",
          "filterOperators": [
            "EQUAL_TO"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "Shipping Address",
      "destinationProperties": {
        "isCreatable": false,
        "isNullable": true,
        "isUpdatable": false,
        "isUpsertable": false,
        "supportedWriteOperations": []
      },
      "identifier": "ShippingAddress",
      "label": "Shipping Address",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "address",
          "filterOperators": [
            "EQUAL_TO"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "Account Phone",
      "destinationProperties": {
        "isCreatable": true,
        "isNullable": true,
        "isUpdatable": true,
        "isUpsertable": true,
        "supportedWriteOperations": []
      },
      "identifier": "Phone",
      "label": "Account Phone",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "phone",
          "filterOperators": [
            "EQUAL_TO"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "Account Fax",
      "destinationProperties": {
        "isCreatable": true,
        "isNullable": true,
        "isUpdatable": true,
        "isUpsertable": true,
        "supportedWriteOperations": []
      },
      "identifier": "Fax",
      "label": "Account Fax",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "phone",
          "filterOperators": [
            "EQUAL_TO"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "Account Number",
      "destinationProperties": {
        "isCreatable": true,
        "isNullable": true,
        "isUpdatable": true,
        "isUpsertable": true,
        "supportedWriteOperations": []
      },
      "identifier": "AccountNumber",
      "label": "Account Number",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "string",
          "filterOperators": [
            "CONTAINS",
            "EQUAL_TO",
            "NOT_EQUAL_TO"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "Website",
      "destinationProperties": {
        "isCreatable": true,
        "isNullable": true,
        "isUpdatable": true,
        "isUpsertable": true,
        "supportedWriteOperations": []
      },
      "identifier": "Website",
      "label": "Website",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "url",
          "filterOperators": [
            "EQUAL_TO"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "Photo URL",
      "destinationProperties": {
        "isCreatable": false,
        "isNullable": true,
        "isUpdatable": false,
        "isUpsertable": false,
        "supportedWriteOperations": []
      },
      "identifier": "PhotoUrl",
      "label": "Photo URL",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "url",
          "filterOperators": [
            "EQUAL_TO"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "SIC Code",
      "destinationProperties": {
        "isCreatable": true,
        "isNullable": true,
        "isUpdatable": true,
        "isUpsertable": true,
        "supportedWriteOperations": []
      },
      "identifier": "Sic",
      "label": "SIC Code",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "string",
          "filterOperators": [
            "CONTAINS",
            "EQUAL_TO",
            "NOT_EQUAL_TO"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "Industry",
      "destinationProperties": {
        "isCreatable": true,
        "isNullable": true,
        "isUpdatable": true,
        "isUpsertable": true,
        "supportedWriteOperations": []
      },
      "identifier": "Industry",
      "label": "Industry",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "picklist",
          "filterOperators": [
            "EQUAL_TO"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "Annual Revenue",
      "destinationProperties": {
        "isCreatable": true,
        "isNullable": true,
        "isUpdatable": true,
        "isUpsertable": true,
        "supportedWriteOperations": []
      },
      "identifier": "AnnualRevenue",
      "label": "Annual Revenue",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "currency",
          "filterOperators": [
            "EQUAL_TO"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "Employees",
      "destinationProperties": {
        "isCreatable": true,
        "isNullable": true,
        "isUpdatable": true,
        "isUpsertable": true,
        "supportedWriteOperations": []
      },
      "identifier": "NumberOfEmployees",
      "label": "Employees",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "int",
          "filterOperators": [
            "NOT_EQUAL_TO",
            "EQUAL_TO",
            "LESS_THAN",
            "LESS_THAN_OR_EQUAL_TO",
            "GREATER_THAN",
            "GREATER_THAN_OR_EQUAL_TO"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "Ownership",
      "destinationProperties": {
        "isCreatable": true,
        "isNullable": true,
        "isUpdatable": true,
        "isUpsertable": true,
        "supportedWriteOperations": []
      },
      "identifier": "Ownership",
      "label": "Ownership",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "picklist",
          "filterOperators": [
            "EQUAL_TO"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "Ticker Symbol",
      "destinationProperties": {
        "isCreatable": true,
        "isNullable": true,
        "isUpdatable": true,
        "isUpsertable": true,
        "supportedWriteOperations": []
      },
      "identifier": "TickerSymbol",
      "label": "Ticker Symbol",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "string",
          "filterOperators": [
            "CONTAINS",
            "EQUAL_TO",
            "NOT_EQUAL_TO"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "Account Description",
      "destinationProperties": {
        "isCreatable": true,
        "isNullable": true,
        "isUpdatable": true,
        "isUpsertable": true,
        "supportedWriteOperations": []
      },
      "identifier": "Description",
      "label": "Account Description",
      "sourceProperties": {
        "isQueryable": false,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "textarea",
          "filterOperators": [
            "EQUAL_TO"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "Account Rating",
      "destinationProperties": {
        "isCreatable": true,
        "isNullable": true,
        "isUpdatable": true,
        "isUpsertable": true,
        "supportedWriteOperations": []
      },
      "identifier": "Rating",
      "label": "Account Rating",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "picklist",
          "filterOperators": [
            "EQUAL_TO"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "Account Site",
      "destinationProperties": {
        "isCreatable": true,
        "isNullable": true,
        "isUpdatable": true,
        "isUpsertable": true,
        "supportedWriteOperations": []
      },
      "identifier": "Site",
      "label": "Account Site",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "string",
          "filterOperators": [
            "CONTAINS",
            "EQUAL_TO",
            "NOT_EQUAL_TO"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "Owner ID",
      "destinationProperties": {
        "isCreatable": true,
        "isNullable": false,
        "isUpdatable": true,
        "isUpsertable": true,
        "supportedWriteOperations": []
      },
      "identifier": "OwnerId",
      "label": "Owner ID",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "reference",
          "filterOperators": [
            "EQUAL_TO"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "Created Date",
      "destinationProperties": {
        "isCreatable": false,
        "isNullable": false,
        "isUpdatable": false,
        "isUpsertable": false,
        "supportedWriteOperations": []
      },
      "identifier": "CreatedDate",
      "label": "Created Date",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "datetime",
          "filterOperators": [
            "EQUAL_TO",
            "LESS_THAN",
            "LESS_THAN_OR_EQUAL_TO",
            "GREATER_THAN",
            "GREATER_THAN_OR_EQUAL_TO",
            "BETWEEN"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "Created By ID",
      "destinationProperties": {
        "isCreatable": false,
        "isNullable": false,
        "isUpdatable": false,
        "isUpsertable": false,
        "supportedWriteOperations": []
      },
      "identifier": "CreatedById",
      "label": "Created By ID",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "reference",
          "filterOperators": [
            "EQUAL_TO"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "Last Modified Date",
      "destinationProperties": {
        "isCreatable": false,
        "isNullable": false,
        "isUpdatable": false,
        "isUpsertable": false,
        "supportedWriteOperations": []
      },
      "identifier": "LastModifiedDate",
      "label": "Last Modified Date",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "datetime",
          "filterOperators": [
            "EQUAL_TO",
            "LESS_THAN",
            "LESS_THAN_OR_EQUAL_TO",
            "GREATER_THAN",
            "GREATER_THAN_OR_EQUAL_TO",
            "BETWEEN"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "Last Modified By ID",
      "destinationProperties": {
        "isCreatable": false,
        "isNullable": false,
        "isUpdatable": false,
        "isUpsertable": false,
        "supportedWriteOperations": []
      },
      "identifier": "LastModifiedById",
      "label": "Last Modified By ID",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "reference",
          "filterOperators": [
            "EQUAL_TO"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "System Modstamp",
      "destinationProperties": {
        "isCreatable": false,
        "isNullable": false,
        "isUpdatable": false,
        "isUpsertable": false,
        "supportedWriteOperations": []
      },
      "identifier": "SystemModstamp",
      "label": "System Modstamp",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "datetime",
          "filterOperators": [
            "EQUAL_TO",
            "LESS_THAN",
            "LESS_THAN_OR_EQUAL_TO",
            "GREATER_THAN",
            "GREATER_THAN_OR_EQUAL_TO",
            "BETWEEN"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "Last Activity",
      "destinationProperties": {
        "isCreatable": false,
        "isNullable": true,
        "isUpdatable": false,
        "isUpsertable": false,
        "supportedWriteOperations": []
      },
      "identifier": "LastActivityDate",
      "label": "Last Activity",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "date",
          "filterOperators": [
            "EQUAL_TO",
            "LESS_THAN",
            "LESS_THAN_OR_EQUAL_TO",
            "GREATER_THAN",
            "GREATER_THAN_OR_EQUAL_TO",
            "BETWEEN"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "Last Viewed Date",
      "destinationProperties": {
        "isCreatable": false,
        "isNullable": true,
        "isUpdatable": false,
        "isUpsertable": false,
        "supportedWriteOperations": []
      },
      "identifier": "LastViewedDate",
      "label": "Last Viewed Date",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "datetime",
          "filterOperators": [
            "EQUAL_TO",
            "LESS_THAN",
            "LESS_THAN_OR_EQUAL_TO",
            "GREATER_THAN",
            "GREATER_THAN_OR_EQUAL_TO",
            "BETWEEN"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "Last Referenced Date",
      "destinationProperties": {
        "isCreatable": false,
        "isNullable": true,
        "isUpdatable": false,
        "isUpsertable": false,
        "supportedWriteOperations": []
      },
      "identifier": "LastReferencedDate",
      "label": "Last Referenced Date",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "datetime",
          "filterOperators": [
            "EQUAL_TO",
            "LESS_THAN",
            "LESS_THAN_OR_EQUAL_TO",
            "GREATER_THAN",
            "GREATER_THAN_OR_EQUAL_TO",
            "BETWEEN"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "Data.com Key",
      "destinationProperties": {
        "isCreatable": true,
        "isNullable": true,
        "isUpdatable": true,
        "isUpsertable": true,
        "supportedWriteOperations": []
      },
      "identifier": "Jigsaw",
      "label": "Data.com Key",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "string",
          "filterOperators": [
            "CONTAINS",
            "EQUAL_TO",
            "NOT_EQUAL_TO"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "Jigsaw Company ID",
      "destinationProperties": {
        "isCreatable": false,
        "isNullable": true,
        "isUpdatable": false,
        "isUpsertable": false,
        "supportedWriteOperations": []
      },
      "identifier": "JigsawCompanyId",
      "label": "Jigsaw Company ID",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "string",
          "filterOperators": [
            "CONTAINS",
            "EQUAL_TO",
            "NOT_EQUAL_TO"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "Clean Status",
      "destinationProperties": {
        "isCreatable": true,
        "isNullable": true,
        "isUpdatable": true,
        "isUpsertable": true,
        "supportedWriteOperations": []
      },
      "identifier": "CleanStatus",
      "label": "Clean Status",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "picklist",
          "filterOperators": [
            "EQUAL_TO"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "Account Source",
      "destinationProperties": {
        "isCreatable": true,
        "isNullable": true,
        "isUpdatable": true,
        "isUpsertable": true,
        "supportedWriteOperations": []
      },
      "identifier": "AccountSource",
      "label": "Account Source",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "picklist",
          "filterOperators": [
            "EQUAL_TO"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "D-U-N-S Number",
      "destinationProperties": {
        "isCreatable": true,
        "isNullable": true,
        "isUpdatable": true,
        "isUpsertable": true,
        "supportedWriteOperations": []
      },
      "identifier": "DunsNumber",
      "label": "D-U-N-S Number",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "string",
          "filterOperators": [
            "CONTAINS",
            "EQUAL_TO",
            "NOT_EQUAL_TO"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "Tradestyle",
      "destinationProperties": {
        "isCreatable": true,
        "isNullable": true,
        "isUpdatable": true,
        "isUpsertable": true,
        "supportedWriteOperations": []
      },
      "identifier": "Tradestyle",
      "label": "Tradestyle",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "string",
          "filterOperators": [
            "CONTAINS",
            "EQUAL_TO",
            "NOT_EQUAL_TO"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "NAICS Code",
      "destinationProperties": {
        "isCreatable": true,
        "isNullable": true,
        "isUpdatable": true,
        "isUpsertable": true,
        "supportedWriteOperations": []
      },
      "identifier": "NaicsCode",
      "label": "NAICS Code",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "string",
          "filterOperators": [
            "CONTAINS",
            "EQUAL_TO",
            "NOT_EQUAL_TO"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "NAICS Description",
      "destinationProperties": {
        "isCreatable": true,
        "isNullable": true,
        "isUpdatable": true,
        "isUpsertable": true,
        "supportedWriteOperations": []
      },
      "identifier": "NaicsDesc",
      "label": "NAICS Description",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "string",
          "filterOperators": [
            "CONTAINS",
            "EQUAL_TO",
            "NOT_EQUAL_TO"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "Year Started",
      "destinationProperties": {
        "isCreatable": true,
        "isNullable": true,
        "isUpdatable": true,
        "isUpsertable": true,
        "supportedWriteOperations": []
      },
      "identifier": "YearStarted",
      "label": "Year Started",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "string",
          "filterOperators": [
            "CONTAINS",
            "EQUAL_TO",
            "NOT_EQUAL_TO"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "SIC Description",
      "destinationProperties": {
        "isCreatable": true,
        "isNullable": true,
        "isUpdatable": true,
        "isUpsertable": true,
        "supportedWriteOperations": []
      },
      "identifier": "SicDesc",
      "label": "SIC Description",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "string",
          "filterOperators": [
            "CONTAINS",
            "EQUAL_TO",
            "NOT_EQUAL_TO"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "D&B Company ID",
      "destinationProperties": {
        "isCreatable": true,
        "isNullable": true,
        "isUpdatable": true,
        "isUpsertable": true,
        "supportedWriteOperations": []
      },
      "identifier": "DandbCompanyId",
      "label": "D&B Company ID",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "reference",
          "filterOperators": [
            "EQUAL_TO"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "Operating Hour ID",
      "destinationProperties": {
        "isCreatable": true,
        "isNullable": true,
        "isUpdatable": true,
        "isUpsertable": true,
        "supportedWriteOperations": []
      },
      "identifier": "OperatingHoursId",
      "label": "Operating Hour ID",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "reference",
          "filterOperators": [
            "EQUAL_TO"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "Customer Priority",
      "destinationProperties": {
        "isCreatable": true,
        "isNullable": true,
        "isUpdatable": true,
        "isUpsertable": true,
        "supportedWriteOperations": []
      },
      "identifier": "CustomerPriority__c",
      "label": "Customer Priority",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "picklist",
          "filterOperators": [
            "EQUAL_TO"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "SLA",
      "destinationProperties": {
        "isCreatable": true,
        "isNullable": true,
        "isUpdatable": true,
        "isUpsertable": true,
        "supportedWriteOperations": []
      },
      "identifier": "SLA__c",
      "label": "SLA",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "picklist",
          "filterOperators": [
            "EQUAL_TO"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "Active",
      "destinationProperties": {
        "isCreatable": true,
        "isNullable": true,
        "isUpdatable": true,
        "isUpsertable": true,
        "supportedWriteOperations": []
      },
      "identifier": "Active__c",
      "label": "Active",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "picklist",
          "filterOperators": [
            "EQUAL_TO"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "Number of Locations",
      "destinationProperties": {
        "isCreatable": true,
        "isNullable": true,
        "isUpdatable": true,
        "isUpsertable": true,
        "supportedWriteOperations": []
      },
      "identifier": "NumberofLocations__c",
      "label": "Number of Locations",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "double",
          "filterOperators": [
            "NOT_EQUAL_TO",
            "EQUAL_TO",
            "LESS_THAN",
            "LESS_THAN_OR_EQUAL_TO",
            "GREATER_THAN",
            "GREATER_THAN_OR_EQUAL_TO"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "Upsell Opportunity",
      "destinationProperties": {
        "isCreatable": true,
        "isNullable": true,
        "isUpdatable": true,
        "isUpsertable": true,
        "supportedWriteOperations": []
      },
      "identifier": "UpsellOpportunity__c",
      "label": "Upsell Opportunity",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "picklist",
          "filterOperators": [
            "EQUAL_TO"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "SLA Serial Number",
      "destinationProperties": {
        "isCreatable": true,
        "isNullable": true,
        "isUpdatable": true,
        "isUpsertable": true,
        "supportedWriteOperations": []
      },
      "identifier": "SLASerialNumber__c",
      "label": "SLA Serial Number",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "string",
          "filterOperators": [
            "CONTAINS",
            "EQUAL_TO",
            "NOT_EQUAL_TO"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "SLA Expiration Date",
      "destinationProperties": {
        "isCreatable": true,
        "isNullable": true,
        "isUpdatable": true,
        "isUpsertable": true,
        "supportedWriteOperations": []
      },
      "identifier": "SLAExpirationDate__c",
      "label": "SLA Expiration Date",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "date",
          "filterOperators": [
            "EQUAL_TO",
            "LESS_THAN",
            "LESS_THAN_OR_EQUAL_TO",
            "GREATER_THAN",
            "GREATER_THAN_OR_EQUAL_TO",
            "BETWEEN"
          ],
          "supportedValues": []
        }
      }
    },
    {
      "description": "extId",
      "destinationProperties": {
        "isCreatable": true,
        "isNullable": true,
        "isUpdatable": true,
        "isUpsertable": true,
        "supportedWriteOperations": [
          "UPSERT"
        ]
      },
      "identifier": "extId__c",
      "label": "extId",
      "sourceProperties": {
        "isQueryable": true,
        "isRetrievable": true
      },
      "supportedFieldTypeDetails": {
        "v1": {
          "fieldType": "string",
          "filterOperators": [
            "CONTAINS",
            "EQUAL_TO",
            "NOT_EQUAL_TO"
          ],
          "supportedValues": []
        }
      }
    }
  ]
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/appflow-2020-08-23/describeconnectorentity.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/appflow-2020-08-23/describeconnectorentity.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/describeconnectorentity.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/appflow-2020-08-23/describeconnectorentity.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/describeconnectorentity.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/appflow-2020-08-23/describeconnectorentity.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/appflow-2020-08-23/describeconnectorentity.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/appflow-2020-08-23/describeconnectorentity.md)

- [AWS SDK for Python](../../../../services/goto/boto3/appflow-2020-08-23/describeconnectorentity.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/describeconnectorentity.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeConnector

DescribeConnectorProfiles

All content copied from https://docs.aws.amazon.com/.
