# Data Types

The Amazon CloudWatch Logs API contains several data types that various actions use. This section describes each data type in detail.

###### Note

The order of each element in a data type structure is not guaranteed. Applications should not assume a particular order.

The following data types are supported:

- [AccountPolicy](api-accountpolicy.md)

- [AddKeyEntry](api-addkeyentry.md)

- [AddKeys](api-addkeys.md)

- [AggregateLogGroupSummary](api-aggregateloggroupsummary.md)

- [Anomaly](api-anomaly.md)

- [AnomalyDetector](api-anomalydetector.md)

- [ConfigurationTemplate](api-configurationtemplate.md)

- [ConfigurationTemplateDeliveryConfigValues](api-configurationtemplatedeliveryconfigvalues.md)

- [CopyValue](api-copyvalue.md)

- [CopyValueEntry](api-copyvalueentry.md)

- [CSV](api-csv.md)

- [DataSource](api-datasource.md)

- [DataSourceFilter](api-datasourcefilter.md)

- [DateTimeConverter](api-datetimeconverter.md)

- [DeleteKeys](api-deletekeys.md)

- [Delivery](api-delivery.md)

- [DeliveryDestination](api-deliverydestination.md)

- [DeliveryDestinationConfiguration](api-deliverydestinationconfiguration.md)

- [DeliverySource](api-deliverysource.md)

- [Destination](api-destination.md)

- [DestinationConfiguration](api-destinationconfiguration.md)

- [Entity](api-entity.md)

- [ExportTask](api-exporttask.md)

- [ExportTaskExecutionInfo](api-exporttaskexecutioninfo.md)

- [ExportTaskStatus](api-exporttaskstatus.md)

- [FieldIndex](api-fieldindex.md)

- [FieldsData](api-fieldsdata.md)

- [FilteredLogEvent](api-filteredlogevent.md)

- [GetLogObjectResponseStream](api-getlogobjectresponsestream.md)

- [Grok](api-grok.md)

- [GroupingIdentifier](api-groupingidentifier.md)

- [Import](api-import.md)

- [ImportBatch](api-importbatch.md)

- [ImportFilter](api-importfilter.md)

- [ImportStatistics](api-importstatistics.md)

- [IndexPolicy](api-indexpolicy.md)

- [InputLogEvent](api-inputlogevent.md)

- [IntegrationDetails](api-integrationdetails.md)

- [IntegrationSummary](api-integrationsummary.md)

- [ListToMap](api-listtomap.md)

- [LiveTailSessionLogEvent](api-livetailsessionlogevent.md)

- [LiveTailSessionMetadata](api-livetailsessionmetadata.md)

- [LiveTailSessionStart](api-livetailsessionstart.md)

- [LiveTailSessionUpdate](api-livetailsessionupdate.md)

- [LogEvent](api-logevent.md)

- [LogFieldsListItem](api-logfieldslistitem.md)

- [LogFieldType](api-logfieldtype.md)

- [LogGroup](api-loggroup.md)

- [LogGroupField](api-loggroupfield.md)

- [LogGroupSummary](api-loggroupsummary.md)

- [LogStream](api-logstream.md)

- [LookupTable](api-lookuptable.md)

- [LowerCaseString](api-lowercasestring.md)

- [MetricFilter](api-metricfilter.md)

- [MetricFilterMatchRecord](api-metricfiltermatchrecord.md)

- [MetricTransformation](api-metrictransformation.md)

- [MoveKeyEntry](api-movekeyentry.md)

- [MoveKeys](api-movekeys.md)

- [OpenSearchApplication](api-opensearchapplication.md)

- [OpenSearchCollection](api-opensearchcollection.md)

- [OpenSearchDataAccessPolicy](api-opensearchdataaccesspolicy.md)

- [OpenSearchDataSource](api-opensearchdatasource.md)

- [OpenSearchEncryptionPolicy](api-opensearchencryptionpolicy.md)

- [OpenSearchIntegrationDetails](api-opensearchintegrationdetails.md)

- [OpenSearchLifecyclePolicy](api-opensearchlifecyclepolicy.md)

- [OpenSearchNetworkPolicy](api-opensearchnetworkpolicy.md)

- [OpenSearchResourceConfig](api-opensearchresourceconfig.md)

- [OpenSearchResourceStatus](api-opensearchresourcestatus.md)

- [OpenSearchWorkspace](api-opensearchworkspace.md)

- [OutputLogEvent](api-outputlogevent.md)

- [ParseCloudfront](api-parsecloudfront.md)

- [ParseJSON](api-parsejson.md)

- [ParseKeyValue](api-parsekeyvalue.md)

- [ParsePostgres](api-parsepostgres.md)

- [ParseRoute53](api-parseroute53.md)

- [ParseToOCSF](api-parsetoocsf.md)

- [ParseVPC](api-parsevpc.md)

- [ParseWAF](api-parsewaf.md)

- [PatternToken](api-patterntoken.md)

- [Policy](api-policy.md)

- [Processor](api-processor.md)

- [QueryCompileError](api-querycompileerror.md)

- [QueryCompileErrorLocation](api-querycompileerrorlocation.md)

- [QueryDefinition](api-querydefinition.md)

- [QueryInfo](api-queryinfo.md)

- [QueryParameter](api-queryparameter.md)

- [QueryStatistics](api-querystatistics.md)

- [RecordField](api-recordfield.md)

- [RejectedEntityInfo](api-rejectedentityinfo.md)

- [RejectedLogEventsInfo](api-rejectedlogeventsinfo.md)

- [RenameKeyEntry](api-renamekeyentry.md)

- [RenameKeys](api-renamekeys.md)

- [ResourceConfig](api-resourceconfig.md)

- [ResourcePolicy](api-resourcepolicy.md)

- [ResultField](api-resultfield.md)

- [S3Configuration](api-s3configuration.md)

- [S3DeliveryConfiguration](api-s3deliveryconfiguration.md)

- [S3TableIntegrationSource](api-s3tableintegrationsource.md)

- [ScheduledQueryDestination](api-scheduledquerydestination.md)

- [ScheduledQuerySummary](api-scheduledquerysummary.md)

- [SearchedLogStream](api-searchedlogstream.md)

- [SplitString](api-splitstring.md)

- [SplitStringEntry](api-splitstringentry.md)

- [StartLiveTailResponseStream](api-startlivetailresponsestream.md)

- [SubscriptionFilter](api-subscriptionfilter.md)

- [SubstituteString](api-substitutestring.md)

- [SubstituteStringEntry](api-substitutestringentry.md)

- [SuppressionPeriod](api-suppressionperiod.md)

- [TransformedLogRecord](api-transformedlogrecord.md)

- [TriggerHistoryRecord](api-triggerhistoryrecord.md)

- [TrimString](api-trimstring.md)

- [TypeConverter](api-typeconverter.md)

- [TypeConverterEntry](api-typeconverterentry.md)

- [UpperCaseString](api-uppercasestring.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UpdateScheduledQuery

AccountPolicy

All content copied from https://docs.aws.amazon.com/.
