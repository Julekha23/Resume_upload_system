package settings

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path"
)

var Mysetting *Settings

type settingConf struct {
	Server_port           int    `json:"server_port"`
	Use_env               string `json:"use_env"`
	Mongo_url             string `json:"mongo_url"`
	Resume_db             string `json:"resume_db"`
    Reviewer_db           string `json:"reviewer_db"`
	Worker_count          int    `json:"worker_count"`
	Job_queue_buffer_size int    `json:"job_queue_buffer_size"`
	MaxResumeSizeMB       int64  `json:"max_resume_size_mb"`
	ResumeStoragePath     string `json:"resume_storage_path"`
}

func ReadConfig() (*settingConf, error) {
	conf := &settingConf{}
	filpath := path.Join("Setting", "settings.jsonc")
	f, err := os.Open(filpath)
	if err != nil {
		return nil, fmt.Errorf("error opening %s: %v", filpath, err)
	}
	b, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("error reading %s: %v", filpath, err)
	}
	if err = json.Unmarshal(b, conf); err != nil {
		return nil, fmt.Errorf("error decoding settings: %v", err)
	}
	return conf, nil
}

func Generate() {
	conf, err := ReadConfig()
	if err != nil {
		log.Fatal(err)
	}
	Mysetting = &Settings{
		server_port:  conf.Server_port,
		use_env:      conf.Use_env,
		Mongo_url:    conf.Mongo_url,
		resume_db:    conf.Resume_db,
		reviewer_db:  conf.Reviewer_db,
		worker_count: conf.Worker_count,
		buffer_size:  conf.Job_queue_buffer_size,
		max_resume:   conf.MaxResumeSizeMB,
		resumepath:   conf.ResumeStoragePath,
	}
}

type Settings struct {
	server_port  int
	use_env      string
	Mongo_url    string
	resume_db    string
	reviewer_db  string
	worker_count int
	buffer_size  int
	max_resume   int64
	resumepath   string
}

func (s *Settings) Get_MongoURL() string {
	 return s.Mongo_url 
	}
func (s *Settings) Get_Server() int       {
	 return s.server_port
	}
func (s *Settings) Get_use_env() string   { 
	return s.use_env 
    }
func (s *Settings) Get_mongo_DB() string  {
	 return s.resume_db
	}
func(s *Settings) Get_Reviewer_DB() string{
	return s.reviewer_db
} 	
func (s *Settings) Get_worker_count() int {
	 return s.worker_count 
	}
func (s *Settings) Get_buffer_size() int  {
	 return s.buffer_size 
	}
func (s *Settings) Get_max_size() int64   { 
	return s.max_resume 
    }
func (s *Settings) Get_resume_path() string { 
	return s.resumepath 
    }
