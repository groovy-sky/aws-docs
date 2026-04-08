# Functions and stored procedures for Amazon RDS for Microsoft SQL Server

Following, you can find a list of the Amazon RDS functions and stored procedures that help
automate SQL Server tasks.

Task typeProcedure or functionWhere it's usedAdministrative tasks

`rds_drop_database`

[Dropping a database in an Amazon RDS for Microsoft SQL Server DB instance](appendix-sqlserver-commondbatasks-dropmirrordb.md)

`rds_failover_time`

[Determining the last failover time for Amazon RDS for SQL Server](appendix-sqlserver-commondbatasks-lastfailover.md)`rds_manage_view_db_permission`[Deny or allow viewing database names for Amazon RDS for SQL Server](appendix-sqlserver-commondbatasks-manageview.md)

`rds_modify_db_name`

[Renaming a Amazon RDS for Microsoft SQL Server database in a Multi-AZ deployment](appendix-sqlserver-commondbatasks-renamingdb.md)

`rds_read_error_log`

[Viewing error and agent logs](appendix-sqlserver-commondbatasks-logs.md#Appendix.SQLServer.CommonDBATasks.Logs.SP)

`rds_set_configuration`

This operation is used to set various DB instance
configurations:

- [Change data capture for Multi-AZ instances](appendix-sqlserver-commondbatasks-cdc.md#Appendix.SQLServer.CommonDBATasks.CDC.Multi-AZ)

- [Setting the retention period for trace and dump files](appendix-sqlserver-commondbatasks-tracefiles.md#Appendix.SQLServer.CommonDBATasks.TraceFiles.PurgeTraceFiles)

- [Compressing backup files](sqlserver-procedural-importing-native-compression.md)

`rds_set_database_online`

[Transitioning a Amazon RDS for SQL Server database from OFFLINE to ONLINE](appendix-sqlserver-commondbatasks-transitiononline.md)

`rds_set_system_database_sync_objects`

`rds_fn_get_system_database_sync_objects`

`rds_fn_server_object_last_sync_time`

[Turning on SQL Server Agent job replication](appendix-sqlserver-commondbatasks-agent.md#SQLServerAgent.Replicate)

`rds_show_configuration`

To see the values that are set using
`rds_set_configuration`, see these topics:

- [Change data capture for Multi-AZ instances](appendix-sqlserver-commondbatasks-cdc.md#Appendix.SQLServer.CommonDBATasks.CDC.Multi-AZ)

- [Setting the retention period for trace and dump files](appendix-sqlserver-commondbatasks-tracefiles.md#Appendix.SQLServer.CommonDBATasks.TraceFiles.PurgeTraceFiles)

`rds_shrink_tempdbfile`

[Shrinking the tempdb database](sqlserver-tempdb-shrinking.md)Change data capture (CDC)

`rds_cdc_disable_db`

[Disabling CDC](appendix-sqlserver-commondbatasks-cdc.md)

`rds_cdc_enable_db`

[Enabling CDC](appendix-sqlserver-commondbatasks-cdc.md)Database Mail

`rds_fn_sysmail_allitems`

[Viewing messages, logs, and attachments](sqlserver-dbmail-view.md)

`rds_fn_sysmail_event_log`

[Viewing messages, logs, and attachments](sqlserver-dbmail-view.md)

`rds_fn_sysmail_mailattachments`

[Viewing messages, logs, and attachments](sqlserver-dbmail-view.md)

`rds_sysmail_control`

This operation is used in starting and stopping the mail
queue:

- [Starting the mail queue](sqlserver-dbmail-startstop.md#SQLServer.DBMail.Start)

- [Stopping the mail queue](sqlserver-dbmail-startstop.md#SQLServer.DBMail.Stop)

`rds_sysmail_delete_mailitems_sp`

[Deleting messages](sqlserver-dbmail-delete.md)Native backup and restore

`rds_backup_database`

[Backing up a database](sqlserver-procedural-importing-native-using.md#SQLServer.Procedural.Importing.Native.Using.Backup)

`rds_cancel_task`

[Canceling a task](sqlserver-procedural-importing-native-using.md#SQLServer.Procedural.Importing.Native.Using.Cancel)

`rds_finish_restore`

[Finishing a database restore](sqlserver-procedural-importing-native-using.md#SQLServer.Procedural.Importing.Native.Finish.Restore)

`rds_restore_database`

[Restoring a database](sqlserver-procedural-importing-native-using.md#SQLServer.Procedural.Importing.Native.Using.Restore)

`rds_restore_log`

[Restoring a log](sqlserver-procedural-importing-native-using.md#SQLServer.Procedural.Importing.Native.Restore.Log)Amazon S3 file transfer

`rds_delete_from_filesystem`

[Deleting files on the RDS DB instance](appendix-sqlserver-options-s3-integration-using-deleting-files.md)

`rds_download_from_s3`

[Downloading files from an Amazon S3 bucket to a SQL Server DB instance](appendix-sqlserver-options-s3-integration-using.md#Appendix.SQLServer.Options.S3-integration.using.download)

`rds_gather_file_details`

[Listing files on the RDS DB instance](appendix-sqlserver-options-s3-integration-using-listing-files.md)

`rds_upload_to_s3`

[Uploading files from a SQL Server DB instance to an Amazon S3 bucket](appendix-sqlserver-options-s3-integration-using.md#Appendix.SQLServer.Options.S3-integration.using.upload)Microsoft Distributed Transaction Coordinator (MSDTC)

`rds_msdtc_transaction_tracing`

[Using transaction tracing](appendix-sqlserver-options-msdtc.md#MSDTC.Tracing)SQL Server Audit

`rds_fn_get_audit_file`

[Viewing audit logs](appendix-sqlserver-options-audit-auditrecords.md)Transparent Data Encryption

`rds_backup_tde_certificate`

`rds_drop_tde_certificate`

`rds_restore_tde_certificate`

`rds_fn_list_user_tde_certificates`

[Support for Transparent Data Encryption in SQL Server](appendix-sqlserver-options-tde.md)Microsoft Business Intelligence (MSBI)

`rds_msbi_task`

This operation is used with SQL Server Analysis Services
(SSAS):

- [Deploying SSAS projects on Amazon RDS](ssas-deploy.md)

- [Adding a domain user as a database administrator](ssas-use.md#SSAS.Admin)

- [Backing up an SSAS database](ssas-backup.md)

- [Restoring an SSAS database](ssas-restore.md)

This operation is also used with SQL Server Integration Services
(SSIS):

- [Administrative permissions on SSISDB](ssis-permissions.md)

- [Deploying an SSIS project](ssis-deploy.md)

This operation is also used with SQL Server Reporting Services
(SSRS):

- [Granting access to domain users](ssrs-access.md#SSRS.Access.Grant)

- [Revoking system-level permissions](ssrs-access-revoke.md)

`rds_fn_task_status`

This operation shows the status of MSBI tasks:

- SSAS: [Monitoring the status of a deployment task](ssas-monitor.md)

- SSIS: [Monitoring the status of a deployment task](ssis-monitor.md)

- SSRS: [Monitoring the status of a task](ssrs-monitor.md)

SSIS

`rds_drop_ssis_database`

[Dropping the SSISDB database](ssis-disabledrop.md#SSIS.Drop)

`rds_sqlagent_proxy`

[Creating an SSIS proxy](ssis-use.md#SSIS.Use.Proxy)SSRS

`rds_drop_ssrs_databases`

[Deleting the SSRS databases](ssrs-disabledelete.md#SSRS.Drop)Resource Governor

`rds_create_resource_pool`

[Create resource Pool](resourcegovernor-using.md#ResourceGovernor.CreateResourcePool)

`rds_alter_resource_pool`

[Alter resource pool](resourcegovernor-using.md#ResourceGovernor.AlterResourcePool)

`rds_drop_resource_pool`

[Drop resource pool](resourcegovernor-using.md#ResourceGovernor.DropResourcePool)

`rds_create_workload_group`

[Create workload group](resourcegovernor-using.md#ResourceGovernor.CreateWorkloadGroup)

`rds_alter_workload_group`

[Alter workload group](resourcegovernor-using.md#ResourceGovernor.AlterWorkloadGroup)

`rds_drop_workload_group`

[Drop workload group](resourcegovernor-using.md#ResourceGovernor.DropWorkloadGroup)

`rds_create_classifier_function`

[Create and register classifier function](resourcegovernor-using.md#ResourceGovernor.ClassifierFunction)

`rds_drop_classifier_function`

[Drop classifier function](resourcegovernor-using.md#ResourceGovernor.DropClassifier)

`rds_alter_resource_governor_configuration`

[Resource governor configuration changes](resourcegovernor-using.md#ResourceGovernor.ConfigChanges)

`rds_bind_tempdb_metadata_to_resource_pool`

[Bind TempDB to a resource pool](resourcegovernor-using.md#ResourceGovernor.BindTempDB)

`rds_unbind_tempdb_metadata_from_resource_pool`

[Unbind TempDB from a resource pool](resourcegovernor-using.md#ResourceGovernor.UnbindTempDB)

`rds_cleanup_resource_governor`

[Cleanup resource governor](resourcegovernor-using.md#ResourceGovernor.Cleanup)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Unsupported and limited feature support

Local time zone

All content copied from https://docs.aws.amazon.com/.
