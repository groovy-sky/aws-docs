# DescribeDBClusters

Describes existing Amazon Aurora DB clusters and Multi-AZ DB clusters. This API supports pagination.

For more information on Amazon Aurora DB clusters, see
[What is Amazon Aurora?](../../../../services/amazonrds/latest/aurorauserguide/chap-auroraoverview.md) in the _Amazon Aurora User Guide_.

For more information on Multi-AZ DB clusters, see [Multi-AZ DB\
cluster deployments](../../../../services/amazonrds/latest/userguide/multi-az-db-clusters-concepts.md) in the _Amazon RDS User_
_Guide_.

This operation can also return information for Amazon Neptune DB instances and Amazon DocumentDB instances.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**DBClusterIdentifier**

The user-supplied DB cluster identifier or the Amazon Resource Name (ARN) of the DB cluster. If this parameter is specified,
information for only the specific DB cluster is returned. This parameter isn't case-sensitive.

Constraints:

- If supplied, must match an existing DB cluster identifier.

Type: String

Required: No

**Filters.Filter.N**

A filter that specifies one or more DB clusters to describe.

Supported Filters:

- `clone-group-id` \- Accepts clone group identifiers.
The results list only includes information about
the DB clusters associated with these clone groups.

- `db-cluster-id` \- Accepts DB cluster identifiers and DB
cluster Amazon Resource Names (ARNs). The results list only includes information about
the DB clusters identified by these ARNs.

- `db-cluster-resource-id` \- Accepts DB cluster resource identifiers.
The results list will only include information about the DB clusters identified
by these DB cluster resource identifiers.

- `domain` \- Accepts Active Directory directory IDs.
The results list only includes information about
the DB clusters associated with these domains.

- `engine` \- Accepts engine names.
The results list only includes information about
the DB clusters for these engines.

Type: Array of [Filter](api-filter.md) objects

Required: No

**IncludeShared**

Specifies whether the output includes information about clusters
shared from other AWS accounts.

Type: Boolean

Required: No

**Marker**

An optional pagination token provided by a previous
`DescribeDBClusters` request.
If this parameter is specified, the response includes
only records beyond the marker,
up to the value specified by `MaxRecords`.

Type: String

Required: No

**MaxRecords**

The maximum number of records to include in the response.
If more records exist than the specified `MaxRecords` value,
a pagination token called a marker is included in the response so you can retrieve the remaining results.

Default: 100

Constraints: Minimum 20, maximum 100

Type: Integer

Required: No

## Response Elements

The following elements are returned by the service.

**DBClusters.DBCluster.N**

Contains a list of DB clusters for the user.

Type: Array of [DBCluster](api-dbcluster.md) objects

**Marker**

A pagination token that can be used in a later `DescribeDBClusters` request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**DBClusterNotFoundFault**

`DBClusterIdentifier` doesn't refer to an existing DB cluster.

HTTP Status Code: 404

## Examples

### Describing an Aurora DB cluster

This example illustrates one usage of DescribeDBClusters.

#### Sample Request

```

https://rds.us-east-1.amazonaws.com/
    ?Action=DescribeDBClusters
    &MaxRecords=100
    &SignatureMethod=HmacSHA256
    &SignatureVersion=4
    &Version=2014-10-31
    &X-Amz-Algorithm=AWS4-HMAC-SHA256
    &X-Amz-Credential=AWS_ACCESS_KEY_ID_REDACTED/20230222/us-east-1/rds/aws4_request
    &X-Amz-Date=20230222T200807Z
    &X-Amz-SignedHeaders=content-type;host;user-agent;x-amz-content-sha256;x-amz-date
    &X-Amz-Signature=2d4f2b9e8abc31122b5546f94c0499bba47de813cb875f9b9c78e8e19c9afe1b
```

#### Sample Response

```

<DescribeDBClustersResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
  <DescribeDBClustersResult>
    <DBClusters>
      <DBCluster>
        <AssociatedRoles>
          <DBClusterRole>
            <RoleArn>arn:aws:iam::123456789012:role/sample-role</RoleArn>
            <Status>ACTIVE</Status>
          </DBClusterRole>
        </AssociatedRoles>
        <Engine>aurora-mysql</Engine>
        <Status>available</Status>
        <BackupRetentionPeriod>1</BackupRetentionPeriod>
        <DBSubnetGroup>my-subgroup</DBSubnetGroup>
        <EngineVersion>8.0.mysql_aurora.3.01.0</EngineVersion>
        <Endpoint>sample-cluster2.cluster-cbfvmgb0y5fy.us-east-1.rds.amazonaws.com</Endpoint>
        <DBClusterIdentifier>sample-cluster2</DBClusterIdentifier>
        <PreferredBackupWindow>04:45-05:15</PreferredBackupWindow>
        <PreferredMaintenanceWindow>sat:05:56-sat:06:26</PreferredMaintenanceWindow>
        <DBClusterMembers/>
        <AllocatedStorage>15</AllocatedStorage>
        <MasterUsername>awsuser</MasterUsername>
      </DBCluster>
      <DBCluster>
        <AssociatedRoles />
        <Engine>aurora-mysql</Engine>
        <Status>available</Status>
        <BackupRetentionPeriod>0</BackupRetentionPeriod>
        <DBSubnetGroup>my-subgroup</DBSubnetGroup>
        <EngineVersion>8.0.mysql_aurora.3.01.0</EngineVersion>
        <Endpoint>sample-cluster3.cluster-cefgqfx9y5fy.us-east-1.rds.amazonaws.com</Endpoint>
        <DBClusterIdentifier>sample-cluster3</DBClusterIdentifier>
        <PreferredBackupWindow>07:06-07:36</PreferredBackupWindow>
        <PreferredMaintenanceWindow>tue:10:18-tue:10:48</PreferredMaintenanceWindow>
        <DBClusterMembers>
          <DBClusterMember>
            <IsClusterWriter>true</IsClusterWriter>
            <DBInstanceIdentifier>sample-cluster3-master</DBInstanceIdentifier>
          </DBClusterMember>
          <DBClusterMember>
            <IsClusterWriter>false</IsClusterWriter>
            <DBInstanceIdentifier>sample-cluster3-read1</DBInstanceIdentifier>
          </DBClusterMember>
        </DBClusterMembers>
        <AllocatedStorage>15</AllocatedStorage>
        <MasterUsername>awsuser</MasterUsername>
      </DBCluster>
    </DBClusters>
  </DescribeDBClustersResult>
  <ResponseMetadata>
    <RequestId>d682b02c-1383-11b4-a6bb-172dfac7f170</RequestId>
  </ResponseMetadata>
</DescribeDBClustersResponse>
```

### Describing a Multi-AZ DB cluster

This example illustrates one usage of DescribeDBClusters.

#### Sample Request

```

https://rds.us-west-2.amazonaws.com/
    ?Action=DescribeDBClusters
    &DBClusterIdentifier=my-multi-az-cluster
    &SignatureMethod=HmacSHA256
    &SignatureVersion=4
    &Version=2014-10-31
    &X-Amz-Algorithm=AWS4-HMAC-SHA256
    &X-Amz-Credential=AWS_ACCESS_KEY_ID_REDACTED/20140722/us-west-2/rds/aws4_request
    &X-Amz-Date=20211026T203316Z
    &X-Amz-SignedHeaders=content-type;host;user-agent;x-amz-content-sha256;x-amz-date
    &X-Amz-Signature=2d4f2b9e8abc31122b5546f94c0499bba47de813cb875f9b9c78e8e19c9afe1b
```

#### Sample Response

```

<DescribeDBClustersResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
  <DescribeDBClustersResult>
    <DBClusters>
      <DBCluster>
        <CrossAccountClone>false</CrossAccountClone>
        <AllocatedStorage>100</AllocatedStorage>
        <DatabaseName>mydb</DatabaseName>
        <AssociatedRoles />
        <AvailabilityZones>
          <AvailabilityZone>us-west-2a</AvailabilityZone>
          <AvailabilityZone>us-west-2b</AvailabilityZone>
          <AvailabilityZone>us-west-2c</AvailabilityZone>
        </AvailabilityZones>
        <ReadReplicaIdentifiers />
        <Iops>1000</Iops>
        <PerformanceInsightsKMSKeyId>arn:aws:kms:us-west-2:123456789012:key/123EXAMPLE-abcd-4567-efgEXAMPLE</PerformanceInsightsKMSKeyId>
        <PerformanceInsightsRetentionPeriod>7</PerformanceInsightsRetentionPeriod>
        <EngineVersion>8.0.26</EngineVersion>
        <MasterUsername>admin</MasterUsername>
        <DBClusterMembers>
          <DBClusterMember>
            <DBInstanceIdentifier>my-multi-az-cluster-instance-3</DBInstanceIdentifier>
            <DBClusterParameterGroupStatus>in-sync</DBClusterParameterGroupStatus>
            <PromotionTier>1</PromotionTier>
            <IsClusterWriter>false</IsClusterWriter>
          </DBClusterMember>
          <DBClusterMember>
            <DBInstanceIdentifier>my-multi-az-cluster-instance-2</DBInstanceIdentifier>
            <DBClusterParameterGroupStatus>in-sync</DBClusterParameterGroupStatus>
            <PromotionTier>1</PromotionTier>
            <IsClusterWriter>false</IsClusterWriter>
          </DBClusterMember>
          <DBClusterMember>
            <DBInstanceIdentifier>my-multi-az-cluster-instance-1</DBInstanceIdentifier>
            <DBClusterParameterGroupStatus>in-sync</DBClusterParameterGroupStatus>
            <PromotionTier>1</PromotionTier>
            <IsClusterWriter>true</IsClusterWriter>
          </DBClusterMember>
        </DBClusterMembers>
        <DBActivityStreamStatus>stopped</DBActivityStreamStatus>
        <HttpEndpointEnabled>false</HttpEndpointEnabled>
        <Port>3306</Port>
        <MonitoringInterval>30</MonitoringInterval>
        <BackupRetentionPeriod>2</BackupRetentionPeriod>
        <KmsKeyId>arn:aws:kms:us-west-2:123456789012:key/123EXAMPLE-abcd-4567-efgEXAMPLE</KmsKeyId>
        <DBClusterIdentifier>my-multi-az-cluster</DBClusterIdentifier>
        <DbClusterResourceId>cluster-TSW4QJNKY3P2DNDRR523BDGEIU</DbClusterResourceId>
        <Status>creating</Status>
        <PreferredBackupWindow>11:34-12:04</PreferredBackupWindow>
        <DeletionProtection>false</DeletionProtection>
        <Endpoint>my-multi-az-cluster.cluster-123456789012.us-west-2.rds.amazonaws.com</Endpoint>
        <EngineMode>provisioned</EngineMode>
        <Engine>mysql</Engine>
        <ReaderEndpoint>my-multi-az-cluster.cluster-ro-123456789012.us-west-2.rds.amazonaws.com</ReaderEndpoint>
        <PubliclyAccessible>true</PubliclyAccessible>
        <IAMDatabaseAuthenticationEnabled>false</IAMDatabaseAuthenticationEnabled>
        <ClusterCreateTime>2021-10-26T20:31:54.943Z</ClusterCreateTime>
        <ActivityStreamStatus>stopped</ActivityStreamStatus>
        <PerformanceInsightsEnabled>true</PerformanceInsightsEnabled>
        <MultiAZ>true</MultiAZ>
        <DomainMemberships />
        <MonitoringRoleArn>arn:aws:iam::123456789012:role/enhance-monitoring-role</MonitoringRoleArn>
        <StorageEncrypted>true</StorageEncrypted>
        <DBSubnetGroup>mysubnet1</DBSubnetGroup>
        <VpcSecurityGroups>
          <VpcSecurityGroupMembership>
            <VpcSecurityGroupId>sg-6921cc28</VpcSecurityGroupId>
            <Status>active</Status>
          </VpcSecurityGroupMembership>
        </VpcSecurityGroups>
        <TagList />
        <HostedZoneId>Z3GZ3VYA3PGHTQ</HostedZoneId>
        <PreferredMaintenanceWindow>sat:07:05-sat:07:35</PreferredMaintenanceWindow>
        <DBClusterParameterGroup>my-cluster-param-1</DBClusterParameterGroup>
        <StorageType>io1</StorageType>
        <DBClusterInstanceClass>db.r6gd.large</DBClusterInstanceClass>
        <CopyTagsToSnapshot>false</CopyTagsToSnapshot>
        <AutoMinorVersionUpgrade>true</AutoMinorVersionUpgrade>
        <DBClusterArn>arn:aws:rds:us-west-2:123456789012:cluster:my-multi-az-cluster</DBClusterArn>
      </DBCluster>
    </DBClusters>
  </DescribeDBClustersResult>
  <ResponseMetadata>
    <RequestId>ae8b2342-55d7-4cf0-b7b3-f24e681ce7b9</RequestId>
  </ResponseMetadata>
</DescribeDBClustersResponse>
```

### Describing a DB cluster with Internet Access Gateway, and disabled VPC Networking

This example illustrates one usage of DescribeDBClusters.

#### Sample Request

```

                    https://rds.us-east-1.amazonaws.com/
    ?Action=DescribeDBClusters
    &DBClusterIdentifier=testvalidation-cluster-3
    &SignatureMethod=HmacSHA256
    &SignatureVersion=4
    &Version=2014-10-31
    &X-Amz-Algorithm=AWS4-HMAC-SHA256
    &X-Amz-Credential=AWS_ACCESS_KEY_ID_REDACTED/20260309/us-east-1/rds/aws4_request
    &X-Amz-Date=20260309T205322Z
    &X-Amz-SignedHeaders=content-type;host;user-agent;x-amz-content-sha256;x-amz-date
    &X-Amz-Signature=2d4f2b9e8abc31122b5546f94c0499bba47de813cb875f9b9c78e8e19c9afe1b
```

#### Sample Response

```

                    <DescribeDBClustersResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
  <DescribeDBClustersResult>
    <DBClusters>
      <DBCluster>
        <AllocatedStorage>1</AllocatedStorage>
        <AvailabilityZones>
          <AvailabilityZone>us-east-1d</AvailabilityZone>
          <AvailabilityZone>us-east-1a</AvailabilityZone>
          <AvailabilityZone>us-east-1c</AvailabilityZone>
        </AvailabilityZones>
        <BackupRetentionPeriod>1</BackupRetentionPeriod>
        <DBClusterIdentifier>testvalidation-cluster-3</DBClusterIdentifier>
        <DBClusterParameterGroup>default.aurora-postgresql17</DBClusterParameterGroup>
        <Status>available</Status>
        <MultiAZ>false</MultiAZ>
        <Engine>aurora-postgresql</Engine>
        <EngineVersion>17.4</EngineVersion>
        <Port>5432</Port>
        <MasterUsername>postgres</MasterUsername>
        <PreferredBackupWindow>07:13-07:43</PreferredBackupWindow>
        <PreferredMaintenanceWindow>mon:07:55-mon:08:25</PreferredMaintenanceWindow>
        <DBClusterMembers/>
        <VpcSecurityGroups/>
        <StorageEncrypted>false</StorageEncrypted>
        <DbClusterResourceId>cluster-AHX35HFI2YV26F3XVXVVO3MEHU</DbClusterResourceId>
        <DBClusterArn>arn:aws:rds:us-east-1:634686195229:cluster:testvalidation-cluster-3</DBClusterArn>
        <AssociatedRoles/>
        <IAMDatabaseAuthenticationEnabled>true</IAMDatabaseAuthenticationEnabled>
        <ClusterCreateTime>2026-03-09T20:53:24.054Z</ClusterCreateTime>
        <EngineMode>provisioned</EngineMode>
        <AutoMinorVersionUpgrade>true</AutoMinorVersionUpgrade>
        <DeletionProtection>false</DeletionProtection>
        <HttpEndpointEnabled>false</HttpEndpointEnabled>
        <DBActivityStreamStatus>stopped</DBActivityStreamStatus>
        <CopyTagsToSnapshot>false</CopyTagsToSnapshot>
        <CrossAccountClone>false</CrossAccountClone>
        <ServerlessV2ScalingConfiguration>
          <MinCapacity>1.0</MinCapacity>
          <MaxCapacity>128.0</MaxCapacity>
        </ServerlessV2ScalingConfiguration>
        <VPCNetworkingEnabled>false</VPCNetworkingEnabled>
        <InternetAccessGatewayEnabled>true</InternetAccessGatewayEnabled>
      </DBCluster>
    </DBClusters>
  </DescribeDBClustersResult>
  <ResponseMetadata>
    <RequestId>f193c741-2294-22c5-b7cc-283efbd8g281</RequestId>
  </ResponseMetadata>
</DescribeDBClustersResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/describedbclusters.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/describedbclusters.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/describedbclusters.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/describedbclusters.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/describedbclusters.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/describedbclusters.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/describedbclusters.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/describedbclusters.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/describedbclusters.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/describedbclusters.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeDBClusterParameters

DescribeDBClusterSnapshotAttributes

All content copied from https://docs.aws.amazon.com/.
