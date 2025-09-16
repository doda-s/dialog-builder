package dialog

type Context struct {
	Name        string       `yaml:"name"`
	DialogNodes []DialogNode `yaml:"dialog_nodes"`
}