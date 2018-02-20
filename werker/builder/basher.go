package builder

import (
	"fmt"
	"bitbucket.org/level11consulting/ocelot/protos"
	ocelog "bitbucket.org/level11consulting/go-til/log"
	"strings"
)

const DefaultBitbucketURL = "https://x-token-auth:%s@bitbucket.org/%s.git"
const DefaultGithubURL = ""

type Basher struct {
	BbDownloadURL string
	GithubDownloadURL string
}

func (b *Basher) GetBbDownloadURL() string {
	if len(b.BbDownloadURL) > 0 {
		return b.BbDownloadURL
	}
	return DefaultBitbucketURL
}

func (b *Basher) GetGithubDownloadURL() string {
	if len(b.GithubDownloadURL) > 0 {
		return b.GithubDownloadURL
	}
	return DefaultGithubURL
}

func (b *Basher) SetBbDownloadURL(downloadURL string) {
	b.BbDownloadURL = downloadURL
}

func (b *Basher) SetGithubDownloadURL(downloadURL string) {
	b.GithubDownloadURL = downloadURL
}

//DownloadCodebase builds bash commands to be executed for downloading the codebase
func (b *Basher) DownloadCodebase(werk *protos.WerkerTask) []string {
	downloadCode := []string {"/bin/sh", "-c"}
	var downloadCmd string
	switch werk.VcsType {
	case "bitbucket":
		//if download url is not the default, then we assume whoever set it knows exactly what they're doing and no replacements
		if b.GetBbDownloadURL() != DefaultBitbucketURL {
			downloadCmd = fmt.Sprintf("/.ocelot/bb_download.sh %s %s %s", werk.VcsToken, b.GetBbDownloadURL(), werk.CheckoutHash)
		} else {
			downloadCmd = fmt.Sprintf("/.ocelot/bb_download.sh %s %s %s", werk.VcsToken, fmt.Sprintf(b.GetBbDownloadURL(), werk.VcsToken, werk.FullName), werk.CheckoutHash)
		}
	case "github":
		ocelog.Log().Error("not implemented")
	default:
		ocelog.Log().Error("werker VCS type not recognized")
	}

	downloadCode = append(downloadCode, downloadCmd)
	return downloadCode
}

//DownloadSSHKey will using the vault token to try to download the ssh key located at the path + `/ssh`
func (b *Basher) DownloadSSHKey(vaultKey, vaultPath string) []string {
	ocelog.Log().Info(fmt.Sprintf("inside of basher's downloadSSHKEYYYY %s %s", vaultKey, vaultPath))
	return []string{"/bin/sh", "-c", fmt.Sprintf("/.ocelot/get_ssh_key.sh %s %s", vaultKey, vaultPath + "/ssh")}
}

func (b *Basher) WriteMavenSettingsXml(settingsXML string) []string {
	return []string{"/bin/sh", "-c", "/.ocelot/render_mvn.sh " + "'" + settingsXML + "'"}
}

//DownloadTemplateFiles will download template files necessary to build containers from werker
func (b *Basher) DownloadTemplateFiles(werkerPort string) []string {
	//downloadLink := fmt.Sprintf("http://docker.for.mac.localhost:%s/do_things.zip", werkerPort)
	downloadLink := fmt.Sprintf("http://172.17.0.1:%s/do_things.zip", werkerPort)
	return []string{"/bin/sh", "-c", "mkdir /.ocelot && wget " + downloadLink + " && unzip do_things.zip -d /.ocelot && cd /.ocelot && chmod +x * && echo \"Ocelot has finished with downloading templates\" && sleep infinity"}
}

//CDAndRunCmds will cd into the root directory of the codebase and execute commands passed in
func (b *Basher) CDAndRunCmds(cmds []string, commitHash string) []string {
	build := append([]string{"cd /" + commitHash}, cmds...)
	buildAndDeploy := append([]string{"/bin/sh", "-c", strings.Join(build, " && ")})
	return buildAndDeploy
}