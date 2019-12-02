package backend

// getPlugin is used to retrieve a plugin from either a repo or local path.
// Should be able to download from most common sources. (eg: git, http, mercurial)
// See (https://github.com/hashicorp/go-getter#url-format) for more information
// on how to form input
func getPlugin(name, path string) error {

}

func buildPlugin() error {

}

// buildPlugin attempts to compile and store a specified plugin
// // TODO: should clean up after itself if it fails
// func (master *CursorMaster) buildPlugin(pluginID string) error {

// 	repoPath := fmt.Sprintf("%s/%s", master.config.Master.RepoDirectoryPath, pluginID)
// 	binarypath := fmt.Sprintf("%s/%s", master.config.Master.PluginDirectoryPath, pluginID)

// 	buildArgs := []string{"build", "-o", binarypath}

// 	golangBinaryPath, err := exec.LookPath(golangBinaryName)
// 	if err != nil {
// 		return err
// 	}

// 	_, err = executeCmd(golangBinaryPath, buildArgs, nil, repoPath)
// 	if err != nil {
// 		return err
// 	}

// 	return nil

// }
