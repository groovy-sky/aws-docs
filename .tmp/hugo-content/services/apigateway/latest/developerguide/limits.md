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

ActionDefault quotaCan be increased[CreateApiKey](../api/api-createapikey.md)5 requests per second per accountNo[CreateDeployment](../api/api-createdeployment.md)1 request every 5 seconds per accountNo[CreateDocumentationVersion](../api/api-createdocumentationversion.md)1 request every 20 seconds per accountNo[CreateDomainName](../api/api-createdomainname.md)1 request every 30 seconds per accountNo[CreateResource](../api/api-createresource.md)5 requests per second per accountNo[CreateRestApi](../api/api-createrestapi.md) for Regional or private API1 request every 3 seconds per accountNo[CreateRestApi](../api/api-createrestapi.md) for edge-optimized API1 request every 30 seconds per accountNo[CreateVpcLink](../../../apigatewayv2/latest/api-reference/vpclinks.md#CreateVpcLink) (V2)
1 request every 15 seconds per accountNo[DeleteApiKey](../api/api-deleteapikey.md)5 requests per second per accountNo[DeleteDomainName](../api/api-deletedomainname.md)1 request every 30 seconds per accountNo[DeleteResource](../api/api-deleteresource.md)5 requests per second per accountNo[DeleteRestApi](../api/api-deleterestapi.md)1 request every 30 seconds per accountNo[GetResources](../api/api-getresources.md)5 requests every 2 seconds per accountNo[DeleteVpcLink](../../../apigatewayv2/latest/api-reference/vpclinks.md#DeleteVpcLink) (V2)
1 request every 30 seconds per accountNo[ImportDocumentationParts](../api/api-importdocumentationparts.md)1 request every 30 seconds per accountNo[ImportRestApi](../api/api-importrestapi.md) for Regional or private API1 request every 3 seconds per accountNo[ImportRestApi](../api/api-importrestapi.md) for edge-optimized API 1 request every 30 seconds per accountNo[PutRestApi](../api/api-putrestapi.md)1 request per second per accountNo[UpdateAccount](../api/api-updateaccount.md)1 request every 20 seconds per accountNo[UpdateDomainName](../api/api-updatedomainname.md)1 request every 30 seconds per accountNo[UpdateUsagePlan](../api/api-updateusageplan.md)1 request every 20 seconds per accountNoCreate Portal1 request every 3 secondsNoUpdate Portal2 requests per minuteNoGet Portal10 requests per secondNoList Portals10 requests per secondNoPublish Portal2 requests per minuteNoDeletePortal2 requests per minuteNoPreviewPortal1 request every 3 secondsNoDisablePortal2 requests per minuteNoGetPortalProduct10 requests per secondNoListPortalProduct5 requests per secondNoCreatePortalProduct2 requests per secondNoUpdatePortalProduct0.5 requests per secondNoDeletePortalProduct1 request per secondNoPutPortalProductSharingPolicy1 request every 3 secondsNoGetPortalProductSharingPolicy10 requests per secondNoDeletePortalProductSharingPolicy1 request every 3 secondsNoCreateProductRestEndpointPage1 request every 3 secondsNoGetProductRestEndpointPage10 requests per secondNoUpdateProductRestEndpointPage1 request every 3 secondsNoDeleteProductRestEndpointPage1 request every 3 secondsNoListProductRestEndpointPages10 requests per secondNoCreateProductPage1 request every 3 secondsNoGetProductPage10 requests per secondNoUpdateProductPage1 request every 3 secondsNoDeleteProductPage1 request every 3 secondsNoListProductPages10 requests per secondNoOther operationsNo quota up to the total account quota.NoTotal operations10 requests per second with a burst quota of 40 requests per second.No

[Document Conventions](../../../../general/latest/gr/docconventions.md)

API references

REST API quotas

All content copied from https://docs.aws.amazon.com/.
