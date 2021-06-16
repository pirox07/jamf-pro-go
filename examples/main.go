package main

import (
	"fmt"
	"os"

	jamf "github.com/pirox07/jamf-pro-go"
)

func main() {
	url := os.Getenv("JAMF_BASE_URL")
	userName := os.Getenv("JAMF_USER")
	password := os.Getenv("JAMF_USER_PASSWORD")
	conf, err := jamf.NewConfig(url, userName, password)
	if err !=nil{
		fmt.Println(err.Error())
	}
	client := jamf.NewClient(conf)


	// GetScripts
	var scriptResults []jamf.Script
	var ops jamf.GetScriptsOpts
	//ops.Page = 3
	ops.PageSize = 10
	//ops.Sort = append(ops.Sort,"categoryName:desc", "id:asc")
	//ops.Filter = "categoryName==\"cat_1\""

	scripts, err := client.GetScripts(ops)
	if err !=nil{
		fmt.Println(err.Error())
	}
	scriptResults = scripts.Results
	fmt.Printf("%+v\n", scriptResults)

	/*
	// GetScript
	script, err := client.GetScript(41)
	if err !=nil{
		fmt.Println(err.Error())
	}

	fmt.Printf("%+v", script)

	 */

	/*
	// CreateScript
	createScriptParams := &jamf.CreateScriptParams{
		// Script ID
		//ID: "0",
		Name: "test_script.sh",
		Info: "hoge",
		Notes: "fuga",
		// [ BEFORE, AFTER, AT_REBOOT ]
		Priority: "AFTER",
		CategoryID: "0",
		CategoryName: "",
		// Parameter label Names
		Parameter4: "arg_1",
		Parameter5: "arg_2",
		Parameter6: "arg_3",
		Parameter7: "arg_4",
		Parameter8: "arg_5",
		Parameter9: "arg_6",
		Parameter10: "arg_7",
		Parameter11: "arg_8",
		OsRequirements: "10.15.x",
		ScriptContents: "#!/bin/bash\n\necho \"Trivial script.\"",
	}
	createScriptResult, err := client.CreateScript(*createScriptParams)
	if err !=nil{
		fmt.Println(err.Error())
	}
	fmt.Printf("%+v\n", createScriptResult)

	 */

	/*
	// UpdateScript
	updateScriptID := "54"
	updateScriptParams := &jamf.UpdateScriptParams{
		// Script ID
		//ID: "54",
		Name: "test_script_update.sh.sh",
		Info: "hogehoge",
		Notes: "fugafuga",
		// [ BEFORE, AFTER, AT_REBOOT ]
		Priority: "AT_REBOOT",
		CategoryID: "1",
		//CategoryName: "",
		// Parameter label Names
		Parameter4: "arg_1_update",
		//Parameter5: "arg_2",
		//Parameter6: "arg_3",
		//Parameter7: "arg_4",
		//Parameter8: "arg_5",
		//Parameter9: "arg_6",
		//Parameter10: "arg_7",
		//Parameter11: "arg_8",
		//OsRequirements: "10.15.x",
		ScriptContents: "#!/bin/bash\n\necho \"Script is updated.\"",
	}
	updateScriptResult, err := client.UpdateScript(updateScriptID, *updateScriptParams)
	if err !=nil{
		fmt.Println(err.Error())
	}
	fmt.Printf("%+v\n",updateScriptResult)

	 */

	/*
	// DeleteScript
	err = client.DeleteScript("54")
	if err != nil{
		fmt.Println(err.Error())
	}

	 */

	/*
	// GetPolicies
	getPoliciesResult, err := client.GetPolicies()
	if err != nil{
		fmt.Println(err.Error())
	}
	fmt.Printf("%+v\n", getPoliciesResult)

	 */


	/*
	// GetPolicy
	var policyID uint32 = 44
	getPolicyResult, err := client.GetPolicy(policyID)
	if err != nil{
		fmt.Println(err.Error())
	}
	fmt.Printf("General ---\n%+v\n",getPolicyResult.General)

	 */

	/*
	// CreatePolicy
	createPolicyParams := &jamf.CreatePolicyParams{
		General: &jamf.PolicyGeneral{
			//ID: 0,
			Name: "sample_policy",
			Enabled: false,
		},
		Scripts: &jamf.PolicyScripts{
			PolicyScript: []*jamf.PolicyScript{
				{
					ID: 41,
					Priority: "AFTER",
				},
			},
		},
	}

	createPolicyResult, err := client.CreatePolicy(createPolicyParams)
	if err !=nil{
		fmt.Println(err.Error())
	}
	fmt.Printf("Created, Policy ID: %+v\n", createPolicyResult.ID)
	 */

	/*
	// UpdatePolicy
	var targetPolicyID  uint32
	targetPolicyID = 79
	updatePolicyParams := &jamf.UpdatePolicyParams{
		General: &jamf.PolicyGeneral{
			//ID: targetPolicyID,
			Name: "sample_policy_update",
			Enabled: true,
		},
		Scripts: &jamf.PolicyScripts{
			PolicyScript: []*jamf.PolicyScript{
				{
					ID: 41,
					Priority: "BEFORE",
				},
			},
		},
	}
	updatePolicyResult, err := client.UpdatePolicy(targetPolicyID, updatePolicyParams)
	if err !=nil{
		fmt.Println(err.Error())
	}
	fmt.Printf("Policy ID: %s\n", updatePolicyResult.ID, " is updated.")

	 */

	/*
	// DeletePolicy
	var n uint32 = 79
	err = client.DeletePolicy(n)
	if err != nil{
		fmt.Println(err.Error())
	}

	 */

}

