# ModifyHosts

Modify the auto-placement setting of a Dedicated Host. When auto-placement is enabled,
any instances that you launch with a tenancy of `host` but without a specific
host ID are placed onto any available Dedicated Host in your account that has
auto-placement enabled. When auto-placement is disabled, you need to provide a host ID
to have the instance launch onto a specific host. If no host ID is provided, the
instance is launched onto a suitable host with auto-placement enabled.

You can also use this API action to modify a Dedicated Host to support either multiple
instance types in an instance family, or to support a specific instance type
only.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**AutoPlacement**

Specify whether to enable or disable auto-placement.

Type: String

Valid Values: `on | off`

Required: No

**HostId.N**

The IDs of the Dedicated Hosts to modify.

Type: Array of strings

Required: Yes

**HostMaintenance**

Indicates whether to enable or disable host maintenance for the Dedicated Host. For
more information, see [Host\
maintenance](../../../../services/ec2/latest/userguide/dedicated-hosts-maintenance.md) in the _Amazon EC2 User Guide_.

Type: String

Valid Values: `on | off`

Required: No

**HostRecovery**

Indicates whether to enable or disable host recovery for the Dedicated Host. For more
information, see [Host recovery](../../../../services/ec2/latest/userguide/dedicated-hosts-recovery.md) in
the _Amazon EC2 User Guide_.

Type: String

Valid Values: `on | off`

Required: No

**InstanceFamily**

Specifies the instance family to be supported by the Dedicated Host. Specify this
parameter to modify a Dedicated Host to support multiple instance types within its
current instance family.

If you want to modify a Dedicated Host to support a specific instance type only, omit
this parameter and specify **InstanceType** instead. You
cannot specify **InstanceFamily** and **InstanceType** in the same request.

Type: String

Required: No

**InstanceType**

Specifies the instance type to be supported by the Dedicated Host. Specify this
parameter to modify a Dedicated Host to support only a specific instance type.

If you want to modify a Dedicated Host to support multiple instance types in its
current instance family, omit this parameter and specify **InstanceFamily** instead. You cannot specify **InstanceType** and **InstanceFamily** in the
same request.

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**successful**

The IDs of the Dedicated Hosts that were successfully modified.

Type: Array of strings

**unsuccessful**

The IDs of the Dedicated Hosts that could not be modified. Check whether the setting
you requested can be used.

Type: Array of [UnsuccessfulItem](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_UnsuccessfulItem.html) objects

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example 1

This example enables the auto-placement setting on a Dedicated Host.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=ModifyHosts
&AutoPlacement=on
&HostId=h-00548908djdsgfs
&AUTHPARAMS
```

#### Sample Response

```

<ModifyHostsResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>d4904fd9-82c2-4ea5-adfe-a9cc3EXAMPLE</requestId>
    <unsuccessful/>
    <successful>
        <item>h-00548908djdsgfs</item>
    </successful>
</ModifyHostsResponse>
```

### Example 2

This example enables host recovery on a Dedicated Host.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=ModifyHosts
&HostRecovery=on
&HostId=h-00548908djdsgfs
&AUTHPARAMS
```

#### Sample Response

```

<ModifyHostsResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>d4904fd9-82c2-4ea5-adfe-a9cc3EXAMPLE</requestId>
    <unsuccessful/>
    <successful>
        <item>h-00548908djdsgfs</item>
    </successful>
</ModifyHostsResponse>
```

### Example 3

This example modifies a Dedicated Host that supports only
`m5.large` instances to support multiple instance types in the
`m5` instance family.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=ModifyHosts
&InstanceFamily=m5
&HostId=h-00548908djdsgfs
&AUTHPARAMS
```

#### Sample Response

```

<ModifyHostsResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>d4904fd9-82c2-4ea5-adfe-a9cc3EXAMPLE</requestId>
    <unsuccessful/>
    <successful>
        <item>h-00548908djdsgfs</item>
    </successful>
</ModifyHostsResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/ModifyHosts)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/ModifyHosts)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/ModifyHosts)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ec2-2016-11-15/ModifyHosts)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/ModifyHosts)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ec2-2016-11-15/ModifyHosts)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ec2-2016-11-15/ModifyHosts)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ec2-2016-11-15/ModifyHosts)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/ModifyHosts)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/ModifyHosts)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ModifyFpgaImageAttribute

ModifyIdentityIdFormat
