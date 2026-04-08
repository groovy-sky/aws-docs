Menu

- [Aws](namespace-aws.md)
- [Connect](namespace-aws-connect.md)

## ConnectClient     extends [AwsClient](class-aws-awsclient.md)   in package    - [Aws](package-aws.md)

This client is used to interact with the **Amazon Connect Service** service.

## Supported API Versions

This class uses a _service description model_ that is associated at
runtime based on the `version` option given when constructing the
client. The `version` option will determine which API operations,
waiters, and paginators are available for a client. Creating a command or a
specific API operation can be done using magic methods (e.g.,
`$client->commandName(/** parameters */)`, or using the
`$client->getCommand` method of the client.

- [**2017-08-08**](api-connect-2017-08-08.md)

  - [ActivateEvaluationForm](api-connect-2017-08-08-activateevaluationform.md)
  - [AssociateAnalyticsDataSet](api-connect-2017-08-08-associateanalyticsdataset.md)
  - [AssociateApprovedOrigin](api-connect-2017-08-08-associateapprovedorigin.md)
  - [AssociateBot](api-connect-2017-08-08-associatebot.md)
  - [AssociateContactWithUser](api-connect-2017-08-08-associatecontactwithuser.md)
  - [AssociateDefaultVocabulary](api-connect-2017-08-08-associatedefaultvocabulary.md)
  - [AssociateEmailAddressAlias](api-connect-2017-08-08-associateemailaddressalias.md)
  - [AssociateFlow](api-connect-2017-08-08-associateflow.md)
  - [AssociateHoursOfOperations](api-connect-2017-08-08-associatehoursofoperations.md)
  - [AssociateInstanceStorageConfig](api-connect-2017-08-08-associateinstancestorageconfig.md)
  - [AssociateLambdaFunction](api-connect-2017-08-08-associatelambdafunction.md)
  - [AssociateLexBot](api-connect-2017-08-08-associatelexbot.md)
  - [AssociatePhoneNumberContactFlow](api-connect-2017-08-08-associatephonenumbercontactflow.md)
  - [AssociateQueueEmailAddresses](api-connect-2017-08-08-associatequeueemailaddresses.md)
  - [AssociateQueueQuickConnects](api-connect-2017-08-08-associatequeuequickconnects.md)
  - [AssociateRoutingProfileQueues](api-connect-2017-08-08-associateroutingprofilequeues.md)
  - [AssociateSecurityKey](api-connect-2017-08-08-associatesecuritykey.md)
  - [AssociateSecurityProfiles](api-connect-2017-08-08-associatesecurityprofiles.md)
  - [AssociateTrafficDistributionGroupUser](api-connect-2017-08-08-associatetrafficdistributiongroupuser.md)
  - [AssociateUserProficiencies](api-connect-2017-08-08-associateuserproficiencies.md)
  - [AssociateWorkspace](api-connect-2017-08-08-associateworkspace.md)
  - [BatchAssociateAnalyticsDataSet](api-connect-2017-08-08-batchassociateanalyticsdataset.md)
  - [BatchCreateDataTableValue](api-connect-2017-08-08-batchcreatedatatablevalue.md)
  - [BatchDeleteDataTableValue](api-connect-2017-08-08-batchdeletedatatablevalue.md)
  - [BatchDescribeDataTableValue](api-connect-2017-08-08-batchdescribedatatablevalue.md)
  - [BatchDisassociateAnalyticsDataSet](api-connect-2017-08-08-batchdisassociateanalyticsdataset.md)
  - [BatchGetAttachedFileMetadata](api-connect-2017-08-08-batchgetattachedfilemetadata.md)
  - [BatchGetFlowAssociation](api-connect-2017-08-08-batchgetflowassociation.md)
  - [BatchPutContact](api-connect-2017-08-08-batchputcontact.md)
  - [BatchUpdateDataTableValue](api-connect-2017-08-08-batchupdatedatatablevalue.md)
  - [ClaimPhoneNumber](api-connect-2017-08-08-claimphonenumber.md)
  - [CompleteAttachedFileUpload](api-connect-2017-08-08-completeattachedfileupload.md)
  - [CreateAgentStatus](api-connect-2017-08-08-createagentstatus.md)
  - [CreateContact](api-connect-2017-08-08-createcontact.md)
  - [CreateContactFlow](api-connect-2017-08-08-createcontactflow.md)
  - [CreateContactFlowModule](api-connect-2017-08-08-createcontactflowmodule.md)
  - [CreateContactFlowModuleAlias](api-connect-2017-08-08-createcontactflowmodulealias.md)
  - [CreateContactFlowModuleVersion](api-connect-2017-08-08-createcontactflowmoduleversion.md)
  - [CreateContactFlowVersion](api-connect-2017-08-08-createcontactflowversion.md)
  - [CreateDataTable](api-connect-2017-08-08-createdatatable.md)
  - [CreateDataTableAttribute](api-connect-2017-08-08-createdatatableattribute.md)
  - [CreateEmailAddress](api-connect-2017-08-08-createemailaddress.md)
  - [CreateEvaluationForm](api-connect-2017-08-08-createevaluationform.md)
  - [CreateHoursOfOperation](api-connect-2017-08-08-createhoursofoperation.md)
  - [CreateHoursOfOperationOverride](api-connect-2017-08-08-createhoursofoperationoverride.md)
  - [CreateInstance](api-connect-2017-08-08-createinstance.md)
  - [CreateIntegrationAssociation](api-connect-2017-08-08-createintegrationassociation.md)
  - [CreateNotification](api-connect-2017-08-08-createnotification.md)
  - [CreateParticipant](api-connect-2017-08-08-createparticipant.md)
  - [CreatePersistentContactAssociation](api-connect-2017-08-08-createpersistentcontactassociation.md)
  - [CreatePredefinedAttribute](api-connect-2017-08-08-createpredefinedattribute.md)
  - [CreatePrompt](api-connect-2017-08-08-createprompt.md)
  - [CreatePushNotificationRegistration](api-connect-2017-08-08-createpushnotificationregistration.md)
  - [CreateQueue](api-connect-2017-08-08-createqueue.md)
  - [CreateQuickConnect](api-connect-2017-08-08-createquickconnect.md)
  - [CreateRoutingProfile](api-connect-2017-08-08-createroutingprofile.md)
  - [CreateRule](api-connect-2017-08-08-createrule.md)
  - [CreateSecurityProfile](api-connect-2017-08-08-createsecurityprofile.md)
  - [CreateTaskTemplate](api-connect-2017-08-08-createtasktemplate.md)
  - [CreateTestCase](api-connect-2017-08-08-createtestcase.md)
  - [CreateTrafficDistributionGroup](api-connect-2017-08-08-createtrafficdistributiongroup.md)
  - [CreateUseCase](api-connect-2017-08-08-createusecase.md)
  - [CreateUser](api-connect-2017-08-08-createuser.md)
  - [CreateUserHierarchyGroup](api-connect-2017-08-08-createuserhierarchygroup.md)
  - [CreateView](api-connect-2017-08-08-createview.md)
  - [CreateViewVersion](api-connect-2017-08-08-createviewversion.md)
  - [CreateVocabulary](api-connect-2017-08-08-createvocabulary.md)
  - [CreateWorkspace](api-connect-2017-08-08-createworkspace.md)
  - [CreateWorkspacePage](api-connect-2017-08-08-createworkspacepage.md)
  - [DeactivateEvaluationForm](api-connect-2017-08-08-deactivateevaluationform.md)
  - [DeleteAttachedFile](api-connect-2017-08-08-deleteattachedfile.md)
  - [DeleteContactEvaluation](api-connect-2017-08-08-deletecontactevaluation.md)
  - [DeleteContactFlow](api-connect-2017-08-08-deletecontactflow.md)
  - [DeleteContactFlowModule](api-connect-2017-08-08-deletecontactflowmodule.md)
  - [DeleteContactFlowModuleAlias](api-connect-2017-08-08-deletecontactflowmodulealias.md)
  - [DeleteContactFlowModuleVersion](api-connect-2017-08-08-deletecontactflowmoduleversion.md)
  - [DeleteContactFlowVersion](api-connect-2017-08-08-deletecontactflowversion.md)
  - [DeleteDataTable](api-connect-2017-08-08-deletedatatable.md)
  - [DeleteDataTableAttribute](api-connect-2017-08-08-deletedatatableattribute.md)
  - [DeleteEmailAddress](api-connect-2017-08-08-deleteemailaddress.md)
  - [DeleteEvaluationForm](api-connect-2017-08-08-deleteevaluationform.md)
  - [DeleteHoursOfOperation](api-connect-2017-08-08-deletehoursofoperation.md)
  - [DeleteHoursOfOperationOverride](api-connect-2017-08-08-deletehoursofoperationoverride.md)
  - [DeleteInstance](api-connect-2017-08-08-deleteinstance.md)
  - [DeleteIntegrationAssociation](api-connect-2017-08-08-deleteintegrationassociation.md)
  - [DeleteNotification](api-connect-2017-08-08-deletenotification.md)
  - [DeletePredefinedAttribute](api-connect-2017-08-08-deletepredefinedattribute.md)
  - [DeletePrompt](api-connect-2017-08-08-deleteprompt.md)
  - [DeletePushNotificationRegistration](api-connect-2017-08-08-deletepushnotificationregistration.md)
  - [DeleteQueue](api-connect-2017-08-08-deletequeue.md)
  - [DeleteQuickConnect](api-connect-2017-08-08-deletequickconnect.md)
  - [DeleteRoutingProfile](api-connect-2017-08-08-deleteroutingprofile.md)
  - [DeleteRule](api-connect-2017-08-08-deleterule.md)
  - [DeleteSecurityProfile](api-connect-2017-08-08-deletesecurityprofile.md)
  - [DeleteTaskTemplate](api-connect-2017-08-08-deletetasktemplate.md)
  - [DeleteTestCase](api-connect-2017-08-08-deletetestcase.md)
  - [DeleteTrafficDistributionGroup](api-connect-2017-08-08-deletetrafficdistributiongroup.md)
  - [DeleteUseCase](api-connect-2017-08-08-deleteusecase.md)
  - [DeleteUser](api-connect-2017-08-08-deleteuser.md)
  - [DeleteUserHierarchyGroup](api-connect-2017-08-08-deleteuserhierarchygroup.md)
  - [DeleteView](api-connect-2017-08-08-deleteview.md)
  - [DeleteViewVersion](api-connect-2017-08-08-deleteviewversion.md)
  - [DeleteVocabulary](api-connect-2017-08-08-deletevocabulary.md)
  - [DeleteWorkspace](api-connect-2017-08-08-deleteworkspace.md)
  - [DeleteWorkspaceMedia](api-connect-2017-08-08-deleteworkspacemedia.md)
  - [DeleteWorkspacePage](api-connect-2017-08-08-deleteworkspacepage.md)
  - [DescribeAgentStatus](api-connect-2017-08-08-describeagentstatus.md)
  - [DescribeAuthenticationProfile](api-connect-2017-08-08-describeauthenticationprofile.md)
  - [DescribeContact](api-connect-2017-08-08-describecontact.md)
  - [DescribeContactEvaluation](api-connect-2017-08-08-describecontactevaluation.md)
  - [DescribeContactFlow](api-connect-2017-08-08-describecontactflow.md)
  - [DescribeContactFlowModule](api-connect-2017-08-08-describecontactflowmodule.md)
  - [DescribeContactFlowModuleAlias](api-connect-2017-08-08-describecontactflowmodulealias.md)
  - [DescribeDataTable](api-connect-2017-08-08-describedatatable.md)
  - [DescribeDataTableAttribute](api-connect-2017-08-08-describedatatableattribute.md)
  - [DescribeEmailAddress](api-connect-2017-08-08-describeemailaddress.md)
  - [DescribeEvaluationForm](api-connect-2017-08-08-describeevaluationform.md)
  - [DescribeHoursOfOperation](api-connect-2017-08-08-describehoursofoperation.md)
  - [DescribeHoursOfOperationOverride](api-connect-2017-08-08-describehoursofoperationoverride.md)
  - [DescribeInstance](api-connect-2017-08-08-describeinstance.md)
  - [DescribeInstanceAttribute](api-connect-2017-08-08-describeinstanceattribute.md)
  - [DescribeInstanceStorageConfig](api-connect-2017-08-08-describeinstancestorageconfig.md)
  - [DescribeNotification](api-connect-2017-08-08-describenotification.md)
  - [DescribePhoneNumber](api-connect-2017-08-08-describephonenumber.md)
  - [DescribePredefinedAttribute](api-connect-2017-08-08-describepredefinedattribute.md)
  - [DescribePrompt](api-connect-2017-08-08-describeprompt.md)
  - [DescribeQueue](api-connect-2017-08-08-describequeue.md)
  - [DescribeQuickConnect](api-connect-2017-08-08-describequickconnect.md)
  - [DescribeRoutingProfile](api-connect-2017-08-08-describeroutingprofile.md)
  - [DescribeRule](api-connect-2017-08-08-describerule.md)
  - [DescribeSecurityProfile](api-connect-2017-08-08-describesecurityprofile.md)
  - [DescribeTestCase](api-connect-2017-08-08-describetestcase.md)
  - [DescribeTrafficDistributionGroup](api-connect-2017-08-08-describetrafficdistributiongroup.md)
  - [DescribeUser](api-connect-2017-08-08-describeuser.md)
  - [DescribeUserHierarchyGroup](api-connect-2017-08-08-describeuserhierarchygroup.md)
  - [DescribeUserHierarchyStructure](api-connect-2017-08-08-describeuserhierarchystructure.md)
  - [DescribeView](api-connect-2017-08-08-describeview.md)
  - [DescribeVocabulary](api-connect-2017-08-08-describevocabulary.md)
  - [DescribeWorkspace](api-connect-2017-08-08-describeworkspace.md)
  - [DisassociateAnalyticsDataSet](api-connect-2017-08-08-disassociateanalyticsdataset.md)
  - [DisassociateApprovedOrigin](api-connect-2017-08-08-disassociateapprovedorigin.md)
  - [DisassociateBot](api-connect-2017-08-08-disassociatebot.md)
  - [DisassociateEmailAddressAlias](api-connect-2017-08-08-disassociateemailaddressalias.md)
  - [DisassociateFlow](api-connect-2017-08-08-disassociateflow.md)
  - [DisassociateHoursOfOperations](api-connect-2017-08-08-disassociatehoursofoperations.md)
  - [DisassociateInstanceStorageConfig](api-connect-2017-08-08-disassociateinstancestorageconfig.md)
  - [DisassociateLambdaFunction](api-connect-2017-08-08-disassociatelambdafunction.md)
  - [DisassociateLexBot](api-connect-2017-08-08-disassociatelexbot.md)
  - [DisassociatePhoneNumberContactFlow](api-connect-2017-08-08-disassociatephonenumbercontactflow.md)
  - [DisassociateQueueEmailAddresses](api-connect-2017-08-08-disassociatequeueemailaddresses.md)
  - [DisassociateQueueQuickConnects](api-connect-2017-08-08-disassociatequeuequickconnects.md)
  - [DisassociateRoutingProfileQueues](api-connect-2017-08-08-disassociateroutingprofilequeues.md)
  - [DisassociateSecurityKey](api-connect-2017-08-08-disassociatesecuritykey.md)
  - [DisassociateSecurityProfiles](api-connect-2017-08-08-disassociatesecurityprofiles.md)
  - [DisassociateTrafficDistributionGroupUser](api-connect-2017-08-08-disassociatetrafficdistributiongroupuser.md)
  - [DisassociateUserProficiencies](api-connect-2017-08-08-disassociateuserproficiencies.md)
  - [DisassociateWorkspace](api-connect-2017-08-08-disassociateworkspace.md)
  - [DismissUserContact](api-connect-2017-08-08-dismissusercontact.md)
  - [EvaluateDataTableValues](api-connect-2017-08-08-evaluatedatatablevalues.md)
  - [GetAttachedFile](api-connect-2017-08-08-getattachedfile.md)
  - [GetContactAttributes](api-connect-2017-08-08-getcontactattributes.md)
  - [GetContactMetrics](api-connect-2017-08-08-getcontactmetrics.md)
  - [GetCurrentMetricData](api-connect-2017-08-08-getcurrentmetricdata.md)
  - [GetCurrentUserData](api-connect-2017-08-08-getcurrentuserdata.md)
  - [GetEffectiveHoursOfOperations](api-connect-2017-08-08-geteffectivehoursofoperations.md)
  - [GetFederationToken](api-connect-2017-08-08-getfederationtoken.md)
  - [GetFlowAssociation](api-connect-2017-08-08-getflowassociation.md)
  - [GetMetricData](api-connect-2017-08-08-getmetricdata.md)
  - [GetMetricDataV2](api-connect-2017-08-08-getmetricdatav2.md)
  - [GetPromptFile](api-connect-2017-08-08-getpromptfile.md)
  - [GetTaskTemplate](api-connect-2017-08-08-gettasktemplate.md)
  - [GetTestCaseExecutionSummary](api-connect-2017-08-08-gettestcaseexecutionsummary.md)
  - [GetTrafficDistribution](api-connect-2017-08-08-gettrafficdistribution.md)
  - [ImportPhoneNumber](api-connect-2017-08-08-importphonenumber.md)
  - [ImportWorkspaceMedia](api-connect-2017-08-08-importworkspacemedia.md)
  - [ListAgentStatuses](api-connect-2017-08-08-listagentstatuses.md)
  - [ListAnalyticsDataAssociations](api-connect-2017-08-08-listanalyticsdataassociations.md)
  - [ListAnalyticsDataLakeDataSets](api-connect-2017-08-08-listanalyticsdatalakedatasets.md)
  - [ListApprovedOrigins](api-connect-2017-08-08-listapprovedorigins.md)
  - [ListAssociatedContacts](api-connect-2017-08-08-listassociatedcontacts.md)
  - [ListAuthenticationProfiles](api-connect-2017-08-08-listauthenticationprofiles.md)
  - [ListBots](api-connect-2017-08-08-listbots.md)
  - [ListChildHoursOfOperations](api-connect-2017-08-08-listchildhoursofoperations.md)
  - [ListContactEvaluations](api-connect-2017-08-08-listcontactevaluations.md)
  - [ListContactFlowModuleAliases](api-connect-2017-08-08-listcontactflowmodulealiases.md)
  - [ListContactFlowModuleVersions](api-connect-2017-08-08-listcontactflowmoduleversions.md)
  - [ListContactFlowModules](api-connect-2017-08-08-listcontactflowmodules.md)
  - [ListContactFlowVersions](api-connect-2017-08-08-listcontactflowversions.md)
  - [ListContactFlows](api-connect-2017-08-08-listcontactflows.md)
  - [ListContactReferences](api-connect-2017-08-08-listcontactreferences.md)
  - [ListDataTableAttributes](api-connect-2017-08-08-listdatatableattributes.md)
  - [ListDataTablePrimaryValues](api-connect-2017-08-08-listdatatableprimaryvalues.md)
  - [ListDataTableValues](api-connect-2017-08-08-listdatatablevalues.md)
  - [ListDataTables](api-connect-2017-08-08-listdatatables.md)
  - [ListDefaultVocabularies](api-connect-2017-08-08-listdefaultvocabularies.md)
  - [ListEntitySecurityProfiles](api-connect-2017-08-08-listentitysecurityprofiles.md)
  - [ListEvaluationFormVersions](api-connect-2017-08-08-listevaluationformversions.md)
  - [ListEvaluationForms](api-connect-2017-08-08-listevaluationforms.md)
  - [ListFlowAssociations](api-connect-2017-08-08-listflowassociations.md)
  - [ListHoursOfOperationOverrides](api-connect-2017-08-08-listhoursofoperationoverrides.md)
  - [ListHoursOfOperations](api-connect-2017-08-08-listhoursofoperations.md)
  - [ListInstanceAttributes](api-connect-2017-08-08-listinstanceattributes.md)
  - [ListInstanceStorageConfigs](api-connect-2017-08-08-listinstancestorageconfigs.md)
  - [ListInstances](api-connect-2017-08-08-listinstances.md)
  - [ListIntegrationAssociations](api-connect-2017-08-08-listintegrationassociations.md)
  - [ListLambdaFunctions](api-connect-2017-08-08-listlambdafunctions.md)
  - [ListLexBots](api-connect-2017-08-08-listlexbots.md)
  - [ListNotifications](api-connect-2017-08-08-listnotifications.md)
  - [ListPhoneNumbers](api-connect-2017-08-08-listphonenumbers.md)
  - [ListPhoneNumbersV2](api-connect-2017-08-08-listphonenumbersv2.md)
  - [ListPredefinedAttributes](api-connect-2017-08-08-listpredefinedattributes.md)
  - [ListPrompts](api-connect-2017-08-08-listprompts.md)
  - [ListQueueEmailAddresses](api-connect-2017-08-08-listqueueemailaddresses.md)
  - [ListQueueQuickConnects](api-connect-2017-08-08-listqueuequickconnects.md)
  - [ListQueues](api-connect-2017-08-08-listqueues.md)
  - [ListQuickConnects](api-connect-2017-08-08-listquickconnects.md)
  - [ListRealtimeContactAnalysisSegmentsV2](api-connect-2017-08-08-listrealtimecontactanalysissegmentsv2.md)
  - [ListRoutingProfileManualAssignmentQueues](api-connect-2017-08-08-listroutingprofilemanualassignmentqueues.md)
  - [ListRoutingProfileQueues](api-connect-2017-08-08-listroutingprofilequeues.md)
  - [ListRoutingProfiles](api-connect-2017-08-08-listroutingprofiles.md)
  - [ListRules](api-connect-2017-08-08-listrules.md)
  - [ListSecurityKeys](api-connect-2017-08-08-listsecuritykeys.md)
  - [ListSecurityProfileApplications](api-connect-2017-08-08-listsecurityprofileapplications.md)
  - [ListSecurityProfileFlowModules](api-connect-2017-08-08-listsecurityprofileflowmodules.md)
  - [ListSecurityProfilePermissions](api-connect-2017-08-08-listsecurityprofilepermissions.md)
  - [ListSecurityProfiles](api-connect-2017-08-08-listsecurityprofiles.md)
  - [ListTagsForResource](api-connect-2017-08-08-listtagsforresource.md)
  - [ListTaskTemplates](api-connect-2017-08-08-listtasktemplates.md)
  - [ListTestCaseExecutionRecords](api-connect-2017-08-08-listtestcaseexecutionrecords.md)
  - [ListTestCaseExecutions](api-connect-2017-08-08-listtestcaseexecutions.md)
  - [ListTestCases](api-connect-2017-08-08-listtestcases.md)
  - [ListTrafficDistributionGroupUsers](api-connect-2017-08-08-listtrafficdistributiongroupusers.md)
  - [ListTrafficDistributionGroups](api-connect-2017-08-08-listtrafficdistributiongroups.md)
  - [ListUseCases](api-connect-2017-08-08-listusecases.md)
  - [ListUserHierarchyGroups](api-connect-2017-08-08-listuserhierarchygroups.md)
  - [ListUserNotifications](api-connect-2017-08-08-listusernotifications.md)
  - [ListUserProficiencies](api-connect-2017-08-08-listuserproficiencies.md)
  - [ListUsers](api-connect-2017-08-08-listusers.md)
  - [ListViewVersions](api-connect-2017-08-08-listviewversions.md)
  - [ListViews](api-connect-2017-08-08-listviews.md)
  - [ListWorkspaceMedia](api-connect-2017-08-08-listworkspacemedia.md)
  - [ListWorkspacePages](api-connect-2017-08-08-listworkspacepages.md)
  - [ListWorkspaces](api-connect-2017-08-08-listworkspaces.md)
  - [MonitorContact](api-connect-2017-08-08-monitorcontact.md)
  - [PauseContact](api-connect-2017-08-08-pausecontact.md)
  - [PutUserStatus](api-connect-2017-08-08-putuserstatus.md)
  - [ReleasePhoneNumber](api-connect-2017-08-08-releasephonenumber.md)
  - [ReplicateInstance](api-connect-2017-08-08-replicateinstance.md)
  - [ResumeContact](api-connect-2017-08-08-resumecontact.md)
  - [ResumeContactRecording](api-connect-2017-08-08-resumecontactrecording.md)
  - [SearchAgentStatuses](api-connect-2017-08-08-searchagentstatuses.md)
  - [SearchAvailablePhoneNumbers](api-connect-2017-08-08-searchavailablephonenumbers.md)
  - [SearchContactEvaluations](api-connect-2017-08-08-searchcontactevaluations.md)
  - [SearchContactFlowModules](api-connect-2017-08-08-searchcontactflowmodules.md)
  - [SearchContactFlows](api-connect-2017-08-08-searchcontactflows.md)
  - [SearchContacts](api-connect-2017-08-08-searchcontacts.md)
  - [SearchDataTables](api-connect-2017-08-08-searchdatatables.md)
  - [SearchEmailAddresses](api-connect-2017-08-08-searchemailaddresses.md)
  - [SearchEvaluationForms](api-connect-2017-08-08-searchevaluationforms.md)
  - [SearchHoursOfOperationOverrides](api-connect-2017-08-08-searchhoursofoperationoverrides.md)
  - [SearchHoursOfOperations](api-connect-2017-08-08-searchhoursofoperations.md)
  - [SearchNotifications](api-connect-2017-08-08-searchnotifications.md)
  - [SearchPredefinedAttributes](api-connect-2017-08-08-searchpredefinedattributes.md)
  - [SearchPrompts](api-connect-2017-08-08-searchprompts.md)
  - [SearchQueues](api-connect-2017-08-08-searchqueues.md)
  - [SearchQuickConnects](api-connect-2017-08-08-searchquickconnects.md)
  - [SearchResourceTags](api-connect-2017-08-08-searchresourcetags.md)
  - [SearchRoutingProfiles](api-connect-2017-08-08-searchroutingprofiles.md)
  - [SearchSecurityProfiles](api-connect-2017-08-08-searchsecurityprofiles.md)
  - [SearchTestCases](api-connect-2017-08-08-searchtestcases.md)
  - [SearchUserHierarchyGroups](api-connect-2017-08-08-searchuserhierarchygroups.md)
  - [SearchUsers](api-connect-2017-08-08-searchusers.md)
  - [SearchViews](api-connect-2017-08-08-searchviews.md)
  - [SearchVocabularies](api-connect-2017-08-08-searchvocabularies.md)
  - [SearchWorkspaceAssociations](api-connect-2017-08-08-searchworkspaceassociations.md)
  - [SearchWorkspaces](api-connect-2017-08-08-searchworkspaces.md)
  - [SendChatIntegrationEvent](api-connect-2017-08-08-sendchatintegrationevent.md)
  - [SendOutboundEmail](api-connect-2017-08-08-sendoutboundemail.md)
  - [StartAttachedFileUpload](api-connect-2017-08-08-startattachedfileupload.md)
  - [StartChatContact](api-connect-2017-08-08-startchatcontact.md)
  - [StartContactEvaluation](api-connect-2017-08-08-startcontactevaluation.md)
  - [StartContactMediaProcessing](api-connect-2017-08-08-startcontactmediaprocessing.md)
  - [StartContactRecording](api-connect-2017-08-08-startcontactrecording.md)
  - [StartContactStreaming](api-connect-2017-08-08-startcontactstreaming.md)
  - [StartEmailContact](api-connect-2017-08-08-startemailcontact.md)
  - [StartOutboundChatContact](api-connect-2017-08-08-startoutboundchatcontact.md)
  - [StartOutboundEmailContact](api-connect-2017-08-08-startoutboundemailcontact.md)
  - [StartOutboundVoiceContact](api-connect-2017-08-08-startoutboundvoicecontact.md)
  - [StartScreenSharing](api-connect-2017-08-08-startscreensharing.md)
  - [StartTaskContact](api-connect-2017-08-08-starttaskcontact.md)
  - [StartTestCaseExecution](api-connect-2017-08-08-starttestcaseexecution.md)
  - [StartWebRTCContact](api-connect-2017-08-08-startwebrtccontact.md)
  - [StopContact](api-connect-2017-08-08-stopcontact.md)
  - [StopContactMediaProcessing](api-connect-2017-08-08-stopcontactmediaprocessing.md)
  - [StopContactRecording](api-connect-2017-08-08-stopcontactrecording.md)
  - [StopContactStreaming](api-connect-2017-08-08-stopcontactstreaming.md)
  - [StopTestCaseExecution](api-connect-2017-08-08-stoptestcaseexecution.md)
  - [SubmitContactEvaluation](api-connect-2017-08-08-submitcontactevaluation.md)
  - [SuspendContactRecording](api-connect-2017-08-08-suspendcontactrecording.md)
  - [TagContact](api-connect-2017-08-08-tagcontact.md)
  - [TagResource](api-connect-2017-08-08-tagresource.md)
  - [TransferContact](api-connect-2017-08-08-transfercontact.md)
  - [UntagContact](api-connect-2017-08-08-untagcontact.md)
  - [UntagResource](api-connect-2017-08-08-untagresource.md)
  - [UpdateAgentStatus](api-connect-2017-08-08-updateagentstatus.md)
  - [UpdateAuthenticationProfile](api-connect-2017-08-08-updateauthenticationprofile.md)
  - [UpdateContact](api-connect-2017-08-08-updatecontact.md)
  - [UpdateContactAttributes](api-connect-2017-08-08-updatecontactattributes.md)
  - [UpdateContactEvaluation](api-connect-2017-08-08-updatecontactevaluation.md)
  - [UpdateContactFlowContent](api-connect-2017-08-08-updatecontactflowcontent.md)
  - [UpdateContactFlowMetadata](api-connect-2017-08-08-updatecontactflowmetadata.md)
  - [UpdateContactFlowModuleAlias](api-connect-2017-08-08-updatecontactflowmodulealias.md)
  - [UpdateContactFlowModuleContent](api-connect-2017-08-08-updatecontactflowmodulecontent.md)
  - [UpdateContactFlowModuleMetadata](api-connect-2017-08-08-updatecontactflowmodulemetadata.md)
  - [UpdateContactFlowName](api-connect-2017-08-08-updatecontactflowname.md)
  - [UpdateContactRoutingData](api-connect-2017-08-08-updatecontactroutingdata.md)
  - [UpdateContactSchedule](api-connect-2017-08-08-updatecontactschedule.md)
  - [UpdateDataTableAttribute](api-connect-2017-08-08-updatedatatableattribute.md)
  - [UpdateDataTableMetadata](api-connect-2017-08-08-updatedatatablemetadata.md)
  - [UpdateDataTablePrimaryValues](api-connect-2017-08-08-updatedatatableprimaryvalues.md)
  - [UpdateEmailAddressMetadata](api-connect-2017-08-08-updateemailaddressmetadata.md)
  - [UpdateEvaluationForm](api-connect-2017-08-08-updateevaluationform.md)
  - [UpdateHoursOfOperation](api-connect-2017-08-08-updatehoursofoperation.md)
  - [UpdateHoursOfOperationOverride](api-connect-2017-08-08-updatehoursofoperationoverride.md)
  - [UpdateInstanceAttribute](api-connect-2017-08-08-updateinstanceattribute.md)
  - [UpdateInstanceStorageConfig](api-connect-2017-08-08-updateinstancestorageconfig.md)
  - [UpdateNotificationContent](api-connect-2017-08-08-updatenotificationcontent.md)
  - [UpdateParticipantAuthentication](api-connect-2017-08-08-updateparticipantauthentication.md)
  - [UpdateParticipantRoleConfig](api-connect-2017-08-08-updateparticipantroleconfig.md)
  - [UpdatePhoneNumber](api-connect-2017-08-08-updatephonenumber.md)
  - [UpdatePhoneNumberMetadata](api-connect-2017-08-08-updatephonenumbermetadata.md)
  - [UpdatePredefinedAttribute](api-connect-2017-08-08-updatepredefinedattribute.md)
  - [UpdatePrompt](api-connect-2017-08-08-updateprompt.md)
  - [UpdateQueueHoursOfOperation](api-connect-2017-08-08-updatequeuehoursofoperation.md)
  - [UpdateQueueMaxContacts](api-connect-2017-08-08-updatequeuemaxcontacts.md)
  - [UpdateQueueName](api-connect-2017-08-08-updatequeuename.md)
  - [UpdateQueueOutboundCallerConfig](api-connect-2017-08-08-updatequeueoutboundcallerconfig.md)
  - [UpdateQueueOutboundEmailConfig](api-connect-2017-08-08-updatequeueoutboundemailconfig.md)
  - [UpdateQueueStatus](api-connect-2017-08-08-updatequeuestatus.md)
  - [UpdateQuickConnectConfig](api-connect-2017-08-08-updatequickconnectconfig.md)
  - [UpdateQuickConnectName](api-connect-2017-08-08-updatequickconnectname.md)
  - [UpdateRoutingProfileAgentAvailabilityTimer](api-connect-2017-08-08-updateroutingprofileagentavailabilitytimer.md)
  - [UpdateRoutingProfileConcurrency](api-connect-2017-08-08-updateroutingprofileconcurrency.md)
  - [UpdateRoutingProfileDefaultOutboundQueue](api-connect-2017-08-08-updateroutingprofiledefaultoutboundqueue.md)
  - [UpdateRoutingProfileName](api-connect-2017-08-08-updateroutingprofilename.md)
  - [UpdateRoutingProfileQueues](api-connect-2017-08-08-updateroutingprofilequeues.md)
  - [UpdateRule](api-connect-2017-08-08-updaterule.md)
  - [UpdateSecurityProfile](api-connect-2017-08-08-updatesecurityprofile.md)
  - [UpdateTaskTemplate](api-connect-2017-08-08-updatetasktemplate.md)
  - [UpdateTestCase](api-connect-2017-08-08-updatetestcase.md)
  - [UpdateTrafficDistribution](api-connect-2017-08-08-updatetrafficdistribution.md)
  - [UpdateUserConfig](api-connect-2017-08-08-updateuserconfig.md)
  - [UpdateUserHierarchy](api-connect-2017-08-08-updateuserhierarchy.md)
  - [UpdateUserHierarchyGroupName](api-connect-2017-08-08-updateuserhierarchygroupname.md)
  - [UpdateUserHierarchyStructure](api-connect-2017-08-08-updateuserhierarchystructure.md)
  - [UpdateUserIdentityInfo](api-connect-2017-08-08-updateuseridentityinfo.md)
  - [UpdateUserNotificationStatus](api-connect-2017-08-08-updateusernotificationstatus.md)
  - [UpdateUserPhoneConfig](api-connect-2017-08-08-updateuserphoneconfig.md)
  - [UpdateUserProficiencies](api-connect-2017-08-08-updateuserproficiencies.md)
  - [UpdateUserRoutingProfile](api-connect-2017-08-08-updateuserroutingprofile.md)
  - [UpdateUserSecurityProfiles](api-connect-2017-08-08-updateusersecurityprofiles.md)
  - [UpdateViewContent](api-connect-2017-08-08-updateviewcontent.md)
  - [UpdateViewMetadata](api-connect-2017-08-08-updateviewmetadata.md)
  - [UpdateWorkspaceMetadata](api-connect-2017-08-08-updateworkspacemetadata.md)
  - [UpdateWorkspacePage](api-connect-2017-08-08-updateworkspacepage.md)
  - [UpdateWorkspaceTheme](api-connect-2017-08-08-updateworkspacetheme.md)
  - [UpdateWorkspaceVisibility](api-connect-2017-08-08-updateworkspacevisibility.md)

### Table of Contents  [header link](class-aws-connect-connectclient-toc.md)

#### Methods  [header link](class-aws-connect-connectclient-toc-methods.md)

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

### Methods  [header link](class-aws-connect-connectclient-methods.md)

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
  - [Methods](class-aws-connect-connectclient-toc-methods.md)
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

[Back To Top](class-aws-connect-connectclient-top.md)
