# Troubleshooting SSAS issues

You might encounter the following issues when using SSAS.

IssueTypeTroubleshooting suggestions**`Unable to configure the SSAS option. The requested SSAS mode is**
**new_mode, but the current DB instance has**
**number**
**current_mode databases. Delete the existing databases before switching to**
**new_mode mode. To regain access to**
**current_mode mode for database deletion, either update the current DB**
**option group, or attach a new option group with %s as the MODE option setting value for the SSAS**
**option.`**RDS eventYou can't change the SSAS mode if you still have SSAS databases that use the current mode. Delete
the SSAS databases, then try again.**`Unable to remove the SSAS option because there are number existing**
**mode databases. The SSAS option can't be removed until all SSAS databases**
**are deleted. Add the SSAS option again, delete all SSAS databases, and try again.`**RDS eventYou can't turn off SSAS if you still have SSAS databases.
Delete the SSAS databases, then try again.**`The SSAS option isn't enabled or is in the process of being enabled. Try again**
**later.`**RDS stored procedureYou can't run SSAS stored procedures when the option is turned off, or when it's being turned
on.**`The SSAS option is configured incorrectly. Make sure that the option group membership status**
**is "in-sync", and review the RDS event logs for relevant SSAS configuration error messages. Following**
**these investigations, try again. If errors continue to occur, contact AWS Support.`**RDS stored procedure

You can't run SSAS stored procedures when your option
group membership isn't in the `in-sync` status. This
puts the SSAS option in an incorrect configuration state.

If your option group membership status changes to `failed` due to SSAS option modification,
there are two possible reasons:

1. The SSAS option was removed without the SSAS databases
    being deleted.

2. The SSAS mode was updated from Tabular to
    Multidimensional, or from Multidimenisonal to Tabular,
    without the existing SSAS databases being deleted.

Reconfigure the SSAS option, because RDS allows only one SSAS mode at a time, and doesn't support
SSAS option removal with SSAS databases present.

Check the RDS event logs for configuration errors for your SSAS instance, and resolve the issues
accordingly.

**`Deployment failed. The change can only be deployed on a server running in**
**deployment_file_mode mode. The current server mode is**
**current_mode.`**RDS stored procedure

You can't deploy a Tabular database to a
Multidimensional server, or a Multidimensional database to a Tabular
server.

Make sure that you're using files with the correct mode, and verify that the `MODE` option
setting is set to the appropriate value.

**`The restore failed. The backup file can only be restored on a server running in**
**restore_file_mode mode. The current server mode is**
**current_mode.`**RDS stored procedure

You can't restore a Tabular database to a
Multidimensional server, or a Multidimensional database to a Tabular
server.

Make sure that you're using files with the correct mode, and verify that the `MODE` option
setting is set to the appropriate value.

**`The restore failed. The backup file and the RDS DB instance versions are**
**incompatible.`**RDS stored procedure

You can't restore an SSAS database with a version incompatible to the SQL Server instance
version.

For more information, see [Compatibility levels for tabular models](https://docs.microsoft.com/en-us/analysis-services/tabular-models/compatibility-level-for-tabular-models-in-analysis-services) and [Compatibility level of a multidimensional database](https://docs.microsoft.com/en-us/analysis-services/multidimensional-models/compatibility-level-of-a-multidimensional-database-analysis-services) in the Microsoft documentation.

**`The restore failed. The backup file specified in the restore operation is damaged or is**
**not an SSAS backup file. Make sure that @rds_file_path is correctly formatted.`**RDS stored procedure

You can't restore an SSAS database with a damaged file.

Make sure that the file isn't damaged or corrupted.

This error can also be raised when `@rds_file_path` isn't correctly formatted (for
example, it has double backslashes as in `D:\S3\\incorrect_format.abf`).

**``The restore failed. The restored database name can't contain any reserved words or**
**invalid characters: . , ; ' ` : / \\ * | ? \" & % $ ! + = ( ) [ ] { } < >, or be longer than 100**
**characters.``**RDS stored procedure

The restored database name can't contain any reserved
words or characters that aren't valid, or be longer than 100
characters.

For SSAS object naming conventions, see [Object naming rules](https://docs.microsoft.com/en-us/analysis-services/multidimensional-models/olap-physical/object-naming-rules-analysis-services) in the Microsoft documentation.

**`An invalid role name was provided. The role name can't contain any reserved**
**strings.`**RDS stored procedure

The role name can't contain any reserved strings.

For SSAS object naming conventions, see [Object naming rules](https://docs.microsoft.com/en-us/analysis-services/multidimensional-models/olap-physical/object-naming-rules-analysis-services) in the Microsoft documentation.

**``An invalid role name was provided. The role name can't contain any of the following**
**reserved characters: . , ; ' ` : / \\ * | ? \" & % $ ! + = ( ) [ ] { } < >``**RDS stored procedure

The role name can't contain any reserved characters.

For SSAS object naming conventions, see [Object naming rules](https://docs.microsoft.com/en-us/analysis-services/multidimensional-models/olap-physical/object-naming-rules-analysis-services) in the Microsoft documentation.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Turning off SSAS

SQL Server Integration Services

All content copied from https://docs.aws.amazon.com/.
