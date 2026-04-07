# ServiceConnectTestTrafficHeaderRules

The HTTP header rules used to identify and route test traffic during Amazon ECS blue/green
deployments. These rules specify which HTTP headers to examine and what values to match for
routing decisions.

For more information, see [Service Connect for Amazon ECS blue/green deployments](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-connect-blue-green.html) in the _Amazon Elastic Container Service Developer Guide_.

## Contents

**name**

The name of the HTTP header to examine for test traffic routing. Common examples include custom headers like `X-Test-Version` or `X-Canary-Request` that can be used to identify test traffic.

Type: String

Required: Yes

**value**

The header value matching configuration that determines how the HTTP header value is evaluated for test traffic routing decisions.

Type: [ServiceConnectTestTrafficHeaderMatchRules](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ServiceConnectTestTrafficHeaderMatchRules.html) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ecs-2014-11-13/ServiceConnectTestTrafficHeaderRules)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ecs-2014-11-13/ServiceConnectTestTrafficHeaderRules)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ecs-2014-11-13/ServiceConnectTestTrafficHeaderRules)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ServiceConnectTestTrafficHeaderMatchRules

ServiceConnectTestTrafficRules
