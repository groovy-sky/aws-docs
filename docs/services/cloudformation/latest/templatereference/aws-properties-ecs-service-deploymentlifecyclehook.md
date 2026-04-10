This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::Service DeploymentLifecycleHook

A deployment lifecycle hook runs custom logic at specific stages of the deployment process. Currently, you can use Lambda functions as hook targets.

For more information, see [Lifecycle hooks for Amazon ECS service deployments](../../../amazonecs/latest/developerguide/deployment-lifecycle-hooks.md) in the _Amazon Elastic Container Service Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "HookDetails" : Json,
  "HookTargetArn" : String,
  "LifecycleStages" : [ String, ... ],
  "RoleArn" : String
}

```

### YAML

```yaml

  HookDetails: Json
  HookTargetArn: String
  LifecycleStages:
    - String
  RoleArn: String

```

## Properties

`HookDetails`

Use this field to specify custom parameters that Amazon ECS passes to your hook target invocations (such as a Lambda function).

This field must be a JSON object as a string.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HookTargetArn`

The Amazon Resource Name (ARN) of the hook target. Currently, only Lambda function ARNs are supported.

You must provide this parameter when configuring a deployment lifecycle hook.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LifecycleStages`

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

_Required_: Yes

_Type_: Array of String

_Allowed values_: `RECONCILE_SERVICE | PRE_SCALE_UP | POST_SCALE_UP | TEST_TRAFFIC_SHIFT | POST_TEST_TRAFFIC_SHIFT | PRODUCTION_TRAFFIC_SHIFT | POST_PRODUCTION_TRAFFIC_SHIFT`

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The Amazon Resource Name (ARN) of the IAM role that grants Amazon ECS permission to call Lambda functions on your behalf.

For more information, see [Permissions required\
for Lambda functions in Amazon ECS blue/green deployments](../../../amazonecs/latest/developerguide/blue-green-permissions.md) in the _Amazon Elastic Container Service Developer Guide_.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeploymentController

EBSTagSpecification

All content copied from https://docs.aws.amazon.com/.
