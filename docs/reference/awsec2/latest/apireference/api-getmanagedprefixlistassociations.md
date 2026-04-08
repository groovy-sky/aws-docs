# GetManagedPrefixListAssociations

Gets information about the resources that are associated with the specified managed prefix list.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**MaxResults**

The maximum number of results to return with a single call.
To retrieve the remaining results, make another call with the returned `nextToken` value.

Type: Integer

Valid Range: Minimum value of 5. Maximum value of 255.

Required: No

**NextToken**

The token for the next page of results.

Type: String

Required: No

**PrefixListId**

The ID of the prefix list.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**nextToken**

The token to use to retrieve the next page of results. This value is `null` when there are no more results to return.

Type: String

**prefixListAssociationSet**

Information about the associations.

Type: Array of [PrefixListAssociation](api-prefixlistassociation.md) objects

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example gets the associations for the specified managed prefix
list.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=GetManagedPrefixListAssociations
&PrefixListId=pl-0123123123123aabb
&AUTHPARAMS
```

#### Sample Response

```

<GetManagedPrefixListAssociationsResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>3ae1bcdf-4042-47f8-9b8a-example</requestId>
    <prefixListAssociationSet>
        <item>
            <resourceId>sg-01234567811abcaba</resourceId>
            <resourceOwner>123456789012</resourceOwner>
        </item>
    </prefixListAssociationSet>
</GetManagedPrefixListAssociationsResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/getmanagedprefixlistassociations.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/getmanagedprefixlistassociations.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/getmanagedprefixlistassociations.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/getmanagedprefixlistassociations.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/getmanagedprefixlistassociations.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/getmanagedprefixlistassociations.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/getmanagedprefixlistassociations.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/getmanagedprefixlistassociations.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/getmanagedprefixlistassociations.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/getmanagedprefixlistassociations.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

GetLaunchTemplateData

GetManagedPrefixListEntries
