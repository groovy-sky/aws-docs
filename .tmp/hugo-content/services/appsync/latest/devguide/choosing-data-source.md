# Choosing between direct data source access and proxying via a Lambda data source

With AWS AppSync and the `APPSYNC_JS` runtime, you can write your own code that implements your
custom business logic by using AWS AppSync functions to access your data sources. This makes it easy for you
to directly interact with data sources like Amazon DynamoDB, Aurora Serverless, OpenSearch Service, HTTP APIs, and
other AWS services without having to deploy additional computational services or infrastructure.
AWS AppSync also makes it easy to interact with an AWS Lambda function by configuring a Lambda data source.
Lambda data sources allow you to run complex business logic using AWS Lambda’s full set capabilities to resolve a
GraphQL request. In most cases, an AWS AppSync function directly connected to its target data source will
provide all of the functionality you need. In situations where you need to implement complex business logic
that is not supported by the `APPSYNC_JS` runtime, you can use a Lambda data source as a proxy to
interact with your target data source.

**Direct data source integration****Lambda data source as a proxy****Use case**AWS AppSync functions interact directly with API data sources.AWS AppSync functions call Lambdas that interact with API data
sources.Runtime_`APPSYNC_JS` (JavaScript)__Any supported Lambda runtime_Maximum size of code_32,000 characters per AWS AppSync_
_function__50 MB (zipped, for direct upload) per_
_Lambda_External modules_Limited - APPSYNC\_JS supported features_
_only__Yes_Call any AWS service_Yes - Using AWS AppSync HTTP_
_datasource__Yes - Using AWS SDK_Access to the request header_Yes__Yes_Network access_No__Yes_File system access_No__Yes_Logging and metrics_Yes__Yes_Build and test entirely within AppSync_Yes__No_Cold start_No__No - With provisioned concurrency_Auto-scaling_Yes - transparently by AWS AppSync__Yes - As configured in Lambda_Pricing_No additional charge__Charged for Lambda usage_

AWS AppSync functions that integrate directly with their target data source are ideal for use cases like
the following:

- Interacting with Amazon DynamoDB, Aurora Serverless, and OpenSearch Service

- Interacting with HTTP APIs and passing incoming headers

- Interacting with AWS services using HTTP data sources (with AWS AppSync automatically signing requests
with the provided data source role)

- Implementing access control before accessing data sources

- Implementing the filtering of retrieved data prior to fulfilling a request

- Implementing simple orchestration with sequential execution of AWS AppSync functions in a resolver
pipeline

- Controlling caching and subscription connections in queries and mutations.

AWS AppSync functions that use a Lambda data source as a proxy are ideal for use cases like the
following:

- Using a language other than JavaScript or Velocity Template Language (VTL)

- Adjusting and controlling CPU or memory to optimize performance

- Importing third-party libraries or requiring unsupported features in `APPSYNC_JS`

- Making multiple network requests and/or getting file system access to fulfill a query

- Batching requests using [batching configuration](resolver-reference-lambda-js.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Migrating from VTL to JavaScript

JavaScript resolver context
object reference

All content copied from https://docs.aws.amazon.com/.
