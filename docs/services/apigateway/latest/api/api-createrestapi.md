# CreateRestApi

Creates a new RestApi resource.

## Request Syntax

```nohighlight

POST /restapis HTTP/1.1
Content-type: application/json

{
   "apiKeySource": "string",
   "binaryMediaTypes": [ "string" ],
   "cloneFrom": "string",
   "description": "string",
   "disableExecuteApiEndpoint": boolean,
   "endpointAccessMode": "string",
   "endpointConfiguration": {
      "ipAddressType": "string",
      "types": [ "string" ],
      "vpcEndpointIds": [ "string" ]
   },
   "minimumCompressionSize": number,
   "name": "string",
   "policy": "string",
   "securityPolicy": "string",
   "tags": {
      "string" : "string"
   },
   "version": "string"
}
```

## URI Request Parameters

The request does not use any URI parameters.

## Request Body

The request accepts the following data in JSON format.

**[apiKeySource](#API_CreateRestApi_RequestSyntax)**

The source of the API key for metering requests according to a usage plan. Valid values
are: `HEADER` to read the API key from the `X-API-Key` header of a
request. `AUTHORIZER` to read the API key from the `UsageIdentifierKey`
from a custom authorizer.

Type: String

Valid Values: `HEADER | AUTHORIZER`

Required: No

**[binaryMediaTypes](#API_CreateRestApi_RequestSyntax)**

The list of binary media types supported by the RestApi. By default, the RestApi supports only UTF-8-encoded text payloads.

Type: Array of strings

Required: No

**[cloneFrom](#API_CreateRestApi_RequestSyntax)**

The ID of the RestApi that you want to clone from.

Type: String

Required: No

**[description](#API_CreateRestApi_RequestSyntax)**

The description of the RestApi.

Type: String

Required: No

**[disableExecuteApiEndpoint](#API_CreateRestApi_RequestSyntax)**

Specifies whether clients can invoke your API by using the default `execute-api` endpoint.
By default, clients can invoke your API with the default
`https://{api_id}.execute-api.{region}.amazonaws.com` endpoint. To require that clients use a
custom domain name to invoke your API, disable the default endpoint

Type: Boolean

Required: No

**[endpointAccessMode](#API_CreateRestApi_RequestSyntax)**

The endpoint access mode of the RestApi. Only available for RestApis that use security policies that start with `SecurityPolicy_`.

Type: String

Valid Values: `BASIC | STRICT`

Required: No

**[endpointConfiguration](#API_CreateRestApi_RequestSyntax)**

The endpoint configuration of this RestApi showing the endpoint types and IP address types of the API.

Type: [EndpointConfiguration](https://docs.aws.amazon.com/apigateway/latest/api/API_EndpointConfiguration.html) object

Required: No

**[minimumCompressionSize](#API_CreateRestApi_RequestSyntax)**

A nullable integer that is used to enable compression (with non-negative between 0 and 10485760 (10M) bytes, inclusive) or disable compression (with a null value) on an API. When compression is enabled, compression or decompression is not applied on the payload if the payload size is smaller than this value. Setting it to zero allows compression for any payload size.

Type: Integer

Required: No

**[name](#API_CreateRestApi_RequestSyntax)**

The name of the RestApi.

Type: String

Required: Yes

**[policy](#API_CreateRestApi_RequestSyntax)**

A stringified JSON policy document that applies to this RestApi regardless of the caller and Method configuration.

Type: String

Required: No

**[securityPolicy](#API_CreateRestApi_RequestSyntax)**

The Transport Layer Security (TLS) version + cipher suite for this RestApi.

Type: String

Valid Values: `TLS_1_0 | TLS_1_2 | SecurityPolicy_TLS13_1_3_2025_09 | SecurityPolicy_TLS13_1_3_FIPS_2025_09 | SecurityPolicy_TLS13_1_2_PFS_PQ_2025_09 | SecurityPolicy_TLS13_1_2_FIPS_PQ_2025_09 | SecurityPolicy_TLS13_1_2_PQ_2025_09 | SecurityPolicy_TLS13_1_2_2021_06 | SecurityPolicy_TLS13_2025_EDGE | SecurityPolicy_TLS12_PFS_2025_EDGE | SecurityPolicy_TLS12_2018_EDGE`

Required: No

**[tags](#API_CreateRestApi_RequestSyntax)**

The key-value map of strings. The valid character set is \[a-zA-Z+-=.\_:/\]. The tag key can be up to 128 characters and must not start with `aws:`. The tag value can be up to 256 characters.

Type: String to string map

Required: No

**[version](#API_CreateRestApi_RequestSyntax)**

A version identifier for the API.

Type: String

Required: No

## Response Syntax

```nohighlight

HTTP/1.1 201
Content-type: application/json

{
   "apiKeySource": "string",
   "apiStatus": "string",
   "apiStatusMessage": "string",
   "binaryMediaTypes": [ "string" ],
   "createdDate": number,
   "description": "string",
   "disableExecuteApiEndpoint": boolean,
   "endpointAccessMode": "string",
   "endpointConfiguration": {
      "ipAddressType": "string",
      "types": [ "string" ],
      "vpcEndpointIds": [ "string" ]
   },
   "id": "string",
   "minimumCompressionSize": number,
   "name": "string",
   "policy": "string",
   "rootResourceId": "string",
   "securityPolicy": "string",
   "tags": {
      "string" : "string"
   },
   "version": "string",
   "warnings": [ "string" ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 201 response.

The following data is returned in JSON format by the service.

**[apiKeySource](#API_CreateRestApi_ResponseSyntax)**

The source of the API key for metering requests according to a usage plan. Valid values
are: > `HEADER` to read the API key from the `X-API-Key` header of a
request. `AUTHORIZER` to read the API key from the `UsageIdentifierKey`
from a custom authorizer.

Type: String

Valid Values: `HEADER | AUTHORIZER`

**[apiStatus](#API_CreateRestApi_ResponseSyntax)**

The ApiStatus of the RestApi.

Type: String

Valid Values: `UPDATING | AVAILABLE | PENDING | FAILED`

**[apiStatusMessage](#API_CreateRestApi_ResponseSyntax)**

The status message of the RestApi. When the status message is `UPDATING` you can still invoke it.

Type: String

**[binaryMediaTypes](#API_CreateRestApi_ResponseSyntax)**

The list of binary media types supported by the RestApi. By default, the RestApi supports only UTF-8-encoded text payloads.

Type: Array of strings

**[createdDate](#API_CreateRestApi_ResponseSyntax)**

The timestamp when the API was created.

Type: Timestamp

**[description](#API_CreateRestApi_ResponseSyntax)**

The API's description.

Type: String

**[disableExecuteApiEndpoint](#API_CreateRestApi_ResponseSyntax)**

Specifies whether clients can invoke your API by using the default `execute-api` endpoint.
By default, clients can invoke your API with the default
`https://{api_id}.execute-api.{region}.amazonaws.com` endpoint. To require that clients use a
custom domain name to invoke your API, disable the default endpoint.

Type: Boolean

**[endpointAccessMode](#API_CreateRestApi_ResponseSyntax)**

The endpoint access mode of the RestApi.

Type: String

Valid Values: `BASIC | STRICT`

**[endpointConfiguration](#API_CreateRestApi_ResponseSyntax)**

The endpoint configuration of this RestApi showing the endpoint types and IP address types of the API.

Type: [EndpointConfiguration](https://docs.aws.amazon.com/apigateway/latest/api/API_EndpointConfiguration.html) object

**[id](#API_CreateRestApi_ResponseSyntax)**

The API's identifier. This identifier is unique across all of your APIs in API Gateway.

Type: String

**[minimumCompressionSize](#API_CreateRestApi_ResponseSyntax)**

A nullable integer that is used to enable compression (with non-negative between 0 and 10485760 (10M) bytes, inclusive) or disable compression (with a null value) on an API. When compression is enabled, compression or decompression is not applied on the payload if the payload size is smaller than this value. Setting it to zero allows compression for any payload size.

Type: Integer

**[name](#API_CreateRestApi_ResponseSyntax)**

The API's name.

Type: String

**[policy](#API_CreateRestApi_ResponseSyntax)**

A stringified JSON policy document that applies to this RestApi regardless of the caller and Method configuration.

Type: String

**[rootResourceId](#API_CreateRestApi_ResponseSyntax)**

The API's root resource ID.

Type: String

**[securityPolicy](#API_CreateRestApi_ResponseSyntax)**

The Transport Layer Security (TLS) version + cipher suite for this RestApi.

Type: String

Valid Values: `TLS_1_0 | TLS_1_2 | SecurityPolicy_TLS13_1_3_2025_09 | SecurityPolicy_TLS13_1_3_FIPS_2025_09 | SecurityPolicy_TLS13_1_2_PFS_PQ_2025_09 | SecurityPolicy_TLS13_1_2_FIPS_PQ_2025_09 | SecurityPolicy_TLS13_1_2_PQ_2025_09 | SecurityPolicy_TLS13_1_2_2021_06 | SecurityPolicy_TLS13_2025_EDGE | SecurityPolicy_TLS12_PFS_2025_EDGE | SecurityPolicy_TLS12_2018_EDGE`

**[tags](#API_CreateRestApi_ResponseSyntax)**

The collection of tags. Each tag element is associated with a given resource.

Type: String to string map

**[version](#API_CreateRestApi_ResponseSyntax)**

A version identifier for the API.

Type: String

**[warnings](#API_CreateRestApi_ResponseSyntax)**

The warning messages reported when `failonwarnings` is turned on during API import.

Type: Array of strings

## Errors

For information about the errors that are common to all actions, see [Common Error Types](https://docs.aws.amazon.com/apigateway/latest/api/CommonErrors.html).

**BadRequestException**

The submitted request is not valid, for example, the input is incomplete or incorrect. See the accompanying error message for details.

HTTP Status Code: 400

**ConflictException**

The request configuration has conflicts. For details, see the accompanying error message.

HTTP Status Code: 409

**LimitExceededException**

The request exceeded the rate limit. Retry after the specified time period.

HTTP Status Code: 429

**TooManyRequestsException**

The request has reached its throttling limit. Retry after the specified time period.

HTTP Status Code: 429

**UnauthorizedException**

The request is denied because the caller has insufficient permissions.

HTTP Status Code: 401

## Examples

### Create a new API

This example illustrates one usage of CreateRestApi.

#### Sample Request

```

POST /restapis HTTP/1.1
Content-Type: application/json
Host: apigateway.us-east-1.amazonaws.com
X-Amz-Date: 20160601T185340Z
Authorization: AWS4-HMAC-SHA256 Credential={access_key_ID}/20160601/us-east-1/apigateway/aws4_request, SignedHeaders=content-length;content-type;host;x-amz-date, Signature={sig4_hash}

{
    "name" : "my-test-api",
    "description" : "A sample API created programmatically by calling API Gateway control service using the REST API"
}
```

#### Sample Response

```

{
  "_links": {
    "curies": [
       ...
    ],
    "self": {
      "href": "/restapis/fugvjdxtri"
    },
    "authorizer:by-id": {
      "href": "/restapis/fugvjdxtri/authorizers/{authorizer_id}",
      "templated": true
    },
    "authorizer:create": {
      "href": "/restapis/fugvjdxtri/authorizers"
    },
    "deployment:by-id": {
      "href": "/restapis/fugvjdxtri/deployments/{deployment_id}{?embed}",
      "templated": true
    },
    "deployment:create": {
      "href": "/restapis/fugvjdxtri/deployments"
    },
    "model:by-name": {
      "href": "/restapis/fugvjdxtri/models/{model_name}?flatten=false",
      "templated": true
    },
    "model:create": {
      "href": "/restapis/fugvjdxtri/models"
    },
    "resource:by-id": {
      "href": "/restapis/fugvjdxtri/resources/{resource_id}{?embed}",
      "templated": true
    },
    "resource:create": {
      "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2"
    },
    "restapi:authorizers": {
      "href": "/restapis/fugvjdxtri/authorizers"
    },
    "restapi:delete": {
      "href": "/restapis/fugvjdxtri"
    },
    "restapi:deployments": {
      "href": "/restapis/fugvjdxtri/deployments{?limit}",
      "templated": true
    },
    "restapi:models": {
      "href": "/restapis/fugvjdxtri/models"
    },
    "restapi:resources": {
      "href": "/restapis/fugvjdxtri/resources{?limit,embed}",
      "templated": true
    },
    "restapi:stages": {
      "href": "/restapis/fugvjdxtri/stages{?deployment_id}",
      "templated": true
    },
    "restapi:update": {
      "href": "/restapis/fugvjdxtri"
    },
    "stage:by-name": {
      "href": "/restapis/fugvjdxtri/stages/{stage_name}",
      "templated": true
    },
    "stage:create": {
      "href": "/restapis/fugvjdxtri/stages"
    }
  },
  "createdDate": "2016-06-01T18:53:41Z",
  "description": "A sample API created programmatically by calling API Gateway control service using the REST API",
  "id": "fugvjdxtri",
  "name": "my-test-api"
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/apigateway-2015-07-09/CreateRestApi)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/apigateway-2015-07-09/CreateRestApi)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/apigateway-2015-07-09/CreateRestApi)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/apigateway-2015-07-09/CreateRestApi)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/apigateway-2015-07-09/CreateRestApi)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/apigateway-2015-07-09/CreateRestApi)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/apigateway-2015-07-09/CreateRestApi)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/apigateway-2015-07-09/CreateRestApi)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/apigateway-2015-07-09/CreateRestApi)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/apigateway-2015-07-09/CreateRestApi)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CreateResource

CreateStage
