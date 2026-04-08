# StopInstances

Stops an Amazon EBS-backed instance. You can restart your instance at any time using
the [StartInstances](api-startinstances.md) API. For more information, see [Stop and start Amazon EC2\
instances](../../../../services/ec2/latest/userguide/stop-start.md) in the _Amazon EC2 User Guide_.

When you stop or hibernate an instance, we shut it down. By default, this includes a
graceful operating system (OS) shutdown. To bypass the graceful shutdown, use the
`skipOsShutdown` parameter; however, this might risk data
integrity.

You can use the StopInstances operation together with the `Hibernate`
parameter to hibernate an instance if the instance is [enabled for\
hibernation](../../../../services/ec2/latest/userguide/enabling-hibernation.md) and meets the [hibernation\
prerequisites](../../../../services/ec2/latest/userguide/hibernating-prerequisites.md). Stopping an instance doesn't preserve data stored in RAM,
while hibernation does. If hibernation fails, a normal shutdown occurs. For more
information, see [Hibernate your Amazon EC2\
instance](../../../../services/ec2/latest/userguide/hibernate.md) in the _Amazon EC2 User Guide_.

If your instance appears stuck in the `stopping` state, there might be an
issue with the underlying host computer. You can use the StopInstances operation
together with the Force parameter to force stop your instance. For more information, see
[Troubleshoot\
Amazon EC2 instance stop issues](../../../../services/ec2/latest/userguide/troubleshootinginstancesstopping.md) in the
_Amazon EC2 User Guide_.

Stopping and hibernating an instance differs from rebooting or terminating it. For
example, a stopped or hibernated instance retains its root volume and any data volumes,
unlike terminated instances where these volumes are automatically deleted. For more
information about the differences between stopping, hibernating, rebooting, and
terminating instances, see [Amazon EC2\
instance state changes](../../../../services/ec2/latest/userguide/ec2-instance-lifecycle.md) in the _Amazon EC2 User Guide_.

We don't charge for instance usage or data transfer fees when an instance is stopped.
However, the root volume and any data volumes remain and continue to persist your data,
and you're charged for volume usage. Every time you start your instance, Amazon EC2 charges a one-minute minimum for instance usage, followed by per-second
billing.

You can't stop or hibernate instance store-backed instances.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the operation, without actually making the
request, and provides an error response. If you have the required permissions, the error response is
`DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**Force**

Forces the instance to stop. The instance will first attempt a graceful shutdown,
which includes flushing file system caches and metadata. If the graceful shutdown fails
to complete within the timeout period, the instance shuts down forcibly without flushing
the file system caches and metadata.

After using this option, you must perform file system check and repair procedures.
This option is not recommended for Windows instances. For more information, see [Troubleshoot\
Amazon EC2 instance stop issues](../../../../services/ec2/latest/userguide/troubleshootinginstancesstopping.md) in the
_Amazon EC2 User Guide_.

Default: `false`

Type: Boolean

Required: No

**Hibernate**

Hibernates the instance if the instance was enabled for hibernation at launch. If the
instance cannot hibernate successfully, a normal shutdown occurs. For more information,
see [Hibernate\
your Amazon EC2 instance](../../../../services/ec2/latest/userguide/hibernate.md) in the
_Amazon EC2 User Guide_.

Default: `false`

Type: Boolean

Required: No

**InstanceId.N**

The IDs of the instances.

Type: Array of strings

Required: Yes

**SkipOsShutdown**

Specifies whether to bypass the graceful OS shutdown process when the instance is
stopped.

###### Important

Bypassing the graceful OS shutdown might result in data loss or corruption (for
example, memory contents not flushed to disk or loss of in-flight IOs) or skipped
shutdown scripts.

Default: `false`

Type: Boolean

Required: No

## Response Elements

The following elements are returned by the service.

**instancesSet**

Information about the stopped instances.

Type: Array of [InstanceStateChange](api-instancestatechange.md) objects

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example stops the specified instance.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=StopInstances
&InstanceId.1=i-1234567890abcdef0
&AUTHPARAMS
```

#### Sample Response

```

<StopInstancesResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
  <requestId>59dbff89-35bd-4eac-99ed-be587EXAMPLE</requestId>
  <instancesSet>
    <item>
      <instanceId>i-1234567890abcdef0</instanceId>
      <currentState>
          <code>64</code>
          <name>stopping</name>
      </currentState>
      <previousState>
          <code>16</code>
          <name>running</name>
      </previousState>
    </item>
  </instancesSet>
</StopInstancesResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/stopinstances.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/stopinstances.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/stopinstances.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/stopinstances.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/stopinstances.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/stopinstances.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/stopinstances.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/stopinstances.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/stopinstances.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/stopinstances.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

StartVpcEndpointServicePrivateDnsVerification

TerminateClientVpnConnections
