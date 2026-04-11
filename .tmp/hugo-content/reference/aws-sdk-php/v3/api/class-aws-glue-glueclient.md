Menu

- [Aws](namespace-aws.md)
- [Glue](namespace-aws-glue.md)

## GlueClient     extends [AwsClient](class-aws-awsclient.md)   in package    - [Aws](package-aws.md)

This client is used to interact with the **AWS Glue** service.

## Supported API Versions

This class uses a _service description model_ that is associated at
runtime based on the `version` option given when constructing the
client. The `version` option will determine which API operations,
waiters, and paginators are available for a client. Creating a command or a
specific API operation can be done using magic methods (e.g.,
`$client->commandName(/** parameters */)`, or using the
`$client->getCommand` method of the client.

- [**2017-03-31**](api-glue-2017-03-31.md)

  - [BatchCreatePartition](api-glue-2017-03-31-batchcreatepartition.md)
  - [BatchDeleteConnection](api-glue-2017-03-31-batchdeleteconnection.md)
  - [BatchDeletePartition](api-glue-2017-03-31-batchdeletepartition.md)
  - [BatchDeleteTable](api-glue-2017-03-31-batchdeletetable.md)
  - [BatchDeleteTableVersion](api-glue-2017-03-31-batchdeletetableversion.md)
  - [BatchGetBlueprints](api-glue-2017-03-31-batchgetblueprints.md)
  - [BatchGetCrawlers](api-glue-2017-03-31-batchgetcrawlers.md)
  - [BatchGetCustomEntityTypes](api-glue-2017-03-31-batchgetcustomentitytypes.md)
  - [BatchGetDataQualityResult](api-glue-2017-03-31-batchgetdataqualityresult.md)
  - [BatchGetDevEndpoints](api-glue-2017-03-31-batchgetdevendpoints.md)
  - [BatchGetJobs](api-glue-2017-03-31-batchgetjobs.md)
  - [BatchGetPartition](api-glue-2017-03-31-batchgetpartition.md)
  - [BatchGetTableOptimizer](api-glue-2017-03-31-batchgettableoptimizer.md)
  - [BatchGetTriggers](api-glue-2017-03-31-batchgettriggers.md)
  - [BatchGetWorkflows](api-glue-2017-03-31-batchgetworkflows.md)
  - [BatchPutDataQualityStatisticAnnotation](api-glue-2017-03-31-batchputdataqualitystatisticannotation.md)
  - [BatchStopJobRun](api-glue-2017-03-31-batchstopjobrun.md)
  - [BatchUpdatePartition](api-glue-2017-03-31-batchupdatepartition.md)
  - [CancelDataQualityRuleRecommendationRun](api-glue-2017-03-31-canceldataqualityrulerecommendationrun.md)
  - [CancelDataQualityRulesetEvaluationRun](api-glue-2017-03-31-canceldataqualityrulesetevaluationrun.md)
  - [CancelMLTaskRun](api-glue-2017-03-31-cancelmltaskrun.md)
  - [CancelStatement](api-glue-2017-03-31-cancelstatement.md)
  - [CheckSchemaVersionValidity](api-glue-2017-03-31-checkschemaversionvalidity.md)
  - [CreateBlueprint](api-glue-2017-03-31-createblueprint.md)
  - [CreateCatalog](api-glue-2017-03-31-createcatalog.md)
  - [CreateClassifier](api-glue-2017-03-31-createclassifier.md)
  - [CreateColumnStatisticsTaskSettings](api-glue-2017-03-31-createcolumnstatisticstasksettings.md)
  - [CreateConnection](api-glue-2017-03-31-createconnection.md)
  - [CreateCrawler](api-glue-2017-03-31-createcrawler.md)
  - [CreateCustomEntityType](api-glue-2017-03-31-createcustomentitytype.md)
  - [CreateDataQualityRuleset](api-glue-2017-03-31-createdataqualityruleset.md)
  - [CreateDatabase](api-glue-2017-03-31-createdatabase.md)
  - [CreateDevEndpoint](api-glue-2017-03-31-createdevendpoint.md)
  - [CreateGlueIdentityCenterConfiguration](api-glue-2017-03-31-createglueidentitycenterconfiguration.md)
  - [CreateIntegration](api-glue-2017-03-31-createintegration.md)
  - [CreateIntegrationResourceProperty](api-glue-2017-03-31-createintegrationresourceproperty.md)
  - [CreateIntegrationTableProperties](api-glue-2017-03-31-createintegrationtableproperties.md)
  - [CreateJob](api-glue-2017-03-31-createjob.md)
  - [CreateMLTransform](api-glue-2017-03-31-createmltransform.md)
  - [CreatePartition](api-glue-2017-03-31-createpartition.md)
  - [CreatePartitionIndex](api-glue-2017-03-31-createpartitionindex.md)
  - [CreateRegistry](api-glue-2017-03-31-createregistry.md)
  - [CreateSchema](api-glue-2017-03-31-createschema.md)
  - [CreateScript](api-glue-2017-03-31-createscript.md)
  - [CreateSecurityConfiguration](api-glue-2017-03-31-createsecurityconfiguration.md)
  - [CreateSession](api-glue-2017-03-31-createsession.md)
  - [CreateTable](api-glue-2017-03-31-createtable.md)
  - [CreateTableOptimizer](api-glue-2017-03-31-createtableoptimizer.md)
  - [CreateTrigger](api-glue-2017-03-31-createtrigger.md)
  - [CreateUsageProfile](api-glue-2017-03-31-createusageprofile.md)
  - [CreateUserDefinedFunction](api-glue-2017-03-31-createuserdefinedfunction.md)
  - [CreateWorkflow](api-glue-2017-03-31-createworkflow.md)
  - [DeleteBlueprint](api-glue-2017-03-31-deleteblueprint.md)
  - [DeleteCatalog](api-glue-2017-03-31-deletecatalog.md)
  - [DeleteClassifier](api-glue-2017-03-31-deleteclassifier.md)
  - [DeleteColumnStatisticsForPartition](api-glue-2017-03-31-deletecolumnstatisticsforpartition.md)
  - [DeleteColumnStatisticsForTable](api-glue-2017-03-31-deletecolumnstatisticsfortable.md)
  - [DeleteColumnStatisticsTaskSettings](api-glue-2017-03-31-deletecolumnstatisticstasksettings.md)
  - [DeleteConnection](api-glue-2017-03-31-deleteconnection.md)
  - [DeleteConnectionType](api-glue-2017-03-31-deleteconnectiontype.md)
  - [DeleteCrawler](api-glue-2017-03-31-deletecrawler.md)
  - [DeleteCustomEntityType](api-glue-2017-03-31-deletecustomentitytype.md)
  - [DeleteDataQualityRuleset](api-glue-2017-03-31-deletedataqualityruleset.md)
  - [DeleteDatabase](api-glue-2017-03-31-deletedatabase.md)
  - [DeleteDevEndpoint](api-glue-2017-03-31-deletedevendpoint.md)
  - [DeleteGlueIdentityCenterConfiguration](api-glue-2017-03-31-deleteglueidentitycenterconfiguration.md)
  - [DeleteIntegration](api-glue-2017-03-31-deleteintegration.md)
  - [DeleteIntegrationResourceProperty](api-glue-2017-03-31-deleteintegrationresourceproperty.md)
  - [DeleteIntegrationTableProperties](api-glue-2017-03-31-deleteintegrationtableproperties.md)
  - [DeleteJob](api-glue-2017-03-31-deletejob.md)
  - [DeleteMLTransform](api-glue-2017-03-31-deletemltransform.md)
  - [DeletePartition](api-glue-2017-03-31-deletepartition.md)
  - [DeletePartitionIndex](api-glue-2017-03-31-deletepartitionindex.md)
  - [DeleteRegistry](api-glue-2017-03-31-deleteregistry.md)
  - [DeleteResourcePolicy](api-glue-2017-03-31-deleteresourcepolicy.md)
  - [DeleteSchema](api-glue-2017-03-31-deleteschema.md)
  - [DeleteSchemaVersions](api-glue-2017-03-31-deleteschemaversions.md)
  - [DeleteSecurityConfiguration](api-glue-2017-03-31-deletesecurityconfiguration.md)
  - [DeleteSession](api-glue-2017-03-31-deletesession.md)
  - [DeleteTable](api-glue-2017-03-31-deletetable.md)
  - [DeleteTableOptimizer](api-glue-2017-03-31-deletetableoptimizer.md)
  - [DeleteTableVersion](api-glue-2017-03-31-deletetableversion.md)
  - [DeleteTrigger](api-glue-2017-03-31-deletetrigger.md)
  - [DeleteUsageProfile](api-glue-2017-03-31-deleteusageprofile.md)
  - [DeleteUserDefinedFunction](api-glue-2017-03-31-deleteuserdefinedfunction.md)
  - [DeleteWorkflow](api-glue-2017-03-31-deleteworkflow.md)
  - [DescribeConnectionType](api-glue-2017-03-31-describeconnectiontype.md)
  - [DescribeEntity](api-glue-2017-03-31-describeentity.md)
  - [DescribeInboundIntegrations](api-glue-2017-03-31-describeinboundintegrations.md)
  - [DescribeIntegrations](api-glue-2017-03-31-describeintegrations.md)
  - [GetBlueprint](api-glue-2017-03-31-getblueprint.md)
  - [GetBlueprintRun](api-glue-2017-03-31-getblueprintrun.md)
  - [GetBlueprintRuns](api-glue-2017-03-31-getblueprintruns.md)
  - [GetCatalog](api-glue-2017-03-31-getcatalog.md)
  - [GetCatalogImportStatus](api-glue-2017-03-31-getcatalogimportstatus.md)
  - [GetCatalogs](api-glue-2017-03-31-getcatalogs.md)
  - [GetClassifier](api-glue-2017-03-31-getclassifier.md)
  - [GetClassifiers](api-glue-2017-03-31-getclassifiers.md)
  - [GetColumnStatisticsForPartition](api-glue-2017-03-31-getcolumnstatisticsforpartition.md)
  - [GetColumnStatisticsForTable](api-glue-2017-03-31-getcolumnstatisticsfortable.md)
  - [GetColumnStatisticsTaskRun](api-glue-2017-03-31-getcolumnstatisticstaskrun.md)
  - [GetColumnStatisticsTaskRuns](api-glue-2017-03-31-getcolumnstatisticstaskruns.md)
  - [GetColumnStatisticsTaskSettings](api-glue-2017-03-31-getcolumnstatisticstasksettings.md)
  - [GetConnection](api-glue-2017-03-31-getconnection.md)
  - [GetConnections](api-glue-2017-03-31-getconnections.md)
  - [GetCrawler](api-glue-2017-03-31-getcrawler.md)
  - [GetCrawlerMetrics](api-glue-2017-03-31-getcrawlermetrics.md)
  - [GetCrawlers](api-glue-2017-03-31-getcrawlers.md)
  - [GetCustomEntityType](api-glue-2017-03-31-getcustomentitytype.md)
  - [GetDataCatalogEncryptionSettings](api-glue-2017-03-31-getdatacatalogencryptionsettings.md)
  - [GetDataQualityModel](api-glue-2017-03-31-getdataqualitymodel.md)
  - [GetDataQualityModelResult](api-glue-2017-03-31-getdataqualitymodelresult.md)
  - [GetDataQualityResult](api-glue-2017-03-31-getdataqualityresult.md)
  - [GetDataQualityRuleRecommendationRun](api-glue-2017-03-31-getdataqualityrulerecommendationrun.md)
  - [GetDataQualityRuleset](api-glue-2017-03-31-getdataqualityruleset.md)
  - [GetDataQualityRulesetEvaluationRun](api-glue-2017-03-31-getdataqualityrulesetevaluationrun.md)
  - [GetDatabase](api-glue-2017-03-31-getdatabase.md)
  - [GetDatabases](api-glue-2017-03-31-getdatabases.md)
  - [GetDataflowGraph](api-glue-2017-03-31-getdataflowgraph.md)
  - [GetDevEndpoint](api-glue-2017-03-31-getdevendpoint.md)
  - [GetDevEndpoints](api-glue-2017-03-31-getdevendpoints.md)
  - [GetEntityRecords](api-glue-2017-03-31-getentityrecords.md)
  - [GetGlueIdentityCenterConfiguration](api-glue-2017-03-31-getglueidentitycenterconfiguration.md)
  - [GetIntegrationResourceProperty](api-glue-2017-03-31-getintegrationresourceproperty.md)
  - [GetIntegrationTableProperties](api-glue-2017-03-31-getintegrationtableproperties.md)
  - [GetJob](api-glue-2017-03-31-getjob.md)
  - [GetJobBookmark](api-glue-2017-03-31-getjobbookmark.md)
  - [GetJobRun](api-glue-2017-03-31-getjobrun.md)
  - [GetJobRuns](api-glue-2017-03-31-getjobruns.md)
  - [GetJobs](api-glue-2017-03-31-getjobs.md)
  - [GetMLTaskRun](api-glue-2017-03-31-getmltaskrun.md)
  - [GetMLTaskRuns](api-glue-2017-03-31-getmltaskruns.md)
  - [GetMLTransform](api-glue-2017-03-31-getmltransform.md)
  - [GetMLTransforms](api-glue-2017-03-31-getmltransforms.md)
  - [GetMapping](api-glue-2017-03-31-getmapping.md)
  - [GetMaterializedViewRefreshTaskRun](api-glue-2017-03-31-getmaterializedviewrefreshtaskrun.md)
  - [GetPartition](api-glue-2017-03-31-getpartition.md)
  - [GetPartitionIndexes](api-glue-2017-03-31-getpartitionindexes.md)
  - [GetPartitions](api-glue-2017-03-31-getpartitions.md)
  - [GetPlan](api-glue-2017-03-31-getplan.md)
  - [GetRegistry](api-glue-2017-03-31-getregistry.md)
  - [GetResourcePolicies](api-glue-2017-03-31-getresourcepolicies.md)
  - [GetResourcePolicy](api-glue-2017-03-31-getresourcepolicy.md)
  - [GetSchema](api-glue-2017-03-31-getschema.md)
  - [GetSchemaByDefinition](api-glue-2017-03-31-getschemabydefinition.md)
  - [GetSchemaVersion](api-glue-2017-03-31-getschemaversion.md)
  - [GetSchemaVersionsDiff](api-glue-2017-03-31-getschemaversionsdiff.md)
  - [GetSecurityConfiguration](api-glue-2017-03-31-getsecurityconfiguration.md)
  - [GetSecurityConfigurations](api-glue-2017-03-31-getsecurityconfigurations.md)
  - [GetSession](api-glue-2017-03-31-getsession.md)
  - [GetStatement](api-glue-2017-03-31-getstatement.md)
  - [GetTable](api-glue-2017-03-31-gettable.md)
  - [GetTableOptimizer](api-glue-2017-03-31-gettableoptimizer.md)
  - [GetTableVersion](api-glue-2017-03-31-gettableversion.md)
  - [GetTableVersions](api-glue-2017-03-31-gettableversions.md)
  - [GetTables](api-glue-2017-03-31-gettables.md)
  - [GetTags](api-glue-2017-03-31-gettags.md)
  - [GetTrigger](api-glue-2017-03-31-gettrigger.md)
  - [GetTriggers](api-glue-2017-03-31-gettriggers.md)
  - [GetUnfilteredPartitionMetadata](api-glue-2017-03-31-getunfilteredpartitionmetadata.md)
  - [GetUnfilteredPartitionsMetadata](api-glue-2017-03-31-getunfilteredpartitionsmetadata.md)
  - [GetUnfilteredTableMetadata](api-glue-2017-03-31-getunfilteredtablemetadata.md)
  - [GetUsageProfile](api-glue-2017-03-31-getusageprofile.md)
  - [GetUserDefinedFunction](api-glue-2017-03-31-getuserdefinedfunction.md)
  - [GetUserDefinedFunctions](api-glue-2017-03-31-getuserdefinedfunctions.md)
  - [GetWorkflow](api-glue-2017-03-31-getworkflow.md)
  - [GetWorkflowRun](api-glue-2017-03-31-getworkflowrun.md)
  - [GetWorkflowRunProperties](api-glue-2017-03-31-getworkflowrunproperties.md)
  - [GetWorkflowRuns](api-glue-2017-03-31-getworkflowruns.md)
  - [ImportCatalogToGlue](api-glue-2017-03-31-importcatalogtoglue.md)
  - [ListBlueprints](api-glue-2017-03-31-listblueprints.md)
  - [ListColumnStatisticsTaskRuns](api-glue-2017-03-31-listcolumnstatisticstaskruns.md)
  - [ListConnectionTypes](api-glue-2017-03-31-listconnectiontypes.md)
  - [ListCrawlers](api-glue-2017-03-31-listcrawlers.md)
  - [ListCrawls](api-glue-2017-03-31-listcrawls.md)
  - [ListCustomEntityTypes](api-glue-2017-03-31-listcustomentitytypes.md)
  - [ListDataQualityResults](api-glue-2017-03-31-listdataqualityresults.md)
  - [ListDataQualityRuleRecommendationRuns](api-glue-2017-03-31-listdataqualityrulerecommendationruns.md)
  - [ListDataQualityRulesetEvaluationRuns](api-glue-2017-03-31-listdataqualityrulesetevaluationruns.md)
  - [ListDataQualityRulesets](api-glue-2017-03-31-listdataqualityrulesets.md)
  - [ListDataQualityStatisticAnnotations](api-glue-2017-03-31-listdataqualitystatisticannotations.md)
  - [ListDataQualityStatistics](api-glue-2017-03-31-listdataqualitystatistics.md)
  - [ListDevEndpoints](api-glue-2017-03-31-listdevendpoints.md)
  - [ListEntities](api-glue-2017-03-31-listentities.md)
  - [ListIntegrationResourceProperties](api-glue-2017-03-31-listintegrationresourceproperties.md)
  - [ListJobs](api-glue-2017-03-31-listjobs.md)
  - [ListMLTransforms](api-glue-2017-03-31-listmltransforms.md)
  - [ListMaterializedViewRefreshTaskRuns](api-glue-2017-03-31-listmaterializedviewrefreshtaskruns.md)
  - [ListRegistries](api-glue-2017-03-31-listregistries.md)
  - [ListSchemaVersions](api-glue-2017-03-31-listschemaversions.md)
  - [ListSchemas](api-glue-2017-03-31-listschemas.md)
  - [ListSessions](api-glue-2017-03-31-listsessions.md)
  - [ListStatements](api-glue-2017-03-31-liststatements.md)
  - [ListTableOptimizerRuns](api-glue-2017-03-31-listtableoptimizerruns.md)
  - [ListTriggers](api-glue-2017-03-31-listtriggers.md)
  - [ListUsageProfiles](api-glue-2017-03-31-listusageprofiles.md)
  - [ListWorkflows](api-glue-2017-03-31-listworkflows.md)
  - [ModifyIntegration](api-glue-2017-03-31-modifyintegration.md)
  - [PutDataCatalogEncryptionSettings](api-glue-2017-03-31-putdatacatalogencryptionsettings.md)
  - [PutDataQualityProfileAnnotation](api-glue-2017-03-31-putdataqualityprofileannotation.md)
  - [PutResourcePolicy](api-glue-2017-03-31-putresourcepolicy.md)
  - [PutSchemaVersionMetadata](api-glue-2017-03-31-putschemaversionmetadata.md)
  - [PutWorkflowRunProperties](api-glue-2017-03-31-putworkflowrunproperties.md)
  - [QuerySchemaVersionMetadata](api-glue-2017-03-31-queryschemaversionmetadata.md)
  - [RegisterConnectionType](api-glue-2017-03-31-registerconnectiontype.md)
  - [RegisterSchemaVersion](api-glue-2017-03-31-registerschemaversion.md)
  - [RemoveSchemaVersionMetadata](api-glue-2017-03-31-removeschemaversionmetadata.md)
  - [ResetJobBookmark](api-glue-2017-03-31-resetjobbookmark.md)
  - [ResumeWorkflowRun](api-glue-2017-03-31-resumeworkflowrun.md)
  - [RunStatement](api-glue-2017-03-31-runstatement.md)
  - [SearchTables](api-glue-2017-03-31-searchtables.md)
  - [StartBlueprintRun](api-glue-2017-03-31-startblueprintrun.md)
  - [StartColumnStatisticsTaskRun](api-glue-2017-03-31-startcolumnstatisticstaskrun.md)
  - [StartColumnStatisticsTaskRunSchedule](api-glue-2017-03-31-startcolumnstatisticstaskrunschedule.md)
  - [StartCrawler](api-glue-2017-03-31-startcrawler.md)
  - [StartCrawlerSchedule](api-glue-2017-03-31-startcrawlerschedule.md)
  - [StartDataQualityRuleRecommendationRun](api-glue-2017-03-31-startdataqualityrulerecommendationrun.md)
  - [StartDataQualityRulesetEvaluationRun](api-glue-2017-03-31-startdataqualityrulesetevaluationrun.md)
  - [StartExportLabelsTaskRun](api-glue-2017-03-31-startexportlabelstaskrun.md)
  - [StartImportLabelsTaskRun](api-glue-2017-03-31-startimportlabelstaskrun.md)
  - [StartJobRun](api-glue-2017-03-31-startjobrun.md)
  - [StartMLEvaluationTaskRun](api-glue-2017-03-31-startmlevaluationtaskrun.md)
  - [StartMLLabelingSetGenerationTaskRun](api-glue-2017-03-31-startmllabelingsetgenerationtaskrun.md)
  - [StartMaterializedViewRefreshTaskRun](api-glue-2017-03-31-startmaterializedviewrefreshtaskrun.md)
  - [StartTrigger](api-glue-2017-03-31-starttrigger.md)
  - [StartWorkflowRun](api-glue-2017-03-31-startworkflowrun.md)
  - [StopColumnStatisticsTaskRun](api-glue-2017-03-31-stopcolumnstatisticstaskrun.md)
  - [StopColumnStatisticsTaskRunSchedule](api-glue-2017-03-31-stopcolumnstatisticstaskrunschedule.md)
  - [StopCrawler](api-glue-2017-03-31-stopcrawler.md)
  - [StopCrawlerSchedule](api-glue-2017-03-31-stopcrawlerschedule.md)
  - [StopMaterializedViewRefreshTaskRun](api-glue-2017-03-31-stopmaterializedviewrefreshtaskrun.md)
  - [StopSession](api-glue-2017-03-31-stopsession.md)
  - [StopTrigger](api-glue-2017-03-31-stoptrigger.md)
  - [StopWorkflowRun](api-glue-2017-03-31-stopworkflowrun.md)
  - [TagResource](api-glue-2017-03-31-tagresource.md)
  - [TestConnection](api-glue-2017-03-31-testconnection.md)
  - [UntagResource](api-glue-2017-03-31-untagresource.md)
  - [UpdateBlueprint](api-glue-2017-03-31-updateblueprint.md)
  - [UpdateCatalog](api-glue-2017-03-31-updatecatalog.md)
  - [UpdateClassifier](api-glue-2017-03-31-updateclassifier.md)
  - [UpdateColumnStatisticsForPartition](api-glue-2017-03-31-updatecolumnstatisticsforpartition.md)
  - [UpdateColumnStatisticsForTable](api-glue-2017-03-31-updatecolumnstatisticsfortable.md)
  - [UpdateColumnStatisticsTaskSettings](api-glue-2017-03-31-updatecolumnstatisticstasksettings.md)
  - [UpdateConnection](api-glue-2017-03-31-updateconnection.md)
  - [UpdateCrawler](api-glue-2017-03-31-updatecrawler.md)
  - [UpdateCrawlerSchedule](api-glue-2017-03-31-updatecrawlerschedule.md)
  - [UpdateDataQualityRuleset](api-glue-2017-03-31-updatedataqualityruleset.md)
  - [UpdateDatabase](api-glue-2017-03-31-updatedatabase.md)
  - [UpdateDevEndpoint](api-glue-2017-03-31-updatedevendpoint.md)
  - [UpdateGlueIdentityCenterConfiguration](api-glue-2017-03-31-updateglueidentitycenterconfiguration.md)
  - [UpdateIntegrationResourceProperty](api-glue-2017-03-31-updateintegrationresourceproperty.md)
  - [UpdateIntegrationTableProperties](api-glue-2017-03-31-updateintegrationtableproperties.md)
  - [UpdateJob](api-glue-2017-03-31-updatejob.md)
  - [UpdateJobFromSourceControl](api-glue-2017-03-31-updatejobfromsourcecontrol.md)
  - [UpdateMLTransform](api-glue-2017-03-31-updatemltransform.md)
  - [UpdatePartition](api-glue-2017-03-31-updatepartition.md)
  - [UpdateRegistry](api-glue-2017-03-31-updateregistry.md)
  - [UpdateSchema](api-glue-2017-03-31-updateschema.md)
  - [UpdateSourceControlFromJob](api-glue-2017-03-31-updatesourcecontrolfromjob.md)
  - [UpdateTable](api-glue-2017-03-31-updatetable.md)
  - [UpdateTableOptimizer](api-glue-2017-03-31-updatetableoptimizer.md)
  - [UpdateTrigger](api-glue-2017-03-31-updatetrigger.md)
  - [UpdateUsageProfile](api-glue-2017-03-31-updateusageprofile.md)
  - [UpdateUserDefinedFunction](api-glue-2017-03-31-updateuserdefinedfunction.md)
  - [UpdateWorkflow](api-glue-2017-03-31-updateworkflow.md)

## Examples

### Basics, Actions and Scenarios

The following code examples show you how to perform actions and implement common scenarios by using the AWS SDK for PHP with AWS Glue.

- [See examples on AWS Docs](../../../sdk-for-php/v3/developer-guide/php-glue-code-examples.md)

### Table of Contents  [header link](class-aws-glue-glueclient-toc.md)

#### Methods  [header link](class-aws-glue-glueclient-toc-methods.md)

[\_\_call()](class-aws-awsclienttrait.md#method___call)
: mixed [\_\_construct()](class-aws-awsclient.md#method___construct)
: mixed The client constructor accepts the following options:[\_\_sleep()](class-aws-awsclient.md#method___sleep)
: mixed [execute()](class-aws-awsclienttrait.md#method_execute)
: mixed [executeAsync()](class-aws-awsclienttrait.md#method_executeAsync)
: mixed [factory()](class-aws-awsclient.md#method_factory)
: static [getApi()](class-aws-awsclienttrait.md#method_getApi)
: [Service](class-aws-api-service.md)[getArguments()](class-aws-awsclient.md#method_getArguments)
: array<string\|int, mixed> Get an array of client constructor arguments used by the client.[getClientBuiltIns()](class-aws-awsclient.md#method_getClientBuiltIns)
: array<string\|int, mixed> Provides the set of built-in keys and values
used for endpoint resolution[getClientContextParams()](class-aws-awsclient.md#method_getClientContextParams)
: array<string\|int, mixed> Provides the set of service context parameter
key-value pairs used for endpoint resolution.[getCommand()](class-aws-awsclienttrait.md#method_getCommand)
: [CommandInterface](class-aws-commandinterface.md)[getConfig()](class-aws-awsclient.md#method_getConfig)
: mixed\|null Get a client configuration value.[getCredentials()](class-aws-awsclient.md#method_getCredentials)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Returns a promise that is fulfilled with an
{@see \\Aws\\Credentials\\CredentialsInterface} object.[getEndpoint()](class-aws-awsclient.md#method_getEndpoint)
: [UriInterface](class-psr-http-message-uriinterface.md)Gets the default endpoint, or base URL, used by the client.[getEndpointProvider()](class-aws-awsclient.md#method_getEndpointProvider)
: mixed [getEndpointProviderArgs()](class-aws-awsclient.md#method_getEndpointProviderArgs)
: array<string\|int, mixed> Retrieves arguments to be used in endpoint resolution.[getHandlerList()](class-aws-awsclient.md#method_getHandlerList)
: [HandlerList](class-aws-handlerlist.md)Get the handler list used to transfer commands.[getIterator()](class-aws-awsclienttrait.md#method_getIterator)
: mixed [getPaginator()](class-aws-awsclienttrait.md#method_getPaginator)
: mixed [getRegion()](class-aws-awsclient.md#method_getRegion)
: string Get the region to which the client is configured to send requests.[getSignatureProvider()](class-aws-awsclient.md#method_getSignatureProvider)
: callable Get the signature\_provider function of the client.[getToken()](class-aws-awsclient.md#method_getToken)
: mixed [getWaiter()](class-aws-awsclienttrait.md#method_getWaiter)
: mixed [waitUntil()](class-aws-awsclienttrait.md#method_waitUntil)
: mixed

### Methods  [header link](class-aws-glue-glueclient-methods.md)

#### \_\_call()  [header link](class-aws-awsclienttrait.md\#method___call)

`
    public
                    __call(mixed $name, array<string|int, mixed> $args) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>

#### \_\_construct()  [header link](class-aws-awsclient.md\#method___construct)

The client constructor accepts the following options:

`
    public
                    __construct(array<string|int, mixed> $args) : mixed`

- api\_provider: (callable) An optional PHP callable that accepts a
type, service, and version argument, and returns an array of
corresponding configuration data. The type value can be one of api,
waiter, or paginator.
- credentials:
(Aws\\Credentials\\CredentialsInterface\|array\|bool\|callable) Specifies
the credentials used to sign requests. Provide an
Aws\\Credentials\\CredentialsInterface object, an associative array of
"key", "secret", and an optional "token" key, `false` to use null
credentials, or a callable credentials provider used to create
credentials or return null. See Aws\\Credentials\\CredentialProvider for
a list of built-in credentials providers. If no credentials are
provided, the SDK will attempt to load them from the environment.
- token:
(Aws\\Token\\TokenInterface\|array\|bool\|callable) Specifies
the token used to authorize requests. Provide an
Aws\\Token\\TokenInterface object, an associative array of
"token" and an optional "expires" key, `false` to use no
token, or a callable token provider used to create a
token or return null. See Aws\\Token\\TokenProvider for
a list of built-in token providers. If no token is
provided, the SDK will attempt to load one from the environment.
- csm:
(Aws\\ClientSideMonitoring\\ConfigurationInterface\|array\|callable) Specifies
the credentials used to sign requests. Provide an
Aws\\ClientSideMonitoring\\ConfigurationInterface object, a callable
configuration provider used to create client-side monitoring configuration,
`false` to disable csm, or an associative array with the following keys:
enabled: (bool) Set to true to enable client-side monitoring, defaults
to false; host: (string) the host location to send monitoring events to,
defaults to 127.0.0.1; port: (int) The port used for the host connection,
defaults to 31000; client\_id: (string) An identifier for this project
- debug: (bool\|array) Set to true to display debug information when
sending requests. Alternatively, you can provide an associative array
with the following keys: logfn: (callable) Function that is invoked
with log messages; stream\_size: (int) When the size of a stream is
greater than this number, the stream data will not be logged (set to
"0" to not log any stream data); scrub\_auth: (bool) Set to false to
disable the scrubbing of auth data from the logged messages; http:
(bool) Set to false to disable the "debug" feature of lower level HTTP
adapters (e.g., verbose curl output).
- stats: (bool\|array) Set to true to gather transfer statistics on
requests sent. Alternatively, you can provide an associative array with
the following keys: retries: (bool) Set to false to disable reporting
on retries attempted; http: (bool) Set to true to enable collecting
statistics from lower level HTTP adapters (e.g., values returned in
GuzzleHttp\\TransferStats). HTTP handlers must support an
`http_stats_receiver` option for this to have an effect; timer: (bool)
Set to true to enable a command timer that reports the total wall clock
time spent on an operation in seconds.
- disable\_host\_prefix\_injection: (bool) Set to true to disable host prefix
injection logic for services that use it. This disables the entire
prefix injection, including the portions supplied by user-defined
parameters. Setting this flag will have no effect on services that do
not use host prefix injection.
- endpoint: (string) The full URI of the webservice. This is only
required when connecting to a custom endpoint (e.g., a local version
of S3).
- endpoint\_discovery: (Aws\\EndpointDiscovery\\ConfigurationInterface,
Aws\\CacheInterface, array, callable) Settings for endpoint discovery.
Provide an instance of Aws\\EndpointDiscovery\\ConfigurationInterface,
an instance Aws\\CacheInterface, a callable that provides a promise for
a Configuration object, or an associative array with the following
keys: enabled: (bool) Set to true to enable endpoint discovery, false
to explicitly disable it, defaults to false; cache\_limit: (int) The
maximum number of keys in the endpoints cache, defaults to 1000.
- endpoint\_provider: (callable) An optional PHP callable that
accepts a hash of options including a "service" and "region" key and
returns NULL or a hash of endpoint data, of which the "endpoint" key
is required. See Aws\\Endpoint\\EndpointProvider for a list of built-in
providers.
- handler: (callable) A handler that accepts a command object,
request object and returns a promise that is fulfilled with an
Aws\\ResultInterface object or rejected with an
Aws\\Exception\\AwsException. A handler does not accept a next handler
as it is terminal and expected to fulfill a command. If no handler is
provided, a default Guzzle handler will be utilized.
- http: (array, default=array(0)) Set to an array of SDK request
options to apply to each request (e.g., proxy, verify, etc.).
- http\_handler: (callable) An HTTP handler is a function that
accepts a PSR-7 request object and returns a promise that is fulfilled
with a PSR-7 response object or rejected with an array of exception
data. NOTE: This option supersedes any provided "handler" option.
- idempotency\_auto\_fill: (bool\|callable) Set to false to disable SDK to
populate parameters that enabled 'idempotencyToken' trait with a random
UUID v4 value on your behalf. Using default value 'true' still allows
parameter value to be overwritten when provided. Note: auto-fill only
works when cryptographically secure random bytes generator functions
(random\_bytes, openssl\_random\_pseudo\_bytes or mcrypt\_create\_iv) can be
found. You may also provide a callable source of random bytes.
- profile: (string) Allows you to specify which profile to use when
credentials are created from the AWS credentials file in your HOME
directory. This setting overrides the AWS\_PROFILE environment
variable. Note: Specifying "profile" will cause the "credentials" key
to be ignored.
- region: (string, required) Region to connect to. See
http://docs.aws.amazon.com/general/latest/gr/rande.html for a list of
available regions.
- retries: (int, Aws\\Retry\\ConfigurationInterface, Aws\\CacheInterface,
array, callable) Configures the retry mode and maximum number of
allowed retries for a client (pass 0 to disable retries). Provide an
integer for 'legacy' mode with the specified number of retries.
Otherwise provide an instance of Aws\\Retry\\ConfigurationInterface, an
instance of Aws\\CacheInterface, a callable function, or an array with
the following keys: mode: (string) Set to 'legacy', 'standard' (uses
retry quota management), or 'adapative' (an experimental mode that adds
client-side rate limiting to standard mode); max\_attempts (int) The
maximum number of attempts for a given request.
- scheme: (string, default=string(5) "https") URI scheme to use when
connecting connect. The SDK will utilize "https" endpoints (i.e.,
utilize SSL/TLS connections) by default. You can attempt to connect to
a service over an unencrypted "http" endpoint by setting `scheme` to
"http".
- signature\_provider: (callable) A callable that accepts a signature
version name (e.g., "v4"), a service name, and region, and
returns a SignatureInterface object or null. This provider is used to
create signers utilized by the client. See
Aws\\Signature\\SignatureProvider for a list of built-in providers
- signature\_version: (string) A string representing a custom
signature version to use with a service (e.g., v4). Note that
per/operation signature version MAY override this requested signature
version.
- use\_aws\_shared\_config\_files: (bool, default=bool(true)) Set to false to
disable checking for shared config file in '~/.aws/config' and
'~/.aws/credentials'. This will override the AWS\_CONFIG\_FILE
environment variable.
- validate: (bool, default=bool(true)) Set to false to disable
client-side parameter validation.
- version: (string, required) The version of the webservice to
utilize (e.g., 2006-03-01).
- account\_id\_endpoint\_mode: (string, default(preferred)) this option
decides whether credentials should resolve an accountId value,
which is going to be used as part of the endpoint resolution.
The valid values for this option are:
  - preferred: when this value is set then, a warning is logged when
    accountId is empty in the resolved identity.
  - required: when this value is set then, an exception is thrown when
    accountId is empty in the resolved identity.
  - disabled: when this value is set then, the validation for if accountId
    was resolved or not, is ignored.
- ua\_append: (string, array) To pass custom user agent parameters.
- app\_id: (string) an optional application specific identifier that can be set.
When set it will be appended to the User-Agent header of every request
in the form of App/{AppId}. This variable is sourced from environment
variable AWS\_SDK\_UA\_APP\_ID or the shared config profile attribute sdk\_ua\_app\_id.
See https://docs.aws.amazon.com/sdkref/latest/guide/settings-reference.html for
more information on environment variables and shared config settings.

##### Parameters

$args
: array<string\|int, mixed>

Client configuration arguments.

##### Tags  [header link](class-aws-awsclient.md\#method___construct\#tags)

throwsInvalidArgumentException

if any required options are missing or
the service is not supported.

#### \_\_sleep()  [header link](class-aws-awsclient.md\#method___sleep)

`
    public
                    __sleep() : mixed`

#### execute()  [header link](class-aws-awsclienttrait.md\#method_execute)

`
    public
                    execute(CommandInterface $command) : mixed`

##### Parameters

$command
: [CommandInterface](class-aws-commandinterface.md)

#### executeAsync()  [header link](class-aws-awsclienttrait.md\#method_executeAsync)

`
    public
                    executeAsync(CommandInterface $command) : mixed`

##### Parameters

$command
: [CommandInterface](class-aws-commandinterface.md)

#### factory()  [header link](class-aws-awsclient.md\#method_factory)

`
    public
            static        factory([array<string|int, mixed> $config = [] ]) : static`

##### Parameters

$config
: array<string\|int, mixed>
= \[\]

##### Tags  [header link](class-aws-awsclient.md\#method_factory\#tags)

deprecated

##### Return values

static

#### getApi()  [header link](class-aws-awsclienttrait.md\#method_getApi)

`
    public
    abstract                getApi() : Service`

##### Return values

[Service](class-aws-api-service.md)

#### getArguments()  [header link](class-aws-awsclient.md\#method_getArguments)

Get an array of client constructor arguments used by the client.

`
    public
            static        getArguments() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getClientBuiltIns()  [header link](class-aws-awsclient.md\#method_getClientBuiltIns)

Provides the set of built-in keys and values
used for endpoint resolution

`
    public
                    getClientBuiltIns() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getClientContextParams()  [header link](class-aws-awsclient.md\#method_getClientContextParams)

Provides the set of service context parameter
key-value pairs used for endpoint resolution.

`
    public
                    getClientContextParams() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getCommand()  [header link](class-aws-awsclienttrait.md\#method_getCommand)

`
    public
    abstract                getCommand(string $name[, array<string|int, mixed> $args = [] ]) : CommandInterface`

##### Parameters

$name
: string$args
: array<string\|int, mixed>
= \[\]

##### Return values

[CommandInterface](class-aws-commandinterface.md)

#### getConfig()  [header link](class-aws-awsclient.md\#method_getConfig)

Get a client configuration value.

`
    public
                    getConfig([mixed $option = null ]) : mixed|null`

##### Parameters

$option
: mixed
= null

The option to retrieve. Pass null to retrieve
all options.

##### Return values

mixed\|null

#### getCredentials()  [header link](class-aws-awsclient.md\#method_getCredentials)

Returns a promise that is fulfilled with an
{@see \\Aws\\Credentials\\CredentialsInterface} object.

`
    public
                    getCredentials() : PromiseInterface`

If you need the credentials synchronously, then call the wait() method
on the returned promise.

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)

#### getEndpoint()  [header link](class-aws-awsclient.md\#method_getEndpoint)

Gets the default endpoint, or base URL, used by the client.

`
    public
                    getEndpoint() : UriInterface`

##### Return values

[UriInterface](class-psr-http-message-uriinterface.md)

#### getEndpointProvider()  [header link](class-aws-awsclient.md\#method_getEndpointProvider)

`
    public
                    getEndpointProvider() : mixed`

#### getEndpointProviderArgs()  [header link](class-aws-awsclient.md\#method_getEndpointProviderArgs)

Retrieves arguments to be used in endpoint resolution.

`
    public
                    getEndpointProviderArgs() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getHandlerList()  [header link](class-aws-awsclient.md\#method_getHandlerList)

Get the handler list used to transfer commands.

`
    public
                    getHandlerList() : HandlerList`

This list can be modified to add middleware or to change the underlying
handler used to send HTTP requests.

##### Return values

[HandlerList](class-aws-handlerlist.md)

#### getIterator()  [header link](class-aws-awsclienttrait.md\#method_getIterator)

`
    public
                    getIterator(mixed $name[, array<string|int, mixed> $args = [] ]) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>
= \[\]

#### getPaginator()  [header link](class-aws-awsclienttrait.md\#method_getPaginator)

`
    public
                    getPaginator(mixed $name[, array<string|int, mixed> $args = [] ]) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>
= \[\]

#### getRegion()  [header link](class-aws-awsclient.md\#method_getRegion)

Get the region to which the client is configured to send requests.

`
    public
                    getRegion() : string`

##### Return values

string

#### getSignatureProvider()  [header link](class-aws-awsclient.md\#method_getSignatureProvider)

Get the signature\_provider function of the client.

`
    public
        final            getSignatureProvider() : callable`

##### Return values

callable

#### getToken()  [header link](class-aws-awsclient.md\#method_getToken)

`
    public
                    getToken() : mixed`

#### getWaiter()  [header link](class-aws-awsclienttrait.md\#method_getWaiter)

`
    public
                    getWaiter(mixed $name[, array<string|int, mixed> $args = [] ]) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>
= \[\]

#### waitUntil()  [header link](class-aws-awsclienttrait.md\#method_waitUntil)

`
    public
                    waitUntil(mixed $name[, array<string|int, mixed> $args = [] ]) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>
= \[\]
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](class-aws-glue-glueclient-toc-methods.md)
- Methods
  - [\_\_call()](class-aws-awsclienttrait.md#method___call)
  - [\_\_construct()](class-aws-awsclient.md#method___construct)
  - [\_\_sleep()](class-aws-awsclient.md#method___sleep)
  - [execute()](class-aws-awsclienttrait.md#method_execute)
  - [executeAsync()](class-aws-awsclienttrait.md#method_executeAsync)
  - [factory()](class-aws-awsclient.md#method_factory)
  - [getApi()](class-aws-awsclienttrait.md#method_getApi)
  - [getArguments()](class-aws-awsclient.md#method_getArguments)
  - [getClientBuiltIns()](class-aws-awsclient.md#method_getClientBuiltIns)
  - [getClientContextParams()](class-aws-awsclient.md#method_getClientContextParams)
  - [getCommand()](class-aws-awsclienttrait.md#method_getCommand)
  - [getConfig()](class-aws-awsclient.md#method_getConfig)
  - [getCredentials()](class-aws-awsclient.md#method_getCredentials)
  - [getEndpoint()](class-aws-awsclient.md#method_getEndpoint)
  - [getEndpointProvider()](class-aws-awsclient.md#method_getEndpointProvider)
  - [getEndpointProviderArgs()](class-aws-awsclient.md#method_getEndpointProviderArgs)
  - [getHandlerList()](class-aws-awsclient.md#method_getHandlerList)
  - [getIterator()](class-aws-awsclienttrait.md#method_getIterator)
  - [getPaginator()](class-aws-awsclienttrait.md#method_getPaginator)
  - [getRegion()](class-aws-awsclient.md#method_getRegion)
  - [getSignatureProvider()](class-aws-awsclient.md#method_getSignatureProvider)
  - [getToken()](class-aws-awsclient.md#method_getToken)
  - [getWaiter()](class-aws-awsclienttrait.md#method_getWaiter)
  - [waitUntil()](class-aws-awsclienttrait.md#method_waitUntil)

[Back To Top](class-aws-glue-glueclient-top.md)
