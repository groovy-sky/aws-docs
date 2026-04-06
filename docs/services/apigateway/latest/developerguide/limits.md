# Amazon API Gateway quotas

The following quotas apply for all Amazon API Gateway API types.

## API Gateway account-level quotas, per Region

The following quotas apply per account, per Region in Amazon API Gateway.

Resource or operationDefault quotaCan be increasedThrottle quota per account, per Region across HTTP APIs, REST APIs, WebSocket APIs, and WebSocket
callback APIs10,000 requests per second (RPS) with an additional burst capacity provided by the [token bucket algorithm](https://en.wikipedia.org/wiki/Token_bucket), using a maximum bucket
capacity of 5,000 requests. \*

###### Note

The burst quota is determined by the API Gateway service team based on the overall RPS quota for the
account in the Region. It is not a quota that a customer can control or request changes to.

[Yes](https://console.aws.amazon.com/servicequotas/home/services/apigateway/quotas/L-8A5B8E43)Throttle quota without access control per account per Region for a portal250,000 requests per secondNoThrottle quota with access control per account per Region for a portal10,000 requests per secondNo

\\* For the following Regions, the default throttle quota is 2500 RPS and the default burst quota is 1250 RPS:
Africa (Cape Town), Europe (Milan), Asia Pacific (Jakarta), Middle East (UAE), Asia Pacific (Hyderabad),
Asia Pacific (Melbourne), Europe (Spain), Europe (Zurich), Israel (Tel Aviv),
Canada West (Calgary), Asia Pacific (Malaysia), Asia Pacific (Thailand), and Mexico (Central).

## API Gateway quotas for creating, deploying and managing an API

The following fixed quotas apply to creating, deploying, and managing an API in API Gateway,
using the AWS CLI, the API Gateway console, or the API Gateway REST API and its SDKs. These quotas can't
be increased.

ActionDefault quotaCan be increased[CreateApiKey](https://docs.aws.amazon.com/apigateway/latest/api/API_CreateApiKey.html)5 requests per second per accountNo[CreateDeployment](https://docs.aws.amazon.com/apigateway/latest/api/API_CreateDeployment.html)1 request every 5 seconds per accountNo[CreateDocumentationVersion](https://docs.aws.amazon.com/apigateway/latest/api/API_CreateDocumentationVersion.html)1 request every 20 seconds per accountNo[CreateDomainName](https://docs.aws.amazon.com/apigateway/latest/api/API_CreateDomainName.html)1 request every 30 seconds per accountNo[CreateResource](https://docs.aws.amazon.com/apigateway/latest/api/API_CreateResource.html)5 requests per second per accountNo[CreateRestApi](https://docs.aws.amazon.com/apigateway/latest/api/API_CreateRestApi.html) for Regional or private API1 request every 3 seconds per accountNo[CreateRestApi](https://docs.aws.amazon.com/apigateway/latest/api/API_CreateRestApi.html) for edge-optimized API1 request every 30 seconds per accountNo[CreateVpcLink](https://docs.aws.amazon.com/apigatewayv2/latest/api-reference/vpclinks.html#CreateVpcLink) (V2)
1 request every 15 seconds per accountNo[DeleteApiKey](https://docs.aws.amazon.com/apigateway/latest/api/API_DeleteApiKey.html)5 requests per second per accountNo[DeleteDomainName](https://docs.aws.amazon.com/apigateway/latest/api/API_DeleteDomainName.html)1 request every 30 seconds per accountNo[DeleteResource](https://docs.aws.amazon.com/apigateway/latest/api/API_DeleteResource.html)5 requests per second per accountNo[DeleteRestApi](https://docs.aws.amazon.com/apigateway/latest/api/API_DeleteRestApi.html)1 request every 30 seconds per accountNo[GetResources](https://docs.aws.amazon.com/apigateway/latest/api/API_GetResources.html)5 requests every 2 seconds per accountNo[DeleteVpcLink](https://docs.aws.amazon.com/apigatewayv2/latest/api-reference/vpclinks.html#DeleteVpcLink) (V2)
1 request every 30 seconds per accountNo[ImportDocumentationParts](https://docs.aws.amazon.com/apigateway/latest/api/API_ImportDocumentationParts.html)1 request every 30 seconds per accountNo[ImportRestApi](https://docs.aws.amazon.com/apigateway/latest/api/API_ImportRestApi.html) for Regional or private API1 request every 3 seconds per accountNo[ImportRestApi](https://docs.aws.amazon.com/apigateway/latest/api/API_ImportRestApi.html) for edge-optimized API 1 request every 30 seconds per accountNo[PutRestApi](https://docs.aws.amazon.com/apigateway/latest/api/API_PutRestApi.html)1 request per second per accountNo[UpdateAccount](https://docs.aws.amazon.com/apigateway/latest/api/API_UpdateAccount.html)1 request every 20 seconds per accountNo[UpdateDomainName](https://docs.aws.amazon.com/apigateway/latest/api/API_UpdateDomainName.html)1 request every 30 seconds per accountNo[UpdateUsagePlan](https://docs.aws.amazon.com/apigateway/latest/api/API_UpdateUsagePlan.html)1 request every 20 seconds per accountNoCreate Portal1 request every 3 secondsNoUpdate Portal2 requests per minuteNoGet Portal10 requests per secondNoList Portals10 requests per secondNoPublish Portal2 requests per minuteNoDeletePortal2 requests per minuteNoPreviewPortal1 request every 3 secondsNoDisablePortal2 requests per minuteNoGetPortalProduct10 requests per secondNoListPortalProduct5 requests per secondNoCreatePortalProduct2 requests per secondNoUpdatePortalProduct0.5 requests per secondNoDeletePortalProduct1 request per secondNoPutPortalProductSharingPolicy1 request every 3 secondsNoGetPortalProductSharingPolicy10 requests per secondNoDeletePortalProductSharingPolicy1 request every 3 secondsNoCreateProductRestEndpointPage1 request every 3 secondsNoGetProductRestEndpointPage10 requests per secondNoUpdateProductRestEndpointPage1 request every 3 secondsNoDeleteProductRestEndpointPage1 request every 3 secondsNoListProductRestEndpointPages10 requests per secondNoCreateProductPage1 request every 3 secondsNoGetProductPage10 requests per secondNoUpdateProductPage1 request every 3 secondsNoDeleteProductPage1 request every 3 secondsNoListProductPages10 requests per secondNoOther operationsNo quota up to the total account quota.NoTotal operations10 requests per second with a burst quota of 40 requests per second.No

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

API references

REST API quotas
