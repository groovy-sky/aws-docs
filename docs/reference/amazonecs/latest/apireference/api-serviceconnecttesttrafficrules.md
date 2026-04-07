# ServiceConnectTestTrafficRules

The test traffic routing configuration for Amazon ECS blue/green deployments. This
configuration allows you to define rules for routing specific traffic to the new service
revision during the deployment process, allowing for safe testing before full production
traffic shift.

For more information, see [Service Connect for Amazon ECS blue/green deployments](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-connect-blue-green.html) in the _Amazon Elastic Container Service Developer Guide_.

## Contents

**header**

The HTTP header-based routing rules that determine which requests should be routed to the new service version during blue/green deployment testing. These rules provide fine-grained control over test traffic routing based on request headers.

Type: [ServiceConnectTestTrafficHeaderRules](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ServiceConnectTestTrafficHeaderRules.html) object

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ecs-2014-11-13/ServiceConnectTestTrafficRules)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ecs-2014-11-13/ServiceConnectTestTrafficRules)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ecs-2014-11-13/ServiceConnectTestTrafficRules)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ServiceConnectTestTrafficHeaderRules

ServiceConnectTlsCertificateAuthority
