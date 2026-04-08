# ModifyInstanceCreditSpecification

Modifies the credit option for CPU usage on a running or stopped burstable performance
instance. The credit options are `standard` and
`unlimited`.

For more information, see [Burstable\
performance instances](../../../../services/ec2/latest/userguide/burstable-performance-instances.md) in the _Amazon EC2 User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**ClientToken**

A unique, case-sensitive token that you provide to ensure idempotency of your
modification request. For more information, see [Ensuring\
Idempotency](run-instance-idempotency.md).

Type: String

Required: No

**DryRun**

Checks whether you have the required permissions for the operation, without actually making the
request, and provides an error response. If you have the required permissions, the error response is
`DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**InstanceCreditSpecification.N**

Information about the credit option for CPU usage.

Type: Array of [InstanceCreditSpecificationRequest](api-instancecreditspecificationrequest.md) objects

Required: Yes

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**successfulInstanceCreditSpecificationSet**

Information about the instances whose credit option for CPU usage was successfully
modified.

Type: Array of [SuccessfulInstanceCreditSpecificationItem](api-successfulinstancecreditspecificationitem.md) objects

**unsuccessfulInstanceCreditSpecificationSet**

Information about the instances whose credit option for CPU usage was not
modified.

Type: Array of [UnsuccessfulInstanceCreditSpecificationItem](api-unsuccessfulinstancecreditspecificationitem.md) objects

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This request modifies the credit option for CPU usage of the specified
instance in the specified Region to `unlimited`. Valid credit options
are `standard` and `unlimited`.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=ModifyInstanceCreditSpecification
&Region=us-east-1
&InstanceCreditSpecification.1.InstanceId=i-1234567890abcdef0
&InstanceCreditSpecification.1.CpuCredits=unlimited
&AUTHPARAMS
```

#### Sample Response

```

<ModifyInstanceCreditSpecificationResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>11111111-2222-3333-4444-5555EXAMPLE</requestId>
    <unsuccessfulInstanceCreditSpecificationSet/>
    <successfulInstanceCreditSpecificationSet>
        <item>
            <instanceId>i-1234567890abcdef0</instanceId>
        </item>
    </successfulInstanceCreditSpecificationSet>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/modifyinstancecreditspecification.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/modifyinstancecreditspecification.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/modifyinstancecreditspecification.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/modifyinstancecreditspecification.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/modifyinstancecreditspecification.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/modifyinstancecreditspecification.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/modifyinstancecreditspecification.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/modifyinstancecreditspecification.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/modifyinstancecreditspecification.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/modifyinstancecreditspecification.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ModifyInstanceCpuOptions

ModifyInstanceEventStartTime
