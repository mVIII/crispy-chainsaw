package git

const(
	apiUrlgithub = "https://api.github.com/repos/"
	apiCredsGithub= "?client_id=a9d95dc5ff7296cbc8f0&client_secret=3d2c055452377731740bde0a3a1e38349e536708"
)
type File struct {
	Path  string
	Value string
}
type GitClient interface {
	RepoStructure()[]File
	//GetFile(string)string
}
