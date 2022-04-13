package main


import (
	// "flag"
	"fmt"
	"tdex/lib"
	"os"
	"strings"
)


func main() {
	
	// Table Summary Export
	var vHost string 
	fmt.Printf("Host(localhost): ")
	nHost, _ := fmt.Scanf("%s",&vHost)
	if nHost == 0 {
		vHost = "localhost"
	}

	var vPort string 
	fmt.Printf("Port(3306): ")
	nPort, _ :=fmt.Scanf("%s",&vPort)
	if nPort == 0 {
		vPort = "3306"
	}

	var vUser string 
	fmt.Printf("User(tdex): ")
	nUser, _ :=fmt.Scanf("%s",&vUser)
	if nUser == 0 {
		vUser = "tdex"
	}
	
	var vPass string 
	fmt.Printf("Pass(tdex): ")
	nPass, _ := fmt.Scanf("%s",&vPass)
	if nPass == 0 {
		vPass = "tdex"
	}


	var vDb string 
	var vLike string
	var nTBlist []string
	fmt.Printf("Database: ")
	nDb, _ := fmt.Scanf("%s",&vDb)
	if nDb == 0 {
		panic("[ERROR]Invaild Database.")
	}

	vDblist := strings.Split(vDb,",")
	if len(vDblist) == 1 {
		fmt.Printf("Table name Like Condition: ")
		_, _ = fmt.Scanf("%s",&vLike)
		nTBlist = strings.Split(vLike,",")
	} 

	var vTBCnt int
	if vLike != "" {
		vTBCnt = 1
	} else {
		vTBCnt = 0
	}

	var vPath string 
	fmt.Printf("FilePath: ")
	nPath, _ := fmt.Scanf("%s",&vPath)
	if nPath == 0 {
		vPath, _ = os.Getwd()
	}


	if len(vDblist) == 1 {
		export := lib.TableSummary(vHost,vPort,vUser,vPass,vDb,vTBCnt,nTBlist,vPath)
		if export == 1 {
			fmt.Println(fmt.Sprintf("%s Table Summary Export Success.",vDb))
		}
	} else {
		for _,dbnm := range vDblist {
			export := lib.TableSummary(vHost,vPort,vUser,vPass,dbnm,vTBCnt,nTBlist,vPath)
			if export == 1 {
				fmt.Println(fmt.Sprintf("%s Table Summary Export Success.",dbnm))
			}
		}	
	}
}