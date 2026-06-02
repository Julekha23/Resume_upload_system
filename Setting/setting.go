package settings

import (
	"encoding/json"
	"fmt"
	// "go/constant"
	"io"
	"log"
	"os"
	"path"
	// "path/filepath"
)
var Mysetting *Settings
type settingConf struct{
	Server_port int `json:"server_port"`
	Use_env string `json:"use_env"`
	Mongo_url string `json:"mongo_url"`
	Db_name string `json:"db_name"`
	Worker_count int `json:"worker_count"`
	Job_queue_buffer_size int `json:"job_queue_buffer_size"`
	MaxResumeSizeMB        int64    `json:"max_resume_size_mb"`
	ResumeStoragePath      string `json:"resume_storage_path"`  
}
func ReadConfig()(*settingConf,error){
	conf:=&settingConf{}
	filpath:=path.Join("setting","settings.jsonc")
	f,err:=os.Open(filpath)
	if err!=nil{
		return nil,fmt.Errorf("error opening %s: %v", filpath, err)
    } 
	b,err:=io.ReadAll(f)
	if err!=nil{
		return nil,fmt.Errorf("error reading %s: %v", filpath, err)
	}
	if err=json.Unmarshal(b,conf);err!=nil{
	    return nil,fmt.Errorf("error decoding settings: %v", err)
	}
	return conf,nil
}
func Generate(){
	conf,err:=ReadConfig()
	if err!=nil{
		log.Fatal(err)
	}
	
	Mysetting=&Settings{
		server_port:conf.Server_port,
		use_env:conf.Use_env,
		Mongo_url:conf.Mongo_url,
		db_name:conf.Db_name,
		worker_count:conf.Worker_count,
		buffer_size:conf.Job_queue_buffer_size,
		max_resume:conf.MaxResumeSizeMB,
		resumepath:conf.ResumeStoragePath,
		
	}
}
type Settings struct{
	server_port int
	use_env string
	Mongo_url string
	db_name string
    worker_count int
	buffer_size int
	max_resume int64
	resumepath string
}
func (s *Settings) Get_MongoURL() string {
	return  s.Mongo_url
}
func (s *Settings) Get_Server() int{
	return s.server_port
}
func(s *Settings) Get_use_env() string{
	return s.use_env
}
func(s *Settings) Get_mongo_DB() string{
	return s.db_name
}
func(s *Settings) Get_worker_count() int {
	return s.worker_count
}
func(s *Settings) Get_buffer_size() int{
	return s.buffer_size
}
func(s *Settings) Get_max_size() int64{
	return s.max_resume
}
func(s *Settings) Get_resume_path() string{
	return s.resumepath
}