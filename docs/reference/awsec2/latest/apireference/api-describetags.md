# DescribeTags

Describes the specified tags for your EC2 resources.

For more information about tags, see [Tag your Amazon EC2 resources](../../../../services/ec2/latest/userguide/using-tags.md) in the
_Amazon Elastic Compute Cloud User Guide_.

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

- `key` \- The tag key.

- `resource-id` \- The ID of the resource.

- `resource-type` \- The resource type. For a list of possible values, see
[TagSpecification](api-tagspecification.md).

- `tag`:<key> - The key/value combination of the tag. For example,
specify "tag:Owner" for the filter name and "TeamA" for the filter value to find
resources with the tag "Owner=TeamA".

- `value` \- The tag value.

Type: Array of [Filter](api-filter.md) objects

Required: No

**MaxResults**

The maximum number of items to return for this request. This value can be between 5 and 1000.
To get the next page of items, make another request with the token returned in the output.
For more information, see [Pagination](query-requests.md#api-pagination).

Type: Integer

Required: No

**NextToken**

The token returned from a previous paginated request.
Pagination continues from the end of the items returned by the previous request.

Type: String

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

**tagSet**

The tags.

Type: Array of [TagDescription](api-tagdescription.md) objects

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example describes all the tags in your account.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribeTags
&AUTHPARAMS
```

#### Sample Response

```

<DescribeTagsResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/"/">
   <requestId>7a62c49f-347e-4fc4-9331-6e8eEXAMPLE</requestId>
   <tagSet>
      <item>
         <resourceId>ami-1a2b3c4d</resourceId>
         <resourceType>image</resourceType>
         <key>webserver</key>
         <value/>
      </item>
       <item>
         <resourceId>ami-1a2b3c4d</resourceId>
         <resourceType>image</resourceType>
         <key>stack</key>
         <value>Production</value>
      </item>
      <item>
         <resourceId>i-1234567890abcdef0</resourceId>
         <resourceType>instance</resourceType>
         <key>webserver</key>
         <value/>
      </item>
       <item>
         <resourceId>i-1234567890abcdef0</resourceId>
         <resourceType>instance</resourceType>
         <key>stack</key>
         <value>Production</value>
      </item>
      <item>
         <resourceId>i-0598c7d356eba48d7</resourceId>
         <resourceType>instance</resourceType>
         <key>database_server</key>
         <value/>
      </item>
       <item>
         <resourceId>i-0598c7d356eba48d7</resourceId>
         <resourceType>instance</resourceType>
         <key>stack</key>
         <value>Test</value>
      </item>
    </tagSet>
</DescribeTagsResponse>
```

### Example

This example describes only the tags for the AMI with ID ami-1a2b3c4d.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribeTags
&Filter.1.Name=resource-id
&Filter.1.Value.1=ami-1a2b3c4d
&AUTHPARAMS
```

#### Sample Response

```

<DescribeTagsResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/"/">
   <requestId>7a62c49f-347e-4fc4-9331-6e8eEXAMPLE</requestId>
   <tagSet>
      <item>
         <resourceId>ami-1a2b3c4d</resourceId>
         <resourceType>image</resourceType>
         <key>webserver</key>
         <value/>
      </item>
      <item>
         <resourceId>ami-1a2b3c4d</resourceId>
         <resourceType>image</resourceType>
         <key>stack</key>
         <value>Production</value>
      </item>
    </tagSet>
</DescribeTagsResponse>
```

### Example

This example describes the tags for all your instances.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribeTags
&Filter.1.Name=resource-type
&Filter.1.Value.1=instance
&AUTHPARAMS
```

#### Sample Response

```

<DescribeTagsResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/"/">
   <requestId>7a62c49f-347e-4fc4-9331-6e8eEXAMPLE</requestId>
   <tagSet>
       <item>
         <resourceId>i-0598c7d356eba48d7</resourceId>
         <resourceType>instance</resourceType>
         <key>webserver</key>
         <value/>
      </item>
       <item>
         <resourceId>i-0598c7d356eba48d7</resourceId>
         <resourceType>instance</resourceType>
         <key>stack</key>
         <value>Production</value>
      </item>
      <item>
         <resourceId>i-1234567890abcdef0</resourceId>
         <resourceType>instance</resourceType>
         <key>database_server</key>
         <value/>
      </item>
       <item>
         <resourceId>i-1234567890abcdef0</resourceId>
         <resourceType>instance</resourceType>
         <key>stack</key>
         <value>Test</value>
      </item>
    </tagSet>
</DescribeTagsResponse>
```

### Example

This example describes the tags for all your instances tagged with the key
_webserver_. You can use wildcards with filters, so you could
specify the value as `?ebserver` to find tags with the key
_webserver_ or _Webserver_.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribeTags
&Filter.1.Name=key
&Filter.1.Value.1=webserver
&AUTHPARAMS
```

#### Sample Response

```

<DescribeTagsResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/"/">
   <requestId>7a62c49f-347e-4fc4-9331-6e8eEXAMPLE</requestId>
   <tagSet>
       <item>
         <resourceId>i-1234567890abcdef0</resourceId>
         <resourceType>instance</resourceType>
         <key>webserver</key>
         <value/>
      </item>
    </tagSet>
</DescribeTagsResponse>
```

### Example

This example describes the tags for all your instances tagged with either
stack=Test or stack=Production.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribeTags
&Filter.1.Name=resource-type
&Filter.1.Value.1=instance
&Filter.2.Name=key
&Filter.2.Value.1=stack
&Filter.3.Name=value
&Filter.3.Value.1=Test
&Filter.3.Value.2=Production
&AUTHPARAMS
```

#### Sample Response

```

<DescribeTagsResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/"/">
   <requestId>7a62c49f-347e-4fc4-9331-6e8eEXAMPLE</requestId>
   <tagSet>
       <item>
         <resourceId>i-1234567890abcdef0</resourceId>
         <resourceType>instance</resourceType>
         <key>stack</key>
         <value>Production</value>
      </item>
       <item>
         <resourceId>i-0598c7d356eba48d7</resourceId>
         <resourceType>instance</resourceType>
         <key>stack</key>
         <value>Test</value>
      </item>
    </tagSet>
</DescribeTagsResponse>
```

### Example

This example describes the tags for all your instances tagged with Purpose=\[empty
string\].

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribeTags
&Filter.1.Name=resource-type
&Filter.1.Value.1=instance
&Filter.2.Name=key
&Filter.2.Value.1=Purpose
&Filter.3.Name=value
&Filter.3.Value.1=
&AUTHPARAMS
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DescribeTags)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DescribeTags)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describetags.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describetags.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describetags.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describetags.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describetags.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describetags.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DescribeTags)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describetags.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeSubnets

DescribeTrafficMirrorFilterRules
