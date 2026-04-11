# DeleteAutoScalingConfiguration

###### Important

AWS App Runner will no longer be open to new
customers starting March 31, 2026. If you would like to use App Runner, sign up prior to that date. Existing customers can
continue to use the service as normal. For more information, see
[AWS App Runner availability change](../dg/apprunner-availability-change.md).

Delete an AWS App Runner automatic scaling configuration resource. You can delete a top level auto scaling configuration, a specific revision of one, or all
revisions associated with the top level configuration. You can't delete the default auto scaling configuration or a configuration that's used by one or
more App Runner services.

## Request Syntax

```nohighlight

{
   "AutoScalingConfigurationArn": "string",
   "DeleteAllRevisions": boolean
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[AutoScalingConfigurationArn](#API_DeleteAutoScalingConfiguration_RequestSyntax)**

The Amazon Resource Name (ARN) of the App Runner auto scaling configuration that you want to delete.

The ARN can be a full auto scaling configuration ARN, or a partial ARN ending with either `.../name
                  ` or
`.../name/revision
                  `. If a revision isn't specified, the latest active revision is deleted.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1011.

Pattern: `arn:aws(-[\w]+)*:[a-z0-9-\\.]{0,63}:[a-z0-9-\\.]{0,63}:[0-9]{12}:(\w|\/|-){1,1011}`

Required: Yes

**[DeleteAllRevisions](#API_DeleteAutoScalingConfiguration_RequestSyntax)**

Set to `true` to delete all of the revisions associated with the `AutoScalingConfigurationArn` parameter value.

When `DeleteAllRevisions` is set to `true`, the only valid value for the Amazon Resource Name (ARN) is a partial ARN ending
with: `.../name`.

Type: Boolean

Required: No

## Response Syntax

```nohighlight

{
   "AutoScalingConfiguration": {
      "AutoScalingConfigurationArn": "string",
      "AutoScalingConfigurationName": "string",
      "AutoScalingConfigurationRevision": number,
      "CreatedAt": number,
      "DeletedAt": number,
      "HasAssociatedService": boolean,
      "IsDefault": boolean,
      "Latest": boolean,
      "MaxConcurrency": number,
      "MaxSize": number,
      "MinSize": number,
      "Status": "string"
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[AutoScalingConfiguration](#API_DeleteAutoScalingConfiguration_ResponseSyntax)**

A description of the App Runner auto scaling configuration that this request just deleted.

Type: [AutoScalingConfiguration](api-autoscalingconfiguration.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InternalServiceErrorException**

An unexpected service exception occurred.

HTTP Status Code: 500

**InvalidRequestException**

One or more input parameters aren't valid. Refer to the API action's document page, correct the input parameters, and try the action again.

HTTP Status Code: 400

**ResourceNotFoundException**

A resource doesn't exist for the specified Amazon Resource Name (ARN) in your AWS account.

HTTP Status Code: 400

## Examples

### Delete the latest active revision of an auto scaling configuration

This example illustrates how to delete the latest active revision of an App Runner auto scaling configuration. To delete the latest active revision,
specify an Amazon Resource Name (ARN) that ends with the configuration name, without the revision component.

In the example, two revisions exist before this action. Therefore, revision `2` (the latest) is deleted. However, it now shows
`"Latest": false`, because, after being deleted, it isn't the latest active revision anymore.

#### Sample Request

```json

$ aws apprunner delete-auto-scaling-configuration --cli-input-json "`cat`"
{
  "AutoScalingConfigurationArn": "arn:aws:apprunner:us-east-1:123456789012:autoscalingconfiguration/high-availability"
}
```

#### Sample Response

```json

{
  "AutoScalingConfiguration": {
    "AutoScalingConfigurationArn": "arn:aws:apprunner:us-east-1:123456789012:autoscalingconfiguration/high-availability/2/e76562f50d78042e819fead0f59672e6",
    "AutoScalingConfigurationName": "high-availability",
    "AutoScalingConfigurationRevision": 2,
    "CreatedAt": "2021-02-25T17:42:59Z",
    "DeletedAt": "2021-03-02T08:07:06Z",
    "Latest": false,
    "Status": "INACTIVE",
    "MaxConcurrency": 30,
    "MaxSize": 25,
    "MinSize": 5
  }
}
```

### Delete a specific revision of an auto scaling configuration

This example illustrates how to delete a specific revision of an App Runner auto scaling configuration. To delete a specific revision, specify an ARN
that includes the revision number.

In the example, several revisions exist before this action. The action deletes revision `1`.

#### Sample Request

```json

$ aws apprunner delete-auto-scaling-configuration --cli-input-json "`cat`"
{
  "AutoScalingConfigurationArn": "arn:aws:apprunner:us-east-1:123456789012:autoscalingconfiguration/high-availability/1"
}
```

#### Sample Response

```json

{
  "AutoScalingConfiguration": {
    "AutoScalingConfigurationArn": "arn:aws:apprunner:us-east-1:123456789012:autoscalingconfiguration/high-availability/1/2f50e7656d7819fead0f59672e68042e",
    "AutoScalingConfigurationName": "high-availability",
    "AutoScalingConfigurationRevision": 1,
    "CreatedAt": "2020-11-03T00:29:17Z",
    "DeletedAt": "2021-03-02T08:07:06Z",
    "Latest": false,
    "Status": "INACTIVE",
    "MaxConcurrency": 100,
    "MaxSize": 25,
    "MinSize": 5
  }
}
```

### Delete all revisions of an auto scaling configuration

This example illustrates how to delete all of the revisions associated with an App Runner auto scaling configuration.

To delete all associated revisions set `DeleteAllRevisions` to `true`. Specify a partial ARN ending with
`.../name
            `.

The response does not include all of the fields listed in the _Response Syntax_. It only contains the fields shown in the
sample response.

#### Sample Request

```json

$ aws apprunner delete-auto-scaling-configuration --cli-input-json "`cat`"
{
    "AutoScalingConfigurationArn": "arn:aws:apprunner:us-east-1:123456789012:autoscalingconfiguration/high-availability",
    "DeleteAllRevisions": true
}
```

#### Sample Response

```json

{
    "AutoScalingConfiguration": {
        "AutoScalingConfigurationArn": "arn:aws:apprunner:us-east-1:123456789012:autoscalingconfiguration/high-availability",
        "AutoScalingConfigurationName": "high-availability",
        "Status": "inactive"
    }
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/apprunner-2020-05-15/deleteautoscalingconfiguration.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apprunner-2020-05-15/deleteautoscalingconfiguration.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apprunner-2020-05-15/deleteautoscalingconfiguration.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apprunner-2020-05-15/deleteautoscalingconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apprunner-2020-05-15/deleteautoscalingconfiguration.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apprunner-2020-05-15/deleteautoscalingconfiguration.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apprunner-2020-05-15/deleteautoscalingconfiguration.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apprunner-2020-05-15/deleteautoscalingconfiguration.md)

- [AWS SDK for Python](../../../goto/boto3/apprunner-2020-05-15/deleteautoscalingconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apprunner-2020-05-15/deleteautoscalingconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateVpcIngressConnection

DeleteConnection

All content copied from https://docs.aws.amazon.com/.
