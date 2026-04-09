# Restore testing inferred metadata

Restoring a recovery point requires restore metadata. To perform restore tests, AWS Backup
automatically infers metadata that is likely to result in a successful restore. The command
`get-restore-testing-inferred-metadata` can be used to preview what AWS Backup will
infer. The command `get-restore-job-metadata` returns the set of metadata
inferred by AWS Backup. Note that for some resource types (Amazon FSx), AWS Backup is not able to infer a
complete set of metadata.

_Inferred restore metadata_ is determined during the restore testing
process. You can override certain restore metadata keys by including the parameter
`RestoreMetadataOverrides` in the body of `RestoreTestingSelection`.
Some metadata overrides are not available in the AWS Backup console.

Each supported resource has both inferred restore metadata keys and values, and
overridable restore metadata keys. Only `RestoreMetadataOverrides` key value
pairs or nested key value pairs marked with `required for successful
          restore` are necessary to include; the others are optional.
Note that key values are not
case sensitive.

###### Important

AWS Backup can infer that
a resource should be restored to the default setting, such as an Amazon EC2
instance or Amazon RDS cluster restored to the default VPC. However, if the default is not present,
for example the default VPC or subnet has been deleted and no metadata override has been
input, the restore will not be successful.

Resource typeInferred restore metadata keys and valuesOverridable metadata

**DynamoDB**

`deletionProtection`, where value is set to
`false`

`encryptionType` is set to `Default`

`targetTableName`, where value is set to random value starting with
`awsbackup-restore-test-`

`encryptionType`

`kmsMasterKeyArn`

**Amazon EBS**

`availabilityZone`, whose value is set to a random availability
zone

`encrypted`, whose value is set to `true`

`availabilityZone`

`iops`

`kmsKeyId`

`throughput`

`volumesize`

`volumetype`

**Amazon EC2**

`disableApiTermination` value is set to `false`

`instanceType` value is set to the instanceType of the recovery
point being restored

`requiredImdsV2` value is set to `true`

`iamInstanceProfileName` (null or `UseBackedUpValue`)

`instanceType`

`requireImdsV2`

`securityGroupIds`

`subnetId`

**Amazon EFS**

`encrypted` value is set to `true`

`file-system-id` value is set to the file system ID of the recovery
point being restored

`kmsKeyId value` is set to
`alias/aws/elasticfilesystem`

`newFileSystem` value is set to `true`

`performanceMode` value is set to
`generalPurpose`

`kmsKeyId`

`performanceMode`

**Amazon FSx for Lustre**

`lustreConfiguration` has nested keys. One nested key is
`automaticBackupRetentionDays`, the value of which is set to
`0`

`kmsKeyId`

`lustreConfiguration` has nested key
`logConfiguration`

`securityGroupIds`

`subnetIds`, `required for successful
                    restore`

**Amazon FSx for NetApp ONTAP**

`name` is set to a random value starting with
`awsbackup_restore_test_`

`ontapConfiguration` has nested keys, including:

- `junctionPath` where `/name` is the name of the
volume being restored

- `sizeInMegabytes`, the value of which is set to the size in
megabytes of the recovery point being restored

- `snapshotPolicy` where the value is set to
`none`

`ontapConfiguration` has specific overrideable nested keys,
including:

- `junctionPath`

- `ontapVolumeType`

- `securityStyle`

- `sizeInMegabytes`

- `storageEfficiencyEnabled`

- `storageVirtualMachineId`, `required for successful
                          restore`

- `tieringPolicy`

**Amazon FSx for OpenZFS**

`openZfzConfiguration`, which has nested keys, including:

- `automaticBackupRetentionDays` with value set to
`0`

- `deploymentType` with value set to the deployment type of the
recovery point being restored

- `throughputCapacity`, whose value is based on the
`deploymentType`. If `deploymentType` is
`SINGLE_AZ_1`, the value is set to `64`; if the
`deploymentType` is `SINGLE_AZ_2 or MULTI_AZ_1`, the
value is set to `160`

`kmsKeyId`

`openZfsConfiguration` has specific overridable nested keys,
including:

- `deploymentType`

- `throughputCapacity`

- `diskiopsConfiguration`

`securityGroupIds`

`subnetIds`

**Amazon FSx for Windows File Server**

`windowsConfiguration`, which has nested keys including:

- `automaticBackupRetentionDays` with value set to
`0`

- `deploymentType` with value set to the deployment type of the
recovery point being restored

- `throughputCapacity` with value set to `8`

`kmsKeyId`

`securityGroupIds`

`subnetIds` `required for successful restore`

`windowsConfiguration`, with specific overridable nested
keys

- `throughputCapacity`

- `activeDirectoryId` `required for successful restore`

- `preferredSubnetId`

**Amazon RDS, Aurora, Amazon DocumentDB, Amazon Neptune**
**clusters**

`availabilityZones` with value set to a list of up to three random
availability zones

`dbClusterIdentifier` with a random value starting with
`awsbackup-restore-test`

`engine` with value set to the engine of the recovery point being
restored

`availabilityZones`

`databaseName`

`dbClusterParameterGroupName`

`dbSubnetGroupName`

`enableCloudwatchLogsExports`

`enableIamDatabaseAuthentication`

`engine`

`engineMode`

`engineVersion`

`kmskeyId`

`port`

`optionGroupName`

`scalingConfiguration`

`vpcSecurityGroupIds`

**Amazon RDS instances**

`dbInstanceIdentifier` with a random value starting with
`awsbackup-restore-test-`

`deletionProtection` with value set to `false`

`multiAz` with value set to `false`

`publiclyAccessible` with value set to false

`allocatedStorage`

`availabilityZones`

`dbInstanceClass`

`dbName`

`dbParameterGroupName`

`dbSubnetGroupName`

`domain`

`domainIamRoleName`

`enableCloudwatchLogsExports`

`enableIamDatabaseAuthentication`

`iops`

`licensemodel`

`multiAz`

`optionGroupName`

`port`

`processorFeatures`

`publiclyAccessible`

`storageType`

`vpcSecurityGroupIds`

**Amazon Simple Storage Service (Amazon S3)**

`destinationBucketName` with a random value starting with
`awsbackup-restore-test-`

`encrypted ` with value set to `true`

`encryptionType` with value set to `SSE-S3`

`newBucket` with value set to `true`

`encryptionType`

`kmsKey`

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Restore testing

Restore testing validation

All content copied from https://docs.aws.amazon.com/.
