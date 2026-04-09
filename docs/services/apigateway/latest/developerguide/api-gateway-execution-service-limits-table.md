# Quotas for configuring and running a REST API in API Gateway

The following quotas apply to configuring and running a REST API in Amazon API Gateway. If the quota is per-API, it can
only be increased on specific APIs, not for all the APIs in an account.

For information about account-level quotas
see [Amazon API Gateway quotas](limits.md).

Resource or operationDefault quotaCan be increasedPublic custom domain names per account per Region120[Yes](https://console.aws.amazon.com/servicequotas/home/services/apigateway/quotas/L-A93447B8)Private custom domain names per account per Region50[Yes](https://console.aws.amazon.com/servicequotas/home/services/apigateway/quotas/L-24E7E662)Regional APIs per Region600NoEdge-optimized APIs per Region120NoPrivate APIs per account per Region600NoDomain name access associations per account100[Yes](https://console.aws.amazon.com/servicequotas/home/services/apigateway/quotas/L-4D98A8A5)Multi-level API mappings per domain200NoRouting rules per domain50[Yes](https://console.aws.amazon.com/servicequotas/home/services/apigateway/quotas/L-68B79FF0)Length, in characters, of the URL for an edge-optimized API8192NoLength, in characters, of the URL for a Regional API10240NoLength, in characters, of the URL for a private API8192NoLength, in characters, of API Gateway resource policy8192[Yes](https://console.aws.amazon.com/servicequotas/home/services/apigateway/quotas/L-8B81B02C)API keys per account per Region10000NoClient certificates per account per Region60[Yes](https://console.aws.amazon.com/servicequotas/home/services/apigateway/quotas/L-824C9E42)Authorizers per API (AWS Lambda and Amazon Cognito)10

Yes

To increase this quota, contact the [AWS Support Center](https://console.aws.amazon.com/support/home)

Documentation parts per API2000

Yes

To increase this quota, contact the [AWS Support Center](https://console.aws.amazon.com/support/home)

Resources per API300[Yes](https://console.aws.amazon.com/servicequotas/home/services/apigateway/quotas/L-01C8A9E0)Stages per API10[Yes](https://console.aws.amazon.com/servicequotas/home/services/apigateway/quotas/L-379E48B0)Stage variables per stage100NoLength, in characters, of the key in a stage variable64NoLength, in characters, of the value in a stage variable512NoUsage plans per account per Region300[Yes](https://console.aws.amazon.com/servicequotas/home/services/apigateway/quotas/L-E8693075)Usage plans per API key10[Yes](https://console.aws.amazon.com/servicequotas/home/services/apigateway/quotas/L-985EB478)VPC links per account per Region20[Yes](https://console.aws.amazon.com/servicequotas/home/services/apigateway/quotas/L-A4C7274F)API caching TTL300 seconds by default and configurable between 0 and 3600 by an API
owner.Not for the upper bound (3600)Cached response size1048576 Bytes. Cache data encryption may increase the size of the item that is being cached.NoIntegration timeout for Regional APIs50 milliseconds - 29 seconds for all integration types, including Lambda,
Lambda proxy, HTTP, HTTP proxy, and AWS integrations.[Yes \*](https://console.aws.amazon.com/servicequotas/home/services/apigateway/quotas/L-E5AE38E3)Integration timeout for edge-optimized APIs50 milliseconds - 29 seconds for all integration types, including Lambda,
Lambda proxy, HTTP, HTTP proxy, and AWS integrations.NoIntegration timeout for private APIs50 milliseconds - 29 seconds for all integration types, including Lambda,
Lambda proxy, HTTP, HTTP proxy, and AWS integrations.[Yes \*](https://console.aws.amazon.com/servicequotas/home/services/apigateway/quotas/L-E5AE38E3)Total combined size of all header values, including header names, values, any line terminators, and whitespaces10240 BytesNoTotal combined size of all header values for a private API8000 BytesNoPayload size10 MBNoTags per stage50NoNumber of iterations in a `#foreach ... #end` loop in mapping
templates1000NoMaximum mapping template size300 KBNoAccess log template size3 KBNoMethod ARN length1600 bytes. If your method contains a path parameter and a client uses a value that exceeds the
ARN length, your API will return a `414 Request URI too long` response.NoMethod-level throttling settings for a stage in a usage plan20

Yes

To increase this quota, contact the [AWS Support Center](https://console.aws.amazon.com/support/home)

Model size per API400 KBNoNumber of certificates in a truststore1,000 certificates up to 1 MB total object size.NoIdle connection timeout310 secondsNoMaximum API definition file size when using [restapi:import](../api/api-importrestapi.md) or [restapi:put](../api/api-putrestapi.md)6 MBNo

\\* You can't set the integration timeout to less than 50 milliseconds. You can raise the integration timeout to
greater than 29 seconds, but this might require a reduction in your Region-level
throttle quota for your account.

## Best practices to reduce your quotas

The following best practices might help reduce your current number of resources to avoid increasing your quota.
Make sure that these suggestions work for your API's architecture.

**APIs per Region**

To reduce the number of APIs per Region, export any unused APIs and then delete them from API Gateway. For more information,
see [Export a REST API from API Gateway](api-gateway-export-api.md).

**Stages per API**

To reduce the number of stages per API, split up your API into multiple APIs.

**Resources per API**

Use `{proxy+}` paths to reduce the number of resources. For more information, see [Set up a proxy integration with a proxy resource](api-gateway-set-up-simple-proxy.md).

**API mappings**

To reduce the number of API mappings for a custom domain name, use API mappings with single levels, such
as `/prod`. API mappings with a single level doesn't count towards the API mappings quota.

**Authorizers per API**

To reduce the number of authorizers per API, reuse authorizers across API methods.

**Documentation parts per API**

To reduce the number of documentation parts per API, use `ignore=documentation` when you
import your API. For more information, see [Import API documentation](api-gateway-documenting-api-quick-start-import-export.md).

You can also use content inheritance to
allow some documentation parts to inherit content from an API entity of a more general specification. For more
information, see [Representation of API documentation in API Gateway](api-gateway-documenting-api-content-representation.md).

**Length, in characters, of API Gateway resource policy**

To reduce the length of a resource policy, use AWS WAFV2 to protect your API. For more information, see
[Use AWS WAF to protect your REST APIs in API Gateway](apigateway-control-access-aws-waf.md).

If your policy contains IP addresses, you can also
use ranges instead of specific values to define IP addresses.

**Usage plans per API key**

To reduce the number of usage plans per API key, use one API key per usage plan, and associate your
usage plan with multiple APIs. We don't recommend sharing one API key across multiple usage plans.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Quotas and important notes

HTTP API quotas

All content copied from https://docs.aws.amazon.com/.
