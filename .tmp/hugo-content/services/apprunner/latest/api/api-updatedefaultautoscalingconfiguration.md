# UpdateDefaultAutoScalingConfiguration

###### Important

AWS App Runner will no longer be open to new
customers starting March 31, 2026. If you would like to use App Runner, sign up prior to that date. Existing customers can
continue to use the service as normal. For more information, see
[AWS App Runner availability change](../dg/apprunner-availability-change.md).

Update an auto scaling configuration to be the default. The existing default auto scaling configuration will be set to non-default
automatically.

## Request Syntax

```nohighlight

{
   "AutoScalingConfigurationArn": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[AutoScalingConfigurationArn](#API_UpdateDefaultAutoScalingConfiguration_RequestSyntax)**

The Amazon Resource Name (ARN) of the App Runner auto scaling configuration that you want to set as the default.

The ARN can be a full auto scaling configuration ARN, or a partial ARN ending with either `.../name
                  ` or
`.../name/revision
                  `. If a revision isn't specified, the latest active revision is set as the
default.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1011.

Pattern: `arn:aws(-[\w]+)*:[a-z0-9-\\.]{0,63}:[a-z0-9-\\.]{0,63}:[0-9]{12}:(\w|\/|-){1,1011}`

Required: Yes

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

**[AutoScalingConfiguration](#API_UpdateDefaultAutoScalingConfiguration_ResponseSyntax)**

A description of the App Runner auto scaling configuration that was set as default.

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

### Update the latest active revision of an auto scaling configuration to be the default

This example illustrates how to update the latest active revision of an App Runner auto scaling configuration to be the default. To designate the latest
active revision as the default, specify an ARN that ends with the configuration name, without the revision component.

In the example, two revisions exist. Therefore, revision 2, (the latest revision), is set as the default. The resulting object shows
`"IsDefault": true` and `"Latest": true`.

#### Sample Request

```json

$ aws apprunner update-default-auto-scaling-configuration --cli-input-json "`cat`"
{
    "AutoScalingConfigurationArn": "arn:aws:apprunner:us-east-1:123456789012:autoscalingconfiguration/high-availability"
}
```

#### Sample Response

```json

{
    "AutoScalingConfiguration": {
        "AutoScalingConfigurationArn": "arn:aws:apprunner:us-east-1:123456789012:autoscalingconfiguration/high-availability/2/6a4d47db94434d30a42cab9a00d21d44",
        "AutoScalingConfigurationName": "high-availability",
        "AutoScalingConfigurationRevision": 2,
        "Latest": true,
        "Status": "active",
        "MaxConcurrency": 100,
        "MinSize": 1,
        "MaxSize": 25,
        "CreatedAt": "2023-09-01T00:00:00Z",
        "HasAssociatedService": false,
        "IsDefault": true
    }
}
```

### Update a specific revision of an auto scaling configuration to be the default

This example illustrates how to update a specific revision of an App Runner auto scaling configuration to be the default. To designate a specific
revision as the default, specify an ARN that includes the revision number.

In the example, several revisions exist, and revision 1 is set as the default. The resulting object shows `"IsDefault": true` and
`"Latest": false`.

#### Sample Request

```json

$ aws apprunner update-default-auto-scaling-configuration --cli-input-json "`cat`"
{
    "AutoScalingConfigurationArn": "arn:aws:apprunner:us-east-1:123456789012:autoscalingconfiguration/high-availability/1"
}
```

#### Sample Response

```json

{
    "AutoScalingConfiguration": {
        "AutoScalingConfigurationArn": "arn:aws:apprunner:us-east-1:123456789012:autoscalingconfiguration/high-availability/1/d2321df129c8440da8c464af7ebcd887",
        "AutoScalingConfigurationName": "high-availability",
        "AutoScalingConfigurationRevision": 1,
        "Latest": false,
        "Status": "active",
        "MaxConcurrency": 100,
        "MinSize": 1,
        "MaxSize": 25,
        "CreatedAt": "2023-09-01T00:00:00Z",
        "HasAssociatedService": false,
        "IsDefault": true
    }
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/apprunner-2020-05-15/updatedefaultautoscalingconfiguration.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apprunner-2020-05-15/updatedefaultautoscalingconfiguration.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apprunner-2020-05-15/updatedefaultautoscalingconfiguration.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apprunner-2020-05-15/updatedefaultautoscalingconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apprunner-2020-05-15/updatedefaultautoscalingconfiguration.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apprunner-2020-05-15/updatedefaultautoscalingconfiguration.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apprunner-2020-05-15/updatedefaultautoscalingconfiguration.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apprunner-2020-05-15/updatedefaultautoscalingconfiguration.md)

- [AWS SDK for Python](../../../goto/boto3/apprunner-2020-05-15/updatedefaultautoscalingconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apprunner-2020-05-15/updatedefaultautoscalingconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UntagResource

UpdateService

All content copied from https://docs.aws.amazon.com/.
