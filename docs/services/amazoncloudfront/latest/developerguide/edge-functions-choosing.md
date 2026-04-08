# Differences between CloudFront Functions and Lambda@Edge

CloudFront Functions and Lambda@Edge both provide a way to run code in response to CloudFront
events.

CloudFront Functions is ideal for lightweight, short-running functions for the following use
cases:

- **Cache key normalization** – Transform
HTTP request attributes (headers, query strings, cookies, and even the URL path)
to create an optimal [cache\
key](understanding-the-cache-key.md), which can improve your cache hit ratio.

- **Header manipulation** – Insert, modify,
or delete HTTP headers in the request or response. For example, you can add a
`True-Client-IP` header to every request.

- **URL redirects or rewrites** – Redirect
viewers to other pages based on information in the request, or rewrite all
requests from one path to another.

- **Request authorization** – Validate
hashed authorization tokens, such as JSON web tokens (JWT), by inspecting
authorization headers or other request metadata.

To get started with CloudFront Functions, see [Customize at the edge with CloudFront Functions](cloudfront-functions.md).

Lambda@Edge is ideal for the following use cases:

- Functions that take several milliseconds or more to complete

- Functions that require adjustable CPU or memory

- Functions that depend on third-party libraries (including the AWS SDK, for
integration with other AWS services)

- Functions that require network access to use external services for
processing

- Functions that require file system access or access to the body of HTTP
requests

To get started with Lambda@Edge, see [Customize at the edge with Lambda@Edge](lambda-at-the-edge.md).

To help you choose the option for your use case, use the following table to understand
the differences between CloudFront Functions and Lambda@Edge. For information about differences that apply for origin modification helper methods, see [Choose between CloudFront Functions and Lambda@Edge](helper-functions-origin-modification.md#origin-modification-considerations).

CloudFront FunctionsLambda@EdgeProgramming languagesJavaScript (ECMAScript 5.1 compliant)Node.js and PythonEvent sources

- Viewer request

- Viewer response

- Viewer request

- Viewer response

- Origin request

- Origin response

Supports [Amazon CloudFront KeyValueStore](kvs-with-functions.md)

Yes

CloudFront KeyValueStore only supports [JavaScript runtime\
2.0](functions-javascript-runtime-20.md)

No

ScaleUp to millions of requests per secondUp to 10,000 requests per second per RegionFunction durationSubmillisecond

Up to 30 seconds (viewer request and viewer response)

Up to 30 seconds (origin request and origin response)

Maximum function memory size

2 MB

128 MB (viewer request and viewer response)

10,240 MB (10 GB) (origin request and origin response)

For more information, see [Quotas on Lambda@Edge](cloudfront-limits.md#limits-lambda-at-edge).

Maximum size of the function code and included libraries10 KB

50 MB (viewer request and viewer response)

50 MB (origin request and origin response)

Network accessNoYesFile system accessNoYesAccess to the request bodyNoYesAccess to geolocation and device dataYes

No (viewer request and viewer response)

Yes (origin request and origin response)

Can build and test entirely within CloudFrontYesNoFunction logging and metricsYesYes

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Use functions to customize at the
edge

Customize with
CloudFront Functions

All content copied from https://docs.aws.amazon.com/.
