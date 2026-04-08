# AdvancedConfiguration

The advanced settings for a load balancer used in blue/green deployments. Specify the
alternate target group, listener rules, and IAM role required for traffic shifting during
blue/green deployments. For more information, see [Required resources for Amazon ECS\
blue/green deployments](../../../../services/amazonecs/latest/developerguide/blue-green-deployment-implementation.md) in the _Amazon Elastic Container Service Developer Guide_.

## Contents

**alternateTargetGroupArn**

The Amazon Resource Name (ARN) of the alternate target group for Amazon ECS blue/green deployments.

Type: String

Required: No

**productionListenerRule**

The Amazon Resource Name (ARN) that that identifies the production listener rule (in the case of an Application Load Balancer) or
listener (in the case for an Network Load Balancer) for routing production traffic.

Type: String

Required: No

**roleArn**

The Amazon Resource Name (ARN) of the IAM role that grants Amazon ECS permission to call the Elastic Load Balancing APIs for you.

Type: String

Required: No

**testListenerRule**

The Amazon Resource Name (ARN) that identifies ) that identifies the test listener rule (in the case of an Application Load Balancer) or
listener (in the case for an Network Load Balancer) for routing test traffic.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ecs-2014-11-13/advancedconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecs-2014-11-13/advancedconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecs-2014-11-13/advancedconfiguration.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

AcceleratorTotalMemoryMiBRequest

Attachment
