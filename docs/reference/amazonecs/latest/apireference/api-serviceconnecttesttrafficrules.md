# ServiceConnectTestTrafficRules

The test traffic routing configuration for Amazon ECS blue/green deployments. This
configuration allows you to define rules for routing specific traffic to the new service
revision during the deployment process, allowing for safe testing before full production
traffic shift.

For more information, see [Service Connect for Amazon ECS blue/green deployments](../../../../services/amazonecs/latest/developerguide/service-connect-blue-green.md) in the _Amazon Elastic Container Service Developer Guide_.

## Contents

**header**

The HTTP header-based routing rules that determine which requests should be routed to the new service version during blue/green deployment testing. These rules provide fine-grained control over test traffic routing based on request headers.

Type: [ServiceConnectTestTrafficHeaderRules](api-serviceconnecttesttrafficheaderrules.md) object

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ecs-2014-11-13/serviceconnecttesttrafficrules.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecs-2014-11-13/serviceconnecttesttrafficrules.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecs-2014-11-13/serviceconnecttesttrafficrules.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ServiceConnectTestTrafficHeaderRules

ServiceConnectTlsCertificateAuthority
