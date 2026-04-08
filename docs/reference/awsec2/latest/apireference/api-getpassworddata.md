# GetPasswordData

Retrieves the encrypted administrator password for a running Windows instance.

The Windows password is generated at boot by the `EC2Config` service or
`EC2Launch` scripts (Windows Server 2016 and later). This usually only
happens the first time an instance is launched. For more information, see [EC2Config](../../../../services/ec2/latest/userguide/usingconfig-winami.md) and [EC2Launch](../../../../services/ec2/latest/userguide/ec2launch.md) in the
_Amazon EC2 User Guide_.

For the `EC2Config` service, the password is not generated for rebundled
AMIs unless `Ec2SetPassword` is enabled before bundling.

The password is encrypted using the key pair that you specified when you launched the
instance. You must provide the corresponding key pair file.

When you launch an instance, password generation and encryption may take a few
minutes. If you try to retrieve the password before it's available, the output returns
an empty string. We recommend that you wait up to 15 minutes after launching an instance
before trying to retrieve the generated password.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the operation, without actually making the
request, and provides an error response. If you have the required permissions, the error response is
`DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**InstanceId**

The ID of the Windows instance.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**instanceId**

The ID of the Windows instance.

Type: String

**passwordData**

The password of the instance. Returns an empty string if the password is not
available.

Type: String

**requestId**

The ID of the request.

Type: String

**timestamp**

The time the data was last updated.

Type: Timestamp

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example returns the encrypted version of the administrator password for
the specified instance.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=GetPasswordData
&InstanceId=i-1234567890abcdef0
&AUTHPARAMS
```

#### Sample Response

```

<GetPasswordDataResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
  <requestId>59dbff89-35bd-4eac-99ed-be587EXAMPLE</requestId>
  <instanceId>i-1234567890abcdef0</instanceId>
  <timestamp>2009-10-24 15:00:00</timestamp>
  <passwordData>TGludXggdmVyc2lvbiAyLjYuMTYteGVuVSAoYnVpbGRlckBwYXRjaGJhdC5hbWF6b25zYSkgKGdj</passwordData>
</GetPasswordDataResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/getpassworddata.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/getpassworddata.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/getpassworddata.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/getpassworddata.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/getpassworddata.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/getpassworddata.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/getpassworddata.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/getpassworddata.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/getpassworddata.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/getpassworddata.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

GetNetworkInsightsAccessScopeContent

GetReservedInstancesExchangeQuote
