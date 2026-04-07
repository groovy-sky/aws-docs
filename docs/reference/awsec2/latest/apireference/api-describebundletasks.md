# DescribeBundleTasks

Describes the specified bundle tasks or all of your bundle tasks.

###### Note

Completed bundle tasks are listed for only a limited time. If your bundle task is no
longer in the list, you can still register an AMI from it. Just use
`RegisterImage` with the Amazon S3 bucket name and image manifest name you provided
to the bundle task.

###### Note

The order of the elements in the response, including those within nested structures,
might vary. Applications should not assume the elements appear in a particular order.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**BundleId.N**

The bundle task IDs.

Default: Describes all your bundle tasks.

Type: Array of strings

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is
`DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**Filter.N**

The filters.

- `bundle-id` \- The ID of the bundle task.

- `error-code` \- If the task failed, the error code returned.

- `error-message` \- If the task failed, the error message returned.

- `instance-id` \- The ID of the instance.

- `progress` \- The level of task completion, as a percentage (for example,
20%).

- `s3-bucket` \- The Amazon S3 bucket to store the AMI.

- `s3-prefix` \- The beginning of the AMI name.

- `start-time` \- The time the task started (for example,
2013-09-15T17:15:20.000Z).

- `state` \- The state of the task ( `pending` \|
`waiting-for-shutdown` \| `bundling` \| `storing` \|
`cancelling` \| `complete` \| `failed`).

- `update-time` \- The time of the most recent update for the task.

Type: Array of [Filter](api-filter.md) objects

Required: No

## Response Elements

The following elements are returned by the service.

**bundleInstanceTasksSet**

Information about the bundle tasks.

Type: Array of [BundleTask](api-bundletask.md) objects

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example 1

This example describes the status of the specified bundle task.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribeBundleTasks
&bundleId.1=bun-c1a540a8
&AUTHPARAMS
```

#### Sample Response

```

<DescribeBundleTasksResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
   <requestId>59dbff89-35bd-4eac-99ed-be587EXAMPLE</requestId>
   <bundleInstanceTasksSet>
      <item>
         <instanceId>i-1234567890abcdef0</instanceId>
         <bundleId>bun-c1a540a8</bundleId>
         <state>cancelling</state>
         <startTime>2008-10-07T11:41:50.000Z</startTime>
         <updateTime>2008-10-07T11:51:50.000Z</updateTime>
         <storage>
            <S3>
               <bucket>myawsbucket</bucket>
               <prefix>winami</prefix>
            </S3>
         </storage>
         <progress>20%</progress>
      </item>
   <bundleInstanceTasksSet>
</DescribeBundleTasksResponse>
```

### Example 2

This example filters the response to include only bundle tasks whose state is either
`complete` or `failed`, and in addition are targeted for the Amazon S3
bucket named `myawsbucket`.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribeBundleTasks
&Filter.1.Name=s3-bucket
&Filter.1.Value.1=myawsbucket
&Filter.2.Name=state
&Filter.2.Name.1=complete
&Filter.2.Name.2=failed
&AUTHPARAMS
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DescribeBundleTasks)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DescribeBundleTasks)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describebundletasks.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describebundletasks.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describebundletasks.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describebundletasks.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describebundletasks.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describebundletasks.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DescribeBundleTasks)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describebundletasks.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeAwsNetworkPerformanceMetricSubscriptions

DescribeByoipCidrs
