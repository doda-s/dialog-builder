package dialog

type Context struct {
	Name        string       `yaml:"name"`
	StartIn 	string		 `yaml:"start_in"`	
	EndIn		string		 `yaml:"end_in"`
	DialogNodes []DialogNode `yaml:"dialog_nodes"`
}