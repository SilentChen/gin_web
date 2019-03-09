package libs

import "gin/plugins/ini"

func CheckErr(err error) {
	if nil != err {
		panic(err)
	}
}

func LoadIniFile(path string) *ini.Section {
	File, err := ini.Load(path)

	CheckErr(err)

	section := File.Section("base").Key("section").String()

	ret := File.Section(section)

	return ret
}