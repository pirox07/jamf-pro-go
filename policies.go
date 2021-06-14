package jamf_pro_go

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"path"
)

const (
	APIVersionPolicies	= "classic"
	APIPathPolices		= "policies"
)

type Policy struct {
	General					*PolicyGeneral				`xml:"general,omitempty"`
	Scope					*PolicyScope				`xml:"scope,omitempty"`
	SelfService				*PolicySelfService			`xml:"self_service,omitempty"`
	PackageConfiguration	*PolicyPackageConfiguration	`xml:"package_configuration,omitempty"`
	Scripts					*PolicyScripts				`xml:"scripts,omitempty"`
	Printers				*PolicyPrinters				`xml:"printers,omitempty"`
	DockItems				*PolicyDockItems			`xml:"dock_items,omitempty"`
	AccountMaintenance		*PolicyAccountMaintenance	`xml:"account_maintenance,omitempty"`
	Maintenance				*PolicyMaintenance			`xml:"maintenance,omitempty"`
	FilesProcesses			*PolicyFilesProcesses		`xml:"files_processes,omitempty"`
	UserInteraction			*PolicyUserInteraction		`xml:"user_interaction,omitempty"`
	DiskEncryption			*PolicyDiskEncryption		`xml:"disk_encryption,omitempty"`
}

type PolicyGeneral struct {
	ID                         	uint32						`xml:"id,omitempty"`
	Name                       	string						`xml:"name,omitempty"`
	Enabled                    	bool						`xml:"enabled,omitempty"`
	Trigger                    	string						`xml:"trigger,omitempty"`
	TriggerCheckin             	bool						`xml:"trigger_checkin,omitempty"`
	TriggerEnrollmentComplete  	bool      					`xml:"trigger_enrollment_complete,omitempty"`
	TriggerLogin               	bool      					`xml:"trigger_login,omitempty"`
	TriggerLogout              	bool      					`xml:"trigger_logout,omitempty"`
	TriggerNetworkStateChanged	bool      					`xml:"trigger_network_state_changed,omitempty"`
	TriggerStartup				bool      					`xml:"trigger_startup,omitempty"`
	TriggerOther               	string      				`xml:"trigger_other,omitempty"`
	Frequency                  	string      				`xml:"frequency,omitempty"`
	Offline                    	bool      					`xml:"offline,omitempty"`
	RetryEvent                 	string      				`xml:"retry_event,omitempty"`
	RetryAttempts              	int32      					`xml:"retry_attempts,omitempty"`
	NotifyOnEachFailedRetry    	bool      					`xml:"notify_on_each_failed_retry,omitempty"`
	LocationUserOnly           	string      				`xml:"location_user_only,omitempty"`
	TargetDrive                	string      				`xml:"target_drive,omitempty"`
	Category                   	*PolicyCategory				`xml:"category,omitempty"`
	DateTimeLimitations 		*PolicyDateTimeLimitations	`xml:"date_time_limitations,omitempty"`
	NetworkLimitations 			*PolicyNetworkLimitations	`xml:"network_limitations,omitempty"`
	OverrideDefaultSettings		*PolicyOverrides			`xml:"override_default_settings,omitempty"`
	NetworkRequirements			string						`xml:"network_requirements,omitempty"`
	Site						*PolicySite					`xml:"site,omitempty"`
}

type PolicyCategory struct {
	ID		int32	`xml:"id,omitempty"`	// default: -1
	Name	string	`xml:"name,omitempty"`	// default: Unknown
}

type PolicyDateTimeLimitations struct {
	ActivationDate      string	`xml:"activation_date,omitempty"`
	ActivationDateEpoch uint64	`xml:"activation_date_epoch,omitempty"`
	ActivationDateUtc   string	`xml:"activation_date_utc,omitempty"`
	ExpirationDate      string 	`xml:"expiration_date,omitempty"`
	ExpirationDateEpoch uint64	`xml:"expiration_date_epoch,omitempty"`
	ExpirationDateUtc   string 	`xml:"expiration_date_utc,omitempty"`
	NoExecuteOn         string 	`xml:"no_execute_on,omitempty"`
	  // Enum: [ Sun, Mon, Tue, Wed, Thu, Fri, Sat ]
	NoExecuteStart      string 	`xml:"no_execute_start,omitempty"`
	NoExecuteEnd        string 	`xml:"no_execute_end,omitempty"`
}

type PolicyNetworkLimitations struct {
	MinimumNetworkConnection	string	`xml:"minimum_network_connection,omitempty"`
	  // Enum: [ No Minimum, Ethernet ]
	AnyIPAddress             	bool	`xml:"any_ip_address,omitempty"`
}

// PolicyOverrideDefaultSettings contains overrides for the policy's default config
type PolicyOverrides struct {
	TargetDrive       	string	`xml:"target_drive,omitempty"`
	DistributionPoint	string	`xml:"distribution_point,omitempty"`
	ForceAfpSmb			bool	`xml:"force_afp_smb,omitempty"`
	Sus					string	`xml:"sus,omitempty"`
	NetbootServer		string	`xml:"netboot_server,omitempty"`
}

type PolicySite struct {
	ID		int32	`xml:"id,omitempty"`  // default: -1
	Name	string	`xml:"name,omitempty"`
}

type PolicyScope struct {
	AllComputers	bool						`xml:"all_computers,omitempty"`
	  // default: false
	Computers		*PolicyScopeComputers		`xml:"computers,omitempty"`
	ComputerGroups	*PolicyScopeComputerGroups	`xml:"computer_groups,omitempty"`
	Buildings		*PolicyScopeBuildings		`xml:"buildings,omitempty"`
	Departments		*PolicyScopeDepartments		`xml:"departments,omitempty"`
	LimitToUsers	*PolicyScopeLimitToUsers		`xml:"limit_to_users,omitempty"`
	Limitations		*PolicyScopeLimitations		`xml:"limitations,omitempty"`
	Exclusions		*PolicyScopeExclusions		`xml:"exclusions,omitempty"`
}

type PolicyScopeComputers struct {
	Computer	[]*PolicyScopeComputer	`xml:"computer,omitempty"`
}

type PolicyScopeComputer struct {
	ID		uint32	`xml:"id,omitempty"`
	Name	string	`xml:"name,omitempty"`
	UDID	string	`xml:"udid,omitempty"`
}

type PolicyScopeComputerGroups struct {
	ComputerGroup []*PolicyScopeComputerGroup `xml:"computer_group,omitempty"`
}

type PolicyScopeComputerGroup struct {
	ID 		uint32	`xml:"id,omitempty"`
	Name	string	`xml:"name,omitempty"`
}

type PolicyScopeBuildings struct {
	Building []*PolicyScopeBuilding `xml:"id,omitempty"`
}

type PolicyScopeBuilding struct {
	ID		uint32	`xml:"id,omitempty"`
	Name	string	`xml:"name,omitempty"`
}

type PolicyScopeDepartments struct {
	Department	[]*PolicyScopeDepartment	`xml:"department,omitempty"`
}

type PolicyScopeDepartment struct {
	ID		uint32	`xml:"id,omitempty"`
	Name	string	`xml:"name,omitempty"`
}

type PolicyScopeLimitToUsers struct {
	UserGroups	*PolicyScopeLimitUserGroups	`xml:"user_groups,omitempty"`
}

type PolicyScopeLimitUserGroups struct {
	UsrGroups	[]*PolicyScopeLimitUserGroup	`xml:"usr_groups,omitempty"`
}

type PolicyScopeLimitUserGroup struct {
	UserGroup	string	`xml:"user_group,omitempty"`
}

type PolicyScopeLimitations struct {
	Users           *PolicyScopeUsers			`xml:"users,omitempty"`
	UserGroups      *PolicyScopeUserGroups		`xml:"user_groups,omitempty"`
	NetworkSegments *PolicyScopeNetworkSegments	`xml:"network_segments,omitempty"`
	Ibeacons        *PolicyScopeIbeacons			`xml:"ibeacons,omitempty"`
}

type PolicyScopeUsers struct {
	User	[]*PolicyScopeUsersUser	`xml:"user,omitempty"`
}

type PolicyScopeUsersUser struct {
	ID		uint32	`xml:"id,omitempty"`
	Name	string	`xml:"name,omitempty"`
}

type PolicyScopeUserGroups struct {
	UserGroup	[]*PolicyScopeUserGroupsUserGroup	`xml:"user_group,omitempty"`
}

type PolicyScopeUserGroupsUserGroup struct {
	ID		uint32	`xml:"id,omitempty"`
	Name	string	`xml:"name,omitempty"`
}

type PolicyScopeNetworkSegments struct {
	NetworkSegment []*PolicyScopeNetworkSegmentsNetworkSegment `xml:"network_segment,omitempty"`
}

type PolicyScopeNetworkSegmentsNetworkSegment struct {
	ID		uint32	`xml:"id,omitempty"`
	Name	string	`xml:"name,omitempty"`
}

type PolicyScopeIbeacons struct {
	Ibeacon	[]*PolicyScopeIbeaconsIbeacon	`xml:"ibeacon,omitempty"`
}

type PolicyScopeIbeaconsIbeacon struct {
	ID		uint32	`xml:"id,omitempty"`
	Name	string	`xml:"name,omitempty"`
}

type PolicyScopeExclusions struct {
	Computers       *PolicyScopeComputers		`xml:"computers,omitempty"`
	ComputerGroups  *PolicyScopeComputerGroups	`xml:"computer_groups,omitempty"`
	Buildings       *PolicyScopeBuildings		`xml:"buildings,omitempty"`
	Departments     *PolicyScopeDepartments		`xml:"departments,omitempty"`
	Users           *PolicyScopeUsers			`xml:"users,omitempty"`
	UserGroups      *PolicyScopeUserGroups		`xml:"user_groups,omitempty"`
	NetworkSegments	*PolicyScopeNetworkSegments	`xml:"network_segments,omitempty"`
	Ibeacons        *PolicyScopeIbeacons			`xml:"ibeacons,omitempty"`
}

type PolicySelfService struct {
	// default: false
	UseForSelfService           bool						`xml:"use_for_self_service,omitempty"`
	SelfServiceDisplayName      string						`xml:"self_service_display_name,omitempty"`
	InstallButtonText           string						`xml:"install_button_text,omitempty"`
	ReinstallButtonText         string						`xml:"reinstall_button_text,omitempty"`
	SelfServiceDescription      string						`xml:"self_service_description,omitempty"`
	// default: false
	ForceUsersToViewDescription bool						`xml:"force_users_to_view_description,omitempty"`
	SelfServiceIcon             *PolicySelfServiceIcon		`xml:"self_service_icon,omitempty"`
	// default: false
	FeatureOnMainPage           bool						`xml:"feature_on_main_page,omitempty"`
	SelfServiceCategories       *PolicySelfServiceCategories	`xml:"self_service_categories,omitempty"`
}

type PolicySelfServiceIcon struct {
	ID		uint32	`xml:"id,omitempty"`
	Name	string `xml:"name,omitempty"`
	Url		string `xml:"url,omitempty"`
}

type PolicySelfServiceCategories struct {
	Category *PolicySelfServiceCategory `xml:"category,omitempty"`
}

type PolicySelfServiceCategory struct {
	ID			uint32	`xml:"id,omitempty"`
	Name		string	`xml:"name,omitempty"`
	// default: true
	DisplayIn 	bool	`xml:"display_in,omitempty"`
	// default: false
	FeatureIn 	bool	`xml:"feature_in,omitempty"`
}

type PolicyPackageConfiguration struct {
	Packages *PolicyPackages `xml:"packages,omitempty"`
}

type PolicyPackages struct {
	Size	uint32			`xml:"size,omitempty"`
	Package	[]*PolicyPackage `xml:"package,omitempty"`
}

type PolicyPackage struct {
	ID				uint32	`xml:"id,omitempty"`
	Name			string	`xml:"name,omitempty"`
	Action			string	`xml:"action,omitempty"`
	Fut				bool	`xml:"fut,omitempty"`
	Feu				bool	`xml:"feu,omitempty"`
	UpdateAutorun	bool	`xml:"update_autorun,omitempty"`
}


type PolicyScripts struct {
	Size		uint32       `xml:"size,omitempty"`
	PolicyScript []*PolicyScript `xml:"script,omitempty"`
}

type PolicyScript struct {
	ID          uint32	`xml:"id,omitempty"`
	Name        string	`xml:"name,omitempty"`
	// [ Before, After ]
	Priority    string	`xml:"priority,omitempty"`
	Parameter4  string	`xml:"parameter4,omitempty"`
	Parameter5  string	`xml:"parameter5,omitempty"`
	Parameter6  string	`xml:"parameter6,omitempty"`
	Parameter7  string	`xml:"parameter7,omitempty"`
	Parameter8  string	`xml:"parameter8,omitempty"`
	Parameter9  string	`xml:"parameter9,omitempty"`
	Parameter10 string	`xml:"parameter10,omitempty"`
	Parameter11	string	`xml:"parameter11,omitempty"`
}

type PolicyPrinters struct {
	Size					uint32		`xml:"size,omitempty"`
	LeaveExistingDefault	string		`xml:"leave_existing_default,omitempty"`
	Printer					[]*PolicyPrinter	`xml:"printer,omitempty"`
}

type PolicyPrinter struct {
	ID			uint32	`xml:"id,omitempty"`
	Name		uint32	`xml:"name,omitempty"`
	// [ install, uninstall ]
	Action		string	`xml:"action,omitempty"`
	MakeDefault	string	`xml:"make_default,omitempty"`
}

type PolicyDockItems struct {
	Size		uint32				`xml:"size,omitempty"`
	DockItem	[]*PolicyDockItem	`xml:"dock_item,omitempty"`
}

type PolicyDockItem struct {
	ID		uint32 `xml:"id,omitempty"`
	Name	string `xml:"name,omitempty"`
	// [ Add To Beginning, Add To End, Remove ]
	Action	string `xml:"action,omitempty"`
}

type PolicyAccountMaintenance struct {
	Accounts				*PolicyAccounts					`xml:"accounts,omitempty"`
	DirectoryBindings		*PolicyDirectoryBindings			`xml:"directory_bindings,omitempty"`
	ManagementAccount		*PolicyManagementAccount			`xml:"management_account,omitempty"`
	OpenFirmwareEfiPassword	*PolicyOpenFirmwareEfiPassword	`xml:"open_firmware_efi_password,omitempty"`
}

type PolicyAccounts struct {
	Size	uint32			`xml:"size,omitempty"`
	Account []*PolicyAccount `xml:"account,omitempty"`
}

type PolicyAccount struct {
	// [ Create, Reset, Delete, DisableFileVault ]
	Action					string	`xml:"action,omitempty"`
	UserName				string	`xml:"user_name,omitempty"`
	RealName				string	`xml:"real_name,omitempty"`
	Password				string	`xml:"password,omitempty"`
	ArchiveHomeDirectory 	bool	`xml:"archive_home_directory,omitempty"`
	ArchiveHomeDirectoryTo	string	`xml:"archive_home_directory_to,omitempty"`
	Home					string	`xml:"home,omitempty"`
	Picture					string	`xml:"picture,omitempty"`
	Admin					bool	`xml:"admin,omitempty"`
	FileVaultEnabled		bool	`xml:"filevault_enabled,omitempty"`
}

type PolicyDirectoryBindings struct {
	Size	uint32						`xml:"size,omitempty"`
	Binding	[]*PolicyDirectoryBinding	`xml:"binding,omitempty"`
}

type PolicyDirectoryBinding struct {
	ID		uint32	`xml:"id,omitempty"`
	Name	string	`xml:"name,omitempty"`
}

type PolicyManagementAccount struct {
	// [ specified, random, reset, fileVaultEnable, fileVaultDisable ]
	Action					string `xml:"action,omitempty"`
	ManagedPassword			string `xml:"managed_password,omitempty"`
	// Only necessary when utilizing the random action
	ManagedPasswordLength	uint32 `xml:"managed_password_length,omitempty"`
}

type PolicyOpenFirmwareEfiPassword struct {
	// [ command, none ]
	OfMode		string	`xml:"of_mode,omitempty"`
	OfPassword	string	`xml:"of_password,omitempty"`
}

type PolicyMaintenance struct {
	Recon                    bool `xml:"recon,omitempty"`
	ResetName                bool `xml:"reset_name,omitempty"`
	InstallAllCachedPackages bool `xml:"install_all_cached_packages,omitempty"`
	Heal                     bool `xml:"heal,omitempty"`
	Prebindings              bool `xml:"prebindings,omitempty"`
	Permissions              bool `xml:"permissions,omitempty"`
	Byhost                   bool `xml:"byhost,omitempty"`
	SystemCache              bool `xml:"system_cache,omitempty"`
	UserCache                bool `xml:"user_cache,omitempty"`
	Verify                   bool `xml:"verify,omitempty"`
}

type PolicyFilesProcesses struct {
	SearchByPath			string	`xml:"search_by_path,omitempty"`
	DeleteFile				bool	`xml:"delete_file,omitempty"`
	LocateFile           	string	`xml:"locate_file,omitempty"`
	UpdateLocateDatabase	bool	`xml:"update_locate_database,omitempty"`
	SpotlightSearch			string	`xml:"spotlight_search,omitempty"`
	SearchForProcess		string	`xml:"search_for_process,omitempty"`
	KillProcess				bool	`xml:"kill_process,omitempty"`
	RunCommand				string	`xml:"run_command,omitempty"`
}

type PolicyUserInteraction struct {
	MessageStart          string	`xml:"message_start,omitempty"`
	AllowUsersToDefer     bool      `xml:"allow_users_to_defer,omitempty"`
	AllowDeferralUntilUtc string	`xml:"allow_deferral_until_utc,omitempty"`
	AllowDeferralMinutes  uint32	`xml:"allow_deferral_minutes,omitempty"`
	MessageFinish         string	`xml:"message_finish,omitempty"`
}

type PolicyDiskEncryption struct {
	// [ apply, remediate ]
	Action									string	`xml:"action,omitempty"`
	DiskEncryptionConfigurationID			uint32	`xml:"disk_encryption_configuration_id,omitempty"`
	AuthRestart								bool	`xml:"auth_restart,omitempty"`
	// [ Individual, Institutional, Individual And Institutional ]
	RemediateKeyType						string	`xml:"remediate_key_type,omitempty"`
	// disk encryption ID to utilize for remediating institutional recovery key types.
	RemediateDiskEncryptionConfigurationID	uint32	`xml:"remediate_disk_encryption_configuration_id,omitempty"`
}

type GetPoliciesResult struct {
	Size uint32									`xml:"size,omitempty"`
	Policy []GetPoliciesResultPolicyOverview 	`xml:"policy,omitempty"`
}

type GetPoliciesResultPolicyOverview struct {
	ID 		uint32 `xml:"id,omitempty"`
	Name	string `xml:"name,omitempty"`
}

func (c *Client) GetPolicies() (*GetPoliciesResult, error) {
	var result GetPoliciesResult

	err := c.call(APIPathPolices, http.MethodGet,
		APIVersionPolicies, nil, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetPolicy(policyID uint32) (*Policy, error) {
	var result Policy

	err := c.call(path.Join(APIPathPolices, "id", fmt.Sprint(policyID)), http.MethodGet,
		APIVersionPolicies, nil, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

type CreatePolicyParams struct {
	XMLName              	xml.Name                  	`xml:"policy,omitempty"`
	General					*PolicyGeneral				`xml:"general,omitempty"`
	Scope					*PolicyScope					`xml:"scope,omitempty"`
	SelfService				*PolicySelfService			`xml:"self_service,omitempty"`
	PackageConfiguration	*PolicyPackageConfiguration	`xml:"package_configuration,omitempty"`
	Scripts					*PolicyScripts				`xml:"scripts,omitempty"`
	Printers				*PolicyPrinters				`xml:"printers,omitempty"`
	DockItems				*PolicyDockItems				`xml:"dock_items,omitempty"`
	AccountMaintenance		*PolicyAccountMaintenance	`xml:"account_maintenance,omitempty"`
	Maintenance				*PolicyMaintenance			`xml:"maintenance,omitempty"`
	FilesProcesses			*PolicyFilesProcesses		`xml:"files_processes,omitempty"`
	UserInteraction			*PolicyUserInteraction		`xml:"user_interaction,omitempty"`
	DiskEncryption			*PolicyDiskEncryption		`xml:"disk_encryption,omitempty"`
}

type CreatePolicyResult struct {
	XMLName	xml.Name	`xml:"policy,omitempty"`
	ID		uint32		`xml:"id,omitempty"`
}

func (c *Client) CreatePolicy (params *CreatePolicyParams) (*CreatePolicyResult, error) {
	var result CreatePolicyResult

	err := c.call(path.Join(APIPathPolices, "id", "0"), http.MethodPost,
		APIVersionPolicies, nil, params, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}