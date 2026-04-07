# DeleteTags

Deletes the specified set of tags from the specified set of resources.

To list the current tags, use [DescribeTags](api-describetags.md). For more information about
tags, see [Tag\
your Amazon EC2 resources](../../../../services/ec2/latest/userguide/using-tags.md) in the _Amazon Elastic Compute Cloud User_
_Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**ResourceId.N**

The IDs of the resources, separated by spaces.

Constraints: Up to 1000 resource IDs. We recommend breaking up this request into smaller batches.

Type: Array of strings

Required: Yes

**Tag.N**

The tags to delete. Specify a tag key and an optional tag value to delete
specific tags. If you specify a tag key without a tag value, we delete any tag with this
key regardless of its value. If you specify a tag key with an empty string as the tag
value, we delete the tag only if its value is an empty string.

If you omit this parameter, we delete all user-defined tags for the specified
resources. We do not delete AWS-generated tags (tags that have the `aws:`
prefix).

Constraints: Up to 1000 tags.

Type: Array of [Tag](api-tag.md) objects

Required: No

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**return**

Is `true` if the request succeeds, and an error otherwise.

Type: Boolean

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example deletes all the user-defined tags for the AMI with the ID
`ami-1a2b3c4d`.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DeleteTags
&ResourceId.1=ami-1a2b3c4d
&AUTHPARAMS
```

#### Sample Response

```

<DeleteTagsResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
   <requestId>7a62c49f-347e-4fc4-9331-6e8eEXAMPLE</requestId>
   <return>true</return>
</DeleteTagsResponse>
```

### Example

This example deletes the `stack` and `webserver` tags for two
particular instances.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DeleteTags
&ResourceId.1=i-1234567890abcdef0
&ResourceId.2=i-0598c7d356eba48d7
&Tag.1.Key=stack
&Tag.2.Key=webserver
&AUTHPARAMS
```

### Example

You can specify a tag key without a corresponding tag value to delete the tag
regardless of its value. This example request deletes all tags that have a key of
`Purpose`, regardless of the tag value.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DeleteTags
&ResourceId.1=i-0598c7d356eba48d7
&Tag.1.Key=Purpose
&AUTHPARAMS
```

### Example

When you create a tag, you can set the tag value to the empty
string. Correspondingly, you can delete only tags that have a specific key and whose value is
the empty string. This example request deletes all tags for the specified instance where the
key is `Purpose` and the tag value is the empty string.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DeleteTags
&ResourceId.1=i-1234567890abcdef0
&Tag.1.Key=Purpose
&Tag.2.Value=
&AUTHPARAMS
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DeleteTags)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DeleteTags)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/DeleteTags)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ec2-2016-11-15/DeleteTags)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/DeleteTags)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ec2-2016-11-15/DeleteTags)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ec2-2016-11-15/DeleteTags)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ec2-2016-11-15/DeleteTags)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DeleteTags)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/DeleteTags)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DeleteSubnetCidrReservation

DeleteTrafficMirrorFilter
