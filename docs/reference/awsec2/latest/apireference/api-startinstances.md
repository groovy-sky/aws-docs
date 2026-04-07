# StartInstances

Starts an Amazon EBS-backed instance that you've previously stopped.

Instances that use Amazon EBS volumes as their root devices can be quickly stopped and
started. When an instance is stopped, the compute resources are released and you are not
billed for instance usage. However, your root partition Amazon EBS volume remains and
continues to persist your data, and you are charged for Amazon EBS volume usage. You can
restart your instance at any time. Every time you start your instance, Amazon EC2
charges a one-minute minimum for instance usage, and thereafter charges per second for
instance usage.

Before stopping an instance, make sure it is in a state from which it can be
restarted. Stopping an instance does not preserve data stored in RAM.

Performing this operation on an instance that uses an instance store as its root
device returns an error.

If you attempt to start a T3 instance with `host` tenancy and the
`unlimited` CPU credit option, the request fails. The
`unlimited` CPU credit option is not supported on Dedicated Hosts. Before
you start the instance, either change its CPU credit option to `standard`, or
change its tenancy to `default` or `dedicated`.

For more information, see [Stop and start Amazon EC2\
instances](../../../../services/ec2/latest/userguide/stop-start.md) in the _Amazon EC2 User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**AdditionalInfo**

Reserved.

Type: String

Required: No

**DryRun**

Checks whether you have the required permissions for the operation, without actually making the
request, and provides an error response. If you have the required permissions, the error response is
`DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**InstanceId.N**

The IDs of the instances.

Type: Array of strings

Required: Yes

## Response Elements

The following elements are returned by the service.

**instancesSet**

Information about the started instances.

Type: Array of [InstanceStateChange](api-instancestatechange.md) objects

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example starts the specified instance.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=StartInstances
&InstanceId.1=i-1234567890abcdef0
&AUTHPARAMS
```

#### Sample Response

```

<StartInstancesResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
  <requestId>59dbff89-35bd-4eac-99ed-be587EXAMPLE</requestId>
  <instancesSet>
    <item>
      <instanceId>i-1234567890abcdef0</instanceId>
      <currentState>
          <code>0</code>
          <name>pending</name>
      </currentState>
      <previousState>
          <code>80</code>
          <name>stopped</name>
      </previousState>
    </item>
  </instancesSet>
</StartInstancesResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/StartInstances)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/StartInstances)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/startinstances.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/startinstances.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/startinstances.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/startinstances.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/startinstances.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/startinstances.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/StartInstances)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/startinstances.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

StartDeclarativePoliciesReport

StartNetworkInsightsAccessScopeAnalysis
