Menu

- [Aws](namespace-aws.md)
- [QuickSight](namespace-aws-quicksight.md)

## QuickSightClient     extends [AwsClient](class-aws-awsclient.md)   in package    - [Aws](package-aws.md)

This client is used to interact with the **Amazon QuickSight** service.

## Supported API Versions

This class uses a _service description model_ that is associated at
runtime based on the `version` option given when constructing the
client. The `version` option will determine which API operations,
waiters, and paginators are available for a client. Creating a command or a
specific API operation can be done using magic methods (e.g.,
`$client->commandName(/** parameters */)`, or using the
`$client->getCommand` method of the client.

- [**2018-04-01**](api-quicksight-2018-04-01.md)

  - [BatchCreateTopicReviewedAnswer](api-quicksight-2018-04-01-batchcreatetopicreviewedanswer.md)
  - [BatchDeleteTopicReviewedAnswer](api-quicksight-2018-04-01-batchdeletetopicreviewedanswer.md)
  - [CancelIngestion](api-quicksight-2018-04-01-cancelingestion.md)
  - [CreateAccountCustomization](api-quicksight-2018-04-01-createaccountcustomization.md)
  - [CreateAccountSubscription](api-quicksight-2018-04-01-createaccountsubscription.md)
  - [CreateActionConnector](api-quicksight-2018-04-01-createactionconnector.md)
  - [CreateAnalysis](api-quicksight-2018-04-01-createanalysis.md)
  - [CreateBrand](api-quicksight-2018-04-01-createbrand.md)
  - [CreateCustomPermissions](api-quicksight-2018-04-01-createcustompermissions.md)
  - [CreateDashboard](api-quicksight-2018-04-01-createdashboard.md)
  - [CreateDataSet](api-quicksight-2018-04-01-createdataset.md)
  - [CreateDataSource](api-quicksight-2018-04-01-createdatasource.md)
  - [CreateFolder](api-quicksight-2018-04-01-createfolder.md)
  - [CreateFolderMembership](api-quicksight-2018-04-01-createfoldermembership.md)
  - [CreateGroup](api-quicksight-2018-04-01-creategroup.md)
  - [CreateGroupMembership](api-quicksight-2018-04-01-creategroupmembership.md)
  - [CreateIAMPolicyAssignment](api-quicksight-2018-04-01-createiampolicyassignment.md)
  - [CreateIngestion](api-quicksight-2018-04-01-createingestion.md)
  - [CreateNamespace](api-quicksight-2018-04-01-createnamespace.md)
  - [CreateRefreshSchedule](api-quicksight-2018-04-01-createrefreshschedule.md)
  - [CreateRoleMembership](api-quicksight-2018-04-01-createrolemembership.md)
  - [CreateTemplate](api-quicksight-2018-04-01-createtemplate.md)
  - [CreateTemplateAlias](api-quicksight-2018-04-01-createtemplatealias.md)
  - [CreateTheme](api-quicksight-2018-04-01-createtheme.md)
  - [CreateThemeAlias](api-quicksight-2018-04-01-createthemealias.md)
  - [CreateTopic](api-quicksight-2018-04-01-createtopic.md)
  - [CreateTopicRefreshSchedule](api-quicksight-2018-04-01-createtopicrefreshschedule.md)
  - [CreateVPCConnection](api-quicksight-2018-04-01-createvpcconnection.md)
  - [DeleteAccountCustomPermission](api-quicksight-2018-04-01-deleteaccountcustompermission.md)
  - [DeleteAccountCustomization](api-quicksight-2018-04-01-deleteaccountcustomization.md)
  - [DeleteAccountSubscription](api-quicksight-2018-04-01-deleteaccountsubscription.md)
  - [DeleteActionConnector](api-quicksight-2018-04-01-deleteactionconnector.md)
  - [DeleteAnalysis](api-quicksight-2018-04-01-deleteanalysis.md)
  - [DeleteBrand](api-quicksight-2018-04-01-deletebrand.md)
  - [DeleteBrandAssignment](api-quicksight-2018-04-01-deletebrandassignment.md)
  - [DeleteCustomPermissions](api-quicksight-2018-04-01-deletecustompermissions.md)
  - [DeleteDashboard](api-quicksight-2018-04-01-deletedashboard.md)
  - [DeleteDataSet](api-quicksight-2018-04-01-deletedataset.md)
  - [DeleteDataSetRefreshProperties](api-quicksight-2018-04-01-deletedatasetrefreshproperties.md)
  - [DeleteDataSource](api-quicksight-2018-04-01-deletedatasource.md)
  - [DeleteDefaultQBusinessApplication](api-quicksight-2018-04-01-deletedefaultqbusinessapplication.md)
  - [DeleteFolder](api-quicksight-2018-04-01-deletefolder.md)
  - [DeleteFolderMembership](api-quicksight-2018-04-01-deletefoldermembership.md)
  - [DeleteGroup](api-quicksight-2018-04-01-deletegroup.md)
  - [DeleteGroupMembership](api-quicksight-2018-04-01-deletegroupmembership.md)
  - [DeleteIAMPolicyAssignment](api-quicksight-2018-04-01-deleteiampolicyassignment.md)
  - [DeleteIdentityPropagationConfig](api-quicksight-2018-04-01-deleteidentitypropagationconfig.md)
  - [DeleteNamespace](api-quicksight-2018-04-01-deletenamespace.md)
  - [DeleteRefreshSchedule](api-quicksight-2018-04-01-deleterefreshschedule.md)
  - [DeleteRoleCustomPermission](api-quicksight-2018-04-01-deleterolecustompermission.md)
  - [DeleteRoleMembership](api-quicksight-2018-04-01-deleterolemembership.md)
  - [DeleteTemplate](api-quicksight-2018-04-01-deletetemplate.md)
  - [DeleteTemplateAlias](api-quicksight-2018-04-01-deletetemplatealias.md)
  - [DeleteTheme](api-quicksight-2018-04-01-deletetheme.md)
  - [DeleteThemeAlias](api-quicksight-2018-04-01-deletethemealias.md)
  - [DeleteTopic](api-quicksight-2018-04-01-deletetopic.md)
  - [DeleteTopicRefreshSchedule](api-quicksight-2018-04-01-deletetopicrefreshschedule.md)
  - [DeleteUser](api-quicksight-2018-04-01-deleteuser.md)
  - [DeleteUserByPrincipalId](api-quicksight-2018-04-01-deleteuserbyprincipalid.md)
  - [DeleteUserCustomPermission](api-quicksight-2018-04-01-deleteusercustompermission.md)
  - [DeleteVPCConnection](api-quicksight-2018-04-01-deletevpcconnection.md)
  - [DescribeAccountCustomPermission](api-quicksight-2018-04-01-describeaccountcustompermission.md)
  - [DescribeAccountCustomization](api-quicksight-2018-04-01-describeaccountcustomization.md)
  - [DescribeAccountSettings](api-quicksight-2018-04-01-describeaccountsettings.md)
  - [DescribeAccountSubscription](api-quicksight-2018-04-01-describeaccountsubscription.md)
  - [DescribeActionConnector](api-quicksight-2018-04-01-describeactionconnector.md)
  - [DescribeActionConnectorPermissions](api-quicksight-2018-04-01-describeactionconnectorpermissions.md)
  - [DescribeAnalysis](api-quicksight-2018-04-01-describeanalysis.md)
  - [DescribeAnalysisDefinition](api-quicksight-2018-04-01-describeanalysisdefinition.md)
  - [DescribeAnalysisPermissions](api-quicksight-2018-04-01-describeanalysispermissions.md)
  - [DescribeAssetBundleExportJob](api-quicksight-2018-04-01-describeassetbundleexportjob.md)
  - [DescribeAssetBundleImportJob](api-quicksight-2018-04-01-describeassetbundleimportjob.md)
  - [DescribeAutomationJob](api-quicksight-2018-04-01-describeautomationjob.md)
  - [DescribeBrand](api-quicksight-2018-04-01-describebrand.md)
  - [DescribeBrandAssignment](api-quicksight-2018-04-01-describebrandassignment.md)
  - [DescribeBrandPublishedVersion](api-quicksight-2018-04-01-describebrandpublishedversion.md)
  - [DescribeCustomPermissions](api-quicksight-2018-04-01-describecustompermissions.md)
  - [DescribeDashboard](api-quicksight-2018-04-01-describedashboard.md)
  - [DescribeDashboardDefinition](api-quicksight-2018-04-01-describedashboarddefinition.md)
  - [DescribeDashboardPermissions](api-quicksight-2018-04-01-describedashboardpermissions.md)
  - [DescribeDashboardSnapshotJob](api-quicksight-2018-04-01-describedashboardsnapshotjob.md)
  - [DescribeDashboardSnapshotJobResult](api-quicksight-2018-04-01-describedashboardsnapshotjobresult.md)
  - [DescribeDashboardsQAConfiguration](api-quicksight-2018-04-01-describedashboardsqaconfiguration.md)
  - [DescribeDataSet](api-quicksight-2018-04-01-describedataset.md)
  - [DescribeDataSetPermissions](api-quicksight-2018-04-01-describedatasetpermissions.md)
  - [DescribeDataSetRefreshProperties](api-quicksight-2018-04-01-describedatasetrefreshproperties.md)
  - [DescribeDataSource](api-quicksight-2018-04-01-describedatasource.md)
  - [DescribeDataSourcePermissions](api-quicksight-2018-04-01-describedatasourcepermissions.md)
  - [DescribeDefaultQBusinessApplication](api-quicksight-2018-04-01-describedefaultqbusinessapplication.md)
  - [DescribeFolder](api-quicksight-2018-04-01-describefolder.md)
  - [DescribeFolderPermissions](api-quicksight-2018-04-01-describefolderpermissions.md)
  - [DescribeFolderResolvedPermissions](api-quicksight-2018-04-01-describefolderresolvedpermissions.md)
  - [DescribeGroup](api-quicksight-2018-04-01-describegroup.md)
  - [DescribeGroupMembership](api-quicksight-2018-04-01-describegroupmembership.md)
  - [DescribeIAMPolicyAssignment](api-quicksight-2018-04-01-describeiampolicyassignment.md)
  - [DescribeIngestion](api-quicksight-2018-04-01-describeingestion.md)
  - [DescribeIpRestriction](api-quicksight-2018-04-01-describeiprestriction.md)
  - [DescribeKeyRegistration](api-quicksight-2018-04-01-describekeyregistration.md)
  - [DescribeNamespace](api-quicksight-2018-04-01-describenamespace.md)
  - [DescribeQPersonalizationConfiguration](api-quicksight-2018-04-01-describeqpersonalizationconfiguration.md)
  - [DescribeQuickSightQSearchConfiguration](api-quicksight-2018-04-01-describequicksightqsearchconfiguration.md)
  - [DescribeRefreshSchedule](api-quicksight-2018-04-01-describerefreshschedule.md)
  - [DescribeRoleCustomPermission](api-quicksight-2018-04-01-describerolecustompermission.md)
  - [DescribeSelfUpgradeConfiguration](api-quicksight-2018-04-01-describeselfupgradeconfiguration.md)
  - [DescribeTemplate](api-quicksight-2018-04-01-describetemplate.md)
  - [DescribeTemplateAlias](api-quicksight-2018-04-01-describetemplatealias.md)
  - [DescribeTemplateDefinition](api-quicksight-2018-04-01-describetemplatedefinition.md)
  - [DescribeTemplatePermissions](api-quicksight-2018-04-01-describetemplatepermissions.md)
  - [DescribeTheme](api-quicksight-2018-04-01-describetheme.md)
  - [DescribeThemeAlias](api-quicksight-2018-04-01-describethemealias.md)
  - [DescribeThemePermissions](api-quicksight-2018-04-01-describethemepermissions.md)
  - [DescribeTopic](api-quicksight-2018-04-01-describetopic.md)
  - [DescribeTopicPermissions](api-quicksight-2018-04-01-describetopicpermissions.md)
  - [DescribeTopicRefresh](api-quicksight-2018-04-01-describetopicrefresh.md)
  - [DescribeTopicRefreshSchedule](api-quicksight-2018-04-01-describetopicrefreshschedule.md)
  - [DescribeUser](api-quicksight-2018-04-01-describeuser.md)
  - [DescribeVPCConnection](api-quicksight-2018-04-01-describevpcconnection.md)
  - [GenerateEmbedUrlForAnonymousUser](api-quicksight-2018-04-01-generateembedurlforanonymoususer.md)
  - [GenerateEmbedUrlForRegisteredUser](api-quicksight-2018-04-01-generateembedurlforregistereduser.md)
  - [GenerateEmbedUrlForRegisteredUserWithIdentity](api-quicksight-2018-04-01-generateembedurlforregistereduserwithidentity.md)
  - [GetDashboardEmbedUrl](api-quicksight-2018-04-01-getdashboardembedurl.md)
  - [GetFlowMetadata](api-quicksight-2018-04-01-getflowmetadata.md)
  - [GetFlowPermissions](api-quicksight-2018-04-01-getflowpermissions.md)
  - [GetIdentityContext](api-quicksight-2018-04-01-getidentitycontext.md)
  - [GetSessionEmbedUrl](api-quicksight-2018-04-01-getsessionembedurl.md)
  - [ListActionConnectors](api-quicksight-2018-04-01-listactionconnectors.md)
  - [ListAnalyses](api-quicksight-2018-04-01-listanalyses.md)
  - [ListAssetBundleExportJobs](api-quicksight-2018-04-01-listassetbundleexportjobs.md)
  - [ListAssetBundleImportJobs](api-quicksight-2018-04-01-listassetbundleimportjobs.md)
  - [ListBrands](api-quicksight-2018-04-01-listbrands.md)
  - [ListCustomPermissions](api-quicksight-2018-04-01-listcustompermissions.md)
  - [ListDashboardVersions](api-quicksight-2018-04-01-listdashboardversions.md)
  - [ListDashboards](api-quicksight-2018-04-01-listdashboards.md)
  - [ListDataSets](api-quicksight-2018-04-01-listdatasets.md)
  - [ListDataSources](api-quicksight-2018-04-01-listdatasources.md)
  - [ListFlows](api-quicksight-2018-04-01-listflows.md)
  - [ListFolderMembers](api-quicksight-2018-04-01-listfoldermembers.md)
  - [ListFolders](api-quicksight-2018-04-01-listfolders.md)
  - [ListFoldersForResource](api-quicksight-2018-04-01-listfoldersforresource.md)
  - [ListGroupMemberships](api-quicksight-2018-04-01-listgroupmemberships.md)
  - [ListGroups](api-quicksight-2018-04-01-listgroups.md)
  - [ListIAMPolicyAssignments](api-quicksight-2018-04-01-listiampolicyassignments.md)
  - [ListIAMPolicyAssignmentsForUser](api-quicksight-2018-04-01-listiampolicyassignmentsforuser.md)
  - [ListIdentityPropagationConfigs](api-quicksight-2018-04-01-listidentitypropagationconfigs.md)
  - [ListIngestions](api-quicksight-2018-04-01-listingestions.md)
  - [ListNamespaces](api-quicksight-2018-04-01-listnamespaces.md)
  - [ListRefreshSchedules](api-quicksight-2018-04-01-listrefreshschedules.md)
  - [ListRoleMemberships](api-quicksight-2018-04-01-listrolememberships.md)
  - [ListSelfUpgrades](api-quicksight-2018-04-01-listselfupgrades.md)
  - [ListTagsForResource](api-quicksight-2018-04-01-listtagsforresource.md)
  - [ListTemplateAliases](api-quicksight-2018-04-01-listtemplatealiases.md)
  - [ListTemplateVersions](api-quicksight-2018-04-01-listtemplateversions.md)
  - [ListTemplates](api-quicksight-2018-04-01-listtemplates.md)
  - [ListThemeAliases](api-quicksight-2018-04-01-listthemealiases.md)
  - [ListThemeVersions](api-quicksight-2018-04-01-listthemeversions.md)
  - [ListThemes](api-quicksight-2018-04-01-listthemes.md)
  - [ListTopicRefreshSchedules](api-quicksight-2018-04-01-listtopicrefreshschedules.md)
  - [ListTopicReviewedAnswers](api-quicksight-2018-04-01-listtopicreviewedanswers.md)
  - [ListTopics](api-quicksight-2018-04-01-listtopics.md)
  - [ListUserGroups](api-quicksight-2018-04-01-listusergroups.md)
  - [ListUsers](api-quicksight-2018-04-01-listusers.md)
  - [ListVPCConnections](api-quicksight-2018-04-01-listvpcconnections.md)
  - [PredictQAResults](api-quicksight-2018-04-01-predictqaresults.md)
  - [PutDataSetRefreshProperties](api-quicksight-2018-04-01-putdatasetrefreshproperties.md)
  - [RegisterUser](api-quicksight-2018-04-01-registeruser.md)
  - [RestoreAnalysis](api-quicksight-2018-04-01-restoreanalysis.md)
  - [SearchActionConnectors](api-quicksight-2018-04-01-searchactionconnectors.md)
  - [SearchAnalyses](api-quicksight-2018-04-01-searchanalyses.md)
  - [SearchDashboards](api-quicksight-2018-04-01-searchdashboards.md)
  - [SearchDataSets](api-quicksight-2018-04-01-searchdatasets.md)
  - [SearchDataSources](api-quicksight-2018-04-01-searchdatasources.md)
  - [SearchFlows](api-quicksight-2018-04-01-searchflows.md)
  - [SearchFolders](api-quicksight-2018-04-01-searchfolders.md)
  - [SearchGroups](api-quicksight-2018-04-01-searchgroups.md)
  - [SearchTopics](api-quicksight-2018-04-01-searchtopics.md)
  - [StartAssetBundleExportJob](api-quicksight-2018-04-01-startassetbundleexportjob.md)
  - [StartAssetBundleImportJob](api-quicksight-2018-04-01-startassetbundleimportjob.md)
  - [StartAutomationJob](api-quicksight-2018-04-01-startautomationjob.md)
  - [StartDashboardSnapshotJob](api-quicksight-2018-04-01-startdashboardsnapshotjob.md)
  - [StartDashboardSnapshotJobSchedule](api-quicksight-2018-04-01-startdashboardsnapshotjobschedule.md)
  - [TagResource](api-quicksight-2018-04-01-tagresource.md)
  - [UntagResource](api-quicksight-2018-04-01-untagresource.md)
  - [UpdateAccountCustomPermission](api-quicksight-2018-04-01-updateaccountcustompermission.md)
  - [UpdateAccountCustomization](api-quicksight-2018-04-01-updateaccountcustomization.md)
  - [UpdateAccountSettings](api-quicksight-2018-04-01-updateaccountsettings.md)
  - [UpdateActionConnector](api-quicksight-2018-04-01-updateactionconnector.md)
  - [UpdateActionConnectorPermissions](api-quicksight-2018-04-01-updateactionconnectorpermissions.md)
  - [UpdateAnalysis](api-quicksight-2018-04-01-updateanalysis.md)
  - [UpdateAnalysisPermissions](api-quicksight-2018-04-01-updateanalysispermissions.md)
  - [UpdateApplicationWithTokenExchangeGrant](api-quicksight-2018-04-01-updateapplicationwithtokenexchangegrant.md)
  - [UpdateBrand](api-quicksight-2018-04-01-updatebrand.md)
  - [UpdateBrandAssignment](api-quicksight-2018-04-01-updatebrandassignment.md)
  - [UpdateBrandPublishedVersion](api-quicksight-2018-04-01-updatebrandpublishedversion.md)
  - [UpdateCustomPermissions](api-quicksight-2018-04-01-updatecustompermissions.md)
  - [UpdateDashboard](api-quicksight-2018-04-01-updatedashboard.md)
  - [UpdateDashboardLinks](api-quicksight-2018-04-01-updatedashboardlinks.md)
  - [UpdateDashboardPermissions](api-quicksight-2018-04-01-updatedashboardpermissions.md)
  - [UpdateDashboardPublishedVersion](api-quicksight-2018-04-01-updatedashboardpublishedversion.md)
  - [UpdateDashboardsQAConfiguration](api-quicksight-2018-04-01-updatedashboardsqaconfiguration.md)
  - [UpdateDataSet](api-quicksight-2018-04-01-updatedataset.md)
  - [UpdateDataSetPermissions](api-quicksight-2018-04-01-updatedatasetpermissions.md)
  - [UpdateDataSource](api-quicksight-2018-04-01-updatedatasource.md)
  - [UpdateDataSourcePermissions](api-quicksight-2018-04-01-updatedatasourcepermissions.md)
  - [UpdateDefaultQBusinessApplication](api-quicksight-2018-04-01-updatedefaultqbusinessapplication.md)
  - [UpdateFlowPermissions](api-quicksight-2018-04-01-updateflowpermissions.md)
  - [UpdateFolder](api-quicksight-2018-04-01-updatefolder.md)
  - [UpdateFolderPermissions](api-quicksight-2018-04-01-updatefolderpermissions.md)
  - [UpdateGroup](api-quicksight-2018-04-01-updategroup.md)
  - [UpdateIAMPolicyAssignment](api-quicksight-2018-04-01-updateiampolicyassignment.md)
  - [UpdateIdentityPropagationConfig](api-quicksight-2018-04-01-updateidentitypropagationconfig.md)
  - [UpdateIpRestriction](api-quicksight-2018-04-01-updateiprestriction.md)
  - [UpdateKeyRegistration](api-quicksight-2018-04-01-updatekeyregistration.md)
  - [UpdatePublicSharingSettings](api-quicksight-2018-04-01-updatepublicsharingsettings.md)
  - [UpdateQPersonalizationConfiguration](api-quicksight-2018-04-01-updateqpersonalizationconfiguration.md)
  - [UpdateQuickSightQSearchConfiguration](api-quicksight-2018-04-01-updatequicksightqsearchconfiguration.md)
  - [UpdateRefreshSchedule](api-quicksight-2018-04-01-updaterefreshschedule.md)
  - [UpdateRoleCustomPermission](api-quicksight-2018-04-01-updaterolecustompermission.md)
  - [UpdateSPICECapacityConfiguration](api-quicksight-2018-04-01-updatespicecapacityconfiguration.md)
  - [UpdateSelfUpgrade](api-quicksight-2018-04-01-updateselfupgrade.md)
  - [UpdateSelfUpgradeConfiguration](api-quicksight-2018-04-01-updateselfupgradeconfiguration.md)
  - [UpdateTemplate](api-quicksight-2018-04-01-updatetemplate.md)
  - [UpdateTemplateAlias](api-quicksight-2018-04-01-updatetemplatealias.md)
  - [UpdateTemplatePermissions](api-quicksight-2018-04-01-updatetemplatepermissions.md)
  - [UpdateTheme](api-quicksight-2018-04-01-updatetheme.md)
  - [UpdateThemeAlias](api-quicksight-2018-04-01-updatethemealias.md)
  - [UpdateThemePermissions](api-quicksight-2018-04-01-updatethemepermissions.md)
  - [UpdateTopic](api-quicksight-2018-04-01-updatetopic.md)
  - [UpdateTopicPermissions](api-quicksight-2018-04-01-updatetopicpermissions.md)
  - [UpdateTopicRefreshSchedule](api-quicksight-2018-04-01-updatetopicrefreshschedule.md)
  - [UpdateUser](api-quicksight-2018-04-01-updateuser.md)
  - [UpdateUserCustomPermission](api-quicksight-2018-04-01-updateusercustompermission.md)
  - [UpdateVPCConnection](api-quicksight-2018-04-01-updatevpcconnection.md)

### Table of Contents  [header link](class-aws-quicksight-quicksightclient-toc.md)

#### Methods  [header link](class-aws-quicksight-quicksightclient-toc-methods.md)

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

### Methods  [header link](class-aws-quicksight-quicksightclient-methods.md)

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
  - [Methods](class-aws-quicksight-quicksightclient-toc-methods.md)
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

[Back To Top](class-aws-quicksight-quicksightclient-top.md)
