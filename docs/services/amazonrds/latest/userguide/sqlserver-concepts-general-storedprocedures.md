# Functions and stored procedures for Amazon RDS for Microsoft SQL Server

Following, you can find a list of the Amazon RDS functions and stored procedures that help
automate SQL Server tasks.

Task typeProcedure or functionWhere it's usedAdministrative tasks

`rds_drop_database`

[Dropping a database in an Amazon RDS for Microsoft SQL Server DB instance](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.SQLServer.CommonDBATasks.DropMirrorDB.html)

`rds_failover_time`

[Determining the last failover time for Amazon RDS for SQL Server](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.SQLServer.CommonDBATasks.LastFailover.html)`rds_manage_view_db_permission`[Deny or allow viewing database names for Amazon RDS for SQL Server](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.SQLServer.CommonDBATasks.ManageView.html)

`rds_modify_db_name`

[Renaming a Amazon RDS for Microsoft SQL Server database in a Multi-AZ deployment](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.SQLServer.CommonDBATasks.RenamingDB.html)

`rds_read_error_log`

[Viewing error and agent logs](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.SQLServer.CommonDBATasks.Logs.html#Appendix.SQLServer.CommonDBATasks.Logs.SP)

`rds_set_configuration`

This operation is used to set various DB instance
configurations:

- [Change data capture for Multi-AZ instances](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.SQLServer.CommonDBATasks.CDC.html#Appendix.SQLServer.CommonDBATasks.CDC.Multi-AZ)

- [Setting the retention period for trace and dump files](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.SQLServer.CommonDBATasks.TraceFiles.html#Appendix.SQLServer.CommonDBATasks.TraceFiles.PurgeTraceFiles)

- [Compressing backup files](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/SQLServer.Procedural.Importing.Native.Compression.html)

`rds_set_database_online`

[Transitioning a Amazon RDS for SQL Server database from OFFLINE to ONLINE](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.SQLServer.CommonDBATasks.TransitionOnline.html)

`rds_set_system_database_sync_objects`

`rds_fn_get_system_database_sync_objects`

`rds_fn_server_object_last_sync_time`

[Turning on SQL Server Agent job replication](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.SQLServer.CommonDBATasks.Agent.html#SQLServerAgent.Replicate)

`rds_show_configuration`

To see the values that are set using
`rds_set_configuration`, see these topics:

- [Change data capture for Multi-AZ instances](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.SQLServer.CommonDBATasks.CDC.html#Appendix.SQLServer.CommonDBATasks.CDC.Multi-AZ)

- [Setting the retention period for trace and dump files](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.SQLServer.CommonDBATasks.TraceFiles.html#Appendix.SQLServer.CommonDBATasks.TraceFiles.PurgeTraceFiles)

`rds_shrink_tempdbfile`

[Shrinking the tempdb database](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/SQLServer.TempDB.Shrinking.html)Change data capture (CDC)

`rds_cdc_disable_db`

[Disabling CDC](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.SQLServer.CommonDBATasks.CDC.html)

`rds_cdc_enable_db`

[Enabling CDC](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.SQLServer.CommonDBATasks.CDC.html)Database Mail

`rds_fn_sysmail_allitems`

[Viewing messages, logs, and attachments](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/SQLServer.DBMail.View.html)

`rds_fn_sysmail_event_log`

[Viewing messages, logs, and attachments](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/SQLServer.DBMail.View.html)

`rds_fn_sysmail_mailattachments`

[Viewing messages, logs, and attachments](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/SQLServer.DBMail.View.html)

`rds_sysmail_control`

This operation is used in starting and stopping the mail
queue:

- [Starting the mail queue](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/SQLServer.DBMail.StartStop.html#SQLServer.DBMail.Start)

- [Stopping the mail queue](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/SQLServer.DBMail.StartStop.html#SQLServer.DBMail.Stop)

`rds_sysmail_delete_mailitems_sp`

[Deleting messages](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/SQLServer.DBMail.Delete.html)Native backup and restore

`rds_backup_database`

[Backing up a database](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/SQLServer.Procedural.Importing.Native.Using.html#SQLServer.Procedural.Importing.Native.Using.Backup)

`rds_cancel_task`

[Canceling a task](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/SQLServer.Procedural.Importing.Native.Using.html#SQLServer.Procedural.Importing.Native.Using.Cancel)

`rds_finish_restore`

[Finishing a database restore](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/SQLServer.Procedural.Importing.Native.Using.html#SQLServer.Procedural.Importing.Native.Finish.Restore)

`rds_restore_database`

[Restoring a database](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/SQLServer.Procedural.Importing.Native.Using.html#SQLServer.Procedural.Importing.Native.Using.Restore)

`rds_restore_log`

[Restoring a log](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/SQLServer.Procedural.Importing.Native.Using.html#SQLServer.Procedural.Importing.Native.Restore.Log)Amazon S3 file transfer

`rds_delete_from_filesystem`

[Deleting files on the RDS DB instance](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.SQLServer.Options.S3-integration.using.deleting-files.html)

`rds_download_from_s3`

[Downloading files from an Amazon S3 bucket to a SQL Server DB instance](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.SQLServer.Options.S3-integration.using.html#Appendix.SQLServer.Options.S3-integration.using.download)

`rds_gather_file_details`

[Listing files on the RDS DB instance](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.SQLServer.Options.S3-integration.using.listing-files.html)

`rds_upload_to_s3`

[Uploading files from a SQL Server DB instance to an Amazon S3 bucket](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.SQLServer.Options.S3-integration.using.html#Appendix.SQLServer.Options.S3-integration.using.upload)Microsoft Distributed Transaction Coordinator (MSDTC)

`rds_msdtc_transaction_tracing`

[Using transaction tracing](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.SQLServer.Options.MSDTC.html#MSDTC.Tracing)SQL Server Audit

`rds_fn_get_audit_file`

[Viewing audit logs](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.SQLServer.Options.Audit.AuditRecords.html)Transparent Data Encryption

`rds_backup_tde_certificate`

`rds_drop_tde_certificate`

`rds_restore_tde_certificate`

`rds_fn_list_user_tde_certificates`

[Support for Transparent Data Encryption in SQL Server](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.SQLServer.Options.TDE.html)Microsoft Business Intelligence (MSBI)

`rds_msbi_task`

This operation is used with SQL Server Analysis Services
(SSAS):

- [Deploying SSAS projects on Amazon RDS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/SSAS.Deploy.html)

- [Adding a domain user as a database administrator](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/SSAS.Use.html#SSAS.Admin)

- [Backing up an SSAS database](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/SSAS.Backup.html)

- [Restoring an SSAS database](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/SSAS.Restore.html)

This operation is also used with SQL Server Integration Services
(SSIS):

- [Administrative permissions on SSISDB](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/SSIS.Permissions.html)

- [Deploying an SSIS project](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/SSIS.Deploy.html)

This operation is also used with SQL Server Reporting Services
(SSRS):

- [Granting access to domain users](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/SSRS.Access.html#SSRS.Access.Grant)

- [Revoking system-level permissions](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/SSRS.Access.Revoke.html)

`rds_fn_task_status`

This operation shows the status of MSBI tasks:

- SSAS: [Monitoring the status of a deployment task](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/SSAS.Monitor.html)

- SSIS: [Monitoring the status of a deployment task](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/SSIS.Monitor.html)

- SSRS: [Monitoring the status of a task](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/SSRS.Monitor.html)

SSIS

`rds_drop_ssis_database`

[Dropping the SSISDB database](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/SSIS.DisableDrop.html#SSIS.Drop)

`rds_sqlagent_proxy`

[Creating an SSIS proxy](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/SSIS.Use.html#SSIS.Use.Proxy)SSRS

`rds_drop_ssrs_databases`

[Deleting the SSRS databases](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/SSRS.DisableDelete.html#SSRS.Drop)Resource Governor

`rds_create_resource_pool`

[Create resource Pool](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/ResourceGovernor.Using.html#ResourceGovernor.CreateResourcePool)

`rds_alter_resource_pool`

[Alter resource pool](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/ResourceGovernor.Using.html#ResourceGovernor.AlterResourcePool)

`rds_drop_resource_pool`

[Drop resource pool](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/ResourceGovernor.Using.html#ResourceGovernor.DropResourcePool)

`rds_create_workload_group`

[Create workload group](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/ResourceGovernor.Using.html#ResourceGovernor.CreateWorkloadGroup)

`rds_alter_workload_group`

[Alter workload group](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/ResourceGovernor.Using.html#ResourceGovernor.AlterWorkloadGroup)

`rds_drop_workload_group`

[Drop workload group](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/ResourceGovernor.Using.html#ResourceGovernor.DropWorkloadGroup)

`rds_create_classifier_function`

[Create and register classifier function](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/ResourceGovernor.Using.html#ResourceGovernor.ClassifierFunction)

`rds_drop_classifier_function`

[Drop classifier function](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/ResourceGovernor.Using.html#ResourceGovernor.DropClassifier)

`rds_alter_resource_governor_configuration`

[Resource governor configuration changes](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/ResourceGovernor.Using.html#ResourceGovernor.ConfigChanges)

`rds_bind_tempdb_metadata_to_resource_pool`

[Bind TempDB to a resource pool](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/ResourceGovernor.Using.html#ResourceGovernor.BindTempDB)

`rds_unbind_tempdb_metadata_from_resource_pool`

[Unbind TempDB from a resource pool](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/ResourceGovernor.Using.html#ResourceGovernor.UnbindTempDB)

`rds_cleanup_resource_governor`

[Cleanup resource governor](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/ResourceGovernor.Using.html#ResourceGovernor.Cleanup)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Unsupported and limited feature support

Local time zone
