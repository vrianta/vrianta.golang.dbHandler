package util

import "os"

/*
Pass the variable you want, and the data you need in alternate if the variable is not set
*
*/
func GetEnvriontmentVar(var_name string, alternatedata string) string {
	data := os.Getenv(var_name)

	if data != "" {
		return data
	}

	return alternatedata
}
