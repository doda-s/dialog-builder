package dialog

type DialogNode struct {
	ID      string   `yaml:"id"`
	Title   string   `yaml:"title"`
	Text    string   `yaml:"text"`
	Options []Option `yaml:"options"`
}