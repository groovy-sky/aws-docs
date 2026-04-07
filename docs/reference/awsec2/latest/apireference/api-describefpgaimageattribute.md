# DescribeFpgaImageAttribute

Describes the specified attribute of the specified Amazon FPGA Image (AFI).

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**Attribute**

The AFI attribute.

Type: String

Valid Values: `description | name | loadPermission | productCodes`

Required: Yes

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**FpgaImageId**

The ID of the AFI.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**fpgaImageAttribute**

Information about the attribute.

Type: [FpgaImageAttribute](api-fpgaimageattribute.md) object

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example describes the load permissions for the specified AFI.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribeFpgaImageAttribute
&FpgaImageId=afi-0d123e21abcc85abc
&Attribute=loadPermission
&AUTHPARAMS
```

#### Sample Response

```

<DescribeFpgaImageAttributeResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>19106033-3723-481e-8cc4-aedexample</requestId>
    <fpgaImageAttribute>
        <fpgaImageId>afi-0d123e21abcc85abc</fpgaImageId>
        <loadPermissions>
            <item>
                <userId>123456789012</userId>
            </item>
        </loadPermissions>
    </fpgaImageAttribute>
</DescribeFpgaImageAttributeResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DescribeFpgaImageAttribute)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DescribeFpgaImageAttribute)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describefpgaimageattribute.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describefpgaimageattribute.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describefpgaimageattribute.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describefpgaimageattribute.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describefpgaimageattribute.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describefpgaimageattribute.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DescribeFpgaImageAttribute)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describefpgaimageattribute.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeFlowLogs

DescribeFpgaImages
