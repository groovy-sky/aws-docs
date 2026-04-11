# ReleaseHosts

When you no longer want to use an On-Demand Dedicated Host it can be released.
On-Demand billing is stopped and the host goes into `released` state. The
host ID of Dedicated Hosts that have been released can no longer be specified in another
request, for example, to modify the host. You must stop or terminate all instances on a
host before it can be released.

When Dedicated Hosts are released, it may take some time for them to stop counting
toward your limit and you may receive capacity errors when trying to allocate new
Dedicated Hosts. Wait a few minutes and then try again.

Released hosts still appear in a [DescribeHosts](api-describehosts.md) response.

## Request Parameters

For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**HostId.N**

The IDs of the Dedicated Hosts to release.

Type: Array of strings

Required: Yes

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**successful**

The IDs of the Dedicated Hosts that were successfully released.

Type: Array of strings

**unsuccessful**

The IDs of the Dedicated Hosts that could not be released, including an error
message.

Type: Array of [UnsuccessfulItem](api-unsuccessfulitem.md) objects

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This releases a Dedicated Host successfully.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=ReleaseHosts
&HostId=h-00548908djdsgfs
&AUTHPARAMS
```

#### Sample Response

```

<ReleaseHostsResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>d4904fd9-82c2-4ea5-adfe-a9cc3EXAMPLE</requestId>
    <unsuccessful/>
    <successful>
        <item>h-00548908djdsgfs</item>
    </successful>
</ReleaseHostsResponse>
```

### Example

This request is unsuccessful.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=ReleaseHosts
&HostId=h-00548908djdsgfs
&AUTHPARAMS
```

#### Sample Response

```

<ReleaseHostsResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>d4904fd9-82c2-4ea5-adfe-a9cc3EXAMPLE</requestId>
    <unsuccessful>
        <item>
       <error>
            <message>Dedicated host 'h-00548908djdsgfs' cannot be released as it is occupied</message>
            <code>Client.InvalidHost.Occupied</code<
       </error>
       <resourceId>h-00548908djdsgfs</resourceId>
        </item>
    </unsuccessful>
    <successful/>
</ReleaseHostsResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/releasehosts.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/releasehosts.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/releasehosts.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/releasehosts.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/releasehosts.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/releasehosts.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/releasehosts.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/releasehosts.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/releasehosts.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/releasehosts.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ReleaseAddress

ReleaseIpamPoolAllocation
