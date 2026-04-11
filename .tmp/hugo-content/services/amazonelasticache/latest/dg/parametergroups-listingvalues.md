# Listing an ElastiCache parameter group's values

You can list the parameters and their values for a parameter group using the ElastiCache
console, the AWS CLI, or the ElastiCache API.

## Listing an ElastiCache parameter group's values (Console)

The following procedure shows how to list the parameters and their values for a
parameter group using the ElastiCache console.

###### To list a parameter group's parameters and their values using the ElastiCache console

1. Sign in to the AWS Management Console and open the ElastiCache console at
    [https://console.aws.amazon.com/elasticache/](https://console.aws.amazon.com/elasticache).

2. To see a list of all available parameter groups, in the left hand
    navigation pane choose **Parameter Groups**.

3. Choose the parameter group for which you want to list the parameters and
    values by choosing the box to the left of the parameter group's
    name.

The parameters and their values will be listed at the bottom of the
    screen. Due to the number of parameters, you may have to scroll up and down
    to find the parameter you're interested in.

## Listing a parameter group's values (AWS CLI)

To list a parameter group's parameters and their values using the AWS CLI, use
the command `describe-cache-parameters`.

###### Example

The following sample code list all the Memcached parameters and their values for the
parameter group _myMem14_.

For Linux, macOS, or Unix:

```nohighlight

aws elasticache describe-cache-parameters \
    --cache-parameter-group-name myMem14
```

For Windows:

```nohighlight

aws elasticache describe-cache-parameters ^
    --cache-parameter-group-name myMem14
```

###### Example

The following sample code list all the parameters and their values for the
parameter group _myRedis28_.

For Linux, macOS, or Unix:

```nohighlight

aws elasticache describe-cache-parameters \
    --cache-parameter-group-name myRedis28
```

For Windows:

```nohighlight

aws elasticache describe-cache-parameters ^
    --cache-parameter-group-name myRed28
```

For more information, see [`describe-cache-parameters`](../../../cli/latest/reference/elasticache/describe-cache-parameters.md).

## Listing a parameter group's values (ElastiCache API)

To list a parameter group's parameters and their values using the ElastiCache API,
use the `DescribeCacheParameters` action.

###### Example

The following sample code list all the Memcached parameters for the parameter group
_myMem14_.

```nohighlight

https://elasticache.us-west-2.amazonaws.com/
   ?Action=DescribeCacheParameters
   &CacheParameterGroupName=myMem14
   &SignatureVersion=4
   &SignatureMethod=HmacSHA256
   &Timestamp=20150202T192317Z
   &Version=2015-02-02
   &X-Amz-Credential=<credential>
```

The response from this action will look something like this. This response has
been truncated.

```nohighlight

<DescribeCacheParametersResponse xmlns="http://elasticache.amazonaws.com/doc/2013-06-15/">
  <DescribeCacheParametersResult>
    <CacheClusterClassSpecificParameters>
      <CacheNodeTypeSpecificParameter>
        <DataType>integer</DataType>
        <Source>system</Source>
        <IsModifiable>false</IsModifiable>
        <Description>The maximum configurable amount of memory to use to store items, in megabytes.</Description>
        <CacheNodeTypeSpecificValues>
          <CacheNodeTypeSpecificValue>
            <Value>1000</Value>
            <CacheClusterClass>cache.c1.medium</CacheClusterClass>
          </CacheNodeTypeSpecificValue>
          <CacheNodeTypeSpecificValue>
            <Value>6000</Value>
            <CacheClusterClass>cache.c1.xlarge</CacheClusterClass>
          </CacheNodeTypeSpecificValue>
          <CacheNodeTypeSpecificValue>
            <Value>7100</Value>
            <CacheClusterClass>cache.m1.large</CacheClusterClass>
          </CacheNodeTypeSpecificValue>
          <CacheNodeTypeSpecificValue>
            <Value>1300</Value>
            <CacheClusterClass>cache.m1.small</CacheClusterClass>
          </CacheNodeTypeSpecificValue>

...output omitted...

    </CacheClusterClassSpecificParameters>
  </DescribeCacheParametersResult>
  <ResponseMetadata>
    <RequestId>6d355589-af49-11e0-97f9-279771c4477e</RequestId>
  </ResponseMetadata>
</DescribeCacheParametersResponse>
```

###### Example

The following sample code list all the parameters for the parameter group
_myRed28_.

```nohighlight

https://elasticache.us-west-2.amazonaws.com/
   ?Action=DescribeCacheParameters
   &CacheParameterGroupName=myRed28
   &SignatureVersion=4
   &SignatureMethod=HmacSHA256
   &Timestamp=20150202T192317Z
   &Version=2015-02-02
   &X-Amz-Credential=<credential>
```

The response from this action will look something like this. This response has
been truncated.

```nohighlight

<DescribeCacheParametersResponse xmlns="http://elasticache.amazonaws.com/doc/2013-06-15/">
  <DescribeCacheParametersResult>
    <CacheClusterClassSpecificParameters>
      <CacheNodeTypeSpecificParameter>
        <DataType>integer</DataType>
        <Source>system</Source>
        <IsModifiable>false</IsModifiable>
        <Description>The maximum configurable amount of memory to use to store items, in megabytes.</Description>
        <CacheNodeTypeSpecificValues>
          <CacheNodeTypeSpecificValue>
            <Value>1000</Value>
            <CacheClusterClass>cache.c1.medium</CacheClusterClass>
          </CacheNodeTypeSpecificValue>
          <CacheNodeTypeSpecificValue>
            <Value>6000</Value>
            <CacheClusterClass>cache.c1.xlarge</CacheClusterClass>
          </CacheNodeTypeSpecificValue>
          <CacheNodeTypeSpecificValue>
            <Value>7100</Value>
            <CacheClusterClass>cache.m1.large</CacheClusterClass>
          </CacheNodeTypeSpecificValue>
          <CacheNodeTypeSpecificValue>
            <Value>1300</Value>
            <CacheClusterClass>cache.m1.small</CacheClusterClass>
          </CacheNodeTypeSpecificValue>

...output omitted...

    </CacheClusterClassSpecificParameters>
  </DescribeCacheParametersResult>
  <ResponseMetadata>
    <RequestId>6d355589-af49-11e0-97f9-279771c4477e</RequestId>
  </ResponseMetadata>
</DescribeCacheParametersResponse>
```

For more information, see [`DescribeCacheParameters`](../../../../reference/amazonelasticache/latest/apireference/api-describecacheparameters.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Listing ElastiCache parameter groups by name

Modifying an ElastiCache parameter group

All content copied from https://docs.aws.amazon.com/.
