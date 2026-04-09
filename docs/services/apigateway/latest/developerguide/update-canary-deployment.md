# Update a canary release

After a canary release is deployed, you may want to adjust the percentage of the
canary traffic or enable or disable the use of a stage cache to optimize the test
performance. You can also modify stage variables used in the canary release when the
execution context is updated. To make such updates, call the [**stage:update**](../api/api-updatestage.md) operation with new values on [canarySettings](../api/api-stage.md#canarySettings).

You can update a canary release using the API Gateway console, the AWS CLI [update-stage](../../../cli/latest/reference/apigateway/update-stage.md) command or an AWS
SDK.

###### Topics

- [Update a canary release using the API Gateway console](#update-canary-deployment-using-console)

- [Update a canary release using the AWS CLI](#update-canary-deployment-using-cli)

## Update a canary release using the API Gateway console

To use the API Gateway console to update existing canary settings on a stage, do the
following:

###### To update existing canary settings

1. Sign in to the API Gateway console and choose an existing REST API.

2. In the main navigation pane, choose **Stages**, and then choose an
    existing stage.

3. Choose the **Canary** tab, and then choose **Edit**. You might need to choose the right arrow button to show the **Canary** tab.

4. Update the **Request distribution** by
    increasing or decreasing the percentage number between 0.0 and 100.0,
    inclusive.

5. Select or clear the **Stage cache** the check box.

6. Add, remove, or modify **Canary stage variables**.

7. Choose **Save**.

## Update a canary release using the AWS CLI

To use the AWS CLI to update a canary, use the [`update-stage`](../../../cli/latest/reference/apigateway/update-stage.md) command and modify the patch operation for each parameter of the
canary.

The following [update-stage](../../../cli/latest/reference/apigateway/update-stage.md) command updates if
the canary uses the stage cache:

```nohighlight

aws apigateway update-stage \
    --rest-api-id {rest-api-id} \
    --stage-name '{stage-name}' \
    --patch-operations op=replace,path=/canarySettings/useStageCache,value=true
```

The following [update-stage](../../../cli/latest/reference/apigateway/update-stage.md) command updates
the canary traffic percentage:

```nohighlight

aws apigateway update-stage \
    --rest-api-id {rest-api-id} \
    --stage-name '{stage-name}' \
    --patch-operations op=replace,path=/canarySettings/percentTraffic,value=25.0
```

The following [update-stage](../../../cli/latest/reference/apigateway/update-stage.md) updates stage
variables. The example shows how to create a new stage variable named `newVar`, override the
`var2` stage variable, and remove the `var1` stage variable:

```nohighlight

aws apigateway update-stage  \
    --rest-api-id {rest-api-id} \
    --stage-name '{stage-name}'  \
    --patch-operations '[{
        "op": "replace",
        "path": "/canarySettings/stageVariableOverrides/newVar",
        "value": "newVal"
      }, {
        "op": "replace",
        "path": "/canarySettings/stageVariableOverrides/var2",
        "value": "val4"
      }, {
        "op": "remove",
        "path": "/canarySettings/stageVariableOverrides/var1"
      }]'
```

You can update all of the above by combining the operations into a single
`patch-operations` value:

```nohighlight

aws apigateway update-stage  \
    --rest-api-id {rest-api-id} \
    --stage-name '{stage-name}' \
    --patch-operations '[{
        "op": "replace",
        "path": "/canarySettings/percentTraffic",
        "value": "20.0"
    }, {
        "op": "replace",
        "path": "/canarySettings/useStageCache",
        "value": "true"
    }, {
        "op": "remove",
        "path": "/canarySettings/stageVariableOverrides/var1"
    }, {
        "op": "replace",
        "path": "/canarySettings/stageVariableOverrides/newVar",
        "value": "newVal"
    }, {
        "op": "replace",
        "path": "/canarySettings/stageVariableOverrides/val2",
        "value": "val4"
      }]'
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Create a canary release deployment

Promote a canary release

All content copied from https://docs.aws.amazon.com/.
