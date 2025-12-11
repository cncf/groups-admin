package main

import (
	"bufio"
	"flag"
	"fmt"
	"main/groupsclient"
	"os"
	"regexp"
	"strings"
)

// filterSrcUserSubs takes a regular expression in re and returns an array of MemberInfo whose GroupName filed
// matches the regular expression in filter
func filterSrcUserSubs(re string, subs []groupsclient.MemberInfo) (int, []groupsclient.MemberInfo) {
	filteredList := make([]groupsclient.MemberInfo, 0)
	var subsRegExp = regexp.MustCompile(re)
	for _, sub := range subs {
		if subsRegExp.MatchString(sub.GroupName) {
			filteredList = append(filteredList, sub)
		}
	}
	return len(filteredList), filteredList
}

func main() {
	baseUrl := flag.String("baseUrl", "", "base url of the groups.io server")
	emailPtr := flag.String("srcEmail", "", "groups.io email of the source user")
	passwordPtr := flag.String("srcPass", "", "groups.io password of the source user")
	listFilterPtr := flag.String("filter", "", "RegEx to filter the lists of subscriptions that the command will work on")
	cmdPtr := flag.String("cmd", "members", "Can be one of: srcUserSubs, getUser, xferSubs, members, removeMember, or pendMsgs")
	destEmailPtr := flag.String("destEmail", "", "email of user who will acquire your subscriptions and permissions on groups.io")

	flag.Parse()
	client := groupsclient.NewGroupsClient(*baseUrl)
	// Authenticate and get the token
	err := client.Authenticate(*emailPtr, *passwordPtr)
	if err != nil {
		fmt.Printf("main: client.Authenticate %+v", err)
		return
	}

	// Get user data associated with the user that we authorized the groups.io client.
	// For perms transfer this is the "source user", srcUser
	srcUser, err := client.GetAuthenticatedUser()
	if err != nil {
		fmt.Printf("main: Error getting user ID for %s: %v\n", *emailPtr, err)
		return
	}
	fmt.Printf("loggedInUser: %v-%v \n", srcUser.ID, srcUser.FullName)
	switch *cmdPtr {
	case "members":
		// Get the list of members where the existing user has Owner permissions
		srcUsersSubs, subscriptionCount, err := client.GetMemberInfoList()
		if err != nil {
			fmt.Printf("main: Error getting user groups for %s: %v\n", srcUser.FullName, err)
			return
		}

		if subscriptionCount == 0 {
			fmt.Printf("main: %s is not subscribed to any groups!\n", *emailPtr)
			return
		}
		fmt.Println("Writing report to members.psv")
		membershipReport(srcUsersSubs, client)
	case "srcUserSubs":
		// Get the list of subgroups where the existing user has Owner permissions
		srcUsersSubs, subscriptionCount, err := client.GetMemberInfoList()
		if err != nil {
			fmt.Printf("main: Error getting user groups for %s: %v\n", srcUser.FullName, err)
			return
		}

		if subscriptionCount == 0 {
			fmt.Printf("main: %s is not subscribed to any groups!\n", *emailPtr)
			return
		}
		if *listFilterPtr != "" {
			fmt.Printf("main: Getting user groups for %s: filtered by %s\n", srcUser.FullName, *listFilterPtr)
			filteredCount, filteredList := filterSrcUserSubs(*listFilterPtr, srcUsersSubs)
			fullSummaryReport(emailPtr, filteredCount, filteredList)
		} else {
			fullSummaryReport(emailPtr, subscriptionCount, srcUsersSubs)
		}
	case "getUser":
		if len(*destEmailPtr) > 9 {
			targetUser, err := client.SearchMemberDetails(*destEmailPtr)
			if err != nil {
				fmt.Printf("main: switch case %s: SearchMemberDetails(%s) returned %v\n", *cmdPtr, *destEmailPtr, err)
				return
			}
			userReport(targetUser)
		} else {
			fmt.Printf("main: switch case %s: --destEmail not specified.\n", *cmdPtr)
		}

	case "xferSubs":
		srcUsersSubs, subscriptionCount, err := client.GetMemberInfoList()

		if err != nil {
			fmt.Printf("main: xferSubs: Error getting user groups for %s: %v\n", srcUser.FullName, err)
			return
		}

		if subscriptionCount == 0 {
			fmt.Printf("main: xferSubs: %s is not subscribed to any groups!\n", *emailPtr)
			return
		}
		targetUser, err := client.SearchMemberDetails(*destEmailPtr)
		if err != nil {
			fmt.Printf("main: xferSubs: Error running %s: %v\n", *cmdPtr, err)
			return
		}
		if *listFilterPtr != "" {
			fmt.Printf("main: xferSubs: Getting user groups for %s: filtered by %s\n", srcUser.FullName, *listFilterPtr)
			filteredCount, filteredList := filterSrcUserSubs(*listFilterPtr, srcUsersSubs)
			targetUserSubs, err := client.GrantOwnerPermsToGroupMember(*targetUser, filteredList)
			if err != nil {
				fmt.Printf("main: xferSubs: Error granting owner perms from %s to : filtered by %s %+v \n", srcUser.FullName, *targetUser, err)
			}
			fmt.Printf("main: xferSubs: targetUser %s should be an OWNER on %d namely \n %+v\n", *targetUser, filteredCount, targetUserSubs)
		} else {
			targetUserSubs, err := client.GrantOwnerPermsToGroupMember(*targetUser, srcUsersSubs)
			if err != nil {
				fmt.Printf("Error running client.GrantOwnerPermsToGroupMember(%s, %s): %v\n", *targetUser, srcUsersSubs, err)
				return
			}
			fmt.Printf("targetUserSubs %+v\n", targetUserSubs)
		}

	case "pendMsgs":

		pendingMessages, count, err := client.GetPendingMsgList()
		if err != nil {
			fmt.Printf("groups-admin: Error returned by client.GetPendingMsgList(): %v \n", err)
			return
		}
		fmt.Printf("pendMsgs: found %d ON MAIN GROUP\n", count)

		//targetUserSubs, err := client.ReleasePendingEmail(listIds, allowedEmail)
		for i, pendingMessage := range pendingMessages {
			fmt.Printf("pendMsgs: %d, from: %+v, subject: %s\n", i, pendingMessage.Sender, pendingMessage.Subject)
		}
	case "removeMember":
		srcUsersSubs, subscriptionCount, err := client.GetMemberInfoList()

		if err != nil {
			fmt.Printf("main: xferSubs: Error getting user groups for %s: %v\n", srcUser.FullName, err)
			return
		}

		if subscriptionCount == 0 {
			fmt.Printf("main: xferSubs: %s is not subscribed to any groups!\n", *emailPtr)
			return
		}
		targetUser, err := client.SearchMemberDetails(*destEmailPtr)
		if err != nil {
			fmt.Printf("main: xferSubs: Error running %s: %v\n", *cmdPtr, err)
			return
		}
		if *listFilterPtr != "" {
			fmt.Printf("main: xferSubs: Getting user groups for %s: filtered by %s\n", srcUser.FullName, *listFilterPtr)
			_, filteredList := filterSrcUserSubs(*listFilterPtr, srcUsersSubs)
			removedCount, removedGroups, err := client.RemoveFromGroups(*targetUser, filteredList)
			if err != nil {
				fmt.Printf("main: removeMember: partial failure removing %s: %v\n", targetUser.FullName, err)
			}
			fmt.Printf("main: removeMember: %s removed from %d groups (filtered): %s\n", *targetUser, removedCount, strings.Join(removedGroups, ", "))
		} else {
			removedCount, removedGroups, err := client.RemoveFromGroups(*targetUser, srcUsersSubs)
			if err != nil {
				fmt.Printf("main: removeMember: partial failure removing %s: %v\n", targetUser.FullName, err)
			}
			fmt.Printf("main: removeMember: %s removed from %d groups: %s\n", *targetUser, removedCount, strings.Join(removedGroups, ", "))
		}
	default:
		fmt.Printf("main.go: unknown sub command %s\n", *cmdPtr)
	}
}
func membershipReport(subs []groupsclient.MemberInfo, client *groupsclient.GroupsClient) {
	// Open or create members.psv for writing (overwrite file)
	file, err := os.Create("members.psv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create members.psv: %v\n", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for _, sub := range subs {
		fmt.Fprintf(os.Stderr, "%s - %+v\n", sub.GroupName, sub.Perms)
		if sub.Perms.ManageMembers != true {
			fmt.Fprintf(os.Stderr, "%s - you are subscribed but have no ManageMember permission", sub.GroupName)
		} else {
			members, err := client.GetMembers(sub.GroupID)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error getting membership list for %s. Error: %v\n", sub.GroupName, err)
				continue
			}
			for _, member := range members {
				fullName := member.FullName
				if fullName == "" {
					fullName = "name missing"
				}
				// Write pipe-separated line to file
				line := fmt.Sprintf("%s | %s | %s\n", sub.NiceGroupName, fullName, member.Email)
				if _, err := writer.WriteString(line); err != nil {
					fmt.Fprintf(os.Stderr, "Failed to write member line for %s: %v\n", member.Email, err)
				}
			}
		}
	}
}

func userReport(u interface{}) {
	fmt.Printf("User is %+v\n", u)
}

// fullSummaryReport reports on the logged-in user, showing their email, subs count and a list of their subscriptions
func fullSummaryReport(emailPtr *string, subscriptionCount int, loggedInUsersSubs []groupsclient.MemberInfo) {
	fmt.Printf("%s is subscribed to %d groups, they are...\n", *emailPtr, subscriptionCount)
	for _, subscription := range loggedInUsersSubs {
		fmt.Printf("%s, ", subscription.GroupName)
	}
}

// ContinuePrompt asks if user wants to continue
// if user enters y just return
// if user presses enter or n ir N then calls os.Exit(1)
func ContinuePrompt() {
	ok := YesNoPrompt("Do you want to continue? Defaults to No: [y/n] : ", false)
	if ok {
		fmt.Println("Continuing ...")
	} else {
		fmt.Println("Exiting ...")
		os.Exit(1)
	}
}

// YesNoPrompt asks yes/no questions using the label.
func YesNoPrompt(label string, def bool) bool {
	choices := "Y/n"
	if !def {
		choices = "y/N"
	}

	r := bufio.NewReader(os.Stdin)
	var s string

	for {
		fmt.Printf("%s (%s) ", label, choices)
		s, _ = r.ReadString('\n')
		s = strings.TrimSpace(s)
		if s == "" {
			return def
		}
		s = strings.ToLower(s)
		if s == "y" || s == "yes" {
			return true
		}
		if s == "n" || s == "no" {
			return false
		}
	}
}
