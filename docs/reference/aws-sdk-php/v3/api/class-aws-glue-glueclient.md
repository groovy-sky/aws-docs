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

- [**2017-03-31**](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html)

  - [BatchCreatePartition](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#batchcreatepartition)
  - [BatchDeleteConnection](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#batchdeleteconnection)
  - [BatchDeletePartition](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#batchdeletepartition)
  - [BatchDeleteTable](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#batchdeletetable)
  - [BatchDeleteTableVersion](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#batchdeletetableversion)
  - [BatchGetBlueprints](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#batchgetblueprints)
  - [BatchGetCrawlers](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#batchgetcrawlers)
  - [BatchGetCustomEntityTypes](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#batchgetcustomentitytypes)
  - [BatchGetDataQualityResult](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#batchgetdataqualityresult)
  - [BatchGetDevEndpoints](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#batchgetdevendpoints)
  - [BatchGetJobs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#batchgetjobs)
  - [BatchGetPartition](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#batchgetpartition)
  - [BatchGetTableOptimizer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#batchgettableoptimizer)
  - [BatchGetTriggers](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#batchgettriggers)
  - [BatchGetWorkflows](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#batchgetworkflows)
  - [BatchPutDataQualityStatisticAnnotation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#batchputdataqualitystatisticannotation)
  - [BatchStopJobRun](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#batchstopjobrun)
  - [BatchUpdatePartition](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#batchupdatepartition)
  - [CancelDataQualityRuleRecommendationRun](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#canceldataqualityrulerecommendationrun)
  - [CancelDataQualityRulesetEvaluationRun](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#canceldataqualityrulesetevaluationrun)
  - [CancelMLTaskRun](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#cancelmltaskrun)
  - [CancelStatement](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#cancelstatement)
  - [CheckSchemaVersionValidity](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#checkschemaversionvalidity)
  - [CreateBlueprint](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#createblueprint)
  - [CreateCatalog](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#createcatalog)
  - [CreateClassifier](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#createclassifier)
  - [CreateColumnStatisticsTaskSettings](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#createcolumnstatisticstasksettings)
  - [CreateConnection](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#createconnection)
  - [CreateCrawler](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#createcrawler)
  - [CreateCustomEntityType](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#createcustomentitytype)
  - [CreateDataQualityRuleset](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#createdataqualityruleset)
  - [CreateDatabase](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#createdatabase)
  - [CreateDevEndpoint](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#createdevendpoint)
  - [CreateGlueIdentityCenterConfiguration](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#createglueidentitycenterconfiguration)
  - [CreateIntegration](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#createintegration)
  - [CreateIntegrationResourceProperty](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#createintegrationresourceproperty)
  - [CreateIntegrationTableProperties](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#createintegrationtableproperties)
  - [CreateJob](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#createjob)
  - [CreateMLTransform](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#createmltransform)
  - [CreatePartition](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#createpartition)
  - [CreatePartitionIndex](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#createpartitionindex)
  - [CreateRegistry](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#createregistry)
  - [CreateSchema](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#createschema)
  - [CreateScript](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#createscript)
  - [CreateSecurityConfiguration](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#createsecurityconfiguration)
  - [CreateSession](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#createsession)
  - [CreateTable](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#createtable)
  - [CreateTableOptimizer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#createtableoptimizer)
  - [CreateTrigger](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#createtrigger)
  - [CreateUsageProfile](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#createusageprofile)
  - [CreateUserDefinedFunction](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#createuserdefinedfunction)
  - [CreateWorkflow](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#createworkflow)
  - [DeleteBlueprint](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#deleteblueprint)
  - [DeleteCatalog](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#deletecatalog)
  - [DeleteClassifier](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#deleteclassifier)
  - [DeleteColumnStatisticsForPartition](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#deletecolumnstatisticsforpartition)
  - [DeleteColumnStatisticsForTable](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#deletecolumnstatisticsfortable)
  - [DeleteColumnStatisticsTaskSettings](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#deletecolumnstatisticstasksettings)
  - [DeleteConnection](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#deleteconnection)
  - [DeleteConnectionType](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#deleteconnectiontype)
  - [DeleteCrawler](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#deletecrawler)
  - [DeleteCustomEntityType](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#deletecustomentitytype)
  - [DeleteDataQualityRuleset](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#deletedataqualityruleset)
  - [DeleteDatabase](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#deletedatabase)
  - [DeleteDevEndpoint](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#deletedevendpoint)
  - [DeleteGlueIdentityCenterConfiguration](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#deleteglueidentitycenterconfiguration)
  - [DeleteIntegration](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#deleteintegration)
  - [DeleteIntegrationResourceProperty](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#deleteintegrationresourceproperty)
  - [DeleteIntegrationTableProperties](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#deleteintegrationtableproperties)
  - [DeleteJob](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#deletejob)
  - [DeleteMLTransform](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#deletemltransform)
  - [DeletePartition](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#deletepartition)
  - [DeletePartitionIndex](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#deletepartitionindex)
  - [DeleteRegistry](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#deleteregistry)
  - [DeleteResourcePolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#deleteresourcepolicy)
  - [DeleteSchema](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#deleteschema)
  - [DeleteSchemaVersions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#deleteschemaversions)
  - [DeleteSecurityConfiguration](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#deletesecurityconfiguration)
  - [DeleteSession](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#deletesession)
  - [DeleteTable](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#deletetable)
  - [DeleteTableOptimizer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#deletetableoptimizer)
  - [DeleteTableVersion](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#deletetableversion)
  - [DeleteTrigger](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#deletetrigger)
  - [DeleteUsageProfile](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#deleteusageprofile)
  - [DeleteUserDefinedFunction](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#deleteuserdefinedfunction)
  - [DeleteWorkflow](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#deleteworkflow)
  - [DescribeConnectionType](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#describeconnectiontype)
  - [DescribeEntity](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#describeentity)
  - [DescribeInboundIntegrations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#describeinboundintegrations)
  - [DescribeIntegrations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#describeintegrations)
  - [GetBlueprint](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getblueprint)
  - [GetBlueprintRun](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getblueprintrun)
  - [GetBlueprintRuns](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getblueprintruns)
  - [GetCatalog](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getcatalog)
  - [GetCatalogImportStatus](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getcatalogimportstatus)
  - [GetCatalogs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getcatalogs)
  - [GetClassifier](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getclassifier)
  - [GetClassifiers](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getclassifiers)
  - [GetColumnStatisticsForPartition](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getcolumnstatisticsforpartition)
  - [GetColumnStatisticsForTable](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getcolumnstatisticsfortable)
  - [GetColumnStatisticsTaskRun](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getcolumnstatisticstaskrun)
  - [GetColumnStatisticsTaskRuns](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getcolumnstatisticstaskruns)
  - [GetColumnStatisticsTaskSettings](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getcolumnstatisticstasksettings)
  - [GetConnection](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getconnection)
  - [GetConnections](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getconnections)
  - [GetCrawler](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getcrawler)
  - [GetCrawlerMetrics](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getcrawlermetrics)
  - [GetCrawlers](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getcrawlers)
  - [GetCustomEntityType](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getcustomentitytype)
  - [GetDataCatalogEncryptionSettings](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getdatacatalogencryptionsettings)
  - [GetDataQualityModel](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getdataqualitymodel)
  - [GetDataQualityModelResult](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getdataqualitymodelresult)
  - [GetDataQualityResult](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getdataqualityresult)
  - [GetDataQualityRuleRecommendationRun](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getdataqualityrulerecommendationrun)
  - [GetDataQualityRuleset](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getdataqualityruleset)
  - [GetDataQualityRulesetEvaluationRun](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getdataqualityrulesetevaluationrun)
  - [GetDatabase](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getdatabase)
  - [GetDatabases](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getdatabases)
  - [GetDataflowGraph](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getdataflowgraph)
  - [GetDevEndpoint](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getdevendpoint)
  - [GetDevEndpoints](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getdevendpoints)
  - [GetEntityRecords](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getentityrecords)
  - [GetGlueIdentityCenterConfiguration](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getglueidentitycenterconfiguration)
  - [GetIntegrationResourceProperty](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getintegrationresourceproperty)
  - [GetIntegrationTableProperties](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getintegrationtableproperties)
  - [GetJob](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getjob)
  - [GetJobBookmark](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getjobbookmark)
  - [GetJobRun](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getjobrun)
  - [GetJobRuns](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getjobruns)
  - [GetJobs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getjobs)
  - [GetMLTaskRun](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getmltaskrun)
  - [GetMLTaskRuns](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getmltaskruns)
  - [GetMLTransform](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getmltransform)
  - [GetMLTransforms](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getmltransforms)
  - [GetMapping](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getmapping)
  - [GetMaterializedViewRefreshTaskRun](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getmaterializedviewrefreshtaskrun)
  - [GetPartition](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getpartition)
  - [GetPartitionIndexes](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getpartitionindexes)
  - [GetPartitions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getpartitions)
  - [GetPlan](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getplan)
  - [GetRegistry](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getregistry)
  - [GetResourcePolicies](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getresourcepolicies)
  - [GetResourcePolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getresourcepolicy)
  - [GetSchema](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getschema)
  - [GetSchemaByDefinition](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getschemabydefinition)
  - [GetSchemaVersion](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getschemaversion)
  - [GetSchemaVersionsDiff](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getschemaversionsdiff)
  - [GetSecurityConfiguration](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getsecurityconfiguration)
  - [GetSecurityConfigurations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getsecurityconfigurations)
  - [GetSession](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getsession)
  - [GetStatement](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getstatement)
  - [GetTable](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#gettable)
  - [GetTableOptimizer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#gettableoptimizer)
  - [GetTableVersion](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#gettableversion)
  - [GetTableVersions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#gettableversions)
  - [GetTables](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#gettables)
  - [GetTags](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#gettags)
  - [GetTrigger](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#gettrigger)
  - [GetTriggers](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#gettriggers)
  - [GetUnfilteredPartitionMetadata](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getunfilteredpartitionmetadata)
  - [GetUnfilteredPartitionsMetadata](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getunfilteredpartitionsmetadata)
  - [GetUnfilteredTableMetadata](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getunfilteredtablemetadata)
  - [GetUsageProfile](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getusageprofile)
  - [GetUserDefinedFunction](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getuserdefinedfunction)
  - [GetUserDefinedFunctions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getuserdefinedfunctions)
  - [GetWorkflow](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getworkflow)
  - [GetWorkflowRun](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getworkflowrun)
  - [GetWorkflowRunProperties](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getworkflowrunproperties)
  - [GetWorkflowRuns](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#getworkflowruns)
  - [ImportCatalogToGlue](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#importcatalogtoglue)
  - [ListBlueprints](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#listblueprints)
  - [ListColumnStatisticsTaskRuns](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#listcolumnstatisticstaskruns)
  - [ListConnectionTypes](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#listconnectiontypes)
  - [ListCrawlers](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#listcrawlers)
  - [ListCrawls](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#listcrawls)
  - [ListCustomEntityTypes](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#listcustomentitytypes)
  - [ListDataQualityResults](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#listdataqualityresults)
  - [ListDataQualityRuleRecommendationRuns](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#listdataqualityrulerecommendationruns)
  - [ListDataQualityRulesetEvaluationRuns](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#listdataqualityrulesetevaluationruns)
  - [ListDataQualityRulesets](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#listdataqualityrulesets)
  - [ListDataQualityStatisticAnnotations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#listdataqualitystatisticannotations)
  - [ListDataQualityStatistics](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#listdataqualitystatistics)
  - [ListDevEndpoints](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#listdevendpoints)
  - [ListEntities](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#listentities)
  - [ListIntegrationResourceProperties](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#listintegrationresourceproperties)
  - [ListJobs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#listjobs)
  - [ListMLTransforms](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#listmltransforms)
  - [ListMaterializedViewRefreshTaskRuns](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#listmaterializedviewrefreshtaskruns)
  - [ListRegistries](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#listregistries)
  - [ListSchemaVersions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#listschemaversions)
  - [ListSchemas](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#listschemas)
  - [ListSessions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#listsessions)
  - [ListStatements](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#liststatements)
  - [ListTableOptimizerRuns](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#listtableoptimizerruns)
  - [ListTriggers](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#listtriggers)
  - [ListUsageProfiles](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#listusageprofiles)
  - [ListWorkflows](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#listworkflows)
  - [ModifyIntegration](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#modifyintegration)
  - [PutDataCatalogEncryptionSettings](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#putdatacatalogencryptionsettings)
  - [PutDataQualityProfileAnnotation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#putdataqualityprofileannotation)
  - [PutResourcePolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#putresourcepolicy)
  - [PutSchemaVersionMetadata](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#putschemaversionmetadata)
  - [PutWorkflowRunProperties](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#putworkflowrunproperties)
  - [QuerySchemaVersionMetadata](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#queryschemaversionmetadata)
  - [RegisterConnectionType](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#registerconnectiontype)
  - [RegisterSchemaVersion](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#registerschemaversion)
  - [RemoveSchemaVersionMetadata](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#removeschemaversionmetadata)
  - [ResetJobBookmark](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#resetjobbookmark)
  - [ResumeWorkflowRun](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#resumeworkflowrun)
  - [RunStatement](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#runstatement)
  - [SearchTables](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#searchtables)
  - [StartBlueprintRun](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#startblueprintrun)
  - [StartColumnStatisticsTaskRun](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#startcolumnstatisticstaskrun)
  - [StartColumnStatisticsTaskRunSchedule](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#startcolumnstatisticstaskrunschedule)
  - [StartCrawler](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#startcrawler)
  - [StartCrawlerSchedule](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#startcrawlerschedule)
  - [StartDataQualityRuleRecommendationRun](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#startdataqualityrulerecommendationrun)
  - [StartDataQualityRulesetEvaluationRun](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#startdataqualityrulesetevaluationrun)
  - [StartExportLabelsTaskRun](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#startexportlabelstaskrun)
  - [StartImportLabelsTaskRun](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#startimportlabelstaskrun)
  - [StartJobRun](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#startjobrun)
  - [StartMLEvaluationTaskRun](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#startmlevaluationtaskrun)
  - [StartMLLabelingSetGenerationTaskRun](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#startmllabelingsetgenerationtaskrun)
  - [StartMaterializedViewRefreshTaskRun](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#startmaterializedviewrefreshtaskrun)
  - [StartTrigger](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#starttrigger)
  - [StartWorkflowRun](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#startworkflowrun)
  - [StopColumnStatisticsTaskRun](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#stopcolumnstatisticstaskrun)
  - [StopColumnStatisticsTaskRunSchedule](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#stopcolumnstatisticstaskrunschedule)
  - [StopCrawler](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#stopcrawler)
  - [StopCrawlerSchedule](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#stopcrawlerschedule)
  - [StopMaterializedViewRefreshTaskRun](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#stopmaterializedviewrefreshtaskrun)
  - [StopSession](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#stopsession)
  - [StopTrigger](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#stoptrigger)
  - [StopWorkflowRun](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#stopworkflowrun)
  - [TagResource](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#tagresource)
  - [TestConnection](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#testconnection)
  - [UntagResource](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#untagresource)
  - [UpdateBlueprint](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#updateblueprint)
  - [UpdateCatalog](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#updatecatalog)
  - [UpdateClassifier](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#updateclassifier)
  - [UpdateColumnStatisticsForPartition](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#updatecolumnstatisticsforpartition)
  - [UpdateColumnStatisticsForTable](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#updatecolumnstatisticsfortable)
  - [UpdateColumnStatisticsTaskSettings](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#updatecolumnstatisticstasksettings)
  - [UpdateConnection](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#updateconnection)
  - [UpdateCrawler](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#updatecrawler)
  - [UpdateCrawlerSchedule](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#updatecrawlerschedule)
  - [UpdateDataQualityRuleset](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#updatedataqualityruleset)
  - [UpdateDatabase](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#updatedatabase)
  - [UpdateDevEndpoint](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#updatedevendpoint)
  - [UpdateGlueIdentityCenterConfiguration](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#updateglueidentitycenterconfiguration)
  - [UpdateIntegrationResourceProperty](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#updateintegrationresourceproperty)
  - [UpdateIntegrationTableProperties](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#updateintegrationtableproperties)
  - [UpdateJob](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#updatejob)
  - [UpdateJobFromSourceControl](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#updatejobfromsourcecontrol)
  - [UpdateMLTransform](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#updatemltransform)
  - [UpdatePartition](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#updatepartition)
  - [UpdateRegistry](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#updateregistry)
  - [UpdateSchema](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#updateschema)
  - [UpdateSourceControlFromJob](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#updatesourcecontrolfromjob)
  - [UpdateTable](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#updatetable)
  - [UpdateTableOptimizer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#updatetableoptimizer)
  - [UpdateTrigger](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#updatetrigger)
  - [UpdateUsageProfile](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#updateusageprofile)
  - [UpdateUserDefinedFunction](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#updateuserdefinedfunction)
  - [UpdateWorkflow](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-glue-2017-03-31.html#updateworkflow)

## Examples

### Basics, Actions and Scenarios

The following code examples show you how to perform actions and implement common scenarios by using the AWS SDK for PHP with AWS Glue.

- [See examples on AWS Docs](https://docs.aws.amazon.com/sdk-for-php/v3/developer-guide/php_glue_code_examples.html)

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Glue.GlueClient.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Glue.GlueClient.html\#toc-methods)

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

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Glue.GlueClient.html\#methods)

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
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Glue.GlueClient.html#toc-methods)
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

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Glue.GlueClient.html#top)
