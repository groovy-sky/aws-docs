# CancelBundleTask

Cancels a bundling operation for an instance store-backed Windows instance.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**BundleId**

The ID of the bundle task.

Type: String

Required: Yes

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is
`DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

## Response Elements

The following elements are returned by the service.

**bundleInstanceTask**

Information about the bundle task.

Type: [BundleTask](api-bundletask.md) object

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example request cancels the specified bundle task.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=CancelBundleTask
&BundleId=bun-cla322b9
&AUTHPARAMS
```

#### Sample Response

```

<CancelBundleTaskResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
  <requestId>59dbff89-35bd-4eac-99ed-be587EXAMPLE</requestId>
  <bundleInstanceTask>
      <instanceId>i-1234567890abcdef0</instanceId>
      <bundleId>bun-cla322b9</bundleId>
      <state>canceling</state>
      <startTime>2008-10-07T11:41:50.000Z</startTime>
      <updateTime>2008-10-07T11:51:50.000Z</updateTime>
      <progress>20%</progress>
      <storage>
        <S3>
          <bucket>myawsbucket</bucket>
          <prefix>my-new-image</prefix>
        </S3>
      </storage>
  </bundleInstanceTask>
</CancelBundleTaskResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/cancelbundletask.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/cancelbundletask.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/cancelbundletask.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/cancelbundletask.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/cancelbundletask.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/cancelbundletask.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/cancelbundletask.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/cancelbundletask.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/cancelbundletask.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/cancelbundletask.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

BundleInstance

CancelCapacityReservation
