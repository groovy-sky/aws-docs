# UnmonitorInstances

Disables detailed monitoring for a running instance. For more information, see [Monitoring\
your instances and volumes](../../../../services/ec2/latest/userguide/using-cloudwatch.md) in the
_Amazon EC2 User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

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

The monitoring information.

Type: Array of [InstanceMonitoring](api-instancemonitoring.md) objects

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example disables detailed monitoring for the specified instances.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=UnmonitorInstances
&InstanceId.1=i-1234567890abcdef0
&InstanceId.2=i-0598c7d356eba48d7
&AUTHPARAMS
```

#### Sample Response

```

<UnmonitorInstancesResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
   <requestId>59dbff89-35bd-4eac-99ed-be587EXAMPLE</requestId>
   <instancesSet>
      <item>
         <instanceId>i-1234567890abcdef0</instanceId>
         <monitoring>
            <state>disabled</state>
         </monitoring>
      </item>
      <item>
         <instanceId>i-0598c7d356eba48d7</instanceId>
         <monitoring>
            <state>disabled</state>
         </monitoring>
      </item>
   </instancesSet>
</UnmonitorInstancesResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/UnmonitorInstances)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/UnmonitorInstances)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/unmonitorinstances.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/unmonitorinstances.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/unmonitorinstances.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/unmonitorinstances.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/unmonitorinstances.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/unmonitorinstances.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/UnmonitorInstances)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/unmonitorinstances.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

UnlockSnapshot

UpdateCapacityManagerOrganizationsAccess
