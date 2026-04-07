# DeploymentLifecycleHook

A deployment lifecycle hook runs custom logic at specific stages of the deployment process. Currently, you can use Lambda functions as hook targets.

For more information, see [Lifecycle hooks for Amazon ECS service deployments](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-lifecycle-hooks.html) in the _Amazon Elastic Container Service Developer Guide_.

## Contents

**hookDetails**

Use this field to specify custom parameters that Amazon ECS will pass to your hook target invocations (such as a Lambda function).

Type: JSON value

Required: No

**hookTargetArn**

The Amazon Resource Name (ARN) of the hook target. Currently, only Lambda function ARNs are supported.

You must provide this parameter when configuring a deployment lifecycle hook.

Type: String

Required: No

**lifecycleStages**

The lifecycle stages at which to run the hook. Choose from these valid values:

- RECONCILE\_SERVICE

The reconciliation stage that only happens when you start a new service deployment with more than 1 service revision in an ACTIVE state.

You can use a lifecycle hook for this stage.

- PRE\_SCALE\_UP

The green service revision has not started. The blue service revision is handling 100% of the production traffic. There is no test traffic.

You can use a lifecycle hook for this stage.

- POST\_SCALE\_UP

The green service revision has started. The blue service revision is handling 100% of the production traffic. There is no test traffic.

You can use a lifecycle hook for this stage.

- TEST\_TRAFFIC\_SHIFT

The blue and green service revisions are running. The blue service revision handles 100% of the production traffic. The green service revision is migrating from 0% to 100% of test traffic.

You can use a lifecycle hook for this stage.

- POST\_TEST\_TRAFFIC\_SHIFT

The test traffic shift is complete. The green service revision handles 100% of the test traffic.

You can use a lifecycle hook for this stage.

- PRODUCTION\_TRAFFIC\_SHIFT

Production traffic is shifting to the green service revision. The green service revision is migrating from 0% to 100% of production traffic.

You can use a lifecycle hook for this stage.

- POST\_PRODUCTION\_TRAFFIC\_SHIFT

The production traffic shift is complete.

You can use a lifecycle hook for this stage.

You must provide this parameter when configuring a deployment lifecycle hook.

Type: Array of strings

Valid Values: `RECONCILE_SERVICE | PRE_SCALE_UP | POST_SCALE_UP | TEST_TRAFFIC_SHIFT | POST_TEST_TRAFFIC_SHIFT | PRODUCTION_TRAFFIC_SHIFT | POST_PRODUCTION_TRAFFIC_SHIFT`

Required: No

**roleArn**

The Amazon Resource Name (ARN) of the IAM role that grants Amazon ECS permission to call Lambda functions on your behalf.

For more information, see [Permissions required\
for Lambda functions in Amazon ECS blue/green deployments](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/blue-green-permissions.html) in the _Amazon Elastic Container Service Developer Guide_.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ecs-2014-11-13/DeploymentLifecycleHook)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ecs-2014-11-13/DeploymentLifecycleHook)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ecs-2014-11-13/DeploymentLifecycleHook)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DeploymentEphemeralStorage

Device
