# SendDiagnosticInterrupt

Sends a diagnostic interrupt to the specified Amazon EC2 instance to trigger a
_kernel panic_ (on Linux instances), or a _blue_
_screen_/ _stop error_ (on Windows instances). For
instances based on Intel and AMD processors, the interrupt is received as a
_non-maskable interrupt_ (NMI).

In general, the operating system crashes and reboots when a kernel panic or stop error
is triggered. The operating system can also be configured to perform diagnostic tasks,
such as generating a memory dump file, loading a secondary kernel, or obtaining a call
trace.

Before sending a diagnostic interrupt to your instance, ensure that its operating
system is configured to perform the required diagnostic tasks.

For more information about configuring your operating system to generate a crash dump
when a kernel panic or stop error occurs, see [Send a diagnostic interrupt\
(for advanced users)](../../../../services/ec2/latest/userguide/diagnostic-interrupt.md) in the _Amazon EC2 User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the operation, without actually making the
request, and provides an error response. If you have the required permissions, the error response is
`DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**InstanceId**

The ID of the instance.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**return**

Is `true` if the request succeeds, and an error otherwise.

Type: Boolean

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example sends a diagnostic interrupt to the specified instance.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=SendDiagnosticInterrupt
&InstanceId=i-1234567890abcdef0
&AUTHPARAMS
```

#### Sample Response

```

<SendDiagnosticInterruptResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
  <requestId>59dbff89-35bd-4eac-99ed-be587EXAMPLE</requestId>
  <return>true</return>
</SendDiagnosticInterruptResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/SendDiagnosticInterrupt)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/SendDiagnosticInterrupt)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/senddiagnosticinterrupt.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/senddiagnosticinterrupt.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/senddiagnosticinterrupt.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/senddiagnosticinterrupt.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/senddiagnosticinterrupt.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/senddiagnosticinterrupt.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/SendDiagnosticInterrupt)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/senddiagnosticinterrupt.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

SearchTransitGatewayRoutes

StartDeclarativePoliciesReport
