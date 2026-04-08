# ServiceConnectTestTrafficHeaderMatchRules

The header matching rules for test traffic routing in Amazon ECS blue/green deployments.
These rules determine how incoming requests are matched based on HTTP headers to route test
traffic to the new service revision.

## Contents

**exact**

The exact value that the HTTP header must match for the test traffic routing rule to apply. This provides precise control over which requests are routed to the new service revision during blue/green deployments.

Type: String

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ecs-2014-11-13/serviceconnecttesttrafficheadermatchrules.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecs-2014-11-13/serviceconnecttesttrafficheadermatchrules.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecs-2014-11-13/serviceconnecttesttrafficheadermatchrules.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ServiceConnectServiceResource

ServiceConnectTestTrafficHeaderRules
