# ModifyDefaultCreditSpecification

Modifies the default credit option for CPU usage of burstable performance instances.
The default credit option is set at the account level per AWS Region, and
is specified per instance family. All new burstable performance instances in the account
launch using the default credit option.

`ModifyDefaultCreditSpecification` is an asynchronous operation, which
works at an AWS Region level and modifies the credit option for each
Availability Zone. All zones in a Region are updated within five minutes. But if
instances are launched during this operation, they might not get the new credit option
until the zone is updated. To verify whether the update has occurred, you can call
`GetDefaultCreditSpecification` and check
`DefaultCreditSpecification` for updates.

For more information, see [Burstable\
performance instances](../../../../services/ec2/latest/userguide/burstable-performance-instances.md) in the _Amazon EC2 User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**CpuCredits**

The credit option for CPU usage of the instance family.

Valid Values: `standard` \| `unlimited`

Type: String

Required: Yes

**DryRun**

Checks whether you have the required permissions for the operation, without actually making the
request, and provides an error response. If you have the required permissions, the error response is
`DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**InstanceFamily**

The instance family.

Type: String

Valid Values: `t2 | t3 | t3a | t4g`

Required: Yes

## Response Elements

The following elements are returned by the service.

**instanceFamilyCreditSpecification**

The default credit option for CPU usage of the instance family.

Type: [InstanceFamilyCreditSpecification](api-instancefamilycreditspecification.md) object

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example 1

This example modifies the default credit option for CPU usage to
`unlimited` for all instances in the T2 instance family in the
specified Region.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=ModifyDefaultCreditSpecification
&Region=us-east-1
&InstanceFamily=t2
&CpuCredits=unlimited
&AUTHPARAMS
```

#### Sample Response

```

<ModifyDefaultCreditSpecificationResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>11111111-2222-3333-4444-5555EXAMPLE</requestId>
    <instanceFamilyCreditSpecification>
        <cpuCredits>unlimited</cpuCredits>
        <instanceFamily>t2</instanceFamily>
    </instanceFamilyCreditSpecification>
</ModifyDefaultCreditSpecificationResponse>
```

### Example 2

This example modifies the default credit option for CPU usage to
`standard` for all instances in the T3 instance family in the
specified Region.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=ModifyDefaultCreditSpecification
&Region=us-east-1
&InstanceFamily=t3
&CpuCredits=standard
&AUTHPARAMS
```

#### Sample Response

```

<ModifyDefaultCreditSpecificationResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>11111111-2222-3333-4444-5555EXAMPLE</requestId>
    <instanceFamilyCreditSpecification>
        <cpuCredits>standard</cpuCredits>
        <instanceFamily>t3</instanceFamily>
    </instanceFamilyCreditSpecification>
</ModifyDefaultCreditSpecificationResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/ModifyDefaultCreditSpecification)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/ModifyDefaultCreditSpecification)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/modifydefaultcreditspecification.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/modifydefaultcreditspecification.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/modifydefaultcreditspecification.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/modifydefaultcreditspecification.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/modifydefaultcreditspecification.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/modifydefaultcreditspecification.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/ModifyDefaultCreditSpecification)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/modifydefaultcreditspecification.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ModifyClientVpnEndpoint

ModifyEbsDefaultKmsKeyId
