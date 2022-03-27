package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
	"path/filepath"
)

func AddNewWorkingGroup(groupName string) {
	CreateFile(filepath.Join("./ini/", fmt.Sprint(groupName, ".ini")))
}

type IniHelper struct {
	file *ini.File
	Path string
}

func (receiver *IniHelper) GetSectionToStrings() []string {
	return receiver.file.SectionStrings()[1:]
}
func (receiver IniHelper) GetSection(name string) *ini.Section {
	return receiver.file.Section(name)
}

func (receiver *IniHelper) NewIni() *IniHelper {
	init, initErr := ini.Load(receiver.Path)
	ColdKiller(initErr)
	receiver.file = init
	return receiver
}

func (receiver *IniHelper) Save() *IniHelper {
	saveErr := receiver.file.SaveTo(receiver.Path)
	ColdKiller(saveErr)
	return receiver
}

func (receiver *IniHelper) NewSection(sectionName string) *IniHelper {
	newSectionErr := receiver.file.NewSections(sectionName)
	ColdKiller(newSectionErr)
	return receiver
}

func (receiver *IniHelper) NewKey(sectionName, key, value string) *IniHelper {
	_, err := receiver.file.Section(sectionName).NewKey(key, value)
	ColdKiller(err)
	return receiver
}

func (receiver *IniHelper) GetKey(sectionName, key string) string {
	return receiver.file.Section(sectionName).Key(key).Value()
}

func (receiver *IniHelper) EditKey(sectionName, key, value string) *IniHelper {
	receiver.file.Section(sectionName).DeleteKey(key)
	receiver.NewKey(sectionName, key, value)
	return receiver
}

func (receiver *IniHelper) DeleteKey(sectionName, key string) *IniHelper {
	receiver.file.Section(sectionName).DeleteKey(key)
	return receiver
}
func (receiver IniHelper) DeleteSection(name string) {
	receiver.file.DeleteSection(name)
}

func NewSection(file *ini.File, sectionName, path string) {
	_, err := file.NewSection(sectionName)
	saveerr := file.SaveTo(path)
	ColdKiller(saveerr)
	ColdKiller(err)
}

func DeleteKey(file *ini.File, key, path string) {
	file.DeleteSection(key)
	err := file.SaveTo(path)
	ColdKiller(err)
}

func EditKey(file *ini.File, sectionName, key, value, path string) {
	file.Section(sectionName).DeleteKey(key)
	_, err := file.Section(sectionName).NewKey(key, value)
	ColdKiller(err)
	saveErr := file.SaveTo(path)
	ColdKiller(saveErr)
}

func NewKeyWithIni(file *ini.File, sectionName, key, value string) {
	_, err := file.Section(sectionName).NewKey(key, value)
	ColdKiller(err)
}

type EditWorkSpaceOrGroupFieldConfig struct {
	SectionName string
	Key         string
	Val         string
	Path        string
}

// EditWorkSpaceOrGroupField 修改工作区的任何字段
// 建议只修改一次内容时使用，避免多次io
func EditWorkSpaceOrGroupField(conf EditWorkSpaceOrGroupFieldConfig) {
	defer func() {
		GreenTips("modified successfully")
	}()
	file, readFileErr := ini.Load(conf.Path)
	ColdKiller(readFileErr)
	EditKey(file, conf.SectionName, conf.Key, conf.Val, conf.Path)
}
