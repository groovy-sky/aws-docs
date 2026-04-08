# DescribeVolumes

Describes the specified EBS volumes or all of your EBS volumes.

If you are describing a long list of volumes, we recommend that you paginate the output to make the list
more manageable. For more information, see [Pagination](query-requests.md#api-pagination).

For more information about EBS volumes, see [Amazon EBS volumes](../../../../services/ebs/latest/userguide/ebs-volumes.md) in the _Amazon EBS User Guide_.

###### Important

We strongly recommend using only paginated requests. Unpaginated requests are
susceptible to throttling and timeouts.

###### Note

The order of the elements in the response, including those within nested
structures, might vary. Applications should not assume the elements appear in a
particular order.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**Filter.N**

The filters.

- `attachment.attach-time` \- The time stamp when the attachment
initiated.

- `attachment.delete-on-termination` \- Whether the volume is deleted on
instance termination.

- `attachment.device` \- The device name specified in the block device mapping
(for example, `/dev/sda1`).

- `attachment.instance-id` \- The ID of the instance the volume is attached
to.

- `attachment.status` \- The attachment state ( `attaching` \|
`attached` \| `detaching`).

- `availability-zone` \- The Availability Zone in which the volume was
created.

- `availability-zone-id` \- The ID of the Availability Zone in which the
volume was created.

- `create-time` \- The time stamp when the volume was created.

- `encrypted` \- Indicates whether the volume is encrypted ( `true`
\| `false`)

- `fast-restored` \- Indicates whether the volume was created from a
snapshot that is enabled for fast snapshot restore ( `true` \|
`false`).

- `multi-attach-enabled` \- Indicates whether the volume is enabled for Multi-Attach ( `true`
\| `false`)

- `operator.managed` \- A Boolean that indicates whether this is a managed
volume.

- `operator.principal` \- The principal that manages the volume. Only valid
for managed volumes, where `managed` is `true`.

- `size` \- The size of the volume, in GiB.

- `snapshot-id` \- The snapshot from which the volume was created.

- `status` \- The state of the volume ( `creating` \|
`available` \| `in-use` \| `deleting` \|
`deleted` \| `error`).

- `tag`:<key> - The key/value combination of a tag assigned to the resource. Use the tag key in the filter name and the tag value as the filter value.
For example, to find all resources that have a tag with the key `Owner` and the value `TeamA`, specify `tag:Owner` for the filter name and `TeamA` for the filter value.

- `tag-key` \- The key of a tag assigned to the resource. Use this filter to find all resources assigned a tag with a specific key, regardless of the tag value.

- `volume-id` \- The volume ID.

- `volume-type` \- The Amazon EBS volume type ( `gp2` \| `gp3` \| `io1` \| `io2` \|
`st1` \| `sc1` \| `standard`)

Type: Array of [Filter](api-filter.md) objects

Required: No

**MaxResults**

The maximum number of items to return for this request.
To get the next page of items, make another request with the token returned in the output.
For more information, see [Pagination](query-requests.md#api-pagination).

Type: Integer

Required: No

**NextToken**

The token returned from a previous paginated request.
Pagination continues from the end of the items returned by the previous request.

Type: String

Required: No

**VolumeId.N**

The volume IDs. If not specified, then all volumes are included in the response.

Type: Array of strings

Required: No

## Response Elements

The following elements are returned by the service.

**nextToken**

The token to include in another request to get the next page of items.
This value is `null` when there are no more items to return.

Type: String

**requestId**

The ID of the request.

Type: String

**volumeSet**

Information about the volumes.

Type: Array of [Volume](api-volume.md) objects

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example describes all volumes associated with your account.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribeVolumes
&AUTHPARAMS
```

#### Sample Response

```nohighlight

<DescribeVolumesResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
   <requestId>59dbff89-35bd-4eac-99ed-be587EXAMPLE</requestId>
   <volumeSet>
      <item>
         <volumeId>vol-1234567890abcdef0</volumeId>
         <size>80</size>
         <snapshotId/>
         <availabilityZone>us-east-1a</availabilityZone>
         <status>in-use</status>
         <createTime>YYYY-MM-DDTHH:MM:SS.SSSZ</createTime>
         <attachmentSet>
            <item>
               <volumeId>vol-1234567890abcdef0</volumeId>
               <instanceId>i-1234567890abcdef0</instanceId>
               <device>/dev/sdh</device>
               <status>attached</status>
               <attachTime>YYYY-MM-DDTHH:MM:SS.SSSZ</attachTime>
               <deleteOnTermination>false</deleteOnTermination>
            </item>
         </attachmentSet>
         <volumeType>standard</volumeType>
         <encrypted>true</encrypted>
         <multiAttachEnabled>false</multiAttachEnabled>
      </item>
   </volumeSet>
</DescribeVolumesResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/describevolumes.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/describevolumes.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describevolumes.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describevolumes.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describevolumes.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describevolumes.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describevolumes.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describevolumes.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/describevolumes.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describevolumes.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeVolumeAttribute

DescribeVolumesModifications
