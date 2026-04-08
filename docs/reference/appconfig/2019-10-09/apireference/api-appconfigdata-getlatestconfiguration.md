# GetLatestConfiguration

Retrieves the latest deployed configuration. This API may return empty configuration
data if the client already has the latest version. For more information about this API
action and to view example AWS CLI commands that show how to use it with the [StartConfigurationSession](api-appconfigdata-startconfigurationsession.md) API action, see [Retrieving feature flags\
and configuration data in AWS AppConfig](../../../../services/appconfig/latest/userguide/retrieving-feature-flags.md) in the
_AWS AppConfig User Guide_.

###### Important

Note the following important information.

- Each configuration token is only valid for one call to
`GetLatestConfiguration`. The `GetLatestConfiguration`
response includes a `NextPollConfigurationToken` that should always
replace the token used for the just-completed call in preparation for the next
one.

- `GetLatestConfiguration` is a priced call. For more information, see
[Pricing](https://aws.amazon.com/systems-manager/pricing).

- You can configure applications to retrieve configuration data from a local
cache rather than directly calling AWS AppConfig. Caching can improve
performance and reduce costs. You can use the AWS AppConfig
AWS Lambda extension to cache data on your behalf. For more
information, see [AWS AppConfig integration with Lambda extensions](../../../../services/appconfig/latest/userguide/appconfig-integration-lambda-extensions.md)
in the AWS AppConfig User Guide.

For containerized environments, you can use the AWS AppConfig agent for
Amazon Elastic Container Service or Amazon Elastic Kubernetes Service. For more information, see
[AWS AppConfig integration with Amazon ECS and Amazon EKS](../../../../services/appconfig/latest/userguide/appconfig-integration-containers-agent.md).

## Request Syntax

```nohighlight

GET /configuration?configuration_token=ConfigurationToken HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[ConfigurationToken](#API_appconfigdata_GetLatestConfiguration_RequestSyntax)**

Token describing the current state of the configuration session. To obtain a token,
first call the [StartConfigurationSession](api-appconfigdata-startconfigurationsession.md) API. Note that every call to
`GetLatestConfiguration` will return a new `ConfigurationToken`
( `NextPollConfigurationToken` in the response) and _must_
be provided to subsequent `GetLatestConfiguration` API calls.

###### Important

This token should only be used once. To support long poll use cases, the token is
valid for up to 24 hours. If a `GetLatestConfiguration` call uses an expired
token, the system returns `BadRequestException`.

Pattern: `\S{1,8192}`

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Next-Poll-Configuration-Token: NextPollConfigurationToken
Next-Poll-Interval-In-Seconds: NextPollIntervalInSeconds
Content-Type: ContentType
Version-Label: VersionLabel

Configuration
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The response returns the following HTTP headers.

**[ContentType](#API_appconfigdata_GetLatestConfiguration_ResponseSyntax)**

A standard MIME type describing the format of the configuration content.

**[NextPollConfigurationToken](#API_appconfigdata_GetLatestConfiguration_ResponseSyntax)**

The latest token describing the current state of the configuration session. This
_must_ be provided to the next call to
`GetLatestConfiguration.`

###### Important

This token should only be used once. To support long poll use cases, the token is
valid for up to 24 hours. If a `GetLatestConfiguration` call uses an expired
token, the system returns `BadRequestException`.

Pattern: `\S{1,8192}`

**[NextPollIntervalInSeconds](#API_appconfigdata_GetLatestConfiguration_ResponseSyntax)**

The amount of time the client should wait before polling for configuration updates
again. The default value is 60 seconds. If you specify a value for
`RequiredMinimumPollIntervalInSeconds`, the service uses the specified value
for the desired poll interval instead.

**[VersionLabel](#API_appconfigdata_GetLatestConfiguration_ResponseSyntax)**

The user-defined label for the AWS AppConfig hosted configuration version. This
attribute doesn't apply if the configuration is not from an AWS AppConfig hosted
configuration version. If the client already has the latest version of the configuration
data, this value is empty.

The response returns the following as the HTTP body.

**[Configuration](#API_appconfigdata_GetLatestConfiguration_ResponseSyntax)**

The data of the configuration. This may be empty if the client already has the latest
version of configuration.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**BadRequestException**

The input fails to satisfy the constraints specified by the service.

**Details**

Details describing why the request was invalid.

**Reason**

Code indicating the reason the request was invalid.

HTTP Status Code: 400

**InternalServerException**

There was an internal failure in the service.

HTTP Status Code: 500

**ResourceNotFoundException**

The requested resource could not be found.

**ReferencedBy**

A map indicating which parameters in the request reference the resource that was not
found.

**ResourceType**

The type of resource that was not found.

HTTP Status Code: 404

**ThrottlingException**

The request was denied due to request throttling.

HTTP Status Code: 429

## Examples

### Example

This example illustrates one usage of GetLatestConfiguration.

#### Sample Request

```

GET /configuration?configuration_token=AYADeNgfsRxdKiJ37A12OZ9vN2cAXwABABVhd3MtY3J5cHRvLXB1YmxpYy1rZXkAREF1RzlLMTg1Tkx2Wjk4OGV2UXkyQ1UxV2ZxaGJTdkt2MkprKzJ4TlY2VkM0ZDVoRzBvbXFBTnFHcXpXcjU0dUM3Zz09AAEAB2F3cy1rbXMAS2Fybjphd3M6a21zOnVzLXdlc3QtMTo4MzIyNzQ3ODA3Mzc6a2V5LzA4MGYxYWI3LTNjNmYtNGE1Yi04ZTg2LWE2MjdmNWJiMmRlMQC4AQIBAHiWVBJIpqEFGgjvLHxtJImRmyljpZnxt7FfxeEOE5H8xQF1SfOlWZFnHujbzJmIvNeSAAAAfjB8BgkqhkiG9w0BBwagbzBtAgEAMGgGCSqGSIb3DQEHATAeBglghkgBZQMEAS4wEQQM0iTa25yw3uwnO6%2BrAgEQgDvLYYILWH%2Fx3SKKgyW1nOeBtS6go21vVPA6b8%2FwoFpxhzwjxy80jOfkVC2lM2l626BnJCvSPFfuZvcz%2FQIAAAAADAAAEAAAAAAAAAAAAAAAAADpQpWKSwnomgzEkYmidGKJ%2F%2F%2F%2F%2FwAAAAEAAAAAAAAAAAAAAAEAAAD8yfxkWB0geYfyI%2BDNJGiryebhmEoi8S8UHZSNN5JjJzTN2iORkjrA3DVvnhBTfoPh7o5bl4jwSYa%2F6as%2BmuQ9ntjYwymTZu7inYhsICYUKEDFxonBFJaEC32jEfg%2FMbPaGLOhNHdISiPAlMlOYmsw7phgl6ldbs9qrKVLlk1WNO3XTJiXyaWY4ANMfAX2JgMbGvNNY3HbfUneDGOENg6IfwKDykELIrJ6feE0JyKOV7mXfp%2F0r2pNiW9d6q%2BWDD1w4T87gCxgUGEPB%2FJ7JG3RhpGvECUmgKA0T06MjA7kWdt2IofDaLzRppFGpgLoPmxpM4qHz%2Fw6dMDmeXybKNZP84UP12zsJtUMhuspEQBnMGUCMQD8ssc6G8e6u8nov1ZdgF4m1ad3qyjiFd9DNRZHgLRFpw7%2BQIO%2FRB1l%2FIExP2ftUlkCMBT9oNlClJg4l9uGW5%2FqAiJ5n9ivK0ElRMwpvx96damGxt125XtMkmYf6a0OWSqnBw%3D%3D HTTP/1.1
Host: appconfigdata.us-west-1.amazonaws.com
Accept-Encoding: identity
User-Agent: aws-cli/2.4.15 Python/3.8.8 Windows/10 exe/AMD64 prompt/off command/appconfigdata.get-latest-configuration
X-Amz-Date: 20220218T190734Z
Authorization: AWS4-HMAC-SHA256 Credential=AKIAIOSFODNN7EXAMPLE/20220218/us-west-1/appconfig/aws4_request, SignedHeaders=host;x-amz-date, Signature=39c3b3042cd2aEXAMPLE

```

#### Sample Response

```

{
	"betaGroup": {
		"enabled": true
	}
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/appconfigdata-2021-11-11/getlatestconfiguration.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/appconfigdata-2021-11-11/getlatestconfiguration.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/appconfigdata-2021-11-11/getlatestconfiguration.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/appconfigdata-2021-11-11/getlatestconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appconfigdata-2021-11-11/getlatestconfiguration.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/appconfigdata-2021-11-11/getlatestconfiguration.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/appconfigdata-2021-11-11/getlatestconfiguration.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/appconfigdata-2021-11-11/getlatestconfiguration.md)

- [AWS SDK for Python](../../../../services/goto/boto3/appconfigdata-2021-11-11/getlatestconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appconfigdata-2021-11-11/getlatestconfiguration.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

AWS AppConfig Data

StartConfigurationSession
