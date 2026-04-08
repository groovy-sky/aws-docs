# Removing the EFS\_INTEGRATION option

The steps for removing the `EFS_INTEGRATION` option depend on whether you're
removing the option from multiple DB instances or a single instance.

Number of DB instancesActionRelated informationMultipleRemove the `EFS_INTEGRATION` option from the option group to
which the DB instances belong. This change affects all instances that use the
option group.[Removing an option from an option group](user-workingwithoptiongroups.md#USER_WorkingWithOptionGroups.RemoveOption)SingleModify the DB instance and specify a different option group that doesn't
include the `EFS_INTEGRATION` option. You can specify the default
(empty) option group or a different custom option group.[Modifying an Amazon RDS DB instance](overview-dbinstance-modifying.md)

After you remove the `EFS_INTEGRATION` option, you can optionally delete the
EFS file system that was connected to your DB instances.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Transferring files

Troubleshooting
Amazon EFS

All content copied from https://docs.aws.amazon.com/.
