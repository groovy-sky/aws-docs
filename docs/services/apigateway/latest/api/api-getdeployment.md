# GetDeployment

Gets information about a Deployment resource.

## Request Syntax

```nohighlight

GET /restapis/restapi_id/deployments/deployment_id?embed=embed HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[deployment\_id](#API_GetDeployment_RequestSyntax)**

The identifier of the Deployment resource to get information about.

Required: Yes

**[embed](#API_GetDeployment_RequestSyntax)**

A query parameter to retrieve the specified embedded resources of the returned Deployment resource in the response. In a REST API call, this `embed` parameter value is a list of comma-separated strings, as in `GET /restapis/{restapi_id}/deployments/{deployment_id}?embed=var1,var2`. The SDK and other platform-dependent libraries might use a different format for the list. Currently, this request supports only retrieval of the embedded API summary this way. Hence, the parameter value must be a single-valued list containing only the `"apisummary"` string. For example, `GET /restapis/{restapi_id}/deployments/{deployment_id}?embed=apisummary`.

**[restapi\_id](#API_GetDeployment_RequestSyntax)**

The string identifier of the associated RestApi.

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "apiSummary": {
      "string" : {
         "string" : {
            "apiKeyRequired": boolean,
            "authorizationType": "string"
         }
      }
   },
   "createdDate": number,
   "description": "string",
   "id": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[apiSummary](#API_GetDeployment_ResponseSyntax)**

A summary of the RestApi at the date and time that the deployment resource was created.

Type: String to string to [MethodSnapshot](api-methodsnapshot.md) object map map

**[createdDate](#API_GetDeployment_ResponseSyntax)**

The date and time that the deployment resource was created.

Type: Timestamp

**[description](#API_GetDeployment_ResponseSyntax)**

The description for the deployment resource.

Type: String

**[id](#API_GetDeployment_ResponseSyntax)**

The identifier for the deployment resource.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**BadRequestException**

The submitted request is not valid, for example, the input is incomplete or incorrect. See the accompanying error message for details.

HTTP Status Code: 400

**NotFoundException**

The requested resource is not found. Make sure that the request URI is correct.

HTTP Status Code: 404

**ServiceUnavailableException**

The requested service is not available. For details see the accompanying error message. Retry after the specified time period.

HTTP Status Code: 503

**TooManyRequestsException**

The request has reached its throttling limit. Retry after the specified time period.

HTTP Status Code: 429

**UnauthorizedException**

The request is denied because the caller has insufficient permissions.

HTTP Status Code: 401

## Examples

### Retrieve a deployment

This example illustrates one usage of GetDeployment.

#### Sample Request

```

GET /restapis/{restapi_id}/deployments/{deployment-id}?embed=apisummary HTTP/1.1
Content-Type: application/json
Host: apigateway.us-east-1.amazonaws.com
X-Amz-Date: 20160520T055303Z
Authorization: AWS4-HMAC-SHA256 Credential={access_key_ID}/20160520/us-east-1/apigateway/aws4_request, SignedHeaders=content-type;host;x-amz-date, Signature={hash}
Cache-Control: no-cache
```

#### Sample Response

```

{
"_links": {
  "curies": {
    "href": "https://docs.aws.amazon.com/apigateway/latest/developerguide/restapi-deployment-{rel}.html",
    "name": "deployment",
    "templated": true
  },
  "self": {
    "href": "/restapis/{restapi-id}/deployments/{deployment-id}"
  },
  "deployment:delete": {
    "href": "/restapis/{restapi-id}/deployments/{deployment-id}"
  },
  "deployment:stages": {
    "href": "/restapis/{restapi-id}/stages?deployment_id={deployment-id}"
  },
  "deployment:update": {
    "href": "/restapis/{restapi-id}/deployments/{deployment-id}"
  }
},
"apiSummary": {
  "/petstorewalkthrough/pets/{petId}": {
    "GET": {
      "apiKeyRequired": false,
      "authorizationType": "NONE"
    }
  },
  "/mydemoawsproxy": {
    "GET": {
      "apiKeyRequired": false,
      "authorizationType": "NONE"
    }
  }
},
"createdDate": "2016-02-12T22:20:25Z",
"id": "{deployment-id}"
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/apigateway-2015-07-09/getdeployment.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apigateway-2015-07-09/getdeployment.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigateway-2015-07-09/getdeployment.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apigateway-2015-07-09/getdeployment.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigateway-2015-07-09/getdeployment.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apigateway-2015-07-09/getdeployment.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apigateway-2015-07-09/getdeployment.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apigateway-2015-07-09/getdeployment.md)

- [AWS SDK for Python](../../../goto/boto3/apigateway-2015-07-09/getdeployment.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigateway-2015-07-09/getdeployment.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetClientCertificates

GetDeployments

All content copied from https://docs.aws.amazon.com/.
