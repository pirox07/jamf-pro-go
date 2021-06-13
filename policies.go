package jamf_pro_go

import (
	"fmt"
	"net/http"
	"path"
)

const (
	APIVersionPolicies	= "classic"
	APIPathPolices		= "policies"
)

type Policy struct {
	General					PolicyGeneral				`xml:"general"`
	Scope					PolicyScope					`xml:"scope"`
	SelfService				PolicySelfService			`xml:"self_service"`
	PackageConfiguration	PolicyPackageConfiguration	`xml:"package_configuration"`
	PolicyScripts			PolicyScripts				`xml:"scripts"`
	Printers				PolicyPrinters				`xml:"printers"`
	DockItems				PolicyDockItems				`xml:"dock_items"`
	AccountMaintenance		PolicyAccountMaintenance	`xml:"account_maintenance"`
	Maintenance				PolicyMaintenance			`xml:"maintenance"`
	FilesProcesses			PolicyFilesProcesses		`xml:"files_processes"`
	UserInteraction			PolicyUserInteraction		`xml:"user_interaction"`
	DiskEncryption			PolicyDiskEncryption		`xml:"disk_encryption"`
}

type PolicyGeneral struct {
	// Policy ID
	ID                         	uint32					`xml:"id"`
	// Display name for the policy
	Name                       	string					`xml:"name"`
	// Policy Status
	Enabled                    	bool					`xml:"enabled"`
	// Event(s) to use to initiate the policy
	Trigger                    	string					`xml:"trigger"`
	// Trigger: At the recurring check-in frequency configured in Jamf Pro
	TriggerCheckin             	bool					`xml:"trigger_checkin"`
	// Trigger: Immediately after a computer completes the enrollment process
	TriggerEnrollmentComplete  	bool      				`xml:"trigger_enrollment_complete"`
	// Trigger: When a user logs in to a computer.
	//          A login hook that checks for policies must be configured in Jamf Pro for this to work
	TriggerLogin               	bool      				`xml:"trigger_login"`
	// Trigger: When a user logs out of a computer.
	//          A logout hook that checks for policies must be configured in Jamf Pro for this to work
	TriggerLogout              	bool      				`xml:"trigger_logout"`
	// Trigger: When a computer's network state changes
	//          (e.g., when the network connection changes, when the computer name changes, when the IP address changes)
	TriggerNetworkStateChanged	bool      				`xml:"trigger_network_state_changed"`
	// Trigger: When a computer starts up.
	//          A startup script that checks for policies must be configured in Jamf Pro for this to work
	TriggerStartup				bool      				`xml:"trigger_startup"`
	// Trigger: At a custom event
	//          Use to initiate the policy. For an iBeacon region change event, use “beaconStateChange"
	TriggerOther               	string      			`xml:"trigger_other"`
	// which to run the policy
	// [ Once per computer, Once per user per computer, Once per user, Once every day, Once every week, Once every month, Ongoing ]
	Frequency                  	string      			`xml:"frequency"`
	// Make the policy available offline. (This only works with the "Ongoing" execution frequency.)
	Offline                    	bool      				`xml:"offline"`
	// Retry the policy if it fails.
	// [ none, trigger, check-in ], default: none
	RetryEvent                 	string      			`xml:"retry_event"`
	// Configure how many times Jamf Pro attempts to re-run the policy after it fails
	// default: -1
	RetryAttempts              	int32      				`xml:"retry_attempts"`
	// Send notifications for each failed policy retry
	NotifyOnEachFailedRetry    	bool      				`xml:"notify_on_each_failed_retry"`
	// ???
	LocationUserOnly           	string      			`xml:"location_user_only"`
	// Specify the drive on which to run the policy.
	// (e.g. "/Volumes/Restore/"). The policy runs on the boot drive by default
	TargetDrive                	string      			`xml:"target_drive"`
	// Category to add the policy to
	Category                   	PolicyGeneralCategory	`xml:"category"`
	// Specify server-side and client-side limitations for the policy.
	// (For example, you can specify an expiration date/time for the policy,
	//  or ensure that the policy does not run on weekends.)
	DateTimeLimitations 		PolicyGeneralDateTimeLimitations		`xml:"date_time_limitations"`
	// ???
	NetworkLimitations 			PolicyGeneralNetworkLimitations			`xml:"network_limitations"`
	// ???
	OverrideDefaultSettings		PolicyGeneralOverrideDefaultSettings	`xml:"override_default_settings"`
	// Network connection to require to run the policy
	// [ Any, Ethernet ]
	NetworkRequirements			string					`xml:"network_requirements"`
	// ???
	Site						PolicyGeneralSite		`xml:"site"`
}

type PolicyGeneralCategory struct {
	// Category ID, default: -1
	ID		uint32	`xml:"id"`
	// Name of the category,  default: Unknown
	Name	string	`xml:"name"`
}

type PolicyGeneralDateTimeLimitations struct {
	// Server-side limitations: Date/time to make the policy active
	// example: OrderedMap {}
	ActivationDate      string	`xml:"activation_date"`
	// Server-side limitations: Date/time to make the policy active
	// example: 1499470624555 (Unix time)
	ActivationDateEpoch uint64	`xml:"activation_date_epoch"`
	// Server-side limitations: Date/time to make the policy active
	// example: 2017-07-07T18:37:04.555-0500 (UTC)
	ActivationDateUtc   string	`xml:"activation_date_utc"`
	// Server-side limitations: Date/time to make the policy expire
	// example: OrderedMap {}
	ExpirationDate      string 	`xml:"expiration_date"`
	// Server-side limitations: Date/time to make the policy active
	// example: 1499470624555 (Unix time)
	ExpirationDateEpoch uint64	`xml:"expiration_date_epoch"`
	// Server-side limitations: Date/time to make the policy active
	// example: 2017-07-07T18:37:04.555-0500 (UTC)
	ExpirationDateUtc   string 	`xml:"expiration_date_utc"`
	// Client-side limitations:  Days on which the policy should not run
	// [ Sun, Mon, Tue, Wed, Thu, Fri, Sat ]
	NoExecuteOn         string 	`xml:"no_execute_on"`
	// Client-side limitations:  Days on which the policy should not run
	// example: 2:00 AM
	NoExecuteStart      string 	`xml:"no_execute_start"`
	// Client-side limitations:  Days on which the policy should not run
	// example: 4:00 AM
	NoExecuteEnd        string 	`xml:"no_execute_end"`
}

type PolicyGeneralNetworkLimitations struct {
	// ???
	// [ No Minimum, Ethernet ]
	MinimumNetworkConnection	string	`xml:"minimum_network_connection"`
	// ???
	AnyIPAddress             	bool	`xml:"any_ip_address"`
}

// ???
type PolicyGeneralOverrideDefaultSettings struct {
	TargetDrive       	string	`xml:"target_drive"`
	DistributionPoint	string	`xml:"distribution_point"`
	ForceAfpSmb			bool	`xml:"force_afp_smb"`
	Sus					string	`xml:"sus"`
	NetbootServer		string	`xml:"netboot_server"`
}

type PolicyGeneralSite struct {
	// Site ID, default: -1
	ID		int32	`xml:"id"`
	// Name of the site
	Name	string	`xml:"name"`
}

type PolicyScope struct {
	// Deploy the policy to All Computers
	// default: false
	AllComputers	bool						`xml:"all_computers"`
	// Deployment Targets: Specific Computers
	Computers		PolicyScopeComputers		`xml:"computers"`
	// Deployment Targets: Specific Computer Groups
	ComputerGroups	PolicyScopeComputerGroups	`xml:"computer_groups"`
	// Deployment Targets: Computers of Specific Buildings
	Buildings		PolicyScopeBuildings		`xml:"buildings"`
	// Deployment Targets: Computers of Specific Departments
	Departments		PolicyScopeDepartments		`xml:"departments"`
	// Limit the task to specific users in the target.
	LimitToUsers	PolicyScopeLimitToUsers		`xml:"limit_to_users"`
	// limitations to the scope of a remote management task
	Limitations		PolicyScopeLimitations		`xml:"limitations"`
	// Exclude applicable targets from scope
	Exclusions		PolicyScopeExclusions		`xml:"exclusions"`
}

type PolicyScopeComputers struct {
	Computer	[]PolicyScopeComputer	`xml:"computer"`
}

type PolicyScopeComputer struct {
	ID		uint32	`xml:"id"`
	Name	string	`xml:"name"`
	udid	string	`xml:"udid"`
}

type PolicyScopeComputerGroups struct {
	ComputerGroup []PolicyScopeComputerGroup `xml:"computer_group"`
}

type PolicyScopeComputerGroup struct {
	ID 		uint32	`xml:"id"`
	Name	string	`xml:"name"`
}

type PolicyScopeBuildings struct {
	Building []PolicyScopeBuilding `xml:"id"`
}

type PolicyScopeBuilding struct {
	ID		uint32	`xml:"id"`
	Name	string	`xml:"name"`
}

type PolicyScopeDepartments struct {
	Department	[]PolicyScopeDepartment	`xml:"department"`
}

type PolicyScopeDepartment struct {
	ID		uint32	`xml:"id"`
	Name	string	`xml:"name"`
}

type PolicyScopeLimitToUsers struct {
	UserGroups	PolicyScopeLimitUserGroups	`xml:"user_groups"`
}

type PolicyScopeLimitUserGroups struct {
	UsrGroups	[]PolicyScopeLimitUserGroup	`xml:"usr_groups"`
}

type PolicyScopeLimitUserGroup struct {
	UserGroup	string	`xml:"user_group"`
}

type PolicyScopeLimitations struct {
	Users           PolicyScopeUsers			`xml:"users"`
	UserGroups      PolicyScopeUserGroups		`xml:"user_groups"`
	NetworkSegments PolicyScopeNetworkSegments	`xml:"network_segments"`
	Ibeacons        PolicyScopeIbeacons			`xml:"ibeacons"`
}

type PolicyScopeUsers struct {
	User	[]PolicyScopeUsersUser	`xml:"user"`
}

type PolicyScopeUsersUser struct {
	ID		uint32	`xml:"id"`
	Name	string	`xml:"name"`
}

type PolicyScopeUserGroups struct {
	UserGroup	[]PolicyScopeUserGroupsUserGroup	`xml:"user_group"`
}

type PolicyScopeUserGroupsUserGroup struct {
	ID		uint32	`xml:"id"`
	name	string	`xml:"name"`
}

type PolicyScopeNetworkSegments struct {
	NetworkSegment []PolicyScopeNetworkSegmentsNetworkSegment `xml:"network_segment"`
}

type PolicyScopeNetworkSegmentsNetworkSegment struct {
	ID		uint32	`xml:"id"`
	Name	string	`xml:"name"`
}

type PolicyScopeIbeacons struct {
	Ibeacon	[]PolicyScopeIbeaconsIbeacon	`xml:"ibeacon"`
}

type PolicyScopeIbeaconsIbeacon struct {
	ID		uint32	`xml:"id"`
	Name	string	`xml:"name"`
}

type PolicyScopeExclusions struct {
	Computers       PolicyScopeComputers		`xml:"computers"`
	ComputerGroups  PolicyScopeComputerGroups	`xml:"computer_groups"`
	Buildings       PolicyScopeBuildings		`xml:"buildings"`
	Departments     PolicyScopeDepartments		`xml:"departments"`
	Users           PolicyScopeUsers			`xml:"users"`
	UserGroups      PolicyScopeUserGroups		`xml:"user_groups"`
	NetworkSegments	PolicyScopeNetworkSegments	`xml:"network_segments"`
	Ibeacons        PolicyScopeIbeacons			`xml:"ibeacons"`
}

type PolicySelfService struct {
	// default: false
	UseForSelfService           bool						`xml:"use_for_self_service"`
	SelfServiceDisplayName      string						`xml:"self_service_display_name"`
	InstallButtonText           string						`xml:"install_button_text"`
	ReinstallButtonText         string						`xml:"reinstall_button_text"`
	SelfServiceDescription      string						`xml:"self_service_description"`
	// default: false
	ForceUsersToViewDescription bool						`xml:"force_users_to_view_description"`
	SelfServiceIcon             PolicySelfServiceIcon		`xml:"self_service_icon"`
	// default: false
	FeatureOnMainPage           bool						`xml:"feature_on_main_page"`
	SelfServiceCategories       PolicySelfServiceCategories	`xml:"self_service_categories"`
}

type PolicySelfServiceIcon struct {
	ID		uint32	`xml:"id"`
	Name	string `xml:"name"`
	Url		string `xml:"url"`
}

type PolicySelfServiceCategories struct {
	Category PolicySelfServiceCategory `xml:"category"`
}

type PolicySelfServiceCategory struct {
	ID			uint32	`xml:"id"`
	Name		string	`xml:"name"`
	// default: true
	DisplayIn 	bool	`xml:"display_in"`
	// default: false
	FeatureIn 	bool	`xml:"feature_in"`
}

type PolicyPackageConfiguration struct {
	Packages PolicyPackages `xml:"packages"`
}

type PolicyPackages struct {
	Size	uint32			`xml:"size"`
	Package	[]PolicyPackage `xml:"package"`
}

type PolicyPackage struct {
	ID				uint32	`xml:"id"`
	Name			string	`xml:"name"`
	Action			string	`xml:"action"`
	Fut				bool	`xml:"fut"`
	Feu				bool	`xml:"feu"`
	UpdateAutorun	bool	`xml:"update_autorun"`
}


type PolicyScripts struct {
	Size		uint32       `xml:"size"`
	PolicyScript []PolicyScript `xml:"script"`
}

type PolicyScript struct {
	ID          uint32	`xml:"id"`
	Name        string	`xml:"name"`
	// [ Before, After ]
	Priority    string	`xml:"priority"`
	Parameter4  string	`xml:"parameter4"`
	Parameter5  string	`xml:"parameter5"`
	Parameter6  string	`xml:"parameter6"`
	Parameter7  string	`xml:"parameter7"`
	Parameter8  string	`xml:"parameter8"`
	Parameter9  string	`xml:"parameter9"`
	Parameter10 string	`xml:"parameter10"`
	Parameter11	string	`xml:"parameter11"`
}

type PolicyPrinters struct {
	Size					uint32		`xml:"size"`
	LeaveExistingDefault	string		`xml:"leave_existing_default"`
	Printer					[]PolicyPrinter	`xml:"printer"`
}

type PolicyPrinter struct {
	ID			uint32	`xml:"id"`
	Name		uint32	`xml:"name"`
	// [ install, uninstall ]
	Action		string	`xml:"action"`
	MakeDefault	string	`xml:"make_default"`
}

type PolicyDockItems struct {
	Size		uint32				`xml:"size"`
	DockItem	[]PolicyDockItem	`xml:"dock_item"`
}

type PolicyDockItem struct {
	ID		uint32 `xml:"id"`
	Name	string `xml:"name"`
	// [ Add To Beginning, Add To End, Remove ]
	Action	string `xml:"action"`
}

type PolicyAccountMaintenance struct {
	Accounts				PolicyAccounts					`xml:"accounts"`
	DirectoryBindings		PolicyDirectoryBindings			`xml:"directory_bindings"`
	ManagementAccount		PolicyManagementAccount			`xml:"management_account"`
	OpenFirmwareEfiPassword	PolicyOpenFirmwareEfiPassword	`xml:"open_firmware_efi_password"`
}

type PolicyAccounts struct {
	Size	uint32			`xml:"size"`
	Account []PolicyAccount `xml:"account"`
}

type PolicyAccount struct {
	// [ Create, Reset, Delete, DisableFileVault ]
	Action					string	`xml:"action"`
	UserName				string	`xml:"user_name"`
	RealName				string	`xml:"real_name"`
	Password				string	`xml:"password"`
	ArchiveHomeDirectory 	bool	`xml:"archive_home_directory"`
	ArchiveHomeDirectoryTo	string	`xml:"archive_home_directory_to"`
	Home					string	`xml:"home"`
	Picture					string	`xml:"picture"`
	Admin					bool	`xml:"admin"`
	FileVaultEnabled		bool	`xml:"filevault_enabled"`
}

type PolicyDirectoryBindings struct {
	Size	uint32						`xml:"size"`
	Binding	[]PolicyDirectoryBinding	`xml:"binding"`
}

type PolicyDirectoryBinding struct {
	ID		uint32	`xml:"id"`
	Name	string	`xml:"name"`
}

type PolicyManagementAccount struct {
	// [ specified, random, reset, fileVaultEnable, fileVaultDisable ]
	Action					string `xml:"action"`
	ManagedPassword			string `xml:"managed_password"`
	// Only necessary when utilizing the random action
	ManagedPasswordLength	uint32 `xml:"managed_password_length"`
}

type PolicyOpenFirmwareEfiPassword struct {
	// [ command, none ]
	OfMode		string	`xml:"of_mode"`
	OfPassword	string	`xml:"of_password"`
}

type PolicyMaintenance struct {
	Recon                    bool `xml:"recon"`
	ResetName                bool `xml:"reset_name"`
	InstallAllCachedPackages bool `xml:"install_all_cached_packages"`
	Heal                     bool `xml:"heal"`
	Prebindings              bool `xml:"prebindings"`
	Permissions              bool `xml:"permissions"`
	Byhost                   bool `xml:"byhost"`
	SystemCache              bool `xml:"system_cache"`
	UserCache                bool `xml:"user_cache"`
	Verify                   bool `xml:"verify"`
}

type PolicyFilesProcesses struct {
	SearchByPath			string	`xml:"search_by_path"`
	DeleteFile				bool	`xml:"delete_file"`
	LocateFile           	string	`xml:"locate_file"`
	UpdateLocateDatabase	bool	`xml:"update_locate_database"`
	SpotlightSearch			string	`xml:"spotlight_search"`
	SearchForProcess		string	`xml:"search_for_process"`
	KillProcess				bool	`xml:"kill_process"`
	RunCommand				string	`xml:"run_command"`
}

type PolicyUserInteraction struct {
	MessageStart          string	`xml:"message_start"`
	AllowUsersToDefer     bool      `xml:"allow_users_to_defer"`
	AllowDeferralUntilUtc string	`xml:"allow_deferral_until_utc"`
	AllowDeferralMinutes  uint32	`xml:"allow_deferral_minutes"`
	MessageFinish         string	`xml:"message_finish"`
}

type PolicyDiskEncryption struct {
	// [ apply, remediate ]
	Action									string	`xml:"action"`
	DiskEncryptionConfigurationID			uint32	`xml:"disk_encryption_configuration_id"`
	AuthRestart								bool	`xml:"auth_restart"`
	// [ Individual, Institutional, Individual And Institutional ]
	RemediateKeyType						string	`xml:"remediate_key_type"`
	// disk encryption ID to utilize for remediating institutional recovery key types.
	RemediateDiskEncryptionConfigurationID	uint32	`xml:"remediate_disk_encryption_configuration_id"`
}

type GetPoliciesResult struct {
	Size uint32									`xml:"size"`
	Policy []GetPoliciesResultPolicyOverview 	`xml:"policy"`
}

type GetPoliciesResultPolicyOverview struct {
	ID 		uint32 `xml:"id"`
	Name	string `xml:"name"`
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

