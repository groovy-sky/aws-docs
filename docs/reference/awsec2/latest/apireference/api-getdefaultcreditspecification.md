# GetDefaultCreditSpecification

Describes the default credit option for CPU usage of a burstable performance instance
family.

For more information, see [Burstable\
performance instances](../../../../services/ec2/latest/userguide/burstable-performance-instances.md) in the _Amazon EC2 User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

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

### Example

This example gets the default credit option for CPU usage of instances in the
T2 instance family in the specified Region.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=GetDefaultCreditSpecification
&Region=us-east-1
&InstanceFamily=t2
&AUTHPARAMS
```

#### Sample Response

```

<GetDefaultCreditSpecificationResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>11111111-2222-3333-4444-5555EXAMPLE</requestId>
    <instanceFamilyCreditSpecification>
        <cpuCredits>unlimited</cpuCredits>
        <instanceFamily>t2</instanceFamily>
    </instanceFamilyCreditSpecification>
</GetDefaultCreditSpecificationResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/GetDefaultCreditSpecification)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/GetDefaultCreditSpecification)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/getdefaultcreditspecification.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/getdefaultcreditspecification.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/getdefaultcreditspecification.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/getdefaultcreditspecification.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/getdefaultcreditspecification.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/getdefaultcreditspecification.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/GetDefaultCreditSpecification)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/getdefaultcreditspecification.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

GetDeclarativePoliciesReportSummary

GetEbsDefaultKmsKeyId
