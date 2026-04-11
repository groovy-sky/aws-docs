# DeleteObservabilityConfiguration

###### Important

AWS App Runner will no longer be open to new
customers starting March 31, 2026. If you would like to use App Runner, sign up prior to that date. Existing customers can
continue to use the service as normal. For more information, see
[AWS App Runner availability change](../dg/apprunner-availability-change.md).

Delete an AWS App Runner observability configuration resource. You can delete a specific revision or the latest active revision. You can't delete a
configuration that's used by one or more App Runner services.

## Request Syntax

```nohighlight

{
   "ObservabilityConfigurationArn": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[ObservabilityConfigurationArn](#API_DeleteObservabilityConfiguration_RequestSyntax)**

The Amazon Resource Name (ARN) of the App Runner observability configuration that you want to delete.

The ARN can be a full observability configuration ARN, or a partial ARN ending with either `.../name
                  ` or
`.../name/revision
                  `. If a revision isn't specified, the latest active revision is deleted.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1011.

Pattern: `arn:aws(-[\w]+)*:[a-z0-9-\\.]{0,63}:[a-z0-9-\\.]{0,63}:[0-9]{12}:(\w|\/|-){1,1011}`

Required: Yes

## Response Syntax

```nohighlight

{
   "ObservabilityConfiguration": {
      "CreatedAt": number,
      "DeletedAt": number,
      "Latest": boolean,
      "ObservabilityConfigurationArn": "string",
      "ObservabilityConfigurationName": "string",
      "ObservabilityConfigurationRevision": number,
      "Status": "string",
      "TraceConfiguration": {
         "Vendor": "string"
      }
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[ObservabilityConfiguration](#API_DeleteObservabilityConfiguration_ResponseSyntax)**

A description of the App Runner observability configuration that this request just deleted.

Type: [ObservabilityConfiguration](api-observabilityconfiguration.md) object

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

### Delete the latest active revision of an observability configuration

This example illustrates how to delete the latest active revision of an App Runner observability configuration. To delete the latest active revision,
specify an Amazon Resource Name (ARN) that ends with the configuration name, without the revision component.

In the example, two revisions exist before this action. Therefore, revision `2` (the latest) is deleted.
However, it now shows `"Latest": false`, because, after being deleted, it isn't the latest active revision anymore.

###### Note

The two revisions in our examples are identical, because App Runner doesn't yet support enough observability functionality to demonstrate two
significantly different revisions (for example, multiple tracing vendors). We're including the two examples only to demonstrate the revisioning
behavior during deletion.

#### Sample Request

```json

$ aws apprunner delete-observability-configuration --cli-input-json "`cat`"
{
  "ObservabilityConfigurationArn": "arn:aws:apprunner:us-east-1:123456789012:observabilityconfiguration/xray-tracing"
}
```

#### Sample Response

```json

{
  "ObservabilityConfiguration": {
    "ObservabilityConfigurationArn": "arn:aws:apprunner:us-east-1:123456789012:observabilityconfiguration/xray-tracing/2/e76562f50d78042e819fead0f59672e6",
    "ObservabilityConfigurationName": "xray-tracing",
    "ObservabilityConfigurationRevision": 2,
    "CreatedAt": "2021-02-25T17:42:59Z",
    "DeletedAt": "2021-03-02T08:07:06Z",
    "Latest": false,
    "Status": "INACTIVE",
    "TraceConfiguration": {
        "Vendor": "AWSXRAY"
    }
  }
}
```

### Delete a specific revision of an observability configuration

This example illustrates how to delete a specific revision of an App Runner observability configuration. To delete a specific revision, specify an ARN
that includes the revision number.

In the example, several revisions exist before this action. The action deletes revision `1`.

#### Sample Request

```json

$ aws apprunner delete-observability-configuration --cli-input-json "`cat`"
{
  "ObservabilityConfigurationArn": "arn:aws:apprunner:us-east-1:123456789012:observabilityconfiguration/xray-tracing/1"
}
```

#### Sample Response

```json

{
  "ObservabilityConfiguration": {
    "ObservabilityConfigurationArn": "arn:aws:apprunner:us-east-1:123456789012:observabilityconfiguration/xray-tracing/1/2f50e7656d7819fead0f59672e68042e",
    "ObservabilityConfigurationName": "xray-tracing",
    "ObservabilityConfigurationRevision": 1,
    "CreatedAt": "2020-11-03T00:29:17Z",
    "DeletedAt": "2021-03-02T08:07:06Z",
    "Latest": false,
    "Status": "INACTIVE",
    "TraceConfiguration": {
        "Vendor": "AWSXRAY"
    }
  }
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/apprunner-2020-05-15/deleteobservabilityconfiguration.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apprunner-2020-05-15/deleteobservabilityconfiguration.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apprunner-2020-05-15/deleteobservabilityconfiguration.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apprunner-2020-05-15/deleteobservabilityconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apprunner-2020-05-15/deleteobservabilityconfiguration.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apprunner-2020-05-15/deleteobservabilityconfiguration.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apprunner-2020-05-15/deleteobservabilityconfiguration.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apprunner-2020-05-15/deleteobservabilityconfiguration.md)

- [AWS SDK for Python](../../../goto/boto3/apprunner-2020-05-15/deleteobservabilityconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apprunner-2020-05-15/deleteobservabilityconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteConnection

DeleteService

All content copied from https://docs.aws.amazon.com/.
