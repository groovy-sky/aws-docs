# DescribePlacementGroups

Describes the specified placement groups or all of your placement groups.

###### Note

To describe a specific placement group that is _shared_ with
your account, you must specify the ID of the placement group using the
`GroupId` parameter. Specifying the name of a
_shared_ placement group using the `GroupNames`
parameter will result in an error.

For more information, see [Placement groups](../../../../services/ec2/latest/userguide/placement-groups.md) in the
_Amazon EC2 User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the operation, without actually making the
request, and provides an error response. If you have the required permissions, the error response is
`DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**Filter.N**

The filters.

- `group-name` \- The name of the placement group.

- `group-arn` \- The Amazon Resource Name (ARN) of the placement
group.

- `spread-level` \- The spread level for the placement group
( `host` \| `rack`).

- `state` \- The state of the placement group ( `pending` \|
`available` \| `deleting` \|
`deleted`).

- `strategy` \- The strategy of the placement group
( `cluster` \| `spread` \|
`partition`).

- `tag:<key>` \- The key/value combination of a tag assigned to the resource. Use the tag key in the filter name and the tag value as the filter value.
For example, to find all resources that have a tag with the key `Owner` and the value `TeamA`, specify `tag:Owner` for the filter name and `TeamA` for the filter value.

- `tag-key` \- The key of a tag assigned to the resource. Use this filter to find all resources that have a tag with a specific key, regardless of the tag value.

Type: Array of [Filter](api-filter.md) objects

Required: No

**GroupId.N**

The IDs of the placement groups.

Type: Array of strings

Required: No

**GroupName.N**

The names of the placement groups.

Constraints:

- You can specify a name only if the placement group is owned by your
account.

- If a placement group is _shared_ with your account,
specifying the name results in an error. You must use the `GroupId`
parameter instead.

Type: Array of strings

Required: No

## Response Elements

The following elements are returned by the service.

**placementGroupSet**

Information about the placement groups.

Type: Array of [PlacementGroup](api-placementgroup.md) objects

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example 1

This example describes the placement group named
`ABC-spread`.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribePlacementGroups
&GroupName.1=ABC-spread
&AUTHPARAMS
```

#### Sample Response

```

<DescribePlacementGroupsResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
   <requestID>d4904fd9-82c2-4ea5-adfe-a9cc3EXAMPLE</requestID>
   <placementGroupSet>
      <item>
         <groupName>ABC-spread</groupName>
         <spreadLevel>rack</spreadLevel>
         <strategy>spread</strategy>
         <state>available</state>
      </item>
   </placementGroupSet>
</DescribePlacementGroupsResponse>
```

### Example 2

This example filters the response to include only placement groups that
include the string `Project` in the name.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribePlacementGroups
&Filter.1.Name=group-name
&Filter.1.Value=*Project*
&AUTHPARAMS
```

#### Sample Response

```

<DescribePlacementGroupsResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
   <requestID>d4904fd9-82c2-4ea5-adfe-a9cc3EXAMPLE</requestID>
   <placementGroupSet>
      <item>
         <groupName>Project-cluster</groupName>
         <strategy>cluster</strategy>
         <state>available</state>
      </item>
   </placementGroupSet>
</DescribePlacementGroupsResponse>
```

### Example 3

This example describes the partition placement group named
`HDSF-Group-A` with three partitions.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribePlacementGroups
&GroupName.1=HDSF-Group-A
&AUTHPARAMS
```

#### Sample Response

```

<DescribePlacementGroupsResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
   <requestID>d4904fd9-82c2-4ea5-adfe-a9cc3EXAMPLE</requestID>
   <placementGroupSet>
      <item>
         <groupName>HDSF-Group-A</groupName>
         <strategy>partition</strategy>
         <partitionCount>3</partitionCount>
         <state>available</state>
      </item>
   </placementGroupSet>
</DescribePlacementGroupsResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/describeplacementgroups.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/describeplacementgroups.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describeplacementgroups.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describeplacementgroups.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describeplacementgroups.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describeplacementgroups.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describeplacementgroups.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describeplacementgroups.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/describeplacementgroups.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describeplacementgroups.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeOutpostLags

DescribePrefixLists
