# GetConsoleOutput

Gets the console output for the specified instance. For Linux instances, the instance
console output displays the exact console output that would normally be displayed on a
physical monitor attached to a computer. For Windows instances, the instance console
output includes the last three system event log errors.

For more information, see [Instance\
console output](../../../../services/ec2/latest/userguide/instance-console.md#instance-console-console-output) in the _Amazon EC2 User Guide_.

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

**Latest**

When enabled, retrieves the latest console output for the instance.

Default: disabled ( `false`)

Type: Boolean

Required: No

## Response Elements

The following elements are returned by the service.

**instanceId**

The ID of the instance.

Type: String

**output**

The console output, base64-encoded. If you are using a command line tool, the tool
decodes the output for you.

Type: String

**requestId**

The ID of the request.

Type: String

**timestamp**

The time at which the output was last updated.

Type: Timestamp

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example 1

This example retrieves the console output for the specified instance.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=GetConsoleOutput
&InstanceId=i-1234567890abcdef0
&AUTHPARAMS
```

#### Sample Response

```

<GetConsoleOutputResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
  <requestId>59dbff89-35bd-4eac-99ed-be587EXAMPLE</requestId>
  <instanceId>i-1234567890abcdef0</instanceId>
  <timestamp>2022-10-14T01:12:41.000Z</timestamp>
  <output>TGludXggdmVyc2lvbiAyLjYuMTYteGVuVSAoYnVpbGRlckBwYXRjaGJhdC5hbWF6b25zYSkgKGdj
YyB2ZXJzaW9uIDQuMC4xIDIwMDUwNzI3IChSZWQgSGF0IDQuMC4xLTUpKSAjMSBTTVAgVGh1IE9j
dCAyNiAwODo0MToyNiBTQVNUIDIwMDYKQklPUy1wcm92aWRlZCBwaHlzaWNhbCBSQU0gbWFwOgpY
ZW46IDAwMDAwMDAwMDAwMDAwMDAgLSAwMDAwMDAwMDZhNDAwMDAwICh1c2FibGUpCjk4ME1CIEhJ
R0hNRU0gYXZhaWxhYmxlLgo3MjdNQiBMT1dNRU0gYXZhaWxhYmxlLgpOWCAoRXhlY3V0ZSBEaXNh
YmxlKSBwcm90ZWN0aW9uOiBhY3RpdmUKSVJRIGxvY2t1cCBkZXRlY3Rpb24gZGlzYWJsZWQKQnVp
bHQgMSB6b25lbGlzdHMKS2VybmVsIGNvbW1hbmQgbGluZTogcm9vdD0vZGV2L3NkYTEgcm8gNApF
bmFibGluZyBmYXN0IEZQVSBzYXZlIGFuZCByZXN0b3JlLi4uIGRvbmUuCg==</output>
</GetConsoleOutputResponse>
```

### Example 2

This example retrieves the latest console output for the specified
instance.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=GetConsoleOutput
&InstanceId=i-1234567890abcdef0
&Latest=true
&AUTHPARAMS
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/getconsoleoutput.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/getconsoleoutput.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/getconsoleoutput.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/getconsoleoutput.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/getconsoleoutput.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/getconsoleoutput.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/getconsoleoutput.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/getconsoleoutput.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/getconsoleoutput.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/getconsoleoutput.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

GetCoipPoolUsage

GetConsoleScreenshot
