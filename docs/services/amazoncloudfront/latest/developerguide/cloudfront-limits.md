# Quotas

You can request a CloudFront quota increase by using the following options:

- You can use the Service Quotas console or the AWS Command Line Interface. For more information, see the
following topics:

- [Requesting a quota increase](https://docs.aws.amazon.com/servicequotas/latest/userguide/request-quota-increase.html) in the
_Service Quotas User Guide_

- [request-service-quota-increase](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/request-service-quota-increase.html) in the
_AWS CLI Command Reference_

- If a CloudFront quota isn't available in Service Quotas, use the AWS Support Center Console to create a
[service quota increase case](https://support.console.aws.amazon.com/support/home).

CloudFront is subject to the following quotas.

###### Topics

- [General quotas](#limits-general)

- [General quotas on distributions](#limits-web-distributions)

- [General quotas on policies](#limits-policies)

- [Quotas on mTLS and trust stores](#quotas-mtls)

- [Quotas on CloudFront Functions](#limits-functions)

- [Quotas on Connection Functions](#limits-connection-functions)

- [Quotas on key value stores](#limits-keyvaluestores)

- [Quotas on Lambda@Edge](#limits-lambda-at-edge)

- [Quotas on SSL certificates](#limits-ssl-certificates)

- [Quotas on invalidations](#limits-invalidations)

- [Quotas on key groups](#limits-key-groups)

- [Quotas on WebSocket connections](#limits-websockets)

- [Quotas on field-level encryption](#limits-field-level-encryption)

- [Quotas on cookies (legacy cache settings)](#limits-allowlisted-cookies)

- [Quotas on query strings (legacy cache settings)](#limits-allowlisted-query-strings)

- [Quotas on headers](#limits-custom-headers)

- [Quotas on multi-tenant distributions](#limits-template)

- [Related information](#related-information-cloudfront-quotas)

## General quotas

EntityDefault quota

Data transfer rate per distribution

(This quota doesn't apply for distributions subscribed to CloudFront
flat-rate pricing plans. For more information, see [CloudFront flat-rate pricing plans](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/flat-rate-pricing-plan.html).)

150 Gbps

[Request a higher quota](https://console.aws.amazon.com/support/home)

Requests per second per distribution

(This quota doesn't apply for distributions subscribed to CloudFront
flat-rate pricing plans. For more information, see [CloudFront flat-rate pricing plans](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/flat-rate-pricing-plan.html).)

250,000

[Request a higher quota](https://console.aws.amazon.com/support/home)

Tags that can be added to a distribution

50

[Request a higher quota](https://console.aws.amazon.com/support/home)

Files that you can serve per distribution

No quota

Maximum length of a request or an origin response, including
headers and query strings, but not including the body content

20,480 bytes

Maximum length of a URL

8,192 bytes

Maximum number of real-time access log delivery configurations per
AWS account

150

Maximum number of associations per web ACL

100

[Request a higher quota](https://console.aws.amazon.com/support/home)

## General quotas on distributions

EntityDefault quota

Alternate domain names (CNAMEs) per distribution

For more information, see [Use custom URLs by adding alternate domain names (CNAMEs)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/CNAMEs.html).

100

[Request a higher quota](https://console.aws.amazon.com/support/home)

Cache behaviors per distribution

75

[Request a higher quota](https://console.aws.amazon.com/support/home)

Connection attempts per origin

For more information, see [Connection attempts](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/DownloadDistValuesOrigin.html#origin-connection-attempts).

1-3

Connection timeout per origin

For more information, see [Connection timeout](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/DownloadDistValuesOrigin.html#origin-connection-timeout).

1-10 seconds

Response timeout per origin

This is also known as _origin read timeout_ or
_origin request timeout_. For
more information, see [Response timeout](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/DownloadDistValuesOrigin.html#DownloadDistValuesOriginResponseTimeout).

1-120 seconds

[Request a higher quota](https://console.aws.amazon.com/support/home)

Keep-alive timeout per origin

For more information, see [Keep-alive timeout (custom and VPC origins only)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/DownloadDistValuesOrigin.html#DownloadDistValuesOriginKeepaliveTimeout).

1-120 seconds

[Request a higher quota](https://console.aws.amazon.com/support/home)

Distributions per AWS account

For more information, see [Create a distribution](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-web-creating-console.html).

500

[Request a higher quota](https://console.aws.amazon.com/support/home)

Distributions per origin access control

100

[Request a higher quota](https://console.aws.amazon.com/support/home)

Distributions within chain of requests to origin endpoint

We don't recommend placing one distribution in front of another.
Exceeding this quota results in a 403 error.

2

File compression: range of file sizes that CloudFront compresses

For more information, see [Serve compressed files](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/ServingCompressedFiles.html).

1,000 to 10,000,000 bytes

Maximum cacheable file size per HTTP GET response.

Only the responses for an HTTP GET are cached. Responses for POST
or PUT are not cached.

50 GB

Maximum size of an HTTP request body.

64 GB

Origin access controls per AWS account

100

[Request a higher quota](https://console.aws.amazon.com/support/home)

Origin access identities per AWS account

100

[Request a higher quota](https://console.aws.amazon.com/support/home)

Origins per distribution

100

[Request a higher quota](https://console.aws.amazon.com/support/home)

Origin groups per distribution

10

[Request a higher quota](https://console.aws.amazon.com/support/home)

Staging distributions per AWS account

For more information, see [Use CloudFront continuous deployment to safely test CDN configuration changes](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/continuous-deployment.html).

20

[Request a higher quota](https://console.aws.amazon.com/support/home)

Distributions associated with the same VPC origin

50

VPC origins per AWS account

25

[Request a higher quota](https://console.aws.amazon.com/support/home)

Maximum number of distributions that can be associated with a
single Anycast static IP list.

100

[Request a higher quota](https://console.aws.amazon.com/support/home)

## General quotas on policies

EntityDefault quota

Custom cache policies per AWS account

(Does not apply to [CloudFront managed cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/using-managed-cache-policies.html))

20

[Request a higher quota](https://console.aws.amazon.com/servicequotas/home/services/cloudfront/quotas/L-7D134442)

Distributions associated with the same cache policy

100

Query strings per cache policy

10

[Request a higher quota](https://console.aws.amazon.com/support/home)

Headers per cache policy

10

[Request a higher quota](https://console.aws.amazon.com/support/home)

Cookies per cache policy

10

[Request a higher quota](https://console.aws.amazon.com/support/home)

Total combined length of all query string, header, and cookie
names in a cache policy

1024

Custom origin request policies per AWS account

(Does not apply to [CloudFront managed\
origin request policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/using-managed-origin-request-policies.html))

20

[Request a higher quota](https://console.aws.amazon.com/servicequotas/home/services/cloudfront/quotas/L-C3659C43)

Distributions associated with the same origin request
policy

100

Query strings per origin request policy

10

[Request a higher quota](https://console.aws.amazon.com/support/home)

Headers per origin request policy

10

[Request a higher quota](https://console.aws.amazon.com/support/home)

Cookies per origin request policy

10

[Request a higher quota](https://console.aws.amazon.com/support/home)

Total combined length of all query string, header, and cookie
names in an origin request policy

1024

Custom response headers policies per AWS account

(Does not apply to [CloudFront managed\
response headers policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/using-managed-response-headers-policies.html))

20

[Request a higher quota](https://console.aws.amazon.com/support/home)

Distributions associated with the same response headers
policy

100

[Request a higher quota](https://console.aws.amazon.com/support/home)

Custom headers per response headers policy

10

[Request a higher quota](https://console.aws.amazon.com/support/home)

Continuous deployment policies per AWS account

20

[Request a higher quota](https://console.aws.amazon.com/support/home)

## Quotas on mTLS and trust stores

Entity

Default quota

Trust stores per AWS account

20

[Request a higher quota](https://console.aws.amazon.com/support/home)

Distributions per trust store

25

CA bundle size

64 KB

[Request a higher quota](https://console.aws.amazon.com/support/home)

Certificate size in CA bundle

16384

[Request a higher quota](https://console.aws.amazon.com/support/home)

Number of certificates in CA bundle

25

Certificate chain depth

4

## Quotas on CloudFront Functions

Entity

Default quota

Functions per AWS account

100

Maximum function size

This quota isn't adjustable. To store additional data for your
CloudFront Functions, create a key value store and add your key-value
pairs. For more information, see [Amazon CloudFront KeyValueStore](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/kvs-with-functions.html).

10 KB

Maximum function memory

2 MB

Distributions associated with the same function

100

In addition to these quotas, there are some other restrictions when using
CloudFront Functions. For more information, see [Restrictions on CloudFront Functions](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/cloudfront-function-restrictions.html).

## Quotas on Connection Functions

Entity

Default quota

Connection Functions per AWS account

100

Maximum Connection Function size

This quota isn't adjustable. To store additional data for your
Connection Functions, create a key value store and add your key-value
pairs. For more information, see [Amazon CloudFront KeyValueStore](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/kvs-with-functions.html).

10 KB

Maximum Connection Function memory

2 MB

Distributions associated with the same Connection Function

100

In addition to these quotas, there are some other restrictions when using
Connection Functions. For more information, see [Associate a CloudFront Connection Function](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/connection-functions.html).

## Quotas on key value stores

Entity

Default quota

Maximum size of a key in a key-value pair512 BytesMaximum size of the value in a key-value pair1 KBMaximum key values pairs that you can update in a single API
request50 keys or 3 MB payload, whichever is reached firstMaximum size of an individual key value store5 MBMaximum number of functions that a single key value store can be
associated with10Maximum number of key value stores per function1Maximum number of key value stores per account

50

[Request a higher quota](https://console.aws.amazon.com/support/home)

## Quotas on Lambda@Edge

General quotas

Entity

Default quota

Distributions per AWS account that can have Lambda@Edge
functions

500

[Request a higher quota](https://console.aws.amazon.com/support/home)

Lambda@Edge functions per distribution

100

[Request a higher quota](https://console.aws.amazon.com/support/home)

Concurrent executions

###### Notes

- AWS Lambda manages the concurrency quotas for
Lambda@Edge. All Lambda functions in the AWS Region
share this quota.

- We recommend that you review the concurrent executions
quota in all AWS Regions where you expect your viewer
requests to come from. In addition, each instance of
your Lambda@Edge function can serve up to 10 requests per
second. The total invocation limit is 10 times your
concurrency limit.

For more information, see the following topics in the
_AWS Lambda Developer Guide_:

- [Understanding Lambda function\
scaling](https://docs.aws.amazon.com/lambda/latest/dg/scaling.html)

- [Lambda API requests](https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-limits.html#api-requests)

1,000 (in each AWS Region)

[Request a higher quota](https://console.aws.amazon.com/support/home)

Distributions associated with the same function

500

Maximum compressed size of a Lambda function and any included
libraries

50 MB

Lambda@Edge requests per second (each supported
AWS Region).

For more information, see [Concurrency quotas](https://docs.aws.amazon.com/lambda/latest/dg/lambda-concurrency.html#concurrency-quotas) in the
_AWS Lambda Developer Guide_.

10,000

Function memory size

Same as [Lambda quotas](https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-limits.html)

Function timeout. The function can make network calls to resources
such as Amazon S3 buckets, DynamoDB tables, or Amazon EC2 instances in
AWS Regions.

30 seconds

Quotas that differ by event type

Entity

Viewer request and viewer response events

Origin request and origin response events

Size of a response that is generated by a Lambda function,
including headers and body

40 KB

1 MB

###### Notes

- For a list of additional Lambda@Edge quotas that can be increased from
Service Quotas, see [Amazon CloudFront\
endpoints and quotas](https://docs.aws.amazon.com/general/latest/gr/cf_region.html#limits_cloudfront) in the
_AWS General Reference_.

- In addition to these quotas, there are some other restrictions when using
Lambda@Edge functions. For more information, see [Restrictions on Lambda@Edge](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/lambda-at-edge-function-restrictions.html).

## Quotas on SSL certificates

EntityDefault quota

SSL certificates per AWS account when serving HTTPS requests
using dedicated IP addresses (no quota when serving HTTPS requests
using SNI)

For more information, see [Use HTTPS with CloudFront](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/using-https.html).

2

[Request a higher quota](https://console.aws.amazon.com/support/home)

SSL certificates that can be associated with a CloudFront
distribution

1

If your SSL certificate is specifically for HTTPS communication between viewers and
CloudFront, and if you used AWS Certificate Manager (ACM) or the IAM certificate store to provision or
import your certificate, additional quotas apply. For more information, see [Quotas on using SSL/TLS certificates with CloudFront (HTTPS between viewers and CloudFront only)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/cnames-and-https-limits.html).

There are also quotas on the number of SSL certificates that you can import into
AWS Certificate Manager (ACM) or upload to AWS Identity and Access Management (IAM). For more information, see [Increase the quotas for SSL/TLS certificates](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/increasing-the-limit-for-ssl-tls-certificates.html).

## Quotas on invalidations

EntityDefault quota

File invalidation: maximum number of files allowed in active
invalidation requests, excluding wildcard invalidations

For more information, see [Invalidate files to remove content](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Invalidation.html).

3,000

File invalidation: maximum number of active wildcard invalidations
allowed

15

File invalidation: maximum number of files that one wildcard
invalidation can process

No quota

## Quotas on key groups

EntityDefault quota

Public keys in a single key group

5

[Request a higher quota](https://console.aws.amazon.com/support/home)

Key groups associated with a single cache behavior

4

[Request a higher quota](https://console.aws.amazon.com/support/home)

Key groups per AWS account

10

[Request a higher quota](https://console.aws.amazon.com/support/home)

Distributions associated with a single key group

100

[Request a higher quota](https://console.aws.amazon.com/support/home)

## Quotas on WebSocket connections

EntityDefault quota

Origin response timeout (idle timeout)

10 minutes

If CloudFront hasn't detected any bytes sent from the origin to the
client within the past 10 minutes, the connection is assumed to be
idle and is closed.

## Quotas on field-level encryption

EntityDefault quota

Maximum length of a field to encrypt

For more information, see [Use field-level encryption to help protect sensitive data](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/field-level-encryption.html).

16 KB

Maximum number of fields in a request body when field-level
encryption is configured

10

Maximum length of a request body when field-level encryption is
configured

1 MB

Maximum number of field-level encryption configurations that can
be associated with one AWS account

10

Maximum number of field-level encryption profiles that can be
associated with one AWS account

10

Maximum number of public keys that can be added to one
AWS account

10

Maximum number of fields to encrypt that can be specified in one
profile

10

Maximum number of CloudFront distributions that can be associated with a
field-level encryption configuration

20

Maximum number of query argument profile mappings that can be
included in a field-level encryption configuration

5

## Quotas on cookies (legacy cache settings)

These quotas apply to CloudFront's legacy cache settings. We recommend using a [cache policy](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html) or [origin request policy](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-origin-requests.html) instead of the
legacy settings.

EntityDefault quota

Cookies per cache behavior

For more information, see [Cache content based on cookies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Cookies.html).

10

[Request a higher quota](https://console.aws.amazon.com/support/home)

Total number of bytes in cookie names (doesn't apply if you
configure CloudFront to forward all cookies to the origin)

512 minus the number of cookies

## Quotas on query strings (legacy cache settings)

These quotas apply to CloudFront's legacy cache settings. We recommend using a [cache policy](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html) or [origin request policy](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-origin-requests.html) instead of the
legacy settings.

EntityDefault quota

Maximum number of characters in a query string

128 characters

Maximum number of characters total for all query strings in the
same parameter

512 characters

Query strings per cache behavior

For more information, see [Cache content based on query string parameters](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/QueryStringParameters.html).

10

[Request a higher quota](https://console.aws.amazon.com/support/home)

## Quotas on headers

EntityDefault quota

Headers per cache behavior (legacy cache settings)

For more information, see [Cache content based on request headers](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/header-caching.html).

10

[Request a higher quota](https://console.aws.amazon.com/support/home)

Forward headers per cache behavior

25

[Request a higher quota](https://console.aws.amazon.com/support/home)

Custom headers: maximum number of custom headers that you can
configure CloudFront to add to origin requests

For more information, see [Add custom headers to origin requests](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/add-origin-custom-headers.html).

30

[Request a higher quota](https://console.aws.amazon.com/support/home)

Custom headers: maximum number of custom headers that you can add
to a response headers policy

10

[Request a higher quota](https://console.aws.amazon.com/support/home)

Custom headers: maximum length of a header name

256 characters

Custom headers: maximum length of a header value

1,783 characters

Custom headers: maximum length of all header values and names
combined

10,240 characters

Maximum length of the `Content-Security-Policy` header
value

1,783 characters

[Request a higher quota](https://console.aws.amazon.com/servicequotas/home/services/cloudfront/quotas/L-E9944CCE)

Maximum length of a `CORS (Access-Control-Allow-Origin)`
header value

1,783 characters

## Quotas on multi-tenant distributions

Entity

Default quota

Maximum number of distribution tenants per AWS account

10,000

[Request a higher quota](https://console.aws.amazon.com/support/home)

Maximum number of multi-tenant distributions per AWS account

20

[Request a higher quota](https://console.aws.amazon.com/support/home)

Maximum number of connection groups per AWS account

100

[Request a higher quota](https://console.aws.amazon.com/support/home)

Maximum number of aliases per distribution tenant

100

[Request a higher quota](https://console.aws.amazon.com/support/home)

Maximum number of parameters per distribution tenant

5

[Request a higher quota](https://console.aws.amazon.com/support/home)

Maximum number of parameters per multi-tenant distribution

5

[Request a higher quota](https://console.aws.amazon.com/support/home)

Maximum number of parameters in a field in a multi-tenant distribution

2

[Request a higher quota](https://console.aws.amazon.com/support/home)

Maximum number of connection groups per Anycast static IP
list

5

[Request a higher quota](https://console.aws.amazon.com/support/home)

For more information about multi-tenant distributions, see [Understand how multi-tenant distributions work](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-config-options.html).

## Related information

For more information, see [Amazon CloudFront endpoints and quotas](https://docs.aws.amazon.com/general/latest/gr/cf_region.html) in the
_AWS General Reference_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Load testing CloudFront

Code examples
