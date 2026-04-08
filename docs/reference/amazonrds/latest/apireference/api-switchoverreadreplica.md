# SwitchoverReadReplica

Switches over an Oracle standby database in an Oracle Data Guard environment, making it the new
primary database. Issue this command in the Region that hosts the current standby database.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**DBInstanceIdentifier**

The DB instance identifier of the current standby database. This value is stored as a lowercase string.

Constraints:

- Must match the identiﬁer of an existing Oracle read replica DB instance.

Type: String

Required: Yes

## Response Elements

The following element is returned by the service.

**DBInstance**

Contains the details of an Amazon RDS DB instance.

This data type is used as a response element in the operations `CreateDBInstance`,
`CreateDBInstanceReadReplica`, `DeleteDBInstance`, `DescribeDBInstances`,
`ModifyDBInstance`, `PromoteReadReplica`, `RebootDBInstance`,
`RestoreDBInstanceFromDBSnapshot`, `RestoreDBInstanceFromS3`, `RestoreDBInstanceToPointInTime`,
`StartDBInstance`, and `StopDBInstance`.

Type: [DBInstance](api-dbinstance.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**DBInstanceNotFound**

`DBInstanceIdentifier` doesn't refer to an existing DB instance.

HTTP Status Code: 404

**InvalidDBInstanceState**

The DB instance isn't in a valid state.

HTTP Status Code: 400

## Examples

### Example

The following example shows one use of SwitchoverReadReplica.

#### Sample Request

```

https://rds.us-east-1.amazonaws.com/
?Action=SwitchoverReadReplica
&DBInstanceIdentifier=new-primary
&Version=2014-10-31
&Signature=12345678caef670d84c14ffba62e107842557f934f1e68e5d38a2d219ae70527
```

#### Sample Response

```

<SwitchoverReadReplicaResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
    <SwitchoverReadReplicaResult>
        <DBInstance>
            <AllocatedStorage>20</AllocatedStorage>
            <ReadReplicaSourceDBInstanceIdentifier>bystanders-old-primary</ReadReplicaSourceDBInstanceIdentifier>
            <AssociatedRoles/>
            <DBParameterGroups>
                <DBParameterGroup>
                    <DBParameterGroupName>default.oracle-ee-19</DBParameterGroupName>
                    <ParameterApplyStatus>in-sync</ParameterApplyStatus>
                </DBParameterGroup>
            </DBParameterGroups>
            <AvailabilityZone>us-west-2c</AvailabilityZone>
            <DBSecurityGroups/>
            <StatusInfos>
                <DBInstanceStatusInfo>
                    <Normal>true</Normal>
                    <StatusType>read replication</StatusType>
                    <Status>replicating</Status>
                </DBInstanceStatusInfo>
            </StatusInfos>
            <EngineVersion>19.0.0.0.ru-2021-10.rur-2021-10.r1</EngineVersion>
            <MasterUsername>masteruser</MasterUsername>
            <InstanceCreateTime>2022-01-09T11:55:29.005Z</InstanceCreateTime>
            <DBInstanceClass>db.m4.xlarge</DBInstanceClass>
            <StorageThroughput>0</StorageThroughput>
            <HttpEndpointEnabled>false</HttpEndpointEnabled>
            <ReadReplicaDBInstanceIdentifiers/>
            <CustomerOwnedIpEnabled>false</CustomerOwnedIpEnabled>
            <MonitoringInterval>0</MonitoringInterval>
            <DBInstanceStatus>available</DBInstanceStatus>
            <BackupRetentionPeriod>1</BackupRetentionPeriod>
            <OptionGroupMemberships>
                <OptionGroupMembership>
                    <OptionGroupName>default:oracle-ee-19</OptionGroupName>
                    <Status>in-sync</Status>
                </OptionGroupMembership>
            </OptionGroupMemberships>
            <BackupTarget>region</BackupTarget>
            <CACertificateIdentifier>rds-ca-2019</CACertificateIdentifier>
            <DbInstancePort>0</DbInstancePort>
            <DbiResourceId>db-ABCDEFG12H3I4J5KLMNOPQR6ST</DbiResourceId>
            <PreferredBackupWindow>11:11-11:11</PreferredBackupWindow>
            <DeletionProtection>false</DeletionProtection>
            <DBInstanceIdentifier>new-primary</DBInstanceIdentifier>
            <DBInstanceArn>arn:aws:rds:us-west-2:123456789012:db:new-primary</DBInstanceArn>
            <Endpoint>
                <HostedZoneId>ABCD7F8REH8UF3</HostedZoneId>
                <Address>new-primary.abcdefgh0ijk.us-west-2.rds.amazonaws.com</Address>
                <Port>1521</Port>
            </Endpoint>
            <Engine>oracle-ee</Engine>
            <PubliclyAccessible>true</PubliclyAccessible>
            <IAMDatabaseAuthenticationEnabled>false</IAMDatabaseAuthenticationEnabled>
            <NetworkType>IPV4</NetworkType>
            <PerformanceInsightsEnabled>false</PerformanceInsightsEnabled>
            <ReplicaMode>open-read-only</ReplicaMode>
            <DBName>ORCL</DBName>
            <MultiAZ>false</MultiAZ>
            <DomainMemberships/>
            <CharacterSetName>AL32UTF8</CharacterSetName>
            <StorageEncrypted>false</StorageEncrypted>
            <DBSubnetGroup>
                <VpcId>vpc-2f206b57</VpcId>
                <Subnets>
                    <Subnet>
                        <SubnetIdentifier>subnet-ac26e0e6</SubnetIdentifier>
                        <SubnetStatus>Active</SubnetStatus>
                        <SubnetOutpost/>
                        <SubnetAvailabilityZone>
                            <Name>us-west-2a</Name>
                        </SubnetAvailabilityZone>
                    </Subnet>
                    <Subnet>
                        <SubnetIdentifier>subnet-1a2bcde3</SubnetIdentifier>
                        <SubnetStatus>Active</SubnetStatus>
                        <SubnetOutpost/>
                        <SubnetAvailabilityZone>
                            <Name>us-west-2b</Name>
                        </SubnetAvailabilityZone>
                    </Subnet>
                    <Subnet>
                        <SubnetIdentifier>subnet-a1b2c3de</SubnetIdentifier>
                        <SubnetStatus>Active</SubnetStatus>
                        <SubnetOutpost/>
                        <SubnetAvailabilityZone><Name>us-west-2d</Name>
                        </SubnetAvailabilityZone>
                    </Subnet>
                    <Subnet>
                        <SubnetIdentifier>subnet-a12345b6</SubnetIdentifier>
                        <SubnetStatus>Active</SubnetStatus>
                        <SubnetOutpost/>
                        <SubnetAvailabilityZone>
                            <Name>us-west-2c</Name>
                        </SubnetAvailabilityZone>
                    </Subnet>
                </Subnets>
                <SubnetGroupStatus>Complete</SubnetGroupStatus>
                <DBSubnetGroupDescription>default</DBSubnetGroupDescription>
                <DBSubnetGroupName>default</DBSubnetGroupName>
            </DBSubnetGroup>
            <VpcSecurityGroups>
                <VpcSecurityGroupMembership>
                    <VpcSecurityGroupId>ab-12c3d45e</VpcSecurityGroupId>
                    <Status>active</Status>
                </VpcSecurityGroupMembership>
            </VpcSecurityGroups>
            <TagList/>
            <NcharCharacterSetName>AL16UTF16</NcharCharacterSetName>
            <LicenseModel>bring-your-own-license</LicenseModel>
            <PendingModifiedValues/>
            <PreferredMaintenanceWindow>tue:07:56-tue:08:26</PreferredMaintenanceWindow>
            <StorageType>gp2</StorageType>
            <AutoMinorVersionUpgrade>true</AutoMinorVersionUpgrade>
            <CopyTagsToSnapshot>false</CopyTagsToSnapshot>
        </DBInstance>
    </SwitchoverReadReplicaResult>
    <ResponseMetadata>
        <RequestId>abcd12ef-34g5-41d6-aed9-b6366d786923</RequestId>
    </ResponseMetadata>
</SwitchoverReadReplicaResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/switchoverreadreplica.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/switchoverreadreplica.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/switchoverreadreplica.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/switchoverreadreplica.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/switchoverreadreplica.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/switchoverreadreplica.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/switchoverreadreplica.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/switchoverreadreplica.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/switchoverreadreplica.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/switchoverreadreplica.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SwitchoverGlobalCluster

Data Types

All content copied from https://docs.aws.amazon.com/.
