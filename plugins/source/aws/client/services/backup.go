// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/backup"
)

//go:generate mockgen -package=mocks -destination=../mocks/backup.go -source=backup.go BackupClient
type BackupClient interface {
	DescribeBackupJob(context.Context, *backup.DescribeBackupJobInput, ...func(*backup.Options)) (*backup.DescribeBackupJobOutput, error)
	DescribeBackupVault(context.Context, *backup.DescribeBackupVaultInput, ...func(*backup.Options)) (*backup.DescribeBackupVaultOutput, error)
	DescribeCopyJob(context.Context, *backup.DescribeCopyJobInput, ...func(*backup.Options)) (*backup.DescribeCopyJobOutput, error)
	DescribeFramework(context.Context, *backup.DescribeFrameworkInput, ...func(*backup.Options)) (*backup.DescribeFrameworkOutput, error)
	DescribeGlobalSettings(context.Context, *backup.DescribeGlobalSettingsInput, ...func(*backup.Options)) (*backup.DescribeGlobalSettingsOutput, error)
	DescribeProtectedResource(context.Context, *backup.DescribeProtectedResourceInput, ...func(*backup.Options)) (*backup.DescribeProtectedResourceOutput, error)
	DescribeRecoveryPoint(context.Context, *backup.DescribeRecoveryPointInput, ...func(*backup.Options)) (*backup.DescribeRecoveryPointOutput, error)
	DescribeRegionSettings(context.Context, *backup.DescribeRegionSettingsInput, ...func(*backup.Options)) (*backup.DescribeRegionSettingsOutput, error)
	DescribeReportJob(context.Context, *backup.DescribeReportJobInput, ...func(*backup.Options)) (*backup.DescribeReportJobOutput, error)
	DescribeReportPlan(context.Context, *backup.DescribeReportPlanInput, ...func(*backup.Options)) (*backup.DescribeReportPlanOutput, error)
	DescribeRestoreJob(context.Context, *backup.DescribeRestoreJobInput, ...func(*backup.Options)) (*backup.DescribeRestoreJobOutput, error)
	GetBackupPlan(context.Context, *backup.GetBackupPlanInput, ...func(*backup.Options)) (*backup.GetBackupPlanOutput, error)
	GetBackupPlanFromJSON(context.Context, *backup.GetBackupPlanFromJSONInput, ...func(*backup.Options)) (*backup.GetBackupPlanFromJSONOutput, error)
	GetBackupPlanFromTemplate(context.Context, *backup.GetBackupPlanFromTemplateInput, ...func(*backup.Options)) (*backup.GetBackupPlanFromTemplateOutput, error)
	GetBackupSelection(context.Context, *backup.GetBackupSelectionInput, ...func(*backup.Options)) (*backup.GetBackupSelectionOutput, error)
	GetBackupVaultAccessPolicy(context.Context, *backup.GetBackupVaultAccessPolicyInput, ...func(*backup.Options)) (*backup.GetBackupVaultAccessPolicyOutput, error)
	GetBackupVaultNotifications(context.Context, *backup.GetBackupVaultNotificationsInput, ...func(*backup.Options)) (*backup.GetBackupVaultNotificationsOutput, error)
	GetLegalHold(context.Context, *backup.GetLegalHoldInput, ...func(*backup.Options)) (*backup.GetLegalHoldOutput, error)
	GetRecoveryPointRestoreMetadata(context.Context, *backup.GetRecoveryPointRestoreMetadataInput, ...func(*backup.Options)) (*backup.GetRecoveryPointRestoreMetadataOutput, error)
	GetSupportedResourceTypes(context.Context, *backup.GetSupportedResourceTypesInput, ...func(*backup.Options)) (*backup.GetSupportedResourceTypesOutput, error)
	ListBackupJobSummaries(context.Context, *backup.ListBackupJobSummariesInput, ...func(*backup.Options)) (*backup.ListBackupJobSummariesOutput, error)
	ListBackupJobs(context.Context, *backup.ListBackupJobsInput, ...func(*backup.Options)) (*backup.ListBackupJobsOutput, error)
	ListBackupPlanTemplates(context.Context, *backup.ListBackupPlanTemplatesInput, ...func(*backup.Options)) (*backup.ListBackupPlanTemplatesOutput, error)
	ListBackupPlanVersions(context.Context, *backup.ListBackupPlanVersionsInput, ...func(*backup.Options)) (*backup.ListBackupPlanVersionsOutput, error)
	ListBackupPlans(context.Context, *backup.ListBackupPlansInput, ...func(*backup.Options)) (*backup.ListBackupPlansOutput, error)
	ListBackupSelections(context.Context, *backup.ListBackupSelectionsInput, ...func(*backup.Options)) (*backup.ListBackupSelectionsOutput, error)
	ListBackupVaults(context.Context, *backup.ListBackupVaultsInput, ...func(*backup.Options)) (*backup.ListBackupVaultsOutput, error)
	ListCopyJobSummaries(context.Context, *backup.ListCopyJobSummariesInput, ...func(*backup.Options)) (*backup.ListCopyJobSummariesOutput, error)
	ListCopyJobs(context.Context, *backup.ListCopyJobsInput, ...func(*backup.Options)) (*backup.ListCopyJobsOutput, error)
	ListFrameworks(context.Context, *backup.ListFrameworksInput, ...func(*backup.Options)) (*backup.ListFrameworksOutput, error)
	ListLegalHolds(context.Context, *backup.ListLegalHoldsInput, ...func(*backup.Options)) (*backup.ListLegalHoldsOutput, error)
	ListProtectedResources(context.Context, *backup.ListProtectedResourcesInput, ...func(*backup.Options)) (*backup.ListProtectedResourcesOutput, error)
	ListProtectedResourcesByBackupVault(context.Context, *backup.ListProtectedResourcesByBackupVaultInput, ...func(*backup.Options)) (*backup.ListProtectedResourcesByBackupVaultOutput, error)
	ListRecoveryPointsByBackupVault(context.Context, *backup.ListRecoveryPointsByBackupVaultInput, ...func(*backup.Options)) (*backup.ListRecoveryPointsByBackupVaultOutput, error)
	ListRecoveryPointsByLegalHold(context.Context, *backup.ListRecoveryPointsByLegalHoldInput, ...func(*backup.Options)) (*backup.ListRecoveryPointsByLegalHoldOutput, error)
	ListRecoveryPointsByResource(context.Context, *backup.ListRecoveryPointsByResourceInput, ...func(*backup.Options)) (*backup.ListRecoveryPointsByResourceOutput, error)
	ListReportJobs(context.Context, *backup.ListReportJobsInput, ...func(*backup.Options)) (*backup.ListReportJobsOutput, error)
	ListReportPlans(context.Context, *backup.ListReportPlansInput, ...func(*backup.Options)) (*backup.ListReportPlansOutput, error)
	ListRestoreJobSummaries(context.Context, *backup.ListRestoreJobSummariesInput, ...func(*backup.Options)) (*backup.ListRestoreJobSummariesOutput, error)
	ListRestoreJobs(context.Context, *backup.ListRestoreJobsInput, ...func(*backup.Options)) (*backup.ListRestoreJobsOutput, error)
	ListTags(context.Context, *backup.ListTagsInput, ...func(*backup.Options)) (*backup.ListTagsOutput, error)
}
