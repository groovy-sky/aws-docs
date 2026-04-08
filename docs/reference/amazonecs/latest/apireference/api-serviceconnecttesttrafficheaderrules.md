# ServiceConnectTestTrafficHeaderRules

The HTTP header rules used to identify and route test traffic during Amazon ECS blue/green
deployments. These rules specify which HTTP headers to examine and what values to match for
routing decisions.

For more information, see [Service Connect for Amazon ECS blue/green deployments](../../../../services/amazonecs/latest/developerguide/service-connect-blue-green.md) in the _Amazon Elastic Container Service Developer Guide_.

## Contents

**name**

The name of the HTTP header to examine for test traffic routing. Common examples include custom headers like `X-Test-Version` or `X-Canary-Request` that can be used to identify test traffic.

Type: String

Required: Yes

**value**

The header value matching configuration that determines how the HTTP header value is evaluated for test traffic routing decisions.

Type: [ServiceConnectTestTrafficHeaderMatchRules](api-serviceconnecttesttrafficheadermatchrules.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ecs-2014-11-13/serviceconnecttesttrafficheaderrules.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecs-2014-11-13/serviceconnecttesttrafficheaderrules.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecs-2014-11-13/serviceconnecttesttrafficheaderrules.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ServiceConnectTestTrafficHeaderMatchRules

ServiceConnectTestTrafficRules
